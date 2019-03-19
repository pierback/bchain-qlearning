// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coffeecoinparent

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

// CoffeecoinparentABI is the input ABI used to generate the binding from.
const CoffeecoinparentABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key_\",\"type\":\"bytes32\"},{\"name\":\"newOrgAddress\",\"type\":\"address\"}],\"name\":\"upgradeCoffeCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key_\",\"type\":\"bytes32\"},{\"name\":\"bvgrlAddress\",\"type\":\"address\"}],\"name\":\"registerCoffeCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"coffeeCoins\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"coffeeCoin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"now\",\"type\":\"uint256\"}],\"name\":\"CoffeCoinCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"coffeeCoin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"now\",\"type\":\"uint256\"}],\"name\":\"CoffeCoinUpgraded\",\"type\":\"event\"}]"

// CoffeecoinparentBin is the compiled bytecode used for deploying new contracts.
const CoffeecoinparentBin = `608060405234801561001057600080fd5b506116eb806100206000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806319f24b9c1461005c5780636eaeb4e9146100ad578063d3b415ab146100fe575b600080fd5b34801561006857600080fd5b506100ab6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061016f565b005b3480156100b957600080fd5b506100fc6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061032b565b005b34801561010a57600080fd5b5061012d6004803603810190808035600019169060200190929190505050610465565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600061017a83610498565b73ffffffffffffffffffffffffffffffffffffffff1663f6333d916040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156101dd57600080fd5b505af11580156101f1573d6000803e3d6000fd5b505050506040513d602081101561020757600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663b262b9ae826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156102b557600080fd5b505af11580156102c9573d6000803e3d6000fd5b5050505081600080856000191660001916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60006103356104dc565b604051809103906000f080158015610351573d6000803e3d6000fd5b5090508173ffffffffffffffffffffffffffffffffffffffff1663b262b9ae826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156103ef57600080fd5b505af1158015610403573d6000803e3d6000fd5b5050505081600080856000191660001916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000836000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6040516111d3806104ed833901905600608060405234801561001057600080fd5b506111b3806100206000396000f300608060405260043610610133576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063043106c0146101385780630c55d9251461016957806317e7dd221461019a57806324d5a450146101e35780633562fd20146102685780633cc1635c146102a35780633eba9ed2146102d45780633f81ac421461031157806344bfa56e146103885780634c77e5ba146104325780635a2bf25a146104a35780636adf83f0146104f45780638267a9ee1461055d5780639007127b1461058e57806393fe4248146105d3578063a209a29c14610604578063a6fa543f146106ae578063a77aa49e14610734578063ba69fcaa1461076f578063bdc963d8146107a0578063c9a52d2c146107e5578063ea20b2a31461085c578063f58660661461089b575b600080fd5b34801561014457600080fd5b506101676004803603810190808035600019169060200190929190505050610912565b005b34801561017557600080fd5b506101986004803603810190808035600019169060200190929190505050610953565b005b3480156101a657600080fd5b506101c9600480360381019080803560001916906020019092919050505061097d565b604051808215151515815260200191505060405180910390f35b3480156101ef57600080fd5b5061024a600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506109af565b60405180826000191660001916815260200191505060405180910390f35b34801561027457600080fd5b506102a1600480360381019080803560001916906020019092919080359060200190929190505050610a24565b005b3480156102af57600080fd5b506102d26004803603810190808035600019169060200190929190505050610a47565b005b3480156102e057600080fd5b5061030f6004803603810190808035600019169060200190929190803515159060200190929190505050610a75565b005b34801561031d57600080fd5b50610386600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035600019169060200190929190505050610aac565b005b34801561039457600080fd5b506103b76004803603810190808035600019169060200190929190505050610b24565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103f75780820151818401526020810190506103dc565b50505050905090810190601f1680156104245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043e57600080fd5b506104616004803603810190808035600019169060200190929190505050610be1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104af57600080fd5b506104f26004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c26565b005b34801561050057600080fd5b5061055b600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610c84565b005b34801561056957600080fd5b5061058c6004803603810190808035600019169060200190929190505050610cf6565b005b34801561059a57600080fd5b506105bd6004803603810190808035600019169060200190929190505050610d18565b6040518082815260200191505060405180910390f35b3480156105df57600080fd5b506106026004803603810190808035600019169060200190929190505050610d3d565b005b34801561061057600080fd5b506106336004803603810190808035600019169060200190929190505050610d5e565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610673578082015181840152602081019050610658565b50505050905090810190601f1680156106a05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156106ba57600080fd5b506106dd6004803603810190808035600019169060200190929190505050610e1b565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610720578082015181840152602081019050610705565b505050509050019250505060405180910390f35b34801561074057600080fd5b5061076d600480360381019080803560001916906020019092919080359060200190929190505050610e92565b005b34801561077b57600080fd5b5061079e6004803603810190808035600019169060200190929190505050610eb6565b005b3480156107ac57600080fd5b506107cf6004803603810190808035600019169060200190929190505050610ee0565b6040518082815260200191505060405180910390f35b3480156107f157600080fd5b5061085a6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610f04565b005b34801561086857600080fd5b5061089960048036038101908080356000191690602001909291908035600019169060200190929190505050610f38565b005b3480156108a757600080fd5b506109106004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610f9e565b005b60026000826000191660001916815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550565b600360008260001916600019168152602001908152602001600020600061097a9190610fd2565b50565b600060066000836000191660001916815260200190815260200160002060009054906101000a900460ff169050919050565b60006004826040518082805190602001908083835b6020831015156109e957805182526020820191506020810190506020830392506109c4565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b806000808460001916600019168152602001908152602001600020819055505050565b60066000826000191660001916815260200190815260200160002060006101000a81549060ff021916905550565b8060066000846000191660001916815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b806004836040518082805190602001908083835b602083101515610ae55780518252602082019150602081019050602083039250610ac0565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902081600019169055505050565b60606003600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bd55780601f10610baa57610100808354040283529160200191610bd5565b820191906000526020600020905b815481529060010190602001808311610bb857829003601f168201915b50505050509050919050565b600060026000836000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8060026000846000191660001916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6004816040518082805190602001908083835b602083101515610cbc5780518252602082019150602081019050602083039250610c97565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000905550565b6007600082600019166000191681526020019081526020016000206000905550565b6000600760008360001916600019168152602001908152602001600020549050919050565b60008082600019166000191681526020019081526020016000206000905550565b60606001600083600019166000191681526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e0f5780601f10610de457610100808354040283529160200191610e0f565b820191906000526020600020905b815481529060010190602001808311610df257829003601f168201915b50505050509050919050565b6060600560008360001916600019168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610e8657602002820191906000526020600020905b81546000191681526020019060010190808311610e6e575b50505050509050919050565b80600760008460001916600019168152602001908152602001600020819055505050565b6001600082600019166000191681526020019081526020016000206000610edd919061101a565b50565b60008060008360001916600019168152602001908152602001600020549050919050565b806003600084600019166000191681526020019081526020016000209080519060200190610f33929190611062565b505050565b60006005600084600019166000191681526020019081526020016000208054905090508160056000856000191660001916815260200190815260200160002060018303815481101515610f8757fe5b906000526020600020018160001916905550505050565b806001600084600019166000191681526020019081526020016000209080519060200190610fcd9291906110e2565b505050565b50805460018160011615610100020316600290046000825580601f10610ff85750611017565b601f0160209004906000526020600020908101906110169190611162565b5b50565b50805460018160011615610100020316600290046000825580601f10611040575061105f565b601f01602090049060005260206000209081019061105e9190611162565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110a357805160ff19168380011785556110d1565b828001600101855582156110d1579182015b828111156110d05782518255916020019190600101906110b5565b5b5090506110de9190611162565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061112357805160ff1916838001178555611151565b82800160010185558215611151579182015b82811115611150578251825591602001919060010190611135565b5b50905061115e9190611162565b5090565b61118491905b80821115611180576000816000905550600101611168565b5090565b905600a165627a7a72305820eebad841af6daef8bf6e8a27efc89007175cec2b04e77996a823e0dd7d220a630029a165627a7a723058207258ec037f5f7269dedc67b4d9b4b307b4b6d4707ce161ca29def0932f4d48f50029`

