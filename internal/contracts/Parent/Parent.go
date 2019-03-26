// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parent

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

// ParentABI is the input ABI used to generate the binding from.
const ParentABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key_\",\"type\":\"bytes32\"},{\"name\":\"bvgrlAddress\",\"type\":\"address\"}],\"name\":\"registerBeverageList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"beverageLists\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key_\",\"type\":\"bytes32\"},{\"name\":\"newOrgAddress\",\"type\":\"address\"}],\"name\":\"upgradeBeverageList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"beverageList\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"now\",\"type\":\"uint256\"}],\"name\":\"BeverageListCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"beverageList\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"now\",\"type\":\"uint256\"}],\"name\":\"BeverageListUpgraded\",\"type\":\"event\"}]"

// ParentBin is the compiled bytecode used for deploying new contracts.
const ParentBin = `608060405234801561001057600080fd5b50611952806100206000396000f3fe608060405234801561001057600080fd5b506004361061005e576000357c0100000000000000000000000000000000000000000000000000000000900480630b2cc919146100635780633a784b32146100b157806344f338a01461011f575b600080fd5b6100af6004803603604081101561007957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061016d565b005b6100dd600480360360208110156100c757600080fd5b81019080803590602001909291905050506102a3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61016b6004803603604081101561013557600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102d6565b005b600060405161017b906104c4565b604051809103906000f080158015610197573d6000803e3d6000fd5b5090508173ffffffffffffffffffffffffffffffffffffffff1663047e2b27826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561023557600080fd5b505af1158015610249573d6000803e3d6000fd5b505050508160008085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006102e183610488565b73ffffffffffffffffffffffffffffffffffffffff166386cc31ac6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561034257600080fd5b505afa158015610356573d6000803e3d6000fd5b505050506040513d602081101561036c57600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663047e2b27826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561041a57600080fd5b505af115801561042e573d6000803e3d6000fd5b505050508160008085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600080600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b611455806104d28339019056fe608060405234801561001057600080fd5b50611435806100206000396000f3fe608060405234801561001057600080fd5b5060043610610190576000357c0100000000000000000000000000000000000000000000000000000000900480637687d797116100fb578063a77aa49e116100b4578063c9a52d2c1161008e578063c9a52d2c14610962578063dc317db714610a27578063ea20b2a314610a89578063f586606614610ac157610190565b8063a77aa49e146108ba578063ba69fcaa146108f2578063bdc963d81461092057610190565b80637687d797146106895780638267a9ee146106f25780639007127b1461072057806393fe424814610762578063a209a29c14610790578063a6fa543f1461083757610190565b80633eba9ed21161014d5780633eba9ed21461036c5780633f81ac42146103a657806344bfa56e1461046b5780634c77e5ba146105125780635a2bf25a146105805780636adf83f0146105ce57610190565b8063043106c0146101955780630c55d925146101c357806317e7dd22146101f157806324d5a450146102375780633562fd20146103065780633cc1635c1461033e575b600080fd5b6101c1600480360360208110156101ab57600080fd5b8101908080359060200190929190505050610b86565b005b6101ef600480360360208110156101d957600080fd5b8101908080359060200190929190505050610bbf565b005b61021d6004803603602081101561020757600080fd5b8101908080359060200190929190505050610be1565b604051808215151515815260200191505060405180910390f35b6102f06004803603602081101561024d57600080fd5b810190808035906020019064010000000081111561026a57600080fd5b82018360208201111561027c57600080fd5b8035906020019184600183028401116401000000008311171561029e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610c0b565b6040518082815260200191505060405180910390f35b61033c6004803603604081101561031c57600080fd5b810190808035906020019092919080359060200190929190505050610c80565b005b61036a6004803603602081101561035457600080fd5b8101908080359060200190929190505050610c9c565b005b6103a46004803603604081101561038257600080fd5b8101908080359060200190929190803515159060200190929190505050610cc2565b005b610469600480360360408110156103bc57600080fd5b81019080803590602001906401000000008111156103d957600080fd5b8201836020820111156103eb57600080fd5b8035906020019184600183028401116401000000008311171561040d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610cf1565b005b6104976004803603602081101561048157600080fd5b8101908080359060200190929190505050610d65565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104d75780820151818401526020810190506104bc565b50505050905090810190601f1680156105045780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61053e6004803603602081101561052857600080fd5b8101908080359060200190929190505050610e1a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105cc6004803603604081101561059657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e57565b005b610687600480360360208110156105e457600080fd5b810190808035906020019064010000000081111561060157600080fd5b82018360208201111561061357600080fd5b8035906020019184600183028401116401000000008311171561063557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610ead565b005b6106d56004803603604081101561069f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f1f565b604051808381526020018281526020019250505060405180910390f35b61071e6004803603602081101561070857600080fd5b8101908080359060200190929190505050610f84565b005b61074c6004803603602081101561073657600080fd5b8101908080359060200190929190505050610f9e565b6040518082815260200191505060405180910390f35b61078e6004803603602081101561077857600080fd5b8101908080359060200190929190505050610fbb565b005b6107bc600480360360208110156107a657600080fd5b8101908080359060200190929190505050610fd5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107fc5780820151818401526020810190506107e1565b50505050905090810190601f1680156108295780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6108636004803603602081101561084d57600080fd5b810190808035906020019092919050505061108a565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156108a657808201518184015260208101905061088b565b505050509050019250505060405180910390f35b6108f0600480360360408110156108d057600080fd5b8101908080359060200190929190803590602001909291905050506110f5565b005b61091e6004803603602081101561090857600080fd5b8101908080359060200190929190505050611111565b005b61094c6004803603602081101561093657600080fd5b8101908080359060200190929190505050611133565b6040518082815260200191505060405180910390f35b610a256004803603604081101561097857600080fd5b81019080803590602001909291908035906020019064010000000081111561099f57600080fd5b8201836020820111156109b157600080fd5b803590602001918460018302840111640100000000831117156109d357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611150565b005b610a8760048036036080811015610a3d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919050505061117c565b005b610abf60048036036040811015610a9f57600080fd5b8101908080359060200190929190803590602001909291905050506111d6565b005b610b8460048036036040811015610ad757600080fd5b810190808035906020019092919080359060200190640100000000811115610afe57600080fd5b820183602082011115610b1057600080fd5b80359060200191846001830284011164010000000083111715610b3257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611228565b005b6003600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550565b600460008281526020019081526020016000206000610bde9190611254565b50565b60006007600083815260200190815260200160002060009054906101000a900460ff169050919050565b60006005826040518082805190602001908083835b602083101515610c455780518252602082019150602081019050602083039250610c20565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b8060016000848152602001908152602001600020819055505050565b6007600082815260200190815260200160002060006101000a81549060ff021916905550565b806007600084815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b806005836040518082805190602001908083835b602083101515610d2a5780518252602082019150602081019050602083039250610d05565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055505050565b6060600460008381526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e0e5780601f10610de357610100808354040283529160200191610e0e565b820191906000526020600020905b815481529060010190602001808311610df157829003601f168201915b50505050509050919050565b60006003600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6005816040518082805190602001908083835b602083101515610ee55780518252602082019150602081019050602083039250610ec0565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000905550565b60008060008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020549050808192509250509250929050565b600860008281526020019081526020016000206000905550565b600060086000838152602001908152602001600020549050919050565b600160008281526020019081526020016000206000905550565b6060600260008381526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561107e5780601f106110535761010080835404028352916020019161107e565b820191906000526020600020905b81548152906001019060200180831161106157829003601f168201915b50505050509050919050565b6060600660008381526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156110e957602002820191906000526020600020905b8154815260200190600101908083116110d5575b50505050509050919050565b8060086000848152602001908152602001600020819055505050565b600260008281526020019081526020016000206000611130919061129c565b50565b600060016000838152602001908152602001600020549050919050565b806004600084815260200190815260200160002090805190602001906111779291906112e4565b505050565b816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000208190555050505050565b60006006600084815260200190815260200160002080549050905081600660008581526020019081526020016000206001830381548110151561121557fe5b9060005260206000200181905550505050565b8060026000848152602001908152602001600020908051906020019061124f929190611364565b505050565b50805460018160011615610100020316600290046000825580601f1061127a5750611299565b601f01602090049060005260206000209081019061129891906113e4565b5b50565b50805460018160011615610100020316600290046000825580601f106112c257506112e1565b601f0160209004906000526020600020908101906112e091906113e4565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061132557805160ff1916838001178555611353565b82800160010185558215611353579182015b82811115611352578251825591602001919060010190611337565b5b50905061136091906113e4565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113a557805160ff19168380011785556113d3565b828001600101855582156113d3579182015b828111156113d25782518255916020019190600101906113b7565b5b5090506113e091906113e4565b5090565b61140691905b808211156114025760008160009055506001016113ea565b5090565b9056fea165627a7a72305820c77afbfe0ba8aac3f2ec2becbaca80ef15f13ff0dc478973b112803f7d77df720029a165627a7a72305820aa9db915c4804018b3237c9724cf45835351f86ec8151d70f16d014c9cf3e3070029`

