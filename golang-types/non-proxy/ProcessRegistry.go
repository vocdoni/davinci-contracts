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

// IProcessRegistryBallotMode is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryBallotMode struct {
	CostFromWeight  bool
	ForceUniqueness bool
	MaxCount        uint8
	CostExponent    uint8
	MaxValue        *big.Int
	MinValue        *big.Int
	MaxTotalCost    *big.Int
	MinTotalCost    *big.Int
}

// IProcessRegistryCensus is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryCensus struct {
	CensusOrigin uint8
	MaxVotes     *big.Int
	CensusRoot   [32]byte
	CensusURI    string
}

// IProcessRegistryEncryptionKey is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryEncryptionKey struct {
	X *big.Int
	Y *big.Int
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chainID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_organizationRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidCensus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrganizationAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrganizationNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"ProcessStateRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CENSUS_ORIGIN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"}],\"name\":\"getProcess\",\"outputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOverwriteCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"initStateRoot\",\"type\":\"uint256\"}],\"name\":\"newProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOverwriteCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"results\",\"type\":\"uint256[]\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161217738038061217783398101604081905261002f916100a0565b600261003b8482610216565b50600180546001600160a01b039384166001600160a01b03199182161790915560038054929093169116179055506102d4565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b038116811461009b57600080fd5b919050565b6000806000606084860312156100b557600080fd5b83516001600160401b038111156100cb57600080fd5b8401601f810186136100dc57600080fd5b80516001600160401b038111156100f5576100f561006e565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101235761012361006e565b60405281815282820160200188101561013b57600080fd5b60005b8281101561015a5760208185018101518383018201520161013e565b5060006020838301015280955050505061017660208501610084565b915061018460408501610084565b90509250925092565b600181811c908216806101a157607f821691505b6020821081036101c157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561021157806000526020600020601f840160051c810160208510156101ee5750805b601f840160051c820191505b8181101561020e57600081556001016101fa565b50505b505050565b81516001600160401b0381111561022f5761022f61006e565b6102438161023d845461018d565b846101c7565b6020601f821160018114610277576000831561025f5750848201515b600019600385901b1c1916600184901b17845561020e565b600084815260208120601f198516915b828110156102a75787850151825560209485019460019092019101610287565b50848210156102c55786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b611e94806102e36000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806392bc26691161008c578063adc879e911610066578063adc879e914610210578063b79938e914610225578063c718c01f14610238578063d71d072e1461024b57600080fd5b806392bc2669146101bf578063992bc45b146101d2578063ac64bd03146101fd57600080fd5b806356a6f1e2116100c857806356a6f1e2146101635780636bae04ea1461017857806372c628ef1461018b578063848df5401461019357600080fd5b806302cc919a146100ef5780630535fece1461010e5780632b7ac3f314610138575b600080fd5b6100f7600981565b60405160ff90911681526020015b60405180910390f35b61012161011c366004611368565b61025e565b6040516101059b9a9998979695949392919061149e565b60035461014b906001600160a01b031681565b6040516001600160a01b039091168152602001610105565b61017661017136600461154d565b6104a0565b005b610176610186366004611591565b61066f565b6100f7600481565b6001546101aa90600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610105565b6101766101cd36600461165d565b61089c565b6101e56101e0366004611368565b610be9565b6040516101059c9b9a99989796959493929190611744565b61017661020b366004611825565b505050565b610218610fef565b60405161010591906118a6565b60015461014b906001600160a01b031681565b6101766102463660046118c0565b61107d565b6101766102593660046118e2565b6111fd565b6000602081815291815260409081902080548251808401909352600182015483526002820154938301939093526003810154600582015460068301546007840154600885015460098601805460ff8a169961010090046001600160a01b0316989791906102ca90611961565b80601f01602080910402602001604051908101604052809291908181526020018280546102f690611961565b80156103435780601f1061031857610100808354040283529160200191610343565b820191906000526020600020905b81548152906001019060200180831161032657829003601f168201915b5050604080516101008082018352600a88015460ff808216151584529181048216151560208401526201000081048216838501526301000000900481166060830152600b880154608080840191909152600c89015460a0840152600d89015460c0840152600e89015460e08401528351908101909352600f880180549798929792965092945091925083911660098111156103e0576103e0611381565b60098111156103f1576103f1611381565b8152602001600182015481526020016002820154815260200160038201805461041990611961565b80601f016020809104026020016040519081016040528092919081815260200182805461044590611961565b80156104925780601f1061046757610100808354040283529160200191610492565b820191906000526020600020905b81548152906001019060200180831161047557829003601f168201915b50505050508152505090508b565b6000828152602081905260409020805461010090046001600160a01b0316806104dc57604051634d36eb6960e01b815260040160405180910390fd5b815460ff168060048111156104f3576104f3611381565b84600481111561050557610505611381565b148061052457506004848181111561051f5761051f611381565b60ff16115b8061055e5750600081600481111561053e5761053e611381565b1415801561055e5750600381600481111561055b5761055b611381565b14155b1561057c576040516307a92f1960e51b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0384811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa1580156105cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f091906119a6565b61060d5760405163f8b857d360e01b815260040160405180910390fd5b82548490849060ff1916600183600481111561062b5761062b611381565b0217905550847fac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c1908560405161066091906119c3565b60405180910390a25050505050565b61067c60608201826119d7565b905060000361069e5760405163111127cb60e01b815260040160405180910390fd5b60408101356000036106c35760405163111127cb60e01b815260040160405180910390fd5b6000828152602081905260409020805461010090046001600160a01b0316806106ff57604051634d36eb6960e01b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0383811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa15801561074f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077391906119a6565b6107905760405163f8b857d360e01b815260040160405180910390fd5b6010820154602084013510156107b95760405163111127cb60e01b815260040160405180910390fd5b6000825460ff1660048111156107d1576107d1611381565b141580156107f557506003825460ff1660048111156107f2576107f2611381565b14155b15610813576040516307a92f1960e51b815260040160405180910390fd5b602083013560108301556040830135601183015561083460608401846119d7565b6012840191610844919083611a82565b50837f35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b604085013561087960608701876119d7565b876020013560405161088e9493929190611b6b565b60405180910390a250505050565b60015460405163c1af6e0360e01b81526001600160a01b0386811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa1580156108ec573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091091906119a6565b61092d5760405163f8b857d360e01b815260040160405180910390fd5b6000838152602081905260409020805461010090046001600160a01b03161561096957604051635040c4d360e11b815260040160405180910390fd5b61097960608a0160408b01611ba5565b60ff1660000361099c5760405163ac38930b60e01b815260040160405180910390fd5b6109ac60608a0160408b01611ba5565b60ff168960800135116109d25760405163b1911d8b60e01b815260040160405180910390fd5b60048c818111156109e5576109e5611381565b60ff161180610a23575060008c6004811115610a0357610a03611381565b14158015610a23575060038c6004811115610a2057610a20611381565b14155b15610a41576040516307a92f1960e51b815260040160405180910390fd5b42808c11610a6257604051632ca4094f60e21b815260040160405180910390fd5b80610a6d8c8e611bd8565b11610a8b57604051637616640160e01b815260040160405180910390fd5b6009610a9a60208b018b611bf8565b6009811115610aab57610aab611381565b60ff161115610acd5760405163111127cb60e01b815260040160405180910390fd5b81548d90839060ff19166001836004811115610aeb57610aeb611381565b0217905550600582018c9055600682018b90558154610100600160a81b0319166101006001600160a01b0388160217825583356001830155602084013560028301556003820183905560098201610b43888a83611a82565b5089600a8301610b538282611c2f565b50899050600f8301610b658282611cf1565b505060018054600160a01b900463ffffffff16906014610b8483611d89565b91906101000a81548163ffffffff021916908363ffffffff16021790555050336001600160a01b0316857fada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa60405160405180910390a350505050505050505050505050565b600080610c09604051806040016040528060008152602001600081525090565b600060606000806000806060610c67604051806101000160405280600015158152602001600015158152602001600060ff168152602001600060ff168152602001600081526020016000815260200160008152602001600081525090565b610c9160408051608081019091528060008152600060208201819052604082015260609081015290565b60008d815260208190526040808220815161018081019092528054829060ff166004811115610cc257610cc2611381565b6004811115610cd357610cd3611381565b8152815461010090046001600160a01b03166020808301919091526040805180820182526001850154815260028501548184015281840152600384015460608401526004840180548251818502810185019093528083526080909401939192909190830182828015610d6457602002820191906000526020600020905b815481526020019060010190808311610d50575b5050505050815260200160058201548152602001600682015481526020016007820154815260200160088201548152602001600982018054610da590611961565b80601f0160208091040260200160405190810160405280929190818152602001828054610dd190611961565b8015610e1e5780601f10610df357610100808354040283529160200191610e1e565b820191906000526020600020905b815481529060010190602001808311610e0157829003601f168201915b5050509183525050604080516101008082018352600a85015460ff8082161515845291810482161515602084810191909152620100008204831684860152630100000090910482166060840152600b860154608080850191909152600c87015460a0850152600d87015460c0850152600e87015460e08501529085019290925282519182018352600f85018054939094019391928391166009811115610ec657610ec6611381565b6009811115610ed757610ed7611381565b81526020016001820154815260200160028201548152602001600382018054610eff90611961565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2b90611961565b8015610f785780601f10610f4d57610100808354040283529160200191610f78565b820191906000526020600020905b815481529060010190602001808311610f5b57829003601f168201915b505050505081525050815250509050806000015181602001518260400151836060015184608001518560a001518660c001518760e001518861010001518961012001518a61014001518b61016001519c509c509c509c509c509c509c509c509c509c509c509c505091939597999b5091939597999b565b60028054610ffc90611961565b80601f016020809104026020016040519081016040528092919081815260200182805461102890611961565b80156110755780601f1061104a57610100808354040283529160200191611075565b820191906000526020600020905b81548152906001019060200180831161105857829003601f168201915b505050505081565b42811161109d57604051637616640160e01b815260040160405180910390fd5b6000828152602081905260409020805461010090046001600160a01b0316806110d957604051634d36eb6960e01b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0383811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa158015611129573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061114d91906119a6565b61116a5760405163f8b857d360e01b815260040160405180910390fd5b6000825460ff16600481111561118257611182611381565b141580156111a657506003825460ff1660048111156111a3576111a3611381565b14155b156111c4576040516307a92f1960e51b815260040160405180910390fd5b6006820183905560405183815284907f0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f469060200161088e565b6000858152602081905260409020805461010090046001600160a01b031661123857604051634d36eb6960e01b815260040160405180910390fd5b6000815460ff16600481111561125057611250611381565b1461126e576040516307a92f1960e51b815260040160405180910390fd5b600354604051635c73957b60e11b81526001600160a01b039091169063b8e72af6906112a4908890889088908890600401611dae565b60006040518083038186803b1580156112bc57600080fd5b505afa1580156112d0573d6000803e3d6000fd5b50600092506112e491505083850185611de0565b600383015481519192501461130c57604051630b6fac0360e41b815260040160405180910390fd5b602080820151600384018190556040808401516007860155606084015160088601555190815288917fd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20910160405180910390a250505050505050565b60006020828403121561137a57600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b600581106113a7576113a7611381565b9052565b6000815180845260005b818110156113d1576020818501810151868301820152016113b5565b506000602082860101526020601f19601f83011685010191505092915050565b80511515825260208101511515602083015260ff60408201511660408301526060810151611424606084018260ff169052565b506080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b60008151600a811061146557611465611381565b80845250602082015160208401526040820151604084015260608201516080606085015261149660808501826113ab565b949350505050565b6114a8818d611397565b6001600160a01b038b1660208201526114ce604082018b80518252602090810151910152565b8860808201528760a08201528660c08201528560e08201528461010082015261026061012082015260006115066102608301866113ab565b6115146101408401866113f1565b8281036102408401526115278185611451565b9e9d5050505050505050505050505050565b80356005811061154857600080fd5b919050565b6000806040838503121561156057600080fd5b8235915061157060208401611539565b90509250929050565b60006080828403121561158b57600080fd5b50919050565b600080604083850312156115a457600080fd5b82359150602083013567ffffffffffffffff8111156115c257600080fd5b6115ce85828601611579565b9150509250929050565b6000610100828403121561158b57600080fd5b60008083601f8401126115fd57600080fd5b50813567ffffffffffffffff81111561161557600080fd5b60208301915083602082850101111561162d57600080fd5b9250929050565b80356001600160a01b038116811461154857600080fd5b60006040828403121561158b57600080fd5b60008060008060008060008060008060006102408c8e03121561167f57600080fd5b6116888c611539565b9a5060208c0135995060408c013598506116a58d60608e016115d8565b97506101608c013567ffffffffffffffff8111156116c257600080fd5b6116ce8e828f01611579565b9750506101808c013567ffffffffffffffff8111156116ec57600080fd5b6116f88e828f016115eb565b909750955061170c90506101a08d01611634565b93506101c08c013592506117248d6101e08e0161164b565b915060006102208d01359050809150509295989b509295989b9093969950565b61174e818e611397565b6001600160a01b038c166020820152611774604082018c80518252602090810151910152565b89608082015261028060a08201526000610280820160008b5180835260208301925060208d0160005b828110156117bb57815185526020948501949091019060010161179d565b505050819050508960c08401528860e084015287610100840152866101208401528281036101408401526117ef81876113ab565b90506117ff6101608401866113f1565b8281036102608401526118128185611451565b9f9e505050505050505050505050505050565b60008060006040848603121561183a57600080fd5b83359250602084013567ffffffffffffffff81111561185857600080fd5b8401601f8101861361186957600080fd5b803567ffffffffffffffff81111561188057600080fd5b8660208260051b840101111561189557600080fd5b939660209190910195509293505050565b6020815260006118b960208301846113ab565b9392505050565b600080604083850312156118d357600080fd5b50508035926020909101359150565b6000806000806000606086880312156118fa57600080fd5b85359450602086013567ffffffffffffffff81111561191857600080fd5b611924888289016115eb565b909550935050604086013567ffffffffffffffff81111561194457600080fd5b611950888289016115eb565b969995985093965092949392505050565b600181811c9082168061197557607f821691505b60208210810361158b57634e487b7160e01b600052602260045260246000fd5b80151581146119a357600080fd5b50565b6000602082840312156119b857600080fd5b81516118b981611995565b602081016119d18284611397565b92915050565b6000808335601e198436030181126119ee57600080fd5b83018035915067ffffffffffffffff821115611a0957600080fd5b60200191503681900382131561162d57600080fd5b634e487b7160e01b600052604160045260246000fd5b601f82111561020b57806000526020600020601f840160051c81016020851015611a5b5750805b601f840160051c820191505b81811015611a7b5760008155600101611a67565b5050505050565b67ffffffffffffffff831115611a9a57611a9a611a1e565b611aae83611aa88354611961565b83611a34565b6000601f841160018114611ae25760008515611aca5750838201355b600019600387901b1c1916600186901b178355611a7b565b600083815260209020601f19861690835b82811015611b135786850135825560209485019460019092019101611af3565b5086821015611b305760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b848152606060208201526000611b85606083018587611b42565b905082604083015295945050505050565b60ff811681146119a357600080fd5b600060208284031215611bb757600080fd5b81356118b981611b96565b634e487b7160e01b600052601160045260246000fd5b808201808211156119d1576119d1611bc2565b600a81106119a357600080fd5b600060208284031215611c0a57600080fd5b81356118b981611beb565b600081356119d181611995565b600081356119d181611b96565b8135611c3a81611995565b815490151560ff1660ff1991909116178155611c75611c5b60208401611c15565b82805461ff00191691151560081b61ff0016919091179055565b611c9c611c8460408401611c22565b825462ff0000191660109190911b62ff000016178255565b611cc5611cab60608401611c22565b825463ff000000191660189190911b63ff00000016178255565b6080820135600182015560a0820135600282015560c0820135600382015560e090910135600490910155565b8135611cfc81611beb565b600a8110611d0c57611d0c611381565b815460ff191660ff9190911617815560208201356001820155604082013560028201556060820135601e1936849003018112611d4757600080fd5b8201803567ffffffffffffffff811115611d6057600080fd5b602082019150803603821315611d7557600080fd5b611d83818360038601611a82565b50505050565b600063ffffffff821663ffffffff8103611da557611da5611bc2565b60010192915050565b604081526000611dc2604083018688611b42565b8281036020840152611dd5818587611b42565b979650505050505050565b600060808284031215611df257600080fd5b82601f830112611e0157600080fd5b6040516080810181811067ffffffffffffffff82111715611e2457611e24611a1e565b604052806080840185811115611e3957600080fd5b845b81811015611e53578035835260209283019201611e3b565b50919594505050505056fea2646970667358221220b42a617edcb29e10292b96453409d30fc459be94d618fabeb93d96efd5f0a4f864736f6c634300081c0033",
}

// ProcessRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProcessRegistryMetaData.ABI instead.
var ProcessRegistryABI = ProcessRegistryMetaData.ABI

// ProcessRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProcessRegistryMetaData.Bin instead.
var ProcessRegistryBin = ProcessRegistryMetaData.Bin

// DeployProcessRegistry deploys a new Ethereum contract, binding an instance of ProcessRegistry to it.
func DeployProcessRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID string, _organizationRegistryAddress common.Address, _verifier common.Address) (common.Address, *types.Transaction, *ProcessRegistry, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProcessRegistryBin), backend, _chainID, _organizationRegistryAddress, _verifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProcessRegistry{ProcessRegistryCaller: ProcessRegistryCaller{contract: contract}, ProcessRegistryTransactor: ProcessRegistryTransactor{contract: contract}, ProcessRegistryFilterer: ProcessRegistryFilterer{contract: contract}}, nil
}

// ProcessRegistry is an auto generated Go binding around an Ethereum contract.
type ProcessRegistry struct {
	ProcessRegistryCaller     // Read-only binding to the contract
	ProcessRegistryTransactor // Write-only binding to the contract
	ProcessRegistryFilterer   // Log filterer for contract events
}

// ProcessRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProcessRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProcessRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProcessRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProcessRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProcessRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProcessRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProcessRegistrySession struct {
	Contract     *ProcessRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProcessRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProcessRegistryCallerSession struct {
	Contract *ProcessRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProcessRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProcessRegistryTransactorSession struct {
	Contract     *ProcessRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProcessRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProcessRegistryRaw struct {
	Contract *ProcessRegistry // Generic contract binding to access the raw methods on
}

// ProcessRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProcessRegistryCallerRaw struct {
	Contract *ProcessRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProcessRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProcessRegistryTransactorRaw struct {
	Contract *ProcessRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProcessRegistry creates a new instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistry(address common.Address, backend bind.ContractBackend) (*ProcessRegistry, error) {
	contract, err := bindProcessRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistry{ProcessRegistryCaller: ProcessRegistryCaller{contract: contract}, ProcessRegistryTransactor: ProcessRegistryTransactor{contract: contract}, ProcessRegistryFilterer: ProcessRegistryFilterer{contract: contract}}, nil
}

// NewProcessRegistryCaller creates a new read-only instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProcessRegistryCaller, error) {
	contract, err := bindProcessRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryCaller{contract: contract}, nil
}

// NewProcessRegistryTransactor creates a new write-only instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProcessRegistryTransactor, error) {
	contract, err := bindProcessRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryTransactor{contract: contract}, nil
}

// NewProcessRegistryFilterer creates a new log filterer instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProcessRegistryFilterer, error) {
	contract, err := bindProcessRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryFilterer{contract: contract}, nil
}

