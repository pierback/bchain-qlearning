// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coffeedash

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

// CoffeedashABI is the input ABI used to generate the binding from.
const CoffeedashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getStateValues\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"setQValue\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getQtableState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"addState\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"}],\"name\":\"stateExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStateCnt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"_lr\",\"type\":\"uint256\"},{\"name\":\"_gm\",\"type\":\"uint256\"},{\"name\":\"_ep\",\"type\":\"uint256\"},{\"name\":\"_epd\",\"type\":\"uint256\"},{\"name\":\"_sr\",\"type\":\"uint256\"},{\"name\":\"_pd\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// CoffeedashBin is the compiled bytecode used for deploying new contracts.
const CoffeedashBin = `608060405260405162001bcc38038062001bcc833981018060405260e08110156200002957600080fd5b8101908080516401000000008111156200004257600080fd5b828101905060208101848111156200005957600080fd5b81518560018202830111640100000000821117156200007757600080fd5b5050929190602001805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050606061012060405190810160405280898152602001888152602001878152602001868152602001858152602001848152602001838152602001600060ff168152602001828152506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190620001619291906200058b565b506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548160ff021916908360ff160217905550610100820151816008019080519060200190620001de92919062000612565b5090505062000232886040805190810160405280600481526020017f302c20300000000000000000000000000000000000000000000000000000000081525062000240640100000000026401000000009004565b5050505050505050620007a5565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b602083101515620002b9578051825260208201915060208101905060208303925062000292565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff16156200030857600080fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b6020831015156200038357805182526020820191506020810190506020830392506200035c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548160ff021916908315150217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b6020831015156200044a578051825260208201915060208101905060208303925062000423565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000190805190602001906200049592919062000679565b506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018290806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906200051492919062000679565b50506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff160217905550505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620005ce57805160ff1916838001178555620005ff565b82800160010185558215620005ff579182015b82811115620005fe578251825591602001919060010190620005e1565b5b5090506200060e919062000700565b5090565b82805482825590600052602060002090810192821562000666579160200282015b8281111562000665578251829080519060200190620006549291906200058b565b509160200191906001019062000633565b5b50905062000675919062000728565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620006bc57805160ff1916838001178555620006ed565b82800160010185558215620006ed579182015b82811115620006ec578251825591602001919060010190620006cf565b5b509050620006fc919062000700565b5090565b6200072591905b808211156200072157600081600090555060010162000707565b5090565b90565b6200075691905b8082111562000752576000818162000748919062000759565b506001016200072f565b5090565b90565b50805460018160011615610100020316600290046000825580601f10620007815750620007a2565b601f016020900490600052602060002090810190620007a1919062000700565b5b50565b61141780620007b56000396000f3fe60806040526004361061008c576000357c0100000000000000000000000000000000000000000000000000000000900480631672df681161006b5780631672df68146103515780632faf5168146104745780638f05ae83146105c6578063d46c97ba146106a65761008c565b8062d19f3014610091578063068531631461014857806312cfa7ce146101ff575b600080fd5b34801561009d57600080fd5b506100cd600480360360208110156100b457600080fd5b81019080803560ff1690602001909291905050506106d7565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561010d5780820151818401526020810190506100f2565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015457600080fd5b506101846004803603602081101561016b57600080fd5b81019080803560ff1690602001909291905050506107d4565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101c45780820151818401526020810190506101a9565b50505050905090810190601f1680156101f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61034f6004803603604081101561021557600080fd5b810190808035906020019064010000000081111561023257600080fd5b82018360208201111561024457600080fd5b8035906020019184600183028401116401000000008311171561026657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102c957600080fd5b8201836020820111156102db57600080fd5b803590602001918460018302840111640100000000831117156102fd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610a1b565b005b34801561035d57600080fd5b5061038d6004803603602081101561037457600080fd5b81019080803560ff169060200190929190505050610ba7565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156103d15780820151818401526020810190506103b6565b50505050905090810190601f1680156103fe5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561043757808201518184015260208101905061041c565b50505050905090810190601f1680156104645780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6105c46004803603604081101561048a57600080fd5b81019080803590602001906401000000008111156104a757600080fd5b8201836020820111156104b957600080fd5b803590602001918460018302840111640100000000831117156104db57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561053e57600080fd5b82018360208201111561055057600080fd5b8035906020019184600183028401116401000000008311171561057257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610eec565b005b3480156105d257600080fd5b5061068c600480360360208110156105e957600080fd5b810190808035906020019064010000000081111561060657600080fd5b82018360208201111561061857600080fd5b8035906020019184600183028401116401000000008311171561063a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061122c565b604051808215151515815260200191505060405180910390f35b3480156106b257600080fd5b506106bb6112f0565b604051808260ff1660ff16815260200191505060405180910390f35b60606000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018260ff1681548110151561072a57fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107c85780601f1061079d576101008083540402835291602001916107c8565b820191906000526020600020905b8154815290600101906020018083116107ab57829003601f168201915b50505050509050919050565b6060806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018360ff1681548110151561082857fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108c65780601f1061089b576101008083540402835291602001916108c6565b820191906000526020600020905b8154815290600101906020018083116108a957829003601f168201915b505050505090506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901816040518082805190602001908083835b602083101515610944578051825260208201915060208101905060208303925061091f565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a0e5780601f106109e357610100808354040283529160200191610a0e565b820191906000526020600020905b8154815290600101906020018083116109f157829003601f168201915b5050505050915050919050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b602083101515610a925780518252602082019150602081019050602083039250610a6d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff161515610ae157600080fd5b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b602083101515610b595780518252602082019150602081019050602083039250610b34565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000019080519060200190610ba2929190611346565b505050565b60608060606000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018460ff16815481101515610bfd57fe5b906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c9b5780601f10610c7057610100808354040283529160200191610c9b565b820191906000526020600020905b815481529060010190602001808311610c7e57829003601f168201915b505050505090506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018460ff16815481101515610cf357fe5b906000526020600020016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b602083101515610d745780518252602082019150602081019050602083039250610d4f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600001818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e3f5780601f10610e1457610100808354040283529160200191610e3f565b820191906000526020600020905b815481529060010190602001808311610e2257829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610edb5780601f10610eb057610100808354040283529160200191610edb565b820191906000526020600020905b815481529060010190602001808311610ebe57829003601f168201915b505050505090509250925050915091565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b602083101515610f635780518252602082019150602081019050602083039250610f3e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff1615610fb157600080fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b60208310151561102a5780518252602082019150602081019050602083039250611005565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548160ff021916908315150217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901836040518082805190602001908083835b6020831015156110ef57805182526020820191506020810190506020830392506110ca565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000019080519060200190611138929190611346565b506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008018290806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906111b5929190611346565b50506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600701600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff160217905550505050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600901826040518082805190602001908083835b6020831015156112a55780518252602082019150602081019050602083039250611280565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff169050919050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060070160009054906101000a900460ff16905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061138757805160ff19168380011785556113b5565b828001600101855582156113b5579182015b828111156113b4578251825591602001919060010190611399565b5b5090506113c291906113c6565b5090565b6113e891905b808211156113e45760008160009055506001016113cc565b5090565b9056fea165627a7a72305820bd039bedbd2d91f6b299dad5092a73e12d4dc80263b06a6b1c1a11fe883ef3420029`

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
// Solidity: function getQtableState(uint8 i) constant returns(string, string)
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
// Solidity: function getQtableState(uint8 i) constant returns(string, string)
func (_Coffeedash *CoffeedashSession) GetQtableState(i uint8) (string, string, error) {
	return _Coffeedash.Contract.GetQtableState(&_Coffeedash.CallOpts, i)
}

