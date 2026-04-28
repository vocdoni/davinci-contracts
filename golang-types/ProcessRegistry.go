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
	UniqueValues bool
	NumFields    uint8
	GroupSize    uint8
	CostExponent uint8
	MaxValue     *big.Int
	MinValue     *big.Int
	MaxValueSum  *big.Int
	MinValueSum  *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_stVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rVerifier\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_blobsDA\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"limbIndex\",\"type\":\"uint8\"}],\"name\":\"InvalidBlobCommitmentLimb\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOB_INDEX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CENSUS_ORIGIN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobsDA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461017d57608081612e1a803803809161001f8285610194565b83398101031261017d57610032816101cb565b9061003f602082016101dc565b606061004d604084016101dc565b92015180151580910361017d5760015f556003805460048054600160201b600160e01b0319909216604095861b600160401b600160e01b031617602088811b67ffffffff000000001691909117909355600164ff0000000160a01b03199091166001600160a01b039095169490941760c09290921b60ff60c01b16919091178355905163342e1df160e21b815263ffffffff909316918301919091523060248301528160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610189575f9161014b575b506004805463ffffffff60a01b191660a09290921b63ffffffff60a01b16919091179055604051612c2990816101f18239f35b90506020813d602011610181575b8161016660209383610194565b8101031261017d57610177906101cb565b5f610118565b5f80fd5b3d9150610159565b6040513d5f823e3d90fd5b601f909101601f19168101906001600160401b038211908210176101b757604052565b634e487b7160e01b5f52604160045260245ffd5b519063ffffffff8216820361017d57565b51906001600160a01b038216820361017d5756fe6080806040526004361015610012575f80fd5b5f905f3560e01c908163026cdee8146124645750806302cc919a1461244957806304ed00fa146122f0578063082b642e14612130578063097063921461210657806326795bc3146120ec57806335aeb98c1461182e5780634c0acc56146117de57806359d821c4146114cf57806362fa11fc1461148a57806368141f2c146113a85780636c7aff7f14610ff357806372c628ef14610fd8578063766422e014610bc5578063784df74a14610a97578063848df54014610a735780639b464994146108855780639c18deaf1461085c578063adc879e914610836578063c557fb4b146101fc578063cddf08bc146101d5578063cdf497c8146101af5763f9aa44991461011b575f80fd5b346101ac57806003193601126101ac576004805460405163233ace1160e01b8152929160209184919082906001600160a01b03165afa9081156101a05790610169575b602090604051908152f35b506020813d602011610198575b81610183602093836127a2565b81010312610194576020905161015e565b5f80fd5b3d9150610176565b604051903d90823e3d90fd5b80fd5b50346101ac57806003193601126101ac57602060ff60045460c01c166040519015158152f35b50346101ac57806003193601126101ac57602063ffffffff60045460a01c16604051908152f35b50346101ac5761020b366126d4565b610219949394929192612bd5565b60ff1984169384156108275760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156106975787916107f8575b50156107e9578386526001602052604086208054909390600881901c6001600160a01b0316156107da5760ff166102b6816125eb565b6107cb5760058401546102cd60068601548261298f565b4210156107bc5742106107ad576102e2612a64565b5081810192610100828503126105415783601f83011215610541576040519361030d610100866127a2565b8461010084019182116107955783905b8282106107855750505083519260208501519760408601519360608701519360808801519760a0810151988d60c08301519260e00151916040519a6103618c612786565b8b5260208b019e8f5260408b01998a5260608b0198895260808b0190815260a08b019b8c5260c08b0193845260e08b0192835260128d015460ff166103a5816125eb565b6003036107695760148d0154905160405163650e5fcf60e01b8152600481019190915290602090829060249082906001600160a01b03165afa9182156101a05791610737575b5060ff60168d015416159081610729575b811561071f575b50610710575b88519760038c019889540361070157518751610424916129b4565b998a1515806106ef575b6106e057908e96959493929160ff60045460c01c16610554575b505060035460401c6001600160a01b0316919050813b1561055057859361048560405196879586948594635c73957b60e11b865260048601612ac0565b03915afa80156105455761052c575b5050600b91875190556104ac6008860194855461298f565b80945551936104c06009820195865461298f565b8095550180545f1981146105185760010190555193516040519485526020850152604084015260608301527f7156cf1102d86d870b69eed151febab016e34063d74c714b17538efccfd9d01960803393a36001815580f35b634e487b7160e01b88526011600452602488fd5b81610536916127a2565b61054157875f610494565b8780fd5b6040513d84823e3d90fd5b8580fd5b9091929394959650519151905190604051926105716060856127a2565b60308452602084019060403683378060801c6106ca5760801b90528060801c6106b65760801b60308301528060801c6106a257908d959493929160801b60408201526105ef602073__$176218451d0b397034f549dcd265d6db5e$__926040518093819263da23b04d60e01b83528460048401526024830190612609565b0381855af4908115610697578791610662575b50813b1561065e578690602460405180948193634e733ef560e11b835260048301525af490811561065357869161063a575b80610448565b81610644916127a2565b61064f57845f610634565b8480fd5b6040513d88823e3d90fd5b8680fd5b9650506020863d60201161068f575b8161067e602093836127a2565b81010312610194578d95515f610602565b3d9150610671565b6040513d89823e3d90fd5b6316416a3d60e01b8e52600260045260248efd5b6316416a3d60e01b8f52600160045260248ffd5b50506316416a3d60e01b8f5260048f905260248ffd5b6357d18d5360e11b8f5260048ffd5b5060088c015460078d0154111561042e565b630b6fac0360e41b8f5260048ffd5b635e32eadd60e01b8e5260048efd5b905043105f610403565b600a8d0154811091506103fc565b90506020813d602011610761575b81610752602093836127a2565b8101031261019457515f6103eb565b3d9150610745565b90505160138c01541461040957635e32eadd60e01b8e5260048efd5b813581526020918201910161031d565b8980fd5b634e487b7160e01b5f52604160045260245ffd5b63e843c5eb60e01b8752600487fd5b63e843c5eb60e01b8852600488fd5b6307a92f1960e51b8752600487fd5b634d36eb6960e01b8852600488fd5b632299770d60e11b8652600486fd5b61081a915060203d602011610820575b61081281836127a2565b810190612977565b5f610280565b503d610808565b63cbf4a64560e01b8752600487fd5b50346101ac57806003193601126101ac57602063ffffffff600354821c16604051908152f35b50346101ac57806003193601126101ac576004546040516001600160a01b039091168152602090f35b50346101ac5760403660031901126101ac5761089f6125ad565b60243560ff198216918215610a645760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610a59578491610a3a575b5015610a2b578183526001602052604083208054600881901c6001600160a01b03168015610a1c573303610a0e5760ff16610943816125eb565b80151590816109f9575b506109ea576006600582015491019081549083159182156109d5575b82156109b6575b50506109a757817f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba9260209255604051908152a280f35b637616640160e01b8452600484fd5b8192506109c6856109cc9361298f565b9261298f565b10155f80610970565b91506109e1848261298f565b42101591610969565b6307a92f1960e51b8452600484fd5b60039150610a06816125eb565b14155f61094d565b6282b42960e81b8552600485fd5b634d36eb6960e01b8652600486fd5b632299770d60e11b8352600483fd5b610a53915060203d6020116108205761081281836127a2565b5f610909565b6040513d86823e3d90fd5b63cbf4a64560e01b8452600484fd5b50346101ac57806003193601126101ac57602063ffffffff60035416604051908152f35b50346101ac5760203660031901126101ac5760ff19610ab46125ad565b16815260016020526040902080549060ff82169060018101610ad5906127c3565b9060038101546005820154600683015460078401546008850154600986015491600a87015493600b88015495600c8901610b0e90612819565b97610b1b600d8b016128b9565b99601201610b289061291b565b9a604051809e8e829f610b3a906125eb565b825260081c6001600160a01b0316602091820152815160408f0152015160608d015260808c015260a08b015260c08a015260e08901526101008801526101208701526101408601526101608501526102c06101808501819052610ba09190850190612609565b906101a08401610baf9161262d565b8281036102a0840152610bc191612683565b0390f35b503461019457610bd4366126d4565b949293610be2929192612bd5565b60ff198116908115610fc95760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610f35575f91610faa575b5015610f9b57805f52600160205260405f209586549560018060a01b038760081c1615610f8c5760ff871693610c7e856125eb565b600285148015610f79575b610f6a57610c96856125eb565b600185141580610f4f575b610f40576004546001600160a01b031691823b156101945783925f928892610cdf60405196879586948594635c73957b60e11b865260048601612ac0565b03915afa8015610f3557610f20575b5083016101208482031261064f5780601f8501121561064f5760405193610d17610120866127a2565b8490610120810192831161065e57905b828210610f10575050508251600387015403610f01576101209160405196610d4f84896127a2565b600888526020880193601f190136853760015b6009811015610dbf578060051b8601515f198201828111610dab578a51811015610d975760051b8a0160200152600101610d62565b634e487b7160e01b89526032600452602489fd5b634e487b7160e01b89526011600452602489fd5b8685858b86600487818f60ff1916178155018251906001600160401b038211610eed57680100000000000000008211610eed578054828255808310610ed2575b50908492918690885260208820885b838110610ebb575050505060407f56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf91815190610e49816125eb565b815260046020820152a26040519060208201906020835251809152604082019390855b818110610ea557505050807fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e9133940390a36001815580f35b8251865260209586019590920191600101610e6c565b825181830155879550602090920191600101610e0e565b81885260208820610ee79181019084016129db565b87610dff565b634e487b7160e01b87526041600452602487fd5b630b6fac0360e41b8452600484fd5b8135815260209182019101610d27565b610f2d9195505f906127a2565b5f935f610cee565b6040513d5f823e3d90fd5b63e843c5eb60e01b5f5260045ffd5b50610f6360058a015460068b01549061298f565b4210610ca1565b6307a92f1960e51b5f5260045ffd5b50610f83856125eb565b60048514610c89565b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b610fc3915060203d6020116108205761081281836127a2565b5f610c49565b63cbf4a64560e01b5f5260045ffd5b34610194575f36600319011261019457602060405160048152f35b346101945760403660031901126101945761100c6125ad565b6024356001600160401b038111610194578060040160a060031983360301126101945760ff198316928315610fc95760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610f35575f91611389575b5015610f9b575f83815260016020526040902080549190600883901c6001600160a01b03168015610f8c57330361137b5760ff601282015416926110d9846125eb565b6002840361136c5782356005811015610194576110f5816125eb565b6110fe856125eb565b840361135d57602485013594851561134e57606481019461111f8686612a32565b90501561133f57600390611132816125eb565b149081611321575b506113125760ff1661114b816125eb565b80151590816112fd575b50610f6a5761116d600582015460068301549061298f565b421015610f405783601382015560156111868484612a32565b91909201916001600160401b038211610799576111a383546127e1565b601f81116112c2575b505f90601f831160011461123357928261122396937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2989693611208965f92611228575b50508160011b915f199060031b1c1916179055612a32565b92906040519384938452604060208501526040840191612aa0565b0390a2005b013590508a806111f0565b601f19831691845f5260205f20925f5b8181106112aa5750937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed29896936112089693600193836112239b9810611291575b505050811b019055612a32565b01355f19600384901b60f8161c191690558a8080611284565b91936020600181928787013581550195019201611243565b6112ed90845f5260205f20601f850160051c810191602086106112f3575b601f0160051c01906129db565b876111ac565b90915081906112e0565b6003915061130a816125eb565b141586611155565b63562e597160e11b5f5260045ffd5b6001600160a01b03915061133790604401612a1e565b16158761113a565b630f8b932160e11b5f5260045ffd5b635e32eadd60e01b5f5260045ffd5b63f37f7b5d60e01b5f5260045ffd5b63050b77c760e21b5f5260045ffd5b6282b42960e81b5f5260045ffd5b6113a2915060203d6020116108205761081281836127a2565b84611096565b34610194576020366003190112610194576004356001600160a01b0381169081810361019457600480545f93845260026020908152604094859020549451636846b35160e11b815260a09290921c63ffffffff16928201929092526001600160a01b039290921660248301526001600160401b039092166044820152908160648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af48015610f35576020915f9161145d575b506040519060ff19168152f35b61147d9150823d8411611483575b61147581836127a2565b8101906129c1565b82611450565b503d61146b565b34610194576020366003190112610194576004356001600160a01b03811690819003610194575f52600260205260206001600160401b0360405f205416604051908152f35b34610194576020366003190112610194576114e86125ad565b6040516114f481612734565b5f81525f602082015260405161150981612750565b5f81525f602082015260408201525f6060820152606060808201525f60a08201525f60c08201525f60e08201525f6101008201525f6101208201525f6101408201525f6101608201526060610180820152611562612a64565b6101a08201526101c0604051916115788361276b565b5f83525f60208401525f60408401526060808401525f6080840152015260ff19165f52600160205260405f206040516115b081612734565b81549160ff83166115c0816125eb565b825260089290921c6001600160a01b031660208201908152916115e5600182016127c3565b9160408101928352600382015460608201908152600483016040519081602082549182815201915f5260205f20905f5b8181106117c85750505081900361162c90826127a2565b60808301908152600584015460a08401908152600685015460c08501908152600786015460e0860190815260088701549061010087019182526009880154926101208801938452600a890154946101408901958652600b8a0154966101608a01978852600c8b0161169c90612819565b6101808b01908152986116b1600d8d016128b9565b9b6101a08c019c8d526012016116c69061291b565b9d6101c08c019e8f526040519d8e916020835261030083019d516116e9816125eb565b6020840152600160a01b60019003905116604083015251805160608301526020015190608001525160a08d0152519860c08c016102e0905289518091526103208c0199602001905f5b8181106117b2575050905160e08c015250516101008a015251610120890152516101408801525161016087015251610180860152516101a085015251838203601f19016101c085015292938493610bc1939261179f916117929190612609565b92516101e086019061262d565b51838203601f19016102e0850152612683565b82518c5260209b8c019b90920191600101611732565b8254845260209093019260019283019201611615565b34610194575f366003190112610194576003546040805163233ace1160e01b81529160209183916004918391901c6001600160a01b03165afa8015610f35575f9061016957602090604051908152f35b3461019457610200366003190112610194576005600435101561019457602435604435610100366083190112610194576001600160401b0361018435116101945760a06101843536036003190112610194576101a4356001600160401b038111610194576118a09036906004016125be565b6040366101c31901126101945760048054335f81815260026020908152604091829020549151636846b35160e11b815260a09490941c63ffffffff169484019490945260248301919091526001600160401b03166044820152939085908560648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4948515610f35575f956120cb575b5060ff1985165f9081526001602052604090205460081c6001600160a01b031633146120bc5760ff6119566129fe565b161580156120a8575b6120995761196b612a0e565b60ff806119766129fe565b1691161161208a5761010435610124351161207b5761014435610164351161206c576064351561205d57600561018435600401351015610194576119c061018435600401356125eb565b600560ff6101843560040135161161135d576119e261018435600401356125eb565b6004610184350135600314158080612048575b61203957611a0961018435600401356125eb565b1561200e5760246101843501351561134e575b611a326064610184350161018435600401612a32565b90501561133f57611a446004356125eb565b600460ff813516118015611fdf575b610f6a5715611fd7575b428510611fc85742611a6f828761298f565b1115611fb95760ff1984165f52600160205260405f2094611a926004358761299c565b6005860155600685015560643560078501558354610100600160a81b0319163360081b610100600160a81b03161784556101c435600185018190556101e43560028601819055604051631044f8bb60e11b815260ff198616600482810191909152929492909190611b08906101843501356125eb565b610184356004013560248301526084351515806084350361019457604483015260a43560ff811680910361019457606483015260c43560ff811680910361019457608483015260e4359460ff86168087036101945760a48401526101043560c48401526101243560e484015261014435610104840152610164356101248401526101448301526101648201526020816101848173__$4e0a7d2578f6e358dccb8beebb7c720d9f$__5af4908115610f35575f91611f87575b5060038601556001600160401b038211610799578190611be3600c8701546127e1565b601f8111611f54575b505f90601f8311600114611ee7575f92611edc575b50508160011b915f199060031b1c191617600c8401555b600d8301805460ff191660ff60843515151617815590611c366129fe565b90825461ff0063ff00000062ff0000611c4d612a0e565b60101b169360181b169360081b169063ffffff001916171717905561010435600e83015561012435600f83015561014435601083015561016435601183015560128201611ca061018435600401356125eb565b60ff61018435600401351660ff19825416179055602461018435013560138301556014820160018060a01b03611cdb60446101843501612a1e565b82546001600160a01b03191691161790556015820191611d05610184356064810190600401612a32565b6001600160401b03819592951161079957611d2082546127e1565b601f8111611eac575b505f601f8211600114611e495781929394955f92611e3e575b50508160011b915f199060031b1c19161790555b611d7d611d68608461018435016129f1565b601683019060ff801983541691151516179055565b600a4391015560035463ffffffff811663ffffffff8114611e2a57600163ffffffff9101169063ffffffff191617600355335f52600260205260405f20908154916001600160401b038316926001600160401b038414611e2a576001600160401b0360016020950116906001600160401b031916179055604051903360ff1982167feefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d15f80a360ff19168152f35b634e487b7160e01b5f52601160045260245ffd5b013590508580611d42565b601f19821695835f5260205f20915f5b888110611e9457508360019596979810611e7b575b505050811b019055611d56565b01355f19600384901b60f8161c19169055858080611e6e565b90926020600181928686013581550194019101611e59565b611ed690835f5260205f20601f840160051c810191602085106112f357601f0160051c01906129db565b85611d29565b013590508580611c01565b909250600c86015f5260205f20905f935b601f1984168510611f3c576001945083601f19811610611f23575b505050811b01600c840155611c18565b01355f19600384901b60f8161c19169055858080611f13565b81810135835560209485019460019093019201611ef8565b611f8190600c88015f5260205f20601f850160051c810191602086106112f357601f0160051c01906129db565b86611bec565b90506020813d602011611fb1575b81611fa2602093836127a2565b81010312610194575186611bc0565b3d9150611f95565b637616640160e01b5f5260045ffd5b632ca4094f60e21b5f5260045ffd5b429450611a5d565b50611feb6004356125eb565b60043515158015611a5357506120026004356125eb565b60036004351415611a53565b6001600160a01b0361202561018435604401612a1e565b16611a1c5763562e597160e11b5f5260045ffd5b63f545b7bf60e01b5f5260045ffd5b50612058608461018435016129f1565b6119f5565b632c45be6f60e01b5f5260045ffd5b636d403ec760e11b5f5260045ffd5b63207ea56d60e01b5f5260045ffd5b632cbdc23160e01b5f5260045ffd5b63ac38930b60e01b5f5260045ffd5b50600860ff6120b56129fe565b161161195f565b635040c4d360e11b5f5260045ffd5b6120e591955060203d6020116114835761147581836127a2565b9386611926565b34610194575f3660031901126101945760206040515f8152f35b34610194575f366003190112610194576003546040805191901c6001600160a01b03168152602090f35b34610194576040366003190112610194576121496125ad565b60243560058110156101945760ff198216918215610fc95760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610f35575f916122d1575b5015610f9b576121cb816125eb565b600460ff821611610f6a57815f52600160205260405f2090815460018060a01b038160081c168015610f8c57330361137b5760ff169061220b8183612aea565b15610f6a57838361223d837f56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf9661299c565b612246836125eb565b60018314612274575b50506040519161225e816125eb565b8252612269816125eb565b6020820152604090a2005b602060067f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba926005810154804210155f146122c7576122b390426129b4565b9182915b0155604051908152a2838561224f565b505f9182916122b7565b6122ea915060203d6020116108205761081281836127a2565b836121bc565b346101945760203660031901126101945760ff1961230c6125ad565b165f52600160205260405f2060405161232481612734565b815460ff8116612333816125eb565b825260081c6001600160a01b03166020820152612352600183016127c3565b604082015260038201546060820152600482016040519081602082549182815201915f5260205f20905f5b81811061243357602061242b888888612398818a03826127a2565b608082015260058201549060a081019182526101c061241f601260068601549560c08501968752600781015460e086015260088101546101008601526009810154610120860152600a810154610140860152600b810154610160860152612401600c8201612819565b610180860152612413600d82016128b9565b6101a08601520161291b565b9101525190519061298f565b604051908152f35b825484526020909301926001928301920161237d565b34610194575f36600319011261019457602060405160058152f35b346101945760403660031901126101945761247d6125ad565b60ff19811691602435918315610fc95760048054632145bb8f60e01b845260ff199092169083015260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610f35575f9161258e575b5015610f9b575f8281526001602052604090208054600881901c6001600160a01b03168015610f8c57330361137b5760ff1661251d816125eb565b8015159081612579575b50610f6a578115801561256c575b61205d57817f36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be9260076020930155604051908152a2005b5060088101548210612535565b60039150612586816125eb565b141584612527565b6125a7915060203d6020116108205761081281836127a2565b836124e2565b6004359060ff198216820361019457565b9181601f84011215610194578235916001600160401b038311610194576020838186019501011161019457565b600511156125f557565b634e487b7160e01b5f52602160045260245ffd5b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60e0809180511515845260ff602082015116602085015260ff604082015116604085015260ff60608201511660608501526080810151608085015260a081015160a085015260c081015160c08501520151910152565b90815161268f816125eb565b81526020820151602082015260018060a01b0360408301511660408201526080806126c9606085015160a0606086015260a0850190612609565b930151151591015290565b9060606003198301126101945760043560ff198116810361019457916024356001600160401b038111610194578161270e916004016125be565b92909291604435906001600160401b03821161019457612730916004016125be565b9091565b6101e081019081106001600160401b0382111761079957604052565b604081019081106001600160401b0382111761079957604052565b60a081019081106001600160401b0382111761079957604052565b61010081019081106001600160401b0382111761079957604052565b90601f801991011681019081106001600160401b0382111761079957604052565b906040516127d081612750565b602060018294805484520154910152565b90600182811c9216801561280f575b60208310146127fb57565b634e487b7160e01b5f52602260045260245ffd5b91607f16916127f0565b9060405191825f82549261282c846127e1565b80845293600181169081156128975750600114612853575b50612851925003836127a2565b565b90505f9291925260205f20905f915b81831061287b575050906020612851928201015f612844565b6020919350806001915483858901015201910190918492612862565b90506020925061285194915060ff191682840152151560051b8201015f612844565b906040516128c681612786565b60e06004829460ff815481811615158652818160081c166020870152818160101c16604087015260181c16606085015260018101546080850152600281015460a0850152600381015460c08501520154910152565b906040516129288161276b565b608060ff600483958281541661293d816125eb565b85526001810154602086015260028101546001600160a01b0316604086015261296860038201612819565b60608601520154161515910152565b90816020910312610194575180151581036101945790565b91908201809211611e2a57565b906129a6816125eb565b60ff80198354169116179055565b91908203918211611e2a57565b90816020910312610194575160ff19811681036101945790565b8181106129e6575050565b5f81556001016129db565b3580151581036101945790565b60a43560ff811681036101945790565b60c43560ff811681036101945790565b356001600160a01b03811681036101945790565b903590601e198136030182121561019457018035906001600160401b0382116101945760200191813603831361019457565b60405190612a7182612786565b5f60e0838281528260208201528260408201528260608201528260808201528260a08201528260c08201520152565b908060209392818452848401375f828201840152601f01601f1916010190565b9290612ad990612ae79593604086526040860191612aa0565b926020818503910152612aa0565b90565b612af3816125eb565b612afc826125eb565b808214612ba957612b0c816125eb565b600281148015612bc2575b8015612baf575b612ba957612b2b816125eb565b8015612b8a57600390612b3d816125eb565b14612b4757505f90565b612b50816125eb565b8015908115612b75575b8115612b64575090565b60019150612b71816125eb565b1490565b9050612b80816125eb565b6002811490612b5a565b50612b94816125eb565b60038114908115612b75578115612b64575090565b50505f90565b50612bb9816125eb565b60018114612b1e565b50612bcc816125eb565b60048114612b17565b60025f5414612be45760025f55565b633ee5aeb560e01b5f5260045ffdfea2646970667358221220186896916b4b94a5e24809cbcd4fa86aa4938b2dea8b1a8b906a504891b0823464736f6c634300081c0033",
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
// Solidity: function getNextProcessId(address organizationId) view returns(bytes31)
func (_ProcessRegistry *ProcessRegistryCaller) GetNextProcessId(opts *bind.CallOpts, organizationId common.Address) ([31]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getNextProcessId", organizationId)

	if err != nil {
		return *new([31]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([31]byte)).(*[31]byte)

	return out0, err

}

// GetNextProcessId is a free data retrieval call binding the contract method 0x68141f2c.
//
// Solidity: function getNextProcessId(address organizationId) view returns(bytes31)
func (_ProcessRegistry *ProcessRegistrySession) GetNextProcessId(organizationId common.Address) ([31]byte, error) {
	return _ProcessRegistry.Contract.GetNextProcessId(&_ProcessRegistry.CallOpts, organizationId)
}

// GetNextProcessId is a free data retrieval call binding the contract method 0x68141f2c.
//
// Solidity: function getNextProcessId(address organizationId) view returns(bytes31)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetNextProcessId(organizationId common.Address) ([31]byte, error) {
	return _ProcessRegistry.Contract.GetNextProcessId(&_ProcessRegistry.CallOpts, organizationId)
}

