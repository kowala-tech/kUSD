// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
)

var _ = (*headerMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash     common.Hash    `json:"parentHash"       gencodec:"required"`
		Coinbase       common.Address `json:"miner"            gencodec:"required"`
		Root           common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash         common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash    common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		ValidatorsHash common.Hash    `json:"validators"       gencodec:"required"`
		LastCommitHash common.Hash    `json:"lastCommit"       gencodec:"required"`
		Bloom          Bloom          `json:"logsBloom"        gencodec:"required"`
		Number         *hexutil.Big   `json:"number"           gencodec:"required"`
		GasLimit       hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed        hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		Time           *hexutil.Big   `json:"timestamp"        gencodec:"required"`
		Extra          hexutil.Bytes  `json:"extraData"        gencodec:"required"`
		Hash           common.Hash    `json:"hash"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.ValidatorsHash = h.ValidatorsHash
	enc.LastCommitHash = h.LastCommitHash
	enc.Bloom = h.Bloom
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = (*hexutil.Big)(h.Time)
	enc.Extra = h.Extra
	enc.Hash = h.Hash()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Header) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash     *common.Hash    `json:"parentHash"       gencodec:"required"`
		Coinbase       *common.Address `json:"miner"            gencodec:"required"`
		Root           *common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash         *common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash    *common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		ValidatorsHash *common.Hash    `json:"validators"       gencodec:"required"`
		LastCommitHash *common.Hash    `json:"lastCommit"       gencodec:"required"`
		Bloom          *Bloom          `json:"logsBloom"        gencodec:"required"`
		Number         *hexutil.Big    `json:"number"           gencodec:"required"`
		GasLimit       *hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed        *hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		Time           *hexutil.Big    `json:"timestamp"        gencodec:"required"`
		Extra          *hexutil.Bytes  `json:"extraData"        gencodec:"required"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *dec.ParentHash
	if dec.Coinbase == nil {
		return errors.New("missing required field 'miner' for Header")
	}
	h.Coinbase = *dec.Coinbase
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	h.Root = *dec.Root
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	h.TxHash = *dec.TxHash
	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	h.ReceiptHash = *dec.ReceiptHash
	if dec.ValidatorsHash == nil {
		return errors.New("missing required field 'validators' for Header")
	}
	h.ValidatorsHash = *dec.ValidatorsHash
	if dec.LastCommitHash == nil {
		return errors.New("missing required field 'lastCommit' for Header")
	}
	h.LastCommitHash = *dec.LastCommitHash
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	h.Bloom = *dec.Bloom
	if dec.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	h.Number = (*big.Int)(dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	h.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	h.GasUsed = uint64(*dec.GasUsed)
	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	h.Time = (*big.Int)(dec.Time)
	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	h.Extra = *dec.Extra
	return nil
}
