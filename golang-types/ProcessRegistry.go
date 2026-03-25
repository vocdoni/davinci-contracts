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

// DAVINCITypesCensusParamsMod is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesCensusParamsMod struct {
	CensusRoot               bool
	ContractAddress          bool
	CensusURI                bool
	OnchainAllowAnyValidRoot bool
}

// DAVINCITypesEncryptionKey is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesEncryptionKey struct {
	X *big.Int
	Y *big.Int
}

// DAVINCITypesParamsMod is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesParamsMod struct {
	Status    bool
	Duration  bool
	Metadata  bool
	MaxVoters bool
	Census    DAVINCITypesCensusParamsMod
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_stVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rVerifier\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_blobsDA\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"limbIndex\",\"type\":\"uint8\"}],\"name\":\"InvalidBlobCommitmentLimb\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncryptionKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValueSum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"ProcessMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOB_INDEX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobsDA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"duration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"metadata\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maxVoters\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"censusRoot\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"contractAddress\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"censusURI\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.CensusParamsMod\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structDAVINCITypes.ParamsMod\",\"name\":\"paramsMod\",\"type\":\"tuple\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"duration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"metadata\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maxVoters\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"censusRoot\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"contractAddress\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"censusURI\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.CensusParamsMod\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structDAVINCITypes.ParamsMod\",\"name\":\"paramsMod\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"duration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"metadata\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"maxVoters\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"censusRoot\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"contractAddress\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"censusURI\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.CensusParamsMod\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structDAVINCITypes.ParamsMod\",\"name\":\"paramsMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"setProcessMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346101905760808161384f803803809161001f82856101ed565b8339810103126101905761003281610224565b61003e60208301610235565b606061004c60408501610235565b9301518015158091036101905760015f5563ffffffff83169384156101b6576001600160a01b038316156101a7576001600160a01b03169182156101a7576003805460048054600160201b600160e01b0319909216604094851b600160401b600160e01b031617602097881b67ffffffff000000001617909255600164ff0000000160a01b0319169390931760c09290921b60ff60c01b169190911782555163342e1df160e21b8152908101929092523060248301528160448173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af490811561019c575f9161015e575b506004805463ffffffff60a01b191660a09290921b63ffffffff60a01b16919091179055604051613605908161024a8239f35b90506020813d602011610194575b81610179602093836101ed565b810103126101905761018a90610224565b5f61012b565b5f80fd5b3d915061016c565b6040513d5f823e3d90fd5b63baa3de5f60e01b5f5260045ffd5b60405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9251608a1b6044820152606490fd5b601f909101601f19168101906001600160401b0382119082101761021057604052565b634e487b7160e01b5f52604160045260245ffd5b519063ffffffff8216820361019057565b51906001600160a01b03821682036101905756fe6080806040526004361015610012575f80fd5b5f905f3560e01c908163026cdee814612bf75750806304ed00fa14612a84578063082b642e1461287a578063097063921461285057806326795bc3146128365780634c0acc56146127e657806359d821c41461242157806362fa11fc146123dc57806368141f2c1461230a5780636c7aff7f14611e2c57806372c628ef14611e11578063766422e0146119fd578063784df74a146118a1578063848df5401461187d578063851c345314610d7957806391ccf66614610b545780639b4649941461094b5780639c18deaf14610922578063adc879e9146108fc578063c557fb4b146101fc578063cddf08bc146101d5578063cdf497c8146101af5763f9aa44991461011b575f80fd5b346101ac57806003193601126101ac576004805460405163233ace1160e01b8152929160209184919082906001600160a01b03165afa9081156101a05790610169575b602090604051908152f35b506020813d602011610198575b8161018360209383612fe4565b81010312610194576020905161015e565b5f80fd5b3d9150610176565b604051903d90823e3d90fd5b80fd5b50346101ac57806003193601126101ac57602060ff60045460c01c166040519015158152f35b50346101ac57806003193601126101ac57602063ffffffff60045460a01c16604051908152f35b50346101ac5761020b36612edf565b91610217949394613591565b60ff1984169384156108ed5760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af49081156106fb5787916108be575b50156108af578386526001602052604086208054909390600881901c6001600160a01b0316156108a05760ff16600581101561088c5761087d5760058401546102cb60068601548261325c565b42101561086e57421061085f578660e06040516102e781612fc8565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201520152808301906101008483031261085b5781601f8501121561085b576040519161010061033b8185612fe4565b839086019182116108575785905b8282106108475750505081519660208301519460408401519460608501519460808101519760a08201519860c08301519260e00151916040519d8e61038d81612fc8565b818152826020820152604081019b8c52606081019a8b526080810193845260a081019c8d5260c0810195865260e0019384525f5160206135b05f395f51905f52119081159161082f575b50610764575f5160206135b05f395f51905f52815110156108205760ff60128c015416600381101561080c576001036107f2578d60018060a01b0360148d0154166020835160246040518094819363650e5fcf60e01b835260048301525afa9182156101a057916107c0575b50801580156107b7575b6107a85760188c015460ff1661078257505160138b015403610773575b8b519760038b0198895403610764575187516104859161327d565b9889151580610752575b61074357908d96959493929160ff60045460c01c166105bc575b505060035460401c6001600160a01b0316919050813b156105b85785936104e660405196879586948594635c73957b60e11b865260048601613355565b03915afa80156105ad57610594575b5050600b91602087015190556105106008850193845461325c565b80935551926105246009820194855461325c565b8094550180545f198114610580576001019055602084519401516040519485526020850152604084015260608301527f7156cf1102d86d870b69eed151febab016e34063d74c714b17538efccfd9d01960803393a36001815580f35b634e487b7160e01b87526011600452602487fd5b8161059e91612fe4565b6105a957865f6104f5565b8680fd5b6040513d84823e3d90fd5b8580fd5b9091929394959650519151905190604051926105d9606085612fe4565b60308452602084019060403683378060801c61072e5760801b90528060801c61071a5760801b60308301528060801c61070657908c959493929160801b6040820152610657602073__$176218451d0b397034f549dcd265d6db5e$__926040518093819263da23b04d60e01b83528460048401526024830190612d74565b0381855af49081156106fb5787916106c6575b50813b156105a9578690602460405180948193634e733ef560e11b835260048301525af49081156106bb5786916106a2575b806104a9565b816106ac91612fe4565b6106b757845f61069c565b8480fd5b6040513d88823e3d90fd5b9650506020863d6020116106f3575b816106e260209383612fe4565b81010312610194578c95515f61066a565b3d91506106d5565b6040513d89823e3d90fd5b6316416a3d60e01b8d52600260045260248dfd5b6316416a3d60e01b8e52600160045260248efd5b506316416a3d60e01b8f5260048f905260248ffd5b6357d18d5360e11b8e5260048efd5b5060088b015460078c0154111561048f565b630b6fac0360e41b8e5260048efd5b635e32eadd60e01b8d5260048dfd5b905060ff60168c01541615610798575b5061046a565b600a8b015411610773575f610792565b635e32eadd60e01b8f5260048ffd5b5043811161044d565b90506020813d6020116107ea575b816107db60209383612fe4565b8101031261019457515f610443565b3d91506107ce565b5160138b01541461046a57635e32eadd60e01b8d5260048dfd5b634e487b7160e01b8f52602160045260248ffd5b635e32eadd60e01b8e5260048efd5b5f5160206135b05f395f51905f52915010155f6103d7565b8135815260209182019101610349565b8980fd5b8780fd5b63e843c5eb60e01b8752600487fd5b63e843c5eb60e01b8852600488fd5b6307a92f1960e51b8752600487fd5b634e487b7160e01b88526021600452602488fd5b634d36eb6960e01b8852600488fd5b632299770d60e11b8652600486fd5b6108e0915060203d6020116108e6575b6108d88183612fe4565b810190613244565b5f61027e565b503d6108ce565b63cbf4a64560e01b8752600487fd5b50346101ac57806003193601126101ac57602063ffffffff600354821c16604051908152f35b50346101ac57806003193601126101ac576004546040516001600160a01b039091168152602090f35b50346101ac5760403660031901126101ac57610965612d56565b60243560ff198216918215610b455760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115610b3a578491610b1b575b5015610b0c578183526001602052604083208054600881901c6001600160a01b03168015610afd573303610aef5760ff601783015460081c1615610aef5760ff166005811015610adb578015159081610acf575b50610ac057600660058201549101908154908315918215610aab575b8215610a8c575b5050610a7d57817f4b5dee969e28ac23e61cb3c24582a7b667358bcc41cd6e31e76001e7624af2429260209255604051908152a280f35b637616640160e01b8452600484fd5b819250610a9c85610aa29361325c565b9261325c565b10155f80610a46565b9150610ab7848261325c565b42101591610a3f565b6307a92f1960e51b8452600484fd5b6003915014155f610a23565b634e487b7160e01b85526021600452602485fd5b6282b42960e81b8552600485fd5b634d36eb6960e01b8652600486fd5b632299770d60e11b8352600483fd5b610b34915060203d6020116108e6576108d88183612fe4565b5f6109cf565b6040513d86823e3d90fd5b63cbf4a64560e01b8452600484fd5b50346101ac5760403660031901126101ac57610b6e612d56565b906024356001600160401b038111610d7557610b8e903690600401612eb2565b9260ff1916908115610d66578183526001602052604083208054600881901c6001600160a01b03168015610afd573303610aef5760ff601783015460101c1615610aef5760ff166005811015610adb578015159081610d5a575b50610ac057600c01936001600160401b038111610d4657610c098554613023565b601f8111610d0b575b508394601f8211600114610c87578185967f970502e3e02e01b5b447e50c390a402175fe6a1e78b7c238e714e28c47622e83959691610c7c575b508260011b905f198460031b1c19161790555b610c76604051928392602084526020840191613335565b0390a280f35b90508301355f610c4c565b8085526020852095601f198316865b818110610cf3575090837f970502e3e02e01b5b447e50c390a402175fe6a1e78b7c238e714e28c47622e839697989210610cda575b5050600182811b019055610c5f565b8401355f19600385901b60f8161c191690555f80610ccb565b85830135895560019098019760209283019201610c96565b610d369086865260208620601f840160051c81019160208510610d3c575b601f0160051c019061331f565b5f610c12565b9091508190610d29565b634e487b7160e01b84526041600452602484fd5b6003915014155f610be8565b63cbf4a64560e01b8352600483fd5b5080fd5b50346101ac576103203660031901126101ac57600560043510156101ac576101203660831901126101ac576001600160401b036101a435116101ac5760a06101a435360360031901126101ac576101c4356001600160401b038111610d7557610de6903690600401612eb2565b9091906040366101e3190112610d755761010036610223190112610d75576004805433808552600260209081526040808720549051636846b35160e11b815260a09490941c63ffffffff16948401949094526024808401929092526001600160401b03909316604483015290949035918560648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4948515610b3a57849561184c575b5060ff19851684526001602052604084205460081c6001600160a01b031661183d5760ff610eab61337f565b16158015611829575b61181a5760ff610ec261338f565b161580156117fd575b6117ee5765ffffffffffff61012435116117df5765ffffffffffff61014435116117d057677fffffffffffffff61016435116117c157677fffffffffffffff61018435116117a3576101243561014435116117b2576101643561018435116117a35760643515611794576080366102a31901126117905760405192610f4f84612fad565b6102a4359283151584036105b8578385526102c4359182151583036105a9578260208701526102e43593841515850361085b578460408801526103043596871515880361178c5780886060610fac9301526101a435600401613463565b87600460ff81351611801561176f575b6117605761088c5760036004351480611751575b611742575f5160206135b05f395f51905f526101e43510801590611728575b61171957604435156116f35760243515611711575b42831061170257426110186044358561325c565b11156116f35760ff19891688526001602052604088209261103b6004358561328a565b6005840155604435600684015560643560078401558254610100600160a81b0319163360081b610100600160a81b03161783556101e435600184015561020435600284015560036101a43560040135101561085b5760405190631044f8bb60e11b825260ff198a1660048301526110bc602483016101a43560040135612dfb565b60843592831515840361085757831515604484015260a4359182151583036116ef57821515606485015260c43560ff811681036116eb5760ff16608485015260e43560ff811681036116eb5760ff1660a4850152610104359360ff85168086036116e75760c48201526101243560e48201526101443561010482015261016435610124820152610184356101448201526101e435610164820152610204356101848201526020816101a48173__$4e0a7d2578f6e358dccb8beebb7c720d9f$__5af49081156116dc578c916116aa575b5060038701556001600160401b038211611696576111ad600c870154613023565b601f8111611663575b508a90601f83116001146115f3576112199392918c91836115e8575b50508160011b915f199060031b1c191617600c8601555b611203600d860194859060ff801983541691151516179055565b835461ff00191690151560081b61ff0016178355565b61122161337f565b90825462ff000064ff0000000063ff00000061123b61338f565b60181b169360201b169360101b169064ffffff00001916171717905561012435600e82015561014435600f8201556101643560108201556101843560118201556012810160ff1981541660ff6101a435600401351617905560246101a435013560138201556014810160018060a01b036112ba60446101a435016132bc565b82546001600160a01b0319169116179055601581016112e36101a43560648101906004016132ed565b906001600160401b0382116115d45781906112fe8454613023565b601f81116115a4575b508990601f831160011461153f578a92611534575b50508160011b915f199060031b1c19161790555b61135761134260846101a435016132e0565b601683019060ff801983541691151516179055565b43600a8201556017810161137e61136c6132d0565b829060ff801983541691151516179055565b61024435801515810361085b57815461ff00191690151560081b61ff001617815561026435801515810361085b57815462ff0000191690151560101b62ff0000161781556102843590811515820361085b579261141a601861146498979694611406856114489861143097509063ff000000825491151560181b169063ff0000001916179055565b0195869060ff801983541691151516179055565b845461ff00191690151560081b61ff0016178455565b825462ff0000191690151560101b62ff000016178255565b9063ff000000825491151560181b169063ff0000001916179055565b60035463ffffffff811663ffffffff811461152057600163ffffffff9101169063ffffffff191617600355338152600260205260408120918254926001600160401b038416936001600160401b03851461150c576001600160401b0360016020960116906001600160401b0319161790556040519133907feefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d160ff1984169180a360ff19168152f35b634e487b7160e01b84526011600452602484fd5b634e487b7160e01b83526011600452602483fd5b013590505f8061131c565b848b5260208b20925090601f1984168b5b81811061158c5750908460019594939210611573575b505050811b019055611330565b01355f19600384901b60f8161c191690555f8080611566565b91936020600181928787013581550195019201611550565b6115ce90858c5260208c20601f850160051c81019160208610610d3c57601f0160051c019061331f565b5f611307565b634e487b7160e01b89526041600452602489fd5b013590505f806111d2565b600c87018c5260208c20918c5b601f198516811061164b57509183916001936112199695601f19811610611632575b505050811b01600c8601556111e9565b01355f19600384901b60f8161c191690555f8080611622565b90926020600181928686013581550194019101611600565b61169090600c88018d5260208d20601f850160051c81019160208610610d3c57601f0160051c019061331f565b5f6111b6565b634e487b7160e01b8b52604160045260248bfd5b90506020813d6020116116d4575b816116c560209383612fe4565b8101031261019457515f61118c565b3d91506116b8565b6040513d8e823e3d90fd5b8c80fd5b8b80fd5b8a80fd5b637616640160e01b8852600488fd5b632ca4094f60e21b8852600488fd5b429250611004565b63208e53e560e11b8852600488fd5b505f5160206135b05f395f51905f52610204351015610fef565b6307a92f1960e51b8852600488fd5b5061175a6132d0565b15610fd0565b6307a92f1960e51b8952600489fd5b50508760043515158015610fbc5750508760036004351415610fbc565b8880fd5b8380fd5b632c45be6f60e01b8452600484fd5b636d403ec760e11b8452600484fd5b63207ea56d60e01b8452600484fd5b634dba795160e01b8452600484fd5b6363f4b4b760e01b8452600484fd5b63b1911d8b60e01b8452600484fd5b632cbdc23160e01b8452600484fd5b5061180661338f565b60ff8061181161337f565b16911611610ecb565b63ac38930b60e01b8452600484fd5b50600860ff61183661337f565b1611610eb4565b635040c4d360e11b8452600484fd5b61186f91955060203d602011611876575b6118678183612fe4565b8101906132a2565b935f610e7f565b503d61185d565b50346101ac57806003193601126101ac57602063ffffffff60035416604051908152f35b50346101ac5760203660031901126101ac5760ff196118be612d56565b1681526001602052604090208054600182016118d990613005565b9160038101549060058101546006820154600783015460088401546009850154600a86015491600b87015493604051958680600c8b01906119199161305b565b036119249088612fe4565b611930600d8a016130dc565b9761193d60128b0161314c565b9960170161194a906131f3565b9a604051809e8e829f60ff8491169061196291612d67565b600160a01b600190039060081c1660208301528051604083015260200151906060015260808d015260a08c015260c08b015260e08a015261010089015261012088015261014087015261016086015261018085016103e090526103e085016119c991612d74565b906101a085016119d891612d98565b8381036102c08501526119ea91612e08565b906102e083016119f991612e57565b0390f35b503461019457611a0c36612edf565b93611a1b959395929192613591565b60ff198116908115611e025760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611d42575f91611de3575b5015611dd457805f52600160205260405f2094855460018060a01b038160081c1615611dc55760ff1696600588101580611d9657600289148015611db9575b611daa57611d9657600188141580611d7b575b611d6c5781850194610120818703126101945785601f82011215610194576101209460405196611b048789612fe4565b878784019182116101945783905b828210611d5c57505050865160038a015403611d4d576004546001600160a01b031690813b15610194575f93611b5e60405196879586948594635c73957b60e11b865260048601613355565b03915afa8015611d4257611d2d575b5060405195611b7c8388612fe4565b600887526020870192601f190136843760015b6009811015611bec578060051b8501515f198201828111611bd8578951811015611bc45760051b890160200152600101611b8f565b634e487b7160e01b88526032600452602488fd5b634e487b7160e01b88526011600452602488fd5b8584848a8560048c8160ff19825416178155018251906001600160401b038211611d1957680100000000000000008211611d19578054828255808310611cfe575b50908492918690885260208820885b838110611ce7575050505060407f6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b91611c7782518092612d67565b60046020820152a26040519060208201906020835251809152604082019390855b818110611cd157505050807fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e9133940390a36001815580f35b8251865260209586019590920191600101611c98565b825181830155879550602090920191600101611c3c565b81885260208820611d1391810190840161331f565b87611c2d565b634e487b7160e01b87526041600452602487fd5b611d3a9194505f90612fe4565b5f925f611b6d565b6040513d5f823e3d90fd5b630b6fac0360e41b5f5260045ffd5b8135815260209182019101611b12565b63e843c5eb60e01b5f5260045ffd5b50611d8f600588015460068901549061325c565b4210611ad4565b634e487b7160e01b5f52602160045260245ffd5b6307a92f1960e51b5f5260045ffd5b50505f60048914611ac1565b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b611dfc915060203d6020116108e6576108d88183612fe4565b5f611a82565b63cbf4a64560e01b5f5260045ffd5b34610194575f36600319011261019457602060405160048152f35b3461019457604036600319011261019457611e45612d56565b6024356001600160401b03811161019457806004019060a060031982360301126101945760ff198316928315611e025760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611d42575f916122eb575b5015611dd457825f52600160205260405f2091825460018060a01b038160081c168015611dc55733036121c85760ff166005811015611d965780151590816122df575b50611daa57611f2b600584015460068501549061325c565b421015611d6c57611f3e601884016131b3565b805115806122d3575b806122c7575b806122bb575b6121c8576012840180549091833560ff83166003821015610194576003811015611d965781036122ac578151158061229a575b6121c85760208201511580612274575b6121c85760608201511580612253575b6121c857604082015115806121d6575b6121c857611fc660ff9286613463565b60ff19909216911617905560248201356013840181905592601481016001600160a01b03611ff6604486016132bc565b82546001600160a01b031916911617905560648301926015820161201a85856132ed565b906001600160401b0382116121b4576120338354613023565b601f8111612184575b505f90601f83116001146120e9576120a660846120d999967f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed29b9996866120b9976120be9b976016975f926120de575b50508160011b915f199060031b1c19161790555b016132e0565b91019060ff801983541691151516179055565b6132ed565b92906040519384938452604060208501526040840191613335565b0390a2005b013590505f8061208c565b601f19831691845f5260205f20925f5b81811061216c57506120d999967f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed29b99966001876120be9b976016976120a6976120b99b60849810612153575b505050811b0190556120a0565b01355f19600384901b60f8161c191690555f8080612146565b919360206001819287870135815501950192016120f9565b6121ae90845f5260205f20601f850160051c81019160208610610d3c57601f0160051c019061331f565b8961203c565b634e487b7160e01b5f52604160045260245ffd5b6282b42960e81b5f5260045ffd5b506121e460648701866132ed565b6001600160401b0381116121b4576040519061220a601f8201601f191660200183612fe4565b8082526020820192368282011161019457815f926020928637830101525190206040516122458161223e8160158d0161305b565b0382612fe4565b602081519101201415611fb6565b50612260608487016132e0565b60ff60168901541615159015151415611fa6565b50612281604487016132bc565b60148801546001600160a01b0391821691161415611f96565b50601387015460248701351415611f86565b63f37f7b5d60e01b5f5260045ffd5b50606081015115611f53565b50604081015115611f4d565b50602081015115611f47565b60039150141585611f13565b612304915060203d6020116108e6576108d88183612fe4565b84611ed0565b34610194576020366003190112610194576004356001600160a01b0381169081810361019457600480545f93845260026020908152604094859020549451636846b35160e11b815260a09290921c63ffffffff16928201929092526001600160a01b039290921660248301526001600160401b039092166044820152908160648173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af48015611d42576020915f916123bf575b506040519060ff19168152f35b6123d69150823d8411611876576118678183612fe4565b826123b2565b34610194576020366003190112610194576004356001600160a01b03811690819003610194575f52600260205260206001600160401b0360405f205416604051908152f35b346101945760203660031901126101945761243a612d56565b60405161244681612f3f565b5f81525f602082015260405161245b81612f5b565b5f81525f602082015260408201525f6060820152606060808201525f60a08201525f60c08201525f60e08201525f6101008201525f6101208201525f6101408201525f61016082015260606101808201526040516124b881612f76565b5f81525f60208201525f60408201525f60608201525f60808201525f60a08201525f60c08201525f60e08201525f6101008201526101a08201526040516124fe81612f92565b5f81525f60208201525f60408201526060808201525f60808201526101c08201526101e06040519161252f83612f92565b5f83525f60208401525f60408401525f606084015260405161255081612fad565b5f81525f60208201525f60408201525f60608201526080840152015260ff19165f52600160205260405f2060405161258781612f3f565b81549060ff82166005811015611d9657815260089190911c6001600160a01b031660208201908152906125bc60018401613005565b60408201908152600384015460608301908152600485019160405180938491602082549182815201915f5260205f20905f5b8181106127cd5750505060209291612607910385612fe4565b608085019384526126a4601760058901549860a08801998a52600681015460c0890152600781015460e089015260088101546101008901526009810154610120890152600a810154610140890152600b8101546101608901526040516126748161223e81600c860161305b565b610180890152612686600d82016130dc565b6101a08901526126986012820161314c565b6101c0890152016131f3565b6101e0860152604051958287526126be8388018751612d67565b516001600160a01b031660408701525180516060870152015160808501525160a08401525161040060c08401528051610420840181905261044084019491602001905f5b8181106127b7575050506101e06127a961277f85966119f9945160e088015260c086015161010088015260e08601516101208801526101008601516101408801526101208601516101608801526101408601516101808801526101608601516101a0880152610180860151601f19888303016101c0890152612d74565b6127916101a086015184880190612d98565b6101c0850151868203601f1901610300880152612e08565b920151610320840190612e57565b8251875260209687019690920191600101612702565b82548452879450602090930192600192830192016125ee565b34610194575f366003190112610194576003546040805163233ace1160e01b81529160209183916004918391901c6001600160a01b03165afa8015611d42575f9061016957602090604051908152f35b34610194575f3660031901126101945760206040515f8152f35b34610194575f366003190112610194576003546040805191901c6001600160a01b03168152602090f35b3461019457604036600319011261019457612893612d56565b60243560058110156101945760ff198216918215611e025760048054604051632145bb8f60e01b815260ff199093169183019190915260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611d42575f91612a65575b5015611dd457600460ff821611611daa575f82815260016020526040902080549091600882901c6001600160a01b03168015611dc55733036121c857601783019160ff835416156121c85760ff169161295f828461339f565b15611daa57600182146129b9575b50906129b6826129a06040947f6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b9661328a565b6129ac84518094612d67565b6020830190612d67565ba2005b6005840180549194914210612a5c576129d390544261327d565b6006820194855482036129e9575b50509261296d565b549093929060081c60ff16156121c8576129a0816129b693887f4b5dee969e28ac23e61cb3c24582a7b667358bcc41cd6e31e76001e7624af24260208960409a7f6adae7b36d9ac37d13a21ea4539a2a7838d8820b3807922ccd4a0e3b114ec71b9c558a51908152a293509394506129e1565b4290555f6129d3565b612a7e915060203d6020116108e6576108d88183612fe4565b83612906565b346101945760203660031901126101945760ff19612aa0612d56565b165f52600160205260405f20604051612ab881612f3f565b815460ff81166005811015611d9657825260081c6001600160a01b03166020820152612ae660018301613005565b604082015260038201546060820152600482016040519081602082549182815201915f5260205f20905f5b818110612be1576020612bd9888888612b2c818a0382612fe4565b608082015260058201549060a081019182526101e0612bcd601760068601549560c08501968752600781015460e086015260088101546101008601526009810154610120860152600a810154610140860152600b810154610160860152604051612b9d8161223e81600c860161305b565b610180860152612baf600d82016130dc565b6101a0860152612bc16012820161314c565b6101c0860152016131f3565b9101525190519061325c565b604051908152f35b8254845260209093019260019283019201612b11565b3461019457604036600319011261019457612c10612d56565b60ff19811691602435918315611e025760048054632145bb8f60e01b845260ff199092169083015260a01c63ffffffff1660248201526020818060448101038173__$3dc519e8cd588db0d942d0eb61ec77f7fa$__5af4908115611d42575f91612d37575b5015611dd4575f8281526001602052604090208054600881901c6001600160a01b03168015611dc55733036121c85760ff601783015460181c16156121c85760ff166005811015611d96578015159081612d2b575b50611daa5781158015612d1e575b612d0f57817fc12d00e36fb3763916c443453e211136b31b621e8f92b16e7eec16bf3919ae0f9260076020930155604051908152a2005b632c45be6f60e01b5f5260045ffd5b5060088101548210612cd8565b60039150141584612cca565b612d50915060203d6020116108e6576108d88183612fe4565b83612c75565b6004359060ff198216820361019457565b906005821015611d965752565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b610100809180511515845260208101511515602085015260ff604082015116604085015260ff606082015116606085015260ff608082015116608085015260a081015160a085015260c081015160c085015260e081015160e08501520151910152565b906003821015611d965752565b90612e14818351612dfb565b6020820151602082015260018060a01b036040830151166040820152608080612e4c606085015160a0606086015260a0850190612d74565b930151151591015290565b6060608060e0928051151585526020810151151560208601526040810151151560408601528281015115158386015201518051151560808501526020810151151560a08501526040810151151560c085015201511515910152565b9181601f84011215610194578235916001600160401b038311610194576020838186019501011161019457565b9060606003198301126101945760043560ff198116810361019457916024356001600160401b0381116101945781612f1991600401612eb2565b92909291604435906001600160401b03821161019457612f3b91600401612eb2565b9091565b61020081019081106001600160401b038211176121b457604052565b604081019081106001600160401b038211176121b457604052565b61012081019081106001600160401b038211176121b457604052565b60a081019081106001600160401b038211176121b457604052565b608081019081106001600160401b038211176121b457604052565b61010081019081106001600160401b038211176121b457604052565b90601f801991011681019081106001600160401b038211176121b457604052565b9060405161301281612f5b565b602060018294805484520154910152565b90600182811c92168015613051575b602083101461303d57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613032565b5f929181549161306a83613023565b80835292600181169081156130bf575060011461308657505050565b5f9081526020812093945091925b8383106130a5575060209250010190565b600181602092949394548385870101520191019190613094565b915050602093945060ff929192191683830152151560051b010190565b906040516130e981612f76565b6101006004829460ff815481811615158652818160081c1615156020870152818160101c166040870152818160181c16606087015260201c166080850152600181015460a0850152600281015460c0850152600381015460e08501520154910152565b9060405161315981612f92565b809260ff815416906003821015611d9657600460ff9160809385526001810154602086015260018060a01b0360028201541660408601526040516131a48161223e816003860161305b565b60608601520154161515910152565b906040516131c081612fad565b606060ff82945481811615158452818160081c1615156020850152818160101c161515604085015260181c161515910152565b9060405161320081612f92565b608061323f6001839560ff815481811615158752818160081c1615156020880152818160101c161515604088015260181c1615156060860152016131b3565b910152565b90816020910312610194575180151581036101945790565b9190820180921161326957565b634e487b7160e01b5f52601160045260245ffd5b9190820391821161326957565b906005811015611d965760ff80198354169116179055565b90816020910312610194575160ff19811681036101945790565b356001600160a01b03811681036101945790565b6102243580151581036101945790565b3580151581036101945790565b903590601e198136030182121561019457018035906001600160401b0382116101945760200191813603831361019457565b81811061332a575050565b5f815560010161331f565b908060209392818452848401375f828201840152601f01601f1916010190565b929061336e9061337c9593604086526040860191613335565b926020818503910152613335565b90565b60c43560ff811681036101945790565b60e43560ff811681036101945790565b906005821015611d965760058110159182611d965780821461345c575f600282148015613450575b818115613440575b5061343857611d9657801561341f576003146133eb5750505f90565b81611d96578015918215613412575b821561340557505090565b909150611d965760011490565b506002811491505f6133fa565b506003811491505f821561341257821561340557505090565b505050505f90565b9050611d965760018214816133cf565b50505f600482146133c7565b5050505f90565b908135600381101561019457600103613522576001600160a01b0361348a604084016132bc565b16156135135780511580613507575b6134c35780511590816134f9575b506134ea576020905b013580151590816134d2575b506134c357565b635e32eadd60e01b5f5260045ffd5b5f5160206135b05f395f51905f52915010155f6134bc565b63f545b7bf60e01b5f5260045ffd5b60609150015115155f6134a7565b50602082013515613499565b63562e597160e11b5f5260045ffd5b60208101516134ea57606001516134ea576001600160a01b03613547604083016132bc565b166134ea5761355960608201826132ed565b905015613582576020810135156134c357613576608082016132e0565b6134ea576020906134b0565b630f8b932160e11b5f5260045ffd5b60025f54146135a05760025f55565b633ee5aeb560e01b5f5260045ffdfe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122090d4e8e0ca7aa74f5bf0aad72bbc9c58189852f7fc5ebe4039fa5bd9683c60ae64736f6c634300081c0033",
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),(bool,bool,bool,bool,(bool,bool,bool,bool))))
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),(bool,bool,bool,bool,(bool,bool,bool,bool))))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [31]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x59d821c4.
//
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),(bool,bool,bool,bool,(bool,bool,bool,bool))))
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
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, (bool,bool,bool,bool,(bool,bool,bool,bool)) paramsMod)
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
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, (bool,bool,bool,bool,(bool,bool,bool,bool)) paramsMod)
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
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, (bool,bool,bool,bool,(bool,bool,bool,bool)) paramsMod)
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

// NewProcess is a paid mutator transaction binding the contract method 0x851c3453.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (bool,bool,bool,bool,(bool,bool,bool,bool)) paramsMod) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, paramsMod DAVINCITypesParamsMod) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, paramsMod)
}

// NewProcess is a paid mutator transaction binding the contract method 0x851c3453.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (bool,bool,bool,bool,(bool,bool,bool,bool)) paramsMod) returns(bytes31)
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, paramsMod DAVINCITypesParamsMod) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, paramsMod)
}

// NewProcess is a paid mutator transaction binding the contract method 0x851c3453.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (bool,bool,bool,bool,(bool,bool,bool,bool)) paramsMod) returns(bytes31)
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