// GetProcess is a free data retrieval call binding the contract method 0x59d821c4.
//
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool)))
func (_ProcessRegistry *ProcessRegistryCaller) GetProcess(opts *bind.CallOpts, processId [31]byte) (DAVINCITypesProcess, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcess", processId)

	if err != nil {
		return *new(DAVINCITypesProcess), err
	}

	out0 := *abi.ConvertType(out[0], new(DAVINCITypesProcess)).(*DAVINCITypesProcess)

	return out0, err

}

// GetProcess is a free data retrieval call binding the contract method 0x59d821c4.
//
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool)))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [31]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x59d821c4.
//
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool)))
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcess(processId [31]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcessEndTime is a free data retrieval call binding the contract method 0x04ed00fa.
//
// Solidity: function getProcessEndTime(bytes31 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistryCaller) GetProcessEndTime(opts *bind.CallOpts, processId [31]byte) (*big.Int, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcessEndTime", processId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProcessEndTime is a free data retrieval call binding the contract method 0x04ed00fa.
//
// Solidity: function getProcessEndTime(bytes31 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistrySession) GetProcessEndTime(processId [31]byte) (*big.Int, error) {
	return _ProcessRegistry.Contract.GetProcessEndTime(&_ProcessRegistry.CallOpts, processId)
}

