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
const BeveragelistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"time\",\"type\":\"bytes32\"},{\"name\":\"_drink\",\"type\":\"bytes32\"},{\"name\":\"_wd\",\"type\":\"bytes32\"}],\"name\":\"setDrinkData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"time\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"drink\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"weekday\",\"type\":\"bytes32\"}],\"name\":\"NewDrink\",\"type\":\"event\"}]"

// BeveragelistBin is the compiled bytecode used for deploying new contracts.
const BeveragelistBin = `608060405234801561001057600080fd5b50610182806100206000396000f3fe608060405234801561001057600080fd5b5060043610610048576000357c010000000000000000000000000000000000000000000000000000000090048063b89eadb61461004d575b600080fd5b61008d6004803603606081101561006357600080fd5b8101908080359060200190929190803590602001909291908035906020019092919050505061008f565b005b6040805190810160405280838152602001828152506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060008201518160000155602082015181600101559050507f45d101a97a5441b8d2ca576de6954245df176fd85d6ebaad65ae496e89079ce983838360405180848152602001838152602001828152602001935050505060405180910390a150505056fea165627a7a72305820f81f88a88a1a87ea4b5d09e1dc43abcd8b3b9a1056a6a930f6a97fefa9e420bf0029`

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
	Time    [32]byte
	Drink   [32]byte
	Weekday [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewDrink is a free log retrieval operation binding the contract event 0x45d101a97a5441b8d2ca576de6954245df176fd85d6ebaad65ae496e89079ce9.
//
// Solidity: event NewDrink(bytes32 time, bytes32 drink, bytes32 weekday)
func (_Beveragelist *BeveragelistFilterer) FilterNewDrink(opts *bind.FilterOpts) (*BeveragelistNewDrinkIterator, error) {

	logs, sub, err := _Beveragelist.contract.FilterLogs(opts, "NewDrink")
	if err != nil {
		return nil, err
	}
	return &BeveragelistNewDrinkIterator{contract: _Beveragelist.contract, event: "NewDrink", logs: logs, sub: sub}, nil
}

// WatchNewDrink is a free log subscription operation binding the contract event 0x45d101a97a5441b8d2ca576de6954245df176fd85d6ebaad65ae496e89079ce9.
//
// Solidity: event NewDrink(bytes32 time, bytes32 drink, bytes32 weekday)
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
