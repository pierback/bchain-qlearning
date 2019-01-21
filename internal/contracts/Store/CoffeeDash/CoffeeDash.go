// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coffeedash

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CoffeedashABI is the input ABI used to generate the binding from.
const CoffeedashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"states\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getQtableState\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string[2]\"}],\"name\":\"setQValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"string[2]\"}],\"name\":\"addState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"_lr\",\"type\":\"uint256\"},{\"name\":\"_gm\",\"type\":\"uint256\"},{\"name\":\"_ep\",\"type\":\"uint256\"},{\"name\":\"_epd\",\"type\":\"uint256\"},{\"name\":\"_sr\",\"type\":\"uint256\"},{\"name\":\"_pd\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// CoffeedashBin is the compiled bytecode used for deploying new contracts.
const CoffeedashBin = `6080604052604051620018ad380380620018ad8339810180604052620000299190810190620005f5565b6200003362000336565b60e060405190810160405280898152602001888152602001878152602001868152602001858152602001848152602001838152509050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190620000c692919062000374565b506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c082015181600601559050506200019f8860408051908101604052806040805190810160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525081526020016040805190810160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250815250620001ad640100000000026401000000009004565b505050505050505062000756565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600701846040518082805190602001908083835b6020831015156200022d578051825260208201915060208101905060208303925062000206565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902050600210156200026f57600080fd5b8181600701846040518082805190602001908083835b602083101515620002ac578051825260208201915060208101905060208303925062000285565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020906002620002ef929190620003fb565b5060018390806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906200032f92919062000455565b5050505050565b60e060405190810160405280606081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003b757805160ff1916838001178555620003e8565b82800160010185558215620003e8579182015b82811115620003e7578251825591602001919060010190620003ca565b5b509050620003f79190620004dc565b5090565b826002810192821562000442579160200282015b82811115620004415782518290805190602001906200043092919062000374565b50916020019190600101906200040f565b5b50905062000451919062000504565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200049857805160ff1916838001178555620004c9565b82800160010185558215620004c9579182015b82811115620004c8578251825591602001919060010190620004ab565b5b509050620004d89190620004dc565b5090565b6200050191905b80821115620004fd576000816000905550600101620004e3565b5090565b90565b6200053291905b808211156200052e576000818162000524919062000535565b506001016200050b565b5090565b90565b50805460018160011615610100020316600290046000825580601f106200055d57506200057e565b601f0160209004906000526020600020908101906200057d9190620004dc565b5b50565b600082601f83011215156200059557600080fd5b8151620005ac620005a682620006e9565b620006bb565b91508082526020830160208301858383011115620005c957600080fd5b620005d683828462000720565b50505092915050565b6000620005ed825162000716565b905092915050565b600080600080600080600060e0888a0312156200061157600080fd5b600088015167ffffffffffffffff8111156200062c57600080fd5b6200063a8a828b0162000581565b97505060206200064d8a828b01620005df565b9650506040620006608a828b01620005df565b9550506060620006738a828b01620005df565b9450506080620006868a828b01620005df565b93505060a0620006998a828b01620005df565b92505060c0620006ac8a828b01620005df565b91505092959891949750929550565b6000604051905081810181811067ffffffffffffffff82111715620006df57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200070157600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b60005b838110156200074057808201518184015260208101905062000723565b8381111562000750576000848401525b50505050565b61114780620007666000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063017a91051461007257806334acde82146100af578063a1cdeacf146100ed578063bdad5aa814610116578063daa050151461013f575b600080fd5b34801561007e57600080fd5b5061009960048036036100949190810190610e4a565b61016a565b6040516100a69190610f53565b60405180910390f35b3480156100bb57600080fd5b506100d660048036036100d19190810190610e4a565b610225565b6040516100e4929190610f75565b60405180910390f35b3480156100f957600080fd5b50610114600480360361010f9190810190610dde565b610491565b005b34801561012257600080fd5b5061013d60048036036101389190810190610dde565b6106ef565b005b34801561014b57600080fd5b5061015461086f565b6040516101619190610fac565b60405180910390f35b60018181548110151561017957fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561021d5780601f106101f25761010080835404028352916020019161021d565b820191906000526020600020905b81548152906001019060200180831161020057829003601f168201915b505050505081565b606061022f610aa4565b6000806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915060018581548110151561028157fe5b90600052602060002001905060018581548110151561029c57fe5b90600052602060002001826007018260405180828054600181600116156101000203166002900480156103065780601f106102e4576101008083540402835291820191610306565b820191906000526020600020905b8154815290600101906020018083116102f2575b50509150509081526020016040518091039020818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103ae5780601f10610383576101008083540402835291602001916103ae565b820191906000526020600020905b81548152906001019060200180831161039157829003601f168201915b5050505050915080600280602002604051908101604052809291906000905b82821015610480578382018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561046c5780601f106104415761010080835404028352916020019161046c565b820191906000526020600020905b81548152906001019060200180831161044f57829003601f168201915b5050505050815260200190600101906103cd565b505050509050935093505050915091565b60008061057f846001805480602002602001604051908101604052809291908181526020016000905b82821015610576578382906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105625780601f1061053757610100808354040283529160200191610562565b820191906000526020600020905b81548152906001019060200180831161054557829003601f168201915b5050505050815260200190600101906104ba565b5050505061087c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821315156105af57600080fd5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600701856040518082805190602001908083835b60208310151561062b5780518252602082019150602081019050602083039250610606565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020506002111561066c57600080fd5b8281600701856040518082805190602001908083835b6020831015156106a75780518252602082019150602081019050602083039250610682565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390209060026106e8929190610acb565b5050505050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600701846040518082805190602001908083835b60208310151561076d5780518252602082019150602081019050602083039250610748565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902050600210156107ae57600080fd5b8181600701846040518082805190602001908083835b6020831015156107e957805182526020820191506020810190506020830392506107c4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902090600261082a929190610acb565b506001839080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190610868929190610b1e565b5050505050565b6000600180549050905090565b600080600090505b82518110156108c8576108ae838281518110151561089e57fe5b90602001906020020151856108f3565b156108bb578091506108ec565b8080600101915050610884565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff91505b5092915050565b6000816040516020018082805190602001908083835b60208310151561092e5780518252602082019150602081019050602083039250610909565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831015156109975780518252602082019150602081019050602083039250610972565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916836040516020018082805190602001908083835b602083101515610a0157805182526020820191506020810190506020830392506109dc565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083101515610a6a5780518252602082019150602081019050602083039250610a45565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614905092915050565b60408051908101604052806002905b6060815260200190600190039081610ab35790505090565b8260028101928215610b0d579160200282015b82811115610b0c578251829080519060200190610afc929190610b9e565b5091602001919060010190610ade565b5b509050610b1a9190610c1e565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b5f57805160ff1916838001178555610b8d565b82800160010185558215610b8d579182015b82811115610b8c578251825591602001919060010190610b71565b5b509050610b9a9190610c4a565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bdf57805160ff1916838001178555610c0d565b82800160010185558215610c0d579182015b82811115610c0c578251825591602001919060010190610bf1565b5b509050610c1a9190610c4a565b5090565b610c4791905b80821115610c435760008181610c3a9190610c6f565b50600101610c24565b5090565b90565b610c6c91905b80821115610c68576000816000905550600101610c50565b5090565b90565b50805460018160011615610100020316600290046000825580601f10610c955750610cb4565b601f016020900490600052602060002090810190610cb39190610c4a565b5b50565b600082601f8301121515610cca57600080fd5b6002610cdd610cd882610ff4565b610fc7565b9150818360005b83811015610d145781358601610cfa8882610d1e565b845260208401935060208301925050600181019050610ce4565b5050505092915050565b600082601f8301121515610d3157600080fd5b8135610d44610d3f82611016565b610fc7565b91508082526020830160208301858383011115610d6057600080fd5b610d6b8382846110ba565b50505092915050565b600082601f8301121515610d8757600080fd5b8135610d9a610d9582611042565b610fc7565b91508082526020830160208301858383011115610db657600080fd5b610dc18382846110ba565b50505092915050565b6000610dd682356110b0565b905092915050565b60008060408385031215610df157600080fd5b600083013567ffffffffffffffff811115610e0b57600080fd5b610e1785828601610d74565b925050602083013567ffffffffffffffff811115610e3457600080fd5b610e4085828601610cb7565b9150509250929050565b600060208284031215610e5c57600080fd5b6000610e6a84828501610dca565b91505092915050565b6000610e7e82611078565b83602082028501610e8e8561106e565b60005b84811015610ec7578383038852610ea9838351610f0e565b9250610eb482611099565b9150602088019750600181019050610e91565b508196508694505050505092915050565b6000610ee38261108e565b808452610ef78160208601602086016110c9565b610f00816110fc565b602085010191505092915050565b6000610f1982611083565b808452610f2d8160208601602086016110c9565b610f36816110fc565b602085010191505092915050565b610f4d816110a6565b82525050565b60006020820190508181036000830152610f6d8184610f0e565b905092915050565b60006040820190508181036000830152610f8f8185610ed8565b90508181036020830152610fa38184610e73565b90509392505050565b6000602082019050610fc16000830184610f44565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610fea57600080fd5b8060405250919050565b600067ffffffffffffffff82111561100b57600080fd5b602082029050919050565b600067ffffffffffffffff82111561102d57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff82111561105957600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b600060029050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000819050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156110e75780820151818401526020810190506110cc565b838111156110f6576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a72305820636feec31cda027468ad9ca399f64ce7f95277466e70fe981db77adaef4dc9a86c6578706572696d656e74616cf50037`

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

