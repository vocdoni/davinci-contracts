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

// DAVINCITypesBallotMode is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesBallotMode struct {
	CostFromWeight bool
	UniqueValues   bool
	NumFields      uint8
	GroupSize      uint8
	CostExponent   uint8
	MaxValue       *big.Int
	MinValue       *big.Int
	MaxValueSum    *big.Int
	MinValueSum    *big.Int
}

// DAVINCITypesCensus is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesCensus struct {
	CensusOrigin             uint8
	CensusRoot               [32]byte
	ContractAddress          common.Address
	CensusURI                string
	OnchainAllowAnyValidRoot bool
}

// DAVINCITypesEncryptionKey is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesEncryptionKey struct {
	X *big.Int
	Y *big.Int
}

// DAVINCITypesProcess is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesProcess struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         DAVINCITypesEncryptionKey
	LatestStateRoot       *big.Int
	Result                []*big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_stVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rVerifier\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_blobsDA\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"limbIndex\",\"type\":\"uint8\"}],\"name\":\"InvalidBlobCommitmentLimb\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValueSum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOB_INDEX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CENSUS_ORIGIN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobsDA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461017757608081612e11803803809161001f828561018e565b83398101031261017757610032816101c5565b9061003f602082016101d6565b606061004d604084016101d6565b920151801515809103610177576002805460038054600160201b600160e01b0319909216604095861b600160401b600160e01b031617602088811b67ffffffff000000001691909117909355600164ff0000000160a01b03199091166001600160a01b039095169490941760c09290921b60ff60c01b16919091179092555163342e1df160e21b815263ffffffff90921660048301523060248301528160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610183575f91610145575b506003805463ffffffff60a01b191660a09290921b63ffffffff60a01b16919091179055604051612c2690816101eb8239f35b90506020813d60201161017b575b816101606020938361018e565b8101031261017757610171906101c5565b5f610112565b5f80fd5b3d9150610153565b6040513d5f823e3d90fd5b601f909101601f19168101906001600160401b038211908210176101b157604052565b634e487b7160e01b5f52604160045260245ffd5b519063ffffffff8216820361017757565b51906001600160a01b03821682036101775756fe6080806040526004361015610012575f80fd5b5f905f3560e01c90816302cc919a146125d4575080630535fece146124b5578063097063921461248b57806326795bc3146124715780634c0acc561461242157806352016fe014611b0557806356a6f1e21461195957806362fa11fc1461191457806368141f2c1461186957806372c628ef1461184e57806374e07826146116ff578063848df540146116dc578063992bc45b1461139b5780639c18deaf14611373578063a5b110de14610f7b578063a5f01e1914610e2a578063adc879e914610e04578063aeaa8b0f146109f0578063c718c01f1461081e578063cddf08bc146107f7578063cdf497c8146107d1578063d71d072e146101b05763f9aa44991461011b575f80fd5b346101ad57806003193601126101ad5760035460405163233ace1160e01b81529190602090839060049082906001600160a01b03165afa9081156101a1579061016a575b602090604051908152f35b506020813d602011610199575b8161018460209383612677565b81010312610195576020905161015f565b5f80fd5b3d9150610177565b604051903d90823e3d90fd5b80fd5b50346101ad576101bf36612993565b849394929192156107c2576003546040516350f3895360e11b81526004810186905260a09190911c63ffffffff16602482015260208160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156105fa578791610793575b501561078457838652856020526040862092835460018060a01b038160081c16156107755760ff1661024e816127fe565b6107665761026560058501546006860154906129e8565b421015610757578660e060405161027b8161265b565b8281528260208201528260408201528260608201528260808201528260a08201528260c0820152015281810192610100828503126104a45783601f830112156104a457604051936102ce61010086612677565b84610100840191821161073f5783905b82821061072f5750505083519260208501519760408601519360608701519360808801519760a0810151988d60c08301519260e00151916040519a6103228c61265b565b8b5260208b019e8f5260408b01998a5260608b0198895260808b0190815260a08b019b8c5260c08b0193845260e08b0192835260128d015460ff16610366816127fe565b600314610673575b505088519760038c01988954036106645751875161038b91612aae565b998a151580610652575b61064357908e96959493929160ff60035460c01c166104b7575b505060025460401c6001600160a01b0316919050813b156104b35785936103ec60405196879586948594635c73957b60e11b865260048601612adb565b03915afa80156104a85761048f575b5050600b9187519055610413600886019485546129e8565b8094555193610427600982019586546129e8565b8095550180545f19811461047b5760010190555193516040519485526020850152604084015260608301527f43b1a882c73bc84ae667608d3a61baececd99134224962eba9d13eb63c595ae160803393a380f35b634e487b7160e01b88526011600452602488fd5b8161049991612677565b6104a457875f6103fb565b8780fd5b6040513d84823e3d90fd5b8580fd5b9091929394959650519151905190604051926104d4606085612677565b60308452602084019060403683378060801c61062d5760801b90528060801c6106195760801b60308301528060801c61060557908d959493929160801b6040820152610552602073__$176218451d0b397034f549dcd265d6db5e$__926040518093819263da23b04d60e01b83528460048401526024830190612878565b0381855af49081156105fa5787916105c5575b50813b156105c1578690602460405180948193634e733ef560e11b835260048301525af49081156105b657869161059d575b806103af565b816105a791612677565b6105b257845f610597565b8480fd5b6040513d88823e3d90fd5b8680fd5b9650506020863d6020116105f2575b816105e160209383612677565b81010312610195578d95515f610565565b3d91506105d4565b6040513d89823e3d90fd5b6316416a3d60e01b8e52600260045260248efd5b6316416a3d60e01b8f52600160045260248ffd5b50506316416a3d60e01b8f5260048f905260248ffd5b6357d18d5360e11b8f5260048ffd5b5060088c015460078d01541115610395565b630b6fac0360e41b8f5260048ffd5b60148d0154905160405163650e5fcf60e01b8152600481019190915290602090829060249082906001600160a01b03165afa9182156101a157916106fd575b5060ff60168d0154161590816106ef575b81156106e5575b506106d6578d5f61036e565b635e32eadd60e01b8e5260048efd5b905043105f6106ca565b600a8d0154811091506106c3565b90506020813d602011610727575b8161071860209383612677565b8101031261019557515f6106b2565b3d915061070b565b81358152602091820191016102de565b8980fd5b634e487b7160e01b5f52604160045260245ffd5b63e843c5eb60e01b8752600487fd5b6307a92f1960e51b8752600487fd5b634d36eb6960e01b8852600488fd5b632299770d60e11b8652600486fd5b6107b5915060203d6020116107bb575b6107ad8183612677565b810190612a96565b5f61021d565b503d6107a3565b63cbf4a64560e01b8652600486fd5b50346101ad57806003193601126101ad57602060ff60035460c01c166040519015158152f35b50346101ad57806003193601126101ad57602063ffffffff60035460a01c16604051908152f35b50346101ad5761082d3661297d565b81156109e1576003546040516350f3895360e11b81526004810184905260a09190911c63ffffffff16602482015260208160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156109d65784916109b7575b50156109a8578183526020839052604083208054600881901c6001600160a01b0316801561099957330361098b5760ff166108c0816127fe565b8015159081610976575b5061096757600660058201549101908154908315918215610952575b8215610933575b505061092457817f0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f469260209255604051908152a280f35b637616640160e01b8452600484fd5b81925061094385610949936129e8565b926129e8565b10155f806108ed565b915061095e84826129e8565b421015916108e6565b6307a92f1960e51b8452600484fd5b60039150610983816127fe565b14155f6108ca565b6282b42960e81b8552600485fd5b634d36eb6960e01b8652600486fd5b632299770d60e11b8352600483fd5b6109d0915060203d6020116107bb576107ad8183612677565b5f610886565b6040513d86823e3d90fd5b63cbf4a64560e01b8352600483fd5b50346101ad5760403660031901126101ad576004356024356001600160401b038111610e00578060040160a06003198336030112610dfc578215610ded576003546040516350f3895360e11b81526004810185905260a09190911c63ffffffff16602482015260208160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610de2578591610dc3575b5015610db45782845260208490526040842080549190600883901c6001600160a01b03168015610da5573303610d975760ff60128201541692610ac3846127fe565b60028403610d8857823560058110156104a457610adf816127fe565b610ae8856127fe565b8403610d79576024850135948515610d6a576064810194610b098686612a64565b905015610d5b57600390610b1c816127fe565b149081610d3d575b50610d2e5760ff16610b35816127fe565b8015159081610d19575b50610d0a57610b5760058201546006830154906129e8565b421015610cfb578360138201556015610b708484612a64565b91909201916001600160401b038211610ce757610b8d83546126b6565b601f8111610cac575b508790601f8311600114610c1e579282610c0d96937f87e1c82bc035c552e997ec2f1cf023790ced0cc01098d516dc6e7200835ffb27989693610bf2968c92610c13575b50508160011b915f199060031b1c1916179055612a64565b92906040519384938452604060208501526040840191612abb565b0390a280f35b013590505f80610bda565b8389526020892091601f1984168a5b818110610c945750937f87e1c82bc035c552e997ec2f1cf023790ced0cc01098d516dc6e7200835ffb27989693610bf2969360019383610c0d9b9810610c7b575b505050811b019055612a64565b01355f19600384901b60f8161c191690555f8080610c6e565b91936020600181928787013581550195019201610c2d565b610cd790848a5260208a20601f850160051c81019160208610610cdd575b601f0160051c0190612a0d565b5f610b96565b9091508190610cca565b634e487b7160e01b88526041600452602488fd5b63e843c5eb60e01b8652600486fd5b6307a92f1960e51b8652600486fd5b60039150610d26816127fe565b14155f610b3f565b63562e597160e11b8752600487fd5b6001600160a01b039150610d5390604401612a50565b16155f610b24565b630f8b932160e11b8952600489fd5b635e32eadd60e01b8852600488fd5b63f37f7b5d60e01b8752600487fd5b63050b77c760e21b8752600487fd5b6282b42960e81b8652600486fd5b634d36eb6960e01b8752600487fd5b632299770d60e11b8452600484fd5b610ddc915060203d6020116107bb576107ad8183612677565b5f610a81565b6040513d87823e3d90fd5b63cbf4a64560e01b8452600484fd5b8380fd5b8280fd5b50346101ad57806003193601126101ad57602063ffffffff600254821c16604051908152f35b50346101ad5760203660031901126101ad57600435815280602052604081209060405190610e57826125ed565b825460ff8116610e66816127fe565b835260081c6001600160a01b03166020830152610e8560018401612698565b60408301526003830154606083015260048301604051918260208354918281520192825260208220915b818110610f65576020610f5d888888610eca818a0382612677565b608082015260058201549060a081019182526101c0610f51601260068601549560c08501968752600781015460e086015260088101546101008601526009810154610120860152600a810154610140860152600b810154610160860152610f33600c82016126ee565b610180860152610f45600d820161278e565b6101a08601520161281c565b910152519051906129e8565b604051908152f35b8254845260209093019260019283019201610eaf565b503461019557610f8a36612993565b9492938092919215611364576003546040516350f3895360e11b81526004810183905260a09190911c63ffffffff16602482015260208160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156112d0575f91611345575b501561133657805f525f60205260405f209586549560018060a01b038760081c16156113275760ff87169361101d856127fe565b600285148015611314575b61130557611035856127fe565b6001851415806112ea575b6112db576003546001600160a01b031691823b156101955783925f92889261107e60405196879586948594635c73957b60e11b865260048601612adb565b03915afa80156112d0576112bb575b508301610120848203126105b25780601f850112156105b257604051936110b661012086612677565b849061012081019283116105c157905b8282106112ab57505050825160038701540361129c5761012091604051966110ee8489612677565b600888526020880193601f190136853760015b600981101561115e578060051b8601515f19820182811161114a578a518110156111365760051b8a0160200152600101611101565b634e487b7160e01b89526032600452602489fd5b634e487b7160e01b89526011600452602489fd5b8685858b86600487818f60ff1916178155018251906001600160401b0382116112885768010000000000000000821161128857805482825580831061126d575b50908492918690885260208820885b838110611256575050505060407f2f25f350b681441b5bb45430fbe61e78d0170cff4ad02cb0aa48a6427909aa8c918151906111e8816127fe565b815260046020820152a26040519060208201906020835251809152604082019390855b81811061124057505050807f5da8009aeb11de7ef64598b0de897a03d2f93efabda8507a70e29cca4530cf5d9133940390a380f35b825186526020958601959092019160010161120b565b8251818301558795506020909201916001016111ad565b81885260208820611282918101908401612a0d565b8761119e565b634e487b7160e01b87526041600452602487fd5b630b6fac0360e41b8452600484fd5b81358152602091820191016110c6565b6112c89195505f90612677565b5f935f61108d565b6040513d5f823e3d90fd5b63e843c5eb60e01b5f5260045ffd5b506112fe60058a015460068b0154906129e8565b4210611040565b6307a92f1960e51b5f5260045ffd5b5061131e856127fe565b60048514611028565b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b61135e915060203d6020116107bb576107ad8183612677565b5f610fe9565b63cbf4a64560e01b5f5260045ffd5b34610195575f366003190112610195576003546040516001600160a01b039091168152602090f35b34610195576020366003190112610195576040516113b8816125ed565b5f81525f60208201526040516113cd81612609565b5f81525f602082015260408201525f6060820152606060808201525f60a08201525f60c08201525f60e08201525f6101008201525f6101208201525f6101408201525f610160820152606061018082015260405161142a81612624565b5f81525f60208201525f60408201525f60608201525f60808201525f60a08201525f60c08201525f60e08201525f6101008201526101a08201526101c06040519161147483612640565b5f83525f60208401525f60408401526060808401525f608084015201526004355f525f60205260405f206040516114aa816125ed565b81549160ff83166114ba816127fe565b825260089290921c6001600160a01b031660208201908152916114df60018201612698565b9160408101928352600382015460608201908152600483016040519081602082549182815201915f5260205f20905f5b8181106116c6575050508190036115269082612677565b60808301908152600584015460a08401908152600685015460c08501908152600786015460e0860190815260088701549061010087019182526009880154926101208801938452600a890154946101408901958652600b8a0154966101608a01978852600c8b01611596906126ee565b6101808b01908152986115ab600d8d0161278e565b9b6101a08c019c8d526012016115c09061281c565b9d6101c08c019e8f526040519d8e916020835261032083019d516115e3816127fe565b6020840152600160a01b60019003905116604083015251805160608301526020015190608001525160a08d0152519860c08c01610300905289518091526103408c0199602001905f5b8181106116b0575050905160e08c015250516101008a015251610120890152516101408801525161016087015251610180860152516101a085015251838203601f19016101c0850152929384936116ac93926116999161168c9190612878565b92516101e086019061289c565b51838203601f19016103008501526128ff565b0390f35b82518c5260209b8c019b9092019160010161162c565b825484526020909301926001928301920161150f565b34610195575f36600319011261019557602063ffffffff60025416604051908152f35b346101955761170d3661297d565b8115611364576003546040516350f3895360e11b81526004810184905260a09190911c63ffffffff16602482015260208160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156112d0575f9161182f575b5015611336575f8281526020819052604090208054600881901c6001600160a01b031680156113275733036118215760ff166117a1816127fe565b801515908161180c575b5061130557811580156117ff575b6117f057817fdbed30289c8edb9c8e40449ce45792cee5b9eba7c798afbadbcb4345699f1e249260076020930155604051908152a2005b632c45be6f60e01b5f5260045ffd5b50600881015482106117b9565b60039150611819816127fe565b1415846117ab565b6282b42960e81b5f5260045ffd5b611848915060203d6020116107bb576107ad8183612677565b83611766565b34610195575f36600319011261019557602060405160048152f35b34610195576020366003190112610195576004356001600160a01b03811690818103610195576003545f92835260016020908152604093849020549351636846b35160e11b815260a09290921c63ffffffff1660048301526001600160a01b039290921660248201526001600160401b0390921660448301528160648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af480156112d0575f9061016a57602090604051908152f35b34610195576020366003190112610195576004356001600160a01b03811690819003610195575f52600160205260206001600160401b0360405f205416604051908152f35b34610195576040366003190112610195576004356024356005811015610195578115611364576003546040516350f3895360e11b81526004810184905260a09190911c63ffffffff16602482015260208160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156112d0575f91611ae6575b5015611336576119e1816127fe565b600460ff82161161130557815f525f60205260405f2090815460018060a01b038160081c1680156113275733036118215760ff1690611a208183612b05565b15611305578383611a52837f2f25f350b681441b5bb45430fbe61e78d0170cff4ad02cb0aa48a6427909aa8c966129f5565b611a5b836127fe565b60018314611a89575b505060405191611a73816127fe565b8252611a7e816127fe565b6020820152604090a2005b602060067f0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46926005810154804210155f14611adc57611ac89042612aae565b9182915b0155604051908152a28385611a64565b505f918291611acc565b611aff915060203d6020116107bb576107ad8183612677565b836119d2565b346101955761022036600319011261019557600435600581101561019557602435604435606435610120366083190112610195576001600160401b036101a435116101955760a06101a43536036003190112610195576101c4356001600160401b03811161019557611b7b903690600401612950565b91909260406101e31936011261019557600354335f81815260016020908152604091829020549151636846b35160e11b815260a09490941c63ffffffff16600485015260248401929092526001600160401b0316604483015290959081908760648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49687156112d0575f976123ed575b505f8781526020819052604090205460081c6001600160a01b031633146123de5760ff611c2d612a30565b161580156123ca575b6123bb57611c42612a40565b60ff80611c4d612a30565b169116116123ac5761ffff610164351161239d5761012435610144351161238e5761016435610184351161237f5783156117f05760056101a43560040135101561019557611ca16101a435600401356127fe565b600560ff6101a43560040135161161237057611cc36101a435600401356127fe565b60046101a4350135600314158080612356575b61234757611cea6101a435600401356127fe565b1561231c5760246101a43501351561230d575b611d1360646101a435016101a435600401612a64565b9050156122fe57611d23886127fe565b600460ff89161180156122d7575b61130557156122d0575b4281106122c15742611d4d83836129e8565b11156122b257855f525f602052611d6860405f2097886129f5565b6005870155600686015560078501558354610100600160a81b0319163360081b610100600160a81b03161784556101e43560018501819055610204356002860181905560405163ec54ab4360e01b81526004818101879052909491939190611dd5906101a43501356127fe565b6101a43560040135602486015260843594851515860361019557851515604482015260a43594851515860361019557851515606483015260c43560ff811680910361019557608483015260e43560ff81168091036101955760a4830152610104359460ff86168087036101955760c48401526101243560e48401526101443561010484015261016435610124840152610184356101448401526101648301526101848201526020816101a48173__$4e0a7d2578f6e358dccb8beebb7c720d9f$__5af49081156112d0575f91612280575b5060038801556001600160401b038211610743578190611ec9600c8901546126b6565b601f811161224d575b505f90601f83116001146121e0575f926121d5575b50508160011b915f199060031b1c191617600c8601555b611f18600d860193849060ff801983541691151516179055565b82549162ff0000611f27612a30565b60101b169064ff0000000063ff000000611f3f612a40565b60181b169360201b169361ff0064ff000000001992151560081b169063ffffff0019161716171717905561012435600e83015561014435600f83015561016435601083015561018435601183015560128201611fa16101a435600401356127fe565b60ff6101a435600401351660ff1982541617905560246101a435013560138301556014820160018060a01b03611fdc60446101a43501612a50565b82546001600160a01b031916911617905560158201916120066101a4356064810190600401612a64565b6001600160401b0381959295116107435761202182546126b6565b601f81116121a5575b505f601f82116001146121425781929394955f92612137575b50508160011b915f199060031b1c19161790555b61207e61206960846101a43501612a23565b601683019060ff801983541691151516179055565b600a4391015560025463ffffffff811663ffffffff811461212357600163ffffffff9101169063ffffffff191617600255335f52600160205260405f20908154916001600160401b038316926001600160401b038414612123576001600160401b0360016020950116906001600160401b0319161790556040519033817fada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa5f80a38152f35b634e487b7160e01b5f52601160045260245ffd5b013590508580612043565b601f19821695835f5260205f20915f5b88811061218d57508360019596979810612174575b505050811b019055612057565b01355f19600384901b60f8161c19169055858080612167565b90926020600181928686013581550194019101612152565b6121cf90835f5260205f20601f840160051c81019160208510610cdd57601f0160051c0190612a0d565b8561202a565b013590508780611ee7565b909250600c88015f5260205f20905f935b601f1984168510612235576001945083601f1981161061221c575b505050811b01600c860155611efe565b01355f19600384901b60f8161c1916905587808061220c565b818101358355602094850194600190930192016121f1565b61227a90600c8a015f5260205f20601f850160051c81019160208610610cdd57601f0160051c0190612a0d565b88611ed2565b90506020813d6020116122aa575b8161229b60209383612677565b81010312610195575188611ea6565b3d915061228e565b637616640160e01b5f5260045ffd5b632ca4094f60e21b5f5260045ffd5b5042611d3b565b506122e1886127fe565b8715158015611d3157506122f4886127fe565b6003881415611d31565b630f8b932160e11b5f5260045ffd5b635e32eadd60e01b5f5260045ffd5b6001600160a01b036123336101a435604401612a50565b16611cfd5763562e597160e11b5f5260045ffd5b63f545b7bf60e01b5f5260045ffd5b50600161236860846101a43501612a23565b151514611cd6565b63f37f7b5d60e01b5f5260045ffd5b636d403ec760e11b5f5260045ffd5b63207ea56d60e01b5f5260045ffd5b634dba795160e01b5f5260045ffd5b632cbdc23160e01b5f5260045ffd5b63ac38930b60e01b5f5260045ffd5b50600860ff6123d7612a30565b1611611c36565b635040c4d360e11b5f5260045ffd5b9096506020813d602011612419575b8161240960209383612677565b8101031261019557519588611c02565b3d91506123fc565b34610195575f366003190112610195576002546040805163233ace1160e01b81529160209183916004918391901c6001600160a01b03165afa80156112d0575f9061016a57602090604051908152f35b34610195575f3660031901126101955760206040515f8152f35b34610195575f366003190112610195576002546040805191901c6001600160a01b03168152602090f35b34610195576020366003190112610195576004355f525f60205260405f2080549060ff821690600181016124e890612698565b9060038101546005820154600683015460078401546008850154600986015491600a87015493600b88015495600c8901612521906126ee565b9761252e600d8b0161278e565b9960120161253b9061281c565b9a604051809e8e829f61254d906127fe565b825260081c6001600160a01b0316602091820152815160408f0152015160608d015260808c015260a08b015260c08a015260e08901526101008801526101208701526101408601526101608501526102e061018085018190526125b39190850190612878565b906101a084016125c29161289c565b8281036102c08401526116ac916128ff565b34610195575f3660031901126101955780600560209252f35b6101e081019081106001600160401b0382111761074357604052565b604081019081106001600160401b0382111761074357604052565b61012081019081106001600160401b0382111761074357604052565b60a081019081106001600160401b0382111761074357604052565b61010081019081106001600160401b0382111761074357604052565b90601f801991011681019081106001600160401b0382111761074357604052565b906040516126a581612609565b602060018294805484520154910152565b90600182811c921680156126e4575b60208310146126d057565b634e487b7160e01b5f52602260045260245ffd5b91607f16916126c5565b9060405191825f825492612701846126b6565b808452936001811690811561276c5750600114612728575b5061272692500383612677565b565b90505f9291925260205f20905f915b818310612750575050906020612726928201015f612719565b6020919350806001915483858901015201910190918492612737565b90506020925061272694915060ff191682840152151560051b8201015f612719565b9060405161279b81612624565b6101006004829460ff815481811615158652818160081c1615156020870152818160101c166040870152818160181c16606087015260201c166080850152600181015460a0850152600281015460c0850152600381015460e08501520154910152565b6005111561280857565b634e487b7160e01b5f52602160045260245ffd5b9060405161282981612640565b608060ff600483958281541661283e816127fe565b85526001810154602086015260028101546001600160a01b03166040860152612869600382016126ee565b60608601520154161515910152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b610100809180511515845260208101511515602085015260ff604082015116604085015260ff606082015116606085015260ff608082015116608085015260a081015160a085015260c081015160c085015260e081015160e08501520151910152565b90815161290b816127fe565b81526020820151602082015260018060a01b036040830151166040820152608080612945606085015160a0606086015260a0850190612878565b930151151591015290565b9181601f84011215610195578235916001600160401b038311610195576020838186019501011161019557565b6040906003190112610195576004359060243590565b90606060031983011261019557600435916024356001600160401b03811161019557816129c291600401612950565b92909291604435906001600160401b038211610195576129e491600401612950565b9091565b9190820180921161212357565b906129ff816127fe565b60ff80198354169116179055565b818110612a18575050565b5f8155600101612a0d565b3580151581036101955790565b60c43560ff811681036101955790565b60e43560ff811681036101955790565b356001600160a01b03811681036101955790565b903590601e198136030182121561019557018035906001600160401b0382116101955760200191813603831361019557565b90816020910312610195575180151581036101955790565b9190820391821161212357565b908060209392818452848401375f828201840152601f01601f1916010190565b9290612af490612b029593604086526040860191612abb565b926020818503910152612abb565b90565b612b0e816127fe565b612b17826127fe565b808214612bc457612b27816127fe565b600281148015612bdd575b8015612bca575b612bc457612b46816127fe565b8015612ba557600390612b58816127fe565b14612b6257505f90565b612b6b816127fe565b8015908115612b90575b8115612b7f575090565b60019150612b8c816127fe565b1490565b9050612b9b816127fe565b6002811490612b75565b50612baf816127fe565b60038114908115612b90578115612b7f575090565b50505f90565b50612bd4816127fe565b60018114612b39565b50612be7816127fe565b60048114612b3256fea26469706673582212200a52f27184416b5ae98e0087b25714a990f6962f07c49629a75a0c2b8d22e53c64736f6c634300081c0033",
}

// ProcessRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProcessRegistryMetaData.ABI instead.
var ProcessRegistryABI = ProcessRegistryMetaData.ABI

// ProcessRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProcessRegistryMetaData.Bin instead.
var ProcessRegistryBin = ProcessRegistryMetaData.Bin

// DeployProcessRegistry deploys a new Ethereum contract, binding an instance of ProcessRegistry to it.
func DeployProcessRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID uint32, _stVerifier common.Address, _rVerifier common.Address, _blobsDA bool) (common.Address, *types.Transaction, *ProcessRegistry, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProcessRegistryBin), backend, _chainID, _stVerifier, _rVerifier, _blobsDA)
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

// BLOBINDEX is a free data retrieval call binding the contract method 0x26795bc3.
//
// Solidity: function BLOB_INDEX() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCaller) BLOBINDEX(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "BLOB_INDEX")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLOBINDEX is a free data retrieval call binding the contract method 0x26795bc3.
//
// Solidity: function BLOB_INDEX() view returns(uint8)
func (_ProcessRegistry *ProcessRegistrySession) BLOBINDEX() (uint8, error) {
	return _ProcessRegistry.Contract.BLOBINDEX(&_ProcessRegistry.CallOpts)
}

// BLOBINDEX is a free data retrieval call binding the contract method 0x26795bc3.
//
// Solidity: function BLOB_INDEX() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCallerSession) BLOBINDEX() (uint8, error) {
	return _ProcessRegistry.Contract.BLOBINDEX(&_ProcessRegistry.CallOpts)
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

