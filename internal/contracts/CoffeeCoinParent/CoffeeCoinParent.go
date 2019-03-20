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

// CoffeecoinparentABI is the input ABI used to generate the binding from.
const CoffeecoinparentABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key_\",\"type\":\"bytes32\"},{\"name\":\"newOrgAddress\",\"type\":\"address\"}],\"name\":\"upgradeCoffeCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key_\",\"type\":\"bytes32\"},{\"name\":\"bvgrlAddress\",\"type\":\"address\"}],\"name\":\"registerCoffeCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"coffeeCoins\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"coffeeCoin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"now\",\"type\":\"uint256\"}],\"name\":\"CoffeCoinCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"coffeeCoin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"now\",\"type\":\"uint256\"}],\"name\":\"CoffeCoinUpgraded\",\"type\":\"event\"}]"

// CoffeecoinparentBin is the compiled bytecode used for deploying new contracts.
const CoffeecoinparentBin = `608060405234801561001057600080fd5b5061179f806100206000396000f3fe608060405234801561001057600080fd5b506004361061005e576000357c01000000000000000000000000000000000000000000000000000000009004806319f24b9c146100635780636eaeb4e9146100b1578063d3b415ab146100ff575b600080fd5b6100af6004803603604081101561007957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061016d565b005b6100fd600480360360408110156100c757600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061031f565b005b61012b6004803603602081101561011557600080fd5b8101908080359060200190929190505050610455565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600061017883610488565b73ffffffffffffffffffffffffffffffffffffffff1663f6333d916040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156101d957600080fd5b505afa1580156101ed573d6000803e3d6000fd5b505050506040513d602081101561020357600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663b262b9ae826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156102b157600080fd5b505af11580156102c5573d6000803e3d6000fd5b505050508160008085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600060405161032d906104c4565b604051809103906000f080158015610349573d6000803e3d6000fd5b5090508173ffffffffffffffffffffffffffffffffffffffff1663b262b9ae826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156103e757600080fd5b505af11580156103fb573d6000803e3d6000fd5b505050508160008085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6112a2806104d28339019056fe608060405234801561001057600080fd5b50611282806100206000396000f3fe608060405234801561001057600080fd5b506004361061016a576000357c0100000000000000000000000000000000000000000000000000000000900480636adf83f0116100e0578063a77aa49e11610099578063a77aa49e1461082b578063ba69fcaa14610863578063bdc963d814610891578063c9a52d2c146108d3578063ea20b2a314610998578063f5866066146109d05761016a565b80636adf83f0146105a85780638267a9ee146106635780639007127b1461069157806393fe4248146106d3578063a209a29c14610701578063a6fa543f146107a85761016a565b80633cc1635c116101325780633cc1635c146103185780633eba9ed2146103465780633f81ac421461038057806344bfa56e146104455780634c77e5ba146104ec5780635a2bf25a1461055a5761016a565b8063043106c01461016f5780630c55d9251461019d57806317e7dd22146101cb57806324d5a450146102115780633562fd20146102e0575b600080fd5b61019b6004803603602081101561018557600080fd5b8101908080359060200190929190505050610a95565b005b6101c9600480360360208110156101b357600080fd5b8101908080359060200190929190505050610ace565b005b6101f7600480360360208110156101e157600080fd5b8101908080359060200190929190505050610af0565b604051808215151515815260200191505060405180910390f35b6102ca6004803603602081101561022757600080fd5b810190808035906020019064010000000081111561024457600080fd5b82018360208201111561025657600080fd5b8035906020019184600183028401116401000000008311171561027857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610b1a565b6040518082815260200191505060405180910390f35b610316600480360360408110156102f657600080fd5b810190808035906020019092919080359060200190929190505050610b8f565b005b6103446004803603602081101561032e57600080fd5b8101908080359060200190929190505050610baa565b005b61037e6004803603604081101561035c57600080fd5b8101908080359060200190929190803515159060200190929190505050610bd0565b005b6104436004803603604081101561039657600080fd5b81019080803590602001906401000000008111156103b357600080fd5b8201836020820111156103c557600080fd5b803590602001918460018302840111640100000000831117156103e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610bff565b005b6104716004803603602081101561045b57600080fd5b8101908080359060200190929190505050610c73565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104b1578082015181840152602081019050610496565b50505050905090810190601f1680156104de5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6105186004803603602081101561050257600080fd5b8101908080359060200190929190505050610d28565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105a66004803603604081101561057057600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d65565b005b610661600480360360208110156105be57600080fd5b81019080803590602001906401000000008111156105db57600080fd5b8201836020820111156105ed57600080fd5b8035906020019184600183028401116401000000008311171561060f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610dbb565b005b61068f6004803603602081101561067957600080fd5b8101908080359060200190929190505050610e2d565b005b6106bd600480360360208110156106a757600080fd5b8101908080359060200190929190505050610e47565b6040518082815260200191505060405180910390f35b6106ff600480360360208110156106e957600080fd5b8101908080359060200190929190505050610e64565b005b61072d6004803603602081101561071757600080fd5b8101908080359060200190929190505050610e7d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561076d578082015181840152602081019050610752565b50505050905090810190601f16801561079a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6107d4600480360360208110156107be57600080fd5b8101908080359060200190929190505050610f32565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156108175780820151818401526020810190506107fc565b505050509050019250505060405180910390f35b6108616004803603604081101561084157600080fd5b810190808035906020019092919080359060200190929190505050610f9d565b005b61088f6004803603602081101561087957600080fd5b8101908080359060200190929190505050610fb9565b005b6108bd600480360360208110156108a757600080fd5b8101908080359060200190929190505050610fdb565b6040518082815260200191505060405180910390f35b610996600480360360408110156108e957600080fd5b81019080803590602001909291908035906020019064010000000081111561091057600080fd5b82018360208201111561092257600080fd5b8035906020019184600183028401116401000000008311171561094457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610ff7565b005b6109ce600480360360408110156109ae57600080fd5b810190808035906020019092919080359060200190929190505050611023565b005b610a93600480360360408110156109e657600080fd5b810190808035906020019092919080359060200190640100000000811115610a0d57600080fd5b820183602082011115610a1f57600080fd5b80359060200191846001830284011164010000000083111715610a4157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611075565b005b6002600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550565b600360008281526020019081526020016000206000610aed91906110a1565b50565b60006006600083815260200190815260200160002060009054906101000a900460ff169050919050565b60006004826040518082805190602001908083835b602083101515610b545780518252602082019150602081019050602083039250610b2f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b80600080848152602001908152602001600020819055505050565b6006600082815260200190815260200160002060006101000a81549060ff021916905550565b806006600084815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b806004836040518082805190602001908083835b602083101515610c385780518252602082019150602081019050602083039250610c13565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055505050565b6060600360008381526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d1c5780601f10610cf157610100808354040283529160200191610d1c565b820191906000526020600020905b815481529060010190602001808311610cff57829003601f168201915b50505050509050919050565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6004816040518082805190602001908083835b602083101515610df35780518252602082019150602081019050602083039250610dce565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000905550565b600760008281526020019081526020016000206000905550565b600060076000838152602001908152602001600020549050919050565b6000808281526020019081526020016000206000905550565b6060600160008381526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f265780601f10610efb57610100808354040283529160200191610f26565b820191906000526020600020905b815481529060010190602001808311610f0957829003601f168201915b50505050509050919050565b606060056000838152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610f9157602002820191906000526020600020905b815481526020019060010190808311610f7d575b50505050509050919050565b8060076000848152602001908152602001600020819055505050565b600160008281526020019081526020016000206000610fd891906110e9565b50565b6000806000838152602001908152602001600020549050919050565b8060036000848152602001908152602001600020908051906020019061101e929190611131565b505050565b60006005600084815260200190815260200160002080549050905081600560008581526020019081526020016000206001830381548110151561106257fe5b9060005260206000200181905550505050565b8060016000848152602001908152602001600020908051906020019061109c9291906111b1565b505050565b50805460018160011615610100020316600290046000825580601f106110c757506110e6565b601f0160209004906000526020600020908101906110e59190611231565b5b50565b50805460018160011615610100020316600290046000825580601f1061110f575061112e565b601f01602090049060005260206000209081019061112d9190611231565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061117257805160ff19168380011785556111a0565b828001600101855582156111a0579182015b8281111561119f578251825591602001919060010190611184565b5b5090506111ad9190611231565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106111f257805160ff1916838001178555611220565b82800160010185558215611220579182015b8281111561121f578251825591602001919060010190611204565b5b50905061122d9190611231565b5090565b61125391905b8082111561124f576000816000905550600101611237565b5090565b9056fea165627a7a72305820461a753f545f098d766aa1dc1987ff8578ee505f3bf64465b4c784ac0df0a9290029a165627a7a72305820b0167740b8d0eca04b40de5dcfde985f1b7aa4babd81889c1fa721fb71330a840029`

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
// Solidity: function coffeeCoins(bytes32 ) constant returns(address)
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
// Solidity: function coffeeCoins(bytes32 ) constant returns(address)
func (_Coffeecoinparent *CoffeecoinparentSession) CoffeeCoins(arg0 [32]byte) (common.Address, error) {
	return _Coffeecoinparent.Contract.CoffeeCoins(&_Coffeecoinparent.CallOpts, arg0)
}

