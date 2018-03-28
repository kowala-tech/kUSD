package network

//go:generate solc --abi --bin --overwrite -o build contracts/Election.sol
//go:generate abigen -abi build/Election.abi -bin build/Election.bin -pkg contracts -type ElectionContract -out contracts/election.go

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/contracts/network/contracts"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/params"
)

var (
	errNoDeposit         = errors.New("user does not have deposits")
	errNoUnlockedDeposit = errors.New("user does not have unlocked deposits")
)

var mapChainIDToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62"),
}

type ValidatorsChecksum [32]byte

// Election is a gateway to validators contracts on the network
type Election interface {
	Join(walletAccount accounts.WalletAccount, amount uint64) (*types.Transaction, error)
	Leave(walletAccount accounts.WalletAccount) (*types.Transaction, error)
	RedeemDeposits(walletAccount accounts.WalletAccount) error
	ValidatorsChecksum() (ValidatorsChecksum, error)
	Validators() (types.ValidatorList, error)
	Deposits(address common.Address) ([]*types.Deposit, error)
	IsGenesisValidator(address common.Address) (bool, error)
	IsValidator(address common.Address) (bool, error)
	MinimumDeposit() (uint64, error)
}

type election struct {
	*contracts.ElectionContract
	chainID *big.Int
}

func NewElection(contractBackend bind.ContractBackend, chainID *big.Int) (*election, error) {
	contract, err := contracts.NewElectionContract(mapChainIDToAddr[chainID.Uint64()], contractBackend)
	if err != nil {
		return nil, err
	}

	return &election{
		ElectionContract: contract,
		chainID:          chainID,
	}, nil
}

func (election *election) Join(walletAccount accounts.WalletAccount, amount uint64) (*types.Transaction, error) {
	// @NOTE (rgeraldes) - the transaction can fail anyway if there's a similar transaction
	// on the same block that gets posted before ours.
	minDeposit, err := election.MinimumDeposit()
	if err != nil {
		return nil, err
	}
	if amount < minDeposit {
		return nil, fmt.Errorf("current deposit - %d - is not enough. The minimum required is %d", amount, minDeposit)
	}

	tx, err := election.ElectionContract.Join(election.transactDepositOpts(walletAccount, amount))
	if err != nil {
		return nil, fmt.Errorf("failed to transact the deposit: %s", err)
	}

	return tx, nil
}

func (election *election) Leave(walletAccount accounts.WalletAccount) (*types.Transaction, error) {
	tx, err := election.ElectionContract.Leave(election.transactOpts(walletAccount))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (election *election) RedeemDeposits(walletAccount accounts.WalletAccount) error {
	depositCount, err := election.GetDepositCount(&bind.CallOpts{From: walletAccount.Account().Address})
	if err != nil {
		return err
	}
	if depositCount.Uint64() < 1 {
		return errNoDeposit
	}
	deposit, err := election.GetDepositAtIndex(&bind.CallOpts{From: walletAccount.Account().Address}, common.Big0)
	if time.Unix(deposit.AvailableAt.Int64(), 0).Before(time.Now()) {
		return errNoUnlockedDeposit
	}

	_, err = election.ElectionContract.RedeemDeposits(election.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (election *election) ValidatorsChecksum() (ValidatorsChecksum, error) {
	return election.ElectionContract.ValidatorsChecksum(&bind.CallOpts{})
}

func (election *election) Validators() (types.ValidatorList, error) {
	count, err := election.GetValidatorCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	validators := make([]*types.Validator, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := election.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		weight := big.NewInt(0)
		validators[i] = types.NewValidator(validator.Code, validator.Deposit.Uint64(), weight)
	}

	return types.NewValidatorList(validators)
}

func (election *election) Deposits(addr common.Address) ([]*types.Deposit, error) {
	count, err := election.GetDepositCount(&bind.CallOpts{From: addr})
	if err != nil {
		return nil, err
	}

	deposits := make([]*types.Deposit, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		deposit, err := election.GetDepositAtIndex(&bind.CallOpts{From: addr}, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		deposits[i] = types.NewDeposit(deposit.Amount.Uint64(), deposit.AvailableAt.Int64())
	}

	return deposits, nil
}

func (election *election) IsGenesisValidator(address common.Address) (bool, error) {
	return election.ElectionContract.IsGenesisValidator(&bind.CallOpts{}, address)
}

func (election *election) IsValidator(address common.Address) (bool, error) {
	return election.ElectionContract.IsValidator(&bind.CallOpts{}, address)
}

func (election *election) transactOpts(walletAccount accounts.WalletAccount) *bind.TransactOpts {
	signerAddress := walletAccount.Account().Address
	opts := &bind.TransactOpts{
		From: signerAddress,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != signerAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			return walletAccount.SignTx(walletAccount.Account(), tx, election.chainID)
		},
	}

	return opts
}

func (election *election) MinimumDeposit() (uint64, error) {
	rawMinDeposit, err := election.GetMinimumDeposit(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return rawMinDeposit.Uint64(), nil
}

func (election *election) transactDepositOpts(walletAccount accounts.WalletAccount, amount uint64) *bind.TransactOpts {
	ops := election.transactOpts(walletAccount)
	var deposit big.Int
	ops.Value = deposit.SetUint64(amount)
	return ops
}
