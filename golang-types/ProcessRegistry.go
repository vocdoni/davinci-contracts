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
	UniqueValues   bool
	CostFromWeight bool
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

// DAVINCITypesParamsMod is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesParamsMod struct {
	Census    bool
	Status    bool
	Duration  bool
	Metadata  bool
	MaxVoters bool
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
	ParamsMod             DAVINCITypesParamsMod
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_stVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rVerifier\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_blobsDA\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"limbIndex\",\"type\":\"uint8\"}],\"name\":\"InvalidBlobCommitmentLimb\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncryptionKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValueSum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"ProcessMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOB_INDEX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobsDA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"census\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"duration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"metadata\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maxVoters\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.ParamsMod\",\"name\":\"paramsMod\",\"type\":\"tuple\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"census\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"duration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"metadata\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maxVoters\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.ParamsMod\",\"name\":\"paramsMod\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"census\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"duration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"metadata\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maxVoters\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.ParamsMod\",\"name\":\"paramsMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"setProcessMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346101905760808161355a803803809161001f82856101ed565b8339810103126101905761003281610224565b61003e60208301610235565b606061004c60408501610235565b9301518015158091036101905760015f5563ffffffff83169384156101b6576001600160a01b038316156101a7576001600160a01b03169182156101a7576003805460048054600160201b600160e01b0319909216604094851b600160401b600160e01b031617602097881b67ffffffff000000001617909255600164ff0000000160a01b0319169390931760c09290921b60ff60c01b169190911782555163342e1df160e21b8152908101929092523060248301528160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af490811561019c575f9161015e575b506004805463ffffffff60a01b191660a09290921b63ffffffff60a01b16919091179055604051613310908161024a8239f35b90506020813d602011610194575b81610179602093836101ed565b810103126101905761018a90610224565b5f61012b565b5f80fd5b3d915061016c565b6040513d5f823e3d90fd5b63baa3de5f60e01b5f5260045ffd5b60405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b6044820152606490fd5b601f909101601f19168101906001600160401b0382119082101761021057604052565b634e487b7160e01b5f52604160045260245ffd5b519063ffffffff8216820361019057565b51906001600160a01b03821682036101905756fe6080806040526004361015610012575f80fd5b5f905f3560e01c908163026cdee8146129da5750806304ed00fa1461286f578063082b642e14612662578063097063921461263857806326795bc31461261e5780634c0acc56146125ce57806359d821c4146121fd57806362fa11fc146121b857806368141f2c146120e65780636c7aff7f14611d5857806372c628ef14611d3d578063766422e01461193d578063784df74a146117c1578063848df5401461179d57806391ccf666146115835780639b4649941461137a5780639c18deaf14611351578063adc879e91461132b578063c557fb4b14610c41578063cddf08bc14610c1a578063cdf497c814610bf4578063f9aa449914610b665763fa06e4501461011b575f80fd5b34610b63576102c0366003190112610b6357600435906005821015610b6357602435916064356044356101203660831901126107f3576001600160401b036101a435116107f35760a06101a435360360031901126107f3576101c4356001600160401b038111610b6157610193903690600401612c3a565b9290916040366101e31901126107ef5760a0366102231901126107ef576004805433808952600260209081526040808b20549051636846b35160e11b815260a09490941c63ffffffff169484019490945260248301919091526001600160401b03909216604482015288918160648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156109b0578891610b32575b5060ff19168088526001602052604088205490989060081c6001600160a01b0316610b235760ff6102566130d2565b16158015610b0f575b610b005760ff61026d6130e2565b16158015610ae3575b610ad45765ffffffffffff6101243511610ac55765ffffffffffff6101443511610ab65761ffff6101643511610aa757677fffffffffffffff6101843511610a8957610124356101443511610a9857610164356101843511610a89578315610a7a576102ef6102e361305b565b6101a4356004016131b6565b87600460ff8916118015610a61575b610a5257610a3e576003871480610a2f575b610a20575f5160206132bb5f395f51905f526101e43510801590610a06575b6109f75782156109e857156109e1575b4281106109d257426103518383612f9f565b11156109c357878752600160205261036d604088209687612fcd565b6005860155600685015560078401558254610100600160a81b0319163360081b610100600160a81b03161783556101e435600184015561020435600284015560036101a4356004013510156107f35760405190631044f8bb60e11b82528560048301526103e4602483016101a43560040135612bde565b608435928315158085036107eb57604484015260a435918215158084036109bf57606485015260c43560ff81168091036109bf57608485015260e43560ff811681036109bf5760ff1660a4850152610104359360ff85168086036109bb5760c48201526101243560e48201526101443561010482015261016435610124820152610184356101448201526101e435610164820152610204356101848201526020816101a48173__$4e0a7d2578f6e358dccb8beebb7c720d9f$__5af49081156109b057889161097a575b5060038701556001600160401b038211610966576104cf600c870154612d90565b601f8111610933575b508690601f83116001146108c35761053b9392918891836108b8575b50508160011b915f199060031b1c191617600c8601555b610525600d860194859060ff801983541691151516179055565b835461ff00191690151560081b61ff0016178355565b6105436130d2565b90825462ff000064ff0000000063ff00000061055d6130e2565b60181b169360201b169360101b169064ffffff00001916171717905561012435600e82015561014435600f8201556101643560108201556101843560118201556012810160ff1981541660ff6101a435600401351617905560246101a435013560138201556014810160018060a01b036105dc60446101a43501612fff565b82546001600160a01b0319169116179055601581016106056101a4356064810190600401613013565b906001600160401b0382116108a45761061e8354612d90565b601f8111610869575b508490601f8311600114610802576017949392918691836107f7575b50508160011b915f199060031b1c19161790555b61067e61066960846101a4350161307b565b601683019060ff801983541691151516179055565b43600a820155016106a261069061305b565b829060ff801983541691151516179055565b6106c36106ad61306b565b825461ff00191690151560081b61ff0016178255565b61026435801515908181036107f3575081546102843590811515918281036107ef57506102a435801515908181036107eb5764ff0000000063ff000000939262ff0000925060201b169460101b169064ffffff00001916179160181b161717905560035463ffffffff811663ffffffff81146107d757600163ffffffff9101169063ffffffff191617600355338152600260205260408120918254926001600160401b038416936001600160401b0385146107c3576001600160401b0360016020960116906001600160401b03191617905560405191817feefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d1339280a38152f35b634e487b7160e01b84526011600452602484fd5b634e487b7160e01b83526011600452602483fd5b8680fd5b8580fd5b8380fd5b013590505f80610643565b8386526020862091601f198416875b818110610851575091600193918560179897969410610838575b505050811b019055610657565b01355f19600384901b60f8161c191690555f808061082b565b91936020600181928787013581550195019201610811565b6108949084875260208720601f850160051c8101916020861061089a575b601f0160051c0190613045565b5f610627565b9091508190610887565b634e487b7160e01b85526041600452602485fd5b013590505f806104f4565b600c870188526020882091885b601f198516811061091b575091839160019361053b9695601f19811610610902575b505050811b01600c86015561050b565b01355f19600384901b60f8161c191690555f80806108f2565b909260206001819286860135815501940191016108d0565b61096090600c8801895260208920601f850160051c8101916020861061089a57601f0160051c0190613045565b5f6104d8565b634e487b7160e01b87526041600452602487fd5b90506020813d6020116109a8575b8161099560209383612d51565b810103126109a457515f6104ae565b5f80fd5b3d9150610988565b6040513d8a823e3d90fd5b8880fd5b8780fd5b637616640160e01b8752600487fd5b632ca4094f60e21b8752600487fd5b504261033f565b637616640160e01b8852600488fd5b63208e53e560e11b8852600488fd5b505f5160206132bb5f395f51905f5261020435101561032f565b6307a92f1960e51b8852600488fd5b50610a3861306b565b15610310565b634e487b7160e01b88526021600452602488fd5b6307a92f1960e51b8952600489fd5b50508787151580156102fe5750508760038814156102fe565b632c45be6f60e01b8852600488fd5b636d403ec760e11b8852600488fd5b63207ea56d60e01b8852600488fd5b634dba795160e01b8852600488fd5b6363f4b4b760e01b8852600488fd5b63b1911d8b60e01b8852600488fd5b632cbdc23160e01b8852600488fd5b50610aec6130e2565b60ff80610af76130d2565b16911611610276565b63ac38930b60e01b8852600488fd5b50600860ff610b1c6130d2565b161161025f565b635040c4d360e11b8852600488fd5b610b54915060203d602011610b5a575b610b4c8183612d51565b810190612fe5565b5f610227565b503d610b42565b845b80fd5b5034610b635780600319360112610b63576004805460405163233ace1160e01b8152929160209184919082906001600160a01b03165afa908115610be85790610bb5575b602090604051908152f35b506020813d602011610be0575b81610bcf60209383612d51565b810103126109a45760209051610baa565b3d9150610bc2565b604051903d90823e3d90fd5b5034610b635780600319360112610b6357602060ff60045460c01c166040519015158152f35b5034610b635780600319360112610b6357602063ffffffff60045460a01c16604051908152f35b5034610b6357610c5036612c67565b91610c5c94939461329c565b60ff19841693841561131c5760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156111345787916112ed575b50156112de578386526001602052604086208054909390600881901c6001600160a01b0316156112cf5760ff166005811015610a3e576112c0576005840154610d10600686015482612f9f565b4210156112b15742106112a2578660e0604051610d2c81612d35565b8281528260208201528260408201528260608201528260808201528260a08201528260c0820152015280830190610100848303126109bf5781601f850112156109bf5760405191610100610d808185612d51565b8390860191821161129e5785905b82821061128e5750505081519660208301519460408401519460608501519460808101519760a08201519860c08301519260e00151916040519d8e610dd281612d35565b818152826020820152604081019b8c52606081019a8b526080810193845260a081019c8d5260c0810195865260e0019384525f5160206132bb5f395f51905f521190811591611276575b5061119d575f5160206132bb5f395f51905f52815110156112675760ff60128c015416600381101561125357600103611239578d60018060a01b0360148d0154166020835160246040518094819363650e5fcf60e01b835260048301525afa918215610be85791611207575b50801580156111fe575b6111ef5760178c015460ff166111bb57505160138b0154036111ac575b8b519760038b019889540361119d57518751610eca91612fc0565b988915158061118b575b61117c57908d96959493929160ff60045460c01c16610ff9575b505060035460401c6001600160a01b0316919050813b156107ef578593610f2b60405196879586948594635c73957b60e11b8652600486016130a8565b03915afa8015610fee57610fd9575b5050600b9160208701519055610f5560088501938454612f9f565b8093555192610f6960098201948554612f9f565b8094550180545f198114610fc5576001019055602084519401516040519485526020850152604084015260608301527f7156cf1102d86d870b69eed151febab016e34063d74c714b17538efccfd9d01960803393a36001815580f35b634e487b7160e01b87526011600452602487fd5b81610fe391612d51565b6107eb57865f610f3a565b6040513d84823e3d90fd5b909192939495965051915190519060405192611016606085612d51565b60308452602084019060403683378060801c6111675760801b90528060801c6111535760801b60308301528060801c61113f57908c959493929160801b6040820152611094602073__$176218451d0b397034f549dcd265d6db5e$__926040518093819263da23b04d60e01b83528460048401526024830190612b57565b0381855af49081156111345787916110ff575b50813b156107eb578690602460405180948193634e733ef560e11b835260048301525af49081156110f45786916110df575b80610eee565b816110e991612d51565b610b6157845f6110d9565b6040513d88823e3d90fd5b9650506020863d60201161112c575b8161111b60209383612d51565b810103126109a4578c95515f6110a7565b3d915061110e565b6040513d89823e3d90fd5b6316416a3d60e01b8d52600260045260248dfd5b6316416a3d60e01b8e52600160045260248efd5b506316416a3d60e01b8f5260048f905260248ffd5b6357d18d5360e11b8e5260048efd5b5060088b015460078c01541115610ed4565b630b6fac0360e41b8e5260048efd5b635e32eadd60e01b8d5260048dfd5b60168c015460ff16159150816111e1575b5015610eaf57635e32eadd60e01b8d5260048dfd5b9050600a8b0154115f6111cc565b635e32eadd60e01b8f5260048ffd5b50438111610e92565b90506020813d602011611231575b8161122260209383612d51565b810103126109a457515f610e88565b3d9150611215565b5160138b015414610eaf57635e32eadd60e01b8d5260048dfd5b634e487b7160e01b8f52602160045260248ffd5b635e32eadd60e01b8e5260048efd5b5f5160206132bb5f395f51905f52915010155f610e1c565b8135815260209182019101610d8e565b8980fd5b63e843c5eb60e01b8752600487fd5b63e843c5eb60e01b8852600488fd5b6307a92f1960e51b8752600487fd5b634d36eb6960e01b8852600488fd5b632299770d60e11b8652600486fd5b61130f915060203d602011611315575b6113078183612d51565b810190612f87565b5f610cc3565b503d6112fd565b63cbf4a64560e01b8752600487fd5b5034610b635780600319360112610b6357602063ffffffff600354821c16604051908152f35b5034610b635780600319360112610b63576004546040516001600160a01b039091168152602090f35b5034610b63576040366003190112610b6357611394612b39565b60243560ff1982169182156115745760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af490811561156957849161154a575b501561153b578183526001602052604083208054600881901c6001600160a01b0316801561152c57330361151e5760ff601783015460101c161561151e5760ff16600581101561150a5780151590816114fe575b506114ef576006600582015491019081549083159182156114da575b82156114bb575b50506114ac57817f4b5dee969e28ac23e61cb3c24582a7b667358bcc41cd6e31e76001e7624af2429260209255604051908152a280f35b637616640160e01b8452600484fd5b8192506114cb856114d193612f9f565b92612f9f565b10155f80611475565b91506114e68482612f9f565b4210159161146e565b6307a92f1960e51b8452600484fd5b6003915014155f611452565b634e487b7160e01b85526021600452602485fd5b6282b42960e81b8552600485fd5b634d36eb6960e01b8652600486fd5b632299770d60e11b8352600483fd5b611563915060203d602011611315576113078183612d51565b5f6113fe565b6040513d86823e3d90fd5b63cbf4a64560e01b8452600484fd5b5034610b63576040366003190112610b635761159d612b39565b906024356001600160401b038111611799576115bd903690600401612c3a565b9260ff191690811561178a578183526001602052604083208054600881901c6001600160a01b0316801561152c57330361151e5760ff601783015460181c161561151e5760ff16600581101561150a57801515908161177e575b506114ef57600c01936001600160401b03811161176a576116388554612d90565b601f811161173a575b508394601f82116001146116b6578185967f970502e3e02e01b5b447e50c390a402175fe6a1e78b7c238e714e28c47622e839596916116ab575b508260011b905f198460031b1c19161790555b6116a5604051928392602084526020840191613088565b0390a280f35b90508301355f61167b565b8085526020852095601f198316865b818110611722575090837f970502e3e02e01b5b447e50c390a402175fe6a1e78b7c238e714e28c47622e839697989210611709575b5050600182811b01905561168e565b8401355f19600385901b60f8161c191690555f806116fa565b858301358955600190980197602092830192016116c5565b6117649086865260208620601f840160051c8101916020851061089a57601f0160051c0190613045565b5f611641565b634e487b7160e01b84526041600452602484fd5b6003915014155f611617565b63cbf4a64560e01b8352600483fd5b5080fd5b5034610b635780600319360112610b6357602063ffffffff60035416604051908152f35b5034610b63576020366003190112610b635760ff196117de612b39565b1681526001602052604090208054600182016117f990612d72565b9160038101549060058101546006820154600783015460088401546009850154600a86015491600b87015493600c880161183290612dc8565b9561183f600d8a01612e68565b9761184c60128b01612ed8565b9960170161185990612f3a565b9a604051809e8e829f60ff8491169061187191612b4a565b600160a01b600190039060081c1660208301528051604083015260200151906060015260808d015260a08c015260c08b015260e08a01526101008901526101208801526101408701526101608601526101808501610380905261038085016118d891612b57565b906101a085016118e791612b7b565b8381036102c08501526118f991612beb565b815115156102e084015260208201511515610300840152604082015115156103208401526060820151151561034084015260809091015115156103608301525b0390f35b50346109a45761194c36612c67565b9361195b95939592919261329c565b60ff198116908115611d2e5760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611c6e575f91611d0f575b5015611d0057805f52600160205260405f2094855460018060a01b038160081c1615611cf15760ff1696600588101580611cc257600289148015611ce5575b611cd657611cc257600188141580611ca7575b611c985781850194610120818703126109a45785601f820112156109a4576101209460405196611a448789612d51565b878784019182116109a45783905b828210611c8857505050865160038a015403611c79576004546001600160a01b031690813b156109a4575f93611a9e60405196879586948594635c73957b60e11b8652600486016130a8565b03915afa8015611c6e57611c59575b5060405195611abc8388612d51565b600887526020870192601f190136843760015b6009811015611b2c578060051b8501515f198201828111611b18578951811015611b045760051b890160200152600101611acf565b634e487b7160e01b88526032600452602488fd5b634e487b7160e01b88526011600452602488fd5b8584848a8560048c8160ff19825416178155018251906001600160401b03821161096657680100000000000000008211610966578054828255808310611c3e575b50908492918690885260208820885b838110611c27575050505060407f6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b91611bb782518092612b4a565b60046020820152a26040519060208201906020835251809152604082019390855b818110611c1157505050807fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e9133940390a36001815580f35b8251865260209586019590920191600101611bd8565b825181830155879550602090920191600101611b7c565b81885260208820611c53918101908401613045565b87611b6d565b611c669194505f90612d51565b5f925f611aad565b6040513d5f823e3d90fd5b630b6fac0360e41b5f5260045ffd5b8135815260209182019101611a52565b63e843c5eb60e01b5f5260045ffd5b50611cbb6005880154600689015490612f9f565b4210611a14565b634e487b7160e01b5f52602160045260245ffd5b6307a92f1960e51b5f5260045ffd5b50505f60048914611a01565b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b611d28915060203d602011611315576113078183612d51565b5f6119c2565b63cbf4a64560e01b5f5260045ffd5b346109a4575f3660031901126109a457602060405160048152f35b346109a45760403660031901126109a457611d71612b39565b6024356001600160401b0381116109a457806004019060a060031982360301126109a45760ff198316928315611d2e5760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611c6e575f916120c7575b5015611d0057825f52600160205260405f2091825460018060a01b038160081c168015611cf15733036120b95760ff60178501541680156120b957601285019182549160ff83169185359260038410156109a4576003811015611cc25783036120aa57611e6b60ff92876131b6565b166005811015611cc257801515908161209e575b50611cd657611e976005870154600688015490612f9f565b421015611c985760ff169060ff19161790556024820135928360138201556014810160018060a01b03611ecc60448601612fff565b82546001600160a01b0319169116179055606483019260158201611ef08585613013565b906001600160401b03821161208a57611f098354612d90565b601f811161205a575b505f90601f8311600114611fbf57611f7c6084611faf99967f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed29b999686611f8f97611f949b976016975f92611fb4575b50508160011b915f199060031b1c19161790555b0161307b565b91019060ff801983541691151516179055565b613013565b92906040519384938452604060208501526040840191613088565b0390a2005b013590505f80611f62565b601f19831691845f5260205f20925f5b8181106120425750611faf99967f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed29b9996600187611f949b97601697611f7c97611f8f9b60849810612029575b505050811b019055611f76565b01355f19600384901b60f8161c191690555f808061201c565b91936020600181928787013581550195019201611fcf565b61208490845f5260205f20601f850160051c8101916020861061089a57601f0160051c0190613045565b89611f12565b634e487b7160e01b5f52604160045260245ffd5b60039150141588611e7f565b63f37f7b5d60e01b5f5260045ffd5b6282b42960e81b5f5260045ffd5b6120e0915060203d602011611315576113078183612d51565b84611dfc565b346109a45760203660031901126109a4576004356001600160a01b038116908181036109a457600480545f93845260026020908152604094859020549451636846b35160e11b815260a09290921c63ffffffff16928201929092526001600160a01b039290921660248301526001600160401b039092166044820152908160648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af48015611c6e576020915f9161219b575b506040519060ff19168152f35b6121b29150823d8411610b5a57610b4c8183612d51565b8261218e565b346109a45760203660031901126109a4576004356001600160a01b038116908190036109a4575f52600260205260206001600160401b0360405f205416604051908152f35b346109a45760203660031901126109a457612216612b39565b60405161222281612cc7565b5f81525f602082015260405161223781612ce3565b5f81525f602082015260408201525f6060820152606060808201525f60a08201525f60c08201525f60e08201525f6101008201525f6101208201525f6101408201525f610160820152606061018082015260405161229481612cfe565b5f81525f60208201525f60408201525f60608201525f60808201525f60a08201525f60c08201525f60e08201525f6101008201526101a08201526040516122da81612d1a565b5f81525f60208201525f60408201526060808201525f60808201526101c08201526101e06040519161230b83612d1a565b5f83525f60208401525f60408401525f60608401525f6080840152015260ff19165f52600160205260405f2060405161234381612cc7565b81549060ff82166005811015611cc257815260089190911c6001600160a01b0316602082019081529061237860018401612d72565b60408201908152600384015460608301908152600485019160405180938491602082549182815201915f5260205f20905f5b8181106125b557505050602092916123c3910385612d51565b60808501938452612458601760058901549860a08801998a52600681015460c0890152600781015460e089015260088101546101008901526009810154610120890152600a810154610140890152600b810154610160890152612428600c8201612dc8565b61018089015261243a600d8201612e68565b6101a089015261244c60128201612ed8565b6101c089015201612f3a565b6101e0860152604051958287526124728388018751612b4a565b516001600160a01b031660408701525180516060870152015160808501525160a0840152516103a060c084015280516103c084018190526103e084019491602001905f5b81811061259f575050506101e061255d6125338596611939945160e088015260c086015161010088015260e08601516101208801526101008601516101408801526101208601516101608801526101408601516101808801526101608601516101a0880152610180860151601f19888303016101c0890152612b57565b6125456101a086015184880190612b7b565b6101c0850151868203601f1901610300880152612beb565b920151805115156103208501526020810151151561034085015260408101511515610360850152606081015115156103808501526080015115156103a0840152565b82518752602096870196909201916001016124b6565b82548452879450602090930192600192830192016123aa565b346109a4575f3660031901126109a4576003546040805163233ace1160e01b81529160209183916004918391901c6001600160a01b03165afa8015611c6e575f90610bb557602090604051908152f35b346109a4575f3660031901126109a45760206040515f8152f35b346109a4575f3660031901126109a4576003546040805191901c6001600160a01b03168152602090f35b346109a45760403660031901126109a45761267b612b39565b60243560058110156109a45760ff198216918215611d2e5760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611c6e575f91612850575b5015611d0057600460ff821611611cd6575f82815260016020526040902080549091600882901c6001600160a01b03168015611cf15733036120b957601783019160ff835460081c16156120b95760ff169161274a82846130f2565b15611cd657600182146127a4575b50906127a18261278b6040947f6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b96612fcd565b61279784518094612b4a565b6020830190612b4a565ba2005b6005840180549194914210612847576127be905442612fc0565b6006820194855482036127d4575b505092612758565b549093929060101c60ff16156120b95761278b816127a193887f4b5dee969e28ac23e61cb3c24582a7b667358bcc41cd6e31e76001e7624af24260208960409a7f6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b9c558a51908152a293509394506127cc565b4290555f6127be565b612869915060203d602011611315576113078183612d51565b836126ee565b346109a45760203660031901126109a45760ff1961288b612b39565b165f52600160205260405f206040516128a381612cc7565b815460ff81166005811015611cc257825260081c6001600160a01b031660208201526128d160018301612d72565b604082015260038201546060820152600482016040519081602082549182815201915f5260205f20905f5b8181106129c45760206129bc888888612917818a0382612d51565b608082015260058201549060a081019182526101e06129b0601760068601549560c08501968752600781015460e086015260088101546101008601526009810154610120860152600a810154610140860152600b810154610160860152612980600c8201612dc8565b610180860152612992600d8201612e68565b6101a08601526129a460128201612ed8565b6101c086015201612f3a565b91015251905190612f9f565b604051908152f35b82548452602090930192600192830192016128fc565b346109a45760403660031901126109a4576129f3612b39565b60ff19811691602435918315611d2e5760048054632145bb8f60e01b845260ff199092169083015260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611c6e575f91612b1a575b5015611d00575f8281526001602052604090208054600881901c6001600160a01b03168015611cf15733036120b95760ff601783015460201c16156120b95760ff166005811015611cc2578015159081612b0e575b50611cd65781158015612b01575b612af257817fc12d00e36fb3763916c443453e211136b31b621e8f92b16e7eec16bf3919ae0f9260076020930155604051908152a2005b632c45be6f60e01b5f5260045ffd5b5060088101548210612abb565b60039150141584612aad565b612b33915060203d602011611315576113078183612d51565b83612a58565b6004359060ff19821682036109a457565b906005821015611cc25752565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b610100809180511515845260208101511515602085015260ff604082015116604085015260ff606082015116606085015260ff608082015116608085015260a081015160a085015260c081015160c085015260e081015160e08501520151910152565b906003821015611cc25752565b90612bf7818351612bde565b6020820151602082015260018060a01b036040830151166040820152608080612c2f606085015160a0606086015260a0850190612b57565b930151151591015290565b9181601f840112156109a4578235916001600160401b0383116109a457602083818601950101116109a457565b9060606003198301126109a45760043560ff19811681036109a457916024356001600160401b0381116109a45781612ca191600401612c3a565b92909291604435906001600160401b0382116109a457612cc391600401612c3a565b9091565b61020081019081106001600160401b0382111761208a57604052565b604081019081106001600160401b0382111761208a57604052565b61012081019081106001600160401b0382111761208a57604052565b60a081019081106001600160401b0382111761208a57604052565b61010081019081106001600160401b0382111761208a57604052565b90601f801991011681019081106001600160401b0382111761208a57604052565b90604051612d7f81612ce3565b602060018294805484520154910152565b90600182811c92168015612dbe575b6020831014612daa57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612d9f565b9060405191825f825492612ddb84612d90565b8084529360018116908115612e465750600114612e02575b50612e0092500383612d51565b565b90505f9291925260205f20905f915b818310612e2a575050906020612e00928201015f612df3565b6020919350806001915483858901015201910190918492612e11565b905060209250612e0094915060ff191682840152151560051b8201015f612df3565b90604051612e7581612cfe565b6101006004829460ff815481811615158652818160081c1615156020870152818160101c166040870152818160181c16606087015260201c166080850152600181015460a0850152600281015460c0850152600381015460e08501520154910152565b90604051612ee581612d1a565b809260ff815416906003821015611cc2579082526001810154602083015260028101546001600160a01b0316604083015260809060ff90600490612f2b60038201612dc8565b60608601520154161515910152565b90604051612f4781612d1a565b608060ff82945481811615158452818160081c1615156020850152818160101c1615156040850152818160181c161515606085015260201c161515910152565b908160209103126109a4575180151581036109a45790565b91908201809211612fac57565b634e487b7160e01b5f52601160045260245ffd5b91908203918211612fac57565b906005811015611cc25760ff80198354169116179055565b908160209103126109a4575160ff19811681036109a45790565b356001600160a01b03811681036109a45790565b903590601e19813603018212156109a457018035906001600160401b0382116109a4576020019181360383136109a457565b818110613050575050565b5f8155600101613045565b6102243580151581036109a45790565b6102443580151581036109a45790565b3580151581036109a45790565b908060209392818452848401375f828201840152601f01601f1916010190565b92906130c1906130cf9593604086526040860191613088565b926020818503910152613088565b90565b60c43560ff811681036109a45790565b60e43560ff811681036109a45790565b906005821015611cc25760058110159182611cc2578082146131af575f6002821480156131a3575b818115613193575b5061318b57611cc25780156131725760031461313e5750505f90565b81611cc2578015918215613165575b821561315857505090565b909150611cc25760011490565b506002811491505f61314d565b506003811491505f821561316557821561315857505090565b505050505f90565b9050611cc2576001821481613122565b50505f6004821461311a565b5050505f90565b90813560038110156109a457600103613247576001600160a01b036131dd60408401612fff565b161561323857158061322c575b613205576020905b01358015159081613214575b5061320557565b635e32eadd60e01b5f5260045ffd5b5f5160206132bb5f395f51905f52915010155f6131fe565b506020810135156131ea565b63562e597160e11b5f5260045ffd5b506132556060820182613013565b90501561328d57602081013515613205576132726080820161307b565b61327e576020906131f2565b63f545b7bf60e01b5f5260045ffd5b630f8b932160e11b5f5260045ffd5b60025f54146132ab5760025f55565b633ee5aeb560e01b5f5260045ffdfe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220db82c9ee9e74f65b3d0b8452c531d991a0b0caa6e5eca8536b7b9f5a92ffeee864736f6c634300081c0033",
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),(bool,bool,bool,bool,bool)))
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),(bool,bool,bool,bool,bool)))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [31]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x59d821c4.
//
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),(bool,bool,bool,bool,bool)))
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
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, (bool,bool,bool,bool,bool) paramsMod)
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
	ParamsMod             DAVINCITypesParamsMod
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
		ParamsMod             DAVINCITypesParamsMod
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
	outstruct.ParamsMod = *abi.ConvertType(out[14], new(DAVINCITypesParamsMod)).(*DAVINCITypesParamsMod)

	return *outstruct, err

}

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, (bool,bool,bool,bool,bool) paramsMod)
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
	ParamsMod             DAVINCITypesParamsMod
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, (bool,bool,bool,bool,bool) paramsMod)
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
	ParamsMod             DAVINCITypesParamsMod
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

