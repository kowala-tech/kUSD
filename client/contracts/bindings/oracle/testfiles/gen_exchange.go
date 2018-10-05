// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testfiles

import (
	"math/big"
	"strings"

	kowala "github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
)

// ExchangeMgrABI is the input ABI used to generate the binding from.
const ExchangeMgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"blacklistExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"isExchange\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"isWhitelistedExchange\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"isBlacklistedExchange\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWhitelistedExchangeCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"whitelistExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"removeExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getWhitelistedExchangeAtIndex\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exchange\",\"type\":\"string\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exchange\",\"type\":\"string\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exchange\",\"type\":\"string\"}],\"name\":\"Addition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exchange\",\"type\":\"string\"}],\"name\":\"Removal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ExchangeMgrBin is the compiled bytecode used for deploying new contracts.
const ExchangeMgrBin = `608060405260008060146101000a81548160ff021916908315150217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611abb8061006d6000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063158ef93e146100f65780633b56711b146101255780633f4ba83a1461018e5780634bbc5a61146101a5578063569bca69146102265780635c975abb146102a757806364dfaff8146102d65780636fcceefe14610357578063715018a6146103825780637ebd1b30146103995780638456cb591461043f5780638da5cb5b1461045657806393a25542146104ad578063a7e71ee314610516578063ad6a23b21461057f578063f2fde38b146105e8578063fc44d0b41461062b575b600080fd5b34801561010257600080fd5b5061010b6106d1565b604051808215151515815260200191505060405180910390f35b34801561013157600080fd5b5061018c600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506106e4565b005b34801561019a57600080fd5b506101a3610a23565b005b3480156101b157600080fd5b5061020c600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610ae1565b604051808215151515815260200191505060405180910390f35b34801561023257600080fd5b5061028d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610b66565b604051808215151515815260200191505060405180910390f35b3480156102b357600080fd5b506102bc610bfc565b604051808215151515815260200191505060405180910390f35b3480156102e257600080fd5b5061033d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610c0f565b604051808215151515815260200191505060405180910390f35b34801561036357600080fd5b5061036c610ca6565b6040518082815260200191505060405180910390f35b34801561038e57600080fd5b50610397610cb3565b005b3480156103a557600080fd5b506103c460048036038101908080359060200190929190505050610db5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104045780820151818401526020810190506103e9565b50505050905090810190601f1680156104315780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561044b57600080fd5b50610454610e70565b005b34801561046257600080fd5b5061046b610f30565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104b957600080fd5b50610514600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610f55565b005b34801561052257600080fd5b5061057d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611249565b005b34801561058b57600080fd5b506105e6600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611474565b005b3480156105f457600080fd5b50610629600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116a7565b005b34801561063757600080fd5b506106566004803603810190808035906020019092919050505061170e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561069657808201518184015260208101905061067b565b50505050905090810190601f1680156106c35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600060159054906101000a900460ff1681565b60008060008060149054906101000a900460ff1615151561070457600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561075f57600080fd5b8361076981610b66565b1515610803576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001807f676976656e206e616d65206973206e6f7420612077686974656c69737465642081526020017f65786368616e676500000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6001856040518082805190602001908083835b60208310151561083b5780518252602082019150602081019050602083039250610816565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902093508360000154925060008460010160016101000a81548160ff02191690831515021790555060026001600280549050038154811015156108ab57fe5b906000526020600020019150816002848154811015156108c757fe5b9060005260206000200190805460018160011615610100020316600290046108f09291906118c3565b508260018360405180828054600181600116156101000203166002900480156109505780601f1061092e576101008083540402835291820191610950565b820191906000526020600020905b81548152906001019060200180831161093c575b5050915050908152602001604051809103902060000181905550600280548091906001900361097f919061194a565b507f81e32445db1ab3f971a6e0ae415ce6f58f8ad4ac1e87f6eaf42067e404f81dae856040518080602001828103825283818151815260200191508051906020019080838360005b838110156109e25780820151818401526020810190506109c7565b50505050905090810190601f168015610a0f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a7e57600080fd5b600060149054906101000a900460ff161515610a9957600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60006001826040518082805190602001908083835b602083101515610b1b5780518252602082019150602081019050602083039250610af6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff169050919050565b6000610b7182610ae1565b8015610bf557506001826040518082805190602001908083835b602083101515610bb05780518252602082019150602081019050602083039250610b8b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160019054906101000a900460ff165b9050919050565b600060149054906101000a900460ff1681565b6000610c1a82610ae1565b8015610c9f57506001826040518082805190602001908083835b602083101515610c595780518252602082019150602081019050602083039250610c34565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160019054906101000a900460ff16155b9050919050565b6000600280549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d0e57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600281815481101515610dc457fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e685780601f10610e3d57610100808354040283529160200191610e68565b820191906000526020600020905b815481529060010190602001808311610e4b57829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ecb57600080fd5b600060149054906101000a900460ff16151515610ee757600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060149054906101000a900460ff16151515610f7157600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fcc57600080fd5b80610fd681610c0f565b1515611070576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001807f676976656e206e616d65206973206e6f74206120626c61636b6c69737465642081526020017f65786368616e676500000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600160028390806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906110af929190611976565b50036001836040518082805190602001908083835b6020831015156110e957805182526020820191506020810190506020830392506110c4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000181905550600180836040518082805190602001908083835b60208310151561115c5780518252602082019150602081019050602083039250611137565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160016101000a81548160ff0219169083151502179055507fce09318849d53fa41700de1700ae3708b944cb20a5a4a4fd51a331625ce72a7a826040518080602001828103825283818151815260200191508051906020019080838360005b8381101561120b5780820151818401526020810190506111f0565b50505050905090810190601f1680156112385780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050565b60008060149054906101000a900460ff1615151561126657600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112c157600080fd5b816112cb81610ae1565b151515611340576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f65786368616e676520616c72656164792065786973747300000000000000000081525060200191505060405180910390fd5b6001836040518082805190602001908083835b6020831015156113785780518252602082019150602081019050602083039250611353565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020915060018260010160006101000a81548160ff0219169083151502179055507f9874713fe684a190f03252911ed81f4e7754af845b47e2734570c0a8d78417a8836040518080602001828103825283818151815260200191508051906020019080838360005b8381101561142c578082015181840152602081019050611411565b50505050905090810190601f1680156114595780820380516001836020036101000a031916815260200191505b509250505060405180910390a161146f83610f55565b505050565b600060149054906101000a900460ff1615151561149057600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156114eb57600080fd5b806114f581610ae1565b1515611569576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f676976656e206e616d65206973206e6f7420616e2065786368616e676500000081525060200191505060405180910390fd5b6001826040518082805190602001908083835b6020831015156115a1578051825260208201915060208101905060208303925061157c565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000808201600090556001820160006101000a81549060ff02191690556001820160016101000a81549060ff021916905550507f9751cac4d5046f38f52784e809c6e39f08eb2846a96993b1ea7528937f380e1b826040518080602001828103825283818151815260200191508051906020019080838360005b8381101561166957808201518184015260208101905061164e565b50505050905090810190601f1680156116965780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561170257600080fd5b61170b816117c9565b50565b606060028281548110151561171f57fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117bd5780601f10611792576101008083540402835291602001916117bd565b820191906000526020600020905b8154815290600101906020018083116117a057829003601f168201915b50505050509050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561180557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106118fc5780548555611939565b8280016001018555821561193957600052602060002091601f016020900482015b8281111561193857825482559160010191906001019061191d565b5b50905061194691906119f6565b5090565b815481835581811115611971578183600052602060002091820191016119709190611a1b565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106119b757805160ff19168380011785556119e5565b828001600101855582156119e5579182015b828111156119e45782518255916020019190600101906119c9565b5b5090506119f291906119f6565b5090565b611a1891905b80821115611a145760008160009055506001016119fc565b5090565b90565b611a4491905b80821115611a405760008181611a379190611a47565b50600101611a21565b5090565b90565b50805460018160011615610100020316600290046000825580601f10611a6d5750611a8c565b601f016020900490600052602060002090810190611a8b91906119f6565b5b505600a165627a7a7230582087d48bce60e613708d842ce9209eac3f78b5b0436755eb466417a51291c37d100029`

