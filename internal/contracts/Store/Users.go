// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package users

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// UsersABI is the input ABI used to generate the binding from.
const UsersABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getQtable\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"},{\"name\":\"\",\"type\":\"bytes32[2][]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"bytes32[2]\"}],\"name\":\"addState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"qvalArr\",\"type\":\"bytes32[2]\"}],\"name\":\"setQValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_st\",\"type\":\"string\"},{\"name\":\"_lr\",\"type\":\"bytes32\"},{\"name\":\"_gm\",\"type\":\"bytes32\"},{\"name\":\"_ep\",\"type\":\"bytes32\"},{\"name\":\"_epd\",\"type\":\"bytes32\"},{\"name\":\"_sr\",\"type\":\"uint256\"},{\"name\":\"_pd\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// UsersBin is the compiled bytecode used for deploying new contracts.
const UsersBin = `60806040523480156200001157600080fd5b5060405162002852380380620028528339810180604052620000379190810190620001a2565b6000878787878787876200004a62000107565b6200005c9796959493929190620002c6565b604051809103906000f08015801562000079573d6000803e3d6000fd5b509050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050505050506200041f565b60405161183d806200101583390190565b6000620001268251620003c4565b905092915050565b600082601f83011215156200014257600080fd5b815162000159620001538262000378565b6200034a565b915080825260208301602083018583830111156200017657600080fd5b62000183838284620003d8565b50505092915050565b60006200019a8251620003ce565b905092915050565b600080600080600080600060e0888a031215620001be57600080fd5b600088015167ffffffffffffffff811115620001d957600080fd5b620001e78a828b016200012e565b9750506020620001fa8a828b0162000118565b96505060406200020d8a828b0162000118565b9550506060620002208a828b0162000118565b9450506080620002338a828b0162000118565b93505060a0620002468a828b016200018c565b92505060c0620002598a828b016200018c565b91505092959891949750929550565b6200027381620003b0565b82525050565b60006200028682620003a5565b8084526200029c816020860160208601620003d8565b620002a7816200040e565b602085010191505092915050565b620002c081620003ba565b82525050565b600060e0820190508181036000830152620002e2818a62000279565b9050620002f3602083018962000268565b62000302604083018862000268565b62000311606083018762000268565b62000320608083018662000268565b6200032f60a0830185620002b5565b6200033e60c0830184620002b5565b98975050505050505050565b6000604051905081810181811067ffffffffffffffff821117156200036e57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200039057600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b6000819050919050565b6000819050919050565b6000819050919050565b6000819050919050565b60005b83811015620003f8578082015181840152602081019050620003db565b8381111562000408576000848401525b50505050565b6000601f19601f8301169050919050565b610be6806200042f6000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a45b69c1461005c578063413df77a146100885780638272546f146100b1575b600080fd5b34801561006857600080fd5b506100716100da565b60405161007f929190610921565b60405180910390f35b34801561009457600080fd5b506100af60048036036100aa91908101906106f3565b6101ea565b005b3480156100bd57600080fd5b506100d860048036036100d391908101906106f3565b6102db565b005b60608060008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16632a45b69c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156101a457600080fd5b505af11580156101b8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506101e19190810190610687565b92509250509091565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663413df77a84846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016102a4929190610958565b600060405180830381600087803b1580156102be57600080fd5b505af11580156102d2573d6000803e3d6000fd5b50505050505050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16638272546f84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610395929190610958565b600060405180830381600087803b1580156103af57600080fd5b505af11580156103c3573d6000803e3d6000fd5b50505050505050565b600082601f83011215156103df57600080fd5b81516103f26103ed826109b5565b610988565b9150818183526020840193506020810190508385604084028201111561041757600080fd5b60005b83811015610447578161042d8882610451565b84526020840193506040830192505060018101905061041a565b5050505092915050565b600082601f830112151561046457600080fd5b6002610477610472826109dd565b610988565b9150818385602084028201111561048d57600080fd5b60005b838110156104bd57816104a388826105c7565b845260208401935060208301925050600181019050610490565b5050505092915050565b600082601f83011215156104da57600080fd5b60026104ed6104e8826109ff565b610988565b9150818385602084028201111561050357600080fd5b60005b83811015610533578161051988826105b3565b845260208401935060208301925050600181019050610506565b5050505092915050565b600082601f830112151561055057600080fd5b815161056361055e82610a21565b610988565b9150818183526020840193506020810190508360005b838110156105a9578151860161058f88826105db565b845260208401935060208301925050600181019050610579565b5050505092915050565b60006105bf8235610b4f565b905092915050565b60006105d38251610b4f565b905092915050565b600082601f83011215156105ee57600080fd5b81516106016105fc82610a49565b610988565b9150808252602083016020830185838301111561061d57600080fd5b610628838284610b68565b50505092915050565b600082601f830112151561064457600080fd5b813561065761065282610a75565b610988565b9150808252602083016020830185838301111561067357600080fd5b61067e838284610b59565b50505092915050565b6000806040838503121561069a57600080fd5b600083015167ffffffffffffffff8111156106b457600080fd5b6106c08582860161053d565b925050602083015167ffffffffffffffff8111156106dd57600080fd5b6106e9858286016103cc565b9150509250929050565b6000806060838503121561070657600080fd5b600083013567ffffffffffffffff81111561072057600080fd5b61072c85828601610631565b925050602061073d858286016104c7565b9150509250929050565b600061075282610acf565b80845260208401935061076483610aa1565b60005b828110156107965761077a8683516107ed565b61078382610b11565b9150604086019550600181019050610767565b50849250505092915050565b6107ab81610ae5565b6107b482610ab8565b60005b828110156107e6576107ca8583516108a6565b6107d382610b2b565b91506020850194506001810190506107b7565b5050505050565b6107f681610ada565b6107ff82610aae565b60005b82811015610831576108158583516108a6565b61081e82610b1e565b9150602085019450600181019050610802565b5050505050565b600061084382610af0565b8084526020840193508360208202850161085c85610ac2565b60005b848110156108955783830388526108778383516108eb565b925061088282610b38565b915060208801975060018101905061085f565b508196508694505050505092915050565b6108af81610b45565b82525050565b60006108c082610b06565b8084526108d4816020860160208601610b68565b6108dd81610b9b565b602085010191505092915050565b60006108f682610afb565b80845261090a816020860160208601610b68565b61091381610b9b565b602085010191505092915050565b6000604082019050818103600083015261093b8185610838565b9050818103602083015261094f8184610747565b90509392505050565b6000606082019050818103600083015261097281856108b5565b905061098160208301846107a2565b9392505050565b6000604051905081810181811067ffffffffffffffff821117156109ab57600080fd5b8060405250919050565b600067ffffffffffffffff8211156109cc57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156109f457600080fd5b602082029050919050565b600067ffffffffffffffff821115610a1657600080fd5b602082029050919050565b600067ffffffffffffffff821115610a3857600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610a6057600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115610a8c57600080fd5b601f19601f8301169050602081019050919050565b6000602082019050919050565b6000819050919050565b6000819050919050565b6000602082019050919050565b600081519050919050565b600060029050919050565b600060029050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000819050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610b86578082015181840152602081019050610b6b565b83811115610b95576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a72305820f61512b3b08ebe01008ebcec712ed861772860ae261512945b3d8f1340bfbea46c6578706572696d656e74616cf5003760806040526040516200183d3803806200183d833981018060405262000029919081019062000674565b866000908051906020019062000041929190620004c8565b5085600181600019169055508460028160001916905550836003816000191690555082600481600019169055508160058190555080600681905550620000c387604080519081016040528060006001026000191660001916815260200160006001026000191660001916815250620000d0640100000000026401000000009004565b50505050505050620007df565b600080620001d5846009805480602002602001604051908101604052809291908181526020016000905b82821015620001bc578382906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015620001a75780601f106200017b57610100808354040283529160200191620001a7565b820191906000526020600020905b8154815290600101906020018083116200018957829003601f168201915b505050505081526020019060010190620000fa565b5050505062000282640100000000026401000000009004565b121515620001e257600080fd5b600983908060018154018082558091505090600182039060005260206000200160009091929091909150908051906020019062000221929190620004c8565b505060016009805490500390508160076000838152602001908152602001600020906002620002529291906200054f565b50826008600083815260200190815260200160002090805190602001906200027c929190620004c8565b50505050565b600080600090505b8251811015620002e457620002c78382815181101515620002a757fe5b90602001906020020151856200030f640100000000026401000000009004565b15620002d65780915062000308565b80806001019150506200028a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff91505b5092915050565b6000816040516020018082805190602001908083835b6020831015156200034c578051825260208201915060208101905060208303925062000325565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083101515620003b7578051825260208201915060208101905060208303925062000390565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916836040516020018082805190602001908083835b602083101515620004235780518252602082019150602081019050602083039250620003fc565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831015156200048e578051825260208201915060208101905060208303925062000467565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200050b57805160ff19168380011785556200053c565b828001600101855582156200053c579182015b828111156200053b5782518255916020019190600101906200051e565b5b5090506200054b91906200059a565b5090565b826002810192821562000587579160200282015b828111156200058657825182906000191690559160200191906001019062000563565b5b509050620005969190620005c2565b5090565b620005bf91905b80821115620005bb576000816000905550600101620005a1565b5090565b90565b620005e791905b80821115620005e3576000816000905550600101620005c9565b5090565b90565b6000620005f8825162000795565b905092915050565b600082601f83011215156200061457600080fd5b81516200062b620006258262000768565b6200073a565b915080825260208301602083018583830111156200064857600080fd5b62000655838284620007a9565b50505092915050565b60006200066c82516200079f565b905092915050565b600080600080600080600060e0888a0312156200069057600080fd5b600088015167ffffffffffffffff811115620006ab57600080fd5b620006b98a828b0162000600565b9750506020620006cc8a828b01620005ea565b9650506040620006df8a828b01620005ea565b9550506060620006f28a828b01620005ea565b9450506080620007058a828b01620005ea565b93505060a0620007188a828b016200065e565b92505060c06200072b8a828b016200065e565b91505092959891949750929550565b6000604051905081810181811067ffffffffffffffff821117156200075e57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200078057600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b6000819050919050565b60005b83811015620007c9578082015181840152602081019050620007ac565b83811115620007d9576000848401525b50505050565b61104e80620007ef6000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a45b69c14610067578063413df77a146100935780638272546f146100bc578063e86ef360146100e5575b600080fd5b34801561007357600080fd5b5061007c610122565b60405161008a929190610e5f565b60405180910390f35b34801561009f57600080fd5b506100ba60048036036100b59190810190610c76565b610393565b005b3480156100c857600080fd5b506100e360048036036100de9190810190610c76565b610527565b005b3480156100f157600080fd5b5061010c60048036036101079190810190610cca565b61066c565b6040516101199190610e96565b60405180910390f35b6060806000600b600061013591906108bb565b600c600061014391906108df565b600090505b60098054905081101561021557600c60098281548110151561016657fe5b9060005260206000200190806001815401808255809150509060018203906000526020600020016000909192909190915090805460018160011615610100020316600290046101b6929190610900565b5050600b6007600083815260200190815260200160002090806001815401808255809150509060018203906000526020600020906002020160009091929091909150906002610206929190610987565b50508080600101915050610148565b600c600b81805480602002602001604051908101604052809291908181526020016000905b828210156102f6578382906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102e25780601f106102b7576101008083540402835291602001916102e2565b820191906000526020600020905b8154815290600101906020018083116102c557829003601f168201915b50505050508152602001906001019061023a565b50505050915080805480602002602001604051908101604052809291908181526020016000905b828210156103845783829060005260206000209060020201600280602002604051908101604052809291908260028015610370576020028201915b81546000191681526020019060010190808311610358575b50505050508152602001906001019061031d565b50505050905092509250509091565b600080610481846009805480602002602001604051908101604052809291908181526020016000905b82821015610478578382906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104645780601f1061043957610100808354040283529160200191610464565b820191906000526020600020905b81548152906001019060200180831161044757829003601f168201915b5050505050815260200190600101906103bc565b50505050610693565b12151561048d57600080fd5b60098390806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906104ca9291906109c4565b5050600160098054905003905081600760008381526020019081526020016000209060026104f9929190610a44565b50826008600083815260200190815260200160002090805190602001906105219291906109c4565b50505050565b6000610614836009805480602002602001604051908101604052809291908181526020016000905b8282101561060b578382906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105f75780601f106105cc576101008083540402835291602001916105f7565b820191906000526020600020905b8154815290600101906020018083116105da57829003601f168201915b50505050508152602001906001019061054f565b50505050610693565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8113151561064457600080fd5b8160076000838152602001908152602001600020906002610666929190610a44565b50505050565b60076020528160005260406000208160028110151561068757fe5b01600091509150505481565b600080600090505b82518110156106df576106c583828151811015156106b557fe5b906020019060200201518561070a565b156106d257809150610703565b808060010191505061069b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff91505b5092915050565b6000816040516020018082805190602001908083835b6020831015156107455780518252602082019150602081019050602083039250610720565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831015156107ae5780518252602082019150602081019050602083039250610789565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916836040516020018082805190602001908083835b60208310151561081857805182526020820191506020810190506020830392506107f3565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083101515610881578051825260208201915060208101905060208303925061085c565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614905092915050565b50805460008255600202906000526020600020908101906108dc9190610a8a565b50565b50805460008255906000526020600020908101906108fd9190610ab6565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109395780548555610976565b8280016001018555821561097657600052602060002091601f016020900482015b8281111561097557825482559160010191906001019061095a565b5b5090506109839190610ae2565b5090565b82600281019282156109b3579182015b828111156109b2578254825591600101919060010190610997565b5b5090506109c09190610b07565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a0557805160ff1916838001178555610a33565b82800160010185558215610a33579182015b82811115610a32578251825591602001919060010190610a17565b5b509050610a409190610ae2565b5090565b8260028101928215610a79579160200282015b82811115610a78578251829060001916905591602001919060010190610a57565b5b509050610a869190610b07565b5090565b610ab391905b80821115610aaf5760008181610aa69190610b2c565b50600201610a90565b5090565b90565b610adf91905b80821115610adb5760008181610ad29190610b3a565b50600101610abc565b5090565b90565b610b0491905b80821115610b00576000816000905550600101610ae8565b5090565b90565b610b2991905b80821115610b25576000816000905550600101610b0d565b5090565b90565b506000815560010160009055565b50805460018160011615610100020316600290046000825580601f10610b605750610b7f565b601f016020900490600052602060002090810190610b7e9190610ae2565b5b50565b600082601f8301121515610b9557600080fd5b6002610ba8610ba382610ede565b610eb1565b91508183856020840282011115610bbe57600080fd5b60005b83811015610bee5781610bd48882610bf8565b845260208401935060208301925050600181019050610bc1565b5050505092915050565b6000610c048235610fad565b905092915050565b600082601f8301121515610c1f57600080fd5b8135610c32610c2d82610f00565b610eb1565b91508082526020830160208301858383011115610c4e57600080fd5b610c59838284610fc1565b50505092915050565b6000610c6e8235610fb7565b905092915050565b60008060608385031215610c8957600080fd5b600083013567ffffffffffffffff811115610ca357600080fd5b610caf85828601610c0c565b9250506020610cc085828601610b82565b9150509250929050565b60008060408385031215610cdd57600080fd5b6000610ceb85828601610c62565b9250506020610cfc85828601610c62565b9150509250929050565b6000610d1182610f50565b808452602084019350610d2383610f2c565b60005b82811015610d5557610d39868351610d61565b610d4282610f7c565b9150604086019550600181019050610d26565b50849250505092915050565b610d6a81610f5b565b610d7382610f39565b60005b82811015610da557610d89858351610e1a565b610d9282610f89565b9150602085019450600181019050610d76565b5050505050565b6000610db782610f66565b80845260208401935083602082028501610dd085610f43565b60005b84811015610e09578383038852610deb838351610e29565b9250610df682610f96565b9150602088019750600181019050610dd3565b508196508694505050505092915050565b610e2381610fa3565b82525050565b6000610e3482610f71565b808452610e48816020860160208601610fd0565b610e5181611003565b602085010191505092915050565b60006040820190508181036000830152610e798185610dac565b90508181036020830152610e8d8184610d06565b90509392505050565b6000602082019050610eab6000830184610e1a565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610ed457600080fd5b8060405250919050565b600067ffffffffffffffff821115610ef557600080fd5b602082029050919050565b600067ffffffffffffffff821115610f1757600080fd5b601f19601f8301169050602081019050919050565b6000602082019050919050565b6000819050919050565b6000602082019050919050565b600081519050919050565b600060029050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000819050919050565b6000819050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610fee578082015181840152602081019050610fd3565b83811115610ffd576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a7230582080fe2ca30593413b0dcb1f188adcbfa53a60e1210dd0f145c514dd3de93c328f6c6578706572696d656e74616cf50037`