// DeployParent deploys a new Ethereum contract, binding an instance of Parent to it.
func DeployParent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Parent, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Parent{ParentCaller: ParentCaller{contract: contract}, ParentTransactor: ParentTransactor{contract: contract}, ParentFilterer: ParentFilterer{contract: contract}}, nil
}

// Parent is an auto generated Go binding around an Ethereum contract.
type Parent struct {
	ParentCaller     // Read-only binding to the contract
	ParentTransactor // Write-only binding to the contract
	ParentFilterer   // Log filterer for contract events
}

// ParentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParentSession struct {
	Contract     *Parent           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParentCallerSession struct {
	Contract *ParentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ParentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParentTransactorSession struct {
	Contract     *ParentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParentRaw struct {
	Contract *Parent // Generic contract binding to access the raw methods on
}

// ParentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParentCallerRaw struct {
	Contract *ParentCaller // Generic read-only contract binding to access the raw methods on
}

// ParentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParentTransactorRaw struct {
	Contract *ParentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParent creates a new instance of Parent, bound to a specific deployed contract.
func NewParent(address common.Address, backend bind.ContractBackend) (*Parent, error) {
	contract, err := bindParent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parent{ParentCaller: ParentCaller{contract: contract}, ParentTransactor: ParentTransactor{contract: contract}, ParentFilterer: ParentFilterer{contract: contract}}, nil
}

// NewParentCaller creates a new read-only instance of Parent, bound to a specific deployed contract.
func NewParentCaller(address common.Address, caller bind.ContractCaller) (*ParentCaller, error) {
	contract, err := bindParent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParentCaller{contract: contract}, nil
}

// NewParentTransactor creates a new write-only instance of Parent, bound to a specific deployed contract.
func NewParentTransactor(address common.Address, transactor bind.ContractTransactor) (*ParentTransactor, error) {
	contract, err := bindParent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParentTransactor{contract: contract}, nil
}

// NewParentFilterer creates a new log filterer instance of Parent, bound to a specific deployed contract.
func NewParentFilterer(address common.Address, filterer bind.ContractFilterer) (*ParentFilterer, error) {
	contract, err := bindParent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParentFilterer{contract: contract}, nil
}

// bindParent binds a generic wrapper to an already deployed contract.
func bindParent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parent *ParentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parent.Contract.ParentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parent *ParentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parent.Contract.ParentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parent *ParentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parent.Contract.ParentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parent *ParentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parent *ParentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parent *ParentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parent.Contract.contract.Transact(opts, method, params...)
}