// BlobsDA is a free data retrieval call binding the contract method 0xcdf497c8.
//
// Solidity: function blobsDA() view returns(bool)
func (_ProcessRegistry *ProcessRegistryCaller) BlobsDA(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "blobsDA")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlobsDA is a free data retrieval call binding the contract method 0xcdf497c8.
//
// Solidity: function blobsDA() view returns(bool)
func (_ProcessRegistry *ProcessRegistrySession) BlobsDA() (bool, error) {
	return _ProcessRegistry.Contract.BlobsDA(&_ProcessRegistry.CallOpts)
}

// BlobsDA is a free data retrieval call binding the contract method 0xcdf497c8.
//
// Solidity: function blobsDA() view returns(bool)
func (_ProcessRegistry *ProcessRegistryCallerSession) BlobsDA() (bool, error) {
	return _ProcessRegistry.Contract.BlobsDA(&_ProcessRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) ChainID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) ChainID() (uint32, error) {
	return _ProcessRegistry.Contract.ChainID(&_ProcessRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) ChainID() (uint32, error) {
	return _ProcessRegistry.Contract.ChainID(&_ProcessRegistry.CallOpts)
}

// GetNextProcessId is a free data retrieval call binding the contract method 0x68141f2c.
//
// Solidity: function getNextProcessId(address organizationId) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) GetNextProcessId(opts *bind.CallOpts, organizationId common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getNextProcessId", organizationId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextProcessId is a free data retrieval call binding the contract method 0x68141f2c.
//
// Solidity: function getNextProcessId(address organizationId) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) GetNextProcessId(organizationId common.Address) ([32]byte, error) {
	return _ProcessRegistry.Contract.GetNextProcessId(&_ProcessRegistry.CallOpts, organizationId)
}

