// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package beveragelist

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BeveragelistABI is the input ABI used to generate the binding from.
const BeveragelistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_dataStorage\",\"type\":\"address\"}],\"name\":\"setDataStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"},{\"name\":\"time\",\"type\":\"bytes32\"}],\"name\":\"drinkKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUserDrink\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"},{\"name\":\"time\",\"type\":\"bytes32\"}],\"name\":\"wdKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"userKey\",\"outputs\":[{\"name\":\"userID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDataStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"timestampKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"time\",\"type\":\"bytes32\"},{\"name\":\"_drink\",\"type\":\"bytes32\"},{\"name\":\"_wd\",\"type\":\"bytes32\"}],\"name\":\"setDrinkData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"time\",\"type\":\"bytes32\"}],\"name\":\"getDrinkData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"createUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"Address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"drink\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"weekday\",\"type\":\"bytes32\"}],\"name\":\"NewDrink\",\"type\":\"event\"}]"

// BeveragelistBin is the compiled bytecode used for deploying new contracts.
const BeveragelistBin = `608060405234801561001057600080fd5b5061113c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d1576000357c01000000000000000000000000000000000000000000000000000000009004806386cc31ac1161008e57806386cc31ac146102b75780638870455f14610301578063a9c06b821461034b578063b89eadb6146103a3578063c0ba228f146103e5578063cdd876181461042e576100d1565b8063047e2b27146100d65780630cdbd6371461011a5780632668f9771461017c5780632994fbe6146101a15780634209fff1146102035780637b283b2d1461025f575b600080fd5b610118600480360360208110156100ec57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061048a565b005b6101666004803603604081101561013057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104cd565b6040518082815260200191505060405180910390f35b610184610569565b604051808381526020018281526020019250505060405180910390f35b6101ed600480360360408110156101b757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610871565b6040518082815260200191505060405180910390f35b6102456004803603602081101561021957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061090d565b604051808215151515815260200191505060405180910390f35b6102a16004803603602081101561027557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a6d565b6040518082815260200191505060405180910390f35b6102bf610ad8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610309610b01565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61038d6004803603602081101561036157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b26565b6040518082815260200191505060405180910390f35b6103e3600480360360608110156103b957600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610bb9565b005b610411600480360360208110156103fb57600080fd5b8101908080359060200190929190505050610d0c565b604051808381526020018281526020019250505060405180910390f35b6104706004803603602081101561044457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ec8565b604051808215151515815260200191505060405180910390f35b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008282604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001807f6472696e6b0000000000000000000000000000000000000000000000000000008152506005019250505060405160208183030381529060405280519060200120905092915050565b60008060606000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a6fa543f6105b433610b26565b6040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060006040518083038186803b15801561060457600080fd5b505afa158015610618573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561064257600080fd5b81019080805164010000000081111561065a57600080fd5b8281019050602081018481111561067057600080fd5b815185602082028301116401000000008211171561068d57600080fd5b5050929190505050905060008160018351038151811015156106ab57fe5b90602001906020020151905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663025ec81a61070033856104cd565b6040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b15801561075057600080fd5b505afa158015610764573d6000803e3d6000fd5b505050506040513d602081101561077a57600080fd5b8101908080519060200190929190505050905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663025ec81a6107d63386610871565b6040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b15801561082657600080fd5b505afa15801561083a573d6000803e3d6000fd5b505050506040513d602081101561085057600080fd5b81019080805190602001909291905050509050818195509550505050509091565b60008282604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001807f77640000000000000000000000000000000000000000000000000000000000008152506002019250505060405160208183030381529060405280519060200120905092915050565b60008082604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f65786973747300000000000000000000000000000000000000000000000000008152506006019150506040516020818303038152906040528051906020012090506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166317e7dd22826040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b158015610a2a57600080fd5b505afa158015610a3e573d6000803e3d6000fd5b505050506040513d6020811015610a5457600080fd5b8101908080519060200190929190505050915050919050565b600081604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401915050604051602081830303815290604052805190602001209050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f74696d655374616d707300000000000000000000000000000000000000000000815250600a01915050604051602081830303815290604052805190602001209050919050565b610bc23361090d565b1515610bd357610bd133610ec8565b505b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166325cf512d610c1a33866104cd565b846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050600060405180830381600087803b158015610c7457600080fd5b505af1158015610c88573d6000803e3d6000fd5b505050507ff26bf981eed1e45a81d026c5f34958933f90c7f75cac1c362aa8a40dddfdb6cb33848484604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200194505050505060405180910390a1505050565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663025ec81a610d5833876104cd565b6040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b158015610da857600080fd5b505afa158015610dbc573d6000803e3d6000fd5b505050506040513d6020811015610dd257600080fd5b8101908080519060200190929190505050905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663025ec81a610e2e3388610871565b6040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b158015610e7e57600080fd5b505afa158015610e92573d6000803e3d6000fd5b505050506040513d6020811015610ea857600080fd5b810190808051906020019092919050505090508181935093505050915091565b6000610ed38261090d565b151515610edf57600080fd5b600033604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f65786973747300000000000000000000000000000000000000000000000000008152506006019150506040516020818303038152906040528051906020012090506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633eba9ed28260016040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018215151515815260200192505050600060405180830381600087803b15801561100a57600080fd5b505af115801561101e573d6000803e3d6000fd5b505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635a2bf25a61106885610a6d565b856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b1580156110ee57600080fd5b505af1158015611102573d6000803e3d6000fd5b50505050600191505091905056fea165627a7a723058202557027b0912ef75e9fd1b4e97175b8a88902b095ed8232b39e8fbdf53db2b200029`

// DeployBeveragelist deploys a new Ethereum contract, binding an instance of Beveragelist to it.
func DeployBeveragelist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Beveragelist, error) {
	parsed, err := abi.JSON(strings.NewReader(BeveragelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BeveragelistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Beveragelist{BeveragelistCaller: BeveragelistCaller{contract: contract}, BeveragelistTransactor: BeveragelistTransactor{contract: contract}, BeveragelistFilterer: BeveragelistFilterer{contract: contract}}, nil
}

// Beveragelist is an auto generated Go binding around an Ethereum contract.
type Beveragelist struct {
	BeveragelistCaller     // Read-only binding to the contract
	BeveragelistTransactor // Write-only binding to the contract
	BeveragelistFilterer   // Log filterer for contract events
}

// BeveragelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeveragelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeveragelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeveragelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeveragelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeveragelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeveragelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeveragelistSession struct {
	Contract     *Beveragelist     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeveragelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeveragelistCallerSession struct {
	Contract *BeveragelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BeveragelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeveragelistTransactorSession struct {
	Contract     *BeveragelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BeveragelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeveragelistRaw struct {
	Contract *Beveragelist // Generic contract binding to access the raw methods on
}

// BeveragelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeveragelistCallerRaw struct {
	Contract *BeveragelistCaller // Generic read-only contract binding to access the raw methods on
}

// BeveragelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeveragelistTransactorRaw struct {
	Contract *BeveragelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeveragelist creates a new instance of Beveragelist, bound to a specific deployed contract.
func NewBeveragelist(address common.Address, backend bind.ContractBackend) (*Beveragelist, error) {
	contract, err := bindBeveragelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Beveragelist{BeveragelistCaller: BeveragelistCaller{contract: contract}, BeveragelistTransactor: BeveragelistTransactor{contract: contract}, BeveragelistFilterer: BeveragelistFilterer{contract: contract}}, nil
}

// NewBeveragelistCaller creates a new read-only instance of Beveragelist, bound to a specific deployed contract.
func NewBeveragelistCaller(address common.Address, caller bind.ContractCaller) (*BeveragelistCaller, error) {
	contract, err := bindBeveragelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeveragelistCaller{contract: contract}, nil
}

// NewBeveragelistTransactor creates a new write-only instance of Beveragelist, bound to a specific deployed contract.
func NewBeveragelistTransactor(address common.Address, transactor bind.ContractTransactor) (*BeveragelistTransactor, error) {
	contract, err := bindBeveragelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeveragelistTransactor{contract: contract}, nil
}

// NewBeveragelistFilterer creates a new log filterer instance of Beveragelist, bound to a specific deployed contract.
func NewBeveragelistFilterer(address common.Address, filterer bind.ContractFilterer) (*BeveragelistFilterer, error) {
	contract, err := bindBeveragelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeveragelistFilterer{contract: contract}, nil
}

// bindBeveragelist binds a generic wrapper to an already deployed contract.
func bindBeveragelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BeveragelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Beveragelist *BeveragelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Beveragelist.Contract.BeveragelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Beveragelist *BeveragelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Beveragelist.Contract.BeveragelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Beveragelist *BeveragelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Beveragelist.Contract.BeveragelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Beveragelist *BeveragelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Beveragelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Beveragelist *BeveragelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Beveragelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Beveragelist *BeveragelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Beveragelist.Contract.contract.Transact(opts, method, params...)
}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() constant returns(address)
func (_Beveragelist *BeveragelistCaller) DataStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "dataStorage")
	return *ret0, err
}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() constant returns(address)
func (_Beveragelist *BeveragelistSession) DataStorage() (common.Address, error) {
	return _Beveragelist.Contract.DataStorage(&_Beveragelist.CallOpts)
}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() constant returns(address)
func (_Beveragelist *BeveragelistCallerSession) DataStorage() (common.Address, error) {
	return _Beveragelist.Contract.DataStorage(&_Beveragelist.CallOpts)
}