// bindProcessRegistry binds a generic wrapper to an already deployed contract.
func bindProcessRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProcessRegistry *ProcessRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProcessRegistry.Contract.ProcessRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProcessRegistry *ProcessRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.ProcessRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProcessRegistry *ProcessRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.ProcessRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProcessRegistry *ProcessRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProcessRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProcessRegistry *ProcessRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProcessRegistry *ProcessRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXCENSUSORIGIN is a free data retrieval call binding the contract method 0x02cc919a.
//
// Solidity: function MAX_CENSUS_ORIGIN() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCaller) MAXCENSUSORIGIN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "MAX_CENSUS_ORIGIN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXCENSUSORIGIN is a free data retrieval call binding the contract method 0x02cc919a.
//
// Solidity: function MAX_CENSUS_ORIGIN() view returns(uint8)
func (_ProcessRegistry *ProcessRegistrySession) MAXCENSUSORIGIN() (uint8, error) {
	return _ProcessRegistry.Contract.MAXCENSUSORIGIN(&_ProcessRegistry.CallOpts)
}

// MAXCENSUSORIGIN is a free data retrieval call binding the contract method 0x02cc919a.
//
// Solidity: function MAX_CENSUS_ORIGIN() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCallerSession) MAXCENSUSORIGIN() (uint8, error) {
	return _ProcessRegistry.Contract.MAXCENSUSORIGIN(&_ProcessRegistry.CallOpts)
}