// DeployExchangeMgr deploys a new Kowala contract, binding an instance of ExchangeMgr to it.
func DeployExchangeMgr(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExchangeMgr, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeMgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeMgrBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExchangeMgr{ExchangeMgrCaller: ExchangeMgrCaller{contract: contract}, ExchangeMgrTransactor: ExchangeMgrTransactor{contract: contract}, ExchangeMgrFilterer: ExchangeMgrFilterer{contract: contract}}, nil
}

// ExchangeMgr is an auto generated Go binding around a Kowala contract.
type ExchangeMgr struct {
	ExchangeMgrCaller     // Read-only binding to the contract
	ExchangeMgrTransactor // Write-only binding to the contract
	ExchangeMgrFilterer   // Log filterer for contract events
}

// ExchangeMgrCaller is an auto generated read-only Go binding around a Kowala contract.
type ExchangeMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeMgrTransactor is an auto generated write-only Go binding around a Kowala contract.
type ExchangeMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeMgrFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type ExchangeMgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeMgrSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type ExchangeMgrSession struct {
	Contract     *ExchangeMgr      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeMgrCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type ExchangeMgrCallerSession struct {
	Contract *ExchangeMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExchangeMgrTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type ExchangeMgrTransactorSession struct {
	Contract     *ExchangeMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExchangeMgrRaw is an auto generated low-level Go binding around a Kowala contract.
type ExchangeMgrRaw struct {
	Contract *ExchangeMgr // Generic contract binding to access the raw methods on
}

// ExchangeMgrCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type ExchangeMgrCallerRaw struct {
	Contract *ExchangeMgrCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeMgrTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type ExchangeMgrTransactorRaw struct {
	Contract *ExchangeMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeMgr creates a new instance of ExchangeMgr, bound to a specific deployed contract.
func NewExchangeMgr(address common.Address, backend bind.ContractBackend) (*ExchangeMgr, error) {
	contract, err := bindExchangeMgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeMgr{ExchangeMgrCaller: ExchangeMgrCaller{contract: contract}, ExchangeMgrTransactor: ExchangeMgrTransactor{contract: contract}, ExchangeMgrFilterer: ExchangeMgrFilterer{contract: contract}}, nil
}

// NewExchangeMgrCaller creates a new read-only instance of ExchangeMgr, bound to a specific deployed contract.
func NewExchangeMgrCaller(address common.Address, caller bind.ContractCaller) (*ExchangeMgrCaller, error) {
	contract, err := bindExchangeMgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrCaller{contract: contract}, nil
}

// NewExchangeMgrTransactor creates a new write-only instance of ExchangeMgr, bound to a specific deployed contract.
func NewExchangeMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeMgrTransactor, error) {
	contract, err := bindExchangeMgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrTransactor{contract: contract}, nil
}

// NewExchangeMgrFilterer creates a new log filterer instance of ExchangeMgr, bound to a specific deployed contract.
func NewExchangeMgrFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeMgrFilterer, error) {
	contract, err := bindExchangeMgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrFilterer{contract: contract}, nil
}

// bindExchangeMgr binds a generic wrapper to an already deployed contract.
func bindExchangeMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeMgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeMgr *ExchangeMgrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeMgr.Contract.ExchangeMgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeMgr *ExchangeMgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.ExchangeMgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeMgr *ExchangeMgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.ExchangeMgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeMgr *ExchangeMgrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeMgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeMgr *ExchangeMgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeMgr *ExchangeMgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.contract.Transact(opts, method, params...)
}

// GetWhitelistedExchangeAtIndex is a free data retrieval call binding the contract method 0xfc44d0b4.
//
// Solidity: function getWhitelistedExchangeAtIndex(index uint256) constant returns(name string)
func (_ExchangeMgr *ExchangeMgrCaller) GetWhitelistedExchangeAtIndex(opts *bind.CallOpts, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "getWhitelistedExchangeAtIndex", index)
	return *ret0, err
}

