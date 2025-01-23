// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OrganizationRegistryMetaData contains all meta data concerning the OrganizationRegistry contract.
var OrganizationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"OrganizationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"}],\"name\":\"OrganizationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"addAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"administrators\",\"type\":\"address[]\"}],\"name\":\"createOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"deleteOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"getOrganization\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdministrator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"organizations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"removeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015601357600080fd5b50608051611a2461003d60003960008181610e2c01528181610e550152610f9b0152611a246000f3fe6080604052600436106100f35760003560e01c80638129fc1c1161008a578063c2a950be11610059578063c2a950be146102ed578063d2c30a6d1461030d578063f1c621041461032d578063f2fde38b1461035f57600080fd5b80638129fc1c146101f85780638da5cb5b1461020d578063ad3cb1cc14610254578063c1af6e031461029257600080fd5b80635a1f7406116100c65780635a1f7406146101755780636cca67bf146101a3578063715018a6146101c35780637acbb8af146101d857600080fd5b80631c2e3d82146100f85780633c10eee51461011a5780634f1ef2861461013a57806352d1902d1461014d575b600080fd5b34801561010457600080fd5b50610118610113366004611351565b61037f565b005b34801561012657600080fd5b50610118610135366004611351565b61046d565b61011861014836600461139a565b61054f565b34801561015957600080fd5b5061016261056e565b6040519081526020015b60405180910390f35b34801561018157600080fd5b50610195610190366004611461565b61058b565b60405161016c9291906114cc565b3480156101af57600080fd5b506101186101be366004611539565b6106b7565b3480156101cf57600080fd5b5061011861082f565b3480156101e457600080fd5b506101186101f3366004611461565b610843565b34801561020457600080fd5b506101186108f8565b34801561021957600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b03909116815260200161016c565b34801561026057600080fd5b50610285604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161016c91906115bd565b34801561029e57600080fd5b506102dd6102ad366004611351565b6001600160a01b039182166000908152602081815260408083209390941682526002909201909152205460ff1690565b604051901515815260200161016c565b3480156102f957600080fd5b506101186103083660046115d0565b610a0d565b34801561031957600080fd5b50610195610328366004611461565b610c9c565b34801561033957600080fd5b5060015461034a9063ffffffff1681565b60405163ffffffff909116815260200161016c565b34801561036b57600080fd5b5061011861037a366004611461565b610de3565b6001600160a01b038216600090815260208181526040808320338452600201909152902054829060ff166103ce5760405162461bcd60e51b81526004016103c5906116b8565b60405180910390fd5b6001600160a01b038316600090815260208190526040812080546103f190611702565b9050116104105760405162461bcd60e51b81526004016103c59061173c565b6001600160a01b0382166104365760405162461bcd60e51b81526004016103c59061178d565b506001600160a01b03918216600090815260208181526040808320939094168252600290920190915220805460ff19166001179055565b6001600160a01b038216600090815260208181526040808320338452600201909152902054829060ff166104b35760405162461bcd60e51b81526004016103c5906116b8565b6001600160a01b038316600090815260208190526040812080546104d690611702565b9050116104f55760405162461bcd60e51b81526004016103c59061173c565b6001600160a01b03821661051b5760405162461bcd60e51b81526004016103c59061178d565b506001600160a01b03918216600090815260208181526040808320939094168252600290920190915220805460ff19169055565b610557610e21565b61056082610ec6565b61056a8282610ece565b5050565b6000610578610f90565b506000805160206119cf83398151915290565b6000602081905290815260409020805481906105a690611702565b80601f01602080910402602001604051908101604052809291908181526020018280546105d290611702565b801561061f5780601f106105f45761010080835404028352916020019161061f565b820191906000526020600020905b81548152906001019060200180831161060257829003601f168201915b50505050509080600101805461063490611702565b80601f016020809104026020016040519081016040528092919081815260200182805461066090611702565b80156106ad5780601f10610682576101008083540402835291602001916106ad565b820191906000526020600020905b81548152906001019060200180831161069057829003601f168201915b5050505050905082565b6001600160a01b038516600090815260208181526040808320338452600201909152902054859060ff166106fd5760405162461bcd60e51b81526004016103c5906116b8565b8361071a5760405162461bcd60e51b81526004016103c5906117e0565b816107795760405162461bcd60e51b815260206004820152602960248201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c6964206d6560448201526874616461746155524960b81b60648201526084016103c5565b6001600160a01b0386166000908152602081905260408120805461079c90611702565b9050116107bb5760405162461bcd60e51b81526004016103c59061173c565b6001600160a01b0386166000908152602081905260409020806107df868883611869565b50600181016107ef848683611869565b5060405133906001600160a01b038916907fdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f690600090a350505050505050565b610837610fd9565b6108416000611034565b565b61084b610fd9565b6001600160a01b0381166000908152602081905260408120805461086e90611702565b90501161088d5760405162461bcd60e51b81526004016103c59061173c565b6001600160a01b0381166000908152602081905260408120906108b082826112e7565b6108be6001830160006112e7565b50506001805463ffffffff169060006108d68361193e565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b031660008115801561093d5750825b90506000826001600160401b031660011480156109595750303b155b905081158015610967575080155b156109855760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156109af57845460ff60401b1916600160401b1785555b6109b8336110a5565b6109c06110b6565b8315610a0657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050565b6001600160a01b038716610a635760405162461bcd60e51b815260206004820181905260248201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c696420696460448201526064016103c5565b84610a805760405162461bcd60e51b81526004016103c5906117e0565b6001600160a01b03871660009081526020819052604090208054610aa390611702565b159050610b0c5760405162461bcd60e51b815260206004820152603160248201527f4f7267616e697a6174696f6e52656769737472793a206f7267616e697a6174696044820152706f6e20616c72656164792065786973747360781b60648201526084016103c5565b6001600160a01b038716600090815260208190526040902080610b30878983611869565b5060018101610b40858783611869565b508115610c025760005b82811015610c00576000848483818110610b6657610b6661195e565b9050602002016020810190610b7b9190611461565b6001600160a01b031603610ba15760405162461bcd60e51b81526004016103c59061178d565b6001826002016000868685818110610bbb57610bbb61195e565b9050602002016020810190610bd09190611461565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055600101610b4a565b505b3360009081526002820160205260408120805460ff19166001908117909155805463ffffffff1691610c3383611974565b91906101000a81548163ffffffff021916908363ffffffff16021790555050336001600160a01b0316886001600160a01b03167f2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a60405160405180910390a35050505050505050565b6001600160a01b038116600090815260208190526040902080546060918291819060018201908290610ccd90611702565b80601f0160208091040260200160405190810160405280929190818152602001828054610cf990611702565b8015610d465780601f10610d1b57610100808354040283529160200191610d46565b820191906000526020600020905b815481529060010190602001808311610d2957829003601f168201915b50505050509150808054610d5990611702565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8590611702565b8015610dd25780601f10610da757610100808354040283529160200191610dd2565b820191906000526020600020905b815481529060010190602001808311610db557829003601f168201915b505050505090509250925050915091565b610deb610fd9565b6001600160a01b038116610e1557604051631e4fbdf760e01b8152600060048201526024016103c5565b610e1e81611034565b50565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610ea857507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610e9c6000805160206119cf833981519152546001600160a01b031690565b6001600160a01b031614155b156108415760405163703e46dd60e11b815260040160405180910390fd5b610e1e610fd9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610f28575060408051601f3d908101601f19168201909252610f2591810190611999565b60015b610f5057604051634c9c8ce360e01b81526001600160a01b03831660048201526024016103c5565b6000805160206119cf8339815191528114610f8157604051632a87526960e21b8152600481018290526024016103c5565b610f8b83836110be565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108415760405163703e46dd60e11b815260040160405180910390fd5b3361100b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108415760405163118cdaa760e01b81523360048201526024016103c5565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6110ad611114565b610e1e8161115d565b610841611114565b6110c782611165565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561110c57610f8b82826111ca565b61056a611240565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661084157604051631afcd79f60e31b815260040160405180910390fd5b610deb611114565b806001600160a01b03163b60000361119b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024016103c5565b6000805160206119cf83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516111e791906119b2565b600060405180830381855af49150503d8060008114611222576040519150601f19603f3d011682016040523d82523d6000602084013e611227565b606091505b509150915061123785838361125f565b95945050505050565b34156108415760405163b398979f60e01b815260040160405180910390fd5b6060826112745761126f826112be565b6112b7565b815115801561128b57506001600160a01b0384163b155b156112b457604051639996b31560e01b81526001600160a01b03851660048201526024016103c5565b50805b9392505050565b8051156112ce5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5080546112f390611702565b6000825580601f10611303575050565b601f016020900490600052602060002090810190610e1e91905b80821115611331576000815560010161131d565b5090565b80356001600160a01b038116811461134c57600080fd5b919050565b6000806040838503121561136457600080fd5b61136d83611335565b915061137b60208401611335565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156113ad57600080fd5b6113b683611335565b915060208301356001600160401b038111156113d157600080fd5b8301601f810185136113e257600080fd5b80356001600160401b038111156113fb576113fb611384565b604051601f8201601f19908116603f011681016001600160401b038111828210171561142957611429611384565b60405281815282820160200187101561144157600080fd5b816020840160208301376000602083830101528093505050509250929050565b60006020828403121561147357600080fd5b6112b782611335565b60005b8381101561149757818101518382015260200161147f565b50506000910152565b600081518084526114b881602086016020860161147c565b601f01601f19169290920160200192915050565b6040815260006114df60408301856114a0565b828103602084015261123781856114a0565b60008083601f84011261150357600080fd5b5081356001600160401b0381111561151a57600080fd5b60208301915083602082850101111561153257600080fd5b9250929050565b60008060008060006060868803121561155157600080fd5b61155a86611335565b945060208601356001600160401b0381111561157557600080fd5b611581888289016114f1565b90955093505060408601356001600160401b038111156115a057600080fd5b6115ac888289016114f1565b969995985093965092949392505050565b6020815260006112b760208301846114a0565b60008060008060008060006080888a0312156115eb57600080fd5b6115f488611335565b965060208801356001600160401b0381111561160f57600080fd5b61161b8a828b016114f1565b90975095505060408801356001600160401b0381111561163a57600080fd5b6116468a828b016114f1565b90955093505060608801356001600160401b0381111561166557600080fd5b8801601f81018a1361167657600080fd5b80356001600160401b0381111561168c57600080fd5b8a60208260051b84010111156116a157600080fd5b602082019350809250505092959891949750929550565b6020808252602a908201527f4f7267616e697a6174696f6e52656769737472793a206e6f7420616e2061646d60408201526934b734b9ba3930ba37b960b11b606082015260800190565b600181811c9082168061171657607f821691505b60208210810361173657634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526031908201527f4f7267616e697a6174696f6e52656769737472793a206f7267616e697a6174696040820152701bdb88191bd95cc81b9bdd08195e1a5cdd607a1b606082015260800190565b60208082526033908201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c69642061646040820152726d696e6973747261746f72206164647265737360681b606082015260800190565b60208082526022908201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c6964206e616040820152616d6560f01b606082015260800190565b601f821115610f8b57806000526020600020601f840160051c810160208510156118495750805b601f840160051c820191505b81811015610a065760008155600101611855565b6001600160401b0383111561188057611880611384565b6118948361188e8354611702565b83611822565b6000601f8411600181146118c857600085156118b05750838201355b600019600387901b1c1916600186901b178355610a06565b600083815260209020601f19861690835b828110156118f957868501358255602094850194600190920191016118d9565b50868210156119165760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff82168061195457611954611928565b6000190192915050565b634e487b7160e01b600052603260045260246000fd5b600063ffffffff821663ffffffff810361199057611990611928565b60010192915050565b6000602082840312156119ab57600080fd5b5051919050565b600082516119c481846020870161147c565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212205f7d26eb50828ee2062bc7571f3b0cf748990acb37db9d375af174d7e8e6202864736f6c634300081c0033",
}

// OrganizationRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OrganizationRegistryMetaData.ABI instead.
var OrganizationRegistryABI = OrganizationRegistryMetaData.ABI

// OrganizationRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrganizationRegistryMetaData.Bin instead.
var OrganizationRegistryBin = OrganizationRegistryMetaData.Bin

// DeployOrganizationRegistry deploys a new Ethereum contract, binding an instance of OrganizationRegistry to it.
func DeployOrganizationRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrganizationRegistry, error) {
	parsed, err := OrganizationRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrganizationRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrganizationRegistry{OrganizationRegistryCaller: OrganizationRegistryCaller{contract: contract}, OrganizationRegistryTransactor: OrganizationRegistryTransactor{contract: contract}, OrganizationRegistryFilterer: OrganizationRegistryFilterer{contract: contract}}, nil
}

// OrganizationRegistry is an auto generated Go binding around an Ethereum contract.
type OrganizationRegistry struct {
	OrganizationRegistryCaller     // Read-only binding to the contract
	OrganizationRegistryTransactor // Write-only binding to the contract
	OrganizationRegistryFilterer   // Log filterer for contract events
}

// OrganizationRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrganizationRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganizationRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganizationRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganizationRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganizationRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganizationRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganizationRegistrySession struct {
	Contract     *OrganizationRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OrganizationRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganizationRegistryCallerSession struct {
	Contract *OrganizationRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OrganizationRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganizationRegistryTransactorSession struct {
	Contract     *OrganizationRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OrganizationRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrganizationRegistryRaw struct {
	Contract *OrganizationRegistry // Generic contract binding to access the raw methods on
}

// OrganizationRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganizationRegistryCallerRaw struct {
	Contract *OrganizationRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// OrganizationRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganizationRegistryTransactorRaw struct {
	Contract *OrganizationRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganizationRegistry creates a new instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistry(address common.Address, backend bind.ContractBackend) (*OrganizationRegistry, error) {
	contract, err := bindOrganizationRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistry{OrganizationRegistryCaller: OrganizationRegistryCaller{contract: contract}, OrganizationRegistryTransactor: OrganizationRegistryTransactor{contract: contract}, OrganizationRegistryFilterer: OrganizationRegistryFilterer{contract: contract}}, nil
}

// NewOrganizationRegistryCaller creates a new read-only instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistryCaller(address common.Address, caller bind.ContractCaller) (*OrganizationRegistryCaller, error) {
	contract, err := bindOrganizationRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryCaller{contract: contract}, nil
}

// NewOrganizationRegistryTransactor creates a new write-only instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OrganizationRegistryTransactor, error) {
	contract, err := bindOrganizationRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryTransactor{contract: contract}, nil
}

