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
const ParentABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"createUserController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUserController\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ready\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"td\",\"type\":\"string\"}],\"name\":\"setTodo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"upgradeOrganisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"controllers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ParentBin is the compiled bytecode used for deploying new contracts.
const ParentBin = `608060405234801561001057600080fd5b50611cae806100206000396000f3fe608060405234801561001057600080fd5b506004361061007f576000357c0100000000000000000000000000000000000000000000000000000000900480631efebe0614610084578063315121891461008e5780636defbf80146100d8578063b9abac741461015b578063e300753f14610216578063fbae406a14610244575b600080fd5b61008c6102b2565b005b610096610388565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100e06103b2565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610120578082015181840152602081019050610105565b50505050905090810190601f16801561014d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102146004803603602081101561017157600080fd5b810190808035906020019064010000000081111561018e57600080fd5b8201836020820111156101a057600080fd5b803590602001918460018302840111640100000000831117156101c257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506104ce565b005b6102426004803603602081101561022c57600080fd5b81019080803590602001909291905050506105df565b005b6102706004803603602081101561025a57600080fd5b8101908080359060200190929190505050610778565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60006040516102c0906107ab565b604051809103906000f0801580156102dc573d6000803e3d6000fd5b5090506000816040516102ee906107b8565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610340573d6000803e3d6000fd5b50905080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636defbf806040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160006040518083038186803b15801561043857600080fd5b505afa15801561044c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561047657600080fd5b81019080805164010000000081111561048e57600080fd5b828101905060208101848111156104a457600080fd5b81518560018202830111640100000000821117156104c157600080fd5b5050929190505050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b9abac74826040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561057857808201518184015260208101905061055d565b50505050905090810190601f1680156105a55780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156105c457600080fd5b505af11580156105d8573d6000803e3d6000fd5b5050505050565b600080600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600081905060008173ffffffffffffffffffffffffffffffffffffffff166375e580796040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561067f57600080fd5b505afa158015610693573d6000803e3d6000fd5b505050506040513d60208110156106a957600080fd5b810190808051906020019092919050505090506000816040516106cb906107b8565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f08015801561071d573d6000803e3d6000fd5b5090508060008087815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6105d7806107c683390190565b610ee680610d9d8339019056fe608060405234801561001057600080fd5b506105b7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c0100000000000000000000000000000000000000000000000000000000900480636517579c116100785780636517579c146101f5578063760bb006146102635780637758374e146102a9578063ffcc7bbf146102e4576100a5565b806307973ccf146100aa5780631b62e209146100c857806326b158431461016157806355dbef30146101b0575b600080fd5b6100b2610326565b6040518082815260200191505060405180910390f35b610114600480360360408110156100de57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610332565b60405180831515151581526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b61019a6004803603604081101561017757600080fd5b8101908080359060200190929190803560ff169060200190929190505050610415565b6040518082815260200191505060405180910390f35b6101df600480360360208110156101c657600080fd5b81019080803560ff16906020019092919050505061044d565b6040518082815260200191505060405180910390f35b6102216004803603602081101561020b57600080fd5b8101908080359060200190929190505050610470565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61028f6004803603602081101561027957600080fd5b81019080803590602001909291905050506104b0565b604051808215151515815260200191505060405180910390f35b6102e2600480360360408110156102bf57600080fd5b8101908080359060200190929190803560ff169060200190929190505050610502565b005b610310600480360360208110156102fa57600080fd5b8101908080359060200190929190505050610569565b6040518082815260200191505060405180910390f35b60008080549050905090565b600080826001600086815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160008590806001815401808255809150509060018203906000526020600020016000909192909190915055036001600086815260200190815260200160002060020181905550600180600086815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915091509250929050565b60006001600084815260200190815260200160002060010160008360ff1660ff16815260200190815260200160002054905092915050565b6000600260008360ff1660ff168152602001908152602001600020549050919050565b60006001600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60008060008054905014156104c857600090506104fd565b81600060016000858152602001908152602001600020600201548154811015156104ee57fe5b90600052602060002001541490505b919050565b600180600084815260200190815260200160002060010160008360ff1660ff168152602001908152602001600020600082825401925050819055506001600260008360ff1660ff168152602001908152602001600020600082825401925050819055505050565b6000808281548110151561057957fe5b9060005260206000200154905091905056fea165627a7a72305820323c61e704b60b77e8c2342a3573943446ac22158644a7f391f8842252cbe9110029608060405234801561001057600080fd5b50604051602080610ee68339810180604052602081101561003057600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610e55806100916000396000f3fe608060405234801561001057600080fd5b50600436106100bb576000357c01000000000000000000000000000000000000000000000000000000009004806399cf037f1161008357806399cf037f146102dc578063a2c11dc41461032e578063b5cb15f71461038a578063b9abac74146103a8578063ffcc7bbf14610463576100bb565b80631b62e209146100c057806341aa8a76146101595780636517579c146101a15780636defbf801461020f57806375e5807914610292575b600080fd5b61010c600480360360408110156100d657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104a5565b60405180831515151581526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b61019f6004803603606081101561016f57600080fd5b8101908080359060200190929190803560ff169060200190929190803560ff169060200190929190505050610673565b005b6101cd600480360360208110156101b757600080fd5b8101908080359060200190929190505050610759565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102176108e3565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561025757808201518184015260208101905061023c565b50505050905090810190601f1680156102845780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61029a610985565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610318600480360360408110156102f257600080fd5b81019080803560ff169060200190929190803560ff1690602001909291905050506109aa565b6040518082815260200191505060405180910390f35b6103746004803603606081101561034457600080fd5b8101908080359060200190929190803560ff169060200190929190803560ff169060200190929190505050610aad565b6040518082815260200191505060405180910390f35b610392610bb9565b6040518082815260200191505060405180910390f35b610461600480360360208110156103be57600080fd5b81019080803590602001906401000000008111156103db57600080fd5b8201836020820111156103ed57600080fd5b8035906020019184600183028401116401000000008311171561040f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610c83565b005b61048f6004803603602081101561047957600080fd5b8101908080359060200190929190505050610c9d565b6040518082815260200191505060405180910390f35b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663760bb006866040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b15801561053b57600080fd5b505afa15801561054f573d6000803e3d6000fd5b505050506040513d602081101561056557600080fd5b81019080805190602001909291905050501561058057600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16631b62e20986866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001925050506040805180830381600087803b15801561062257600080fd5b505af1158015610636573d6000803e3d6000fd5b505050506040513d604081101561064c57600080fd5b81019080805190602001909291908051906020019092919050505092509250509250929050565b600061069583600181111561068457fe5b83600281111561069057fe5b610d74565b905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16637758374e86846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018260ff1660ff16815260200192505050600060405180830381600087803b15801561073a57600080fd5b505af115801561074e573d6000803e3d6000fd5b505050505050505050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663760bb006846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b1580156107ed57600080fd5b505afa158015610801573d6000803e3d6000fd5b505050506040513d602081101561081757600080fd5b8101908080519060200190929190505050151561083357600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16636517579c846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b1580156108a057600080fd5b505afa1580156108b4573d6000803e3d6000fd5b505050506040513d60208110156108ca57600080fd5b8101908080519060200190929190505050915050919050565b606060018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561097b5780601f106109505761010080835404028352916020019161097b565b820191906000526020600020905b81548152906001019060200180831161095e57829003601f168201915b5050505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806109cd8460018111156109bc57fe5b8460028111156109c857fe5b610d74565b905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166355dbef30836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260ff1660ff16815260200191505060206040518083038186803b158015610a6857600080fd5b505afa158015610a7c573d6000803e3d6000fd5b505050506040513d6020811015610a9257600080fd5b81019080805190602001909291905050509250505092915050565b600080610ad0846001811115610abf57fe5b846002811115610acb57fe5b610d74565b905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166326b1584387846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018260ff1660ff1681526020019250505060206040518083038186803b158015610b7357600080fd5b505afa158015610b87573d6000803e3d6000fd5b505050506040513d6020811015610b9d57600080fd5b8101908080519060200190929190505050925050509392505050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166307973ccf6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b158015610c4257600080fd5b505afa158015610c56573d6000803e3d6000fd5b505050506040513d6020811015610c6c57600080fd5b810190808051906020019092919050505091505090565b8060019080519060200190610c99929190610d84565b5050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663ffcc7bbf846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b158015610d3157600080fd5b505afa158015610d45573d6000803e3d6000fd5b505050506040513d6020811015610d5b57600080fd5b8101908080519060200190929190505050915050919050565b600081600a840201905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610dc557805160ff1916838001178555610df3565b82800160010185558215610df3579182015b82811115610df2578251825591602001919060010190610dd7565b5b509050610e009190610e04565b5090565b610e2691905b80821115610e22576000816000905550600101610e0a565b5090565b9056fea165627a7a723058202143b8a185d50a16089b26b15898e794db772d4be0eaa068b01990f6bf44f9440029a165627a7a723058205188e2a1a53977c112991fdc2dc8c27cd33462e06c06622dde9ef18df3caeed80029`

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