// GetProcessEndTime is a free data retrieval call binding the contract method 0x04ed00fa.
//
// Solidity: function getProcessEndTime(bytes31 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcessEndTime(processId [31]byte) (*big.Int, error) {
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

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census)
func (_ProcessRegistry *ProcessRegistryCaller) Processes(opts *bind.CallOpts, arg0 [31]byte) (struct {
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

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census)
func (_ProcessRegistry *ProcessRegistrySession) Processes(arg0 [31]byte) (struct {
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

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census)
func (_ProcessRegistry *ProcessRegistryCallerSession) Processes(arg0 [31]byte) (struct {
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

// NewProcess is a paid mutator transaction binding the contract method 0x35aeb98c.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey)
}

// NewProcess is a paid mutator transaction binding the contract method 0x35aeb98c.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey) returns(bytes31)
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey)
}

// NewProcess is a paid mutator transaction binding the contract method 0x35aeb98c.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6c7aff7f.
//
// Solidity: function setProcessCensus(bytes31 processId, (uint8,bytes32,address,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessCensus(opts *bind.TransactOpts, processId [31]byte, census DAVINCITypesCensus) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessCensus", processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6c7aff7f.
//
// Solidity: function setProcessCensus(bytes31 processId, (uint8,bytes32,address,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessCensus(processId [31]byte, census DAVINCITypesCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6c7aff7f.
//
// Solidity: function setProcessCensus(bytes31 processId, (uint8,bytes32,address,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessCensus(processId [31]byte, census DAVINCITypesCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0x9b464994.
//
// Solidity: function setProcessDuration(bytes31 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessDuration(opts *bind.TransactOpts, processId [31]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessDuration", processId, _duration)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0x9b464994.
//
// Solidity: function setProcessDuration(bytes31 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessDuration(processId [31]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessDuration(&_ProcessRegistry.TransactOpts, processId, _duration)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0x9b464994.
//
// Solidity: function setProcessDuration(bytes31 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessDuration(processId [31]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessDuration(&_ProcessRegistry.TransactOpts, processId, _duration)
}