// BeverageLists is a free data retrieval call binding the contract method 0x3a784b32.
//
// Solidity: function beverageLists(bytes32 ) constant returns(address)
func (_Parent *ParentCaller) BeverageLists(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Parent.contract.Call(opts, out, "beverageLists", arg0)
	return *ret0, err
}

// BeverageLists is a free data retrieval call binding the contract method 0x3a784b32.
//
// Solidity: function beverageLists(bytes32 ) constant returns(address)
func (_Parent *ParentSession) BeverageLists(arg0 [32]byte) (common.Address, error) {
	return _Parent.Contract.BeverageLists(&_Parent.CallOpts, arg0)
}

// BeverageLists is a free data retrieval call binding the contract method 0x3a784b32.
//
// Solidity: function beverageLists(bytes32 ) constant returns(address)
func (_Parent *ParentCallerSession) BeverageLists(arg0 [32]byte) (common.Address, error) {
	return _Parent.Contract.BeverageLists(&_Parent.CallOpts, arg0)
}

// RegisterBeverageList is a paid mutator transaction binding the contract method 0x0b2cc919.
//
// Solidity: function registerBeverageList(bytes32 key_, address bvgrlAddress) returns()
func (_Parent *ParentTransactor) RegisterBeverageList(opts *bind.TransactOpts, key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Parent.contract.Transact(opts, "registerBeverageList", key_, bvgrlAddress)
}

// RegisterBeverageList is a paid mutator transaction binding the contract method 0x0b2cc919.
//
// Solidity: function registerBeverageList(bytes32 key_, address bvgrlAddress) returns()
func (_Parent *ParentSession) RegisterBeverageList(key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Parent.Contract.RegisterBeverageList(&_Parent.TransactOpts, key_, bvgrlAddress)
}

// RegisterBeverageList is a paid mutator transaction binding the contract method 0x0b2cc919.
//
// Solidity: function registerBeverageList(bytes32 key_, address bvgrlAddress) returns()
func (_Parent *ParentTransactorSession) RegisterBeverageList(key_ [32]byte, bvgrlAddress common.Address) (*types.Transaction, error) {
	return _Parent.Contract.RegisterBeverageList(&_Parent.TransactOpts, key_, bvgrlAddress)
}

// UpgradeBeverageList is a paid mutator transaction binding the contract method 0x44f338a0.
//
// Solidity: function upgradeBeverageList(bytes32 key_, address newOrgAddress) returns()
func (_Parent *ParentTransactor) UpgradeBeverageList(opts *bind.TransactOpts, key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Parent.contract.Transact(opts, "upgradeBeverageList", key_, newOrgAddress)
}