// NewOrganizationRegistryFilterer creates a new log filterer instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OrganizationRegistryFilterer, error) {
	contract, err := bindOrganizationRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryFilterer{contract: contract}, nil
}

// bindOrganizationRegistry binds a generic wrapper to an already deployed contract.
func bindOrganizationRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrganizationRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganizationRegistry *OrganizationRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganizationRegistry.Contract.OrganizationRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganizationRegistry *OrganizationRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.OrganizationRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganizationRegistry *OrganizationRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.OrganizationRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganizationRegistry *OrganizationRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganizationRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganizationRegistry *OrganizationRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganizationRegistry *OrganizationRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_OrganizationRegistry *OrganizationRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_OrganizationRegistry *OrganizationRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _OrganizationRegistry.Contract.UPGRADEINTERFACEVERSION(&_OrganizationRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _OrganizationRegistry.Contract.UPGRADEINTERFACEVERSION(&_OrganizationRegistry.CallOpts)
}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns(string, string)
func (_OrganizationRegistry *OrganizationRegistryCaller) GetOrganization(opts *bind.CallOpts, id common.Address) (string, string, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "getOrganization", id)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns(string, string)
func (_OrganizationRegistry *OrganizationRegistrySession) GetOrganization(id common.Address) (string, string, error) {
	return _OrganizationRegistry.Contract.GetOrganization(&_OrganizationRegistry.CallOpts, id)
}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns(string, string)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) GetOrganization(id common.Address) (string, string, error) {
	return _OrganizationRegistry.Contract.GetOrganization(&_OrganizationRegistry.CallOpts, id)
}