// SetProcessMaxVoters is a paid mutator transaction binding the contract method 0x026cdee8.
//
// Solidity: function setProcessMaxVoters(bytes31 processId, uint256 _maxVoters) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessMaxVoters(opts *bind.TransactOpts, processId [31]byte, _maxVoters *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessMaxVoters", processId, _maxVoters)
}

// SetProcessMaxVoters is a paid mutator transaction binding the contract method 0x026cdee8.
//
// Solidity: function setProcessMaxVoters(bytes31 processId, uint256 _maxVoters) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessMaxVoters(processId [31]byte, _maxVoters *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMaxVoters(&_ProcessRegistry.TransactOpts, processId, _maxVoters)
}

// SetProcessMaxVoters is a paid mutator transaction binding the contract method 0x026cdee8.
//
// Solidity: function setProcessMaxVoters(bytes31 processId, uint256 _maxVoters) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessMaxVoters(processId [31]byte, _maxVoters *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMaxVoters(&_ProcessRegistry.TransactOpts, processId, _maxVoters)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0x766422e0.
//
// Solidity: function setProcessResults(bytes31 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessResults(opts *bind.TransactOpts, processId [31]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessResults", processId, proof, input)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0x766422e0.
//
// Solidity: function setProcessResults(bytes31 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessResults(processId [31]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0x766422e0.
//
// Solidity: function setProcessResults(bytes31 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessResults(processId [31]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x082b642e.
//
// Solidity: function setProcessStatus(bytes31 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessStatus(opts *bind.TransactOpts, processId [31]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessStatus", processId, newStatus)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x082b642e.
//
// Solidity: function setProcessStatus(bytes31 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessStatus(processId [31]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessStatus(&_ProcessRegistry.TransactOpts, processId, newStatus)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x082b642e.
//
// Solidity: function setProcessStatus(bytes31 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessStatus(processId [31]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessStatus(&_ProcessRegistry.TransactOpts, processId, newStatus)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xc557fb4b.
//
// Solidity: function submitStateTransition(bytes31 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SubmitStateTransition(opts *bind.TransactOpts, processId [31]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "submitStateTransition", processId, proof, input)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xc557fb4b.
//
// Solidity: function submitStateTransition(bytes31 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistrySession) SubmitStateTransition(processId [31]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransition(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xc557fb4b.
//
// Solidity: function submitStateTransition(bytes31 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SubmitStateTransition(processId [31]byte, proof []byte, input []byte) (*types.Transaction, error) {
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
	ProcessId  [31]byte
	CensusRoot [32]byte
	CensusURI  string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCensusUpdated is a free log retrieval operation binding the contract event 0x660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2.