// DrinkKey is a free data retrieval call binding the contract method 0x0cdbd637.
//
// Solidity: function drinkKey(address userAddr, bytes32 time) constant returns(bytes32)
func (_Beveragelist *BeveragelistCaller) DrinkKey(opts *bind.CallOpts, userAddr common.Address, time [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "drinkKey", userAddr, time)
	return *ret0, err
}

// DrinkKey is a free data retrieval call binding the contract method 0x0cdbd637.
//
// Solidity: function drinkKey(address userAddr, bytes32 time) constant returns(bytes32)
func (_Beveragelist *BeveragelistSession) DrinkKey(userAddr common.Address, time [32]byte) ([32]byte, error) {
	return _Beveragelist.Contract.DrinkKey(&_Beveragelist.CallOpts, userAddr, time)
}

// DrinkKey is a free data retrieval call binding the contract method 0x0cdbd637.
//
// Solidity: function drinkKey(address userAddr, bytes32 time) constant returns(bytes32)
func (_Beveragelist *BeveragelistCallerSession) DrinkKey(userAddr common.Address, time [32]byte) ([32]byte, error) {
	return _Beveragelist.Contract.DrinkKey(&_Beveragelist.CallOpts, userAddr, time)
}

// GetDataStorage is a free data retrieval call binding the contract method 0x86cc31ac.
//
// Solidity: function getDataStorage() constant returns(address)
func (_Beveragelist *BeveragelistCaller) GetDataStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "getDataStorage")
	return *ret0, err
}

