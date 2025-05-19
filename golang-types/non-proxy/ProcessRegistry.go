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

// IProcessRegistryProcess is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryProcess struct {
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
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chainID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_organizationRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidCensus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrganizationAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrganizationNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"ProcessStateRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CENSUS_ORIGIN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOverwriteCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structIProcessRegistry.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"initStateRoot\",\"type\":\"uint256\"}],\"name\":\"newProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOverwriteCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"results\",\"type\":\"uint256[]\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161226538038061226583398101604081905261002f916100a0565b600261003b8482610216565b50600180546001600160a01b039384166001600160a01b03199182161790915560038054929093169116179055506102d4565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b038116811461009b57600080fd5b919050565b6000806000606084860312156100b557600080fd5b83516001600160401b038111156100cb57600080fd5b8401601f810186136100dc57600080fd5b80516001600160401b038111156100f5576100f561006e565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101235761012361006e565b60405281815282820160200188101561013b57600080fd5b60005b8281101561015a5760208185018101518383018201520161013e565b5060006020838301015280955050505061017660208501610084565b915061018460408501610084565b90509250925092565b600181811c908216806101a157607f821691505b6020821081036101c157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561021157806000526020600020601f840160051c810160208510156101ee5750805b601f840160051c820191505b8181101561020e57600081556001016101fa565b50505b505050565b81516001600160401b0381111561022f5761022f61006e565b6102438161023d845461018d565b846101c7565b6020601f821160018114610277576000831561025f5750848201515b600019600385901b1c1916600184901b17845561020e565b600084815260208120601f198516915b828110156102a75787850151825560209485019460019092019101610287565b50848210156102c55786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b611f82806102e36000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063848df54011610097578063adc879e911610066578063adc879e914610226578063b79938e91461023b578063c718c01f1461024e578063d71d072e1461026157600080fd5b8063848df540146101b457806392bc2669146101e0578063992bc45b146101f3578063ac64bd031461021357600080fd5b806356a6f1e2116100d357806356a6f1e21461016e5780636bae04ea1461018357806372c628ef146101965780637d2323541461019e57600080fd5b806302cc919a146100fa5780630535fece146101195780632b7ac3f314610143575b600080fd5b610102600981565b60405160ff90911681526020015b60405180910390f35b61012c6101273660046113ee565b610274565b6040516101109b9a99989796959493929190611524565b600354610156906001600160a01b031681565b6040516001600160a01b039091168152602001610110565b61018161017c3660046115d3565b6104b6565b005b610181610191366004611617565b610685565b610102600481565b6101a66108b2565b604051908152602001610110565b6001546101cb90600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610110565b6101816101ee3660046116e3565b610925565b6102066102013660046113ee565b610c80565b6040516101109190611806565b6101816102213660046118fa565b505050565b61022e610f83565b604051610110919061197b565b600154610156906001600160a01b031681565b61018161025c366004611995565b611011565b61018161026f3660046119b7565b611191565b6000602081815291815260409081902080548251808401909352600182015483526002820154938301939093526003810154600582015460068301546007840154600885015460098601805460ff8a169961010090046001600160a01b0316989791906102e090611a36565b80601f016020809104026020016040519081016040528092919081815260200182805461030c90611a36565b80156103595780601f1061032e57610100808354040283529160200191610359565b820191906000526020600020905b81548152906001019060200180831161033c57829003601f168201915b5050604080516101008082018352600a88015460ff808216151584529181048216151560208401526201000081048216838501526301000000900481166060830152600b880154608080840191909152600c89015460a0840152600d89015460c0840152600e89015460e08401528351908101909352600f880180549798929792965092945091925083911660098111156103f6576103f6611407565b600981111561040757610407611407565b8152602001600182015481526020016002820154815260200160038201805461042f90611a36565b80601f016020809104026020016040519081016040528092919081815260200182805461045b90611a36565b80156104a85780601f1061047d576101008083540402835291602001916104a8565b820191906000526020600020905b81548152906001019060200180831161048b57829003601f168201915b50505050508152505090508b565b6000828152602081905260409020805461010090046001600160a01b0316806104f257604051634d36eb6960e01b815260040160405180910390fd5b815460ff1680600481111561050957610509611407565b84600481111561051b5761051b611407565b148061053a57506004848181111561053557610535611407565b60ff16115b806105745750600081600481111561055457610554611407565b141580156105745750600381600481111561057157610571611407565b14155b15610592576040516307a92f1960e51b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0384811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa1580156105e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106069190611a7b565b6106235760405163f8b857d360e01b815260040160405180910390fd5b82548490849060ff1916600183600481111561064157610641611407565b0217905550847fac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c190856040516106769190611a98565b60405180910390a25050505050565b6106926060820182611aac565b90506000036106b45760405163111127cb60e01b815260040160405180910390fd5b60408101356000036106d95760405163111127cb60e01b815260040160405180910390fd5b6000828152602081905260409020805461010090046001600160a01b03168061071557604051634d36eb6960e01b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0383811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa158015610765573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107899190611a7b565b6107a65760405163f8b857d360e01b815260040160405180910390fd5b6010820154602084013510156107cf5760405163111127cb60e01b815260040160405180910390fd5b6000825460ff1660048111156107e7576107e7611407565b1415801561080b57506003825460ff16600481111561080857610808611407565b14155b15610829576040516307a92f1960e51b815260040160405180910390fd5b602083013560108301556040830135601183015561084a6060840184611aac565b601284019161085a919083611b57565b50837f35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b604085013561088f6060870187611aac565b87602001356040516108a49493929190611c40565b60405180910390a250505050565b6003546040805163233ace1160e01b815290516000926001600160a01b03169163233ace119160048083019260209291908290030181865afa1580156108fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109209190611c6b565b905090565b60015460405163c1af6e0360e01b81526001600160a01b0386811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa158015610975573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109999190611a7b565b6109b65760405163f8b857d360e01b815260040160405180910390fd5b6000838152602081905260409020805461010090046001600160a01b0316156109f257604051635040c4d360e11b815260040160405180910390fd5b610a0260608a0160408b01611c93565b60ff16600003610a255760405163ac38930b60e01b815260040160405180910390fd5b610a3560608a0160408b01611c93565b60ff16896080013511610a5b5760405163b1911d8b60e01b815260040160405180910390fd5b60048c81811115610a6e57610a6e611407565b60ff161180610aac575060008c6004811115610a8c57610a8c611407565b14158015610aac575060038c6004811115610aa957610aa9611407565b14155b15610aca576040516307a92f1960e51b815260040160405180910390fd5b4260008c9003610ad857809b505b808c1015610af957604051632ca4094f60e21b815260040160405180910390fd5b80610b048c8e611cc6565b11610b2257604051637616640160e01b815260040160405180910390fd5b6009610b3160208b018b611ce6565b6009811115610b4257610b42611407565b60ff161115610b645760405163111127cb60e01b815260040160405180910390fd5b81548d90839060ff19166001836004811115610b8257610b82611407565b0217905550600582018c9055600682018b90558154610100600160a81b0319166101006001600160a01b0388160217825583356001830155602084013560028301556003820183905560098201610bda888a83611b57565b5089600a8301610bea8282611d1d565b50899050600f8301610bfc8282611ddf565b505060018054600160a01b900463ffffffff16906014610c1b83611e77565b91906101000a81548163ffffffff021916908363ffffffff16021790555050336001600160a01b0316857fada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa60405160405180910390a350505050505050505050505050565b610c886112fc565b60008281526020819052604090819020815161018081019092528054829060ff166004811115610cba57610cba611407565b6004811115610ccb57610ccb611407565b8152815461010090046001600160a01b03166020808301919091526040805180820182526001850154815260028501548184015281840152600384015460608401526004840180548251818502810185019093528083526080909401939192909190830182828015610d5c57602002820191906000526020600020905b815481526020019060010190808311610d48575b5050505050815260200160058201548152602001600682015481526020016007820154815260200160088201548152602001600982018054610d9d90611a36565b80601f0160208091040260200160405190810160405280929190818152602001828054610dc990611a36565b8015610e165780601f10610deb57610100808354040283529160200191610e16565b820191906000526020600020905b815481529060010190602001808311610df957829003601f168201915b5050509183525050604080516101008082018352600a85015460ff8082161515845291810482161515602084810191909152620100008204831684860152630100000090910482166060840152600b860154608080850191909152600c87015460a0850152600d87015460c0850152600e87015460e08501529085019290925282519182018352600f85018054939094019391928391166009811115610ebe57610ebe611407565b6009811115610ecf57610ecf611407565b81526020016001820154815260200160028201548152602001600382018054610ef790611a36565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2390611a36565b8015610f705780601f10610f4557610100808354040283529160200191610f70565b820191906000526020600020905b815481529060010190602001808311610f5357829003601f168201915b5050509190925250505090525092915050565b60028054610f9090611a36565b80601f0160208091040260200160405190810160405280929190818152602001828054610fbc90611a36565b80156110095780601f10610fde57610100808354040283529160200191611009565b820191906000526020600020905b815481529060010190602001808311610fec57829003601f168201915b505050505081565b42811161103157604051637616640160e01b815260040160405180910390fd5b6000828152602081905260409020805461010090046001600160a01b03168061106d57604051634d36eb6960e01b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0383811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa1580156110bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e19190611a7b565b6110fe5760405163f8b857d360e01b815260040160405180910390fd5b6000825460ff16600481111561111657611116611407565b1415801561113a57506003825460ff16600481111561113757611137611407565b14155b15611158576040516307a92f1960e51b815260040160405180910390fd5b6006820183905560405183815284907f0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46906020016108a4565b6000858152602081905260409020805461010090046001600160a01b03166111cc57604051634d36eb6960e01b815260040160405180910390fd5b6000815460ff1660048111156111e4576111e4611407565b14611202576040516307a92f1960e51b815260040160405180910390fd5b600354604051635c73957b60e11b81526001600160a01b039091169063b8e72af690611238908890889088908890600401611e9c565b60006040518083038186803b15801561125057600080fd5b505afa158015611264573d6000803e3d6000fd5b506000925061127891505083850185611ece565b60038301548151919250146112a057604051630b6fac0360e41b815260040160405180910390fd5b602080820151600384018190556040808401516007860155606084015160088601555190815288917fd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20910160405180910390a250505050505050565b604080516101808101909152806000815260200160006001600160a01b0316815260200161133d604051806040016040528060008152602001600081525090565b8152602001600081526020016060815260200160008152602001600081526020016000815260200160008152602001606081526020016113c5604051806101000160405280600015158152602001600015158152602001600060ff168152602001600060ff168152602001600081526020016000815260200160008152602001600081525090565b815260408051608081018252600080825260208281018290529282015260608082015291015290565b60006020828403121561140057600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b6005811061142d5761142d611407565b9052565b6000815180845260005b818110156114575760208185018101518683018201520161143b565b506000602082860101526020601f19601f83011685010191505092915050565b80511515825260208101511515602083015260ff604082015116604083015260608101516114aa606084018260ff169052565b506080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b60008151600a81106114eb576114eb611407565b80845250602082015160208401526040820151604084015260608201516080606085015261151c6080850182611431565b949350505050565b61152e818d61141d565b6001600160a01b038b166020820152611554604082018b80518252602090810151910152565b8860808201528760a08201528660c08201528560e082015284610100820152610260610120820152600061158c610260830186611431565b61159a610140840186611477565b8281036102408401526115ad81856114d7565b9e9d5050505050505050505050505050565b8035600581106115ce57600080fd5b919050565b600080604083850312156115e657600080fd5b823591506115f6602084016115bf565b90509250929050565b60006080828403121561161157600080fd5b50919050565b6000806040838503121561162a57600080fd5b82359150602083013567ffffffffffffffff81111561164857600080fd5b611654858286016115ff565b9150509250929050565b6000610100828403121561161157600080fd5b60008083601f84011261168357600080fd5b50813567ffffffffffffffff81111561169b57600080fd5b6020830191508360208285010111156116b357600080fd5b9250929050565b80356001600160a01b03811681146115ce57600080fd5b60006040828403121561161157600080fd5b60008060008060008060008060008060006102408c8e03121561170557600080fd5b61170e8c6115bf565b9a5060208c0135995060408c0135985061172b8d60608e0161165e565b97506101608c013567ffffffffffffffff81111561174857600080fd5b6117548e828f016115ff565b9750506101808c013567ffffffffffffffff81111561177257600080fd5b61177e8e828f01611671565b909750955061179290506101a08d016116ba565b93506101c08c013592506117aa8d6101e08e016116d1565b915060006102208d01359050809150509295989b509295989b9093969950565b600081518084526020840193506020830160005b828110156117fc5781518652602095860195909101906001016117de565b5093949350505050565b6020815261181860208201835161141d565b6000602083015161183460408401826001600160a01b03169052565b506040830151805160608401526020810151608084015250606083015160a0830152608083015161028060c08401526118716102a08401826117ca565b905060a084015160e084015260c084015161010084015260e0840151610120840152610100840151610140840152610120840151601f19848303016101608501526118bc8282611431565b9150506101408401516118d3610180850182611477565b50610160840151838203601f19016102808501526118f182826114d7565b95945050505050565b60008060006040848603121561190f57600080fd5b83359250602084013567ffffffffffffffff81111561192d57600080fd5b8401601f8101861361193e57600080fd5b803567ffffffffffffffff81111561195557600080fd5b8660208260051b840101111561196a57600080fd5b939660209190910195509293505050565b60208152600061198e6020830184611431565b9392505050565b600080604083850312156119a857600080fd5b50508035926020909101359150565b6000806000806000606086880312156119cf57600080fd5b85359450602086013567ffffffffffffffff8111156119ed57600080fd5b6119f988828901611671565b909550935050604086013567ffffffffffffffff811115611a1957600080fd5b611a2588828901611671565b969995985093965092949392505050565b600181811c90821680611a4a57607f821691505b60208210810361161157634e487b7160e01b600052602260045260246000fd5b8015158114611a7857600080fd5b50565b600060208284031215611a8d57600080fd5b815161198e81611a6a565b60208101611aa6828461141d565b92915050565b6000808335601e19843603018112611ac357600080fd5b83018035915067ffffffffffffffff821115611ade57600080fd5b6020019150368190038213156116b357600080fd5b634e487b7160e01b600052604160045260246000fd5b601f82111561022157806000526020600020601f840160051c81016020851015611b305750805b601f840160051c820191505b81811015611b505760008155600101611b3c565b5050505050565b67ffffffffffffffff831115611b6f57611b6f611af3565b611b8383611b7d8354611a36565b83611b09565b6000601f841160018114611bb75760008515611b9f5750838201355b600019600387901b1c1916600186901b178355611b50565b600083815260209020601f19861690835b82811015611be85786850135825560209485019460019092019101611bc8565b5086821015611c055760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b848152606060208201526000611c5a606083018587611c17565b905082604083015295945050505050565b600060208284031215611c7d57600080fd5b5051919050565b60ff81168114611a7857600080fd5b600060208284031215611ca557600080fd5b813561198e81611c84565b634e487b7160e01b600052601160045260246000fd5b80820180821115611aa657611aa6611cb0565b600a8110611a7857600080fd5b600060208284031215611cf857600080fd5b813561198e81611cd9565b60008135611aa681611a6a565b60008135611aa681611c84565b8135611d2881611a6a565b815490151560ff1660ff1991909116178155611d63611d4960208401611d03565b82805461ff00191691151560081b61ff0016919091179055565b611d8a611d7260408401611d10565b825462ff0000191660109190911b62ff000016178255565b611db3611d9960608401611d10565b825463ff000000191660189190911b63ff00000016178255565b6080820135600182015560a0820135600282015560c0820135600382015560e090910135600490910155565b8135611dea81611cd9565b600a8110611dfa57611dfa611407565b815460ff191660ff9190911617815560208201356001820155604082013560028201556060820135601e1936849003018112611e3557600080fd5b8201803567ffffffffffffffff811115611e4e57600080fd5b602082019150803603821315611e6357600080fd5b611e71818360038601611b57565b50505050565b600063ffffffff821663ffffffff8103611e9357611e93611cb0565b60010192915050565b604081526000611eb0604083018688611c17565b8281036020840152611ec3818587611c17565b979650505050505050565b600060808284031215611ee057600080fd5b82601f830112611eef57600080fd5b6040516080810181811067ffffffffffffffff82111715611f1257611f12611af3565b604052806080840185811115611f2757600080fd5b845b81811015611f41578035835260209283019201611f29565b50919594505050505056fea2646970667358221220fef00bb8c7deb21be694ab92b6cd7779b6e575a6ee5db0aa269ae99faa73dffd64736f6c634300081c0033",
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
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,uint256,bytes32,string)))
func (_ProcessRegistry *ProcessRegistryCaller) GetProcess(opts *bind.CallOpts, processId [32]byte) (IProcessRegistryProcess, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcess", processId)

	if err != nil {
		return *new(IProcessRegistryProcess), err
	}

	out0 := *abi.ConvertType(out[0], new(IProcessRegistryProcess)).(*IProcessRegistryProcess)

	return out0, err

}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,uint256,bytes32,string)))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [32]byte) (IProcessRegistryProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,uint256,bytes32,string)))
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcess(processId [32]byte) (IProcessRegistryProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetVerifierVKeyHash is a free data retrieval call binding the contract method 0x7d232354.
//
// Solidity: function getVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) GetVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifierVKeyHash is a free data retrieval call binding the contract method 0x7d232354.
//
// Solidity: function getVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) GetVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// GetVerifierVKeyHash is a free data retrieval call binding the contract method 0x7d232354.
//
// Solidity: function getVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetVerifierVKeyHash(&_ProcessRegistry.CallOpts)
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