// DeployUsers deploys a new Ethereum contract, binding an instance of Users to it.
func DeployUsers(auth *bind.TransactOpts, backend bind.ContractBackend, _st string, _lr [32]byte, _gm [32]byte, _ep [32]byte, _epd [32]byte, _sr *big.Int, _pd *big.Int) (common.Address, *types.Transaction, *Users, error) {
	parsed, err := abi.JSON(strings.NewReader(UsersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsersBin), backend, _st, _lr, _gm, _ep, _epd, _sr, _pd)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Users{UsersCaller: UsersCaller{contract: contract}, UsersTransactor: UsersTransactor{contract: contract}, UsersFilterer: UsersFilterer{contract: contract}}, nil
}

// Users is an auto generated Go binding around an Ethereum contract.
type Users struct {
	UsersCaller     // Read-only binding to the contract
	UsersTransactor // Write-only binding to the contract
	UsersFilterer   // Log filterer for contract events
}

// UsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsersSession struct {
	Contract     *Users            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsersCallerSession struct {
	Contract *UsersCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsersTransactorSession struct {
	Contract     *UsersTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsersRaw struct {
	Contract *Users // Generic contract binding to access the raw methods on
}

// UsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsersCallerRaw struct {
	Contract *UsersCaller // Generic read-only contract binding to access the raw methods on
}

// UsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsersTransactorRaw struct {
	Contract *UsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsers creates a new instance of Users, bound to a specific deployed contract.
func NewUsers(address common.Address, backend bind.ContractBackend) (*Users, error) {
	contract, err := bindUsers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Users{UsersCaller: UsersCaller{contract: contract}, UsersTransactor: UsersTransactor{contract: contract}, UsersFilterer: UsersFilterer{contract: contract}}, nil
}

// NewUsersCaller creates a new read-only instance of Users, bound to a specific deployed contract.
func NewUsersCaller(address common.Address, caller bind.ContractCaller) (*UsersCaller, error) {
	contract, err := bindUsers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsersCaller{contract: contract}, nil
}

// NewUsersTransactor creates a new write-only instance of Users, bound to a specific deployed contract.
func NewUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*UsersTransactor, error) {
	contract, err := bindUsers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsersTransactor{contract: contract}, nil
}