// NewProcess is a paid mutator transaction binding the contract method 0xfa06e450.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (bool,bool,bool,bool,bool) paramsMod) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, paramsMod DAVINCITypesParamsMod) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, paramsMod)
}

// NewProcess is a paid mutator transaction binding the contract method 0xfa06e450.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (bool,bool,bool,bool,bool) paramsMod) returns(bytes31)
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, paramsMod DAVINCITypesParamsMod) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, paramsMod)
}

// NewProcess is a paid mutator transaction binding the contract method 0xfa06e450.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (bool,bool,bool,bool,bool) paramsMod) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, paramsMod DAVINCITypesParamsMod) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, paramsMod)
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

// SetProcessMetadata is a paid mutator transaction binding the contract method 0x91ccf666.
//
// Solidity: function setProcessMetadata(bytes31 processId, string metadata) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessMetadata(opts *bind.TransactOpts, processId [31]byte, metadata string) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessMetadata", processId, metadata)
}

// SetProcessMetadata is a paid mutator transaction binding the contract method 0x91ccf666.
//
// Solidity: function setProcessMetadata(bytes31 processId, string metadata) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessMetadata(processId [31]byte, metadata string) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMetadata(&_ProcessRegistry.TransactOpts, processId, metadata)
}

// SetProcessMetadata is a paid mutator transaction binding the contract method 0x91ccf666.
//
// Solidity: function setProcessMetadata(bytes31 processId, string metadata) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessMetadata(processId [31]byte, metadata string) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMetadata(&_ProcessRegistry.TransactOpts, processId, metadata)
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

