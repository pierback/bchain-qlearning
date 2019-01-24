// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coffeedash

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CoffeedashABI is the input ABI used to generate the binding from.
const CoffeedashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getStateValues\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"setQValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getQtableState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"addState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"}],\"name\":\"stateExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStateCnt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"_lr\",\"type\":\"uint256\"},{\"name\":\"_gm\",\"type\":\"uint256\"},{\"name\":\"_ep\",\"type\":\"uint256\"},{\"name\":\"_epd\",\"type\":\"uint256\"},{\"name\":\"_sr\",\"type\":\"uint256\"},{\"name\":\"_pd\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// CoffeedashBin is the compiled bytecode used for deploying new contracts.
const CoffeedashBin = `60806040526040516200199c3803806200199c83398101806040528101908080518201929190602001805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050606061012060405190810160405280898152602001888152602001878152602001868152602001858152602001848152602001838152602001600060ff168152602001828152506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906200010a92919062000534565b506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548160ff021916908360ff16021790555061010082015181600801908051906020019062000187929190620005bb565b50905050620001db886040805190810160405280600481526020017f302c203000000000000000000000000000000000000000000000000000000000815250620001e9640100000000026401000000009004565b50505050505050506200074e565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b6020831015156200026257805182526020820191506020810190506020830392506200023b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff1615620002b157600080fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b6020831015156200032c578051825260208201915060208101905060208303925062000305565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548160ff021916908315150217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b602083101515620003f35780518252602082019150602081019050602083039250620003cc565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000190805190602001906200043e92919062000622565b506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600801829080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190620004bd92919062000622565b50506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff160217905550505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200057757805160ff1916838001178555620005a8565b82800160010185558215620005a8579182015b82811115620005a75782518255916020019190600101906200058a565b5b509050620005b79190620006a9565b5090565b8280548282559060005260206000209081019282156200060f579160200282015b828111156200060e578251829080519060200190620005fd92919062000534565b5091602001919060010190620005dc565b5b5090506200061e9190620006d1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200066557805160ff191683800117855562000696565b8280016001018555821562000696579182015b828111156200069557825182559160200191906001019062000678565b5b509050620006a59190620006a9565b5090565b620006ce91905b80821115620006ca576000816000905550600101620006b0565b5090565b90565b620006ff91905b80821115620006fb5760008181620006f1919062000702565b50600101620006d8565b5090565b90565b50805460018160011615610100020316600290046000825580601f106200072a57506200074b565b601f0160209004906000526020600020908101906200074a9190620006a9565b5b50565b61123e806200075e6000396000f300608060405260043610610082576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062d19f3014610087578063068531631461013057806312cfa7ce146101d95780631672df68146102885780632faf51681461039d5780638f05ae831461044c578063d46c97ba146104cd575b600080fd5b34801561009357600080fd5b506100b5600480360381019080803560ff1690602001909291905050506104fe565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f55780820151818401526020810190506100da565b50505050905090810190601f1680156101225780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013c57600080fd5b5061015e600480360381019080803560ff1690602001909291905050506105fb565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561019e578082015181840152602081019050610183565b50505050905090810190601f1680156101cb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101e557600080fd5b50610286600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610842565b005b34801561029457600080fd5b506102b6600480360381019080803560ff1690602001909291905050506109ce565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156102fa5780820151818401526020810190506102df565b50505050905090810190601f1680156103275780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610360578082015181840152602081019050610345565b50505050905090810190601f16801561038d5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156103a957600080fd5b5061044a600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d13565b005b34801561045857600080fd5b506104b3600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611053565b604051808215151515815260200191505060405180910390f35b3480156104d957600080fd5b506104e2611117565b604051808260ff1660ff16815260200191505060405180910390f35b60606000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018260ff1681548110151561055157fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105ef5780601f106105c4576101008083540402835291602001916105ef565b820191906000526020600020905b8154815290600101906020018083116105d257829003601f168201915b50505050509050919050565b6060806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018360ff1681548110151561064f57fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106ed5780601f106106c2576101008083540402835291602001916106ed565b820191906000526020600020905b8154815290600101906020018083116106d057829003601f168201915b505050505090506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901816040518082805190602001908083835b60208310151561076b5780518252602082019150602081019050602083039250610746565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108355780601f1061080a57610100808354040283529160200191610835565b820191906000526020600020905b81548152906001019060200180831161081857829003601f168201915b5050505050915050919050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b6020831015156108b95780518252602082019150602081019050602083039250610894565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff16151561090857600080fd5b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b602083101515610980578051825260208201915060208101905060208303925061095b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000190805190602001906109c992919061116d565b505050565b60608060606000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018460ff16815481101515610a2457fe5b906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ac25780601f10610a9757610100808354040283529160200191610ac2565b820191906000526020600020905b815481529060010190602001808311610aa557829003601f168201915b505050505090506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018460ff16815481101515610b1a57fe5b906000526020600020016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b602083101515610b9b5780518252602082019150602081019050602083039250610b76565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600001818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c665780601f10610c3b57610100808354040283529160200191610c66565b820191906000526020600020905b815481529060010190602001808311610c4957829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d025780601f10610cd757610100808354040283529160200191610d02565b820191906000526020600020905b815481529060010190602001808311610ce557829003601f168201915b505050505090509250925050915091565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b602083101515610d8a5780518252602082019150602081019050602083039250610d65565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff1615610dd857600080fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b602083101515610e515780518252602082019150602081019050602083039250610e2c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548160ff021916908315150217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b602083101515610f165780518252602082019150602081019050602083039250610ef1565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000019080519060200190610f5f92919061116d565b506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600801829080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190610fdc92919061116d565b50506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff160217905550505050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b6020831015156110cc57805182526020820191506020810190506020830392506110a7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff169050919050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060070160009054906101000a900460ff16905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106111ae57805160ff19168380011785556111dc565b828001600101855582156111dc579182015b828111156111db5782518255916020019190600101906111c0565b5b5090506111e991906111ed565b5090565b61120f91905b8082111561120b5760008160009055506001016111f3565b5090565b905600a165627a7a7230582013bd3e46338cdb46f8c6673f8cc1ad3245e8eba6803e2229e0855c83f95efc7f0029`