// MAXSTATUS is a free data retrieval call binding the contract method 0x72c628ef.
//
// Solidity: function MAX_STATUS() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCaller) MAXSTATUS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "MAX_STATUS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXSTATUS is a free data retrieval call binding the contract method 0x72c628ef.
//
// Solidity: function MAX_STATUS() view returns(uint8)
func (_ProcessRegistry *ProcessRegistrySession) MAXSTATUS() (uint8, error) {
	return _ProcessRegistry.Contract.MAXSTATUS(&_ProcessRegistry.CallOpts)
}

// MAXSTATUS is a free data retrieval call binding the contract method 0x72c628ef.
//
// Solidity: function MAX_STATUS() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCallerSession) MAXSTATUS() (uint8, error) {
	return _ProcessRegistry.Contract.MAXSTATUS(&_ProcessRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProcessRegistry *ProcessRegistryCaller) ChainID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProcessRegistry *ProcessRegistrySession) ChainID() (string, error) {
	return _ProcessRegistry.Contract.ChainID(&_ProcessRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProcessRegistry *ProcessRegistryCallerSession) ChainID() (string, error) {
	return _ProcessRegistry.Contract.ChainID(&_ProcessRegistry.CallOpts)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256[] result, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistryCaller) GetProcess(opts *bind.CallOpts, processId [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	Result             []*big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcess", processId)

	outstruct := new(struct {
		Status             uint8
		OrganizationId     common.Address
		EncryptionKey      IProcessRegistryEncryptionKey
		LatestStateRoot    *big.Int
		Result             []*big.Int
		StartTime          *big.Int
		Duration           *big.Int
		VoteCount          *big.Int
		VoteOverwriteCount *big.Int
		MetadataURI        string
		BallotMode         IProcessRegistryBallotMode
		Census             IProcessRegistryCensus
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.OrganizationId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EncryptionKey = *abi.ConvertType(out[2], new(IProcessRegistryEncryptionKey)).(*IProcessRegistryEncryptionKey)
	outstruct.LatestStateRoot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)
	outstruct.StartTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VoteCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.VoteOverwriteCount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.BallotMode = *abi.ConvertType(out[10], new(IProcessRegistryBallotMode)).(*IProcessRegistryBallotMode)
	outstruct.Census = *abi.ConvertType(out[11], new(IProcessRegistryCensus)).(*IProcessRegistryCensus)

	return *outstruct, err

}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256[] result, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	Result             []*big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256[] result, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcess(processId [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	Result             []*big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// OrganizationRegistryAddress is a free data retrieval call binding the contract method 0xb79938e9.
//
// Solidity: function organizationRegistryAddress() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) OrganizationRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "organizationRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrganizationRegistryAddress is a free data retrieval call binding the contract method 0xb79938e9.
//
// Solidity: function organizationRegistryAddress() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) OrganizationRegistryAddress() (common.Address, error) {
	return _ProcessRegistry.Contract.OrganizationRegistryAddress(&_ProcessRegistry.CallOpts)
}

// OrganizationRegistryAddress is a free data retrieval call binding the contract method 0xb79938e9.
//
// Solidity: function organizationRegistryAddress() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) OrganizationRegistryAddress() (common.Address, error) {
	return _ProcessRegistry.Contract.OrganizationRegistryAddress(&_ProcessRegistry.CallOpts)
}

// ProcessCount is a free data retrieval call binding the contract method 0x848df540.
//
// Solidity: function processCount() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) ProcessCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ProcessCount is a free data retrieval call binding the contract method 0x848df540.
//
// Solidity: function processCount() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) ProcessCount() (uint32, error) {
	return _ProcessRegistry.Contract.ProcessCount(&_ProcessRegistry.CallOpts)
}

// ProcessCount is a free data retrieval call binding the contract method 0x848df540.
//
// Solidity: function processCount() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) ProcessCount() (uint32, error) {
	return _ProcessRegistry.Contract.ProcessCount(&_ProcessRegistry.CallOpts)
}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistryCaller) Processes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processes", arg0)

	outstruct := new(struct {
		Status             uint8
		OrganizationId     common.Address
		EncryptionKey      IProcessRegistryEncryptionKey
		LatestStateRoot    *big.Int
		StartTime          *big.Int
		Duration           *big.Int
		VoteCount          *big.Int
		VoteOverwriteCount *big.Int
		MetadataURI        string
		BallotMode         IProcessRegistryBallotMode
		Census             IProcessRegistryCensus
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.OrganizationId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EncryptionKey = *abi.ConvertType(out[2], new(IProcessRegistryEncryptionKey)).(*IProcessRegistryEncryptionKey)
	outstruct.LatestStateRoot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VoteCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VoteOverwriteCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.BallotMode = *abi.ConvertType(out[9], new(IProcessRegistryBallotMode)).(*IProcessRegistryBallotMode)
	outstruct.Census = *abi.ConvertType(out[10], new(IProcessRegistryCensus)).(*IProcessRegistryCensus)

	return *outstruct, err

}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistrySession) Processes(arg0 [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistryCallerSession) Processes(arg0 [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) Verifier() (common.Address, error) {
	return _ProcessRegistry.Contract.Verifier(&_ProcessRegistry.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) Verifier() (common.Address, error) {
	return _ProcessRegistry.Contract.Verifier(&_ProcessRegistry.CallOpts)
}

// NewProcess is a paid mutator transaction binding the contract method 0x92bc2669.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census, string metadata, address organizationId, bytes32 processId, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, organizationId common.Address, processId [32]byte, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, ballotMode, census, metadata, organizationId, processId, encryptionKey, initStateRoot)
}

// NewProcess is a paid mutator transaction binding the contract method 0x92bc2669.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census, string metadata, address organizationId, bytes32 processId, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns()
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, organizationId common.Address, processId [32]byte, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, ballotMode, census, metadata, organizationId, processId, encryptionKey, initStateRoot)
}

