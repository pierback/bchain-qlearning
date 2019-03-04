// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coffeecoin

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CoffeecoinABI is the input ABI used to generate the binding from.
const CoffeecoinABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payMate\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payCoffee\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCoffeePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMatePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChairBalance\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"accounts\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"initialCredit\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnBalance\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWaterPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setWaterPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payWater\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setMatePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setCoffeePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_chairAddress\",\"type\":\"address\"},{\"name\":\"cp\",\"type\":\"uint256\"},{\"name\":\"mp\",\"type\":\"uint256\"},{\"name\":\"wp\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// CoffeecoinBin is the compiled bytecode used for deploying new contracts.
const CoffeecoinBin = `6080604052604051608080610d9a83398101806040528101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061009c836100d5640100000000026401000000009004565b6100b4826100df640100000000026401000000009004565b6100cc816100d5640100000000026401000000009004565b505050506100e9565b8060038190555050565b8060018190555050565b610ca2806100f86000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063095ea7b3146100f65780630bcd8bf91461015b57806311d40e631461018a578063126f4275146101b957806323b872dd146101e45780632cf0bb8c146102695780634103f8c3146102945780635c658165146102bf5780635e5c06e21461033657806370a08231146103985780637130474e146103ef5780637d848fec1461041a57806386712b1314610445578063a9059cbb14610472578063d5ee2785146104d7578063e146110814610506578063f847dff714610533575b600080fd5b34801561010257600080fd5b50610141600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610560565b604051808215151515815260200191505060405180910390f35b34801561016757600080fd5b506101706105ed565b604051808215151515815260200191505060405180910390f35b34801561019657600080fd5b5061019f610622565b604051808215151515815260200191505060405180910390f35b3480156101c557600080fd5b506101ce610657565b6040518082815260200191505060405180910390f35b3480156101f057600080fd5b5061024f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610661565b604051808215151515815260200191505060405180910390f35b34801561027557600080fd5b5061027e6108df565b6040518082815260200191505060405180910390f35b3480156102a057600080fd5b506102a96108e9565b6040518082815260200191505060405180910390f35b3480156102cb57600080fd5b50610320600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061091a565b6040518082815260200191505060405180910390f35b34801561034257600080fd5b50610377600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061093f565b60405180838152602001821515151581526020019250505060405180910390f35b3480156103a457600080fd5b506103d9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610970565b6040518082815260200191505060405180910390f35b3480156103fb57600080fd5b506104046109bc565b6040518082815260200191505060405180910390f35b34801561042657600080fd5b5061042f610a06565b6040518082815260200191505060405180910390f35b34801561045157600080fd5b5061047060048036038101908080359060200190929190505050610a10565b005b34801561047e57600080fd5b506104bd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a1a565b604051808215151515815260200191505060405180910390f35b3480156104e357600080fd5b506104ec610b82565b604051808215151515815260200191505060405180910390f35b34801561051257600080fd5b5061053160048036038101908080359060200190929190505050610bb7565b005b34801561053f57600080fd5b5061055e60048036038101908080359060200190929190505050610bc1565b005b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001905092915050565b600061061d336000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600254610661565b905090565b6000610652336000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600154610661565b905090565b6000600154905090565b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff1615156106ce576106c1610bcb565b6106cc336014610560565b505b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154821115151561071f57600080fd5b600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156107aa57600080fd5b81600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254039250508190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254019250508190555081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550600190509392505050565b6000600254905090565b60006109156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610970565b905090565b6005602052816000526040600020602052806000526040600020600091509150505481565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16905082565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050919050565b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154905090565b6000600354905090565b8060038190555050565b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff161515610a8757610a7a610bcb565b610a85336014610560565b505b81600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410151515610ad857600080fd5b81600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254039250508190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600082825401925050819055506001905092915050565b6000610bb2336000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600354610661565b905090565b8060028190555050565b8060018190555050565b66038d7ea4c68000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff0219169083151502179055505600a165627a7a72305820ad7c66e0bca2b775b994164c2fcdf7cc80aebaa0677bc97607fd3f09691006100029`

