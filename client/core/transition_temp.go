package core

import (
	"errors"
	"math"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/stability"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

var (
	errInsufficientBalanceForGas          = errors.New("insufficient balance to pay for gas")
	errInsufficientBalanceForStabilityFee = errors.New("insufficient balance to pay for the stability fee")
)

type State interface {
	AddBalance(common.Address, *big.Int)
	SetNonce(common.Address, uint64)
	GetNonce(common.Address)
}

type ExecutionEnvironment interface {
	State
	Create(vm.ContractRef, []byte, uint64, *big.Int) ([]byte, common.Address, uint64, error)
	Call(vm.ContractRef, common.Address, []byte, uint64, *big.Int) ([]byte, uint64, error)
	StabilizationLevel() uint64
	Proposer() common.Address
}

/*
The State Transitioning Model

A state transition is a change made when a transaction is applied to the current world state
The state transitioning model does all all the necessary work to work out a valid new state root.

1) Nonce handling
2) Pre pay stability fee
3) Pre pay gas
4) Create a new state object if the recipient is \0*32
5) Value transfer
== If contract creation ==
  5a) Attempt to run transaction data
  5b) If valid, use result as code for the new state object
== end ==
6) Run Script section
7) Derive new state root

*/
type StateTransition struct {
	gasPool      *GasPool
	msg          Message
	gas          uint64
	gasPrice     *big.Int
	initialGas   uint64
	stabilityFee *big.Int
	value        *big.Int
	data         []byte
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(env ExecutionEnvironment, msg Message, gasPool *GasPool) *StateTransition {
	return &StateTransition{
		gasPool:      gasPool,
		vm:           vm,
		msg:          msg,
		gasPrice:     msg.GasPrice(),
		value:        msg.Value(),
		data:         msg.Data(),
		stabilityFee: common.Big0,
	}
}

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any VM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(env ExecutionEnvironment, msg Message) ([]byte, uint64, *big.Int, bool, error) {

	st := NewStateTransition(vm, msg, gasPool)

	// Make sure this transaction's nonce is correct.
	if st.msg.CheckNonce() {
		nonce := st.state.GetNonce(st.msg.From())
		if nonce < st.msg.Nonce() {
			return nil, 0, nil, false, ErrNonceTooHigh
		} else if nonce > st.msg.Nonce() {
			return nil, 0, nil, false, ErrNonceTooLow
		}
	}

	contractCreation := msg.To() == nil

	// intrinsic gas
	gas, err := IntrinsicGas(st.data, contractCreation)
	if err != nil {
		return nil, 0, common.Big0, false, err
	}

	if err = st.preCheck(gas); err != nil {
		return
	}

	sender := vm.AccountRef(msg.From())

	if err = st.useGas(gas); err != nil {
		return nil, 0, common.Big0, false, err
	}

	var (
		vmachine = st.vm
		// vm errors do not effect consensus and are therefor
		// not assigned to err, except for insufficient balance
		// error.
		vmerr error
	)
	if contractCreation {
		ret, _, st.gas, vmerr = env.Create(sender, st.data, st.gas, st.value)
	} else {
		// Increment the nonce for the next transaction
		vm.SetNonce(msg.From(), vm.GetNonce(sender.Address())+1)
		ret, st.gas, vmerr = env.Call(sender, st.to(), st.data, st.gas, st.value)
	}
	if vmerr != nil {
		log.Debug("VM returned with error", "err", vmerr)
		// The only possible consensus-error would be if there wasn't
		// sufficient balance to make the transfer happen. The first
		// balance transfer may never fail.
		if vmerr == vm.ErrInsufficientBalance {
			return nil, 0, common.Big0, false, vmerr
		}
	}

	st.refundGas()

	computeFees := new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice)

	// calculate the final stability fee and refund the remaining value.
	if st.stabilityFee.Cmp(common.Big0) > 0 && st.gas > 0 {
		st.calcAndRefundStabilityFee(computeFees)
	}

	vm.AddBalance(st.vm.Author(), computeFees)

	return ret, st.gasUsed(), st.stabilityFee, vmerr != nil, err
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	if st.msg == nil || st.msg.To() == nil /* contract creation */ {
		return common.Address{}
	}
	return *st.msg.To()
}

func (st *StateTransition) useGas(amount uint64) error {
	if st.gas < amount {
		return vm.ErrOutOfGas
	}
	st.gas -= amount

	return nil
}

func (st *StateTransition) buyGas() error {
	mgval := new(big.Int).Mul(new(big.Int).SetUint64(st.msg.Gas()), st.gasPrice)
	if st.vm.GetBalance(st.msg.From()).Cmp(mgval) < 0 {
		return errInsufficientBalanceForGas
	}
	if err := st.gasPool.SubGas(st.msg.Gas()); err != nil {
		return err
	}
	st.gas += st.msg.Gas()

	st.initialGas = st.msg.Gas()
	st.vm.SubBalance(st.msg.From(), mgval)
	return nil
}

func (st *StateTransition) preCheck(intrinsicGas uint64) error {

	// prepay stability fee if it applies
	if st.vm.StabilizationLevel() != 0 {
		computeFee := new(big.Int).Mul(new(big.Int).SetUint64(st.msg.Gas()), st.gasPrice)
		if stabilityFee := stability.CalcFee(computeFee, st.vm.StabilizationLevel(), st.msg.Value()); stabilityFee.Cmp(common.Big0) > 0 {
			if st.state.GetBalance(st.msg.From()).Cmp(stabilityFee) < 0 {
				return errInsufficientBalanceForStabilityFee
			}
			st.state.SubBalance(st.msg.From(), stabilityFee)
			st.stabilityFee = stabilityFee
		}
	}

	return st.buyGas()
}

// TransitionDb will transition the state by applying the current message and
// returning the result including the the used gas. It returns an error if it
// failed. An error indicates a consensus issue.
func (st *StateTransition) TransitionDb() (ret []byte, usedGas uint64, stabilityFee *big.Int, failed bool, err error) {

}

func (st *StateTransition) calcAndRefundStabilityFee(finalComputeFee *big.Int) {
	computeFee := new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice)
	if stabilityFee := stability.CalcFee(computeFee, st.vm.StabilizationLevel(), st.msg.Value()); stabilityFee.Cmp(common.Big0) > 0 {
		st.state.AddBalance(st.msg.From(), new(big.Int).Sub(st.stabilityFee, stabilityFee))
		st.stabilityFee = stabilityFee
	}
}

