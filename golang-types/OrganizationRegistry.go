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
	Bin: "0x6080604052348015600e575f5ffd5b503380603357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b603a81603f565b50608e565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6114a68061009b5f395ff3fe608060405234801561000f575f5ffd5b50600436106100cb575f3560e01c80637acbb8af11610088578063d2c30a6d11610063578063d2c30a6d1461019f578063f1c62104146101bf578063f2fde38b146101e4578063f6a3d24e146101f7575f5ffd5b80637acbb8af1461014f5780638da5cb5b14610162578063c1af6e031461017c575f5ffd5b80631c2e3d82146100cf5780633c10eee5146100e45780635a1f7406146100f75780636cca67bf14610121578063715018a61461013457806375f1373f1461013c575b5f5ffd5b6100e26100dd366004610fbb565b61020a565b005b6100e26100f2366004610fbb565b610378565b61010a610105366004610fec565b61058e565b60405161011892919061103a565b60405180910390f35b6100e261012f3660046110ac565b6106b5565b6100e26107cc565b6100e261014a36600461112c565b6107df565b6100e261015d366004610fec565b610a1a565b5f546040516001600160a01b039091168152602001610118565b61018f61018a366004610fbb565b610af9565b6040519015158152602001610118565b6101b26101ad366004610fec565b610c0d565b60405161011891906111fc565b6002546101cf9063ffffffff1681565b60405163ffffffff9091168152602001610118565b6100e26101f2366004610fec565b610dd9565b61018f610205366004610fec565b610e1b565b336001600160a01b03831614610232576040516282b42960e81b815260040160405180910390fd5b6001600160a01b0382165f90815260016020526040812080546102549061128d565b90501161027457604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b03811661029b5760405163e6c4247b60e01b815260040160405180910390fd5b6001600160a01b0382165f908152600160205260408120600201905b815481101561031657826001600160a01b03168282815481106102dc576102dc6112c5565b5f918252602090912001546001600160a01b03160361030e5760405163b598ce3d60e01b815260040160405180910390fd5b6001016102b7565b5080546001810182555f8281526020812090910180546001600160a01b0319166001600160a01b038581169182179092556040519092918616917fa2f792d5cab45450daed47916f25d9bd3443d275c51ecbccfbeb727a3b4eda0e91a3505050565b8161038281610e4a565b61039f5760405163e85bbb2160e01b815260040160405180910390fd5b336001600160a01b038416146103b3573391505b6001600160a01b0383165f90815260016020526040812080546103d59061128d565b9050116103f557604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b03821661041c5760405163e6c4247b60e01b815260040160405180910390fd5b6001600160a01b0383165f908152600160205260408120600201815b815481101561052e57846001600160a01b031682828154811061045d5761045d6112c5565b5f918252602090912001546001600160a01b0316036105265781546001935082906104899085906112ed565b81548110610499576104996112c5565b905f5260205f20015f9054906101000a90046001600160a01b03168282815481106104c6576104c66112c5565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508180548061050157610501611300565b5f8281526020902081015f1990810180546001600160a01b031916905501905561052e565b600101610438565b508161054d5760405163e85bbb2160e01b815260040160405180910390fd5b60405133906001600160a01b0386811691908816907f9312053150f18090e219209afedf7c42617a23f1202a19d2289891a2aa6d7fd7905f90a45050505050565b60016020525f90815260409020805481906105a89061128d565b80601f01602080910402602001604051908101604052809291908181526020018280546105d49061128d565b801561061f5780601f106105f65761010080835404028352916020019161061f565b820191905f5260205f20905b81548152906001019060200180831161060257829003601f168201915b5050505050908060010180546106349061128d565b80601f01602080910402602001604051908101604052809291908181526020018280546106609061128d565b80156106ab5780601f10610682576101008083540402835291602001916106ab565b820191905f5260205f20905b81548152906001019060200180831161068e57829003601f168201915b5050505050905082565b846106bf81610e4a565b6106dc5760405163e85bbb2160e01b815260040160405180910390fd5b836106fa57604051631b0b678d60e31b815260040160405180910390fd5b8161071857604051630eec403f60e41b815260040160405180910390fd5b6001600160a01b0386165f908152600160205260408120805461073a9061128d565b90501161075a57604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b0386165f9081526001602052604090208061077d868883611374565b506001810161078d848683611374565b5060405133906001600160a01b038916907fdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f6905f90a350505050505050565b6107d4610ebf565b6107dd5f610eeb565b565b33856107fe57604051631b0b678d60e31b815260040160405180910390fd5b6001600160a01b0381165f90815260016020526040902080546108209061128d565b159050610840576040516369c24f0b60e01b815260040160405180910390fd5b6001600160a01b0381165f90815260016020526040902080610863888a83611374565b5060018101610873868883611374565b508215610979575f5b83811015610977575f858583818110610897576108976112c5565b90506020020160208101906108ac9190610fec565b6001600160a01b0316036108d35760405163e6c4247b60e01b815260040160405180910390fd5b826001600160a01b03168585838181106108ef576108ef6112c5565b90506020020160208101906109049190610fec565b6001600160a01b03161461096f5781600201858583818110610928576109286112c5565b905060200201602081019061093d9190610fec565b81546001810183555f928352602090922090910180546001600160a01b0319166001600160a01b039092169190911790555b60010161087c565b505b600281810180546001810182555f918252602082200180546001600160a01b0319166001600160a01b038616179055815463ffffffff1691906109bb8361142e565b91906101000a81548163ffffffff021916908363ffffffff16021790555050816001600160a01b03167f47b688936cae1ca5de00ac709e05309381fb9f18b4c5adb358a5b542ce67caea60405160405180910390a25050505050505050565b6001600160a01b0381165f9081526001602052604081208054610a3c9061128d565b905011610a5c57604051639551f8b360e01b815260040160405180910390fd5b336001600160a01b03821614610a84576040516282b42960e81b815260040160405180910390fd5b6001600160a01b0381165f90815260016020526040812090610aa68282610f3a565b610ab3600183015f610f3a565b610ac0600283015f610f71565b50506002805463ffffffff16905f610ad783611452565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b6001600160a01b0382165f9081526001602052604081208054829190610b1e9061128d565b905011610b3e57604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b0383165f90815260016020908152604080832060020180548251818502810185019093528083529192909190830182828015610ba857602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610b8a575b509394505f93505050505b8151811015610c0157836001600160a01b0316828281518110610bd857610bd86112c5565b60200260200101516001600160a01b031603610bf957600192505050610c07565b600101610bb3565b505f9150505b92915050565b610c3160405180606001604052806060815260200160608152602001606081525090565b6001600160a01b0382165f9081526001602052604090819020815160608101909252805482908290610c629061128d565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8e9061128d565b8015610cd95780601f10610cb057610100808354040283529160200191610cd9565b820191905f5260205f20905b815481529060010190602001808311610cbc57829003601f168201915b50505050508152602001600182018054610cf29061128d565b80601f0160208091040260200160405190810160405280929190818152602001828054610d1e9061128d565b8015610d695780601f10610d4057610100808354040283529160200191610d69565b820191905f5260205f20905b815481529060010190602001808311610d4c57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020018280548015610dc957602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610dab575b5050505050815250509050919050565b610de1610ebf565b6001600160a01b038116610e0f57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610e1881610eeb565b50565b6001600160a01b0381165f9081526001602052604081208054829190610e409061128d565b9050119050919050565b6001600160a01b0381165f908152600160205260408120600201815b8154811015610eb657336001600160a01b0316828281548110610e8b57610e8b6112c5565b5f918252602090912001546001600160a01b031603610eae575060019392505050565b600101610e66565b505f9392505050565b5f546001600160a01b031633146107dd5760405163118cdaa760e01b8152336004820152602401610e06565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b508054610f469061128d565b5f825580601f10610f55575050565b601f0160209004905f5260205f2090810190610e189190610f88565b5080545f8255905f5260205f2090810190610e1891905b5b80821115610f9c575f8155600101610f89565b5090565b80356001600160a01b0381168114610fb6575f5ffd5b919050565b5f5f60408385031215610fcc575f5ffd5b610fd583610fa0565b9150610fe360208401610fa0565b90509250929050565b5f60208284031215610ffc575f5ffd5b61100582610fa0565b9392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b604081525f61104c604083018561100c565b828103602084015261105e818561100c565b95945050505050565b5f5f83601f840112611077575f5ffd5b50813567ffffffffffffffff81111561108e575f5ffd5b6020830191508360208285010111156110a5575f5ffd5b9250929050565b5f5f5f5f5f606086880312156110c0575f5ffd5b6110c986610fa0565b9450602086013567ffffffffffffffff8111156110e4575f5ffd5b6110f088828901611067565b909550935050604086013567ffffffffffffffff81111561110f575f5ffd5b61111b88828901611067565b969995985093965092949392505050565b5f5f5f5f5f5f60608789031215611141575f5ffd5b863567ffffffffffffffff811115611157575f5ffd5b61116389828a01611067565b909750955050602087013567ffffffffffffffff811115611182575f5ffd5b61118e89828a01611067565b909550935050604087013567ffffffffffffffff8111156111ad575f5ffd5b8701601f810189136111bd575f5ffd5b803567ffffffffffffffff8111156111d3575f5ffd5b8960208260051b84010111156111e7575f5ffd5b60208201935080925050509295509295509295565b602081525f825160606020840152611217608084018261100c565b90506020840151601f19848303016040850152611234828261100c565b6040860151858203601f190160608701528051808352602091820194505f93509101905b808310156112835783516001600160a01b031682526020938401936001939093019290910190611258565b5095945050505050565b600181811c908216806112a157607f821691505b6020821081036112bf57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b81810381811115610c0757610c076112d9565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b601f82111561136f57805f5260205f20601f840160051c8101602085101561134d5750805b601f840160051c820191505b8181101561136c575f8155600101611359565b50505b505050565b67ffffffffffffffff83111561138c5761138c611314565b6113a08361139a835461128d565b83611328565b5f601f8411600181146113d1575f85156113ba5750838201355b5f19600387901b1c1916600186901b17835561136c565b5f83815260208120601f198716915b8281101561140057868501358255602094850194600190920191016113e0565b508682101561141c575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f63ffffffff821663ffffffff8103611449576114496112d9565b60010192915050565b5f63ffffffff821680611467576114676112d9565b5f19019291505056fea26469706673582212205096a929c68aed395836bb7130f4070b341574015cf909f0d08c490c59ce749d64736f6c634300081c0033",
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