// CoffeeCoins is a free data retrieval call binding the contract method 0xd3b415ab.
//
// Solidity: function coffeeCoins(bytes32 ) constant returns(address)
func (_Coffeecoinparent *CoffeecoinparentCallerSession) CoffeeCoins(arg0 [32]byte) (common.Address, error) {
	return _Coffeecoinparent.Contract.CoffeeCoins(&_Coffeecoinparent.CallOpts, arg0)
}

// RegisterCoffeCoin is a paid mutator transaction binding the contract method 0x6eaeb4e9.
//
// Solidity: function registerCoffeCoin(bytes32 key_, address bvgrlAddress) returns()
func (_Coffeecoinparent *CoffeecoinparentTransactor) RegisterCoffeCoin(opts *bind.TransactOpts, key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.contract.Transact(opts, "registerCoffeCoin", key_, bvgrlAddress)
}

// RegisterCoffeCoin is a paid mutator transaction binding the contract method 0x6eaeb4e9.
//
// Solidity: function registerCoffeCoin(bytes32 key_, address bvgrlAddress) returns()
func (_Coffeecoinparent *CoffeecoinparentSession) RegisterCoffeCoin(key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.RegisterCoffeCoin(&_Coffeecoinparent.TransactOpts, key_, bvgrlAddress)
}