// GetQtableState is a free data retrieval call binding the contract method 0x34acde82.
//
// Solidity: function getQtableState(i uint256) constant returns(string, string[2])
func (_Coffeedash *CoffeedashCaller) GetQtableState(opts *bind.CallOpts, i *big.Int) (string, [2]string, error) {
	var (
		ret0 = new(string)
		ret1 = new([2]string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Coffeedash.contract.Call(opts, out, "getQtableState", i)
	return *ret0, *ret1, err
}

// GetQtableState is a free data retrieval call binding the contract method 0x34acde82.
//
// Solidity: function getQtableState(i uint256) constant returns(string, string[2])
func (_Coffeedash *CoffeedashSession) GetQtableState(i *big.Int) (string, [2]string, error) {
	return _Coffeedash.Contract.GetQtableState(&_Coffeedash.CallOpts, i)
}

// GetQtableState is a free data retrieval call binding the contract method 0x34acde82.
//
// Solidity: function getQtableState(i uint256) constant returns(string, string[2])
func (_Coffeedash *CoffeedashCallerSession) GetQtableState(i *big.Int) (string, [2]string, error) {
	return _Coffeedash.Contract.GetQtableState(&_Coffeedash.CallOpts, i)
}

// GetStateCount is a free data retrieval call binding the contract method 0xdaa05015.
//
// Solidity: function getStateCount() constant returns(uint256)
func (_Coffeedash *CoffeedashCaller) GetStateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "getStateCount")
	return *ret0, err
}