// DeployCoffeecoinparent deploys a new Ethereum contract, binding an instance of Coffeecoinparent to it.
func DeployCoffeecoinparent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Coffeecoinparent, error) {
	parsed, err := abi.JSON(strings.NewReader(CoffeecoinparentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoffeecoinparentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Coffeecoinparent{CoffeecoinparentCaller: CoffeecoinparentCaller{contract: contract}, CoffeecoinparentTransactor: CoffeecoinparentTransactor{contract: contract}, CoffeecoinparentFilterer: CoffeecoinparentFilterer{contract: contract}}, nil
}

// Coffeecoinparent is an auto generated Go binding around an Ethereum contract.
type Coffeecoinparent struct {
	CoffeecoinparentCaller     // Read-only binding to the contract
	CoffeecoinparentTransactor // Write-only binding to the contract
	CoffeecoinparentFilterer   // Log filterer for contract events
}

// CoffeecoinparentCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoffeecoinparentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeecoinparentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoffeecoinparentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeecoinparentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoffeecoinparentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoffeecoinparentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoffeecoinparentSession struct {
	Contract     *Coffeecoinparent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoffeecoinparentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoffeecoinparentCallerSession struct {
	Contract *CoffeecoinparentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CoffeecoinparentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoffeecoinparentTransactorSession struct {
	Contract     *CoffeecoinparentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CoffeecoinparentRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoffeecoinparentRaw struct {
	Contract *Coffeecoinparent // Generic contract binding to access the raw methods on
}

// CoffeecoinparentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoffeecoinparentCallerRaw struct {
	Contract *CoffeecoinparentCaller // Generic read-only contract binding to access the raw methods on
}

// CoffeecoinparentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoffeecoinparentTransactorRaw struct {
	Contract *CoffeecoinparentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoffeecoinparent creates a new instance of Coffeecoinparent, bound to a specific deployed contract.
func NewCoffeecoinparent(address common.Address, backend bind.ContractBackend) (*Coffeecoinparent, error) {
	contract, err := bindCoffeecoinparent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coffeecoinparent{CoffeecoinparentCaller: CoffeecoinparentCaller{contract: contract}, CoffeecoinparentTransactor: CoffeecoinparentTransactor{contract: contract}, CoffeecoinparentFilterer: CoffeecoinparentFilterer{contract: contract}}, nil
}

// NewCoffeecoinparentCaller creates a new read-only instance of Coffeecoinparent, bound to a specific deployed contract.
func NewCoffeecoinparentCaller(address common.Address, caller bind.ContractCaller) (*CoffeecoinparentCaller, error) {
	contract, err := bindCoffeecoinparent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoffeecoinparentCaller{contract: contract}, nil
}

// NewCoffeecoinparentTransactor creates a new write-only instance of Coffeecoinparent, bound to a specific deployed contract.
func NewCoffeecoinparentTransactor(address common.Address, transactor bind.ContractTransactor) (*CoffeecoinparentTransactor, error) {
	contract, err := bindCoffeecoinparent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoffeecoinparentTransactor{contract: contract}, nil
}

// NewCoffeecoinparentFilterer creates a new log filterer instance of Coffeecoinparent, bound to a specific deployed contract.
func NewCoffeecoinparentFilterer(address common.Address, filterer bind.ContractFilterer) (*CoffeecoinparentFilterer, error) {
	contract, err := bindCoffeecoinparent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoffeecoinparentFilterer{contract: contract}, nil
}

// bindCoffeecoinparent binds a generic wrapper to an already deployed contract.
func bindCoffeecoinparent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoffeecoinparentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coffeecoinparent *CoffeecoinparentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coffeecoinparent.Contract.CoffeecoinparentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coffeecoinparent *CoffeecoinparentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.CoffeecoinparentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coffeecoinparent *CoffeecoinparentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.CoffeecoinparentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coffeecoinparent *CoffeecoinparentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coffeecoinparent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coffeecoinparent *CoffeecoinparentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coffeecoinparent *CoffeecoinparentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.contract.Transact(opts, method, params...)
}

// CoffeeCoins is a free data retrieval call binding the contract method 0xd3b415ab.
//
// Solidity: function coffeeCoins( bytes32) constant returns(address)
func (_Coffeecoinparent *CoffeecoinparentCaller) CoffeeCoins(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Coffeecoinparent.contract.Call(opts, out, "coffeeCoins", arg0)
	return *ret0, err
}

// CoffeeCoins is a free data retrieval call binding the contract method 0xd3b415ab.
//
// Solidity: function coffeeCoins( bytes32) constant returns(address)
func (_Coffeecoinparent *CoffeecoinparentSession) CoffeeCoins(arg0 [32]byte) (common.Address, error) {
	return _Coffeecoinparent.Contract.CoffeeCoins(&_Coffeecoinparent.CallOpts, arg0)
}

// CoffeeCoins is a free data retrieval call binding the contract method 0xd3b415ab.
//
// Solidity: function coffeeCoins( bytes32) constant returns(address)
func (_Coffeecoinparent *CoffeecoinparentCallerSession) CoffeeCoins(arg0 [32]byte) (common.Address, error) {
	return _Coffeecoinparent.Contract.CoffeeCoins(&_Coffeecoinparent.CallOpts, arg0)
}

// RegisterCoffeCoin is a paid mutator transaction binding the contract method 0x6eaeb4e9.
//
// Solidity: function registerCoffeCoin(key_ bytes32, bvgrlAddress address) returns()
func (_Coffeecoinparent *CoffeecoinparentTransactor) RegisterCoffeCoin(opts *bind.TransactOpts, key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.contract.Transact(opts, "registerCoffeCoin", key_, bvgrlAddress)
}

// RegisterCoffeCoin is a paid mutator transaction binding the contract method 0x6eaeb4e9.
//
// Solidity: function registerCoffeCoin(key_ bytes32, bvgrlAddress address) returns()
func (_Coffeecoinparent *CoffeecoinparentSession) RegisterCoffeCoin(key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.RegisterCoffeCoin(&_Coffeecoinparent.TransactOpts, key_, bvgrlAddress)
}

// RegisterCoffeCoin is a paid mutator transaction binding the contract method 0x6eaeb4e9.
//
// Solidity: function registerCoffeCoin(key_ bytes32, bvgrlAddress address) returns()
func (_Coffeecoinparent *CoffeecoinparentTransactorSession) RegisterCoffeCoin(key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.RegisterCoffeCoin(&_Coffeecoinparent.TransactOpts, key_, bvgrlAddress)
}

// UpgradeCoffeCoin is a paid mutator transaction binding the contract method 0x19f24b9c.
//
// Solidity: function upgradeCoffeCoin(key_ bytes32, newOrgAddress address) returns()
func (_Coffeecoinparent *CoffeecoinparentTransactor) UpgradeCoffeCoin(opts *bind.TransactOpts, key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.contract.Transact(opts, "upgradeCoffeCoin", key_, newOrgAddress)
}

// UpgradeCoffeCoin is a paid mutator transaction binding the contract method 0x19f24b9c.
//
// Solidity: function upgradeCoffeCoin(key_ bytes32, newOrgAddress address) returns()
func (_Coffeecoinparent *CoffeecoinparentSession) UpgradeCoffeCoin(key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.UpgradeCoffeCoin(&_Coffeecoinparent.TransactOpts, key_, newOrgAddress)
}

// UpgradeCoffeCoin is a paid mutator transaction binding the contract method 0x19f24b9c.
//
// Solidity: function upgradeCoffeCoin(key_ bytes32, newOrgAddress address) returns()
func (_Coffeecoinparent *CoffeecoinparentTransactorSession) UpgradeCoffeCoin(key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.UpgradeCoffeCoin(&_Coffeecoinparent.TransactOpts, key_, newOrgAddress)
}

// CoffeecoinparentCoffeCoinCreatedIterator is returned from FilterCoffeCoinCreated and is used to iterate over the raw logs and unpacked data for CoffeCoinCreated events raised by the Coffeecoinparent contract.
type CoffeecoinparentCoffeCoinCreatedIterator struct {
	Event *CoffeecoinparentCoffeCoinCreated // Event containing the contract specifics and raw log

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
func (it *CoffeecoinparentCoffeCoinCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoffeecoinparentCoffeCoinCreated)
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
		it.Event = new(CoffeecoinparentCoffeCoinCreated)
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
func (it *CoffeecoinparentCoffeCoinCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoffeecoinparentCoffeCoinCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoffeecoinparentCoffeCoinCreated represents a CoffeCoinCreated event raised by the Coffeecoinparent contract.
type CoffeecoinparentCoffeCoinCreated struct {
	CoffeeCoin common.Address
	Now        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCoffeCoinCreated is a free log retrieval operation binding the contract event 0xaf77b243b6b49fa214ccfd618b8ebec2e465aff70aeb80062e35ae83e3cd8d92.
//
// Solidity: e CoffeCoinCreated(coffeeCoin address, now uint256)
func (_Coffeecoinparent *CoffeecoinparentFilterer) FilterCoffeCoinCreated(opts *bind.FilterOpts) (*CoffeecoinparentCoffeCoinCreatedIterator, error) {

	logs, sub, err := _Coffeecoinparent.contract.FilterLogs(opts, "CoffeCoinCreated")
	if err != nil {
		return nil, err
	}
	return &CoffeecoinparentCoffeCoinCreatedIterator{contract: _Coffeecoinparent.contract, event: "CoffeCoinCreated", logs: logs, sub: sub}, nil
}

// WatchCoffeCoinCreated is a free log subscription operation binding the contract event 0xaf77b243b6b49fa214ccfd618b8ebec2e465aff70aeb80062e35ae83e3cd8d92.
//
// Solidity: e CoffeCoinCreated(coffeeCoin address, now uint256)
func (_Coffeecoinparent *CoffeecoinparentFilterer) WatchCoffeCoinCreated(opts *bind.WatchOpts, sink chan<- *CoffeecoinparentCoffeCoinCreated) (event.Subscription, error) {

	logs, sub, err := _Coffeecoinparent.contract.WatchLogs(opts, "CoffeCoinCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoffeecoinparentCoffeCoinCreated)
				if err := _Coffeecoinparent.contract.UnpackLog(event, "CoffeCoinCreated", log); err != nil {
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

// CoffeecoinparentCoffeCoinUpgradedIterator is returned from FilterCoffeCoinUpgraded and is used to iterate over the raw logs and unpacked data for CoffeCoinUpgraded events raised by the Coffeecoinparent contract.
type CoffeecoinparentCoffeCoinUpgradedIterator struct {
	Event *CoffeecoinparentCoffeCoinUpgraded // Event containing the contract specifics and raw log

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
func (it *CoffeecoinparentCoffeCoinUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoffeecoinparentCoffeCoinUpgraded)
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
		it.Event = new(CoffeecoinparentCoffeCoinUpgraded)
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
func (it *CoffeecoinparentCoffeCoinUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoffeecoinparentCoffeCoinUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoffeecoinparentCoffeCoinUpgraded represents a CoffeCoinUpgraded event raised by the Coffeecoinparent contract.
type CoffeecoinparentCoffeCoinUpgraded struct {
	CoffeeCoin common.Address
	Now        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCoffeCoinUpgraded is a free log retrieval operation binding the contract event 0xfac28ffd2b73b646bd4d039a13cee7a1fb1e30bc6ea9a043b802610281de2ef8.
//
// Solidity: e CoffeCoinUpgraded(coffeeCoin address, now uint256)
func (_Coffeecoinparent *CoffeecoinparentFilterer) FilterCoffeCoinUpgraded(opts *bind.FilterOpts) (*CoffeecoinparentCoffeCoinUpgradedIterator, error) {

	logs, sub, err := _Coffeecoinparent.contract.FilterLogs(opts, "CoffeCoinUpgraded")
	if err != nil {
		return nil, err
	}
	return &CoffeecoinparentCoffeCoinUpgradedIterator{contract: _Coffeecoinparent.contract, event: "CoffeCoinUpgraded", logs: logs, sub: sub}, nil
}

// WatchCoffeCoinUpgraded is a free log subscription operation binding the contract event 0xfac28ffd2b73b646bd4d039a13cee7a1fb1e30bc6ea9a043b802610281de2ef8.
//
// Solidity: e CoffeCoinUpgraded(coffeeCoin address, now uint256)
func (_Coffeecoinparent *CoffeecoinparentFilterer) WatchCoffeCoinUpgraded(opts *bind.WatchOpts, sink chan<- *CoffeecoinparentCoffeCoinUpgraded) (event.Subscription, error) {

	logs, sub, err := _Coffeecoinparent.contract.WatchLogs(opts, "CoffeCoinUpgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoffeecoinparentCoffeCoinUpgraded)
				if err := _Coffeecoinparent.contract.UnpackLog(event, "CoffeCoinUpgraded", log); err != nil {
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