// RegisterCoffeCoin is a paid mutator transaction binding the contract method 0x6eaeb4e9.
//
// Solidity: function registerCoffeCoin(bytes32 key_, address bvgrlAddress) returns()
func (_Coffeecoinparent *CoffeecoinparentTransactorSession) RegisterCoffeCoin(key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.RegisterCoffeCoin(&_Coffeecoinparent.TransactOpts, key_, bvgrlAddress)
}

// UpgradeCoffeCoin is a paid mutator transaction binding the contract method 0x19f24b9c.
//
// Solidity: function upgradeCoffeCoin(bytes32 key_, address newOrgAddress) returns()
func (_Coffeecoinparent *CoffeecoinparentTransactor) UpgradeCoffeCoin(opts *bind.TransactOpts, key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.contract.Transact(opts, "upgradeCoffeCoin", key_, newOrgAddress)
}

// UpgradeCoffeCoin is a paid mutator transaction binding the contract method 0x19f24b9c.
//
// Solidity: function upgradeCoffeCoin(bytes32 key_, address newOrgAddress) returns()
func (_Coffeecoinparent *CoffeecoinparentSession) UpgradeCoffeCoin(key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Coffeecoinparent.Contract.UpgradeCoffeCoin(&_Coffeecoinparent.TransactOpts, key_, newOrgAddress)
}

// UpgradeCoffeCoin is a paid mutator transaction binding the contract method 0x19f24b9c.
//
// Solidity: function upgradeCoffeCoin(bytes32 key_, address newOrgAddress) returns()
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
// Solidity: event CoffeCoinCreated(address coffeeCoin, uint256 now)
func (_Coffeecoinparent *CoffeecoinparentFilterer) FilterCoffeCoinCreated(opts *bind.FilterOpts) (*CoffeecoinparentCoffeCoinCreatedIterator, error) {

	logs, sub, err := _Coffeecoinparent.contract.FilterLogs(opts, "CoffeCoinCreated")
	if err != nil {
		return nil, err
	}
	return &CoffeecoinparentCoffeCoinCreatedIterator{contract: _Coffeecoinparent.contract, event: "CoffeCoinCreated", logs: logs, sub: sub}, nil
}

// WatchCoffeCoinCreated is a free log subscription operation binding the contract event 0xaf77b243b6b49fa214ccfd618b8ebec2e465aff70aeb80062e35ae83e3cd8d92.
//
// Solidity: event CoffeCoinCreated(address coffeeCoin, uint256 now)
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
// Solidity: event CoffeCoinUpgraded(address coffeeCoin, uint256 now)
func (_Coffeecoinparent *CoffeecoinparentFilterer) FilterCoffeCoinUpgraded(opts *bind.FilterOpts) (*CoffeecoinparentCoffeCoinUpgradedIterator, error) {

	logs, sub, err := _Coffeecoinparent.contract.FilterLogs(opts, "CoffeCoinUpgraded")
	if err != nil {
		return nil, err
	}
	return &CoffeecoinparentCoffeCoinUpgradedIterator{contract: _Coffeecoinparent.contract, event: "CoffeCoinUpgraded", logs: logs, sub: sub}, nil
}

// WatchCoffeCoinUpgraded is a free log subscription operation binding the contract event 0xfac28ffd2b73b646bd4d039a13cee7a1fb1e30bc6ea9a043b802610281de2ef8.
//
// Solidity: event CoffeCoinUpgraded(address coffeeCoin, uint256 now)
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
