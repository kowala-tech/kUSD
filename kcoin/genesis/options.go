package genesis

import (
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/common"
)

type Options struct {
	Network           string
	Governance        *GovernanceOpts
	Consensus         *ConsensusOpts
	DataFeedSystem    *DataFeedSystemOpts
	PrefundedAccounts []PrefundedAccount
	ExtraData         string
}

type ConsensusOpts struct {
	Engine           string
	MaxNumValidators uint64
	FreezePeriod     uint64
	BaseDeposit      uint64
	Validators       []string
}

type GovernanceOpts struct {
	Origin           string
	Governors        []string
	NumConfirmations uint64
}

type DataFeedSystemOpts struct {
	MaxNumOracles uint64
	FreezePeriod  uint64 // in days
	BaseDeposit   uint64 // in kUSD
}

type PrefundedAccount struct {
	AccountAddress string
	Balance        string
}

type validValidatorMgrOpts struct {
	maxNumValidators *big.Int
	freezePeriod     *big.Int
	baseDeposit      *big.Int
	validators       []common.Address
}

type validOracleMgrOpts struct {
	maxNumOracles *big.Int
	freezePeriod  *big.Int
	baseDeposit   *big.Int
}

type validMultiSigOpts struct {
	multiSigCreator  *common.Address
	multiSigOwners   []common.Address
	numConfirmations *big.Int
}

type validPrefundedAccount struct {
	accountAddress *common.Address
	balance        *big.Int
}

type validGenesisOptions struct {
	network           string
	consensusEngine   string
	prefundedAccounts []*validPrefundedAccount
	multiSig          *validMultiSigOpts
	validatorMgr      *validValidatorMgrOpts
	oracleMgr         *validOracleMgrOpts
	ExtraData         string
}

func validateOptions(options Options) (*validGenesisOptions, error) {
	network, err := mapNetwork(options.Network)
	if err != nil {
		return nil, err
	}

	consensusEngine := TendermintConsensus
	if options.Consensus.Engine != "" {
		consensusEngine, err = mapConsensusEngine(options.Consensus.Engine)
		if err != nil {
			return nil, err
		}
	}

	// governance
	multiSigCreator, err := getAddress(options.Governance.Origin)
	if err != nil {
		return nil, err
	}

	multiSigOwners := make([]common.Address, len(options.Governance.Governors), 0)
	for _, governor := range options.Governance.Governors {
		owner, err := getAddress(governor)
		if err != nil {
			return nil, err
		}
		multiSigOwners = append(multiSigOwners, *owner)
	}

	numConfirmations := new(big.Int).SetUint64(options.Governance.NumConfirmations)

	// consensus
	maxNumValidators := new(big.Int).SetUint64(options.Consensus.MaxNumValidators)

	consensusBaseDeposit := new(big.Int).SetUint64(options.Consensus.BaseDeposit)

	consensusFreezePeriod := new(big.Int).SetUint64(options.Consensus.FreezePeriod)

	validators := make([]common.Address, len(options.Consensus.Validators), 0)
	for _, validator := range options.Consensus.Validators {
		validator, err := getAddress(validator)
		if err != nil {
			return nil, err
		}

		validators = append(validators, *validator)
	}

	// data feed system
	maxNumOracles := new(big.Int).SetUint64(options.DataFeedSystem.MaxNumOracles)

	oracleBaseDeposit := new(big.Int).SetUint64(options.DataFeedSystem.BaseDeposit)

	oracleFreezePeriod := new(big.Int).SetUint64(options.DataFeedSystem.FreezePeriod)

	// optional
	validPrefundedAccounts, err := mapPrefundedAccounts(options.PrefundedAccounts)
	if err != nil {
		return nil, err
	}

	if !prefundedIncludesValidatorWallet(validPrefundedAccounts, &validators[0]) {
		return nil, ErrWalletAddressValidatorNotInPrefundedAccounts
	}

	return &validGenesisOptions{
		network:         network,
		consensusEngine: consensusEngine,
		multiSig: &validMultiSigOpts{
			multiSigCreator:  multiSigCreator,
			multiSigOwners:   multiSigOwners,
			numConfirmations: numConfirmations,
		},
		validatorMgr: &validValidatorMgrOpts{
			maxNumValidators: maxNumValidators,
			freezePeriod:     consensusFreezePeriod,
			baseDeposit:      consensusBaseDeposit,
			validators:       validators,
		},
		oracleMgr: &validOracleMgrOpts{
			maxNumOracles: maxNumOracles,
			freezePeriod:  oracleFreezePeriod,
			baseDeposit:   oracleBaseDeposit,
		},
		prefundedAccounts: validPrefundedAccounts,
		ExtraData:         options.ExtraData,
	}, nil
}

func getAddress(s string) (*common.Address, error) {
	if !common.IsHexAddress(s) {
		return nil, fmt.Errorf("", "")
	}
	address := common.HexToAddress(s)

	return &address, nil
}