// DeployCoffeedash deploys a new Ethereum contract, binding an instance of Coffeedash to it.
func DeployCoffeedash(auth *bind.TransactOpts, backend bind.ContractBackend, _st string, _lr *big.Int, _gm *big.Int, _ep *big.Int, _epd *big.Int, _sr *big.Int, _pd *big.Int) (common.Address, *types.Transaction, *Coffeedash, error) {
	parsed, err := abi.JSON(strings.NewReader(CoffeedashABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoffeedashBin), backend, _st, _lr, _gm, _ep, _epd, _sr, _pd)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Coffeedash{CoffeedashCaller: CoffeedashCaller{contract: contract}, CoffeedashTransactor: CoffeedashTransactor{contract: contract}, CoffeedashFilterer: CoffeedashFilterer{contract: contract}}, nil
}

// Coffeedash is an auto generated Go binding around an Ethereum contract.
type Coffeedash struct {
	CoffeedashCaller     // Read-only binding to the contract
	CoffeedashTransactor // Write-only binding to the contract
	CoffeedashFilterer   // Log filterer for contract events
}

// CoffeedashCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoffeedashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeedashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoffeedashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeedashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoffeedashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeedashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoffeedashSession struct {
	Contract     *Coffeedash       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoffeedashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoffeedashCallerSession struct {
	Contract *CoffeedashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CoffeedashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoffeedashTransactorSession struct {
	Contract     *CoffeedashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CoffeedashRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoffeedashRaw struct {
	Contract *Coffeedash // Generic contract binding to access the raw methods on
}

// CoffeedashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoffeedashCallerRaw struct {
	Contract *CoffeedashCaller // Generic read-only contract binding to access the raw methods on
}

// CoffeedashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoffeedashTransactorRaw struct {
	Contract *CoffeedashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoffeedash creates a new instance of Coffeedash, bound to a specific deployed contract.
func NewCoffeedash(address common.Address, backend bind.ContractBackend) (*Coffeedash, error) {
	contract, err := bindCoffeedash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coffeedash{CoffeedashCaller: CoffeedashCaller{contract: contract}, CoffeedashTransactor: CoffeedashTransactor{contract: contract}, CoffeedashFilterer: CoffeedashFilterer{contract: contract}}, nil
}

// NewCoffeedashCaller creates a new read-only instance of Coffeedash, bound to a specific deployed contract.
func NewCoffeedashCaller(address common.Address, caller bind.ContractCaller) (*CoffeedashCaller, error) {
	contract, err := bindCoffeedash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoffeedashCaller{contract: contract}, nil
}

// NewCoffeedashTransactor creates a new write-only instance of Coffeedash, bound to a specific deployed contract.
func NewCoffeedashTransactor(address common.Address, transactor bind.ContractTransactor) (*CoffeedashTransactor, error) {
	contract, err := bindCoffeedash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoffeedashTransactor{contract: contract}, nil
}

// NewCoffeedashFilterer creates a new log filterer instance of Coffeedash, bound to a specific deployed contract.
func NewCoffeedashFilterer(address common.Address, filterer bind.ContractFilterer) (*CoffeedashFilterer, error) {
	contract, err := bindCoffeedash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoffeedashFilterer{contract: contract}, nil
}

// bindCoffeedash binds a generic wrapper to an already deployed contract.
func bindCoffeedash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoffeedashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coffeedash *CoffeedashRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coffeedash.Contract.CoffeedashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coffeedash *CoffeedashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeedash.Contract.CoffeedashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coffeedash *CoffeedashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coffeedash.Contract.CoffeedashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coffeedash *CoffeedashCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coffeedash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coffeedash *CoffeedashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeedash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coffeedash *CoffeedashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coffeedash.Contract.contract.Transact(opts, method, params...)
}

