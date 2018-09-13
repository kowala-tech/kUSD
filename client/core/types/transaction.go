package types

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"sync/atomic"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/rlp"
)

//go:generate gencodec -type txdata -field-override txdataMarshaling -out gen_tx_json.go

// deriveSigner makes a *best* guess about which signer to use.
func deriveSigner(V *big.Int) Signer {
	return NewAndromedaSigner(deriveChainID(V))
}

type Transaction struct {
	data txdata
	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type txdata struct {
	AccountNonce uint64          `json:"nonce"        gencodec:"required"`
	ComputeLimit uint64          `json:"computeLimit" gencodec:"required"`
	Recipient    *common.Address `json:"to"           rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"        gencodec:"required"`
	Payload      []byte          `json:"input"        gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

type txdataMarshaling struct {
	AccountNonce hexutil.Uint64
	ComputeLimit hexutil.Uint64
	Amount       *hexutil.Big
	Payload      hexutil.Bytes
	V            *hexutil.Big
	R            *hexutil.Big
	S            *hexutil.Big
}

func NewTransaction(nonce uint64, to common.Address, amount *big.Int, computeLimit uint64, data []byte) *Transaction {
	return newTransaction(nonce, &to, amount, computeLimit, data)
}

func NewContractCreation(nonce uint64, amount *big.Int, computeLimit uint64, data []byte) *Transaction {
	return newTransaction(nonce, nil, amount, computeLimit, data)
}

func newTransaction(nonce uint64, to *common.Address, amount *big.Int, computeLimit uint64, data []byte) *Transaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := txdata{
		AccountNonce: nonce,
		Recipient:    to,
		Payload:      data,
		Amount:       new(big.Int),
		ComputeLimit: computeLimit,
		V:            new(big.Int),
		R:            new(big.Int),
		S:            new(big.Int),
	}
	if amount != nil {
		d.Amount.Set(amount)
	}

	return &Transaction{data: d}
}

// ChainID returns which chain id this transaction was signed for (if at all)
func (tx *Transaction) ChainID() *big.Int {
	return deriveChainID(tx.data.V)
}
func (tx *Transaction) SignatureValues() (R, S, V *big.Int) {
	R, S, V = tx.data.R, tx.data.S, tx.data.V
	return
}

// EncodeRLP implements rlp.Encoder
func (tx *Transaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.data)
}

// DecodeRLP implements rlp.Decoder
func (tx *Transaction) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&tx.data)
	if err == nil {
		tx.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

// MarshalJSON encodes the web3 RPC transaction format.
func (tx *Transaction) MarshalJSON() ([]byte, error) {
	hash := tx.Hash()
	data := tx.data
	data.Hash = &hash
	return data.MarshalJSON()
}

// UnmarshalJSON decodes the web3 RPC transaction format.
func (tx *Transaction) UnmarshalJSON(input []byte) error {
	var dec txdata
	if err := dec.UnmarshalJSON(input); err != nil {
		return err
	}
	var V byte
	if isProtectedV(dec.V) {
		chainID := deriveChainID(dec.V).Uint64()
		V = byte(dec.V.Uint64() - 35 - 2*chainID)
	} else {
		V = byte(dec.V.Uint64() - 27)
	}
	if !crypto.ValidateSignatureValues(V, dec.R, dec.S, false) {
		return ErrInvalidSig
	}
	*tx = Transaction{data: dec}
	return nil
}

func (tx *Transaction) Data() []byte         { return common.CopyBytes(tx.data.Payload) }
func (tx *Transaction) ComputeLimit() uint64 { return tx.data.ComputeLimit }
func (tx *Transaction) Value() *big.Int      { return new(big.Int).Set(tx.data.Amount) }
func (tx *Transaction) Nonce() uint64        { return tx.data.AccountNonce }
func (tx *Transaction) CheckNonce() bool     { return true }

func (tx *Transaction) From() (*common.Address, error) {
	if tx.data.V == nil {
		return nil, errors.New("[invalid sender: nil V field]")
	}

	signer := deriveSigner(tx.data.V)
	f, err := TxSender(signer, tx)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

// To returns the recipient address of the transaction.
// It returns nil if the transaction is a contract creation.
func (tx *Transaction) To() *common.Address {
	if tx.data.Recipient == nil {
		return nil
	}
	to := *tx.data.Recipient
	return &to
}

// Hash hashes the RLP encoding of tx.
// It uniquely identifies the transaction.
func (tx *Transaction) Hash() common.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(tx)
	tx.hash.Store(v)
	return v
}

// ProtectedHash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (tx *Transaction) ProtectedHash(chainID *big.Int) common.Hash {
	return tx.HashWithData(chainID, uint(0), uint(0))
}

func (tx *Transaction) UnprotectedHash() common.Hash {
	return tx.HashWithData()
}

func (tx *Transaction) HashWithData(data ...interface{}) common.Hash {
	txData := []interface{}{
		tx.data.AccountNonce,
		tx.data.ComputeLimit,
		tx.data.Recipient,
		tx.data.Amount,
		tx.data.Payload,
	}
	return rlpHash(append(txData, data...))
}

// Protected returns whether the transaction is protected from replay protection.
func (tx *Transaction) Protected() bool {
	return isProtectedV(tx.data.V)
}

func isProtectedV(V *big.Int) bool {
	if V.BitLen() <= 8 {
		v := V.Uint64()
		return v != 27 && v != 28
	}
	// anything not 27 or 28 are considered unprotected
	return true
}

// Size returns the true RLP encoded storage size of the transaction, either by
// encoding and returning it, or returning a previsouly cached value.
func (tx *Transaction) Size() common.StorageSize {
	if size := tx.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, &tx.data)
	tx.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

// AsMessage returns the transaction as a core.Message.
//
// AsMessage requires a signer to derive the sender.
//
// XXX Rename message to something less arbitrary?
func (tx *Transaction) AsMessage(s Signer) (Message, error) {
	msg := Message{
		nonce:        tx.data.AccountNonce,
		computeLimit: tx.data.ComputeLimit,
		to:           tx.data.Recipient,
		amount:       tx.data.Amount,
		data:         tx.data.Payload,
		checkNonce:   true,
	}

	var err error
	msg.from, err = TxSender(s, tx)
	return msg, err
}

// WithSignature returns a new transaction with the given signature.
// This signature needs to be formatted as described in the yellow paper (v+27).
func (tx *Transaction) WithSignature(signer Signer, sig []byte) (*Transaction, error) {
	r, s, v, err := signer.SignatureValues(sig)
	if err != nil {
		return nil, err
	}
	cpy := &Transaction{data: tx.data}
	cpy.data.R, cpy.data.S, cpy.data.V = r, s, v
	return cpy, nil
}

// Cost returns amount + computational effort * compute unit price.
func (tx *Transaction) Cost() *big.Int {
	total := new(big.Int).Mul(params.ComputeUnitPrice, new(big.Int).SetUint64(tx.data.ComputeLimit))
	total.Add(total, tx.data.Amount)
	return total
}

func (tx *Transaction) RawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	return tx.data.V, tx.data.R, tx.data.S
}

func (tx *Transaction) String() string {
	var from, to string
	if tx.data.V != nil {
		// make a best guess about the signer and use that to derive
		// the sender.
		signer := deriveSigner(tx.data.V)
		if f, err := TxSender(signer, tx); err != nil { // derive but don't cache
			from = "[invalid sender: invalid sig] " + err.Error()
		} else {
			from = fmt.Sprintf("%x", f[:])
		}
	} else {
		from = "[invalid sender: nil V field]"
	}

	if tx.data.Recipient == nil {
		to = "[contract creation]"
	} else {
		to = fmt.Sprintf("%x", tx.data.Recipient[:])
	}
	enc, _ := rlp.EncodeToBytes(&tx.data)
	return fmt.Sprintf(`
	TX(%x)
	Contract:      %v
	From:          %s
	To:            %s
	Nonce:         %v
	ComputeLimit:  %v
	Value:         %#x
	Data:          0x%x
	V:             %#x
	R:             %#x
	S:             %#x
	Hex:           %x
`,
		tx.Hash(),
		tx.data.Recipient == nil,
		from,
		to,
		tx.data.AccountNonce,
		tx.data.ComputeLimit,
		tx.data.Amount,
		tx.data.Payload,
		tx.data.V,
		tx.data.R,
		tx.data.S,
		enc,
	)
}

// Transactions is a Transaction slice type for basic sorting.
type Transactions []*Transaction

// Len returns the length of s.
func (s Transactions) Len() int { return len(s) }

// Swap swaps the i'th and the j'th element in s.
func (s Transactions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// GetRlp implements Rlpable and returns the i'th element of s in rlp.
func (s Transactions) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(s[i])
	return enc
}

// TxDifference returns a new set t which is the difference between a to b.
func TxDifference(a, b Transactions) (keep Transactions) {
	keep = make(Transactions, 0, len(a))

	remove := make(map[common.Hash]struct{})
	for _, tx := range b {
		remove[tx.Hash()] = struct{}{}
	}

	for _, tx := range a {
		if _, ok := remove[tx.Hash()]; !ok {
			keep = append(keep, tx)
		}
	}

	return keep
}

// Message is a fully derived transaction and implements core.Message
//
// NOTE: In a future PR this will be removed.
type Message struct {
	to           *common.Address
	from         common.Address
	nonce        uint64
	amount       *big.Int
	computeLimit uint64
	data         []byte
	checkNonce   bool
}

func NewMessage(from common.Address, to *common.Address, nonce uint64, amount *big.Int, computeLimit uint64, data []byte, checkNonce bool) Message {
	return Message{
		from:         from,
		to:           to,
		nonce:        nonce,
		amount:       amount,
		computeLimit: computeLimit,
		data:         data,
		checkNonce:   checkNonce,
	}
}

func (m Message) From() common.Address { return m.from }
func (m Message) To() *common.Address  { return m.to }
func (m Message) Value() *big.Int      { return m.amount }
func (m Message) ComputeLimit() uint64 { return m.computeLimit }
func (m Message) Nonce() uint64        { return m.nonce }
func (m Message) Data() []byte         { return m.data }
func (m Message) CheckNonce() bool     { return m.checkNonce }