// DeployCoffeecoin deploys a new Ethereum contract, binding an instance of Coffeecoin to it.
func DeployCoffeecoin(auth *bind.TransactOpts, backend bind.ContractBackend, _chairAddress common.Address, cp *big.Int, mp *big.Int, wp *big.Int) (common.Address, *types.Transaction, *Coffeecoin, error) {
	parsed, err := abi.JSON(strings.NewReader(CoffeecoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoffeecoinBin), backend, _chairAddress, cp, mp, wp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Coffeecoin{CoffeecoinCaller: CoffeecoinCaller{contract: contract}, CoffeecoinTransactor: CoffeecoinTransactor{contract: contract}, CoffeecoinFilterer: CoffeecoinFilterer{contract: contract}}, nil
}

// Coffeecoin is an auto generated Go binding around an Ethereum contract.
type Coffeecoin struct {
	CoffeecoinCaller     // Read-only binding to the contract
	CoffeecoinTransactor // Write-only binding to the contract
	CoffeecoinFilterer   // Log filterer for contract events
}

// CoffeecoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoffeecoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeecoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoffeecoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeecoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoffeecoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeecoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoffeecoinSession struct {
	Contract     *Coffeecoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoffeecoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoffeecoinCallerSession struct {
	Contract *CoffeecoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CoffeecoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoffeecoinTransactorSession struct {
	Contract     *CoffeecoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CoffeecoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoffeecoinRaw struct {
	Contract *Coffeecoin // Generic contract binding to access the raw methods on
}

// CoffeecoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoffeecoinCallerRaw struct {
	Contract *CoffeecoinCaller // Generic read-only contract binding to access the raw methods on
}

// CoffeecoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoffeecoinTransactorRaw struct {
	Contract *CoffeecoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoffeecoin creates a new instance of Coffeecoin, bound to a specific deployed contract.
func NewCoffeecoin(address common.Address, backend bind.ContractBackend) (*Coffeecoin, error) {
	contract, err := bindCoffeecoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coffeecoin{CoffeecoinCaller: CoffeecoinCaller{contract: contract}, CoffeecoinTransactor: CoffeecoinTransactor{contract: contract}, CoffeecoinFilterer: CoffeecoinFilterer{contract: contract}}, nil
}

// NewCoffeecoinCaller creates a new read-only instance of Coffeecoin, bound to a specific deployed contract.
func NewCoffeecoinCaller(address common.Address, caller bind.ContractCaller) (*CoffeecoinCaller, error) {
	contract, err := bindCoffeecoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoffeecoinCaller{contract: contract}, nil
}

// NewCoffeecoinTransactor creates a new write-only instance of Coffeecoin, bound to a specific deployed contract.
func NewCoffeecoinTransactor(address common.Address, transactor bind.ContractTransactor) (*CoffeecoinTransactor, error) {
	contract, err := bindCoffeecoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoffeecoinTransactor{contract: contract}, nil
}

// NewCoffeecoinFilterer creates a new log filterer instance of Coffeecoin, bound to a specific deployed contract.
func NewCoffeecoinFilterer(address common.Address, filterer bind.ContractFilterer) (*CoffeecoinFilterer, error) {
	contract, err := bindCoffeecoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoffeecoinFilterer{contract: contract}, nil
}

// bindCoffeecoin binds a generic wrapper to an already deployed contract.
func bindCoffeecoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoffeecoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coffeecoin *CoffeecoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coffeecoin.Contract.CoffeecoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coffeecoin *CoffeecoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeecoin.Contract.CoffeecoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coffeecoin *CoffeecoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coffeecoin.Contract.CoffeecoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coffeecoin *CoffeecoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coffeecoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coffeecoin *CoffeecoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeecoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coffeecoin *CoffeecoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coffeecoin.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts( address) constant returns(balance uint256, initialCredit bool)
func (_Coffeecoin *CoffeecoinCaller) Accounts(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance       *big.Int
	InitialCredit bool
}, error) {
	ret := new(struct {
		Balance       *big.Int
		InitialCredit bool
	})
	out := ret
	err := _Coffeecoin.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts( address) constant returns(balance uint256, initialCredit bool)
func (_Coffeecoin *CoffeecoinSession) Accounts(arg0 common.Address) (struct {
	Balance       *big.Int
	InitialCredit bool
}, error) {
	return _Coffeecoin.Contract.Accounts(&_Coffeecoin.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts( address) constant returns(balance uint256, initialCredit bool)
func (_Coffeecoin *CoffeecoinCallerSession) Accounts(arg0 common.Address) (struct {
	Balance       *big.Int
	InitialCredit bool
}, error) {
	return _Coffeecoin.Contract.Accounts(&_Coffeecoin.CallOpts, arg0)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_Coffeecoin *CoffeecoinCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeecoin.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_Coffeecoin *CoffeecoinSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Coffeecoin.Contract.Allowed(&_Coffeecoin.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_Coffeecoin *CoffeecoinCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Coffeecoin.Contract.Allowed(&_Coffeecoin.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeecoin.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Coffeecoin.Contract.BalanceOf(&_Coffeecoin.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Coffeecoin.Contract.BalanceOf(&_Coffeecoin.CallOpts, tokenOwner)
}

// GetChairBalance is a free data retrieval call binding the contract method 0x4103f8c3.
//
// Solidity: function getChairBalance() constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinCaller) GetChairBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeecoin.contract.Call(opts, out, "getChairBalance")
	return *ret0, err
}

// GetChairBalance is a free data retrieval call binding the contract method 0x4103f8c3.
//
// Solidity: function getChairBalance() constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinSession) GetChairBalance() (*big.Int, error) {
	return _Coffeecoin.Contract.GetChairBalance(&_Coffeecoin.CallOpts)
}

// GetChairBalance is a free data retrieval call binding the contract method 0x4103f8c3.
//
// Solidity: function getChairBalance() constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinCallerSession) GetChairBalance() (*big.Int, error) {
	return _Coffeecoin.Contract.GetChairBalance(&_Coffeecoin.CallOpts)
}

// GetCoffeePrice is a free data retrieval call binding the contract method 0x126f4275.
//
// Solidity: function getCoffeePrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinCaller) GetCoffeePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeecoin.contract.Call(opts, out, "getCoffeePrice")
	return *ret0, err
}

// GetCoffeePrice is a free data retrieval call binding the contract method 0x126f4275.
//
// Solidity: function getCoffeePrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinSession) GetCoffeePrice() (*big.Int, error) {
	return _Coffeecoin.Contract.GetCoffeePrice(&_Coffeecoin.CallOpts)
}

// GetCoffeePrice is a free data retrieval call binding the contract method 0x126f4275.
//
// Solidity: function getCoffeePrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinCallerSession) GetCoffeePrice() (*big.Int, error) {
	return _Coffeecoin.Contract.GetCoffeePrice(&_Coffeecoin.CallOpts)
}

// GetMatePrice is a free data retrieval call binding the contract method 0x2cf0bb8c.
//
// Solidity: function getMatePrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinCaller) GetMatePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeecoin.contract.Call(opts, out, "getMatePrice")
	return *ret0, err
}