// GetQtableState is a free data retrieval call binding the contract method 0x1672df68.
//
// Solidity: function getQtableState(i uint8) constant returns(string, string)
func (_Coffeedash *CoffeedashCaller) GetQtableState(opts *bind.CallOpts, i uint8) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Coffeedash.contract.Call(opts, out, "getQtableState", i)
	return *ret0, *ret1, err
}

// GetQtableState is a free data retrieval call binding the contract method 0x1672df68.
//
// Solidity: function getQtableState(i uint8) constant returns(string, string)
func (_Coffeedash *CoffeedashSession) GetQtableState(i uint8) (string, string, error) {
	return _Coffeedash.Contract.GetQtableState(&_Coffeedash.CallOpts, i)
}

// GetQtableState is a free data retrieval call binding the contract method 0x1672df68.
//
// Solidity: function getQtableState(i uint8) constant returns(string, string)
func (_Coffeedash *CoffeedashCallerSession) GetQtableState(i uint8) (string, string, error) {
	return _Coffeedash.Contract.GetQtableState(&_Coffeedash.CallOpts, i)
}

// GetState is a free data retrieval call binding the contract method 0x00d19f30.
//
// Solidity: function getState(i uint8) constant returns(string)
func (_Coffeedash *CoffeedashCaller) GetState(opts *bind.CallOpts, i uint8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getState", i)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x00d19f30.
//
// Solidity: function getState(i uint8) constant returns(string)
func (_Coffeedash *CoffeedashSession) GetState(i uint8) (string, error) {
	return _Coffeedash.Contract.GetState(&_Coffeedash.CallOpts, i)
}

// GetState is a free data retrieval call binding the contract method 0x00d19f30.
//
// Solidity: function getState(i uint8) constant returns(string)
func (_Coffeedash *CoffeedashCallerSession) GetState(i uint8) (string, error) {
	return _Coffeedash.Contract.GetState(&_Coffeedash.CallOpts, i)
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashCaller) GetStateCnt(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getStateCnt")
	return *ret0, err
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashSession) GetStateCnt() (uint8, error) {
	return _Coffeedash.Contract.GetStateCnt(&_Coffeedash.CallOpts)
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashCallerSession) GetStateCnt() (uint8, error) {
	return _Coffeedash.Contract.GetStateCnt(&_Coffeedash.CallOpts)
}

// GetStateValues is a free data retrieval call binding the contract method 0x06853163.
//
// Solidity: function getStateValues(i uint8) constant returns(string)
func (_Coffeedash *CoffeedashCaller) GetStateValues(opts *bind.CallOpts, i uint8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getStateValues", i)
	return *ret0, err
}

// GetStateValues is a free data retrieval call binding the contract method 0x06853163.
//
// Solidity: function getStateValues(i uint8) constant returns(string)
func (_Coffeedash *CoffeedashSession) GetStateValues(i uint8) (string, error) {
	return _Coffeedash.Contract.GetStateValues(&_Coffeedash.CallOpts, i)
}

// GetStateValues is a free data retrieval call binding the contract method 0x06853163.
//
// Solidity: function getStateValues(i uint8) constant returns(string)
func (_Coffeedash *CoffeedashCallerSession) GetStateValues(i uint8) (string, error) {
	return _Coffeedash.Contract.GetStateValues(&_Coffeedash.CallOpts, i)
}

// StateExists is a free data retrieval call binding the contract method 0x8f05ae83.
//
// Solidity: function stateExists(_st string) constant returns(bool)
func (_Coffeedash *CoffeedashCaller) StateExists(opts *bind.CallOpts, _st string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "stateExists", _st)
	return *ret0, err
}