//
// Solidity: event CensusUpdated(bytes31 indexed processId, bytes32 censusRoot, string censusURI)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterCensusUpdated(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryCensusUpdatedIterator, error) {

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

// WatchCensusUpdated is a free log subscription operation binding the contract event 0x660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2.
//
// Solidity: event CensusUpdated(bytes31 indexed processId, bytes32 censusRoot, string censusURI)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchCensusUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryCensusUpdated, processId [][31]byte) (event.Subscription, error) {

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

// ParseCensusUpdated is a log parse operation binding the contract event 0x660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2.
//
// Solidity: event CensusUpdated(bytes31 indexed processId, bytes32 censusRoot, string censusURI)
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
	ProcessId [31]byte
	Creator   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessCreated is a free log retrieval operation binding the contract event 0xeefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d1.
//
// Solidity: event ProcessCreated(bytes31 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessCreated(opts *bind.FilterOpts, processId [][31]byte, creator []common.Address) (*ProcessRegistryProcessCreatedIterator, error) {

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

// WatchProcessCreated is a free log subscription operation binding the contract event 0xeefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d1.
//
// Solidity: event ProcessCreated(bytes31 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessCreated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessCreated, processId [][31]byte, creator []common.Address) (event.Subscription, error) {

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

// ParseProcessCreated is a log parse operation binding the contract event 0xeefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d1.
//
// Solidity: event ProcessCreated(bytes31 indexed processId, address indexed creator)
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
	ProcessId [31]byte
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessDurationChanged is a free log retrieval operation binding the contract event 0x45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba.
//
// Solidity: event ProcessDurationChanged(bytes31 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessDurationChanged(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessDurationChangedIterator, error) {

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