// GetDataStorage is a free data retrieval call binding the contract method 0x86cc31ac.
//
// Solidity: function getDataStorage() constant returns(address)
func (_Beveragelist *BeveragelistSession) GetDataStorage() (common.Address, error) {
	return _Beveragelist.Contract.GetDataStorage(&_Beveragelist.CallOpts)
}

// GetDataStorage is a free data retrieval call binding the contract method 0x86cc31ac.
//
// Solidity: function getDataStorage() constant returns(address)
func (_Beveragelist *BeveragelistCallerSession) GetDataStorage() (common.Address, error) {
	return _Beveragelist.Contract.GetDataStorage(&_Beveragelist.CallOpts)
}

// GetDrinkData is a free data retrieval call binding the contract method 0xc0ba228f.
//
// Solidity: function getDrinkData(bytes32 time) constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistCaller) GetDrinkData(opts *bind.CallOpts, time [32]byte) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Beveragelist.contract.Call(opts, out, "getDrinkData", time)
	return *ret0, *ret1, err
}

// GetDrinkData is a free data retrieval call binding the contract method 0xc0ba228f.
//
// Solidity: function getDrinkData(bytes32 time) constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistSession) GetDrinkData(time [32]byte) ([32]byte, [32]byte, error) {
	return _Beveragelist.Contract.GetDrinkData(&_Beveragelist.CallOpts, time)
}