// GetNextProcessId is a free data retrieval call binding the contract method 0x68141f2c.
//
// Solidity: function getNextProcessId(address organizationId) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetNextProcessId(organizationId common.Address) ([32]byte, error) {
	return _ProcessRegistry.Contract.GetNextProcessId(&_ProcessRegistry.CallOpts, organizationId)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool)))
func (_ProcessRegistry *ProcessRegistryCaller) GetProcess(opts *bind.CallOpts, processId [32]byte) (DAVINCITypesProcess, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcess", processId)

	if err != nil {
		return *new(DAVINCITypesProcess), err
	}

	out0 := *abi.ConvertType(out[0], new(DAVINCITypesProcess)).(*DAVINCITypesProcess)

	return out0, err

}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool)))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [32]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool)))
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcess(processId [32]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcessEndTime is a free data retrieval call binding the contract method 0xa5f01e19.
//
// Solidity: function getProcessEndTime(bytes32 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistryCaller) GetProcessEndTime(opts *bind.CallOpts, processId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcessEndTime", processId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProcessEndTime is a free data retrieval call binding the contract method 0xa5f01e19.
//
// Solidity: function getProcessEndTime(bytes32 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistrySession) GetProcessEndTime(processId [32]byte) (*big.Int, error) {
	return _ProcessRegistry.Contract.GetProcessEndTime(&_ProcessRegistry.CallOpts, processId)
}