// WatchProcessDurationChanged is a free log subscription operation binding the contract event 0x45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba.
//
// Solidity: event ProcessDurationChanged(bytes31 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessDurationChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessDurationChanged, processId [][31]byte) (event.Subscription, error) {

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

// ParseProcessDurationChanged is a log parse operation binding the contract event 0x45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba.
//
// Solidity: event ProcessDurationChanged(bytes31 indexed processId, uint256 duration)
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
	ProcessId [31]byte
	MaxVoters *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessMaxVotersChanged is a free log retrieval operation binding the contract event 0x36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be.
//
// Solidity: event ProcessMaxVotersChanged(bytes31 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessMaxVotersChanged(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessMaxVotersChangedIterator, error) {

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

// WatchProcessMaxVotersChanged is a free log subscription operation binding the contract event 0x36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be.
//
// Solidity: event ProcessMaxVotersChanged(bytes31 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessMaxVotersChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessMaxVotersChanged, processId [][31]byte) (event.Subscription, error) {

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

// ParseProcessMaxVotersChanged is a log parse operation binding the contract event 0x36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be.
//
// Solidity: event ProcessMaxVotersChanged(bytes31 indexed processId, uint256 maxVoters)
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
	ProcessId [31]byte
	Sender    common.Address
	Result    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessResultsSet is a free log retrieval operation binding the contract event 0xdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e.