func (st *StateTransition) refundGas() {
	// Apply refund counter, capped to half of the used gas.
	refund := st.gasUsed() / 2
	if refund > st.state.GetRefund() {
		refund = st.state.GetRefund()
	}
	st.gas += refund

	// Return kUSD for remaining gas, exchanged at the original rate.
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.gas), st.gasPrice)
	st.state.AddBalance(st.msg.From(), remaining)

	// Also return remaining gas to the block gas counter so it is
	// available for the next transaction.
	st.gp.AddGas(st.gas)
}

// gasUsed returns the amount of gas used up by the state transition.
func (st *StateTransition) gasUsed() uint64 {
	return st.initialGas - st.gas
}

// IntrinsicGas computes the 'intrinsic gas' for a message with the given data.
func IntrinsicGas(data []byte, contractCreation bool) (uint64, error) {
	// Set the starting gas for the raw transaction
	var gas uint64
	if contractCreation {
		gas = params.TxGasContractCreation
	} else {
		gas = params.TxGas
	}
	// Bump the required gas by the amount of transactional data
	if len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		if (math.MaxUint64-gas)/params.TxDataNonZeroGas < nz {
			return 0, vm.ErrOutOfGas
		}
		gas += nz * params.TxDataNonZeroGas

		z := uint64(len(data)) - nz
		if (math.MaxUint64-gas)/params.TxDataZeroGas < z {
			return 0, vm.ErrOutOfGas
		}
		gas += z * params.TxDataZeroGas
	}
	return gas, nil
}