// GetWhitelistedExchangeAtIndex is a free data retrieval call binding the contract method 0xfc44d0b4.
//
// Solidity: function getWhitelistedExchangeAtIndex(index uint256) constant returns(name string)
func (_ExchangeMgr *ExchangeMgrSession) GetWhitelistedExchangeAtIndex(index *big.Int) (string, error) {
	return _ExchangeMgr.Contract.GetWhitelistedExchangeAtIndex(&_ExchangeMgr.CallOpts, index)
}

// GetWhitelistedExchangeAtIndex is a free data retrieval call binding the contract method 0xfc44d0b4.
//
// Solidity: function getWhitelistedExchangeAtIndex(index uint256) constant returns(name string)
func (_ExchangeMgr *ExchangeMgrCallerSession) GetWhitelistedExchangeAtIndex(index *big.Int) (string, error) {
	return _ExchangeMgr.Contract.GetWhitelistedExchangeAtIndex(&_ExchangeMgr.CallOpts, index)
}

// GetWhitelistedExchangeCount is a free data retrieval call binding the contract method 0x6fcceefe.
//
// Solidity: function getWhitelistedExchangeCount() constant returns(count uint256)
func (_ExchangeMgr *ExchangeMgrCaller) GetWhitelistedExchangeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "getWhitelistedExchangeCount")
	return *ret0, err
}