// GetProcessEndTime is a free data retrieval call binding the contract method 0xa5f01e19.
//
// Solidity: function getProcessEndTime(bytes32 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcessEndTime(processId [32]byte) (*big.Int, error) {
	return _ProcessRegistry.Contract.GetProcessEndTime(&_ProcessRegistry.CallOpts, processId)
}

// GetRVerifierVKeyHash is a free data retrieval call binding the contract method 0xf9aa4499.
//
// Solidity: function getRVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) GetRVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getRVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRVerifierVKeyHash is a free data retrieval call binding the contract method 0xf9aa4499.
//
// Solidity: function getRVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) GetRVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetRVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// GetRVerifierVKeyHash is a free data retrieval call binding the contract method 0xf9aa4499.
//
// Solidity: function getRVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetRVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetRVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// GetSTVerifierVKeyHash is a free data retrieval call binding the contract method 0x4c0acc56.
//
// Solidity: function getSTVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) GetSTVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getSTVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSTVerifierVKeyHash is a free data retrieval call binding the contract method 0x4c0acc56.
//
// Solidity: function getSTVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) GetSTVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetSTVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// GetSTVerifierVKeyHash is a free data retrieval call binding the contract method 0x4c0acc56.
//
// Solidity: function getSTVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetSTVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetSTVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// PidPrefix is a free data retrieval call binding the contract method 0xcddf08bc.
//
// Solidity: function pidPrefix() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) PidPrefix(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "pidPrefix")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PidPrefix is a free data retrieval call binding the contract method 0xcddf08bc.
//
// Solidity: function pidPrefix() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) PidPrefix() (uint32, error) {
	return _ProcessRegistry.Contract.PidPrefix(&_ProcessRegistry.CallOpts)
}