//
// Solidity: event ProcessResultsSet(bytes31 indexed processId, address indexed sender, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessResultsSet(opts *bind.FilterOpts, processId [][31]byte, sender []common.Address) (*ProcessRegistryProcessResultsSetIterator, error) {

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

// WatchProcessResultsSet is a free log subscription operation binding the contract event 0xdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e.
//
// Solidity: event ProcessResultsSet(bytes31 indexed processId, address indexed sender, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessResultsSet(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessResultsSet, processId [][31]byte, sender []common.Address) (event.Subscription, error) {

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

// ParseProcessResultsSet is a log parse operation binding the contract event 0xdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e.
//
// Solidity: event ProcessResultsSet(bytes31 indexed processId, address indexed sender, uint256[] result)
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
	ProcessId                [31]byte
	Sender                   common.Address
	OldStateRoot             *big.Int
	NewStateRoot             *big.Int
	NewVotersCount           *big.Int
	NewOverwrittenVotesCount *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterProcessStateTransitioned is a free log retrieval operation binding the contract event 0x7156cf1102d86d870b69eed151febab016e34063d74c714b17538efccfd9d019.
//
// Solidity: event ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, uint256 oldStateRoot, uint256 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStateTransitioned(opts *bind.FilterOpts, processId [][31]byte, sender []common.Address) (*ProcessRegistryProcessStateTransitionedIterator, error) {

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