// UpgradeBeverageList is a paid mutator transaction binding the contract method 0x44f338a0.
//
// Solidity: function upgradeBeverageList(bytes32 key_, address newOrgAddress) returns()
func (_Parent *ParentSession) UpgradeBeverageList(key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Parent.Contract.UpgradeBeverageList(&_Parent.TransactOpts, key_, newOrgAddress)
}

// UpgradeBeverageList is a paid mutator transaction binding the contract method 0x44f338a0.
//
// Solidity: function upgradeBeverageList(bytes32 key_, address newOrgAddress) returns()
func (_Parent *ParentTransactorSession) UpgradeBeverageList(key_ [32]byte, newOrgAddress common.Address) (*types.Transaction, error) {
	return _Parent.Contract.UpgradeBeverageList(&_Parent.TransactOpts, key_, newOrgAddress)
}

// ParentBeverageListCreatedIterator is returned from FilterBeverageListCreated and is used to iterate over the raw logs and unpacked data for BeverageListCreated events raised by the Parent contract.
type ParentBeverageListCreatedIterator struct {
	Event *ParentBeverageListCreated // Event containing the contract specifics and raw log

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
func (it *ParentBeverageListCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBeverageListCreated)
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
		it.Event = new(ParentBeverageListCreated)
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
func (it *ParentBeverageListCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBeverageListCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBeverageListCreated represents a BeverageListCreated event raised by the Parent contract.
type ParentBeverageListCreated struct {
	BeverageList common.Address
	Now          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBeverageListCreated is a free log retrieval operation binding the contract event 0x6025a3db5babbc1d0b5c8e8e52a86caa417287a0ab307c4f03f38a46d15e29da.
//
// Solidity: event BeverageListCreated(address beverageList, uint256 now)
func (_Parent *ParentFilterer) FilterBeverageListCreated(opts *bind.FilterOpts) (*ParentBeverageListCreatedIterator, error) {

	logs, sub, err := _Parent.contract.FilterLogs(opts, "BeverageListCreated")
	if err != nil {
		return nil, err
	}
	return &ParentBeverageListCreatedIterator{contract: _Parent.contract, event: "BeverageListCreated", logs: logs, sub: sub}, nil
}

// WatchBeverageListCreated is a free log subscription operation binding the contract event 0x6025a3db5babbc1d0b5c8e8e52a86caa417287a0ab307c4f03f38a46d15e29da.
//
// Solidity: event BeverageListCreated(address beverageList, uint256 now)
func (_Parent *ParentFilterer) WatchBeverageListCreated(opts *bind.WatchOpts, sink chan<- *ParentBeverageListCreated) (event.Subscription, error) {

	logs, sub, err := _Parent.contract.WatchLogs(opts, "BeverageListCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBeverageListCreated)
				if err := _Parent.contract.UnpackLog(event, "BeverageListCreated", log); err != nil {
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

// ParentBeverageListUpgradedIterator is returned from FilterBeverageListUpgraded and is used to iterate over the raw logs and unpacked data for BeverageListUpgraded events raised by the Parent contract.
type ParentBeverageListUpgradedIterator struct {
	Event *ParentBeverageListUpgraded // Event containing the contract specifics and raw log

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
func (it *ParentBeverageListUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentBeverageListUpgraded)
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
		it.Event = new(ParentBeverageListUpgraded)
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
func (it *ParentBeverageListUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentBeverageListUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentBeverageListUpgraded represents a BeverageListUpgraded event raised by the Parent contract.
type ParentBeverageListUpgraded struct {
	BeverageList common.Address
	Now          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBeverageListUpgraded is a free log retrieval operation binding the contract event 0x844ceedb969cfa8c89cc0316f50d9ed69c40b9249b279a86ff4254416b00be82.
//
// Solidity: event BeverageListUpgraded(address beverageList, uint256 now)
func (_Parent *ParentFilterer) FilterBeverageListUpgraded(opts *bind.FilterOpts) (*ParentBeverageListUpgradedIterator, error) {

	logs, sub, err := _Parent.contract.FilterLogs(opts, "BeverageListUpgraded")
	if err != nil {
		return nil, err
	}
	return &ParentBeverageListUpgradedIterator{contract: _Parent.contract, event: "BeverageListUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeverageListUpgraded is a free log subscription operation binding the contract event 0x844ceedb969cfa8c89cc0316f50d9ed69c40b9249b279a86ff4254416b00be82.
//
// Solidity: event BeverageListUpgraded(address beverageList, uint256 now)
func (_Parent *ParentFilterer) WatchBeverageListUpgraded(opts *bind.WatchOpts, sink chan<- *ParentBeverageListUpgraded) (event.Subscription, error) {

	logs, sub, err := _Parent.contract.WatchLogs(opts, "BeverageListUpgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentBeverageListUpgraded)
				if err := _Parent.contract.UnpackLog(event, "BeverageListUpgraded", log); err != nil {
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