// GetMatePrice is a free data retrieval call binding the contract method 0x2cf0bb8c.
//
// Solidity: function getMatePrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinSession) GetMatePrice() (*big.Int, error) {
	return _Coffeecoin.Contract.GetMatePrice(&_Coffeecoin.CallOpts)
}

// GetMatePrice is a free data retrieval call binding the contract method 0x2cf0bb8c.
//
// Solidity: function getMatePrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinCallerSession) GetMatePrice() (*big.Int, error) {
	return _Coffeecoin.Contract.GetMatePrice(&_Coffeecoin.CallOpts)
}

// GetOwnBalance is a free data retrieval call binding the contract method 0x7130474e.
//
// Solidity: function getOwnBalance() constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinCaller) GetOwnBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeecoin.contract.Call(opts, out, "getOwnBalance")
	return *ret0, err
}

// GetOwnBalance is a free data retrieval call binding the contract method 0x7130474e.
//
// Solidity: function getOwnBalance() constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinSession) GetOwnBalance() (*big.Int, error) {
	return _Coffeecoin.Contract.GetOwnBalance(&_Coffeecoin.CallOpts)
}

// GetOwnBalance is a free data retrieval call binding the contract method 0x7130474e.
//
// Solidity: function getOwnBalance() constant returns(balance uint256)
func (_Coffeecoin *CoffeecoinCallerSession) GetOwnBalance() (*big.Int, error) {
	return _Coffeecoin.Contract.GetOwnBalance(&_Coffeecoin.CallOpts)
}

// GetWaterPrice is a free data retrieval call binding the contract method 0x7d848fec.
//
// Solidity: function getWaterPrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinCaller) GetWaterPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeecoin.contract.Call(opts, out, "getWaterPrice")
	return *ret0, err
}

// GetWaterPrice is a free data retrieval call binding the contract method 0x7d848fec.
//
// Solidity: function getWaterPrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinSession) GetWaterPrice() (*big.Int, error) {
	return _Coffeecoin.Contract.GetWaterPrice(&_Coffeecoin.CallOpts)
}

// GetWaterPrice is a free data retrieval call binding the contract method 0x7d848fec.
//
// Solidity: function getWaterPrice() constant returns(uint256)
func (_Coffeecoin *CoffeecoinCallerSession) GetWaterPrice() (*big.Int, error) {
	return _Coffeecoin.Contract.GetWaterPrice(&_Coffeecoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.Approve(&_Coffeecoin.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.Approve(&_Coffeecoin.TransactOpts, spender, tokens)
}

// PayCoffee is a paid mutator transaction binding the contract method 0x11d40e63.
//
// Solidity: function payCoffee() returns(success bool)
func (_Coffeecoin *CoffeecoinTransactor) PayCoffee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "payCoffee")
}

// PayCoffee is a paid mutator transaction binding the contract method 0x11d40e63.
//
// Solidity: function payCoffee() returns(success bool)
func (_Coffeecoin *CoffeecoinSession) PayCoffee() (*types.Transaction, error) {
	return _Coffeecoin.Contract.PayCoffee(&_Coffeecoin.TransactOpts)
}