// GetWhitelistedExchangeCount is a free data retrieval call binding the contract method 0x6fcceefe.
//
// Solidity: function getWhitelistedExchangeCount() constant returns(count uint256)
func (_ExchangeMgr *ExchangeMgrSession) GetWhitelistedExchangeCount() (*big.Int, error) {
	return _ExchangeMgr.Contract.GetWhitelistedExchangeCount(&_ExchangeMgr.CallOpts)
}

// GetWhitelistedExchangeCount is a free data retrieval call binding the contract method 0x6fcceefe.
//
// Solidity: function getWhitelistedExchangeCount() constant returns(count uint256)
func (_ExchangeMgr *ExchangeMgrCallerSession) GetWhitelistedExchangeCount() (*big.Int, error) {
	return _ExchangeMgr.Contract.GetWhitelistedExchangeCount(&_ExchangeMgr.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_ExchangeMgr *ExchangeMgrCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_ExchangeMgr *ExchangeMgrSession) Initialized() (bool, error) {
	return _ExchangeMgr.Contract.Initialized(&_ExchangeMgr.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_ExchangeMgr *ExchangeMgrCallerSession) Initialized() (bool, error) {
	return _ExchangeMgr.Contract.Initialized(&_ExchangeMgr.CallOpts)
}

// IsBlacklistedExchange is a free data retrieval call binding the contract method 0x64dfaff8.
//
// Solidity: function isBlacklistedExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrCaller) IsBlacklistedExchange(opts *bind.CallOpts, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "isBlacklistedExchange", name)
	return *ret0, err
}

// IsBlacklistedExchange is a free data retrieval call binding the contract method 0x64dfaff8.
//
// Solidity: function isBlacklistedExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrSession) IsBlacklistedExchange(name string) (bool, error) {
	return _ExchangeMgr.Contract.IsBlacklistedExchange(&_ExchangeMgr.CallOpts, name)
}

// IsBlacklistedExchange is a free data retrieval call binding the contract method 0x64dfaff8.
//
// Solidity: function isBlacklistedExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrCallerSession) IsBlacklistedExchange(name string) (bool, error) {
	return _ExchangeMgr.Contract.IsBlacklistedExchange(&_ExchangeMgr.CallOpts, name)
}

// IsExchange is a free data retrieval call binding the contract method 0x4bbc5a61.
//
// Solidity: function isExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrCaller) IsExchange(opts *bind.CallOpts, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "isExchange", name)
	return *ret0, err
}

// IsExchange is a free data retrieval call binding the contract method 0x4bbc5a61.
//
// Solidity: function isExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrSession) IsExchange(name string) (bool, error) {
	return _ExchangeMgr.Contract.IsExchange(&_ExchangeMgr.CallOpts, name)
}

// IsExchange is a free data retrieval call binding the contract method 0x4bbc5a61.
//
// Solidity: function isExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrCallerSession) IsExchange(name string) (bool, error) {
	return _ExchangeMgr.Contract.IsExchange(&_ExchangeMgr.CallOpts, name)
}

// IsWhitelistedExchange is a free data retrieval call binding the contract method 0x569bca69.
//
// Solidity: function isWhitelistedExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrCaller) IsWhitelistedExchange(opts *bind.CallOpts, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "isWhitelistedExchange", name)
	return *ret0, err
}

// IsWhitelistedExchange is a free data retrieval call binding the contract method 0x569bca69.
//
// Solidity: function isWhitelistedExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrSession) IsWhitelistedExchange(name string) (bool, error) {
	return _ExchangeMgr.Contract.IsWhitelistedExchange(&_ExchangeMgr.CallOpts, name)
}