// GetQtableState is a free data retrieval call binding the contract method 0x1672df68.
//
// Solidity: function getQtableState(uint8 i) constant returns(string, string)
func (_Coffeedash *CoffeedashCallerSession) GetQtableState(i uint8) (string, string, error) {
	return _Coffeedash.Contract.GetQtableState(&_Coffeedash.CallOpts, i)
}

// GetState is a free data retrieval call binding the contract method 0x00d19f30.
//
// Solidity: function getState(uint8 i) constant returns(string)
func (_Coffeedash *CoffeedashCaller) GetState(opts *bind.CallOpts, i uint8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getState", i)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x00d19f30.
//
// Solidity: function getState(uint8 i) constant returns(string)
func (_Coffeedash *CoffeedashSession) GetState(i uint8) (string, error) {
	return _Coffeedash.Contract.GetState(&_Coffeedash.CallOpts, i)
}

// GetState is a free data retrieval call binding the contract method 0x00d19f30.
//
// Solidity: function getState(uint8 i) constant returns(string)
func (_Coffeedash *CoffeedashCallerSession) GetState(i uint8) (string, error) {
	return _Coffeedash.Contract.GetState(&_Coffeedash.CallOpts, i)
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashCaller) GetStateCnt(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getStateCnt")
	return *ret0, err
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashSession) GetStateCnt() (uint8, error) {
	return _Coffeedash.Contract.GetStateCnt(&_Coffeedash.CallOpts)
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint8)
func (_Coffeedash *CoffeedashCallerSession) GetStateCnt() (uint8, error) {
	return _Coffeedash.Contract.GetStateCnt(&_Coffeedash.CallOpts)
}

// GetStateValues is a free data retrieval call binding the contract method 0x06853163.
//
// Solidity: function getStateValues(uint8 i) constant returns(string)
func (_Coffeedash *CoffeedashCaller) GetStateValues(opts *bind.CallOpts, i uint8) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getStateValues", i)
	return *ret0, err
}

