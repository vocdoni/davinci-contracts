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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyAdministrator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadataURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrganizationID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrganizationName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrganizationAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrganizationNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"AdministratorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remover\",\"type\":\"address\"}],\"name\":\"AdministratorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"OrganizationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"}],\"name\":\"OrganizationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"addAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"administrators\",\"type\":\"address[]\"}],\"name\":\"createOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"deleteOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"getOrganization\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"administrators\",\"type\":\"address[]\"}],\"internalType\":\"structIOrganizationRegistry.Organization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdministrator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"organizations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"removeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015601357600080fd5b50608051611d6161003d600039600081816111d0015281816111f9015261133f0152611d616000f3fe6080604052600436106100fe5760003560e01c80637acbb8af11610095578063c1af6e0311610064578063c1af6e03146102bd578063d2c30a6d146102ed578063f1c621041461031a578063f2fde38b1461034c578063f6a3d24e1461036c57600080fd5b80637acbb8af146102035780638129fc1c146102235780638da5cb5b14610238578063ad3cb1cc1461027f57600080fd5b80635a1f7406116100d15780635a1f7406146101805780636cca67bf146101ae578063715018a6146101ce57806375f1373f146101e357600080fd5b80631c2e3d82146101035780633c10eee5146101255780634f1ef2861461014557806352d1902d14610158575b600080fd5b34801561010f57600080fd5b5061012361011e366004611714565b61038c565b005b34801561013157600080fd5b50610123610140366004611714565b6104fe565b61012361015336600461175d565b610720565b34801561016457600080fd5b5061016d61073f565b6040519081526020015b60405180910390f35b34801561018c57600080fd5b506101a061019b366004611824565b61075c565b60405161017792919061188f565b3480156101ba57600080fd5b506101236101c93660046118fc565b610888565b3480156101da57600080fd5b506101236109a2565b3480156101ef57600080fd5b506101236101fe366004611980565b6109b6565b34801561020f57600080fd5b5061012361021e366004611824565b610bf8565b34801561022f57600080fd5b50610123610cdc565b34801561024457600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b039091168152602001610177565b34801561028b57600080fd5b506102b0604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516101779190611a56565b3480156102c957600080fd5b506102dd6102d8366004611714565b610df1565b6040519015158152602001610177565b3480156102f957600080fd5b5061030d610308366004611824565b610f07565b6040516101779190611a69565b34801561032657600080fd5b506001546103379063ffffffff1681565b60405163ffffffff9091168152602001610177565b34801561035857600080fd5b50610123610367366004611824565b6110da565b34801561037857600080fd5b506102dd610387366004611824565b61111d565b336001600160a01b038316146103b4576040516282b42960e81b815260040160405180910390fd5b6001600160a01b038216600090815260208190526040812080546103d790611afc565b9050116103f757604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b03811661041e5760405163e6c4247b60e01b815260040160405180910390fd5b6001600160a01b0382166000908152602081905260408120600201905b815481101561049b57826001600160a01b031682828154811061046057610460611b36565b6000918252602090912001546001600160a01b0316036104935760405163b598ce3d60e01b815260040160405180910390fd5b60010161043b565b50805460018101825560008281526020812090910180546001600160a01b0319166001600160a01b038581169182179092556040519092918616917fa2f792d5cab45450daed47916f25d9bd3443d275c51ecbccfbeb727a3b4eda0e91a3505050565b816105088161114d565b6105255760405163e85bbb2160e01b815260040160405180910390fd5b336001600160a01b03841614610539573391505b6001600160a01b0383166000908152602081905260408120805461055c90611afc565b90501161057c57604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b0382166105a35760405163e6c4247b60e01b815260040160405180910390fd5b6001600160a01b0383166000908152602081905260408120600201815b81548110156106bf57846001600160a01b03168282815481106105e5576105e5611b36565b6000918252602090912001546001600160a01b0316036106b7578154600193508290610612908590611b62565b8154811061062257610622611b36565b9060005260206000200160009054906101000a90046001600160a01b031682828154811061065257610652611b36565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508180548061069057610690611b75565b600082815260209020810160001990810180546001600160a01b03191690550190556106bf565b6001016105c0565b50816106de5760405163e85bbb2160e01b815260040160405180910390fd5b60405133906001600160a01b0386811691908816907f9312053150f18090e219209afedf7c42617a23f1202a19d2289891a2aa6d7fd790600090a45050505050565b6107286111c5565b6107318261126a565b61073b8282611272565b5050565b6000610749611334565b50600080516020611d0c83398151915290565b60006020819052908152604090208054819061077790611afc565b80601f01602080910402602001604051908101604052809291908181526020018280546107a390611afc565b80156107f05780601f106107c5576101008083540402835291602001916107f0565b820191906000526020600020905b8154815290600101906020018083116107d357829003601f168201915b50505050509080600101805461080590611afc565b80601f016020809104026020016040519081016040528092919081815260200182805461083190611afc565b801561087e5780601f106108535761010080835404028352916020019161087e565b820191906000526020600020905b81548152906001019060200180831161086157829003601f168201915b5050505050905082565b846108928161114d565b6108af5760405163e85bbb2160e01b815260040160405180910390fd5b836108cd57604051631b0b678d60e31b815260040160405180910390fd5b816108eb57604051630eec403f60e41b815260040160405180910390fd5b6001600160a01b0386166000908152602081905260408120805461090e90611afc565b90501161092e57604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b038616600090815260208190526040902080610952868883611bd2565b5060018101610962848683611bd2565b5060405133906001600160a01b038916907fdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f690600090a350505050505050565b6109aa61137d565b6109b460006113d8565b565b33856109d557604051631b0b678d60e31b815260040160405180910390fd5b6001600160a01b038116600090815260208190526040902080546109f890611afc565b159050610a18576040516369c24f0b60e01b815260040160405180910390fd5b6001600160a01b038116600090815260208190526040902080610a3c888a83611bd2565b5060018101610a4c868883611bd2565b508215610b555760005b83811015610b53576000858583818110610a7257610a72611b36565b9050602002016020810190610a879190611824565b6001600160a01b031603610aae5760405163e6c4247b60e01b815260040160405180910390fd5b826001600160a01b0316858583818110610aca57610aca611b36565b9050602002016020810190610adf9190611824565b6001600160a01b031614610b4b5781600201858583818110610b0357610b03611b36565b9050602002016020810190610b189190611824565b81546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b039092169190911790555b600101610a56565b505b6002810180546001808201835560009283526020832090910180546001600160a01b0319166001600160a01b038616179055805463ffffffff1691610b9983611c91565b91906101000a81548163ffffffff021916908363ffffffff16021790555050816001600160a01b03167f47b688936cae1ca5de00ac709e05309381fb9f18b4c5adb358a5b542ce67caea60405160405180910390a25050505050505050565b6001600160a01b03811660009081526020819052604081208054610c1b90611afc565b905011610c3b57604051639551f8b360e01b815260040160405180910390fd5b336001600160a01b03821614610c63576040516282b42960e81b815260040160405180910390fd5b6001600160a01b038116600090815260208190526040812090610c86828261168b565b610c9460018301600061168b565b610ca26002830160006116c5565b50506001805463ffffffff16906000610cba83611cb6565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b0316600081158015610d215750825b90506000826001600160401b03166001148015610d3d5750303b155b905081158015610d4b575080155b15610d695760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610d9357845460ff60401b1916600160401b1785555b610d9c33611449565b610da461145a565b8315610dea57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050565b6001600160a01b03821660009081526020819052604081208054829190610e1790611afc565b905011610e3757604051639551f8b360e01b815260040160405180910390fd5b6001600160a01b03831660009081526020818152604080832060020180548251818502810185019093528083529192909190830182828015610ea257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e84575b5050505050905060005b8151811015610efa57836001600160a01b0316828281518110610ed157610ed1611b36565b60200260200101516001600160a01b031603610ef257600192505050610f01565b600101610eac565b5060009150505b92915050565b610f2b60405180606001604052806060815260200160608152602001606081525090565b6001600160a01b03821660009081526020819052604090819020815160608101909252805482908290610f5d90611afc565b80601f0160208091040260200160405190810160405280929190818152602001828054610f8990611afc565b8015610fd65780601f10610fab57610100808354040283529160200191610fd6565b820191906000526020600020905b815481529060010190602001808311610fb957829003601f168201915b50505050508152602001600182018054610fef90611afc565b80601f016020809104026020016040519081016040528092919081815260200182805461101b90611afc565b80156110685780601f1061103d57610100808354040283529160200191611068565b820191906000526020600020905b81548152906001019060200180831161104b57829003601f168201915b50505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156110ca57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116110ac575b5050505050815250509050919050565b6110e261137d565b6001600160a01b03811661111157604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61111a816113d8565b50565b6001600160a01b0381166000908152602081905260408120805482919061114390611afc565b9050119050919050565b6001600160a01b0381166000908152602081905260408120600201815b81548110156111bb57336001600160a01b031682828154811061118f5761118f611b36565b6000918252602090912001546001600160a01b0316036111b3575060019392505050565b60010161116a565b5060009392505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061124c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611240600080516020611d0c833981519152546001600160a01b031690565b6001600160a01b031614155b156109b45760405163703e46dd60e11b815260040160405180910390fd5b61111a61137d565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156112cc575060408051601f3d908101601f191682019092526112c991810190611cd6565b60015b6112f457604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611108565b600080516020611d0c833981519152811461132557604051632a87526960e21b815260048101829052602401611108565b61132f8383611462565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109b45760405163703e46dd60e11b815260040160405180910390fd5b336113af7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146109b45760405163118cdaa760e01b8152336004820152602401611108565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6114516114b8565b61111a81611501565b6109b46114b8565b61146b82611509565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156114b05761132f828261156e565b61073b6115e4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166109b457604051631afcd79f60e31b815260040160405180910390fd5b6110e26114b8565b806001600160a01b03163b60000361153f57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611108565b600080516020611d0c83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161158b9190611cef565b600060405180830381855af49150503d80600081146115c6576040519150601f19603f3d011682016040523d82523d6000602084013e6115cb565b606091505b50915091506115db858383611603565b95945050505050565b34156109b45760405163b398979f60e01b815260040160405180910390fd5b6060826116185761161382611662565b61165b565b815115801561162f57506001600160a01b0384163b155b1561165857604051639996b31560e01b81526001600160a01b0385166004820152602401611108565b50805b9392505050565b8051156116725780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b50805461169790611afc565b6000825580601f106116a7575050565b601f01602090049060005260206000209081019061111a91906116df565b508054600082559060005260206000209081019061111a91905b5b808211156116f457600081556001016116e0565b5090565b80356001600160a01b038116811461170f57600080fd5b919050565b6000806040838503121561172757600080fd5b611730836116f8565b915061173e602084016116f8565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561177057600080fd5b611779836116f8565b915060208301356001600160401b0381111561179457600080fd5b8301601f810185136117a557600080fd5b80356001600160401b038111156117be576117be611747565b604051601f8201601f19908116603f011681016001600160401b03811182821017156117ec576117ec611747565b60405281815282820160200187101561180457600080fd5b816020840160208301376000602083830101528093505050509250929050565b60006020828403121561183657600080fd5b61165b826116f8565b60005b8381101561185a578181015183820152602001611842565b50506000910152565b6000815180845261187b81602086016020860161183f565b601f01601f19169290920160200192915050565b6040815260006118a26040830185611863565b82810360208401526115db8185611863565b60008083601f8401126118c657600080fd5b5081356001600160401b038111156118dd57600080fd5b6020830191508360208285010111156118f557600080fd5b9250929050565b60008060008060006060868803121561191457600080fd5b61191d866116f8565b945060208601356001600160401b0381111561193857600080fd5b611944888289016118b4565b90955093505060408601356001600160401b0381111561196357600080fd5b61196f888289016118b4565b969995985093965092949392505050565b6000806000806000806060878903121561199957600080fd5b86356001600160401b038111156119af57600080fd5b6119bb89828a016118b4565b90975095505060208701356001600160401b038111156119da57600080fd5b6119e689828a016118b4565b90955093505060408701356001600160401b03811115611a0557600080fd5b8701601f81018913611a1657600080fd5b80356001600160401b03811115611a2c57600080fd5b8960208260051b8401011115611a4157600080fd5b60208201935080925050509295509295509295565b60208152600061165b6020830184611863565b602081526000825160606020840152611a856080840182611863565b90506020840151601f19848303016040850152611aa28282611863565b6040860151858203601f19016060870152805180835260209182019450600093509101905b80831015611af25783516001600160a01b031682526020938401936001939093019290910190611ac7565b5095945050505050565b600181811c90821680611b1057607f821691505b602082108103611b3057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115610f0157610f01611b4c565b634e487b7160e01b600052603160045260246000fd5b601f82111561132f57806000526020600020601f840160051c81016020851015611bb25750805b601f840160051c820191505b81811015610dea5760008155600101611bbe565b6001600160401b03831115611be957611be9611747565b611bfd83611bf78354611afc565b83611b8b565b6000601f841160018114611c315760008515611c195750838201355b600019600387901b1c1916600186901b178355610dea565b600083815260209020601f19861690835b82811015611c625786850135825560209485019460019092019101611c42565b5086821015611c7f5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b600063ffffffff821663ffffffff8103611cad57611cad611b4c565b60010192915050565b600063ffffffff821680611ccc57611ccc611b4c565b6000190192915050565b600060208284031215611ce857600080fd5b5051919050565b60008251611d0181846020870161183f565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220b175621efcb1497baaf3e1e2b2fc2ab2688212f36574001e291ab42b9125235864736f6c634300081c0033",
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
