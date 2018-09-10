// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/math"
	"github.com/kowala-tech/kcoin/client/params"
)

// memoryGasCosts calculates the quadratic gas for memory expansion. It does so
// only for the memory region that is expanded, not the total memory.
func memoryGasCost(mem *Memory, newMemSize uint64) (uint64, error) {

	if newMemSize == 0 {
		return 0, nil
	}
	// The maximum that will fit in a uint64 is max_word_count - 1
	// anything above that will result in an overflow.
	// Additionally, a newMemSize which results in a
	// newMemSizeWords larger than 0x7ffffffff will cause the square operation
	// to overflow.
	// The constant 0xffffffffe0 is the highest number that can be used without
	// overflowing the gas calculation
	if newMemSize > 0xffffffffe0 {
		return 0, errComputationalEffortUintOverflow
	}

	newMemSizeWords := toWordSize(newMemSize)
	newMemSize = newMemSizeWords * 32

	if newMemSize > uint64(mem.Len()) {
		square := newMemSizeWords * newMemSizeWords
		linCoef := newMemSizeWords * params.MemoryCompEffort
		quadCoef := square / params.QuadCoeffDiv
		newTotalFee := linCoef + quadCoef

		fee := newTotalFee - mem.lastGasCost
		mem.lastGasCost = newTotalFee

		return fee, nil
	}
	return 0, nil
}

func constGasFunc(gas uint64) gasFunc {
	return func(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
		return gas, nil
	}
}