// IsWhitelistedExchange is a free data retrieval call binding the contract method 0x569bca69.
//
// Solidity: function isWhitelistedExchange(name string) constant returns(isIndeed bool)
func (_ExchangeMgr *ExchangeMgrCallerSession) IsWhitelistedExchange(name string) (bool, error) {
	return _ExchangeMgr.Contract.IsWhitelistedExchange(&_ExchangeMgr.CallOpts, name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ExchangeMgr *ExchangeMgrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ExchangeMgr *ExchangeMgrSession) Owner() (common.Address, error) {
	return _ExchangeMgr.Contract.Owner(&_ExchangeMgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ExchangeMgr *ExchangeMgrCallerSession) Owner() (common.Address, error) {
	return _ExchangeMgr.Contract.Owner(&_ExchangeMgr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ExchangeMgr *ExchangeMgrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ExchangeMgr *ExchangeMgrSession) Paused() (bool, error) {
	return _ExchangeMgr.Contract.Paused(&_ExchangeMgr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ExchangeMgr *ExchangeMgrCallerSession) Paused() (bool, error) {
	return _ExchangeMgr.Contract.Paused(&_ExchangeMgr.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist( uint256) constant returns(string)
func (_ExchangeMgr *ExchangeMgrCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ExchangeMgr.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist( uint256) constant returns(string)
func (_ExchangeMgr *ExchangeMgrSession) Whitelist(arg0 *big.Int) (string, error) {
	return _ExchangeMgr.Contract.Whitelist(&_ExchangeMgr.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist( uint256) constant returns(string)
func (_ExchangeMgr *ExchangeMgrCallerSession) Whitelist(arg0 *big.Int) (string, error) {
	return _ExchangeMgr.Contract.Whitelist(&_ExchangeMgr.CallOpts, arg0)
}

// AddExchange is a paid mutator transaction binding the contract method 0xa7e71ee3.
//
// Solidity: function addExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactor) AddExchange(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "addExchange", name)
}

// AddExchange is a paid mutator transaction binding the contract method 0xa7e71ee3.
//
// Solidity: function addExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrSession) AddExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.AddExchange(&_ExchangeMgr.TransactOpts, name)
}

// AddExchange is a paid mutator transaction binding the contract method 0xa7e71ee3.
//
// Solidity: function addExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) AddExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.AddExchange(&_ExchangeMgr.TransactOpts, name)
}

// BlacklistExchange is a paid mutator transaction binding the contract method 0x3b56711b.
//
// Solidity: function blacklistExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactor) BlacklistExchange(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "blacklistExchange", name)
}

// BlacklistExchange is a paid mutator transaction binding the contract method 0x3b56711b.
//
// Solidity: function blacklistExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrSession) BlacklistExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.BlacklistExchange(&_ExchangeMgr.TransactOpts, name)
}

// BlacklistExchange is a paid mutator transaction binding the contract method 0x3b56711b.
//
// Solidity: function blacklistExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) BlacklistExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.BlacklistExchange(&_ExchangeMgr.TransactOpts, name)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ExchangeMgr *ExchangeMgrTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ExchangeMgr *ExchangeMgrSession) Pause() (*types.Transaction, error) {
	return _ExchangeMgr.Contract.Pause(&_ExchangeMgr.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) Pause() (*types.Transaction, error) {
	return _ExchangeMgr.Contract.Pause(&_ExchangeMgr.TransactOpts)
}

// RemoveExchange is a paid mutator transaction binding the contract method 0xad6a23b2.
//
// Solidity: function removeExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactor) RemoveExchange(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "removeExchange", name)
}

// RemoveExchange is a paid mutator transaction binding the contract method 0xad6a23b2.
//
// Solidity: function removeExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrSession) RemoveExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.RemoveExchange(&_ExchangeMgr.TransactOpts, name)
}

// RemoveExchange is a paid mutator transaction binding the contract method 0xad6a23b2.
//
// Solidity: function removeExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) RemoveExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.RemoveExchange(&_ExchangeMgr.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeMgr *ExchangeMgrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeMgr *ExchangeMgrSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExchangeMgr.Contract.RenounceOwnership(&_ExchangeMgr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExchangeMgr.Contract.RenounceOwnership(&_ExchangeMgr.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ExchangeMgr *ExchangeMgrTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ExchangeMgr *ExchangeMgrSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.TransferOwnership(&_ExchangeMgr.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.TransferOwnership(&_ExchangeMgr.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ExchangeMgr *ExchangeMgrTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ExchangeMgr *ExchangeMgrSession) Unpause() (*types.Transaction, error) {
	return _ExchangeMgr.Contract.Unpause(&_ExchangeMgr.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) Unpause() (*types.Transaction, error) {
	return _ExchangeMgr.Contract.Unpause(&_ExchangeMgr.TransactOpts)
}

// WhitelistExchange is a paid mutator transaction binding the contract method 0x93a25542.
//
// Solidity: function whitelistExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactor) WhitelistExchange(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ExchangeMgr.contract.Transact(opts, "whitelistExchange", name)
}

// WhitelistExchange is a paid mutator transaction binding the contract method 0x93a25542.
//
// Solidity: function whitelistExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrSession) WhitelistExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.WhitelistExchange(&_ExchangeMgr.TransactOpts, name)
}

