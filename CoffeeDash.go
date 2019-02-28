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
const CoffeedashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getStateValues\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"setQValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getQtableState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string\"}],\"name\":\"addState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"}],\"name\":\"stateExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStateCnt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"_lr\",\"type\":\"uint256\"},{\"name\":\"_gm\",\"type\":\"uint256\"},{\"name\":\"_ep\",\"type\":\"uint256\"},{\"name\":\"_epd\",\"type\":\"uint256\"},{\"name\":\"_sr\",\"type\":\"uint256\"},{\"name\":\"_pd\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// CoffeedashBin is the compiled bytecode used for deploying new contracts.
const CoffeedashBin = `608060405260405162001a7238038062001a72833981018060405260e08110156200002957600080fd5b8101908080516401000000008111156200004257600080fd5b828101905060208101848111156200005957600080fd5b81518560018202830111640100000000821117156200007757600080fd5b5050929190602001805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050606061012060405190810160405280898152602001888152602001878152602001868152602001858152602001848152602001838152602001600060ff168152602001828152506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906200016192919062000509565b506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548160ff021916908360ff160217905550610100820151816008019080519060200190620001de92919062000590565b5090505062000232886040805190810160405280600481526020017f302c20300000000000000000000000000000000000000000000000000000000081525062000240640100000000026401000000009004565b505050505050505062000723565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020826040518082805190602001908083835b602083101515620002b7578051825260208201915060208101905060208303925062000290565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff16156200030357600080fd5b6001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020836040518082805190602001908083835b6020831015156200037c578051825260208201915060208101905060208303925062000355565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff02191690831515021790555080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020836040518082805190602001908083835b6020831015156200043e578051825260208201915060208101905060208303925062000417565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020908051906020019062000486929190620005f7565b50600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082908060018154018082558091505090600182039060005260206000200160009091929091909150908051906020019062000503929190620005f7565b50505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200054c57805160ff19168380011785556200057d565b828001600101855582156200057d579182015b828111156200057c5782518255916020019190600101906200055f565b5b5090506200058c91906200067e565b5090565b828054828255906000526020600020908101928215620005e4579160200282015b82811115620005e3578251829080519060200190620005d292919062000509565b5091602001919060010190620005b1565b5b509050620005f39190620006a6565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200063a57805160ff19168380011785556200066b565b828001600101855582156200066b579182015b828111156200066a5782518255916020019190600101906200064d565b5b5090506200067a91906200067e565b5090565b620006a391905b808211156200069f57600081600090555060010162000685565b5090565b90565b620006d491905b80821115620006d05760008181620006c69190620006d7565b50600101620006ad565b5090565b90565b50805460018160011615610100020316600290046000825580601f10620006ff575062000720565b601f0160209004906000526020600020908101906200071f91906200067e565b5b50565b61133f80620007336000396000f3fe608060405234801561001057600080fd5b5060043610610099576000357c0100000000000000000000000000000000000000000000000000000000900480631672df68116100785780631672df68146103445780632faf51681461045a5780638f05ae83146105ac578063d46c97ba1461067f57610099565b8062d19f301461009e578063068531631461014857806312cfa7ce146101f2575b600080fd5b6100cd600480360360208110156100b457600080fd5b81019080803560ff16906020019092919050505061069d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561010d5780820151818401526020810190506100f2565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101776004803603602081101561015e57600080fd5b81019080803560ff169060200190929190505050610798565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101b757808201518184015260208101905061019c565b50505050905090810190601f1680156101e45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103426004803603604081101561020857600080fd5b810190808035906020019064010000000081111561022557600080fd5b82018360208201111561023757600080fd5b8035906020019184600183028401116401000000008311171561025957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102bc57600080fd5b8201836020820111156102ce57600080fd5b803590602001918460018302840111640100000000831117156102f057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506109d8565b005b6103736004803603602081101561035a57600080fd5b81019080803560ff169060200190929190505050610b5a565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156103b757808201518184015260208101905061039c565b50505050905090810190601f1680156103e45780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561041d578082015181840152602081019050610402565b50505050905090810190601f16801561044a5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6105aa6004803603604081101561047057600080fd5b810190808035906020019064010000000081111561048d57600080fd5b82018360208201111561049f57600080fd5b803590602001918460018302840111640100000000831117156104c157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561052457600080fd5b82018360208201111561053657600080fd5b8035906020019184600183028401116401000000008311171561055857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e96565b005b610665600480360360208110156105c257600080fd5b81019080803590602001906401000000008111156105df57600080fd5b8201836020820111156105f157600080fd5b8035906020019184600183028401116401000000008311171561061357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611154565b604051808215151515815260200191505060405180910390f35b610687611224565b6040518082815260200191505060405180910390f35b6060600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208260ff168154811015156106ee57fe5b906000526020600020018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561078c5780601f106107615761010080835404028352916020019161078c565b820191906000526020600020905b81548152906001019060200180831161076f57829003601f168201915b50505050509050919050565b606080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208360ff168154811015156107ea57fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108885780601f1061085d57610100808354040283529160200191610888565b820191906000526020600020905b81548152906001019060200180831161086b57829003601f168201915b50505050509050600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020816040518082805190602001908083835b60208310151561090457805182526020820191506020810190506020830392506108df565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109cb5780601f106109a0576101008083540402835291602001916109cb565b820191906000526020600020905b8154815290600101906020018083116109ae57829003601f168201915b5050505050915050919050565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020826040518082805190602001908083835b602083101515610a4d5780518252602082019150602081019050602083039250610a28565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff161515610a9957600080fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020836040518082805190602001908083835b602083101515610b0f5780518252602082019150602081019050602083039250610aea565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390209080519060200190610b5592919061126e565b505050565b6060806060600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208460ff16815481101515610bae57fe5b906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c4c5780601f10610c2157610100808354040283529160200191610c4c565b820191906000526020600020905b815481529060010190602001808311610c2f57829003601f168201915b50505050509050600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208460ff16815481101515610ca257fe5b90600052602060002001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020826040518082805190602001908083835b602083101515610d215780518252602082019150602081019050602083039250610cfc565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610de95780601f10610dbe57610100808354040283529160200191610de9565b820191906000526020600020905b815481529060010190602001808311610dcc57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e855780601f10610e5a57610100808354040283529160200191610e85565b820191906000526020600020905b815481529060010190602001808311610e6857829003601f168201915b505050505090509250925050915091565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020826040518082805190602001908083835b602083101515610f0b5780518252602082019150602081019050602083039250610ee6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff1615610f5657600080fd5b6001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020836040518082805190602001908083835b602083101515610fcd5780518252602082019150602081019050602083039250610fa8565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff02191690831515021790555080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020836040518082805190602001908083835b60208310151561108d5780518252602082019150602081019050602083039250611068565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902090805190602001906110d392919061126e565b50600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082908060018154018082558091505090600182039060005260206000200160009091929091909150908051906020019061114e92919061126e565b50505050565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020826040518082805190602001908083835b6020831015156111cb57805182526020820191506020810190506020830392506111a6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff161561121a576001905061121f565b600090505b919050565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112af57805160ff19168380011785556112dd565b828001600101855582156112dd579182015b828111156112dc5782518255916020019190600101906112c1565b5b5090506112ea91906112ee565b5090565b61131091905b8082111561130c5760008160009055506001016112f4565b5090565b9056fea165627a7a723058203948c009e478028565975e3c56382ed88b0edf5e048cc5aed79fcbe4708173f80029`

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
// Solidity: function getStateCnt() constant returns(uint256)
func (_Coffeedash *CoffeedashCaller) GetStateCnt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getStateCnt")
	return *ret0, err
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint256)
func (_Coffeedash *CoffeedashSession) GetStateCnt() (*big.Int, error) {
	return _Coffeedash.Contract.GetStateCnt(&_Coffeedash.CallOpts)
}

// GetStateCnt is a free data retrieval call binding the contract method 0xd46c97ba.
//
// Solidity: function getStateCnt() constant returns(uint256)
func (_Coffeedash *CoffeedashCallerSession) GetStateCnt() (*big.Int, error) {
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