// NewProcess is a paid mutator transaction binding the contract method 0x92bc2669.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census, string metadata, address organizationId, bytes32 processId, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, organizationId common.Address, processId [32]byte, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, ballotMode, census, metadata, organizationId, processId, encryptionKey, initStateRoot)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6bae04ea.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,uint256,bytes32,string) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessCensus(opts *bind.TransactOpts, processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessCensus", processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6bae04ea.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,uint256,bytes32,string) census) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessCensus(processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6bae04ea.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,uint256,bytes32,string) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessCensus(processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0xc718c01f.
//
// Solidity: function setProcessDuration(bytes32 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessDuration(opts *bind.TransactOpts, processId [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessDuration", processId, _duration)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0xc718c01f.
//
// Solidity: function setProcessDuration(bytes32 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessDuration(processId [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessDuration(&_ProcessRegistry.TransactOpts, processId, _duration)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0xc718c01f.
//
// Solidity: function setProcessDuration(bytes32 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessDuration(processId [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessDuration(&_ProcessRegistry.TransactOpts, processId, _duration)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xac64bd03.
//
// Solidity: function setProcessResults(bytes32 processId, uint256[] results) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessResults(opts *bind.TransactOpts, processId [32]byte, results []*big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessResults", processId, results)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xac64bd03.
//
// Solidity: function setProcessResults(bytes32 processId, uint256[] results) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessResults(processId [32]byte, results []*big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, results)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xac64bd03.
//
// Solidity: function setProcessResults(bytes32 processId, uint256[] results) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessResults(processId [32]byte, results []*big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, results)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x56a6f1e2.
//
// Solidity: function setProcessStatus(bytes32 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessStatus(opts *bind.TransactOpts, processId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessStatus", processId, newStatus)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x56a6f1e2.
//
// Solidity: function setProcessStatus(bytes32 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessStatus(processId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessStatus(&_ProcessRegistry.TransactOpts, processId, newStatus)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x56a6f1e2.
//
// Solidity: function setProcessStatus(bytes32 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessStatus(processId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessStatus(&_ProcessRegistry.TransactOpts, processId, newStatus)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xd71d072e.
//
// Solidity: function submitStateTransition(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SubmitStateTransition(opts *bind.TransactOpts, processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "submitStateTransition", processId, proof, input)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xd71d072e.
//
// Solidity: function submitStateTransition(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistrySession) SubmitStateTransition(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransition(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xd71d072e.
//
// Solidity: function submitStateTransition(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SubmitStateTransition(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransition(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// ProcessRegistryCensusUpdatedIterator is returned from FilterCensusUpdated and is used to iterate over the raw logs and unpacked data for CensusUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryCensusUpdatedIterator struct {
	Event *ProcessRegistryCensusUpdated // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryCensusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryCensusUpdated)
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
		it.Event = new(ProcessRegistryCensusUpdated)
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
func (it *ProcessRegistryCensusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryCensusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryCensusUpdated represents a CensusUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryCensusUpdated struct {
	ProcessId  [32]byte
	CensusRoot [32]byte
	CensusURI  string
	MaxVotes   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCensusUpdated is a free log retrieval operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI, uint256 maxVotes)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterCensusUpdated(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryCensusUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "CensusUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryCensusUpdatedIterator{contract: _ProcessRegistry.contract, event: "CensusUpdated", logs: logs, sub: sub}, nil
}

// WatchCensusUpdated is a free log subscription operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI, uint256 maxVotes)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchCensusUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryCensusUpdated, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "CensusUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryCensusUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "CensusUpdated", log); err != nil {
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

// ParseCensusUpdated is a log parse operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI, uint256 maxVotes)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseCensusUpdated(log types.Log) (*ProcessRegistryCensusUpdated, error) {
	event := new(ProcessRegistryCensusUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "CensusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessCreatedIterator is returned from FilterProcessCreated and is used to iterate over the raw logs and unpacked data for ProcessCreated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessCreatedIterator struct {
	Event *ProcessRegistryProcessCreated // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessCreated)
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
		it.Event = new(ProcessRegistryProcessCreated)
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
func (it *ProcessRegistryProcessCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessCreated represents a ProcessCreated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessCreated struct {
	ProcessId [32]byte
	Creator   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessCreated is a free log retrieval operation binding the contract event 0xada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa.
//
// Solidity: event ProcessCreated(bytes32 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessCreated(opts *bind.FilterOpts, processId [][32]byte, creator []common.Address) (*ProcessRegistryProcessCreatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessCreated", processIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessCreatedIterator{contract: _ProcessRegistry.contract, event: "ProcessCreated", logs: logs, sub: sub}, nil
}

// WatchProcessCreated is a free log subscription operation binding the contract event 0xada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa.
//
// Solidity: event ProcessCreated(bytes32 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessCreated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessCreated, processId [][32]byte, creator []common.Address) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessCreated", processIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessCreated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessCreated", log); err != nil {
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

// ParseProcessCreated is a log parse operation binding the contract event 0xada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa.
//
// Solidity: event ProcessCreated(bytes32 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessCreated(log types.Log) (*ProcessRegistryProcessCreated, error) {
	event := new(ProcessRegistryProcessCreated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessDurationChangedIterator is returned from FilterProcessDurationChanged and is used to iterate over the raw logs and unpacked data for ProcessDurationChanged events raised by the ProcessRegistry contract.
type ProcessRegistryProcessDurationChangedIterator struct {
	Event *ProcessRegistryProcessDurationChanged // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessDurationChanged)
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
		it.Event = new(ProcessRegistryProcessDurationChanged)
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
func (it *ProcessRegistryProcessDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessDurationChanged represents a ProcessDurationChanged event raised by the ProcessRegistry contract.
type ProcessRegistryProcessDurationChanged struct {
	ProcessId [32]byte
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessDurationChanged is a free log retrieval operation binding the contract event 0x0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46.
//
// Solidity: event ProcessDurationChanged(bytes32 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessDurationChanged(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessDurationChangedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessDurationChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessDurationChangedIterator{contract: _ProcessRegistry.contract, event: "ProcessDurationChanged", logs: logs, sub: sub}, nil
}

// WatchProcessDurationChanged is a free log subscription operation binding the contract event 0x0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46.
//
// Solidity: event ProcessDurationChanged(bytes32 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessDurationChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessDurationChanged, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessDurationChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessDurationChanged)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessDurationChanged", log); err != nil {
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

// ParseProcessDurationChanged is a log parse operation binding the contract event 0x0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46.
//
// Solidity: event ProcessDurationChanged(bytes32 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessDurationChanged(log types.Log) (*ProcessRegistryProcessDurationChanged, error) {
	event := new(ProcessRegistryProcessDurationChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessStateRootUpdatedIterator is returned from FilterProcessStateRootUpdated and is used to iterate over the raw logs and unpacked data for ProcessStateRootUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateRootUpdatedIterator struct {
	Event *ProcessRegistryProcessStateRootUpdated // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessStateRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessStateRootUpdated)
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
		it.Event = new(ProcessRegistryProcessStateRootUpdated)
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
func (it *ProcessRegistryProcessStateRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessStateRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessStateRootUpdated represents a ProcessStateRootUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateRootUpdated struct {
	ProcessId    [32]byte
	NewStateRoot *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProcessStateRootUpdated is a free log retrieval operation binding the contract event 0xd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20.
//
// Solidity: event ProcessStateRootUpdated(bytes32 indexed processId, uint256 newStateRoot)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStateRootUpdated(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessStateRootUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessStateRootUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessStateRootUpdatedIterator{contract: _ProcessRegistry.contract, event: "ProcessStateRootUpdated", logs: logs, sub: sub}, nil
}

// WatchProcessStateRootUpdated is a free log subscription operation binding the contract event 0xd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20.
//
// Solidity: event ProcessStateRootUpdated(bytes32 indexed processId, uint256 newStateRoot)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStateRootUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStateRootUpdated, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessStateRootUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessStateRootUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateRootUpdated", log); err != nil {
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

// ParseProcessStateRootUpdated is a log parse operation binding the contract event 0xd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20.
//
// Solidity: event ProcessStateRootUpdated(bytes32 indexed processId, uint256 newStateRoot)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStateRootUpdated(log types.Log) (*ProcessRegistryProcessStateRootUpdated, error) {
	event := new(ProcessRegistryProcessStateRootUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessStatusChangedIterator is returned from FilterProcessStatusChanged and is used to iterate over the raw logs and unpacked data for ProcessStatusChanged events raised by the ProcessRegistry contract.
type ProcessRegistryProcessStatusChangedIterator struct {
	Event *ProcessRegistryProcessStatusChanged // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessStatusChanged)
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
		it.Event = new(ProcessRegistryProcessStatusChanged)
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
func (it *ProcessRegistryProcessStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessStatusChanged represents a ProcessStatusChanged event raised by the ProcessRegistry contract.
type ProcessRegistryProcessStatusChanged struct {
	ProcessId [32]byte
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessStatusChanged is a free log retrieval operation binding the contract event 0xac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c190.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStatusChanged(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessStatusChangedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessStatusChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessStatusChangedIterator{contract: _ProcessRegistry.contract, event: "ProcessStatusChanged", logs: logs, sub: sub}, nil
}

// WatchProcessStatusChanged is a free log subscription operation binding the contract event 0xac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c190.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStatusChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStatusChanged, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessStatusChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessStatusChanged)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusChanged", log); err != nil {
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

// ParseProcessStatusChanged is a log parse operation binding the contract event 0xac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c190.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStatusChanged(log types.Log) (*ProcessRegistryProcessStatusChanged, error) {
	event := new(ProcessRegistryProcessStatusChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