// NewUsersFilterer creates a new log filterer instance of Users, bound to a specific deployed contract.
func NewUsersFilterer(address common.Address, filterer bind.ContractFilterer) (*UsersFilterer, error) {
	contract, err := bindUsers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsersFilterer{contract: contract}, nil
}

// bindUsers binds a generic wrapper to an already deployed contract.
func bindUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Users.Contract.UsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Users.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.contract.Transact(opts, method, params...)
}

// AddState is a paid mutator transaction binding the contract method 0x413df77a.
//
// Solidity: function addState(st string, qvalArr bytes32[2]) returns()
func (_Users *UsersTransactor) AddState(opts *bind.TransactOpts, st string, qvalArr [2][32]byte) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "addState", st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0x413df77a.
//
// Solidity: function addState(st string, qvalArr bytes32[2]) returns()
func (_Users *UsersSession) AddState(st string, qvalArr [2][32]byte) (*types.Transaction, error) {
	return _Users.Contract.AddState(&_Users.TransactOpts, st, qvalArr)
}

// AddState is a paid mutator transaction binding the contract method 0x413df77a.
//
// Solidity: function addState(st string, qvalArr bytes32[2]) returns()
func (_Users *UsersTransactorSession) AddState(st string, qvalArr [2][32]byte) (*types.Transaction, error) {
	return _Users.Contract.AddState(&_Users.TransactOpts, st, qvalArr)
}

