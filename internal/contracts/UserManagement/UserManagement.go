// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package userManagement

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// UserManagementABI is the input ABI used to generate the binding from.
const UserManagementABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"userCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_email\",\"type\":\"bytes32\"},{\"name\":\"_ethAddress\",\"type\":\"address\"}],\"name\":\"insertUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_email\",\"type\":\"bytes32\"},{\"name\":\"_size\",\"type\":\"uint8\"},{\"name\":\"_strength\",\"type\":\"uint8\"}],\"name\":\"insertCoffee\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setTestAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_email\",\"type\":\"bytes32\"}],\"name\":\"getUser\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_size\",\"type\":\"uint8\"},{\"name\":\"_strength\",\"type\":\"uint8\"}],\"name\":\"getOverallCoffeeCnt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_email\",\"type\":\"bytes32\"},{\"name\":\"_size\",\"type\":\"uint8\"},{\"name\":\"_strength\",\"type\":\"uint8\"}],\"name\":\"getUserCoffeeCnt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTestAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UserManagementBin is the compiled bytecode used for deploying new contracts.
const UserManagementBin = `======= UserManagement.sol:UserManagement =======
Binary:
608060405234801561001057600080fd5b506106dd806100206000396000f300608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806307973ccf1461009e5780631b62e209146100c957806341aa8a761461011a578063584f0228146101585780636517579c1461019b5780636defbf801461020c57806399cf037f1461029c578063a2c11dc4146102ed578063d61478141461034c575b600080fd5b3480156100aa57600080fd5b506100b36103a3565b6040518082815260200191505060405180910390f35b3480156100d557600080fd5b506101186004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103af565b005b6101566004803603810190808035600019169060200190929190803560ff169060200190929190803560ff169060200190929190505050610466565b005b34801561016457600080fd5b50610199600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610501565b005b3480156101a757600080fd5b506101ca6004803603810190808035600019169060200190929190505050610545565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561021857600080fd5b5061022161058d565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610261578082015181840152602081019050610246565b50505050905090810190601f16801561028e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102a857600080fd5b506102d7600480360381019080803560ff169060200190929190803560ff1690602001909291905050506105ca565b6040518082815260200191505060405180910390f35b3480156102f957600080fd5b506103366004803603810190808035600019169060200190929190803560ff169060200190929190803560ff169060200190929190505050610612565b6040518082815260200191505060405180910390f35b34801561035857600080fd5b50610361610677565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008080549050905090565b8060016000846000191660001916815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016000839080600181540180825580915050906001820390600052602060002001600090919290919091509060001916905503600160008460001916600019168152602001908152602001600020600201819055505050565b600061048883600181111561047757fe5b83600281111561048357fe5b6106a1565b90506001610497858585610612565b0160016000866000191660001916815260200190815260200160002060010160008360ff1660ff1681526020019081526020016000208190555060016104dd84846105ca565b01600260008360ff1660ff1681526020019081526020016000208190555050505050565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60606040805190810160405280600c81526020017f49276d2072656164792121210000000000000000000000000000000000000000815250905090565b6000806105ed8460018111156105dc57fe5b8460028111156105e857fe5b6106a1565b9050600260008260ff1660ff1681526020019081526020016000205491505092915050565b60008061063584600181111561062457fe5b84600281111561063057fe5b6106a1565b905060016000866000191660001916815260200190815260200160002060010160008260ff1660ff168152602001908152602001600020549150509392505050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081600a8402019050929150505600a165627a7a7230582019bc7533f477d13c2274f5f298db09268b2d0b357e23a45980f55370e29be1630029`