// GetStateCount is a free data retrieval call binding the contract method 0xdaa05015.
//
// Solidity: function getStateCount() constant returns(uint256)
func (_Coffeedash *CoffeedashSession) GetStateCount() (*big.Int, error) {
	return _Coffeedash.Contract.GetStateCount(&_Coffeedash.CallOpts)
}

// GetStateCount is a free data retrieval call binding the contract method 0xdaa05015.
//
// Solidity: function getStateCount() constant returns(uint256)
func (_Coffeedash *CoffeedashCallerSession) GetStateCount() (*big.Int, error) {
	return _Coffeedash.Contract.GetStateCount(&_Coffeedash.CallOpts)
}

// States is a free data retrieval call binding the contract method 0x017a9105.
//
// Solidity: function states( uint256) constant returns(string)
func (_Coffeedash *CoffeedashCaller) States(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coffeedash.contract.Call(opts, out, "states", arg0)
	return *ret0, err
}

// States is a free data retrieval call binding the contract method 0x017a9105.
//
// Solidity: function states( uint256) constant returns(string)
func (_Coffeedash *CoffeedashSession) States(arg0 *big.Int) (string, error) {
	return _Coffeedash.Contract.States(&_Coffeedash.CallOpts, arg0)
}

// States is a free data retrieval call binding the contract method 0x017a9105.
//
// Solidity: function states( uint256) constant returns(string)
func (_Coffeedash *CoffeedashCallerSession) States(arg0 *big.Int) (string, error) {
	return _Coffeedash.Contract.States(&_Coffeedash.CallOpts, arg0)
}

// AddState is a paid mutator transaction binding the contract method 0xbdad5aa8.
//
// Solidity: function addState(_st string, qvalArr string[2]) returns()
func (_Coffeedash *CoffeedashTransactor) AddState(opts *bind.TransactOpts, _st string, qvalArr [2]string) (*types.Transaction, error) {
	return _Coffeedash.contract.Transact(opts, "addState", _st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0xbdad5aa8.
//
// Solidity: function addState(_st string, qvalArr string[2]) returns()
func (_Coffeedash *CoffeedashSession) AddState(_st string, qvalArr [2]string) (*types.Transaction, error) {
	return _Coffeedash.Contract.AddState(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0xbdad5aa8.
//
// Solidity: function addState(_st string, qvalArr string[2]) returns()
func (_Coffeedash *CoffeedashTransactorSession) AddState(_st string, qvalArr [2]string) (*types.Transaction, error) {
	return _Coffeedash.Contract.AddState(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0xa1cdeacf.
//
// Solidity: function setQValue(_st string, qvalArr string[2]) returns()
func (_Coffeedash *CoffeedashTransactor) SetQValue(opts *bind.TransactOpts, _st string, qvalArr [2]string) (*types.Transaction, error) {
	return _Coffeedash.contract.Transact(opts, "setQValue", _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0xa1cdeacf.
//
// Solidity: function setQValue(_st string, qvalArr string[2]) returns()
func (_Coffeedash *CoffeedashSession) SetQValue(_st string, qvalArr [2]string) (*types.Transaction, error) {
	return _Coffeedash.Contract.SetQValue(&_Coffeedash.TransactOpts, _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0xa1cdeacf.
//
// Solidity: function setQValue(_st string, qvalArr string[2]) returns()
func (_Coffeedash *CoffeedashTransactorSession) SetQValue(_st string, qvalArr [2]string) (*types.Transaction, error) {
	return _Coffeedash.Contract.SetQValue(&_Coffeedash.TransactOpts, _st, qvalArr)
}