// GetDrinkData is a free data retrieval call binding the contract method 0xc0ba228f.
//
// Solidity: function getDrinkData(bytes32 time) constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistCallerSession) GetDrinkData(time [32]byte) ([32]byte, [32]byte, error) {
	return _Beveragelist.Contract.GetDrinkData(&_Beveragelist.CallOpts, time)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddr) constant returns(bool isIndeed)
func (_Beveragelist *BeveragelistCaller) IsUser(opts *bind.CallOpts, userAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "isUser", userAddr)
	return *ret0, err
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddr) constant returns(bool isIndeed)
func (_Beveragelist *BeveragelistSession) IsUser(userAddr common.Address) (bool, error) {
	return _Beveragelist.Contract.IsUser(&_Beveragelist.CallOpts, userAddr)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddr) constant returns(bool isIndeed)
func (_Beveragelist *BeveragelistCallerSession) IsUser(userAddr common.Address) (bool, error) {
	return _Beveragelist.Contract.IsUser(&_Beveragelist.CallOpts, userAddr)
}

// LastUserDrink is a free data retrieval call binding the contract method 0x2668f977.
//
// Solidity: function lastUserDrink() constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistCaller) LastUserDrink(opts *bind.CallOpts) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Beveragelist.contract.Call(opts, out, "lastUserDrink")
	return *ret0, *ret1, err
}

// LastUserDrink is a free data retrieval call binding the contract method 0x2668f977.
//
// Solidity: function lastUserDrink() constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistSession) LastUserDrink() ([32]byte, [32]byte, error) {
	return _Beveragelist.Contract.LastUserDrink(&_Beveragelist.CallOpts)
}

// LastUserDrink is a free data retrieval call binding the contract method 0x2668f977.
//
// Solidity: function lastUserDrink() constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistCallerSession) LastUserDrink() ([32]byte, [32]byte, error) {
	return _Beveragelist.Contract.LastUserDrink(&_Beveragelist.CallOpts)
}

// TimestampKey is a free data retrieval call binding the contract method 0xa9c06b82.
//
// Solidity: function timestampKey(address userAddr) constant returns(bytes32)
func (_Beveragelist *BeveragelistCaller) TimestampKey(opts *bind.CallOpts, userAddr common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "timestampKey", userAddr)
	return *ret0, err
}

// TimestampKey is a free data retrieval call binding the contract method 0xa9c06b82.
//
// Solidity: function timestampKey(address userAddr) constant returns(bytes32)
func (_Beveragelist *BeveragelistSession) TimestampKey(userAddr common.Address) ([32]byte, error) {
	return _Beveragelist.Contract.TimestampKey(&_Beveragelist.CallOpts, userAddr)
}

// TimestampKey is a free data retrieval call binding the contract method 0xa9c06b82.
//
// Solidity: function timestampKey(address userAddr) constant returns(bytes32)
func (_Beveragelist *BeveragelistCallerSession) TimestampKey(userAddr common.Address) ([32]byte, error) {
	return _Beveragelist.Contract.TimestampKey(&_Beveragelist.CallOpts, userAddr)
}

// UserKey is a free data retrieval call binding the contract method 0x7b283b2d.
//
// Solidity: function userKey(address userAddr) constant returns(bytes32 userID)
func (_Beveragelist *BeveragelistCaller) UserKey(opts *bind.CallOpts, userAddr common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "userKey", userAddr)
	return *ret0, err
}