// StateExists is a free data retrieval call binding the contract method 0x8f05ae83.
//
// Solidity: function stateExists(_st string) constant returns(bool)
func (_Coffeedash *CoffeedashSession) StateExists(_st string) (bool, error) {
	return _Coffeedash.Contract.StateExists(&_Coffeedash.CallOpts, _st)
}

// StateExists is a free data retrieval call binding the contract method 0x8f05ae83.
//
// Solidity: function stateExists(_st string) constant returns(bool)
func (_Coffeedash *CoffeedashCallerSession) StateExists(_st string) (bool, error) {
	return _Coffeedash.Contract.StateExists(&_Coffeedash.CallOpts, _st)
}

// AddState is a paid mutator transaction binding the contract method 0x2faf5168.
//
// Solidity: function addState(_st string, qvalArr string) returns()
func (_Coffeedash *CoffeedashTransactor) AddState(opts *bind.TransactOpts, _st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.contract.Transact(opts, "addState", _st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0x2faf5168.
//
// Solidity: function addState(_st string, qvalArr string) returns()
func (_Coffeedash *CoffeedashSession) AddState(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.AddState(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0x2faf5168.
//
// Solidity: function addState(_st string, qvalArr string) returns()
func (_Coffeedash *CoffeedashTransactorSession) AddState(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.AddState(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x12cfa7ce.
//
// Solidity: function setQValue(_st string, qvalArr string) returns()
func (_Coffeedash *CoffeedashTransactor) SetQValue(opts *bind.TransactOpts, _st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.contract.Transact(opts, "setQValue", _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x12cfa7ce.
//
// Solidity: function setQValue(_st string, qvalArr string) returns()
func (_Coffeedash *CoffeedashSession) SetQValue(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.SetQValue(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x12cfa7ce.
//
// Solidity: function setQValue(_st string, qvalArr string) returns()
func (_Coffeedash *CoffeedashTransactorSession) SetQValue(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.SetQValue(&_Coffeedash.TransactOpts, _st, qvalArr)
}