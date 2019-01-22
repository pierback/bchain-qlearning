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
const CoffeedashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"states\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"setQValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getQtableState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"addState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stateCnt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"_lr\",\"type\":\"uint256\"},{\"name\":\"_gm\",\"type\":\"uint256\"},{\"name\":\"_ep\",\"type\":\"uint256\"},{\"name\":\"_epd\",\"type\":\"uint256\"},{\"name\":\"_sr\",\"type\":\"uint256\"},{\"name\":\"_pd\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// CoffeedashBin is the compiled bytecode used for deploying new contracts.
const CoffeedashBin = `60806040526040516200115e3803806200115e8339810180604052810190808051820192919060200180519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919050505060e060405190810160405280888152602001878152602001868152602001858152602001848152602001838152602001828152506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190620000f79291906200045d565b506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015590505062000187876040805190810160405280600481526020017f302c20300000000000000000000000000000000000000000000000000000000081525062000194640100000000026401000000009004565b5050505050505062000593565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701826040518082805190602001908083835b6020831015156200020d5780518252602082019150602081019050602083039250620001e6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff16156200025c57600080fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701836040518082805190602001908083835b602083101515620002d75780518252602082019150602081019050602083039250620002b0565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548160ff021916908315150217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701836040518082805190602001908083835b6020831015156200039e578051825260208201915060208101905060208303925062000377565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000019080519060200190620003e9929190620004e4565b5081600161011960009054906101000a900460ff1660ff16610118811015156200040f57fe5b01908051906020019062000425929190620004e4565b50610119600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff160217905550505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004a057805160ff1916838001178555620004d1565b82800160010185558215620004d1579182015b82811115620004d0578251825591602001919060010190620004b3565b5b509050620004e091906200056b565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200052757805160ff191683800117855562000558565b8280016001018555821562000558579182015b82811115620005575782518255916020019190600101906200053a565b5b5090506200056791906200056b565b5090565b6200059091905b808211156200058c57600081600090555060010162000572565b5090565b90565b610bbb80620005a36000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063017a91051461007257806312cfa7ce146101185780631672df68146101c75780632faf5168146102dc57806366ce8a4a1461038b575b600080fd5b34801561007e57600080fd5b5061009d600480360381019080803590602001909291905050506103bc565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100dd5780820151818401526020810190506100c2565b50505050905090810190601f16801561010a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012457600080fd5b506101c5600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061046f565b005b3480156101d357600080fd5b506101f5600480360381019080803560ff1690602001909291905050506105fb565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561023957808201518184015260208101905061021e565b50505050905090810190601f1680156102665780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561029f578082015181840152602081019050610284565b50505050905090810190601f1680156102cc5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156102e857600080fd5b50610389600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610819565b005b34801561039757600080fd5b506103a0610ad6565b604051808260ff1660ff16815260200191505060405180910390f35b600181610118811015156103cc57fe5b016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104675780601f1061043c57610100808354040283529160200191610467565b820191906000526020600020905b81548152906001019060200180831161044a57829003601f168201915b505050505081565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701826040518082805190602001908083835b6020831015156104e657805182526020820191506020810190506020830392506104c1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff16151561053557600080fd5b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701836040518082805190602001908083835b6020831015156105ad5780518252602082019150602081019050602083039250610588565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000190805190602001906105f6929190610aea565b505050565b60608060018360ff166101188110151561061157fe5b016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060070160018560ff166101188110151561066657fe5b0160405180828054600181600116156101000203166002900480156106c25780601f106106a05761010080835404028352918201916106c2565b820191906000526020600020905b8154815290600101906020018083116106ae575b50509150509081526020016040518091039020600001818054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561076d5780601f106107425761010080835404028352916020019161076d565b820191906000526020600020905b81548152906001019060200180831161075057829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108095780601f106107de57610100808354040283529160200191610809565b820191906000526020600020905b8154815290600101906020018083116107ec57829003601f168201915b5050505050905091509150915091565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701826040518082805190602001908083835b602083101515610890578051825260208201915060208101905060208303925061086b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff16156108de57600080fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701836040518082805190602001908083835b6020831015156109575780518252602082019150602081019050602083039250610932565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548160ff021916908315150217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701836040518082805190602001908083835b602083101515610a1c57805182526020820191506020810190506020830392506109f7565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000019080519060200190610a65929190610aea565b5081600161011960009054906101000a900460ff1660ff1661011881101515610a8a57fe5b019080519060200190610a9e929190610aea565b50610119600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff160217905550505050565b61011960009054906101000a900460ff1681565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b2b57805160ff1916838001178555610b59565b82800160010185558215610b59579182015b82811115610b58578251825591602001919060010190610b3d565b5b509050610b669190610b6a565b5090565b610b8c91905b80821115610b88576000816000905550600101610b70565b5090565b905600a165627a7a723058203196774d02b8c5cc38f8bb2292258249ae0630a0ffdd2ce15afc39927a35c98c0029`

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

// StateCnt is a free data retrieval call binding the contract method 0x66ce8a4a.
//
// Solidity: function stateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashCaller) StateCnt(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "stateCnt")
	return *ret0, err
}

// StateCnt is a free data retrieval call binding the contract method 0x66ce8a4a.
//
// Solidity: function stateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashSession) StateCnt() (uint8, error) {
	return _Coffeedash.Contract.StateCnt(&_Coffeedash.CallOpts)
}

// StateCnt is a free data retrieval call binding the contract method 0x66ce8a4a.
//
// Solidity: function stateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashCallerSession) StateCnt() (uint8, error) {
	return _Coffeedash.Contract.StateCnt(&_Coffeedash.CallOpts)
}

// States is a free data retrieval call binding the contract method 0x017a9105.
//
// Solidity: function states( uint256) constant returns(string)
func (_Coffeedash *CoffeedashCaller) States(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "states", arg0)
	return *ret0, err
}

// States is a free data retrieval call binding the contract method 0x017a9105.
//
// Solidity: function states( uint256) constant returns(string)
func (_Coffeedash *CoffeedashSession) States(arg0 *big.Int) (string, error) {
	return _Coffeedash.Contract.States(&_Coffeedash.CallOpts, arg0)
}

// States is a free data retrieval call binding the contract method 0x017a9105.
//
// Solidity: function states( uint256) constant returns(string)
func (_Coffeedash *CoffeedashCallerSession) States(arg0 *big.Int) (string, error) {
	return _Coffeedash.Contract.States(&_Coffeedash.CallOpts, arg0)
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