// UserKey is a free data retrieval call binding the contract method 0x7b283b2d.
//
// Solidity: function userKey(address userAddr) constant returns(bytes32 userID)
func (_Beveragelist *BeveragelistSession) UserKey(userAddr common.Address) ([32]byte, error) {
	return _Beveragelist.Contract.UserKey(&_Beveragelist.CallOpts, userAddr)
}

// UserKey is a free data retrieval call binding the contract method 0x7b283b2d.
//
// Solidity: function userKey(address userAddr) constant returns(bytes32 userID)
func (_Beveragelist *BeveragelistCallerSession) UserKey(userAddr common.Address) ([32]byte, error) {
	return _Beveragelist.Contract.UserKey(&_Beveragelist.CallOpts, userAddr)
}

// WdKey is a free data retrieval call binding the contract method 0x2994fbe6.
//
// Solidity: function wdKey(address userAddr, bytes32 time) constant returns(bytes32)
func (_Beveragelist *BeveragelistCaller) WdKey(opts *bind.CallOpts, userAddr common.Address, time [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "wdKey", userAddr, time)
	return *ret0, err
}

// WdKey is a free data retrieval call binding the contract method 0x2994fbe6.
//
// Solidity: function wdKey(address userAddr, bytes32 time) constant returns(bytes32)
func (_Beveragelist *BeveragelistSession) WdKey(userAddr common.Address, time [32]byte) ([32]byte, error) {
	return _Beveragelist.Contract.WdKey(&_Beveragelist.CallOpts, userAddr, time)
}

// WdKey is a free data retrieval call binding the contract method 0x2994fbe6.
//
// Solidity: function wdKey(address userAddr, bytes32 time) constant returns(bytes32)
func (_Beveragelist *BeveragelistCallerSession) WdKey(userAddr common.Address, time [32]byte) ([32]byte, error) {
	return _Beveragelist.Contract.WdKey(&_Beveragelist.CallOpts, userAddr, time)
}

// CreateUser is a paid mutator transaction binding the contract method 0xcdd87618.
//
// Solidity: function createUser(address userAddr) returns(bool success)
func (_Beveragelist *BeveragelistTransactor) CreateUser(opts *bind.TransactOpts, userAddr common.Address) (*types.Transaction, error) {
	return _Beveragelist.contract.Transact(opts, "createUser", userAddr)
}

// CreateUser is a paid mutator transaction binding the contract method 0xcdd87618.
//
// Solidity: function createUser(address userAddr) returns(bool success)
func (_Beveragelist *BeveragelistSession) CreateUser(userAddr common.Address) (*types.Transaction, error) {
	return _Beveragelist.Contract.CreateUser(&_Beveragelist.TransactOpts, userAddr)
}

// CreateUser is a paid mutator transaction binding the contract method 0xcdd87618.
//
// Solidity: function createUser(address userAddr) returns(bool success)
func (_Beveragelist *BeveragelistTransactorSession) CreateUser(userAddr common.Address) (*types.Transaction, error) {
	return _Beveragelist.Contract.CreateUser(&_Beveragelist.TransactOpts, userAddr)
}

// SetDataStorage is a paid mutator transaction binding the contract method 0x047e2b27.
//
// Solidity: function setDataStorage(address _dataStorage) returns()
func (_Beveragelist *BeveragelistTransactor) SetDataStorage(opts *bind.TransactOpts, _dataStorage common.Address) (*types.Transaction, error) {
	return _Beveragelist.contract.Transact(opts, "setDataStorage", _dataStorage)
}

// SetDataStorage is a paid mutator transaction binding the contract method 0x047e2b27.
//
// Solidity: function setDataStorage(address _dataStorage) returns()
func (_Beveragelist *BeveragelistSession) SetDataStorage(_dataStorage common.Address) (*types.Transaction, error) {
	return _Beveragelist.Contract.SetDataStorage(&_Beveragelist.TransactOpts, _dataStorage)
}

