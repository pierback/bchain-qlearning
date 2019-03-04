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

// BeveragelistABI is the input ABI used to generate the binding from.
const BeveragelistABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUserDrink\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDrink\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"time\",\"type\":\"bytes32\"},{\"name\":\"_drink\",\"type\":\"bytes32\"},{\"name\":\"_wd\",\"type\":\"bytes32\"}],\"name\":\"setDrinkData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"time\",\"type\":\"bytes32\"}],\"name\":\"getDrinkData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"Address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"drink\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"weekday\",\"type\":\"bytes32\"}],\"name\":\"NewDrink\",\"type\":\"event\"}]"

// BeveragelistBin is the compiled bytecode used for deploying new contracts.
const BeveragelistBin = `608060405234801561001057600080fd5b50610a55806100206000396000f300608060405260043610610077576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062ce8e3e1461007c57806307973ccf146100e85780632668f97714610113578063962e88fb14610155578063b89eadb6146101ca578063c0ba228f14610217575b600080fd5b34801561008857600080fd5b50610091610273565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100d45780820151818401526020810190506100b9565b505050509050019250505060405180910390f35b3480156100f457600080fd5b506100fd610301565b6040518082815260200191505060405180910390f35b34801561011f57600080fd5b5061012861030e565b60405180836000191660001916815260200182600019166000191681526020019250505060405180910390f35b34801561016157600080fd5b5061016a61045a565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183600019166000191681526020018260001916600019168152602001935050505060405180910390f35b3480156101d657600080fd5b50610215600480360381019080803560001916906020019092919080356000191690602001909291908035600019169060200190929190505050610555565b005b34801561022357600080fd5b506102466004803603810190808035600019169060200190929190505050610901565b60405180836000191660001916815260200182600019166000191681526020019250505060405180910390f35b606060028054806020026020016040519081016040528092919081815260200182805480156102f757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116102ad575b5050505050905090565b6000600280549050905090565b600080600061031b6109a8565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010180549050038154811015156103b157fe5b906000526020600020015491506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836000191660001916815260200190815260200160002060408051908101604052908160008201546000191660001916815260200160018201546000191660001916815250509050806000015181602001519350935050509091565b60008060006104676109c8565b600360016003805490500381548110151561047e57fe5b9060005260206000209060040201606060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201604080519081016040529081600082015460001916600019168152602001600182015460001916600019168152505081526020016003820154600019166000191681525050905033816040015182602001516000015193509350935050909192565b61055d6109a8565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1615156106745760018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff02191690831515021790555060023390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b60408051908101604052808460001916815260200183600019168152509050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086600019166000191681526020019081526020016000206000820151816000019060001916905560208201518160010190600019169055905050600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010184908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555060036060604051908101604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200186600019168152509080600181540180825580915050906001820390600052602060002090600402016000909192909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160008201518160000190600019169055602082015181600101906000191690555050604082015181600301906000191690555050507ff26bf981eed1e45a81d026c5f34958933f90c7f75cac1c362aa8a40dddfdb6cb33858585604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184600019166000191681526020018360001916600019168152602001826000191660001916815260200194505050505060405180910390a150505050565b60008061090c6109a8565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000856000191660001916815260200190815260200160002060408051908101604052908160008201546000191660001916815260200160018201546000191660001916815250509050806000015181602001519250925050915091565b604080519081016040528060008019168152602001600080191681525090565b608060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016109f9610a09565b8152602001600080191681525090565b6040805190810160405280600080191681526020016000801916815250905600a165627a7a72305820080f526a486328810ae36cac861b152c5f45bef6dbc4233b941ed2cbbbca700a0029`

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

// GetDrinkData is a free data retrieval call binding the contract method 0xc0ba228f.
//
// Solidity: function getDrinkData(time bytes32) constant returns(bytes32, bytes32)
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
// Solidity: function getDrinkData(time bytes32) constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistSession) GetDrinkData(time [32]byte) ([32]byte, [32]byte, error) {
	return _Beveragelist.Contract.GetDrinkData(&_Beveragelist.CallOpts, time)
}