// GetStateValues is a free data retrieval call binding the contract method 0x06853163.
//
// Solidity: function getStateValues(uint8 i) constant returns(string)
func (_Coffeedash *CoffeedashSession) GetStateValues(i uint8) (string, error) {
	return _Coffeedash.Contract.GetStateValues(&_Coffeedash.CallOpts, i)
}

// GetStateValues is a free data retrieval call binding the contract method 0x06853163.
//
// Solidity: function getStateValues(uint8 i) constant returns(string)
func (_Coffeedash *CoffeedashCallerSession) GetStateValues(i uint8) (string, error) {
	return _Coffeedash.Contract.GetStateValues(&_Coffeedash.CallOpts, i)
}

// StateExists is a free data retrieval call binding the contract method 0x8f05ae83.
//
// Solidity: function stateExists(string _st) constant returns(bool)
func (_Coffeedash *CoffeedashCaller) StateExists(opts *bind.CallOpts, _st string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "stateExists", _st)
	return *ret0, err
}

// StateExists is a free data retrieval call binding the contract method 0x8f05ae83.
//
// Solidity: function stateExists(string _st) constant returns(bool)
func (_Coffeedash *CoffeedashSession) StateExists(_st string) (bool, error) {
	return _Coffeedash.Contract.StateExists(&_Coffeedash.CallOpts, _st)
}

// StateExists is a free data retrieval call binding the contract method 0x8f05ae83.
//
// Solidity: function stateExists(string _st) constant returns(bool)
func (_Coffeedash *CoffeedashCallerSession) StateExists(_st string) (bool, error) {
	return _Coffeedash.Contract.StateExists(&_Coffeedash.CallOpts, _st)
}

// AddState is a paid mutator transaction binding the contract method 0x2faf5168.
//
// Solidity: function addState(string _st, string qvalArr) returns()
func (_Coffeedash *CoffeedashTransactor) AddState(opts *bind.TransactOpts, _st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.contract.Transact(opts, "addState", _st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0x2faf5168.
//
// Solidity: function addState(string _st, string qvalArr) returns()
func (_Coffeedash *CoffeedashSession) AddState(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.AddState(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0x2faf5168.
//
// Solidity: function addState(string _st, string qvalArr) returns()
func (_Coffeedash *CoffeedashTransactorSession) AddState(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.AddState(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x12cfa7ce.
//
// Solidity: function setQValue(string _st, string qvalArr) returns()
func (_Coffeedash *CoffeedashTransactor) SetQValue(opts *bind.TransactOpts, _st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.contract.Transact(opts, "setQValue", _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x12cfa7ce.
//
// Solidity: function setQValue(string _st, string qvalArr) returns()
func (_Coffeedash *CoffeedashSession) SetQValue(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.SetQValue(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x12cfa7ce.
//
// Solidity: function setQValue(string _st, string qvalArr) returns()
func (_Coffeedash *CoffeedashTransactorSession) SetQValue(_st string, qvalArr string) (*types.Transaction, error) {
	return _Coffeedash.Contract.SetQValue(&_Coffeedash.TransactOpts, _st, qvalArr)
}