// WhitelistExchange is a paid mutator transaction binding the contract method 0x93a25542.
//
// Solidity: function whitelistExchange(name string) returns()
func (_ExchangeMgr *ExchangeMgrTransactorSession) WhitelistExchange(name string) (*types.Transaction, error) {
	return _ExchangeMgr.Contract.WhitelistExchange(&_ExchangeMgr.TransactOpts, name)
}

// ExchangeMgrAdditionIterator is returned from FilterAddition and is used to iterate over the raw logs and unpacked data for Addition events raised by the ExchangeMgr contract.
type ExchangeMgrAdditionIterator struct {
	Event *ExchangeMgrAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrAddition represents a Addition event raised by the ExchangeMgr contract.
type ExchangeMgrAddition struct {
	Exchange string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddition is a free log retrieval operation binding the contract event 0x9874713fe684a190f03252911ed81f4e7754af845b47e2734570c0a8d78417a8.
//
// Solidity: e Addition(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) FilterAddition(opts *bind.FilterOpts) (*ExchangeMgrAdditionIterator, error) {

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "Addition")
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrAdditionIterator{contract: _ExchangeMgr.contract, event: "Addition", logs: logs, sub: sub}, nil
}

// WatchAddition is a free log subscription operation binding the contract event 0x9874713fe684a190f03252911ed81f4e7754af845b47e2734570c0a8d78417a8.
//
// Solidity: e Addition(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) WatchAddition(opts *bind.WatchOpts, sink chan<- *ExchangeMgrAddition) (event.Subscription, error) {

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "Addition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrAddition)
				if err := _ExchangeMgr.contract.UnpackLog(event, "Addition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeMgrBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the ExchangeMgr contract.
type ExchangeMgrBlacklistedIterator struct {
	Event *ExchangeMgrBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrBlacklisted represents a Blacklisted event raised by the ExchangeMgr contract.
type ExchangeMgrBlacklisted struct {
	Exchange string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0x81e32445db1ab3f971a6e0ae415ce6f58f8ad4ac1e87f6eaf42067e404f81dae.
//
// Solidity: e Blacklisted(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) FilterBlacklisted(opts *bind.FilterOpts) (*ExchangeMgrBlacklistedIterator, error) {

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "Blacklisted")
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrBlacklistedIterator{contract: _ExchangeMgr.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0x81e32445db1ab3f971a6e0ae415ce6f58f8ad4ac1e87f6eaf42067e404f81dae.
//
// Solidity: e Blacklisted(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *ExchangeMgrBlacklisted) (event.Subscription, error) {

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "Blacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrBlacklisted)
				if err := _ExchangeMgr.contract.UnpackLog(event, "Blacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeMgrOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ExchangeMgr contract.
type ExchangeMgrOwnershipRenouncedIterator struct {
	Event *ExchangeMgrOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrOwnershipRenounced represents a OwnershipRenounced event raised by the ExchangeMgr contract.
type ExchangeMgrOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ExchangeMgr *ExchangeMgrFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ExchangeMgrOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrOwnershipRenouncedIterator{contract: _ExchangeMgr.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ExchangeMgr *ExchangeMgrFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ExchangeMgrOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrOwnershipRenounced)
				if err := _ExchangeMgr.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeMgrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExchangeMgr contract.
type ExchangeMgrOwnershipTransferredIterator struct {
	Event *ExchangeMgrOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrOwnershipTransferred represents a OwnershipTransferred event raised by the ExchangeMgr contract.
type ExchangeMgrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ExchangeMgr *ExchangeMgrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeMgrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrOwnershipTransferredIterator{contract: _ExchangeMgr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ExchangeMgr *ExchangeMgrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeMgrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrOwnershipTransferred)
				if err := _ExchangeMgr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeMgrPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ExchangeMgr contract.
type ExchangeMgrPauseIterator struct {
	Event *ExchangeMgrPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrPause represents a Pause event raised by the ExchangeMgr contract.
type ExchangeMgrPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_ExchangeMgr *ExchangeMgrFilterer) FilterPause(opts *bind.FilterOpts) (*ExchangeMgrPauseIterator, error) {

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrPauseIterator{contract: _ExchangeMgr.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_ExchangeMgr *ExchangeMgrFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ExchangeMgrPause) (event.Subscription, error) {

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrPause)
				if err := _ExchangeMgr.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeMgrRemovalIterator is returned from FilterRemoval and is used to iterate over the raw logs and unpacked data for Removal events raised by the ExchangeMgr contract.
type ExchangeMgrRemovalIterator struct {
	Event *ExchangeMgrRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrRemoval represents a Removal event raised by the ExchangeMgr contract.
type ExchangeMgrRemoval struct {
	Exchange string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoval is a free log retrieval operation binding the contract event 0x9751cac4d5046f38f52784e809c6e39f08eb2846a96993b1ea7528937f380e1b.
//
// Solidity: e Removal(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) FilterRemoval(opts *bind.FilterOpts) (*ExchangeMgrRemovalIterator, error) {

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "Removal")
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrRemovalIterator{contract: _ExchangeMgr.contract, event: "Removal", logs: logs, sub: sub}, nil
}

// WatchRemoval is a free log subscription operation binding the contract event 0x9751cac4d5046f38f52784e809c6e39f08eb2846a96993b1ea7528937f380e1b.
//
// Solidity: e Removal(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) WatchRemoval(opts *bind.WatchOpts, sink chan<- *ExchangeMgrRemoval) (event.Subscription, error) {

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "Removal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrRemoval)
				if err := _ExchangeMgr.contract.UnpackLog(event, "Removal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeMgrUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ExchangeMgr contract.
type ExchangeMgrUnpauseIterator struct {
	Event *ExchangeMgrUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrUnpause represents a Unpause event raised by the ExchangeMgr contract.
type ExchangeMgrUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_ExchangeMgr *ExchangeMgrFilterer) FilterUnpause(opts *bind.FilterOpts) (*ExchangeMgrUnpauseIterator, error) {

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrUnpauseIterator{contract: _ExchangeMgr.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_ExchangeMgr *ExchangeMgrFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ExchangeMgrUnpause) (event.Subscription, error) {

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrUnpause)
				if err := _ExchangeMgr.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeMgrWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ExchangeMgr contract.
type ExchangeMgrWhitelistedIterator struct {
	Event *ExchangeMgrWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeMgrWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMgrWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeMgrWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeMgrWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMgrWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMgrWhitelisted represents a Whitelisted event raised by the ExchangeMgr contract.
type ExchangeMgrWhitelisted struct {
	Exchange string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xce09318849d53fa41700de1700ae3708b944cb20a5a4a4fd51a331625ce72a7a.
//
// Solidity: e Whitelisted(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) FilterWhitelisted(opts *bind.FilterOpts) (*ExchangeMgrWhitelistedIterator, error) {

	logs, sub, err := _ExchangeMgr.contract.FilterLogs(opts, "Whitelisted")
	if err != nil {
		return nil, err
	}
	return &ExchangeMgrWhitelistedIterator{contract: _ExchangeMgr.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xce09318849d53fa41700de1700ae3708b944cb20a5a4a4fd51a331625ce72a7a.
//
// Solidity: e Whitelisted(exchange string)
func (_ExchangeMgr *ExchangeMgrFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ExchangeMgrWhitelisted) (event.Subscription, error) {

	logs, sub, err := _ExchangeMgr.contract.WatchLogs(opts, "Whitelisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMgrWhitelisted)
				if err := _ExchangeMgr.contract.UnpackLog(event, "Whitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