// IsAdministrator is a free data retrieval call binding the contract method 0xc1af6e03.
//
// Solidity: function isAdministrator(address id, address account) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistryCaller) IsAdministrator(opts *bind.CallOpts, id common.Address, account common.Address) (bool, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "isAdministrator", id, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdministrator is a free data retrieval call binding the contract method 0xc1af6e03.
//
// Solidity: function isAdministrator(address id, address account) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistrySession) IsAdministrator(id common.Address, account common.Address) (bool, error) {
	return _OrganizationRegistry.Contract.IsAdministrator(&_OrganizationRegistry.CallOpts, id, account)
}

// IsAdministrator is a free data retrieval call binding the contract method 0xc1af6e03.
//
// Solidity: function isAdministrator(address id, address account) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) IsAdministrator(id common.Address, account common.Address) (bool, error) {
	return _OrganizationRegistry.Contract.IsAdministrator(&_OrganizationRegistry.CallOpts, id, account)
}

// OrganizationCount is a free data retrieval call binding the contract method 0xf1c62104.
//
// Solidity: function organizationCount() view returns(uint32)
func (_OrganizationRegistry *OrganizationRegistryCaller) OrganizationCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "organizationCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OrganizationCount is a free data retrieval call binding the contract method 0xf1c62104.
//
// Solidity: function organizationCount() view returns(uint32)
func (_OrganizationRegistry *OrganizationRegistrySession) OrganizationCount() (uint32, error) {
	return _OrganizationRegistry.Contract.OrganizationCount(&_OrganizationRegistry.CallOpts)
}