// PidPrefix is a free data retrieval call binding the contract method 0xcddf08bc.
//
// Solidity: function pidPrefix() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) PidPrefix() (uint32, error) {
	return _ProcessRegistry.Contract.PidPrefix(&_ProcessRegistry.CallOpts)
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

// ProcessNonce is a free data retrieval call binding the contract method 0x62fa11fc.
//
// Solidity: function processNonce(address ) view returns(uint64)
func (_ProcessRegistry *ProcessRegistryCaller) ProcessNonce(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processNonce", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ProcessNonce is a free data retrieval call binding the contract method 0x62fa11fc.
//
// Solidity: function processNonce(address ) view returns(uint64)
func (_ProcessRegistry *ProcessRegistrySession) ProcessNonce(arg0 common.Address) (uint64, error) {
	return _ProcessRegistry.Contract.ProcessNonce(&_ProcessRegistry.CallOpts, arg0)
}

// ProcessNonce is a free data retrieval call binding the contract method 0x62fa11fc.
//
// Solidity: function processNonce(address ) view returns(uint64)
func (_ProcessRegistry *ProcessRegistryCallerSession) ProcessNonce(arg0 common.Address) (uint64, error) {
	return _ProcessRegistry.Contract.ProcessNonce(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census)
func (_ProcessRegistry *ProcessRegistryCaller) Processes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         DAVINCITypesEncryptionKey
	LatestStateRoot       *big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
}, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processes", arg0)

	outstruct := new(struct {
		Status                uint8
		OrganizationId        common.Address
		EncryptionKey         DAVINCITypesEncryptionKey
		LatestStateRoot       *big.Int
		StartTime             *big.Int
		Duration              *big.Int
		MaxVoters             *big.Int
		VotersCount           *big.Int
		OverwrittenVotesCount *big.Int
		CreationBlock         *big.Int
		BatchNumber           *big.Int
		MetadataURI           string
		BallotMode            DAVINCITypesBallotMode
		Census                DAVINCITypesCensus
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.OrganizationId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EncryptionKey = *abi.ConvertType(out[2], new(DAVINCITypesEncryptionKey)).(*DAVINCITypesEncryptionKey)
	outstruct.LatestStateRoot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MaxVoters = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VotersCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.OverwrittenVotesCount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.CreationBlock = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.BatchNumber = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.BallotMode = *abi.ConvertType(out[12], new(DAVINCITypesBallotMode)).(*DAVINCITypesBallotMode)
	outstruct.Census = *abi.ConvertType(out[13], new(DAVINCITypesCensus)).(*DAVINCITypesCensus)

	return *outstruct, err

}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census)
func (_ProcessRegistry *ProcessRegistrySession) Processes(arg0 [32]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         DAVINCITypesEncryptionKey
	LatestStateRoot       *big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census)
func (_ProcessRegistry *ProcessRegistryCallerSession) Processes(arg0 [32]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         DAVINCITypesEncryptionKey
	LatestStateRoot       *big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// RVerifier is a free data retrieval call binding the contract method 0x9c18deaf.
//
// Solidity: function rVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) RVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "rVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RVerifier is a free data retrieval call binding the contract method 0x9c18deaf.
//
// Solidity: function rVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) RVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.RVerifier(&_ProcessRegistry.CallOpts)
}

// RVerifier is a free data retrieval call binding the contract method 0x9c18deaf.
//
// Solidity: function rVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) RVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.RVerifier(&_ProcessRegistry.CallOpts)
}