// DeployUserManagement deploys a new Ethereum contract, binding an instance of UserManagement to it.
func DeployUserManagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserManagement, error) {
	parsed, err := abi.JSON(strings.NewReader(UserManagementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserManagementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserManagement{UserManagementCaller: UserManagementCaller{contract: contract}, UserManagementTransactor: UserManagementTransactor{contract: contract}, UserManagementFilterer: UserManagementFilterer{contract: contract}}, nil
}

// UserManagement is an auto generated Go binding around an Ethereum contract.
type UserManagement struct {
	UserManagementCaller     // Read-only binding to the contract
	UserManagementTransactor // Write-only binding to the contract
	UserManagementFilterer   // Log filterer for contract events
}

// UserManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserManagementSession struct {
	Contract     *UserManagement   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserManagementCallerSession struct {
	Contract *UserManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UserManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserManagementTransactorSession struct {
	Contract     *UserManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UserManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserManagementRaw struct {
	Contract *UserManagement // Generic contract binding to access the raw methods on
}

// UserManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserManagementCallerRaw struct {
	Contract *UserManagementCaller // Generic read-only contract binding to access the raw methods on
}

// UserManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserManagementTransactorRaw struct {
	Contract *UserManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserManagement creates a new instance of UserManagement, bound to a specific deployed contract.
func NewUserManagement(address common.Address, backend bind.ContractBackend) (*UserManagement, error) {
	contract, err := bindUserManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserManagement{UserManagementCaller: UserManagementCaller{contract: contract}, UserManagementTransactor: UserManagementTransactor{contract: contract}, UserManagementFilterer: UserManagementFilterer{contract: contract}}, nil
}

// NewUserManagementCaller creates a new read-only instance of UserManagement, bound to a specific deployed contract.
func NewUserManagementCaller(address common.Address, caller bind.ContractCaller) (*UserManagementCaller, error) {
	contract, err := bindUserManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserManagementCaller{contract: contract}, nil
}

// NewUserManagementTransactor creates a new write-only instance of UserManagement, bound to a specific deployed contract.
func NewUserManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*UserManagementTransactor, error) {
	contract, err := bindUserManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserManagementTransactor{contract: contract}, nil
}

// NewUserManagementFilterer creates a new log filterer instance of UserManagement, bound to a specific deployed contract.
func NewUserManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*UserManagementFilterer, error) {
	contract, err := bindUserManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserManagementFilterer{contract: contract}, nil
}

// bindUserManagement binds a generic wrapper to an already deployed contract.
func bindUserManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserManagement *UserManagementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserManagement.Contract.UserManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserManagement *UserManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserManagement.Contract.UserManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserManagement *UserManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserManagement.Contract.UserManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserManagement *UserManagementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserManagement *UserManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserManagement *UserManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserManagement.Contract.contract.Transact(opts, method, params...)
}

// GetOverallCoffeeCnt is a free data retrieval call binding the contract method 0x99cf037f.
//
// Solidity: function getOverallCoffeeCnt(_size uint8, _strength uint8) constant returns(uint256)
func (_UserManagement *UserManagementCaller) GetOverallCoffeeCnt(opts *bind.CallOpts, _size uint8, _strength uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserManagement.contract.Call(opts, out, "getOverallCoffeeCnt", _size, _strength)
	return *ret0, err
}

// GetOverallCoffeeCnt is a free data retrieval call binding the contract method 0x99cf037f.
//
// Solidity: function getOverallCoffeeCnt(_size uint8, _strength uint8) constant returns(uint256)
func (_UserManagement *UserManagementSession) GetOverallCoffeeCnt(_size uint8, _strength uint8) (*big.Int, error) {
	return _UserManagement.Contract.GetOverallCoffeeCnt(&_UserManagement.CallOpts, _size, _strength)
}

// GetOverallCoffeeCnt is a free data retrieval call binding the contract method 0x99cf037f.
//
// Solidity: function getOverallCoffeeCnt(_size uint8, _strength uint8) constant returns(uint256)
func (_UserManagement *UserManagementCallerSession) GetOverallCoffeeCnt(_size uint8, _strength uint8) (*big.Int, error) {
	return _UserManagement.Contract.GetOverallCoffeeCnt(&_UserManagement.CallOpts, _size, _strength)
}