// OrganizationCount is a free data retrieval call binding the contract method 0xf1c62104.
//
// Solidity: function organizationCount() view returns(uint32)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) OrganizationCount() (uint32, error) {
	return _OrganizationRegistry.Contract.OrganizationCount(&_OrganizationRegistry.CallOpts)
}

// Organizations is a free data retrieval call binding the contract method 0x5a1f7406.
//
// Solidity: function organizations(address ) view returns(string name, string metadataURI)
func (_OrganizationRegistry *OrganizationRegistryCaller) Organizations(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name        string
	MetadataURI string
}, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "organizations", arg0)

	outstruct := new(struct {
		Name        string
		MetadataURI string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.MetadataURI = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Organizations is a free data retrieval call binding the contract method 0x5a1f7406.
//
// Solidity: function organizations(address ) view returns(string name, string metadataURI)
func (_OrganizationRegistry *OrganizationRegistrySession) Organizations(arg0 common.Address) (struct {
	Name        string
	MetadataURI string
}, error) {
	return _OrganizationRegistry.Contract.Organizations(&_OrganizationRegistry.CallOpts, arg0)
}

// Organizations is a free data retrieval call binding the contract method 0x5a1f7406.
//
// Solidity: function organizations(address ) view returns(string name, string metadataURI)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) Organizations(arg0 common.Address) (struct {
	Name        string
	MetadataURI string
}, error) {
	return _OrganizationRegistry.Contract.Organizations(&_OrganizationRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganizationRegistry *OrganizationRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganizationRegistry *OrganizationRegistrySession) Owner() (common.Address, error) {
	return _OrganizationRegistry.Contract.Owner(&_OrganizationRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) Owner() (common.Address, error) {
	return _OrganizationRegistry.Contract.Owner(&_OrganizationRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OrganizationRegistry *OrganizationRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OrganizationRegistry *OrganizationRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _OrganizationRegistry.Contract.ProxiableUUID(&_OrganizationRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _OrganizationRegistry.Contract.ProxiableUUID(&_OrganizationRegistry.CallOpts)
}

// AddAdministrator is a paid mutator transaction binding the contract method 0x1c2e3d82.
//
// Solidity: function addAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) AddAdministrator(opts *bind.TransactOpts, id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "addAdministrator", id, administrator)
}

// AddAdministrator is a paid mutator transaction binding the contract method 0x1c2e3d82.
//
// Solidity: function addAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) AddAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.AddAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// AddAdministrator is a paid mutator transaction binding the contract method 0x1c2e3d82.
//
// Solidity: function addAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) AddAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.AddAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xc2a950be.
//
// Solidity: function createOrganization(address id, string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) CreateOrganization(opts *bind.TransactOpts, id common.Address, name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "createOrganization", id, name, metadataURI, administrators)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xc2a950be.
//
// Solidity: function createOrganization(address id, string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) CreateOrganization(id common.Address, name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.CreateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI, administrators)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xc2a950be.
//
// Solidity: function createOrganization(address id, string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) CreateOrganization(id common.Address, name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.CreateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI, administrators)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x7acbb8af.
//
// Solidity: function deleteOrganization(address id) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) DeleteOrganization(opts *bind.TransactOpts, id common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "deleteOrganization", id)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x7acbb8af.
//
// Solidity: function deleteOrganization(address id) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) DeleteOrganization(id common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.DeleteOrganization(&_OrganizationRegistry.TransactOpts, id)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x7acbb8af.
//
// Solidity: function deleteOrganization(address id) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) DeleteOrganization(id common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.DeleteOrganization(&_OrganizationRegistry.TransactOpts, id)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OrganizationRegistry *OrganizationRegistrySession) Initialize() (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.Initialize(&_OrganizationRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.Initialize(&_OrganizationRegistry.TransactOpts)
}

// RemoveAdministrator is a paid mutator transaction binding the contract method 0x3c10eee5.
//
// Solidity: function removeAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) RemoveAdministrator(opts *bind.TransactOpts, id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "removeAdministrator", id, administrator)
}

// RemoveAdministrator is a paid mutator transaction binding the contract method 0x3c10eee5.
//
// Solidity: function removeAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) RemoveAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.RemoveAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// RemoveAdministrator is a paid mutator transaction binding the contract method 0x3c10eee5.
//
// Solidity: function removeAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) RemoveAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.RemoveAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OrganizationRegistry *OrganizationRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.RenounceOwnership(&_OrganizationRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.RenounceOwnership(&_OrganizationRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.TransferOwnership(&_OrganizationRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.TransferOwnership(&_OrganizationRegistry.TransactOpts, newOwner)
}