// WatchProcessStateTransitioned is a free log subscription operation binding the contract event 0x7156cf1102d86d870b69eed151febab016e34063d74c714b17538efccfd9d019.
//
// Solidity: event ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, uint256 oldStateRoot, uint256 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStateTransitioned(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStateTransitioned, processId [][31]byte, sender []common.Address) (event.Subscription, error) {

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

// ParseProcessStateTransitioned is a log parse operation binding the contract event 0x7156cf1102d86d870b69eed151febab016e34063d74c714b17538efccfd9d019.
//
// Solidity: event ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, uint256 oldStateRoot, uint256 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount)
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
	ProcessId [31]byte
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessStatusChanged is a free log retrieval operation binding the contract event 0x56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf.
//
// Solidity: event ProcessStatusChanged(bytes31 indexed processId, uint8 oldStatus, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStatusChanged(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessStatusChangedIterator, error) {

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

// WatchProcessStatusChanged is a free log subscription operation binding the contract event 0x56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf.
//
// Solidity: event ProcessStatusChanged(bytes31 indexed processId, uint8 oldStatus, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStatusChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStatusChanged, processId [][31]byte) (event.Subscription, error) {

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

// ParseProcessStatusChanged is a log parse operation binding the contract event 0x56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf.
//
// Solidity: event ProcessStatusChanged(bytes31 indexed processId, uint8 oldStatus, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStatusChanged(log types.Log) (*ProcessRegistryProcessStatusChanged, error) {
	event := new(ProcessRegistryProcessStatusChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