// GetDrinkData is a free data retrieval call binding the contract method 0xc0ba228f.
//
// Solidity: function getDrinkData(time bytes32) constant returns(bytes32, bytes32)
func (_Beveragelist *BeveragelistCallerSession) GetDrinkData(time [32]byte) ([32]byte, [32]byte, error) {
	return _Beveragelist.Contract.GetDrinkData(&_Beveragelist.CallOpts, time)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(address[])
func (_Beveragelist *BeveragelistCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "getUsers")
	return *ret0, err
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(address[])
func (_Beveragelist *BeveragelistSession) GetUsers() ([]common.Address, error) {
	return _Beveragelist.Contract.GetUsers(&_Beveragelist.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(address[])
func (_Beveragelist *BeveragelistCallerSession) GetUsers() ([]common.Address, error) {
	return _Beveragelist.Contract.GetUsers(&_Beveragelist.CallOpts)
}

// LastDrink is a free data retrieval call binding the contract method 0x962e88fb.
//
// Solidity: function lastDrink() constant returns(address, bytes32, bytes32)
func (_Beveragelist *BeveragelistCaller) LastDrink(opts *bind.CallOpts) (common.Address, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Beveragelist.contract.Call(opts, out, "lastDrink")
	return *ret0, *ret1, *ret2, err
}

// LastDrink is a free data retrieval call binding the contract method 0x962e88fb.
//
// Solidity: function lastDrink() constant returns(address, bytes32, bytes32)
func (_Beveragelist *BeveragelistSession) LastDrink() (common.Address, [32]byte, [32]byte, error) {
	return _Beveragelist.Contract.LastDrink(&_Beveragelist.CallOpts)
}

// LastDrink is a free data retrieval call binding the contract method 0x962e88fb.
//
// Solidity: function lastDrink() constant returns(address, bytes32, bytes32)
func (_Beveragelist *BeveragelistCallerSession) LastDrink() (common.Address, [32]byte, [32]byte, error) {
	return _Beveragelist.Contract.LastDrink(&_Beveragelist.CallOpts)
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

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() constant returns(uint256)
func (_Beveragelist *BeveragelistCaller) UserCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Beveragelist.contract.Call(opts, out, "userCount")
	return *ret0, err
}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() constant returns(uint256)
func (_Beveragelist *BeveragelistSession) UserCount() (*big.Int, error) {
	return _Beveragelist.Contract.UserCount(&_Beveragelist.CallOpts)
}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() constant returns(uint256)
func (_Beveragelist *BeveragelistCallerSession) UserCount() (*big.Int, error) {
	return _Beveragelist.Contract.UserCount(&_Beveragelist.CallOpts)
}

// SetDrinkData is a paid mutator transaction binding the contract method 0xb89eadb6.
//
// Solidity: function setDrinkData(time bytes32, _drink bytes32, _wd bytes32) returns()
func (_Beveragelist *BeveragelistTransactor) SetDrinkData(opts *bind.TransactOpts, time [32]byte, _drink [32]byte, _wd [32]byte) (*types.Transaction, error) {
	return _Beveragelist.contract.Transact(opts, "setDrinkData", time, _drink, _wd)
}

// SetDrinkData is a paid mutator transaction binding the contract method 0xb89eadb6.
//
// Solidity: function setDrinkData(time bytes32, _drink bytes32, _wd bytes32) returns()
func (_Beveragelist *BeveragelistSession) SetDrinkData(time [32]byte, _drink [32]byte, _wd [32]byte) (*types.Transaction, error) {
	return _Beveragelist.Contract.SetDrinkData(&_Beveragelist.TransactOpts, time, _drink, _wd)
}

// SetDrinkData is a paid mutator transaction binding the contract method 0xb89eadb6.
//
// Solidity: function setDrinkData(time bytes32, _drink bytes32, _wd bytes32) returns()
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
// Solidity: e NewDrink(Address address, time bytes32, drink bytes32, weekday bytes32)
func (_Beveragelist *BeveragelistFilterer) FilterNewDrink(opts *bind.FilterOpts) (*BeveragelistNewDrinkIterator, error) {

	logs, sub, err := _Beveragelist.contract.FilterLogs(opts, "NewDrink")
	if err != nil {
		return nil, err
	}
	return &BeveragelistNewDrinkIterator{contract: _Beveragelist.contract, event: "NewDrink", logs: logs, sub: sub}, nil
}

// WatchNewDrink is a free log subscription operation binding the contract event 0xf26bf981eed1e45a81d026c5f34958933f90c7f75cac1c362aa8a40dddfdb6cb.
//
// Solidity: e NewDrink(Address address, time bytes32, drink bytes32, weekday bytes32)
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