// UpdateOrganization is a paid mutator transaction binding the contract method 0x6cca67bf.
//
// Solidity: function updateOrganization(address id, string name, string metadataURI) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) UpdateOrganization(opts *bind.TransactOpts, id common.Address, name string, metadataURI string) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "updateOrganization", id, name, metadataURI)
}

// UpdateOrganization is a paid mutator transaction binding the contract method 0x6cca67bf.
//
// Solidity: function updateOrganization(address id, string name, string metadataURI) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) UpdateOrganization(id common.Address, name string, metadataURI string) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.UpdateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI)
}

// UpdateOrganization is a paid mutator transaction binding the contract method 0x6cca67bf.
//
// Solidity: function updateOrganization(address id, string name, string metadataURI) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) UpdateOrganization(id common.Address, name string, metadataURI string) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.UpdateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OrganizationRegistry *OrganizationRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.UpgradeToAndCall(&_OrganizationRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.UpgradeToAndCall(&_OrganizationRegistry.TransactOpts, newImplementation, data)
}

// OrganizationRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OrganizationRegistry contract.
type OrganizationRegistryInitializedIterator struct {
	Event *OrganizationRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryInitialized)
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
		it.Event = new(OrganizationRegistryInitialized)
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
func (it *OrganizationRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryInitialized represents a Initialized event raised by the OrganizationRegistry contract.
type OrganizationRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*OrganizationRegistryInitializedIterator, error) {

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryInitializedIterator{contract: _OrganizationRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryInitialized)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseInitialized(log types.Log) (*OrganizationRegistryInitialized, error) {
	event := new(OrganizationRegistryInitialized)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganizationRegistryOrganizationCreatedIterator is returned from FilterOrganizationCreated and is used to iterate over the raw logs and unpacked data for OrganizationCreated events raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationCreatedIterator struct {
	Event *OrganizationRegistryOrganizationCreated // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryOrganizationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryOrganizationCreated)
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
		it.Event = new(OrganizationRegistryOrganizationCreated)
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
func (it *OrganizationRegistryOrganizationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryOrganizationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryOrganizationCreated represents a OrganizationCreated event raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationCreated struct {
	Id      common.Address
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrganizationCreated is a free log retrieval operation binding the contract event 0x2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a.
//
// Solidity: event OrganizationCreated(address indexed id, address indexed creator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterOrganizationCreated(opts *bind.FilterOpts, id []common.Address, creator []common.Address) (*OrganizationRegistryOrganizationCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "OrganizationCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryOrganizationCreatedIterator{contract: _OrganizationRegistry.contract, event: "OrganizationCreated", logs: logs, sub: sub}, nil
}

// WatchOrganizationCreated is a free log subscription operation binding the contract event 0x2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a.
//
// Solidity: event OrganizationCreated(address indexed id, address indexed creator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchOrganizationCreated(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryOrganizationCreated, id []common.Address, creator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "OrganizationCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryOrganizationCreated)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationCreated", log); err != nil {
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

// ParseOrganizationCreated is a log parse operation binding the contract event 0x2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a.
//
// Solidity: event OrganizationCreated(address indexed id, address indexed creator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseOrganizationCreated(log types.Log) (*OrganizationRegistryOrganizationCreated, error) {
	event := new(OrganizationRegistryOrganizationCreated)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganizationRegistryOrganizationUpdatedIterator is returned from FilterOrganizationUpdated and is used to iterate over the raw logs and unpacked data for OrganizationUpdated events raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationUpdatedIterator struct {
	Event *OrganizationRegistryOrganizationUpdated // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryOrganizationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryOrganizationUpdated)
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
		it.Event = new(OrganizationRegistryOrganizationUpdated)
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
func (it *OrganizationRegistryOrganizationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryOrganizationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryOrganizationUpdated represents a OrganizationUpdated event raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationUpdated struct {
	Id      common.Address
	Updater common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrganizationUpdated is a free log retrieval operation binding the contract event 0xdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f6.
//
// Solidity: event OrganizationUpdated(address indexed id, address indexed updater)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterOrganizationUpdated(opts *bind.FilterOpts, id []common.Address, updater []common.Address) (*OrganizationRegistryOrganizationUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "OrganizationUpdated", idRule, updaterRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryOrganizationUpdatedIterator{contract: _OrganizationRegistry.contract, event: "OrganizationUpdated", logs: logs, sub: sub}, nil
}

// WatchOrganizationUpdated is a free log subscription operation binding the contract event 0xdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f6.
//
// Solidity: event OrganizationUpdated(address indexed id, address indexed updater)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchOrganizationUpdated(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryOrganizationUpdated, id []common.Address, updater []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "OrganizationUpdated", idRule, updaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryOrganizationUpdated)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationUpdated", log); err != nil {
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

// ParseOrganizationUpdated is a log parse operation binding the contract event 0xdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f6.
//
// Solidity: event OrganizationUpdated(address indexed id, address indexed updater)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseOrganizationUpdated(log types.Log) (*OrganizationRegistryOrganizationUpdated, error) {
	event := new(OrganizationRegistryOrganizationUpdated)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganizationRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OrganizationRegistry contract.
type OrganizationRegistryOwnershipTransferredIterator struct {
	Event *OrganizationRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryOwnershipTransferred)
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
		it.Event = new(OrganizationRegistryOwnershipTransferred)
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
func (it *OrganizationRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the OrganizationRegistry contract.
type OrganizationRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OrganizationRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryOwnershipTransferredIterator{contract: _OrganizationRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryOwnershipTransferred)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*OrganizationRegistryOwnershipTransferred, error) {
	event := new(OrganizationRegistryOwnershipTransferred)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganizationRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the OrganizationRegistry contract.
type OrganizationRegistryUpgradedIterator struct {
	Event *OrganizationRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryUpgraded)
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
		it.Event = new(OrganizationRegistryUpgraded)
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
func (it *OrganizationRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryUpgraded represents a Upgraded event raised by the OrganizationRegistry contract.
type OrganizationRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OrganizationRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryUpgradedIterator{contract: _OrganizationRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryUpgraded)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseUpgraded(log types.Log) (*OrganizationRegistryUpgraded, error) {
	event := new(OrganizationRegistryUpgraded)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