// ProcessRegistryProcessDurationUpdatedIterator is returned from FilterProcessDurationUpdated and is used to iterate over the raw logs and unpacked data for ProcessDurationUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessDurationUpdatedIterator struct {
	Event *ProcessRegistryProcessDurationUpdated // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessDurationUpdated)
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
		it.Event = new(ProcessRegistryProcessDurationUpdated)
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
func (it *ProcessRegistryProcessDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessDurationUpdated represents a ProcessDurationUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessDurationUpdated struct {
	ProcessId [31]byte
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessDurationUpdated is a free log retrieval operation binding the contract event 0x4b5dee969e28ac23e61cb3c24582a7b667358bcc41cd6e31e76001e7624af242.
//
// Solidity: event ProcessDurationUpdated(bytes31 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessDurationUpdated(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessDurationUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessDurationUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessDurationUpdatedIterator{contract: _ProcessRegistry.contract, event: "ProcessDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchProcessDurationUpdated is a free log subscription operation binding the contract event 0x4b5dee969e28ac23e61cb3c24582a7b667358bcc41cd6e31e76001e7624af242.
//
// Solidity: event ProcessDurationUpdated(bytes31 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessDurationUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessDurationUpdated, processId [][31]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessDurationUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessDurationUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessDurationUpdated", log); err != nil {
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

// ParseProcessDurationUpdated is a log parse operation binding the contract event 0x4b5dee969e28ac23e61cb3c24582a7b667358bcc41cd6e31e76001e7624af242.
//
// Solidity: event ProcessDurationUpdated(bytes31 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessDurationUpdated(log types.Log) (*ProcessRegistryProcessDurationUpdated, error) {
	event := new(ProcessRegistryProcessDurationUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessMaxVotersUpdatedIterator is returned from FilterProcessMaxVotersUpdated and is used to iterate over the raw logs and unpacked data for ProcessMaxVotersUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessMaxVotersUpdatedIterator struct {
	Event *ProcessRegistryProcessMaxVotersUpdated // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessMaxVotersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessMaxVotersUpdated)
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
		it.Event = new(ProcessRegistryProcessMaxVotersUpdated)
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
func (it *ProcessRegistryProcessMaxVotersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessMaxVotersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessMaxVotersUpdated represents a ProcessMaxVotersUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessMaxVotersUpdated struct {
	ProcessId [31]byte
	MaxVoters *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessMaxVotersUpdated is a free log retrieval operation binding the contract event 0xc12d00e36fb3763916c443453e211136b31b621e8f92b16e7eec16bf3919ae0f.
//
// Solidity: event ProcessMaxVotersUpdated(bytes31 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessMaxVotersUpdated(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessMaxVotersUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessMaxVotersUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessMaxVotersUpdatedIterator{contract: _ProcessRegistry.contract, event: "ProcessMaxVotersUpdated", logs: logs, sub: sub}, nil
}

// WatchProcessMaxVotersUpdated is a free log subscription operation binding the contract event 0xc12d00e36fb3763916c443453e211136b31b621e8f92b16e7eec16bf3919ae0f.
//
// Solidity: event ProcessMaxVotersUpdated(bytes31 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessMaxVotersUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessMaxVotersUpdated, processId [][31]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessMaxVotersUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessMaxVotersUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessMaxVotersUpdated", log); err != nil {
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

// ParseProcessMaxVotersUpdated is a log parse operation binding the contract event 0xc12d00e36fb3763916c443453e211136b31b621e8f92b16e7eec16bf3919ae0f.
//
// Solidity: event ProcessMaxVotersUpdated(bytes31 indexed processId, uint256 maxVoters)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessMaxVotersUpdated(log types.Log) (*ProcessRegistryProcessMaxVotersUpdated, error) {
	event := new(ProcessRegistryProcessMaxVotersUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessMaxVotersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessMetadataUpdatedIterator is returned from FilterProcessMetadataUpdated and is used to iterate over the raw logs and unpacked data for ProcessMetadataUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessMetadataUpdatedIterator struct {
	Event *ProcessRegistryProcessMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessMetadataUpdated)
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
		it.Event = new(ProcessRegistryProcessMetadataUpdated)
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
func (it *ProcessRegistryProcessMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessMetadataUpdated represents a ProcessMetadataUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessMetadataUpdated struct {
	ProcessId [31]byte
	Metadata  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessMetadataUpdated is a free log retrieval operation binding the contract event 0x970502e3e02e01b5b447e50c390a402175fe6a1e78b7c238e714e28c47622e83.
//
// Solidity: event ProcessMetadataUpdated(bytes31 indexed processId, string metadata)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessMetadataUpdated(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessMetadataUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessMetadataUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessMetadataUpdatedIterator{contract: _ProcessRegistry.contract, event: "ProcessMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchProcessMetadataUpdated is a free log subscription operation binding the contract event 0x970502e3e02e01b5b447e50c390a402175fe6a1e78b7c238e714e28c47622e83.
//
// Solidity: event ProcessMetadataUpdated(bytes31 indexed processId, string metadata)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessMetadataUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessMetadataUpdated, processId [][31]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessMetadataUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessMetadataUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessMetadataUpdated", log); err != nil {
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

// ParseProcessMetadataUpdated is a log parse operation binding the contract event 0x970502e3e02e01b5b447e50c390a402175fe6a1e78b7c238e714e28c47622e83.
//
// Solidity: event ProcessMetadataUpdated(bytes31 indexed processId, string metadata)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessMetadataUpdated(log types.Log) (*ProcessRegistryProcessMetadataUpdated, error) {
	event := new(ProcessRegistryProcessMetadataUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessMetadataUpdated", log); err != nil {
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

// ProcessRegistryProcessStatusUpdatedIterator is returned from FilterProcessStatusUpdated and is used to iterate over the raw logs and unpacked data for ProcessStatusUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessStatusUpdatedIterator struct {
	Event *ProcessRegistryProcessStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessStatusUpdated)
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
		it.Event = new(ProcessRegistryProcessStatusUpdated)
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
func (it *ProcessRegistryProcessStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessStatusUpdated represents a ProcessStatusUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessStatusUpdated struct {
	ProcessId [31]byte
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessStatusUpdated is a free log retrieval operation binding the contract event 0x6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b.
//
// Solidity: event ProcessStatusUpdated(bytes31 indexed processId, uint8 oldStatus, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStatusUpdated(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessStatusUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessStatusUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessStatusUpdatedIterator{contract: _ProcessRegistry.contract, event: "ProcessStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchProcessStatusUpdated is a free log subscription operation binding the contract event 0x6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b.
//
// Solidity: event ProcessStatusUpdated(bytes31 indexed processId, uint8 oldStatus, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStatusUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStatusUpdated, processId [][31]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessStatusUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessStatusUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusUpdated", log); err != nil {
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

// ParseProcessStatusUpdated is a log parse operation binding the contract event 0x6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b.
//
// Solidity: event ProcessStatusUpdated(bytes31 indexed processId, uint8 oldStatus, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStatusUpdated(log types.Log) (*ProcessRegistryProcessStatusUpdated, error) {
	event := new(ProcessRegistryProcessStatusUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