func gasCallDataCopy(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if gas, overflow = math.SafeAdd(gas, CompEffortFastestStep); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	words, overflow := bigUint64(stack.Back(2))
	if overflow {
		return 0, errComputationalEffortUintOverflow
	}

	if words, overflow = math.SafeMul(toWordSize(words), params.CopyCompEffort); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	if gas, overflow = math.SafeAdd(gas, words); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasReturnDataCopy(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if gas, overflow = math.SafeAdd(gas, CompEffortFastestStep); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	words, overflow := bigUint64(stack.Back(2))
	if overflow {
		return 0, errComputationalEffortUintOverflow
	}

	if words, overflow = math.SafeMul(toWordSize(words), params.CopyCompEffort); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	if gas, overflow = math.SafeAdd(gas, words); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasSStore(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var (
		y, x = stack.Back(1), stack.Back(0)
		val  = vm.StateDB.GetState(contract.Address(), common.BigToHash(x))
	)
	// This checks for 3 scenario's and calculates gas accordingly
	// 1. From a zero-value address to a non-zero value         (NEW VALUE)
	// 2. From a non-zero value address to a zero-value address (DELETE)
	// 3. From a non-zero to a non-zero                         (CHANGE)
	if val == (common.Hash{}) && y.Sign() != 0 {
		// 0 => non 0
		return params.SstoreSetCompEffort, nil
	} else if val != (common.Hash{}) && y.Sign() == 0 {
		// non 0 => 0
		vm.StateDB.AddRefund(params.SstoreRefundCompEffort)
		return params.SstoreClearCompEffort, nil
	} else {
		// non 0 => non 0 (or 0 => 0)
		return params.SstoreResetCompEffort, nil
	}
}

func makeGasLog(n uint64) gasFunc {
	return func(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
		requestedSize, overflow := bigUint64(stack.Back(1))
		if overflow {
			return 0, errComputationalEffortUintOverflow
		}

		gas, err := memoryGasCost(mem, memorySize)
		if err != nil {
			return 0, err
		}

		if gas, overflow = math.SafeAdd(gas, params.LogCompEffort); overflow {
			return 0, errComputationalEffortUintOverflow
		}
		if gas, overflow = math.SafeAdd(gas, n*params.LogTopicCompEffort); overflow {
			return 0, errComputationalEffortUintOverflow
		}

		var memorySizeGas uint64
		if memorySizeGas, overflow = math.SafeMul(requestedSize, params.LogDataComptEffort); overflow {
			return 0, errComputationalEffortUintOverflow
		}
		if gas, overflow = math.SafeAdd(gas, memorySizeGas); overflow {
			return 0, errComputationalEffortUintOverflow
		}
		return gas, nil
	}
}

func gasSha3(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}

	if gas, overflow = math.SafeAdd(gas, params.Sha3CompEffort); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	wordGas, overflow := bigUint64(stack.Back(1))
	if overflow {
		return 0, errComputationalEffortUintOverflow
	}
	if wordGas, overflow = math.SafeMul(toWordSize(wordGas), params.Sha3WordCompEffort); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	if gas, overflow = math.SafeAdd(gas, wordGas); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasCodeCopy(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if gas, overflow = math.SafeAdd(gas, CompEffortFastestStep); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	wordGas, overflow := bigUint64(stack.Back(2))
	if overflow {
		return 0, errComputationalEffortUintOverflow
	}
	if wordGas, overflow = math.SafeMul(toWordSize(wordGas), params.CopyCompEffort); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	if gas, overflow = math.SafeAdd(gas, wordGas); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasExtCodeCopy(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if gas, overflow = math.SafeAdd(gas, gt.ExtcodeCopy); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	wordGas, overflow := bigUint64(stack.Back(3))
	if overflow {
		return 0, errComputationalEffortUintOverflow
	}

	if wordGas, overflow = math.SafeMul(toWordSize(wordGas), params.CopyCompEffort); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	if gas, overflow = math.SafeAdd(gas, wordGas); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasMLoad(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, errComputationalEffortUintOverflow
	}
	if gas, overflow = math.SafeAdd(gas, CompEffortFastestStep); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasMStore8(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, errComputationalEffortUintOverflow
	}
	if gas, overflow = math.SafeAdd(gas, CompEffortFastestStep); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasMStore(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, errComputationalEffortUintOverflow
	}
	if gas, overflow = math.SafeAdd(gas, CompEffortFastestStep); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasCreate(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}
	if gas, overflow = math.SafeAdd(gas, params.CreateCompEffort); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasBalance(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return gt.Balance, nil
}

func gasExtCodeSize(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return gt.ExtcodeSize, nil
}

func gasSLoad(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return gt.SLoad, nil
}

func gasExp(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	expByteLen := uint64((stack.data[stack.len()-2].BitLen() + 7) / 8)

	var (
		gas      = expByteLen * gt.ExpByte // no overflow check required. Max is 256 * ExpByte gas
		overflow bool
	)
	if gas, overflow = math.SafeAdd(gas, CompEffortSlowStep); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasCall(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var (
		gas            = gt.Calls
		transfersValue = stack.Back(2).Sign() != 0
		address        = common.BigToAddress(stack.Back(1))
	)

	if transfersValue && vm.StateDB.Empty(address) {
		gas += params.CallNewAccountCompEffort
	}

	if transfersValue {
		gas += params.CallValueTransferComputEffort
	}
	memoryGas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if gas, overflow = math.SafeAdd(gas, memoryGas); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	vm.callGasTemp, err = calcComputationalEffort(gt, contract.ComputationalResources, gas, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if gas, overflow = math.SafeAdd(gas, vm.callGasTemp); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasCallCode(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	gas := gt.Calls
	if stack.Back(2).Sign() != 0 {
		gas += params.CallValueTransferComputEffort
	}
	memoryGas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if gas, overflow = math.SafeAdd(gas, memoryGas); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	vm.callGasTemp, err = calcComputationalEffort(gt, contract.ComputationalResources, gas, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if gas, overflow = math.SafeAdd(gas, vm.callGasTemp); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasReturn(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return memoryGasCost(mem, memorySize)
}

func gasRevert(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return memoryGasCost(mem, memorySize)
}

func gasSuicide(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var gas uint64

	gas = gt.Suicide
	var (
		address = common.BigToAddress(stack.Back(0))
	)

	// if empty and transfers value
	if vm.StateDB.Empty(address) && vm.StateDB.GetBalance(contract.Address()).Sign() != 0 {
		gas += gt.CreateBySuicide
	}

	if !vm.StateDB.HasSuicided(contract.Address()) {
		vm.StateDB.AddRefund(params.SuicideRefundCompEffort)
	}
	return gas, nil
}

func gasDelegateCall(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if gas, overflow = math.SafeAdd(gas, gt.Calls); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	vm.callGasTemp, err = calcComputationalEffort(gt, contract.ComputationalResources, gas, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if gas, overflow = math.SafeAdd(gas, vm.callGasTemp); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasStaticCall(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	gas, err := memoryGasCost(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if gas, overflow = math.SafeAdd(gas, gt.Calls); overflow {
		return 0, errComputationalEffortUintOverflow
	}

	vm.callGasTemp, err = calcComputationalEffort(gt, contract.ComputationalResources, gas, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if gas, overflow = math.SafeAdd(gas, vm.callGasTemp); overflow {
		return 0, errComputationalEffortUintOverflow
	}
	return gas, nil
}

func gasPush(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return CompEffortFastestStep, nil
}

func gasSwap(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return CompEffortFastestStep, nil
}

func gasDup(gt params.GasTable, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return CompEffortFastestStep, nil
}