// StVerifier is a free data retrieval call binding the contract method 0x09706392.
//
// Solidity: function stVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) StVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "stVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StVerifier is a free data retrieval call binding the contract method 0x09706392.
//
// Solidity: function stVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) StVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.StVerifier(&_ProcessRegistry.CallOpts)
}

// StVerifier is a free data retrieval call binding the contract method 0x09706392.
//
// Solidity: function stVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) StVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.StVerifier(&_ProcessRegistry.CallOpts)
}

// NewProcess is a paid mutator transaction binding the contract method 0x52016fe0.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey) returns(bytes32)
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey)
}

// NewProcess is a paid mutator transaction binding the contract method 0x52016fe0.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey) returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey)
}

// NewProcess is a paid mutator transaction binding the contract method 0x52016fe0.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey) returns(bytes32)
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0xaeaa8b0f.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,bytes32,address,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessCensus(opts *bind.TransactOpts, processId [32]byte, census DAVINCITypesCensus) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessCensus", processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0xaeaa8b0f.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,bytes32,address,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessCensus(processId [32]byte, census DAVINCITypesCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0xaeaa8b0f.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,bytes32,address,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessCensus(processId [32]byte, census DAVINCITypesCensus) (*types.Transaction, error) {
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

// SetProcessMaxVoters is a paid mutator transaction binding the contract method 0x74e07826.
//
// Solidity: function setProcessMaxVoters(bytes32 processId, uint256 _maxVoters) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessMaxVoters(opts *bind.TransactOpts, processId [32]byte, _maxVoters *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessMaxVoters", processId, _maxVoters)
}

// SetProcessMaxVoters is a paid mutator transaction binding the contract method 0x74e07826.
//
// Solidity: function setProcessMaxVoters(bytes32 processId, uint256 _maxVoters) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessMaxVoters(processId [32]byte, _maxVoters *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMaxVoters(&_ProcessRegistry.TransactOpts, processId, _maxVoters)
}

// SetProcessMaxVoters is a paid mutator transaction binding the contract method 0x74e07826.
//
// Solidity: function setProcessMaxVoters(bytes32 processId, uint256 _maxVoters) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessMaxVoters(processId [32]byte, _maxVoters *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMaxVoters(&_ProcessRegistry.TransactOpts, processId, _maxVoters)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xa5b110de.
//
// Solidity: function setProcessResults(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessResults(opts *bind.TransactOpts, processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessResults", processId, proof, input)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xa5b110de.
//
// Solidity: function setProcessResults(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessResults(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xa5b110de.
//
// Solidity: function setProcessResults(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessResults(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, proof, input)
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
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCensusUpdated is a free log retrieval operation binding the contract event 0x87e1c82bc035c552e997ec2f1cf023790ced0cc01098d516dc6e7200835ffb27.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI)
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

// WatchCensusUpdated is a free log subscription operation binding the contract event 0x87e1c82bc035c552e997ec2f1cf023790ced0cc01098d516dc6e7200835ffb27.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI)
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

// ParseCensusUpdated is a log parse operation binding the contract event 0x87e1c82bc035c552e997ec2f1cf023790ced0cc01098d516dc6e7200835ffb27.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI)
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