// SetDataStorage is a paid mutator transaction binding the contract method 0x047e2b27.
//
// Solidity: function setDataStorage(address _dataStorage) returns()
func (_Beveragelist *BeveragelistTransactorSession) SetDataStorage(_dataStorage common.Address) (*types.Transaction, error) {
	return _Beveragelist.Contract.SetDataStorage(&_Beveragelist.TransactOpts, _dataStorage)
}

// SetDrinkData is a paid mutator transaction binding the contract method 0xb89eadb6.
//
// Solidity: function setDrinkData(bytes32 time, bytes32 _drink, bytes32 _wd) returns()
func (_Beveragelist *BeveragelistTransactor) SetDrinkData(opts *bind.TransactOpts, time [32]byte, _drink [32]byte, _wd [32]byte) (*types.Transaction, error) {
	return _Beveragelist.contract.Transact(opts, "setDrinkData", time, _drink, _wd)
}

// SetDrinkData is a paid mutator transaction binding the contract method 0xb89eadb6.
//
// Solidity: function setDrinkData(bytes32 time, bytes32 _drink, bytes32 _wd) returns()
func (_Beveragelist *BeveragelistSession) SetDrinkData(time [32]byte, _drink [32]byte, _wd [32]byte) (*types.Transaction, error) {
	return _Beveragelist.Contract.SetDrinkData(&_Beveragelist.TransactOpts, time, _drink, _wd)
}

// SetDrinkData is a paid mutator transaction binding the contract method 0xb89eadb6.
//
// Solidity: function setDrinkData(bytes32 time, bytes32 _drink, bytes32 _wd) returns()
func (_Beveragelist *BeveragelistTransactorSession) SetDrinkData(time [32]byte, _drink [32]byte, _wd [32]byte) (*types.Transaction, error) {
	return _Beveragelist.Contract.SetDrinkData(&_Beveragelist.TransactOpts, time, _drink, _wd)
}

// BeveragelistNewDrinkIterator is returned from FilterNewDrink and is used to iterate over the raw logs and unpacked data for NewDrink events raised by the Beveragelist contract.
type BeveragelistNewDrinkIterator struct {
	Event *BeveragelistNewDrink // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BeveragelistNewDrinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeveragelistNewDrink)
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
		it.Event = new(BeveragelistNewDrink)
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
func (it *BeveragelistNewDrinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeveragelistNewDrinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeveragelistNewDrink represents a NewDrink event raised by the Beveragelist contract.
type BeveragelistNewDrink struct {
	Address common.Address
	Time    [32]byte
	Drink   [32]byte
	Weekday [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewDrink is a free log retrieval operation binding the contract event 0xf26bf981eed1e45a81d026c5f34958933f90c7f75cac1c362aa8a40dddfdb6cb.
//
// Solidity: event NewDrink(address Address, bytes32 time, bytes32 drink, bytes32 weekday)
func (_Beveragelist *BeveragelistFilterer) FilterNewDrink(opts *bind.FilterOpts) (*BeveragelistNewDrinkIterator, error) {

	logs, sub, err := _Beveragelist.contract.FilterLogs(opts, "NewDrink")
	if err != nil {
		return nil, err
	}
	return &BeveragelistNewDrinkIterator{contract: _Beveragelist.contract, event: "NewDrink", logs: logs, sub: sub}, nil
}

// WatchNewDrink is a free log subscription operation binding the contract event 0xf26bf981eed1e45a81d026c5f34958933f90c7f75cac1c362aa8a40dddfdb6cb.
//
// Solidity: event NewDrink(address Address, bytes32 time, bytes32 drink, bytes32 weekday)
func (_Beveragelist *BeveragelistFilterer) WatchNewDrink(opts *bind.WatchOpts, sink chan<- *BeveragelistNewDrink) (event.Subscription, error) {

	logs, sub, err := _Beveragelist.contract.WatchLogs(opts, "NewDrink")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeveragelistNewDrink)
				if err := _Beveragelist.contract.UnpackLog(event, "NewDrink", log); err != nil {
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