// PayCoffee is a paid mutator transaction binding the contract method 0x11d40e63.
//
// Solidity: function payCoffee() returns(success bool)
func (_Coffeecoin *CoffeecoinTransactorSession) PayCoffee() (*types.Transaction, error) {
	return _Coffeecoin.Contract.PayCoffee(&_Coffeecoin.TransactOpts)
}

// PayMate is a paid mutator transaction binding the contract method 0x0bcd8bf9.
//
// Solidity: function payMate() returns(success bool)
func (_Coffeecoin *CoffeecoinTransactor) PayMate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "payMate")
}

// PayMate is a paid mutator transaction binding the contract method 0x0bcd8bf9.
//
// Solidity: function payMate() returns(success bool)
func (_Coffeecoin *CoffeecoinSession) PayMate() (*types.Transaction, error) {
	return _Coffeecoin.Contract.PayMate(&_Coffeecoin.TransactOpts)
}

// PayMate is a paid mutator transaction binding the contract method 0x0bcd8bf9.
//
// Solidity: function payMate() returns(success bool)
func (_Coffeecoin *CoffeecoinTransactorSession) PayMate() (*types.Transaction, error) {
	return _Coffeecoin.Contract.PayMate(&_Coffeecoin.TransactOpts)
}

// PayWater is a paid mutator transaction binding the contract method 0xd5ee2785.
//
// Solidity: function payWater() returns(success bool)
func (_Coffeecoin *CoffeecoinTransactor) PayWater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "payWater")
}

// PayWater is a paid mutator transaction binding the contract method 0xd5ee2785.
//
// Solidity: function payWater() returns(success bool)
func (_Coffeecoin *CoffeecoinSession) PayWater() (*types.Transaction, error) {
	return _Coffeecoin.Contract.PayWater(&_Coffeecoin.TransactOpts)
}

// PayWater is a paid mutator transaction binding the contract method 0xd5ee2785.
//
// Solidity: function payWater() returns(success bool)
func (_Coffeecoin *CoffeecoinTransactorSession) PayWater() (*types.Transaction, error) {
	return _Coffeecoin.Contract.PayWater(&_Coffeecoin.TransactOpts)
}

// SetCoffeePrice is a paid mutator transaction binding the contract method 0xf847dff7.
//
// Solidity: function setCoffeePrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinTransactor) SetCoffeePrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "setCoffeePrice", price)
}

// SetCoffeePrice is a paid mutator transaction binding the contract method 0xf847dff7.
//
// Solidity: function setCoffeePrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinSession) SetCoffeePrice(price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.SetCoffeePrice(&_Coffeecoin.TransactOpts, price)
}

// SetCoffeePrice is a paid mutator transaction binding the contract method 0xf847dff7.
//
// Solidity: function setCoffeePrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinTransactorSession) SetCoffeePrice(price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.SetCoffeePrice(&_Coffeecoin.TransactOpts, price)
}

// SetMatePrice is a paid mutator transaction binding the contract method 0xe1461108.
//
// Solidity: function setMatePrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinTransactor) SetMatePrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "setMatePrice", price)
}

// SetMatePrice is a paid mutator transaction binding the contract method 0xe1461108.
//
// Solidity: function setMatePrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinSession) SetMatePrice(price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.SetMatePrice(&_Coffeecoin.TransactOpts, price)
}

// SetMatePrice is a paid mutator transaction binding the contract method 0xe1461108.
//
// Solidity: function setMatePrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinTransactorSession) SetMatePrice(price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.SetMatePrice(&_Coffeecoin.TransactOpts, price)
}

// SetWaterPrice is a paid mutator transaction binding the contract method 0x86712b13.
//
// Solidity: function setWaterPrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinTransactor) SetWaterPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "setWaterPrice", price)
}

// SetWaterPrice is a paid mutator transaction binding the contract method 0x86712b13.
//
// Solidity: function setWaterPrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinSession) SetWaterPrice(price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.SetWaterPrice(&_Coffeecoin.TransactOpts, price)
}

// SetWaterPrice is a paid mutator transaction binding the contract method 0x86712b13.
//
// Solidity: function setWaterPrice(price uint256) returns()
func (_Coffeecoin *CoffeecoinTransactorSession) SetWaterPrice(price *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.SetWaterPrice(&_Coffeecoin.TransactOpts, price)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.Transfer(&_Coffeecoin.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.Transfer(&_Coffeecoin.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.TransferFrom(&_Coffeecoin.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Coffeecoin *CoffeecoinTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Coffeecoin.Contract.TransferFrom(&_Coffeecoin.TransactOpts, from, to, tokens)
}