// GetQtable is a paid mutator transaction binding the contract method 0x2a45b69c.
//
// Solidity: function getQtable() returns(string[], bytes32[2][])
func (_Users *UsersTransactor) GetQtable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "getQtable")
}

// GetQtable is a paid mutator transaction binding the contract method 0x2a45b69c.
//
// Solidity: function getQtable() returns(string[], bytes32[2][])
func (_Users *UsersSession) GetQtable() (*types.Transaction, error) {
	return _Users.Contract.GetQtable(&_Users.TransactOpts)
}

// GetQtable is a paid mutator transaction binding the contract method 0x2a45b69c.
//
// Solidity: function getQtable() returns(string[], bytes32[2][])
func (_Users *UsersTransactorSession) GetQtable() (*types.Transaction, error) {
	return _Users.Contract.GetQtable(&_Users.TransactOpts)
}

// SetQValue is a paid mutator transaction binding the contract method 0x8272546f.
//
// Solidity: function setQValue(_st string, qvalArr bytes32[2]) returns()
func (_Users *UsersTransactor) SetQValue(opts *bind.TransactOpts, _st string, qvalArr [2][32]byte) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "setQValue", _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x8272546f.
//
// Solidity: function setQValue(_st string, qvalArr bytes32[2]) returns()
func (_Users *UsersSession) SetQValue(_st string, qvalArr [2][32]byte) (*types.Transaction, error) {
	return _Users.Contract.SetQValue(&_Users.TransactOpts, _st, qvalArr)
}

// SetQValue is a paid mutator transaction binding the contract method 0x8272546f.
//
// Solidity: function setQValue(_st string, qvalArr bytes32[2]) returns()
func (_Users *UsersTransactorSession) SetQValue(_st string, qvalArr [2][32]byte) (*types.Transaction, error) {
	return _Users.Contract.SetQValue(&_Users.TransactOpts, _st, qvalArr)
}