// GetTestAddress is a free data retrieval call binding the contract method 0xd6147814.
//
// Solidity: function getTestAddress() constant returns(address)
func (_UserManagement *UserManagementCaller) GetTestAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserManagement.contract.Call(opts, out, "getTestAddress")
	return *ret0, err
}

// GetTestAddress is a free data retrieval call binding the contract method 0xd6147814.
//
// Solidity: function getTestAddress() constant returns(address)
func (_UserManagement *UserManagementSession) GetTestAddress() (common.Address, error) {
	return _UserManagement.Contract.GetTestAddress(&_UserManagement.CallOpts)
}

// GetTestAddress is a free data retrieval call binding the contract method 0xd6147814.
//
// Solidity: function getTestAddress() constant returns(address)
func (_UserManagement *UserManagementCallerSession) GetTestAddress() (common.Address, error) {
	return _UserManagement.Contract.GetTestAddress(&_UserManagement.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x6517579c.
//
// Solidity: function getUser(_email bytes32) constant returns(address)
func (_UserManagement *UserManagementCaller) GetUser(opts *bind.CallOpts, _email [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserManagement.contract.Call(opts, out, "getUser", _email)
	return *ret0, err
}

// GetUser is a free data retrieval call binding the contract method 0x6517579c.
//
// Solidity: function getUser(_email bytes32) constant returns(address)
func (_UserManagement *UserManagementSession) GetUser(_email [32]byte) (common.Address, error) {
	return _UserManagement.Contract.GetUser(&_UserManagement.CallOpts, _email)
}

// GetUser is a free data retrieval call binding the contract method 0x6517579c.
//
// Solidity: function getUser(_email bytes32) constant returns(address)
func (_UserManagement *UserManagementCallerSession) GetUser(_email [32]byte) (common.Address, error) {
	return _UserManagement.Contract.GetUser(&_UserManagement.CallOpts, _email)
}

// GetUserCoffeeCnt is a free data retrieval call binding the contract method 0xa2c11dc4.
//
// Solidity: function getUserCoffeeCnt(_email bytes32, _size uint8, _strength uint8) constant returns(uint256)
func (_UserManagement *UserManagementCaller) GetUserCoffeeCnt(opts *bind.CallOpts, _email [32]byte, _size uint8, _strength uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserManagement.contract.Call(opts, out, "getUserCoffeeCnt", _email, _size, _strength)
	return *ret0, err
}

// GetUserCoffeeCnt is a free data retrieval call binding the contract method 0xa2c11dc4.
//
// Solidity: function getUserCoffeeCnt(_email bytes32, _size uint8, _strength uint8) constant returns(uint256)
func (_UserManagement *UserManagementSession) GetUserCoffeeCnt(_email [32]byte, _size uint8, _strength uint8) (*big.Int, error) {
	return _UserManagement.Contract.GetUserCoffeeCnt(&_UserManagement.CallOpts, _email, _size, _strength)
}

// GetUserCoffeeCnt is a free data retrieval call binding the contract method 0xa2c11dc4.
//
// Solidity: function getUserCoffeeCnt(_email bytes32, _size uint8, _strength uint8) constant returns(uint256)
func (_UserManagement *UserManagementCallerSession) GetUserCoffeeCnt(_email [32]byte, _size uint8, _strength uint8) (*big.Int, error) {
	return _UserManagement.Contract.GetUserCoffeeCnt(&_UserManagement.CallOpts, _email, _size, _strength)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() constant returns(string)
func (_UserManagement *UserManagementCaller) Ready(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UserManagement.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() constant returns(string)
func (_UserManagement *UserManagementSession) Ready() (string, error) {
	return _UserManagement.Contract.Ready(&_UserManagement.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() constant returns(string)
func (_UserManagement *UserManagementCallerSession) Ready() (string, error) {
	return _UserManagement.Contract.Ready(&_UserManagement.CallOpts)
}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() constant returns(count uint256)
func (_UserManagement *UserManagementCaller) UserCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserManagement.contract.Call(opts, out, "userCount")
	return *ret0, err
}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() constant returns(count uint256)
func (_UserManagement *UserManagementSession) UserCount() (*big.Int, error) {
	return _UserManagement.Contract.UserCount(&_UserManagement.CallOpts)
}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() constant returns(count uint256)
func (_UserManagement *UserManagementCallerSession) UserCount() (*big.Int, error) {
	return _UserManagement.Contract.UserCount(&_UserManagement.CallOpts)
}

// InsertCoffee is a paid mutator transaction binding the contract method 0x41aa8a76.
//
// Solidity: function insertCoffee(_email bytes32, _size uint8, _strength uint8) returns()
func (_UserManagement *UserManagementTransactor) InsertCoffee(opts *bind.TransactOpts, _email [32]byte, _size uint8, _strength uint8) (*types.Transaction, error) {
	return _UserManagement.contract.Transact(opts, "insertCoffee", _email, _size, _strength)
}

// InsertCoffee is a paid mutator transaction binding the contract method 0x41aa8a76.
//
// Solidity: function insertCoffee(_email bytes32, _size uint8, _strength uint8) returns()
func (_UserManagement *UserManagementSession) InsertCoffee(_email [32]byte, _size uint8, _strength uint8) (*types.Transaction, error) {
	return _UserManagement.Contract.InsertCoffee(&_UserManagement.TransactOpts, _email, _size, _strength)
}

// InsertCoffee is a paid mutator transaction binding the contract method 0x41aa8a76.
//
// Solidity: function insertCoffee(_email bytes32, _size uint8, _strength uint8) returns()
func (_UserManagement *UserManagementTransactorSession) InsertCoffee(_email [32]byte, _size uint8, _strength uint8) (*types.Transaction, error) {
	return _UserManagement.Contract.InsertCoffee(&_UserManagement.TransactOpts, _email, _size, _strength)
}

// InsertUser is a paid mutator transaction binding the contract method 0x1b62e209.
//
// Solidity: function insertUser(_email bytes32, _ethAddress address) returns()
func (_UserManagement *UserManagementTransactor) InsertUser(opts *bind.TransactOpts, _email [32]byte, _ethAddress common.Address) (*types.Transaction, error) {
	return _UserManagement.contract.Transact(opts, "insertUser", _email, _ethAddress)
}

// InsertUser is a paid mutator transaction binding the contract method 0x1b62e209.
//
// Solidity: function insertUser(_email bytes32, _ethAddress address) returns()
func (_UserManagement *UserManagementSession) InsertUser(_email [32]byte, _ethAddress common.Address) (*types.Transaction, error) {
	return _UserManagement.Contract.InsertUser(&_UserManagement.TransactOpts, _email, _ethAddress)
}

// InsertUser is a paid mutator transaction binding the contract method 0x1b62e209.
//
// Solidity: function insertUser(_email bytes32, _ethAddress address) returns()
func (_UserManagement *UserManagementTransactorSession) InsertUser(_email [32]byte, _ethAddress common.Address) (*types.Transaction, error) {
	return _UserManagement.Contract.InsertUser(&_UserManagement.TransactOpts, _email, _ethAddress)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(_address address) returns()
func (_UserManagement *UserManagementTransactor) SetTestAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _UserManagement.contract.Transact(opts, "setTestAddress", _address)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(_address address) returns()
func (_UserManagement *UserManagementSession) SetTestAddress(_address common.Address) (*types.Transaction, error) {
	return _UserManagement.Contract.SetTestAddress(&_UserManagement.TransactOpts, _address)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(_address address) returns()
func (_UserManagement *UserManagementTransactorSession) SetTestAddress(_address common.Address) (*types.Transaction, error) {
	return _UserManagement.Contract.SetTestAddress(&_UserManagement.TransactOpts, _address)
}