// ProcessRegistryProcessMaxVotersChangedIterator is returned from FilterProcessMaxVotersChanged and is used to iterate over the raw logs and unpacked data for ProcessMaxVotersChanged events raised by the ProcessRegistry contract.
type ProcessRegistryProcessMaxVotersChangedIterator struct {
	Event *ProcessRegistryProcessMaxVotersChanged // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessMaxVotersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessMaxVotersChanged)
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
		it.Event = new(ProcessRegistryProcessMaxVotersChanged)
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
func (it *ProcessRegistryProcessMaxVotersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessMaxVotersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessMaxVotersChanged represents a ProcessMaxVotersChanged event raised by the ProcessRegistry contract.
type ProcessRegistryProcessMaxVotersChanged struct {
	ProcessId [32]byte
	MaxVoters *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessMaxVotersChanged is a free log retrieval operation binding the contract event 0xdbed30289c8edb9c8e40449ce45792cee5b9eba7c798afbadbcb4345699f1e24.
//
// Solidity: event ProcessMaxVotersChanged(bytes32 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessMaxVotersChanged(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessMaxVotersChangedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessMaxVotersChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessMaxVotersChangedIterator{contract: _ProcessRegistry.contract, event: "ProcessMaxVotersChanged", logs: logs, sub: sub}, nil
}

// WatchProcessMaxVotersChanged is a free log subscription operation binding the contract event 0xdbed30289c8edb9c8e40449ce45792cee5b9eba7c798afbadbcb4345699f1e24.
//
// Solidity: event ProcessMaxVotersChanged(bytes32 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessMaxVotersChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessMaxVotersChanged, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessMaxVotersChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessMaxVotersChanged)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessMaxVotersChanged", log); err != nil {
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

// ParseProcessMaxVotersChanged is a log parse operation binding the contract event 0xdbed30289c8edb9c8e40449ce45792cee5b9eba7c798afbadbcb4345699f1e24.
//
// Solidity: event ProcessMaxVotersChanged(bytes32 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessMaxVotersChanged(log types.Log) (*ProcessRegistryProcessMaxVotersChanged, error) {
	event := new(ProcessRegistryProcessMaxVotersChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessMaxVotersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessResultsSetIterator is returned from FilterProcessResultsSet and is used to iterate over the raw logs and unpacked data for ProcessResultsSet events raised by the ProcessRegistry contract.
type ProcessRegistryProcessResultsSetIterator struct {
	Event *ProcessRegistryProcessResultsSet // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessResultsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessResultsSet)
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
		it.Event = new(ProcessRegistryProcessResultsSet)
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
func (it *ProcessRegistryProcessResultsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessResultsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessResultsSet represents a ProcessResultsSet event raised by the ProcessRegistry contract.
type ProcessRegistryProcessResultsSet struct {
	ProcessId [32]byte
	Sender    common.Address
	Result    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessResultsSet is a free log retrieval operation binding the contract event 0x5da8009aeb11de7ef64598b0de897a03d2f93efabda8507a70e29cca4530cf5d.
//
// Solidity: event ProcessResultsSet(bytes32 indexed processId, address indexed sender, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessResultsSet(opts *bind.FilterOpts, processId [][32]byte, sender []common.Address) (*ProcessRegistryProcessResultsSetIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessResultsSet", processIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessResultsSetIterator{contract: _ProcessRegistry.contract, event: "ProcessResultsSet", logs: logs, sub: sub}, nil
}

// WatchProcessResultsSet is a free log subscription operation binding the contract event 0x5da8009aeb11de7ef64598b0de897a03d2f93efabda8507a70e29cca4530cf5d.
//
// Solidity: event ProcessResultsSet(bytes32 indexed processId, address indexed sender, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessResultsSet(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessResultsSet, processId [][32]byte, sender []common.Address) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessResultsSet", processIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessResultsSet)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessResultsSet", log); err != nil {
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

// ParseProcessResultsSet is a log parse operation binding the contract event 0x5da8009aeb11de7ef64598b0de897a03d2f93efabda8507a70e29cca4530cf5d.
//
// Solidity: event ProcessResultsSet(bytes32 indexed processId, address indexed sender, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessResultsSet(log types.Log) (*ProcessRegistryProcessResultsSet, error) {
	event := new(ProcessRegistryProcessResultsSet)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessResultsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessStateTransitionedIterator is returned from FilterProcessStateTransitioned and is used to iterate over the raw logs and unpacked data for ProcessStateTransitioned events raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateTransitionedIterator struct {
	Event *ProcessRegistryProcessStateTransitioned // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessStateTransitionedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessStateTransitioned)
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
		it.Event = new(ProcessRegistryProcessStateTransitioned)
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
func (it *ProcessRegistryProcessStateTransitionedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessStateTransitionedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessStateTransitioned represents a ProcessStateTransitioned event raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateTransitioned struct {
	ProcessId                [32]byte
	Sender                   common.Address
	OldStateRoot             *big.Int
	NewStateRoot             *big.Int
	NewVotersCount           *big.Int
	NewOverwrittenVotesCount *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterProcessStateTransitioned is a free log retrieval operation binding the contract event 0x43b1a882c73bc84ae667608d3a61baececd99134224962eba9d13eb63c595ae1.
//
// Solidity: event ProcessStateTransitioned(bytes32 indexed processId, address indexed sender, uint256 oldStateRoot, uint256 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStateTransitioned(opts *bind.FilterOpts, processId [][32]byte, sender []common.Address) (*ProcessRegistryProcessStateTransitionedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessStateTransitioned", processIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessStateTransitionedIterator{contract: _ProcessRegistry.contract, event: "ProcessStateTransitioned", logs: logs, sub: sub}, nil
}

// WatchProcessStateTransitioned is a free log subscription operation binding the contract event 0x43b1a882c73bc84ae667608d3a61baececd99134224962eba9d13eb63c595ae1.
//
// Solidity: event ProcessStateTransitioned(bytes32 indexed processId, address indexed sender, uint256 oldStateRoot, uint256 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStateTransitioned(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStateTransitioned, processId [][32]byte, sender []common.Address) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessStateTransitioned", processIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessStateTransitioned)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateTransitioned", log); err != nil {
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

// ParseProcessStateTransitioned is a log parse operation binding the contract event 0x43b1a882c73bc84ae667608d3a61baececd99134224962eba9d13eb63c595ae1.
//
// Solidity: event ProcessStateTransitioned(bytes32 indexed processId, address indexed sender, uint256 oldStateRoot, uint256 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStateTransitioned(log types.Log) (*ProcessRegistryProcessStateTransitioned, error) {
	event := new(ProcessRegistryProcessStateTransitioned)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateTransitioned", log); err != nil {
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
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessStatusChanged is a free log retrieval operation binding the contract event 0x2f25f350b681441b5bb45430fbe61e78d0170cff4ad02cb0aa48a6427909aa8c.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 oldStatus, uint8 newStatus)
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

// WatchProcessStatusChanged is a free log subscription operation binding the contract event 0x2f25f350b681441b5bb45430fbe61e78d0170cff4ad02cb0aa48a6427909aa8c.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 oldStatus, uint8 newStatus)
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

// ParseProcessStatusChanged is a log parse operation binding the contract event 0x2f25f350b681441b5bb45430fbe61e78d0170cff4ad02cb0aa48a6427909aa8c.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 oldStatus, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStatusChanged(log types.Log) (*ProcessRegistryProcessStatusChanged, error) {
	event := new(ProcessRegistryProcessStatusChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