// Controllers is a free data retrieval call binding the contract method 0xfbae406a.
//
// Solidity: function controllers(bytes32 ) constant returns(address)
func (_Parent *ParentCaller) Controllers(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Parent.contract.Call(opts, out, "controllers", arg0)
	return *ret0, err
}

// Controllers is a free data retrieval call binding the contract method 0xfbae406a.
//
// Solidity: function controllers(bytes32 ) constant returns(address)
func (_Parent *ParentSession) Controllers(arg0 [32]byte) (common.Address, error) {
	return _Parent.Contract.Controllers(&_Parent.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xfbae406a.
//
// Solidity: function controllers(bytes32 ) constant returns(address)
func (_Parent *ParentCallerSession) Controllers(arg0 [32]byte) (common.Address, error) {
	return _Parent.Contract.Controllers(&_Parent.CallOpts, arg0)
}

// GetUserController is a free data retrieval call binding the contract method 0x31512189.
//
// Solidity: function getUserController() constant returns(address)
func (_Parent *ParentCaller) GetUserController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Parent.contract.Call(opts, out, "getUserController")
	return *ret0, err
}

// GetUserController is a free data retrieval call binding the contract method 0x31512189.
//
// Solidity: function getUserController() constant returns(address)
func (_Parent *ParentSession) GetUserController() (common.Address, error) {
	return _Parent.Contract.GetUserController(&_Parent.CallOpts)
}

// GetUserController is a free data retrieval call binding the contract method 0x31512189.
//
// Solidity: function getUserController() constant returns(address)
func (_Parent *ParentCallerSession) GetUserController() (common.Address, error) {
	return _Parent.Contract.GetUserController(&_Parent.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() constant returns(string)
func (_Parent *ParentCaller) Ready(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Parent.contract.Call(opts, out, "ready")
	return *ret0, err
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() constant returns(string)
func (_Parent *ParentSession) Ready() (string, error) {
	return _Parent.Contract.Ready(&_Parent.CallOpts)
}

// Ready is a free data retrieval call binding the contract method 0x6defbf80.
//
// Solidity: function ready() constant returns(string)
func (_Parent *ParentCallerSession) Ready() (string, error) {
	return _Parent.Contract.Ready(&_Parent.CallOpts)
}

// CreateUserController is a paid mutator transaction binding the contract method 0x1efebe06.
//
// Solidity: function createUserController() returns()
func (_Parent *ParentTransactor) CreateUserController(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parent.contract.Transact(opts, "createUserController")
}

// CreateUserController is a paid mutator transaction binding the contract method 0x1efebe06.
//
// Solidity: function createUserController() returns()
func (_Parent *ParentSession) CreateUserController() (*types.Transaction, error) {
	return _Parent.Contract.CreateUserController(&_Parent.TransactOpts)
}

// CreateUserController is a paid mutator transaction binding the contract method 0x1efebe06.
//
// Solidity: function createUserController() returns()
func (_Parent *ParentTransactorSession) CreateUserController() (*types.Transaction, error) {
	return _Parent.Contract.CreateUserController(&_Parent.TransactOpts)
}

// SetTodo is a paid mutator transaction binding the contract method 0xb9abac74.
//
// Solidity: function setTodo(string td) returns()
func (_Parent *ParentTransactor) SetTodo(opts *bind.TransactOpts, td string) (*types.Transaction, error) {
	return _Parent.contract.Transact(opts, "setTodo", td)
}

// SetTodo is a paid mutator transaction binding the contract method 0xb9abac74.
//
// Solidity: function setTodo(string td) returns()
func (_Parent *ParentSession) SetTodo(td string) (*types.Transaction, error) {
	return _Parent.Contract.SetTodo(&_Parent.TransactOpts, td)
}

// SetTodo is a paid mutator transaction binding the contract method 0xb9abac74.
//
// Solidity: function setTodo(string td) returns()
func (_Parent *ParentTransactorSession) SetTodo(td string) (*types.Transaction, error) {
	return _Parent.Contract.SetTodo(&_Parent.TransactOpts, td)
}

// UpgradeOrganisation is a paid mutator transaction binding the contract method 0xe300753f.
//
// Solidity: function upgradeOrganisation(bytes32 key_) returns()
func (_Parent *ParentTransactor) UpgradeOrganisation(opts *bind.TransactOpts, key_ [32]byte) (*types.Transaction, error) {
	return _Parent.contract.Transact(opts, "upgradeOrganisation", key_)
}

// UpgradeOrganisation is a paid mutator transaction binding the contract method 0xe300753f.
//
// Solidity: function upgradeOrganisation(bytes32 key_) returns()
func (_Parent *ParentSession) UpgradeOrganisation(key_ [32]byte) (*types.Transaction, error) {
	return _Parent.Contract.UpgradeOrganisation(&_Parent.TransactOpts, key_)
}

// UpgradeOrganisation is a paid mutator transaction binding the contract method 0xe300753f.
//
// Solidity: function upgradeOrganisation(bytes32 key_) returns()
func (_Parent *ParentTransactorSession) UpgradeOrganisation(key_ [32]byte) (*types.Transaction, error) {
	return _Parent.Contract.UpgradeOrganisation(&_Parent.TransactOpts, key_)
}
