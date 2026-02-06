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

// IOrganizationRegistryOrganization is an auto generated low-level Go binding around an user-defined struct.
type IOrganizationRegistryOrganization struct {
	Name           string
	MetadataURI    string
	Administrators []common.Address
}

// OrganizationRegistryMetaData contains all meta data concerning the OrganizationRegistry contract.
var OrganizationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadataURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrganizationID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrganizationName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrganizationAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrganizationNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"AdministratorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remover\",\"type\":\"address\"}],\"name\":\"AdministratorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"OrganizationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"}],\"name\":\"OrganizationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"addAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"administrators\",\"type\":\"address[]\"}],\"name\":\"createOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"deleteOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"getOrganization\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"administrators\",\"type\":\"address[]\"}],\"internalType\":\"structIOrganizationRegistry.Organization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdministrator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"organizations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"removeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346071573315605e575f8054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a36111af90816100768239f35b631e4fbdf760e01b5f525f60045260245ffd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631c2e3d8214610cd85780633c10eee514610b485780635a1f740614610ada5780636cca67bf14610810578063715018a6146107b957806375f1373f146103e35780637acbb8af146102fb5780638da5cb5b146102d4578063c1af6e031461029e578063d2c30a6d14610196578063f1c6210414610173578063f2fde38b146100ee5763f6a3d24e146100a8575f80fd5b346100ea5760203660031901126100ea576001600160a01b036100c9610dbc565b165f52600160205260206100e060405f2054610de8565b1515604051908152f35b5f80fd5b346100ea5760203660031901126100ea57610107610dbc565b61010f611153565b6001600160a01b03168015610160575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b346100ea575f3660031901126100ea57602063ffffffff60025416604051908152f35b346100ea5760203660031901126100ea576101af610dbc565b6060604080516101be81610e20565b828152826020820152015260018060a01b03165f52600160205261023d60405f20604051906101ec82610e20565b6101f581610e5c565b8252610250610219600261020b60018501610e5c565b936020860194855201611026565b91604084019283526040519485946020865251606060208701526080860190610efc565b9051848203601f19016040860152610efc565b905190601f19838203016060840152602080835192838152019201905f5b81811061027c575050500390f35b82516001600160a01b031684528594506020938401939092019160010161026e565b346100ea5760403660031901126100ea5760206102ca6102bc610dbc565b6102c4610dd2565b90611077565b6040519015158152f35b346100ea575f3660031901126100ea575f546040516001600160a01b039091168152602090f35b346100ea5760203660031901126100ea576001600160a01b0361031c610dbc565b16805f52600160205261033260405f2054610de8565b156103d4578033036103c6575f526001602052600260405f2061035481610fd8565b61036060018201610fd8565b0180545f8255806103ac575b505060025463ffffffff811680156103985763ffffffff199091165f1990910163ffffffff1617600255005b634e487b7160e01b5f52601160045260245ffd5b6103bf915f5260205f2090810190610f9e565b808061036c565b6282b42960e81b5f5260045ffd5b639551f8b360e01b5f5260045ffd5b346100ea5760603660031901126100ea576004356001600160401b0381116100ea57610413903690600401610f20565b906024356001600160401b0381116100ea57610433903690600401610f20565b90604435916001600160401b0383116100ea57366023840112156100ea578260040135946001600160401b0386116100ea576024840193602436918860051b0101116100ea5780156107aa57335f52600160205261049460405f2054610de8565b61079b57335f52600160205260405f20946001600160401b0382116106e8576104bd8654610de8565b601f811161076b575b505f90601f8311600114610707576104f592915f91836106fc575b50508160011b915f199060031b1c19161790565b84555b60018401916001600160401b0382116106e8576105158354610de8565b601f81116106ad575b505f90601f83116001146106495761054c92915f918361063e5750508160011b915f199060031b1c19161790565b90555b826105b7575b5061056591506002339101610f76565b60025463ffffffff811663ffffffff811461039857600163ffffffff9101169063ffffffff191617600255337f47b688936cae1ca5de00ac709e05309381fb9f18b4c5adb358a5b542ce67caea5f80a2005b60028201905f5b8481106105cc575050610555565b6001600160a01b036105e76105e2838886610fb4565b610fc4565b161561062f57600190336001600160a01b036106076105e2848a88610fb4565b161461062a5761062461061e6105e2838987610fb4565b85610f76565b016105be565b610624565b63e6c4247b60e01b5f5260045ffd5b0135905087806104e1565b601f19831691845f5260205f20925f5b818110610695575090846001959493921061067c575b505050811b01905561054f565b01355f19600384901b60f8161c1916905586808061066f565b91936020600181928787013581550195019201610659565b6106d890845f5260205f20601f850160051c810191602086106106de575b601f0160051c0190610f9e565b8661051e565b90915081906106cb565b634e487b7160e01b5f52604160045260245ffd5b0135905088806104e1565b601f19831691875f5260205f20925f5b818110610753575090846001959493921061073a575b505050811b0184556104f8565b01355f19600384901b60f8161c1916905587808061072d565b91936020600181928787013581550195019201610717565b61079590875f5260205f20601f850160051c810191602086106106de57601f0160051c0190610f9e565b876104c6565b6369c24f0b60e01b5f5260045ffd5b631b0b678d60e31b5f5260045ffd5b346100ea575f3660031901126100ea576107d1611153565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100ea5760603660031901126100ea57610829610dbc565b6024356001600160401b0381116100ea57610848903690600401610f20565b906044356001600160401b0381116100ea57610868903690600401610f20565b919093610874816110f6565b15610acb5783156107aa578215610abc5760018060a01b031692835f5260016020526108a360405f2054610de8565b156103d457835f52600160205260405f20916001600160401b0382116106e8576108cd8354610de8565b601f8111610a8c575b505f90601f8311600114610a2557918061090892600195945f926106fc5750508160011b915f199060031b1c19161790565b81555b01906001600160401b0381116106e8576109258254610de8565b601f81116109f5575b505f601f821160011461099357819061095d9394955f926109885750508160011b915f199060031b1c19161790565b90555b33907fdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f65f80a3005b0135905085806104e1565b601f19821694835f5260205f20915f5b8781106109dd5750836001959697106109c4575b505050811b019055610960565b01355f19600384901b60f8161c191690558480806109b7565b909260206001819286860135815501940191016109a3565b610a1f90835f5260205f20601f840160051c810191602085106106de57601f0160051c0190610f9e565b8461092e565b601f19831691845f5260205f20925f5b818110610a745750916001959492918387959310610a5b575b505050811b01815561090b565b01355f19600384901b60f8161c19169055878080610a4e565b91936020600181928787013581550195019201610a35565b610ab690845f5260205f20601f850160051c810191602086106106de57601f0160051c0190610f9e565b866108d6565b630eec403f60e41b5f5260045ffd5b63e85bbb2160e01b5f5260045ffd5b346100ea5760203660031901126100ea576001600160a01b03610afb610dbc565b165f526001602052610b3660405f20610b44610b226001610b1b84610e5c565b9301610e5c565b604051938493604085526040850190610efc565b908382036020850152610efc565b0390f35b346100ea5760403660031901126100ea57610b61610dbc565b610b69610dd2565b610b72826110f6565b15610acb576001600160a01b039091169033829003610cd1575b815f526001602052610ba160405f2054610de8565b156103d4576001600160a01b0316801561062f575f90825f526001602052600260405f20015f5b815480821015610cc85783610bdd8385610f4d565b905460039190911b1c6001600160a01b031614610bfd5750600101610bc8565b9293506001925f1981019190821161039857610c38610c1f610c5c9385610f4d565b905460039190911b1c6001600160a01b03169184610f4d565b81546001600160a01b0393841660039290921b91821b9390911b1916919091179055565b80548015610cb4575f190190610c728282610f4d565b81549060018060a01b039060031b1b19169055555b15610acb5733917f9312053150f18090e219209afedf7c42617a23f1202a19d2289891a2aa6d7fd75f80a4005b634e487b7160e01b5f52603160045260245ffd5b50505090610c87565b5033610b8c565b346100ea5760403660031901126100ea57610cf1610dbc565b610cf9610dd2565b906001600160a01b0316338190036103c657805f526001602052610d2060405f2054610de8565b156103d4576001600160a01b03821690811561062f575f818152600160205260408120600201805494915b858110610d83575090610d5d91610f76565b7fa2f792d5cab45450daed47916f25d9bd3443d275c51ecbccfbeb727a3b4eda0e5f80a3005b84610d8e8284610f4d565b905460039190911b1c6001600160a01b031614610dad57600101610d4b565b63b598ce3d60e01b5f5260045ffd5b600435906001600160a01b03821682036100ea57565b602435906001600160a01b03821682036100ea57565b90600182811c92168015610e16575b6020831014610e0257565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610df7565b606081019081106001600160401b038211176106e857604052565b90601f801991011681019081106001600160401b038211176106e857604052565b9060405191825f825492610e6f84610de8565b8084529360018116908115610eda5750600114610e96575b50610e9492500383610e3b565b565b90505f9291925260205f20905f915b818310610ebe575050906020610e94928201015f610e87565b6020919350806001915483858901015201910190918492610ea5565b905060209250610e9494915060ff191682840152151560051b8201015f610e87565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9181601f840112156100ea578235916001600160401b0383116100ea57602083818601950101116100ea57565b8054821015610f62575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b90815491680100000000000000008310156106e85782610c38916001610e9495018155610f4d565b818110610fa9575050565b5f8155600101610f9e565b9190811015610f625760051b0190565b356001600160a01b03811681036100ea5790565b610fe28154610de8565b9081610fec575050565b81601f5f9311600114610ffd575055565b8183526020832061101991601f0160051c810190600101610f9e565b8082528160208120915555565b90604051918281549182825260208201905f5260205f20925f5b818110611055575050610e9492500383610e3b565b84546001600160a01b0316835260019485019487945060209093019201611040565b6001600160a01b03165f8181526001602052604090205461109790610de8565b156103d4575f5260016020526110b2600260405f2001611026565b805191905f5b838110156110ee57600581901b8201602001516001600160a01b038481169116146110e5576001016110b8565b50505050600190565b505050505f90565b6001600160a01b03165f9081526001602052604081206002018054915b828110611121575050505f90565b61112b8183610f4d565b905460039190911b1c6001600160a01b0316331461114b57600101611113565b505050600190565b5f546001600160a01b0316330361116657565b63118cdaa760e01b5f523360045260245ffdfea26469706673582212200043fd0a194dd64a7d90093a62fe00dcac86265643b331bbf512e5dea9b8dc6464736f6c634300081c0033",
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

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address id) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistryCaller) Exists(opts *bind.CallOpts, id common.Address) (bool, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address id) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistrySession) Exists(id common.Address) (bool, error) {
	return _OrganizationRegistry.Contract.Exists(&_OrganizationRegistry.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address id) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) Exists(id common.Address) (bool, error) {
	return _OrganizationRegistry.Contract.Exists(&_OrganizationRegistry.CallOpts, id)
}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns((string,string,address[]))
func (_OrganizationRegistry *OrganizationRegistryCaller) GetOrganization(opts *bind.CallOpts, id common.Address) (IOrganizationRegistryOrganization, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "getOrganization", id)

	if err != nil {
		return *new(IOrganizationRegistryOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new(IOrganizationRegistryOrganization)).(*IOrganizationRegistryOrganization)

	return out0, err

}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns((string,string,address[]))
func (_OrganizationRegistry *OrganizationRegistrySession) GetOrganization(id common.Address) (IOrganizationRegistryOrganization, error) {
	return _OrganizationRegistry.Contract.GetOrganization(&_OrganizationRegistry.CallOpts, id)
}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns((string,string,address[]))
func (_OrganizationRegistry *OrganizationRegistryCallerSession) GetOrganization(id common.Address) (IOrganizationRegistryOrganization, error) {
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

// CreateOrganization is a paid mutator transaction binding the contract method 0x75f1373f.
//
// Solidity: function createOrganization(string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) CreateOrganization(opts *bind.TransactOpts, name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "createOrganization", name, metadataURI, administrators)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0x75f1373f.
//
// Solidity: function createOrganization(string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) CreateOrganization(name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.CreateOrganization(&_OrganizationRegistry.TransactOpts, name, metadataURI, administrators)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0x75f1373f.
//
// Solidity: function createOrganization(string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) CreateOrganization(name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.CreateOrganization(&_OrganizationRegistry.TransactOpts, name, metadataURI, administrators)
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

// OrganizationRegistryAdministratorAddedIterator is returned from FilterAdministratorAdded and is used to iterate over the raw logs and unpacked data for AdministratorAdded events raised by the OrganizationRegistry contract.
type OrganizationRegistryAdministratorAddedIterator struct {
	Event *OrganizationRegistryAdministratorAdded // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryAdministratorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryAdministratorAdded)
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
		it.Event = new(OrganizationRegistryAdministratorAdded)
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
func (it *OrganizationRegistryAdministratorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryAdministratorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryAdministratorAdded represents a AdministratorAdded event raised by the OrganizationRegistry contract.
type OrganizationRegistryAdministratorAdded struct {
	Id            common.Address
	Administrator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdministratorAdded is a free log retrieval operation binding the contract event 0xa2f792d5cab45450daed47916f25d9bd3443d275c51ecbccfbeb727a3b4eda0e.
//
// Solidity: event AdministratorAdded(address indexed id, address indexed administrator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterAdministratorAdded(opts *bind.FilterOpts, id []common.Address, administrator []common.Address) (*OrganizationRegistryAdministratorAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "AdministratorAdded", idRule, administratorRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryAdministratorAddedIterator{contract: _OrganizationRegistry.contract, event: "AdministratorAdded", logs: logs, sub: sub}, nil
}

// WatchAdministratorAdded is a free log subscription operation binding the contract event 0xa2f792d5cab45450daed47916f25d9bd3443d275c51ecbccfbeb727a3b4eda0e.
//
// Solidity: event AdministratorAdded(address indexed id, address indexed administrator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchAdministratorAdded(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryAdministratorAdded, id []common.Address, administrator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "AdministratorAdded", idRule, administratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryAdministratorAdded)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "AdministratorAdded", log); err != nil {
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

// ParseAdministratorAdded is a log parse operation binding the contract event 0xa2f792d5cab45450daed47916f25d9bd3443d275c51ecbccfbeb727a3b4eda0e.
//
// Solidity: event AdministratorAdded(address indexed id, address indexed administrator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseAdministratorAdded(log types.Log) (*OrganizationRegistryAdministratorAdded, error) {
	event := new(OrganizationRegistryAdministratorAdded)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "AdministratorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganizationRegistryAdministratorRemovedIterator is returned from FilterAdministratorRemoved and is used to iterate over the raw logs and unpacked data for AdministratorRemoved events raised by the OrganizationRegistry contract.
type OrganizationRegistryAdministratorRemovedIterator struct {
	Event *OrganizationRegistryAdministratorRemoved // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryAdministratorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryAdministratorRemoved)
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
		it.Event = new(OrganizationRegistryAdministratorRemoved)
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
func (it *OrganizationRegistryAdministratorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryAdministratorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryAdministratorRemoved represents a AdministratorRemoved event raised by the OrganizationRegistry contract.
type OrganizationRegistryAdministratorRemoved struct {
	Id            common.Address
	Administrator common.Address
	Remover       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdministratorRemoved is a free log retrieval operation binding the contract event 0x9312053150f18090e219209afedf7c42617a23f1202a19d2289891a2aa6d7fd7.
//
// Solidity: event AdministratorRemoved(address indexed id, address indexed administrator, address indexed remover)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterAdministratorRemoved(opts *bind.FilterOpts, id []common.Address, administrator []common.Address, remover []common.Address) (*OrganizationRegistryAdministratorRemovedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}
	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "AdministratorRemoved", idRule, administratorRule, removerRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryAdministratorRemovedIterator{contract: _OrganizationRegistry.contract, event: "AdministratorRemoved", logs: logs, sub: sub}, nil
}

// WatchAdministratorRemoved is a free log subscription operation binding the contract event 0x9312053150f18090e219209afedf7c42617a23f1202a19d2289891a2aa6d7fd7.
//
// Solidity: event AdministratorRemoved(address indexed id, address indexed administrator, address indexed remover)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchAdministratorRemoved(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryAdministratorRemoved, id []common.Address, administrator []common.Address, remover []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}
	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "AdministratorRemoved", idRule, administratorRule, removerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryAdministratorRemoved)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "AdministratorRemoved", log); err != nil {
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

// ParseAdministratorRemoved is a log parse operation binding the contract event 0x9312053150f18090e219209afedf7c42617a23f1202a19d2289891a2aa6d7fd7.
//
// Solidity: event AdministratorRemoved(address indexed id, address indexed administrator, address indexed remover)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseAdministratorRemoved(log types.Log) (*OrganizationRegistryAdministratorRemoved, error) {
	event := new(OrganizationRegistryAdministratorRemoved)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "AdministratorRemoved", log); err != nil {
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
	Id  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOrganizationCreated is a free log retrieval operation binding the contract event 0x47b688936cae1ca5de00ac709e05309381fb9f18b4c5adb358a5b542ce67caea.
//
// Solidity: event OrganizationCreated(address indexed id)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterOrganizationCreated(opts *bind.FilterOpts, id []common.Address) (*OrganizationRegistryOrganizationCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "OrganizationCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryOrganizationCreatedIterator{contract: _OrganizationRegistry.contract, event: "OrganizationCreated", logs: logs, sub: sub}, nil
}

// WatchOrganizationCreated is a free log subscription operation binding the contract event 0x47b688936cae1ca5de00ac709e05309381fb9f18b4c5adb358a5b542ce67caea.
//
// Solidity: event OrganizationCreated(address indexed id)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchOrganizationCreated(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryOrganizationCreated, id []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "OrganizationCreated", idRule)
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

// ParseOrganizationCreated is a log parse operation binding the contract event 0x47b688936cae1ca5de00ac709e05309381fb9f18b4c5adb358a5b542ce67caea.
//
// Solidity: event OrganizationCreated(address indexed id)
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
