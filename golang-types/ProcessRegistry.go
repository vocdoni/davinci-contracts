// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
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

// DAVINCITypesDKGParams is an auto generated low-level Go binding around an user-defined struct.
type DAVINCITypesDKGParams struct {
	Mode    uint8
	EpochId [12]byte
	OrgPKx  *big.Int
	OrgPKy  *big.Int
	PopAx   *big.Int
	PopAy   *big.Int
	PopZ    *big.Int
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
	LatestStateRoot       [32]byte
	Result                []*big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	MetadataHash          [32]byte
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
	Grace                 uint32
	LastVoteAt            uint64
}

// IL1BlobVerifierL1BlobProof is an auto generated low-level Go binding around an user-defined struct.
type IL1BlobVerifierL1BlobProof struct {
	Headers [][]byte
	Txs     []IL1BlobVerifierL1TxProof
}

// IL1BlobVerifierL1TxProof is an auto generated low-level Go binding around an user-defined struct.
type IL1BlobVerifierL1TxProof struct {
	HeaderIndex *big.Int
	Key         []byte
	Proof       [][]byte
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_ziskVerifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_batchProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_resultsProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rootCVadcopFinal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_ballotVKHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_dkgManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_councilManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_defaultGrace\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceFloor\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceCeil\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceMaxTotal\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_noticeMin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_l1BlobVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlobCountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CircuitFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CouncilDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DKGDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DecryptionNotOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyTransition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GraceOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccumulator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobCommitmentLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidBlobOpening\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobsDigest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDKGParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncryptionKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGrace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInclusionProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKZGProofLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOccupiedBefore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxPossibleResultCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"MissingBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoBlobs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsAlreadyRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtMaxLevelsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongDAMode\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"}],\"name\":\"ProcessGraceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"ProcessMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nBlobs\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"l1BlockNumbers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"versionedHashes\",\"type\":\"bytes32[]\"}],\"name\":\"ProcessStateTransitionedL1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"firstIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"name\":\"ResultsDecryptionRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"aidFor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotVKHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"councilAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGrace\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dkgAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"finalizeResultsFromDKG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"}],\"name\":\"genesisRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastVoteAt\",\"type\":\"uint64\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessGraceEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceCeil\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceFloor\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceMaxTotal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BlobVerifier\",\"outputs\":[{\"internalType\":\"contractIL1BlobVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint256\",\"name\":\"orgPKx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orgPKy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popZ\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.DKGParams\",\"name\":\"dkg\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noticeMin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastVoteAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256[64]\",\"name\":\"accumulator\",\"type\":\"uint256[64]\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"requestResultsDecryption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultsProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"sk\",\"type\":\"uint256\"}],\"name\":\"revealProcessKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootCVadcopFinal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"}],\"name\":\"setProcessGrace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"setProcessMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"commitments\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"kzgProofs\",\"type\":\"bytes[]\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"commitments\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"kzgProofs\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"headers\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"headerIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"}],\"internalType\":\"structIL1BlobVerifier.L1TxProof[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIL1BlobVerifier.L1BlobProof\",\"name\":\"l1\",\"type\":\"tuple\"}],\"name\":\"submitStateTransitionL1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ziskVerifier\",\"outputs\":[{\"internalType\":\"contractIZiskVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610220806040523461046b576101c081617dee8038038091610021828561046f565b83398101031261046b5761003481610492565b90610041602082016104a3565b916040820151606083015160808401519060a08501519261006460c087016104a3565b9561007160e082016104a3565b9761007f6101008301610492565b9161008d6101208201610492565b61009a6101408301610492565b6100a76101608401610492565b916100c26101a06100bb6101808701610492565b95016104a3565b60015f556001600160a01b039095169586158015610463575b801561045b575b8015610452575b801561044a575b8015610442575b6104335763ffffffff82168015908115610423575b50801561040e575b80156103f9575b80156103eb575b6103dc5761016052610180526101a0526101c0526101e05260018060a01b03166102005260805260a05260c05260e0526101005260035467ffffffff000000006bffffffff0000000000000000604051602081019063ffffffff60e01b8660e01b1682523060601b60248201526018815261019e60388261046f565b51902060401b169260201b1690640100000000600160601b031916171760035560018060a01b031680155f1461039b57505f5b610120526001600160a01b03168061033b57505f5b61014052604051615e8e90816104b882396080518181816111b501528181611529015281816126420152612c5a015260a0518181816126bd01528181612cce01526141fb015260c05181818161159a015261455a015260e05181818161157901528181611d9f0152818161269c0152612cad015261010051818181610c9a01528181612e27015261341101526101205181818161097701528181611806015281816118df015281816139fd01528181613a5b01526151ea015261014051818181610e0d015281816138b401528181613918015281816151bc0152615267015261016051818181610a31015261362a0152610180518181816110a40152611de101526101a0518181816111360152612dec01526101c051818181612373015261497401526101e0518181816109c50152610f740152610200518181816124a1015281816125e301528181612bba0152612e5f0152f35b60405190610a448083016001600160401b038111848210176103875760209284926173aa843981520301905ff0801561037c576001600160a01b03166101e6565b6040513d5f823e3d90fd5b634e487b7160e01b5f52604160045260245ffd5b604051906110648083016001600160401b03811184821017610387576020928492616346843981520301905ff0801561037c576001600160a01b03166101d1565b63795ee5af60e01b5f5260045ffd5b5063ffffffff851615610122565b5063ffffffff841663ffffffff84161161011b565b5063ffffffff831663ffffffff821611610114565b905063ffffffff8216105f61010c565b6306f9b90760e01b5f5260045ffd5b508a156100f7565b5089156100f0565b508888146100e9565b5088156100e2565b5087156100db565b5f80fd5b601f909101601f19168101906001600160401b0382119082101761038757604052565b519063ffffffff8216820361046b57565b51906001600160a01b038216820361046b5756fe6105c06040526004361015610012575f80fd5b5f6102c0525f3560e01c8063026cdee81461406757806304ed00fa14613e5e578063082b642e14613c5b57806308c0fdd314612e8e5780630c39fa3614612e4a5780630e2ebcf714612e105780631542bbe214612dd05780631fdf344914612ade5780633ea4ee4114612aa657806346c15da8146129e55780634c0acc561461117257806353fb626d14612397578063549d59951461235457806359d821c414611e055780635ff5f98114611dc25780636211533814611d8557806362fa11fc14611d3a57806368141f2c14611cba5780636c7aff7f14611991578063702574b3146118c457806372c628ef146118a657806373417705146117bd578063766422e014611440578063784df74a146111e45780637f64b72f1461119d578063848df54014611177578063946544bf146111725780639a0377891461101e5780639b46499414610e3c578063aa240221146101d7578063abaab13a14610df5578063adc879e914610dcd578063bf74291e14610a55578063c8f0582f14610a12578063cddf08bc146109e9578063d4138a20146109a6578063e16d5b7c1461095f578063e965eead14610316578063f1431097146101dc5763f9aa4499146101d7575f80fd5b614543565b3461030f57602036600319011261030f576101f5614176565b6101fd6149f8565b61020681614ea5565b601881015460ff811660048110156102f557156102e057815460ff1661022b8161421e565b600281149081156102cc575b506102b75761024582614959565b42106102a25760901c60ff161561028d5761025f8161521b565b156102785761026d916152d6565b6102c0516001815580f35b6314badd7360e31b6102c0515260046102c051fd5b630e0d4dc560e41b6102c0515260046102c051fd5b63611f72f360e11b6102c0515260046102c051fd5b6307a92f1960e51b6102c0515260046102c051fd5b600491506102d98161421e565b1484610237565b6365b75c3960e01b6102c0515260046102c051fd5b634e487b7160e01b6102c05152602160045260246102c051fd5b6102c05180fd5b3461030f5761084036600319011261030f57610330614176565b366108241161030f57610824356001600160401b03811161030f576103599036906004016141b4565b6103616149f8565b61036a83614ea5565b91601883015491600460ff841610156102f55760ff8316156102e05760ff8360901c1661094a57835460ff8116929091906103a48461421e565b600284148015610937575b6102b7576103bc8461421e565b600184141591828061091c575b610907576103d687614959565b42106102a2576102c0515b604081106108e15750602060405181810161080060248237610800825261040a6108208361436e565b60405191518091835e8101906102c0518252806102c05192039060025afa1561068a57610440916102c0515160038901546150e4565b156108cc5760ff60901b198416600160901b1760188601556104618361421e565b61088d575b505060ff600e83015460081c169061047d82614749565b9161048b604051938461436e565b808352601f1961049a82614749565b016102c0515b81811061086b5750506102c05190819081905b8082106106de5750506102c0518215949093908515610598575b50509061ffff60ff926018870154908260801b9060801b16908260681b8660681b169064ffffffffff60681b1916178460781b8460781b161717938460188801556019870154604051956001600160601b0360a01b9060981b16865260208601521660408401521660608201527fdca6075f07367349a836825d3ee7c35c204d2e727ae81a5b7a48aca95b9c270b608060ff19861692a280610589575b610579576102c0516001815580f35b610582916152d6565b808061026d565b506105938161521b565b61056a565b838152919350906001600160a01b036105b360ff84166151aa565b601988015460405163574bbf6f60e11b815260ff60901b19909516600160901b1760981b6001600160a01b03191660048601526024850152606060448501528251606485018190526102c051929091169284926084840192602090920191905b8181106106985750505091818060209403916102c051905af190811561068a576102c05191610649575b509161ffff60ff6104cd565b90506020813d602011610682575b816106646020938361436e565b8101031261030f57519061ffff8216820361030f579061ffff61063d565b3d9150610657565b6040513d6102c051823e3d90fd5b918094509290925181906102c051915b600483106106c85750505060206080600192019401910191859392610613565b60208060019284518152019201920191906106a8565b90926001600160fe1b03841684036107cf576106fc8460021b614817565b3590600285901b600181019081106107cf5761071790614817565b356102c051600287811b0191828860021b116107cf5761073683614817565b351580610840575b851580610836575b610818576108035760405194608086018681106001600160401b038211176107e957604052855260208501526107cf5761077f90614817565b356040830152600285901b60038101919082106107cf576001926107a56107c693614817565b3560608201526107b5828a614803565b526107c08189614803565b50614829565b935b01906104b3565b634e487b7160e01b6102c05152601160045260246102c051fd5b634e487b7160e01b6102c05152604160045260246102c051fd5b632a23591560e21b6102c0515260046102c051fd5b9397969450505050156108035760019061ffff82851b1617926107c8565b5060018214610746565b5090506102c0519060038860021b01808960021b116107cf57610864600191614817565b351461073e565b60405160209190608061087e818361436e565b368237828288010152016104a0565b60ff19166001178355604051906108a38161421e565b81526001602082015260ff198416905f516020615e195f395f51905f5290604090a28380610466565b6303cd656760e61b6102c0515260046102c051fd5b5f516020615e395f395f51905f526108f882614817565b351015610803576001016103e1565b63e843c5eb60e01b6102c0515260046102c051fd5b50610930600588015460068901549061458b565b42106103c9565b506109418461421e565b600484146103af565b636c8ae94b60e11b6102c0515260046102c051fd5b3461030f576102c05136600319011261030f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461030f576102c05136600319011261030f57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461030f576102c05136600319011261030f57602063ffffffff60035460401c16604051908152f35b3461030f576102c05136600319011261030f57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461030f5761018036600319011261030f57610a6f614176565b61010036602319011261030f5760403661012319011261030f5761016435600581101561030f5760405190610aa38261431c565b61012435825261014435602083019081526040519290610ac460e08561436e565b6006845260c0948536602087013760405195610ae160e08861436e565b6006875236602088013760081c610af786614792565b526102c051610b0585614792565b5260643560ff8116929083810361030f576044359360ff85169081860361030f578110610db85760a43565ffffffffffff8111610da35760c4359165ffffffffffff8311610d8e5760e43593677fffffffffffffff8511610d79576101043597677fffffffffffffff8911610d645750602435801515810361030f5715610d5a576001905b6084359260ff8416840361030f5761ff0062ff00006301fe000060209c60b81b9960791b9860491b9760191b9660111b169460101b169260081b1617171717171717610bd5876147b3565b526002610be1866147b3565b525190519060405191838301918252604083015260408252610c0460608361436e565b60405191518091835e8101906102c0518252806102c05192039060025afa1561068a576102c05151610c35846147c3565b526003610c41836147c3565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae278610c6c846147d3565b526004610c78836147d3565b52610c828161421e565b610c8b836147e3565b526006610c97826147e3565b527f0000000000000000000000000000000000000000000000000000000000000000610cc2836147f3565b526007610cce826147f3565b528051825103610d4557610ce28151614760565b6102c0515b8251811015610d2c5780610d1b6001600160401b03610d0860019487614803565b5116610d148388614803565b5190615bbf565b610d258285614803565b5201610ce7565b6020610d3d84846102c05191615c52565b604051908152f35b63d088249360e01b6102c0515260046102c051fd5b6102c05190610b8a565b63dd6f54df60e01b6102c0515260046102c051fd5b63271fb80560e01b6102c0515260046102c051fd5b63871a7fa360e01b6102c0515260046102c051fd5b63481eb79f60e01b6102c0515260046102c051fd5b632cbdc23160e01b6102c0515260046102c051fd5b3461030f576102c05136600319011261030f57602063ffffffff600354821c16604051908152f35b3461030f576102c05136600319011261030f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461030f57604036600319011261030f57610e55614176565b60243560ff1982169182156110095763ffffffff808060035460401c16169160401c1603610ff4576102c08051839052600160205251604090208054600881901c6001600160a01b03168015610fdf573303610fcb5760ff16610eb78161421e565b8015159081610fb6575b506102b75760066005820154910190610edb82548261458b565b90428211156109075783610eee9161458b565b8315918215610fab575b8215610fa1575b8215610f56575b5050610f4157817f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba9260209255604051908152a26102c05180f35b637616640160e01b6102c0515260046102c051fd5b8110915081610f68575b508480610f06565b9050610f9a63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000164261458b565b1184610f60565b8181149250610eff565b428211159250610ef8565b60039150610fc38161421e565b141584610ec1565b6282b42960e81b6102c0515260046102c051fd5b634d36eb6960e01b6102c0515260046102c051fd5b632299770d60e11b6102c0515260046102c051fd5b63cbf4a64560e01b6102c0515260046102c051fd5b3461030f57604036600319011261030f57611037614176565b6024359063ffffffff821680920361030f5761105281614ea5565b805433600882901c6001600160a01b031603610fcb5760ff166110748161421e565b801515908161115d575b506102b757611096600582015460068301549061458b565b4210156109075763ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168310801561112e575b61111957601a01805463ffffffff19168317905560405191825260ff1916907fdf161c6af27d090672f982bb8002da0e7f2a530ffbd779e02ab9eba9397e51f190602090a26102c05180f35b63795ee5af60e01b6102c0515260046102c051fd5b5063ffffffff7f00000000000000000000000000000000000000000000000000000000000000001683116110cd565b6003915061116a8161421e565b14158461107e565b6141e4565b3461030f576102c05136600319011261030f57602063ffffffff60035416604051908152f35b3461030f576102c05136600319011261030f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461030f57602036600319011261030f5760ff19611200614176565b166102c0515260016020526102c051604090208054600182016112229061438f565b6101c0526003820154610200526005820154600683015460078401549060088501546009860154600a870154600b88015491600c8901611261906143e5565b93600d8a015496600e8b0161127590614485565b9661128260138d016144e7565b9960188d01546101a05260198d01549c601a01549b6040516101e05260ff81166112ab9061421e565b60ff81166101e05152600160a01b600190039060081c166101e051602001526101c051516101e051604001526101c051602001516101e05160600152610200516101e051608001526101e05160a001526101e05160c001526101e05160e001526101e05161010001526101e05161012001526101e05161014001526101e05161016001526101e0516101800161040090526101e0516104000161134d91614228565b916101e0516101a001526101e0516101c0016113689161424c565b6101e05181036101e0516102c00152611380916142a2565b916101e0516102e0016101a05160ff169061139a916142f3565b6001600160601b0360a01b6101a05160981b166101e05161030001526101a05160681c61ffff166101e05161032001526101a05160781c60ff166101e05161034001526101a05160801c61ffff166101e05161036001526101a05160901c60ff1615156101e05161038001526101e0516103a0015263ffffffff81166101e0516103c0015260201c6001600160401b03166101e0516103e001526101e05190036101e051f35b3461030f57606036600319011261030f57611459614176565b6024356001600160401b03811161030f57611478903690600401614187565b90916044356001600160401b03811161030f57611499903690600401614187565b926114a26149f8565b6114ab83614ea5565b9160ff60188401541660048110156102f5576102e057825460ff8116959092906114d48761421e565b6002871480156117aa575b6102b7576114ec8761421e565b60018714158061178f575b6109075761150485614959565b42106102a257611514828961500b565b61151d8861509f565b60038601540361177a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561030f576115c29389936040519586948593849363be98686160e01b85526102c051987f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801614713565b03915afa801561068a5761175f575b5060ff600e83015460081c16946115e786614760565b956102c0515b8181106116f65750505060ff191660049081178255845191016001600160401b0382116107e957600160401b82116107e95780548282558083106116d4575b5060208501906102c0515260206102c051206102c0515b8381106116c057866040875f516020615e195f395f51905f528860ff19169283928151906116708161421e565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e60405160208152806116b2339560208301906146c0565b0390a36102c0516001815580f35b600190602084519401938184015501611643565b6116f090826102c051528360206102c0512091820191016146aa565b8561162c565b600181901b906001600160ff1b03811681036107cf5781600a019182600a116107cf57600b61172d8460031b87013560c01c615b72565b91019283106107cf5761174a60019360031b86013560c01c615b72565b60201b17611758828b614803565b52016115ed565b6102c05161176c9161436e565b6102c05161030f57856115d1565b630b6fac0360e41b6102c0515260046102c051fd5b506117a3600586015460068701549061458b565b42106114f7565b506117b48761421e565b600487146114df565b3461030f57604036600319011261030f576117e66117d9614176565b6117e16149f8565b614ea5565b60188101549060ff821660048110156102f5576002036102e057601901547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561030f576040519263193f942b60e21b84526001600160601b0360a01b9060981b16600484015260248301526024356044830152816064816102c051936102c051905af1801561068a5761188b576102c0516001815580f35b6102c0516118989161436e565b6102c05161030f578061026d565b3461030f576102c05136600319011261030f57602060405160048152f35b3461030f57602036600319011261030f576118dd614176565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690811561197d5760209060246040518094819363702574b360e01b835260ff191660048301525afa801561068a576102c0519061194a575b602090604051908152f35b506020813d602011611975575b816119646020938361436e565b8101031261030f576020905161193f565b3d9150611957565b6203eb8f60e01b6102c0515260046102c051fd5b3461030f57604036600319011261030f576119aa614176565b6024356001600160401b03811161030f578060040160a0600319833603011261030f576119d683614ea5565b8054919033600884901c6001600160a01b031603610fcb57601381015460029060ff16611a028161421e565b03611ca5578135600581101561030f57600290611a1e8161421e565b03611c90576001600160a01b03611a3760448601614664565b16611c7b576024840135938415611c665760640192611a568484614678565b905015611c515760ff16611a698161421e565b8015159081611c3c575b506102b757611a8b600582015460068301549061458b565b421015610907578360148201556016611aa48484614678565b91909201916001600160401b0382116107e957611ac183546143ad565b601f8111611bfb575b506102c05190601f8311600114611b62579282611b4d9896937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2989693611b2d966102c05192611b57575b50508160011b915f199060031b1c1916179055614678565b949060405193849384526040602085015260ff19169560408401916146f3565b0390a26102c05180f35b013590508a80611b15565b601f19831691846102c0515260206102c05120926102c0515b818110611be35750937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2989693611b2d969360019383611b4d9d9b9810611bca575b505050811b019055614678565b01355f19600384901b60f8161c191690558a8080611bbd565b91936020600181928787013581550195019201611b7b565b611c2c90846102c0515260206102c05120601f850160051c81019160208610611c32575b601f0160051c01906146aa565b87611aca565b9091508190611c1f565b60039150611c498161421e565b141586611a73565b630f8b932160e11b6102c0515260046102c051fd5b635e32eadd60e01b6102c0515260046102c051fd5b63562e597160e11b6102c0515260046102c051fd5b63f37f7b5d60e01b6102c0515260046102c051fd5b63050b77c760e21b6102c0515260046102c051fd5b3461030f57602036600319011261030f576004356001600160a01b0381169081810361030f576003546102c0805193909352600260209081529251604090205466ffffffffffffff1660589290921b600160581b600160f81b0316600891821c63ffffffff60381b161791909117901b60ff19166040519060ff19168152f35b3461030f57602036600319011261030f576004356001600160a01b0381169081900361030f576102c05152600260205260206001600160401b0360406102c051205416604051908152f35b3461030f576102c05136600319011261030f5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b3461030f576102c05136600319011261030f57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461030f57602036600319011261030f57611e1e614176565b604051611e2a81614300565b6102c05181526102c0516020820152604051611e458161431c565b6102c05181526102c051602082015260408201526102c0516060820152606060808201526102c05160a08201526102c05160c08201526102c05160e08201526102c0516101008201526102c0516101208201526102c0516101408201526102c05161016082015260606101808201526102c0516101a0820152604051611eca81614337565b6102c05181526102c05160208201526102c05160408201526102c05160608201526102c05160808201526102c05160a08201526102c05160c08201526102c05160e08201526101c0820152604051611f2181614353565b6102c05181526102c05160208201526102c05160408201526060808201526102c05160808201526101e08201526102c0516102008201526102c0516102208201526102c0516102408201526102c0516102608201526102c0516102808201526102c0516102a08201526102c0516102c08201526102c0516102e08201526103006102c05191015260ff19166102c05152600160205260406102c0512060405190611fca82614300565b805460ff8116611fd98161421e565b835260081c6001600160a01b03166020830152611ff86001820161438f565b60408301526003810154606083015260048101604051808260208294549384815201906102c0515260206102c05120926102c0515b81811061233b5750506120429250038261436e565b6080830152600581015460a0830152600681015460c0830152600781015460e083015260088101546101008301526009810154610120830152600a810154610140830152600b81015461016083015261209d600c82016143e5565b610180830152600d8101546101a08301526120ba600e8201614485565b6101c08301526120cc601382016144e7565b6101e08301526018810154600460ff821610156102f5576001600160401b039160ff8281601a94166102008701526001600160601b0360a01b8160981b1661022087015261ffff8160681c16610240870152818160781c1661026087015261ffff8160801c1661028087015260901c1615156102a085015260198101546102c0850152015463ffffffff81166102e084015260201c166103008201526040516020815261044081019180516121808161421e565b602083015260018060a01b036020820151166040830152602060408201518051606085015201516080830152606081015160a083015260808101519261042060c08401528351809152602061046084019401906102c0515b818110612325575050506001600160401b0361030061229361225d859660a086015160e088015260c086015161010088015260e08601516101208801526101008601516101408801526101208601516101608801526101408601516101808801526101608601516101a0880152610180860151601f19888303016101c0890152614228565b6101a08501516101e087015261227d6101c086015161020088019061424c565b6101e0850151868203601f1901848801526142a2565b926122a86102008201516103208701906142f3565b6102208101516001600160a01b03191661034086015261024081015161ffff90811661036087015261026082015160ff16610380870152610280820151166103a08601526102a081015115156103c08601526102c08101516103e08601526102e081015163ffffffff166104008601520151166104208301520390f35b82518652602095860195909201916001016121d8565b845483526001948501948694506020909301920161202d565b3461030f576102c05136600319011261030f57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461030f5760e036600319011261030f576123b0614176565b6102a0526024356001600160401b03811161030f576123d3903690600401614187565b61046052610440526044356001600160401b03811161030f576123fa903690600401614187565b61038052610360526064356001600160401b03811161030f576124219036906004016141b4565b61040052610420526084356001600160401b03811161030f576124489036906004016141b4565b6103a0526103e05260a4356001600160401b03811161030f5761246f9036906004016141b4565b6103c052610340526001600160401b0360c4351161030f57604060c4353603600319011261030f5761249f6149f8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156129d0576124e261046051610440516102a051614c1f565b61022052610240526102e052610300526103205261251e6103c051610340516103a0516103e0516104005161042051610460516104405161573d565b60a0526125496103c051610340516103a0516103e051610400516104205161046051610440516157d8565b604051610280818152638051f3c160e01b90915251602060049182015261258f906125789060c4350180614f47565b604060246102805101526064610280510191614fac565b6125a3602460c4350160c435600401614f47565b602319610280518403016044610280510152808352602083019060208160051b85010193836102c051915b838310612957576102c05161028051808903817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa61026052610260511561068a576102c051610180526102c0516101605261026051612819575b60a051610180515103612804577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561030f576040519063be98686160e01b8252816102c0519181806126e5610380516103605161046051610440517f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801614713565b03915afa801561068a576127e9575b5061272d610180516103c051610340516103a0516103e05161040051610420516127206104405161506a565b610300516102a05161590a565b6127596104005161022051610240516102e0516103005161046051610440516102a05161032051614dc0565b6040516040815261277060408201610160516146c0565b9080820360208201526020610180515192838152019160206101805101906102c0515b8181106127d3576102a05160ff19167f4707050fda6801db89adf05b220e5793cac2e1ce5346f94f5c4d504b5e52519b85870386a26102c0516001815580f35b8251855260209485019490920191600101612793565b6102c0516127f69161436e565b6102c05161030f57806126f4565b63b8ff08e960e01b6102c0515260046102c051fd5b3d61014052610140516102c051610280513e61283b610140516102805161436e565b604061028051610140516102805101031261030f576102805151610120526001600160401b03610120511161030f57610140516102805101601f61012051610280510101121561030f57610120516102805101516101005261289f61010051614749565b60e05260405160c0526128b660e05160c05161436e565b60c051506101005160c05152602060c0510160206101005160051b6101205161028051010101610140516102805101811161030f57602061012051610280510101915b8183106129475750505060206102805101516001600160401b03811161030f5761293190610140516102805101906102805101614eea565b60805260c0516101805260805161016052612633565b82518152602092830192016128f9565b909192939495601f198282030186528635605e198436030181121561030f5760206129c26001938683940190813581526129b46129a961299986850185614f7b565b60608886015260608501916146f3565b926040810190614f47565b916040818503910152614fac565b9801960194930191906125ce565b63e734380b60e01b6102c0515260046102c051fd5b3461030f57606036600319011261030f576129fe614176565b6024356001600160401b03811161030f57612a1d903690600401614187565b60443591612a2a84614ea5565b805433600882901c6001600160a01b031603610fcb5760ff90612a4e868686614a16565b16612a588161421e565b8015159081612a91575b506102b757612a7a600582015460068301549061458b565b42101561090757612a8a94614ab9565b6102c05180f35b60039150612a9e8161421e565b141586612a62565b3461030f57602036600319011261030f5760ff19612ac2614176565b166102c0515260016020526020610d3d60406102c05120614959565b34612dae5760c0366003190112612dae57612af7614176565b610520526024356001600160401b038111612dae57612b1a903690600401614187565b6105a052610580526044356001600160401b038111612dae57612b41903690600401614187565b6064356001600160401b038111612dae57612b609036906004016141b4565b61054052610560526084356001600160401b038111612dae57612b879036906004016141b4565b906105005260a4356001600160401b038111612dae57612bab9036906004016141b4565b6104e05290612bb86149f8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612dc157612bfa6105a0516105805161052051614c1f565b9290959391610480526104a0526104c0529293612c2f6104e05183856105005161054051610560516105a0516105805161573d565b958649612db257612c586104e05184866105005161054051610560516105a051610580516157d8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b15612dae57612cf6925f9260405180958194829363be98686160e01b84526105a051610580517f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801614713565b03915afa8015612da357612d8e575b50612d126105805161506a565b91612d1c86614760565b956102c0515b818110612d785750509161026d9591612d56936104e051926105005190610540519061056051906104a0516105205161590a565b6105405191610480516104a0516105a05161058051610520516104c051614dc0565b8060019149612d87828b614803565b5201612d22565b5f612d989161436e565b5f6102c05285612d05565b6040513d5f823e3d90fd5b5f80fd5b63b8ff08e960e01b5f5260045ffd5b63e734380b60e01b5f5260045ffd5b34612dae575f366003190112612dae57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34612dae575f366003190112612dae5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b34612dae575f366003190112612dae576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34612dae57610300366003190112612dae5760056004351015612dae57610100366083190112612dae576001600160401b036101843511612dae5760a06101843536036003190112612dae576101a4356001600160401b038111612dae57612efa903690600401614187565b6040366101e3190112612dae5760e036610223190112612dae57612f1c6149f8565b600354335f81815260026020526040902054602435939266ffffffffffffff90911660589290921b600160581b600160f81b0316600891821c63ffffffff60381b161791909117901b60ff191660ff1981165f9081526001602052604090205490929060081c6001600160a01b03163314613c4c5760a43560ff811614159384612dae5760ff60a43516158015613c3a575b613c2b5760c43560ff811614159485612dae57612dae5760ff60a4351660ff60c435161161382b57610104356101243511613c1c57610144356101643511613c0d5760643515613bfe5761300761010435606435614837565b600561018435600401351015612dae57613027610184356004013561421e565b610184356004013515613bef5761304360846101843501614657565b613be05760246101843501359384613061610184356004013561421e565b6004610184350135600303613b7c575061308060446101843501614664565b803b15613b6d5760405163c1da869160e01b6020820152600481526130af916130aa60248361436e565b61551c565b9015613b6d57935b6130cd6064610184350161018435600401614678565b905015613b5e576130df60043561421e565b600460ff813516118015613b2f575b613b20576131006101c4358285614a16565b60243515613b18575b428410613b0957604435158015613af4575b613ae55760ff1982165f52600160205260405f2060405161313b8161431c565b6101e4358152610204356020820152946004610224351015612dae576102243561389457610244356001600160601b0360a01b8116809103612dae5715801590613888575b801561387c575b8015613870575b8015613864575b8015613858575b613849575b6131aa86614a3f565b1561383a576131bb60043583614598565b6005820155604435600682015560643560078201558054610100600160a81b0319163360081b610100600160a81b031617815584516001820155602085015160028201556040519461320e60e08761436e565b6006865260c03660208801376040519861322960e08b61436e565b60068a5260c03660208c01378460081c6132428b614792565b525f61324d88614792565b5280612dae5760ff60a4351660a43503612dae5760ff60a4351660ff60c435161161382b5765ffffffffffff610104351161381c5765ffffffffffff610124351161380d57677fffffffffffffff61014435116137fe57677fffffffffffffff61016435116137ef5760ff60a4351660a43503612dae57612dae57608435151560843503612dae57608435156137e95760015b60ff60e4351660e43503612dae576020915f9160a43560ff1660c43560081b61ff00161760109190911b62ff0000161760e43560111b6301fe000016176101043560191b176101243560491b176101443560791b176101643560b81b176133468c6147b3565b526002613352896147b3565b528281519101516040519084820192835260408201526040815261337760608261436e565b604051918291518091835e8101838152039060025afa15612da3575f5161339d896147c3565b5260036133a9866147c3565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae2786133d4896147d3565b5260046133e0866147d3565b526133f1610184356004013561421e565b6101843560040135613402896147e3565b52600661340e866147e3565b527f0000000000000000000000000000000000000000000000000000000000000000613439896147f3565b526007613445866147f3565b5284518851036137da576134598551614760565b955f5b865181101561349b578061348a8b610d14836001600160401b036134826001978e614803565b511692614803565b613494828b614803565b520161345c565b50876134a85f8989615c52565b6003840155600e8301608435151560843503612dae57805460ff191660ff608435151516178155805460ff60e4351660e43503612dae5763ff00000060e43560181b169061ff0060a43560081b169063ffffff0019161762ff000060c43560101b161717905561010435600f84015561012435601084015561014435601184015561016435601284015560138301613546610184356004013561421e565b60ff1981541660ff610184356004013516179055601483019081556015830160018060a01b0361357b60446101843501614664565b82546001600160a01b0319169116179055601683016135a4610184356064810190600401614678565b906001600160401b0382116137c65781906135bf84546143ad565b601f8111613796575b505f90601f831160011461372f575f92613724575b50508160011b915f199060031b1c19161790555b61361861360360846101843501614657565b601785019060ff801983541691151516179055565b5543600a820155601a810163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff1982541617905560035463ffffffff811663ffffffff811461371057600163ffffffff9101169063ffffffff191617600355335f52600260205260405f20918254946001600160401b038616936001600160401b038514613710576020966001600160401b036001613700970116906001600160401b0319161790553360ff1986167feefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d15f80a36101c4359285614ab9565b60015f556040519060ff19168152f35b634e487b7160e01b5f52601160045260245ffd5b0135905089806135dd565b909150601f19831691845f5260205f20925f5b81811061377e5750908460019594939210613765575b505050811b0190556135f1565b01355f19600384901b60f8161c19169055898080613758565b91936020600181928787013581550195019201613742565b6137c090855f5260205f20601f850160051c81019160208610611c3257601f0160051c01906146aa565b8a6135c8565b634e487b7160e01b5f52604160045260245ffd5b63d088249360e01b5f5260045ffd5b5f6132e0565b63dd6f54df60e01b5f5260045ffd5b63271fb80560e01b5f5260045ffd5b63871a7fa360e01b5f5260045ffd5b63481eb79f60e01b5f5260045ffd5b632cbdc23160e01b5f5260045ffd5b63208e53e560e11b5f5260045ffd5b63e4291a1960e01b5f5260045ffd5b506102e435151561319c565b506102c4351515613195565b506102a435151561318e565b50610284351515613187565b50610264351515613180565b94506101e43515801590613ad9575b61383a57610224356003036139fb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156139ec5760405163ecd8e20560e01b815260ff19841660048201523360248201529461390e604487016102246145f0565b608086610124815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af18015612da3575f905f905f985f916139b3575b50979091975b604051916139688361431c565b825260208201526018840180546cffffffffffffffffffffffffff19166102243560ff161760989990991c6cffffffffffffffffffffffff00169890981790975560198301556131a1565b925050506139da91965060803d6080116139e5575b6139d2818361436e565b8101906145bd565b97929190978c613955565b503d6139c8565b63e00ffde960e01b5f5260045ffd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615613acb5760405163441d214d60e11b815260ff198416600482015294613a51602487016102246145f0565b608086610104815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af18015612da3575f905f905f985f91613aa2575b509790919761395b565b92505050613ac091965060803d6080116139e5576139d2818361436e565b97929190978c613a98565b6203eb8f60e01b5f5260045ffd5b506102043515156138a3565b637616640160e01b5f5260045ffd5b5042613b026044358661458b565b111561311b565b632ca4094f60e21b5f5260045ffd5b429350613109565b6307a92f1960e51b5f5260045ffd5b50613b3b60043561421e565b600435151580156130ee5750613b5260043561421e565b600360043514156130ee565b630f8b932160e11b5f5260045ffd5b63562e597160e11b5f5260045ffd5b936001600160a01b03613b9461018435604401614664565b16613b6d578415613bc557613baf610184356004013561421e565b6004610184358101351480613bd4575b156130b7575b635e32eadd60e01b5f5260045ffd5b508460a01c1515613bbf565b63f545b7bf60e01b5f5260045ffd5b63f37f7b5d60e01b5f5260045ffd5b632c45be6f60e01b5f5260045ffd5b636d403ec760e11b5f5260045ffd5b63207ea56d60e01b5f5260045ffd5b63ac38930b60e01b5f5260045ffd5b505f9450601060a43560ff1611612fae565b635040c4d360e11b5f5260045ffd5b34612dae576040366003190112612dae57613c74614176565b6024356005811015612dae5760ff198216918215613e4f5763ffffffff808060035460401c16169160401c1603613e4057613cae8161421e565b600460ff821611613b20575f8281526001602052604090208054600881901c6001600160a01b03168015613e31573303613e235760ff16613cef838261486e565b15613b20576005820154926006830192613d0a84548661458b565b93613d148361421e565b60018314958680613e1a575b613def57613d2d8461421e565b6003841480613e10575b613def57613d448461421e565b6002841480613dfe575b613def575f516020615e195f395f51905f5296604096613d6f868b96614598565b613d788661421e565b81613de5575b50613da6575b505050825191613d938161421e565b8252613d9e8161421e565b6020820152a2005b7f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba91613dd4602092426145b0565b8091558651908152a2848680613d84565b9050421089613d7e565b63e843c5eb60e01b5f5260045ffd5b50613e0883614959565b421015613d4e565b5085421015613d37565b50804210613d20565b6282b42960e81b5f5260045ffd5b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b63cbf4a64560e01b5f5260045ffd5b34612dae576020366003190112612dae5760ff19613e7a614176565b165f52600160205260405f20604051613e9281614300565b815460ff8116613ea18161421e565b825260081c6001600160a01b03166020820152613ec06001830161438f565b60408201526003820154606082015260048201604051808260208294549384815201905f5260205f20925f5b81811061404e575050613f019250038261436e565b6080820152600582015460a0820190815260068301549060c08301918252600784015460e084015260088401546101008401526009840154610120840152600a840154610140840152600b840154610160840152613f61600c85016143e5565b610180840152600d8401546101a0840152613f7e600e8501614485565b6101c0840152613f90601385016144e7565b6101e084015260188401549260ff841694600486101561403a576001600160401b03601a6103009260ff610d3d9860209a6102008801526001600160601b0360a01b8160981b1661022088015261ffff8160681c16610240880152818160781c1661026088015261ffff8160801c1661028088015260901c1615156102a086015260198101546102c0860152015463ffffffff81166102e0850152871c169101525190519061458b565b634e487b7160e01b5f52602160045260245ffd5b8454835260019485019486945060209093019201613eec565b34612dae576040366003190112612dae57614080614176565b60243560ff198216918215613e4f5763ffffffff808060035460401c16169160401c1603613e40575f8281526001602052604090208054600881901c6001600160a01b03168015613e31573303613e235760ff166140dd8161421e565b8015159081614161575b50613b20576140ff600582015460068301549061458b565b421015613def5781158015614154575b613bfe57817f36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be92600783614149600f602096015484614837565b0155604051908152a2005b506008810154821061410f565b6003915061416e8161421e565b1415846140e7565b6004359060ff1982168203612dae57565b9181601f84011215612dae578235916001600160401b038311612dae5760208381860195010111612dae57565b9181601f84011215612dae578235916001600160401b038311612dae576020808501948460051b010111612dae57565b34612dae575f366003190112612dae5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b6005111561403a57565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60e0809180511515845260ff602082015116602085015260ff604082015116604085015260ff60608201511660608501526080810151608085015260a081015160a085015260c081015160c08501520151910152565b9081516142ae8161421e565b81526020820151602082015260018060a01b0360408301511660408201526080806142e8606085015160a0606086015260a0850190614228565b930151151591015290565b90600482101561403a5752565b61032081019081106001600160401b038211176137c657604052565b604081019081106001600160401b038211176137c657604052565b61010081019081106001600160401b038211176137c657604052565b60a081019081106001600160401b038211176137c657604052565b90601f801991011681019081106001600160401b038211176137c657604052565b9060405161439c8161431c565b602060018294805484520154910152565b90600182811c921680156143db575b60208310146143c757565b634e487b7160e01b5f52602260045260245ffd5b91607f16916143bc565b9060405191825f8254926143f8846143ad565b8084529360018116908115614463575060011461441f575b5061441d9250038361436e565b565b90505f9291925260205f20905f915b81831061444757505090602061441d928201015f614410565b602091935080600191548385890101520191019091849261442e565b90506020925061441d94915060ff191682840152151560051b8201015f614410565b9060405161449281614337565b60e06004829460ff815481811615158652818160081c166020870152818160101c16604087015260181c16606085015260018101546080850152600281015460a0850152600381015460c08501520154910152565b906040516144f481614353565b608060ff60048395828154166145098161421e565b85526001810154602086015260028101546001600160a01b03166040860152614534600382016143e5565b60608601520154161515910152565b34612dae575f366003190112612dae5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b906001820180921161371057565b9190820180921161371057565b906145a28161421e565b60ff80198354169116179055565b9190820391821161371057565b9190826080910312612dae5781516001600160a01b031981168103612dae57916020810151916060604083015192015190565b80356004811015612dae5782614605916142f3565b60208101356001600160601b0360a01b8116809103612dae5760c0918291602085015260408101356040850152606081013560608501526080810135608085015260a081013560a08501520135910152565b358015158103612dae5790565b356001600160a01b0381168103612dae5790565b903590601e1981360301821215612dae57018035906001600160401b038211612dae57602001918136038313612dae57565b8181106146b5575050565b5f81556001016146aa565b90602080835192838152019201905f5b8181106146dd5750505090565b82518452602093840193909201916001016146d0565b908060209392818452848401375f828201840152601f01601f1916010190565b94929093614738926147469795875260208701526080604087015260808601916146f3565b9260608185039101526146f3565b90565b6001600160401b0381116137c65760051b60200190565b9061476a82614749565b614777604051918261436e565b8281528092614788601f1991614749565b0190602036910137565b80511561479f5760200190565b634e487b7160e01b5f52603260045260245ffd5b80516001101561479f5760400190565b80516002101561479f5760600190565b80516003101561479f5760800190565b80516004101561479f5760a00190565b80516005101561479f5760c00190565b805182101561479f5760209160051b010190565b604081101561479f5760051b60240190565b5f1981146137105760010190565b801561485a5764e8d4a51000041061484b57565b63eba5c29b60e01b5f5260045ffd5b634e487b7160e01b5f52601260045260245ffd5b6148778161421e565b6148808261421e565b80821461492d576148908161421e565b600281148015614946575b8015614933575b61492d576148af8161421e565b801561490e576003906148c18161421e565b146148cb57505f90565b6148d48161421e565b80159081156148f9575b81156148e8575090565b600191506148f58161421e565b1490565b90506149048161421e565b60028114906148de565b506149188161421e565b600381149081156148f95781156148e8575090565b50505f90565b5061493d8161421e565b600181146148a2565b506149508161421e565b6004811461489b565b61496c600582015460068301549061458b565b9063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016801983116149f0576149cc601a6149d29301546001600160401b038160201c168581115f146149e35763ffffffff90915b169061458b565b9261458b565b808210156149de575090565b905090565b5063ffffffff85916149c5565b5050505f1990565b60025f5414614a075760025f55565b633ee5aeb560e01b5f5260045ffd5b5015908115614a36575b50614a2757565b635e765b2560e11b5f5260045ffd5b9050155f614a20565b602081519101519080158015614aa2575b8015614a8b575b61492d575f516020615e395f395f51905f52808281930992800981808080848709620292f80960010893620292fc09081490565b505f516020615e395f395f51905f52821015614a57565b505f516020615e395f395f51905f52811015614a50565b600c820193916001600160401b0383116137c6578590614ad986546143ad565b601f8111614bef575b505f95601f8511600114614b655790600d91857f77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f0985f91614b5a575b508660011b905f198860031b1c19161790555b0155614b4a6040519384936040855260408501916146f3565b94602083015260ff1916930390a2565b90508701355f614b1e565b601f19851696815f5260205f20975f5b818110614bd4575097600d939291877f77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f09a10614bbb575b5050600186811b019055614b31565b8801355f19600389901b60f8161c191690555f80614bac565b888301358a556001909901988a955060209283019201614b75565b614c1990875f5260205f20601f870160051c81019160208810611c3257601f0160051c01906146aa565b5f614ae2565b614c2b90939293614ea5565b9260ff845416614c3a8161421e565b8015159081614daa575b81614d73575b50613b205760058401544210613def57614c6384614959565b421015613def57614c74908261500b565b614c7d8161509f565b9160038401548303614d64575f5f5b60088110614d435750614ca890614ca290615537565b8561569d565b60088401549182614cc061015083013560c01c615b72565b03614d3457614cf081614ceb614cde6098609095013560c01c615b72565b938492013560c01c615b72565b6145b0565b91614cfb828461458b565b15614d2557614d0a838561458b565b600787015410614d1657565b6357d18d5360e11b5f5260045ffd5b633f8cdc5560e11b5f5260045ffd5b63341203dd60e21b5f5260045ffd5b906001908260140160031b85013560e01c8360051b60e0031b179101614c8c565b630b6fac0360e41b5f5260045ffd5b60039150614d808161421e565b1480614d8e575b155f614c4a565b50614da2600585015460068601549061458b565b421015614d87565b9050614db58161421e565b600181141590614c44565b9396949097959192505f925f905b60088210614e82575050601a91614dea9184600387015561458b565b92836008820155614e006009820195865461458b565b809555600b8101614e118154614829565b90550180546bffffffffffffffff000000004260201b16906bffffffffffffffff00000000191617905560405194855260208501526040840152606083015260808201527f36c6781d994e030a156d2f6fa11abcc1cd6814482a5d61f90da3f336318b1d2360a0339360ff191692a3565b909360019085600a0160031b83013560e01c8660051b60e0031b17940190614dce565b60ff198116908115613e4f5763ffffffff808060035460401c16169160401c1603613e40575f52600160205260405f209060018060a01b03825460081c1615613e3157565b9080601f83011215612dae578151614f0181614749565b92614f0f604051948561436e565b81845260208085019260051b820101928311612dae57602001905b828210614f375750505090565b8151815260209182019101614f2a565b9035601e1982360301811215612dae5701602081359101916001600160401b038211612dae578160051b36038313612dae57565b9035601e1982360301811215612dae5701602081359101916001600160401b038211612dae578136038313612dae57565b90602083828152019260208260051b82010193835f925b848410614fd35750505050505090565b909192939495602080614ffb600193601f19868203018852614ff58b88614f7b565b906146f3565b9801940194019294939190614fc3565b906102000361505b576001615023823560c01c615b72565b1490811591615042575b5061503457565b6254aabb60e81b5f5260045ffd5b61505391506008013560c01c615b72565b15155f61502d565b630f61e7fd60e21b5f5260045ffd5b5f905f905b6008821061507c57505090565b909160019083602e0160031b83013560e01c8460051b60e0031b1792019061506f565b5f905f905b600882106150b157505090565b90916001908360020160031b83013560e01c8460051b60e0031b179201906150a4565b919081101561479f5760051b0190565b92919091801580156151a0575b615198575f925f5b82811061516c575081841461516357615113906004615bbf565b92805b615121575050501490565b5f19019283906004821c60011661514e57615148906151418385876150d4565b3590615c2f565b93615116565b6151489061515d8385876150d4565b35615c2f565b50505050505f90565b6151778184866150d4565b35615185575b6001016150f9565b935060018401808511613710579361517d565b505050505f90565b50604081116150f1565b600481101561403a576003036151e8577f00000000000000000000000000000000000000000000000000000000000000005b6001600160a01b031690565b7f00000000000000000000000000000000000000000000000000000000000000006151dc565b51908115158203612dae57565b6018015460ff8116600481101561403a576003149081159161523b575090565b604051635f0ddacd60e01b815260989190911b6001600160a01b031916600482015290506020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115612da3575f916152a2575090565b90506020813d6020116152ce575b816152bd6020938361436e565b81010312612dae576147469061520e565b3d91506152b0565b9060ff600e82015460081c16916152ec83614760565b92601883015460ff8160781c1690816153f4575b50505060048254928160ff85169460ff1916178155018351906001600160401b0382116137c657600160401b82116137c65780548282558083106153d8575b5060208501905f5260205f205f5b8381106153c45750505050905f516020615e195f395f51905f5260409260ff191692839281519061537d8161421e565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e60405160208152806153bf339560208301906146c0565b0390a3565b60019060208451940193818401550161534d565b6153ee90825f528360205f2091820191016146aa565b5f61533f565b5f6001600160a01b0361540960ff84166151aa565b1692608460198801546040519586938492630222162f60e21b84526001600160601b0360a01b8860981b166004850152602484015261ffff8760681c16604484015260648301525afa918215612da3575f905f936154cf575b50156154c0575f919060801c61ffff16825b848110615482575050615300565b600182821c1615615496575b600101615474565b926154b8816154a760019386614803565b516154b2878c614803565b52614829565b93905061548e565b630e0d4dc560e41b5f5260045ffd5b9250503d805f843e6154e1818461436e565b820191604081840312612dae576154f78161520e565b9260208201516001600160401b038111612dae576155159201614eea565b915f615462565b6020915f91838251920190620186a0fa601f3d1116905f5190565b77ffffffffffffffff0000000000000000ffffffffffffffff8160081c9160081b917cff000000ff000000ff000000ff000000ff000000ff000000ff000000ff7dff000000ff000000ff000000ff000000ff000000ff000000ff000000ff007fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff0085167eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff84161760101c941691161760101b9179ffff000000000000ffff000000000000ffff000000000000ffff7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff847dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff84161760201c941691161760201b7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff82821673ffffffff000000000000000000000000ffffffff85161760401b93161760401c16178060801b9060801c1790565b90600360ff6013840154166156b18161421e565b0361573157601582015460405163650e5fcf60e01b60208201526024808201939093529182526156f091906001600160a01b03166130aa60448361436e565b9015918215615728575b821561571e575b8215615710575b5050613bc557565b600a01541190505f80615708565b4382119250615701565b811592506156fa565b906014015403613bc557565b915050610120919394506157599250959495013560c01c615b72565b938415615793578480921492831593615788575b50821561577d575b5050612db257565b14159050825f615775565b82141592505f61576d565b63fdac229f60e01b5f5260045ffd5b6001600160401b0381116137c657601f01601f191660200190565b9082101561479f576157d49160051b810190614678565b9091565b96949192959050605086029486860460501487151715613710576158146157fe876157a2565b9661580c604051988961436e565b8088526157a2565b602087019790601f19013689375f5b81811061589b57505050505050506020915f91604051918291518091835e8101838152039060025afa15612da3575f515f915f905b600882106158785750500361586957565b630f2e4cb960e31b5f5260045ffd5b909260019084601c0160031b83013560e01c8560051b60e0031b17930190615858565b806158a960309284896157bd565b929092036158fb5760306158be8286886157bd565b9050036158ec5760019160506158d583898c6150d4565b359160308285028d01916020830137015201615823565b6350320ab160e01b5f5260045ffd5b6338b2168b60e21b5f5260045ffd5b9095919492945f5b858110615926575050505050505050505050565b615930818c614803565b5115615b60578660205f85898f60a08e61596561595c8a86615955826159dd98614803565b519b6157bd565b91829193615537565b926040519687948b8601996f646176696e63692f706f696e742f763160801b8b5260081c603087015260508601528d60708601526001600160401b0360c01b9060c01b1660908501526001600160401b0360c01b8c60c01b16609885015284840137810186838201520301601f19810183528261436e565b604051918291518091835e8101838152039060025afa15612da3575f809189615a8f898f8f90608091898f615a488f615a5292615a418580938f7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff000000019051069b6150d4565b359a6157bd565b94859391946157bd565b9384929193604051998a9860208a019c8d5260408a01526060890152868801378501918483018a815237010185815203601f19810183528261436e565b5190600a5afa3d15615b58573d90615aa6826157a2565b91615ab4604051938461436e565b82523d5f602084013e5b158015615b4c575b615b3957604081805181010312612dae5761100060406020830151920151911490811591615b0e575b50615afc57600101615912565b638308e1e960e01b5f5260045260245ffd5b7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001915014155f615aef565b50638308e1e960e01b5f5260045260245ffd5b50604081511415615ac6565b606090615abe565b634af69d9960e11b5f5260045260245ffd5b66ff00ff00ff00ff67ff00ff00ff00ff008260081b169160081c161765ffff0000ffff67ffff0000ffff00008260101b169160101c161767ffffffff000000008160201b169060201c1790565b602091615be75f92615bdd6001600160401b038060c01b9216615b72565b60c01b1691615537565b604051908482019283526028820152600160f81b604882015260298152615c0f60498261436e565b604051918291518091835e8101838152039060025afa15612da3575f5190565b5f9060209260405190848201928352604082015260408152615c0f60608261436e565b90815180156151985760018114615e075760408414615df8575f915f5b828110615dcf5750615c89615c8484846145b0565b614760565b91615c97615c8485836145b0565b91615caa615ca486614760565b95614760565b955f9081805b888a8c888710615cef5795505050505050615cd894939250615cd2915061457d565b91615c52565b91615cd2615ce6939461457d565b61474691615c2f565b90615d346001615d278a9594615d15615d088c8c614803565b516001600160401b031690565b906001600160401b03809216901c1690565b166001600160401b031690565b615d8d57505050615d84600191615d7f615d51615d088888614803565b615d5b888a614803565b51615d66848d614803565b52615d71838d614803565b906001600160401b03169052565b614829565b935b0192615cb0565b615dc992615d71868094615d7f94615dc38c9a615dbc8c60019c9f8f615db691615d0891614803565b98614803565b5192614803565b52614803565b91615d86565b92615df1600191615deb83615d278a615d15615d088b8d614803565b9061458b565b9301615c6f565b63394fd24160e21b5f5260045ffd5b509050615e149150614792565b519056fe56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220ebc51d37ac156c8fb3b3b5ce74f2e05048fb4318f6ef748e43d2469b7f2d180564736f6c634300081c003360e0806040523461011157602061002e60049261106480380380916100248285610115565b833981019061014c565b336080526001600160a01b031660a081905260405163ebe86c1360e01b815292839182905afa908115610106575f916100d7575b506001600160a01b031660c052604051610ef8908161016c82396080518181816101530152818161037701528181610756015281816107cd0152610bc0015260a051818181610186015281816108b701528181610a2f0152610c3f015260c05181818160c30152818161047001526107fd0152f35b6100f9915060203d6020116100ff575b6100f18183610115565b81019061014c565b5f610062565b503d6100e7565b6040513d5f823e3d90fd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761013857604052565b634e487b7160e01b5f52604160045260245ffd5b9081602091031261011157516001600160a01b0381168103610111579056fe60806040526004361015610011575f80fd5b5f5f3560e01c8063088858bc146108e6578063481c6a75146108a257806364fe50ac146107b1578063702574b3146107855780637b10399914610740578063883a429a1461034e578063ae977ede146100f2578063ebe86c13146100ad5763f08c9b3c1461007d575f80fd5b346100aa57806003193601126100aa576020610097610c30565b6040516001600160a01b03199091168152f35b80fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760603660031901126100aa5761010c610977565b9060243591604435916001600160401b0383116100aa57366023840112156100aa578260040135936001600160401b03851161034a573660248660071b8601011161034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b579093927f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316928592916001600160a01b031990911690835b8688101561032c578760071b82019060845f516020610ea35f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d2899602485013509925f516020610ea35f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d289960648301350960405194633d98dab360e11b865287600487015288602487015260448601526044820135606486015282850152013560a483015260208260c481898b5af19182156103215786926102e2575b508861029c5750600190975b01966101c1565b979061ffff8916908282018092116102ce5761ffff16036102bf57600190610295565b6357149e2560e01b8552600485fd5b634e487b7160e01b87526011600452602487fd5b9091506020813d8211610319575b816102fd6020938361099f565b810103126103155761030e906109eb565b905f610289565b8580fd5b3d91506102f0565b6040513d88823e3d90fd5b60209061ffff60405191168152f35b633217675b60e21b8252600482fd5b5080fd5b50346100aa576101003660031901126100aa5761036961098e565b60e036602319011261034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b576103ae90610bab565b602435600481101561073c5760021490811561072e576044356001600160a01b03198116810361072a57915b6040908151906103ea838361099f565b6001825260208201601f198401368237825115610716573090521561070e5784935b82519461010086018681106001600160401b038211176106fa57845260028110156106e6578552602085019486865283810192835260608101601081526080820188815260a0830189815260c08401918a835260e08501938b855260018060a01b037f00000000000000000000000000000000000000000000000000000000000000001697883b156106e257895197631bb64f5960e21b89526001600160601b0360a01b169b8c60048a01528b60248a015261010060448a01526102048901975160028110156106ce5791899693918f989593610104899b989b01525115156101248801525194610100610144880152855180915260206102248801960190885b81811061069d57505050916001600160401b03869798818096959461ffff829651166101648b0152511661018489015251166101a487015251166101c485015251166101e48301526064356064830152608435608483015260a43560a483015260c43560c483015260e43560e4830152038183865af180156106935761067a575b5093816044958151968780926303e95d1360e21b82528860048301528760248301525afa92831561066e57608095829461060b575b505f516020610ea35f395f51905f52917f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e0097542918451968752602087015209908301526060820152f35b7f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e00975429194505f516020610ea35f395f51905f52925061065e90843d8611610667575b610656818361099f565b810190610c1a565b949150916105c3565b503d61064c565b509051903d90823e3d90fd5b61068586809261099f565b61068f575f61058e565b8480fd5b83513d88823e3d90fd5b92949750929497509794602080600192838060a01b038c511681520199019101908e9794928a97949299969961050d565b634e487b7160e01b8f52602160045260248ffd5b8c80fd5b634e487b7160e01b87526021600452602487fd5b634e487b7160e01b88526041600452602488fd5b60019361040c565b634e487b7160e01b87526032600452602487fd5b8380fd5b610736610c30565b916103da565b8280fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760203660031901126100aa5760206107a96107a461098e565b610bab565b604051908152f35b503461088f57606036600319011261088f576107cb610977565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610893577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561088f575f9160648392604051948593849263a59b7a4d60e01b84526001600160601b0360a01b166004840152602435602484015260443560448401525af1801561088457610876575080f35b61088291505f9061099f565b005b6040513d5f823e3d90fd5b5f80fd5b633217675b60e21b5f5260045ffd5b3461088f575f36600319011261088f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461088f57608036600319011261088f576108ff610977565b60443561ffff8116810361088f576064359061ffff8216820361088f5761092992602435906109fa565b906040519182916040830190151583526040602084015281518091526020606084019201905f5b81811061095e575050500390f35b8251845285945060209384019390920191600101610950565b600435906001600160a01b03198216820361088f57565b6004359060ff198216820361088f57565b90601f801991011681019081106001600160401b038211176109c057604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b0381116109c05760051b60200190565b519061ffff8216820361088f57565b9261ffff1693610a09856109d4565b610a16604051918261099f565b858152601f19610a25876109d4565b01366020830137937f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925f5b61ffff811688811015610b9e578061ffff88160161ffff8111610b8a57604051639bbada6760e01b81526001600160a01b0319861660048201526024810185905261ffff9190911660448201526060816064818a5afa908115610884575f91610b14575b50602081015115610b0657604001519088511115610af257600582901b621fffe01688016020015260010161ffff16610a5a565b634e487b7160e01b5f52603260045260245ffd5b505f98509395505050505050565b90506060813d8211610b82575b81610b2e6060938361099f565b8101031261088f5760405190606082018281106001600160401b038211176109c057604052610b5c816109eb565b8252602081015190811515820361088f576040916020840152015160408201525f610abe565b3d9150610b21565b634e487b7160e01b5f52601160045260245ffd5b5050505093505050600191565b60405190602082019046825260018060a01b037f000000000000000000000000000000000000000000000000000000000000000016604084015260ff1916606083015260608252610bfd60808361099f565b6b0fffffffffffffffffffffff60a01b9151902060041c16301790565b919082604091031261088f576020825192015190565b60405163a4adcd7f60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316602082600481845afa918215610884575f92610e5e575b506040516323488be560e01b8152602081600481855afa8015610884575f90610e18575b63ffffffff60401b915060401b165f5b6001600160401b0381166008811080610e06575b15610df7576001600160401b038516036001600160401b038111610b8a576001600160401b036001600160601b0360a01b9116831760a01b1660405163d397925360e01b8152816004820152602081602481885afa8015610884575f90610dbb575b60ff9150166010811015610dab57604051906356cbb5f360e01b82528260048301526024820152604081604481885afa9081610d8e575b50610d8657506001600160401b03905b166001600160401b038114610b8a57600101610cb2565b935050505090565b610da59060403d811161066757610656818361099f565b50610d5f565b50506001600160401b0390610d6f565b506020813d8211610def575b81610dd46020938361099f565b8101031261088f575160ff8116810361088f5760ff90610d28565b3d9150610dc7565b63081ea97160e31b5f5260045ffd5b50806001600160401b03861611610cc6565b506020813d602011610e56575b81610e326020938361099f565b8101031261088f575163ffffffff8116810361088f5763ffffffff60401b90610ca2565b3d9150610e25565b9091506020813d602011610e9a575b81610e7a6020938361099f565b8101031261088f57516001600160401b038116810361088f57905f610c7e565b3d9150610e6d56fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220e80e7f0d450628d1e82cbd4c507827a00c4b11ab550fc488d6c26c4ede03779a64736f6c634300081c003360c0346100a257601f610a4438819003918201601f19168301916001600160401b038311848410176100a6578084926020946040528339810103126100a257516001600160a01b038116908190036100a2573360805260a05260405161098990816100bb823960805181818160b90152818161034101526104ea015260a051818181610168015281816104050152818161058101528181610614015261075a0152f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c908163088858bc14610698575080632680c7d814610643578063481c6a75146105ff5780635f0ddacd1461054257806364fe50ac146105195780637b103999146104d5578063ae977ede146102e25763ecd8e20514610074575f80fd5b34610266576101203660031901126102665760043560ff198116809103610266576024356001600160a01b03811691908290036102665760e0366043190112610266577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036102d357604435600481101561026657600219016102c457606435906001600160601b0360a01b821680920361026657811580156102b9575b80156102ae575b80156102a3575b8015610298575b801561028c575b61027d576040516358d6692360e01b8152600481018390526024810182905260448101939093526060836064815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610272575f5f915f93610228575b6080955060016040516101b581610908565b8681525f602080830182815260408085019687528784529183905291209151825491516cffffffffffffffffffffffffff1990921660a09190911c1760609190911b60ff60601b1617815501905160081c60ff60f81b825416179055604051938452602084015260408301526060820152f35b925050506060833d60601161026a575b8161024560609383610924565b810103126102665782608093519060406020820151910151919091926101a3565b5f80fd5b3d9150610238565b6040513d5f823e3d90fd5b63e4291a1960e01b5f5260045ffd5b50610104351515610138565b5060e4351515610131565b5060c435151561012a565b5060a4351515610123565b50608435151561011c565b6365b75c3960e01b5f5260045ffd5b633217675b60e21b5f5260045ffd5b34610266576060366003190112610266576102fb6108f1565b6024359060443567ffffffffffffffff8111610266573660238201121561026657806004013567ffffffffffffffff8111610266573660248260071b84010111610266577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036102d357835f525f60205260405f2060ff19600182015460081b1693841580156104b9575b6104aa57815460ff60601b1916606084901b60ff60601b16179091559291908060405194859463548efb7160e01b865260648601916001600160601b0360a01b166004870152602486015260606044860152526024608484019201905f5b81811061048e57506020939283900391508290505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610272575f9161045c575b500361044d5760206040515f8152f35b63ba343de560e01b5f5260045ffd5b90506020813d602011610486575b8161047760209383610924565b8101031261026657518261043d565b3d915061046a565b91935091608080828187600195370194019101918493926103ee565b636d08029760e01b5f5260045ffd5b50815460a01b6001600160a01b03199081169082161415610390565b34610266575f366003190112610266576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610266576060366003190112610266576105326108f1565b50633254484760e01b5f5260045ffd5b346102665760203660031901126102665761055b6108f1565b604051635f0ddacd60e01b81526001600160a01b031990911660048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610272575f906105c5575b6020906040519015158152f35b506020813d6020116105f7575b816105df60209383610924565b81010312610266576105f2602091610946565b6105b8565b3d91506105d2565b34610266575f366003190112610266576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610266576020366003190112610266576004355f525f602052606060405f206001815491015460081b60ff604051926001600160601b0360a01b8160a01b168452841c16602083015260ff19166040820152f35b34610266576080366003190112610266576106b16108f1565b6024359160443561ffff8116809103610266576064359261ffff841680940361026657845f525f60205260405f20926106e981610908565b83549060ff6001600160601b0360a01b8360a01b169283835260601c169060406020820196838852600160ff1991015460081b16910152159182156108dc575b50506104aa5715908115916108cd575b506108be57604051635584afe560e11b815260048101929092525f826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610272575f925f926107fc575b508290816107f0575b5061044d57906040519182916040830190151583526040602084015281518091526020606084019201905f5b8181106107d7575050500390f35b82518452859450602093840193909201916001016107c9565b9050815114158361079d565b925090503d805f843e61080f8184610924565b8201916040818403126102665761082581610946565b9060208101519067ffffffffffffffff821161026657019280601f850112156102665783519367ffffffffffffffff85116108aa578460051b90604051956108706020840188610924565b865260208087019282010192831161026657602001905b82821061089a5750505080929190610794565b8151815260209182019101610887565b634e487b7160e01b5f52604160045260245ffd5b6347269e1960e11b5f5260045ffd5b60ff9150511681141583610739565b6001600160a01b031916141590508580610729565b600435906001600160a01b03198216820361026657565b6060810190811067ffffffffffffffff8211176108aa57604052565b90601f8019910116810190811067ffffffffffffffff8211176108aa57604052565b519081151582036102665756fea2646970667358221220b0218f600d6bdc2356edf65eaad7dea8f130c0a94e413d8193a83ebcf79b54d164736f6c634300081c0033",
}

// ProcessRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProcessRegistryMetaData.ABI instead.
var ProcessRegistryABI = ProcessRegistryMetaData.ABI

// ProcessRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProcessRegistryMetaData.Bin instead.
var ProcessRegistryBin = ProcessRegistryMetaData.Bin

// DeployProcessRegistry deploys a new Ethereum contract, binding an instance of ProcessRegistry to it.
func DeployProcessRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID uint32, _ziskVerifier common.Address, _batchProgramVK [32]byte, _resultsProgramVK [32]byte, _rootCVadcopFinal [32]byte, _ballotVKHash [32]byte, _dkgManager common.Address, _councilManager common.Address, _defaultGrace uint32, _graceFloor uint32, _graceCeil uint32, _graceMaxTotal uint32, _noticeMin uint32, _l1BlobVerifier common.Address) (common.Address, *types.Transaction, *ProcessRegistry, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProcessRegistryBin), backend, _chainID, _ziskVerifier, _batchProgramVK, _resultsProgramVK, _rootCVadcopFinal, _ballotVKHash, _dkgManager, _councilManager, _defaultGrace, _graceFloor, _graceCeil, _graceMaxTotal, _noticeMin, _l1BlobVerifier)
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

// AidFor is a free data retrieval call binding the contract method 0x702574b3.
//
// Solidity: function aidFor(bytes31 processId) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) AidFor(opts *bind.CallOpts, processId [31]byte) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "aidFor", processId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AidFor is a free data retrieval call binding the contract method 0x702574b3.
//
// Solidity: function aidFor(bytes31 processId) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) AidFor(processId [31]byte) ([32]byte, error) {
	return _ProcessRegistry.Contract.AidFor(&_ProcessRegistry.CallOpts, processId)
}

// AidFor is a free data retrieval call binding the contract method 0x702574b3.
//
// Solidity: function aidFor(bytes31 processId) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) AidFor(processId [31]byte) ([32]byte, error) {
	return _ProcessRegistry.Contract.AidFor(&_ProcessRegistry.CallOpts, processId)
}

// BallotVKHash is a free data retrieval call binding the contract method 0x0e2ebcf7.
//
// Solidity: function ballotVKHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) BallotVKHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "ballotVKHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BallotVKHash is a free data retrieval call binding the contract method 0x0e2ebcf7.
//
// Solidity: function ballotVKHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) BallotVKHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.BallotVKHash(&_ProcessRegistry.CallOpts)
}

// BallotVKHash is a free data retrieval call binding the contract method 0x0e2ebcf7.
//
// Solidity: function ballotVKHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) BallotVKHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.BallotVKHash(&_ProcessRegistry.CallOpts)
}

// BatchProgramVK is a free data retrieval call binding the contract method 0x946544bf.
//
// Solidity: function batchProgramVK() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) BatchProgramVK(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "batchProgramVK")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchProgramVK is a free data retrieval call binding the contract method 0x946544bf.
//
// Solidity: function batchProgramVK() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) BatchProgramVK() ([32]byte, error) {
	return _ProcessRegistry.Contract.BatchProgramVK(&_ProcessRegistry.CallOpts)
}

// BatchProgramVK is a free data retrieval call binding the contract method 0x946544bf.
//
// Solidity: function batchProgramVK() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) BatchProgramVK() ([32]byte, error) {
	return _ProcessRegistry.Contract.BatchProgramVK(&_ProcessRegistry.CallOpts)
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

// CouncilAdapter is a free data retrieval call binding the contract method 0xabaab13a.
//
// Solidity: function councilAdapter() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) CouncilAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "councilAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CouncilAdapter is a free data retrieval call binding the contract method 0xabaab13a.
//
// Solidity: function councilAdapter() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) CouncilAdapter() (common.Address, error) {
	return _ProcessRegistry.Contract.CouncilAdapter(&_ProcessRegistry.CallOpts)
}

// CouncilAdapter is a free data retrieval call binding the contract method 0xabaab13a.
//
// Solidity: function councilAdapter() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) CouncilAdapter() (common.Address, error) {
	return _ProcessRegistry.Contract.CouncilAdapter(&_ProcessRegistry.CallOpts)
}

// DefaultGrace is a free data retrieval call binding the contract method 0xc8f0582f.
//
// Solidity: function defaultGrace() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) DefaultGrace(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "defaultGrace")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DefaultGrace is a free data retrieval call binding the contract method 0xc8f0582f.
//
// Solidity: function defaultGrace() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) DefaultGrace() (uint32, error) {
	return _ProcessRegistry.Contract.DefaultGrace(&_ProcessRegistry.CallOpts)
}

// DefaultGrace is a free data retrieval call binding the contract method 0xc8f0582f.
//
// Solidity: function defaultGrace() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) DefaultGrace() (uint32, error) {
	return _ProcessRegistry.Contract.DefaultGrace(&_ProcessRegistry.CallOpts)
}

// DkgAdapter is a free data retrieval call binding the contract method 0xe16d5b7c.
//
// Solidity: function dkgAdapter() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) DkgAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "dkgAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DkgAdapter is a free data retrieval call binding the contract method 0xe16d5b7c.
//
// Solidity: function dkgAdapter() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) DkgAdapter() (common.Address, error) {
	return _ProcessRegistry.Contract.DkgAdapter(&_ProcessRegistry.CallOpts)
}

// DkgAdapter is a free data retrieval call binding the contract method 0xe16d5b7c.
//
// Solidity: function dkgAdapter() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) DkgAdapter() (common.Address, error) {
	return _ProcessRegistry.Contract.DkgAdapter(&_ProcessRegistry.CallOpts)
}

// GenesisRoot is a free data retrieval call binding the contract method 0xbf74291e.
//
// Solidity: function genesisRoot(bytes31 processId, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint256,uint256) encryptionKey, uint8 censusOrigin) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) GenesisRoot(opts *bind.CallOpts, processId [31]byte, ballotMode DAVINCITypesBallotMode, encryptionKey DAVINCITypesEncryptionKey, censusOrigin uint8) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "genesisRoot", processId, ballotMode, encryptionKey, censusOrigin)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GenesisRoot is a free data retrieval call binding the contract method 0xbf74291e.
//
// Solidity: function genesisRoot(bytes31 processId, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint256,uint256) encryptionKey, uint8 censusOrigin) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) GenesisRoot(processId [31]byte, ballotMode DAVINCITypesBallotMode, encryptionKey DAVINCITypesEncryptionKey, censusOrigin uint8) ([32]byte, error) {
	return _ProcessRegistry.Contract.GenesisRoot(&_ProcessRegistry.CallOpts, processId, ballotMode, encryptionKey, censusOrigin)
}

// GenesisRoot is a free data retrieval call binding the contract method 0xbf74291e.
//
// Solidity: function genesisRoot(bytes31 processId, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint256,uint256) encryptionKey, uint8 censusOrigin) view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GenesisRoot(processId [31]byte, ballotMode DAVINCITypesBallotMode, encryptionKey DAVINCITypesEncryptionKey, censusOrigin uint8) ([32]byte, error) {
	return _ProcessRegistry.Contract.GenesisRoot(&_ProcessRegistry.CallOpts, processId, ballotMode, encryptionKey, censusOrigin)
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),bytes32,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,bytes32,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),uint8,bytes12,uint16,uint8,uint16,bool,bytes32,uint32,uint64))
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),bytes32,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,bytes32,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),uint8,bytes12,uint16,uint8,uint16,bool,bytes32,uint32,uint64))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [31]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x59d821c4.
//
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),bytes32,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,bytes32,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),uint8,bytes12,uint16,uint8,uint16,bool,bytes32,uint32,uint64))
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

// GetProcessGraceEnd is a free data retrieval call binding the contract method 0x3ea4ee41.
//
// Solidity: function getProcessGraceEnd(bytes31 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistryCaller) GetProcessGraceEnd(opts *bind.CallOpts, processId [31]byte) (*big.Int, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcessGraceEnd", processId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProcessGraceEnd is a free data retrieval call binding the contract method 0x3ea4ee41.
//
// Solidity: function getProcessGraceEnd(bytes31 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistrySession) GetProcessGraceEnd(processId [31]byte) (*big.Int, error) {
	return _ProcessRegistry.Contract.GetProcessGraceEnd(&_ProcessRegistry.CallOpts, processId)
}

// GetProcessGraceEnd is a free data retrieval call binding the contract method 0x3ea4ee41.
//
// Solidity: function getProcessGraceEnd(bytes31 processId) view returns(uint256)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcessGraceEnd(processId [31]byte) (*big.Int, error) {
	return _ProcessRegistry.Contract.GetProcessGraceEnd(&_ProcessRegistry.CallOpts, processId)
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

// GraceCeil is a free data retrieval call binding the contract method 0x1542bbe2.
//
// Solidity: function graceCeil() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) GraceCeil(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "graceCeil")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GraceCeil is a free data retrieval call binding the contract method 0x1542bbe2.
//
// Solidity: function graceCeil() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) GraceCeil() (uint32, error) {
	return _ProcessRegistry.Contract.GraceCeil(&_ProcessRegistry.CallOpts)
}

// GraceCeil is a free data retrieval call binding the contract method 0x1542bbe2.
//
// Solidity: function graceCeil() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GraceCeil() (uint32, error) {
	return _ProcessRegistry.Contract.GraceCeil(&_ProcessRegistry.CallOpts)
}

// GraceFloor is a free data retrieval call binding the contract method 0x5ff5f981.
//
// Solidity: function graceFloor() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) GraceFloor(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "graceFloor")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GraceFloor is a free data retrieval call binding the contract method 0x5ff5f981.
//
// Solidity: function graceFloor() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) GraceFloor() (uint32, error) {
	return _ProcessRegistry.Contract.GraceFloor(&_ProcessRegistry.CallOpts)
}

// GraceFloor is a free data retrieval call binding the contract method 0x5ff5f981.
//
// Solidity: function graceFloor() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GraceFloor() (uint32, error) {
	return _ProcessRegistry.Contract.GraceFloor(&_ProcessRegistry.CallOpts)
}

// GraceMaxTotal is a free data retrieval call binding the contract method 0x549d5995.
//
// Solidity: function graceMaxTotal() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) GraceMaxTotal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "graceMaxTotal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GraceMaxTotal is a free data retrieval call binding the contract method 0x549d5995.
//
// Solidity: function graceMaxTotal() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) GraceMaxTotal() (uint32, error) {
	return _ProcessRegistry.Contract.GraceMaxTotal(&_ProcessRegistry.CallOpts)
}

// GraceMaxTotal is a free data retrieval call binding the contract method 0x549d5995.
//
// Solidity: function graceMaxTotal() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GraceMaxTotal() (uint32, error) {
	return _ProcessRegistry.Contract.GraceMaxTotal(&_ProcessRegistry.CallOpts)
}

// L1BlobVerifier is a free data retrieval call binding the contract method 0x0c39fa36.
//
// Solidity: function l1BlobVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) L1BlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "l1BlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1BlobVerifier is a free data retrieval call binding the contract method 0x0c39fa36.
//
// Solidity: function l1BlobVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) L1BlobVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.L1BlobVerifier(&_ProcessRegistry.CallOpts)
}

// L1BlobVerifier is a free data retrieval call binding the contract method 0x0c39fa36.
//
// Solidity: function l1BlobVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) L1BlobVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.L1BlobVerifier(&_ProcessRegistry.CallOpts)
}

// NoticeMin is a free data retrieval call binding the contract method 0xd4138a20.
//
// Solidity: function noticeMin() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) NoticeMin(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "noticeMin")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NoticeMin is a free data retrieval call binding the contract method 0xd4138a20.
//
// Solidity: function noticeMin() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) NoticeMin() (uint32, error) {
	return _ProcessRegistry.Contract.NoticeMin(&_ProcessRegistry.CallOpts)
}

// NoticeMin is a free data retrieval call binding the contract method 0xd4138a20.
//
// Solidity: function noticeMin() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) NoticeMin() (uint32, error) {
	return _ProcessRegistry.Contract.NoticeMin(&_ProcessRegistry.CallOpts)
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
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, bytes32 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, bytes32 metadataHash, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, uint8 keyMode, bytes12 dkgEpochId, uint16 dkgFirstIndex, uint8 dkgCount, uint16 dkgZeroSkipped, bool dkgResultsRequested, bytes32 dkgAid, uint32 grace, uint64 lastVoteAt)
func (_ProcessRegistry *ProcessRegistryCaller) Processes(opts *bind.CallOpts, arg0 [31]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         DAVINCITypesEncryptionKey
	LatestStateRoot       [32]byte
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	MetadataHash          [32]byte
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
	Grace                 uint32
	LastVoteAt            uint64
}, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processes", arg0)

	outstruct := new(struct {
		Status                uint8
		OrganizationId        common.Address
		EncryptionKey         DAVINCITypesEncryptionKey
		LatestStateRoot       [32]byte
		StartTime             *big.Int
		Duration              *big.Int
		MaxVoters             *big.Int
		VotersCount           *big.Int
		OverwrittenVotesCount *big.Int
		CreationBlock         *big.Int
		BatchNumber           *big.Int
		MetadataURI           string
		MetadataHash          [32]byte
		BallotMode            DAVINCITypesBallotMode
		Census                DAVINCITypesCensus
		KeyMode               uint8
		DkgEpochId            [12]byte
		DkgFirstIndex         uint16
		DkgCount              uint8
		DkgZeroSkipped        uint16
		DkgResultsRequested   bool
		DkgAid                [32]byte
		Grace                 uint32
		LastVoteAt            uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.OrganizationId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EncryptionKey = *abi.ConvertType(out[2], new(DAVINCITypesEncryptionKey)).(*DAVINCITypesEncryptionKey)
	outstruct.LatestStateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MaxVoters = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VotersCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.OverwrittenVotesCount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.CreationBlock = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.BatchNumber = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.MetadataHash = *abi.ConvertType(out[12], new([32]byte)).(*[32]byte)
	outstruct.BallotMode = *abi.ConvertType(out[13], new(DAVINCITypesBallotMode)).(*DAVINCITypesBallotMode)
	outstruct.Census = *abi.ConvertType(out[14], new(DAVINCITypesCensus)).(*DAVINCITypesCensus)
	outstruct.KeyMode = *abi.ConvertType(out[15], new(uint8)).(*uint8)
	outstruct.DkgEpochId = *abi.ConvertType(out[16], new([12]byte)).(*[12]byte)
	outstruct.DkgFirstIndex = *abi.ConvertType(out[17], new(uint16)).(*uint16)
	outstruct.DkgCount = *abi.ConvertType(out[18], new(uint8)).(*uint8)
	outstruct.DkgZeroSkipped = *abi.ConvertType(out[19], new(uint16)).(*uint16)
	outstruct.DkgResultsRequested = *abi.ConvertType(out[20], new(bool)).(*bool)
	outstruct.DkgAid = *abi.ConvertType(out[21], new([32]byte)).(*[32]byte)
	outstruct.Grace = *abi.ConvertType(out[22], new(uint32)).(*uint32)
	outstruct.LastVoteAt = *abi.ConvertType(out[23], new(uint64)).(*uint64)

	return *outstruct, err

}

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, bytes32 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, bytes32 metadataHash, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, uint8 keyMode, bytes12 dkgEpochId, uint16 dkgFirstIndex, uint8 dkgCount, uint16 dkgZeroSkipped, bool dkgResultsRequested, bytes32 dkgAid, uint32 grace, uint64 lastVoteAt)
func (_ProcessRegistry *ProcessRegistrySession) Processes(arg0 [31]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         DAVINCITypesEncryptionKey
	LatestStateRoot       [32]byte
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	MetadataHash          [32]byte
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
	Grace                 uint32
	LastVoteAt            uint64
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, bytes32 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, bytes32 metadataHash, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, uint8 keyMode, bytes12 dkgEpochId, uint16 dkgFirstIndex, uint8 dkgCount, uint16 dkgZeroSkipped, bool dkgResultsRequested, bytes32 dkgAid, uint32 grace, uint64 lastVoteAt)
func (_ProcessRegistry *ProcessRegistryCallerSession) Processes(arg0 [31]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         DAVINCITypesEncryptionKey
	LatestStateRoot       [32]byte
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	MetadataHash          [32]byte
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
	Grace                 uint32
	LastVoteAt            uint64
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// ResultsProgramVK is a free data retrieval call binding the contract method 0xaa240221.
//
// Solidity: function resultsProgramVK() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) ResultsProgramVK(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "resultsProgramVK")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ResultsProgramVK is a free data retrieval call binding the contract method 0xaa240221.
//
// Solidity: function resultsProgramVK() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) ResultsProgramVK() ([32]byte, error) {
	return _ProcessRegistry.Contract.ResultsProgramVK(&_ProcessRegistry.CallOpts)
}

// ResultsProgramVK is a free data retrieval call binding the contract method 0xaa240221.
//
// Solidity: function resultsProgramVK() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) ResultsProgramVK() ([32]byte, error) {
	return _ProcessRegistry.Contract.ResultsProgramVK(&_ProcessRegistry.CallOpts)
}

// RootCVadcopFinal is a free data retrieval call binding the contract method 0x62115338.
//
// Solidity: function rootCVadcopFinal() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) RootCVadcopFinal(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "rootCVadcopFinal")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RootCVadcopFinal is a free data retrieval call binding the contract method 0x62115338.
//
// Solidity: function rootCVadcopFinal() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) RootCVadcopFinal() ([32]byte, error) {
	return _ProcessRegistry.Contract.RootCVadcopFinal(&_ProcessRegistry.CallOpts)
}

// RootCVadcopFinal is a free data retrieval call binding the contract method 0x62115338.
//
// Solidity: function rootCVadcopFinal() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) RootCVadcopFinal() ([32]byte, error) {
	return _ProcessRegistry.Contract.RootCVadcopFinal(&_ProcessRegistry.CallOpts)
}

// ZiskVerifier is a free data retrieval call binding the contract method 0x7f64b72f.
//
// Solidity: function ziskVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) ZiskVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "ziskVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZiskVerifier is a free data retrieval call binding the contract method 0x7f64b72f.
//
// Solidity: function ziskVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) ZiskVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.ZiskVerifier(&_ProcessRegistry.CallOpts)
}

// ZiskVerifier is a free data retrieval call binding the contract method 0x7f64b72f.
//
// Solidity: function ziskVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) ZiskVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.ZiskVerifier(&_ProcessRegistry.CallOpts)
}

// FinalizeResultsFromDKG is a paid mutator transaction binding the contract method 0xf1431097.
//
// Solidity: function finalizeResultsFromDKG(bytes31 processId) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) FinalizeResultsFromDKG(opts *bind.TransactOpts, processId [31]byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "finalizeResultsFromDKG", processId)
}

// FinalizeResultsFromDKG is a paid mutator transaction binding the contract method 0xf1431097.
//
// Solidity: function finalizeResultsFromDKG(bytes31 processId) returns()
func (_ProcessRegistry *ProcessRegistrySession) FinalizeResultsFromDKG(processId [31]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.FinalizeResultsFromDKG(&_ProcessRegistry.TransactOpts, processId)
}

// FinalizeResultsFromDKG is a paid mutator transaction binding the contract method 0xf1431097.
//
// Solidity: function finalizeResultsFromDKG(bytes31 processId) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) FinalizeResultsFromDKG(processId [31]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.FinalizeResultsFromDKG(&_ProcessRegistry.TransactOpts, processId)
}

// NewProcess is a paid mutator transaction binding the contract method 0x08c0fdd3.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadataURI, bytes32 metadataHash, (uint256,uint256) encryptionKey, (uint8,bytes12,uint256,uint256,uint256,uint256,uint256) dkg) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadataURI string, metadataHash [32]byte, encryptionKey DAVINCITypesEncryptionKey, dkg DAVINCITypesDKGParams) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, maxVoters, ballotMode, census, metadataURI, metadataHash, encryptionKey, dkg)
}

// NewProcess is a paid mutator transaction binding the contract method 0x08c0fdd3.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadataURI, bytes32 metadataHash, (uint256,uint256) encryptionKey, (uint8,bytes12,uint256,uint256,uint256,uint256,uint256) dkg) returns(bytes31)
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadataURI string, metadataHash [32]byte, encryptionKey DAVINCITypesEncryptionKey, dkg DAVINCITypesDKGParams) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadataURI, metadataHash, encryptionKey, dkg)
}

// NewProcess is a paid mutator transaction binding the contract method 0x08c0fdd3.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadataURI, bytes32 metadataHash, (uint256,uint256) encryptionKey, (uint8,bytes12,uint256,uint256,uint256,uint256,uint256) dkg) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadataURI string, metadataHash [32]byte, encryptionKey DAVINCITypesEncryptionKey, dkg DAVINCITypesDKGParams) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadataURI, metadataHash, encryptionKey, dkg)
}

// RequestResultsDecryption is a paid mutator transaction binding the contract method 0xe965eead.
//
// Solidity: function requestResultsDecryption(bytes31 processId, uint256[64] accumulator, bytes32[] siblings) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) RequestResultsDecryption(opts *bind.TransactOpts, processId [31]byte, accumulator [64]*big.Int, siblings [][32]byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "requestResultsDecryption", processId, accumulator, siblings)
}

// RequestResultsDecryption is a paid mutator transaction binding the contract method 0xe965eead.
//
// Solidity: function requestResultsDecryption(bytes31 processId, uint256[64] accumulator, bytes32[] siblings) returns()
func (_ProcessRegistry *ProcessRegistrySession) RequestResultsDecryption(processId [31]byte, accumulator [64]*big.Int, siblings [][32]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.RequestResultsDecryption(&_ProcessRegistry.TransactOpts, processId, accumulator, siblings)
}

// RequestResultsDecryption is a paid mutator transaction binding the contract method 0xe965eead.
//
// Solidity: function requestResultsDecryption(bytes31 processId, uint256[64] accumulator, bytes32[] siblings) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) RequestResultsDecryption(processId [31]byte, accumulator [64]*big.Int, siblings [][32]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.RequestResultsDecryption(&_ProcessRegistry.TransactOpts, processId, accumulator, siblings)
}

// RevealProcessKey is a paid mutator transaction binding the contract method 0x73417705.
//
// Solidity: function revealProcessKey(bytes31 processId, uint256 sk) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) RevealProcessKey(opts *bind.TransactOpts, processId [31]byte, sk *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "revealProcessKey", processId, sk)
}

// RevealProcessKey is a paid mutator transaction binding the contract method 0x73417705.
//
// Solidity: function revealProcessKey(bytes31 processId, uint256 sk) returns()
func (_ProcessRegistry *ProcessRegistrySession) RevealProcessKey(processId [31]byte, sk *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.RevealProcessKey(&_ProcessRegistry.TransactOpts, processId, sk)
}

// RevealProcessKey is a paid mutator transaction binding the contract method 0x73417705.
//
// Solidity: function revealProcessKey(bytes31 processId, uint256 sk) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) RevealProcessKey(processId [31]byte, sk *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.RevealProcessKey(&_ProcessRegistry.TransactOpts, processId, sk)
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

// SetProcessGrace is a paid mutator transaction binding the contract method 0x9a037789.
//
// Solidity: function setProcessGrace(bytes31 processId, uint32 grace) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessGrace(opts *bind.TransactOpts, processId [31]byte, grace uint32) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessGrace", processId, grace)
}

// SetProcessGrace is a paid mutator transaction binding the contract method 0x9a037789.
//
// Solidity: function setProcessGrace(bytes31 processId, uint32 grace) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessGrace(processId [31]byte, grace uint32) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessGrace(&_ProcessRegistry.TransactOpts, processId, grace)
}

// SetProcessGrace is a paid mutator transaction binding the contract method 0x9a037789.
//
// Solidity: function setProcessGrace(bytes31 processId, uint32 grace) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessGrace(processId [31]byte, grace uint32) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessGrace(&_ProcessRegistry.TransactOpts, processId, grace)
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

// SetProcessMetadata is a paid mutator transaction binding the contract method 0x46c15da8.
//
// Solidity: function setProcessMetadata(bytes31 processId, string metadataURI, bytes32 metadataHash) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessMetadata(opts *bind.TransactOpts, processId [31]byte, metadataURI string, metadataHash [32]byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessMetadata", processId, metadataURI, metadataHash)
}

// SetProcessMetadata is a paid mutator transaction binding the contract method 0x46c15da8.
//
// Solidity: function setProcessMetadata(bytes31 processId, string metadataURI, bytes32 metadataHash) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessMetadata(processId [31]byte, metadataURI string, metadataHash [32]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMetadata(&_ProcessRegistry.TransactOpts, processId, metadataURI, metadataHash)
}

// SetProcessMetadata is a paid mutator transaction binding the contract method 0x46c15da8.
//
// Solidity: function setProcessMetadata(bytes31 processId, string metadataURI, bytes32 metadataHash) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessMetadata(processId [31]byte, metadataURI string, metadataHash [32]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessMetadata(&_ProcessRegistry.TransactOpts, processId, metadataURI, metadataHash)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0x766422e0.
//
// Solidity: function setProcessResults(bytes31 processId, bytes publicValues, bytes proofBytes) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessResults(opts *bind.TransactOpts, processId [31]byte, publicValues []byte, proofBytes []byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessResults", processId, publicValues, proofBytes)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0x766422e0.
//
// Solidity: function setProcessResults(bytes31 processId, bytes publicValues, bytes proofBytes) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessResults(processId [31]byte, publicValues []byte, proofBytes []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, publicValues, proofBytes)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0x766422e0.
//
// Solidity: function setProcessResults(bytes31 processId, bytes publicValues, bytes proofBytes) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessResults(processId [31]byte, publicValues []byte, proofBytes []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, publicValues, proofBytes)
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

// SubmitStateTransition is a paid mutator transaction binding the contract method 0x1fdf3449.
//
// Solidity: function submitStateTransition(bytes31 processId, bytes publicValues, bytes proofBytes, bytes[] commitments, bytes32[] ys, bytes[] kzgProofs) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SubmitStateTransition(opts *bind.TransactOpts, processId [31]byte, publicValues []byte, proofBytes []byte, commitments [][]byte, ys [][32]byte, kzgProofs [][]byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "submitStateTransition", processId, publicValues, proofBytes, commitments, ys, kzgProofs)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0x1fdf3449.
//
// Solidity: function submitStateTransition(bytes31 processId, bytes publicValues, bytes proofBytes, bytes[] commitments, bytes32[] ys, bytes[] kzgProofs) returns()
func (_ProcessRegistry *ProcessRegistrySession) SubmitStateTransition(processId [31]byte, publicValues []byte, proofBytes []byte, commitments [][]byte, ys [][32]byte, kzgProofs [][]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransition(&_ProcessRegistry.TransactOpts, processId, publicValues, proofBytes, commitments, ys, kzgProofs)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0x1fdf3449.
//
// Solidity: function submitStateTransition(bytes31 processId, bytes publicValues, bytes proofBytes, bytes[] commitments, bytes32[] ys, bytes[] kzgProofs) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SubmitStateTransition(processId [31]byte, publicValues []byte, proofBytes []byte, commitments [][]byte, ys [][32]byte, kzgProofs [][]byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransition(&_ProcessRegistry.TransactOpts, processId, publicValues, proofBytes, commitments, ys, kzgProofs)
}

// SubmitStateTransitionL1 is a paid mutator transaction binding the contract method 0x53fb626d.
//
// Solidity: function submitStateTransitionL1(bytes31 processId, bytes publicValues, bytes proofBytes, bytes[] commitments, bytes32[] ys, bytes[] kzgProofs, (bytes[],(uint256,bytes,bytes[])[]) l1) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SubmitStateTransitionL1(opts *bind.TransactOpts, processId [31]byte, publicValues []byte, proofBytes []byte, commitments [][]byte, ys [][32]byte, kzgProofs [][]byte, l1 IL1BlobVerifierL1BlobProof) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "submitStateTransitionL1", processId, publicValues, proofBytes, commitments, ys, kzgProofs, l1)
}

// SubmitStateTransitionL1 is a paid mutator transaction binding the contract method 0x53fb626d.
//
// Solidity: function submitStateTransitionL1(bytes31 processId, bytes publicValues, bytes proofBytes, bytes[] commitments, bytes32[] ys, bytes[] kzgProofs, (bytes[],(uint256,bytes,bytes[])[]) l1) returns()
func (_ProcessRegistry *ProcessRegistrySession) SubmitStateTransitionL1(processId [31]byte, publicValues []byte, proofBytes []byte, commitments [][]byte, ys [][32]byte, kzgProofs [][]byte, l1 IL1BlobVerifierL1BlobProof) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransitionL1(&_ProcessRegistry.TransactOpts, processId, publicValues, proofBytes, commitments, ys, kzgProofs, l1)
}

// SubmitStateTransitionL1 is a paid mutator transaction binding the contract method 0x53fb626d.
//
// Solidity: function submitStateTransitionL1(bytes31 processId, bytes publicValues, bytes proofBytes, bytes[] commitments, bytes32[] ys, bytes[] kzgProofs, (bytes[],(uint256,bytes,bytes[])[]) l1) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SubmitStateTransitionL1(processId [31]byte, publicValues []byte, proofBytes []byte, commitments [][]byte, ys [][32]byte, kzgProofs [][]byte, l1 IL1BlobVerifierL1BlobProof) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransitionL1(&_ProcessRegistry.TransactOpts, processId, publicValues, proofBytes, commitments, ys, kzgProofs, l1)
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ProcessRegistryProcessGraceChangedIterator is returned from FilterProcessGraceChanged and is used to iterate over the raw logs and unpacked data for ProcessGraceChanged events raised by the ProcessRegistry contract.
type ProcessRegistryProcessGraceChangedIterator struct {
	Event *ProcessRegistryProcessGraceChanged // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessGraceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessGraceChanged)
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
		it.Event = new(ProcessRegistryProcessGraceChanged)
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
func (it *ProcessRegistryProcessGraceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessGraceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessGraceChanged represents a ProcessGraceChanged event raised by the ProcessRegistry contract.
type ProcessRegistryProcessGraceChanged struct {
	ProcessId [31]byte
	Grace     uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessGraceChanged is a free log retrieval operation binding the contract event 0xdf161c6af27d090672f982bb8002da0e7f2a530ffbd779e02ab9eba9397e51f1.
//
// Solidity: event ProcessGraceChanged(bytes31 indexed processId, uint32 grace)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessGraceChanged(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessGraceChangedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessGraceChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessGraceChangedIterator{contract: _ProcessRegistry.contract, event: "ProcessGraceChanged", logs: logs, sub: sub}, nil
}

// WatchProcessGraceChanged is a free log subscription operation binding the contract event 0xdf161c6af27d090672f982bb8002da0e7f2a530ffbd779e02ab9eba9397e51f1.
//
// Solidity: event ProcessGraceChanged(bytes31 indexed processId, uint32 grace)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessGraceChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessGraceChanged, processId [][31]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessGraceChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessGraceChanged)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessGraceChanged", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseProcessGraceChanged is a log parse operation binding the contract event 0xdf161c6af27d090672f982bb8002da0e7f2a530ffbd779e02ab9eba9397e51f1.
//
// Solidity: event ProcessGraceChanged(bytes31 indexed processId, uint32 grace)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessGraceChanged(log types.Log) (*ProcessRegistryProcessGraceChanged, error) {
	event := new(ProcessRegistryProcessGraceChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessGraceChanged", log); err != nil {
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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
	ProcessId    [31]byte
	MetadataURI  string
	MetadataHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProcessMetadataUpdated is a free log retrieval operation binding the contract event 0x77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f0.
//
// Solidity: event ProcessMetadataUpdated(bytes31 indexed processId, string metadataURI, bytes32 metadataHash)
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

// WatchProcessMetadataUpdated is a free log subscription operation binding the contract event 0x77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f0.
//
// Solidity: event ProcessMetadataUpdated(bytes31 indexed processId, string metadataURI, bytes32 metadataHash)
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseProcessMetadataUpdated is a log parse operation binding the contract event 0x77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f0.
//
// Solidity: event ProcessMetadataUpdated(bytes31 indexed processId, string metadataURI, bytes32 metadataHash)
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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
	OldStateRoot             [32]byte
	NewStateRoot             [32]byte
	NewVotersCount           *big.Int
	NewOverwrittenVotesCount *big.Int
	NBlobs                   *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterProcessStateTransitioned is a free log retrieval operation binding the contract event 0x36c6781d994e030a156d2f6fa11abcc1cd6814482a5d61f90da3f336318b1d23.
//
// Solidity: event ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, bytes32 oldStateRoot, bytes32 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount, uint256 nBlobs)
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

// WatchProcessStateTransitioned is a free log subscription operation binding the contract event 0x36c6781d994e030a156d2f6fa11abcc1cd6814482a5d61f90da3f336318b1d23.
//
// Solidity: event ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, bytes32 oldStateRoot, bytes32 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount, uint256 nBlobs)
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseProcessStateTransitioned is a log parse operation binding the contract event 0x36c6781d994e030a156d2f6fa11abcc1cd6814482a5d61f90da3f336318b1d23.
//
// Solidity: event ProcessStateTransitioned(bytes31 indexed processId, address indexed sender, bytes32 oldStateRoot, bytes32 newStateRoot, uint256 newVotersCount, uint256 newOverwrittenVotesCount, uint256 nBlobs)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStateTransitioned(log types.Log) (*ProcessRegistryProcessStateTransitioned, error) {
	event := new(ProcessRegistryProcessStateTransitioned)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateTransitioned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessStateTransitionedL1Iterator is returned from FilterProcessStateTransitionedL1 and is used to iterate over the raw logs and unpacked data for ProcessStateTransitionedL1 events raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateTransitionedL1Iterator struct {
	Event *ProcessRegistryProcessStateTransitionedL1 // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryProcessStateTransitionedL1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessStateTransitionedL1)
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
		it.Event = new(ProcessRegistryProcessStateTransitionedL1)
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
func (it *ProcessRegistryProcessStateTransitionedL1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessStateTransitionedL1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessStateTransitionedL1 represents a ProcessStateTransitionedL1 event raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateTransitionedL1 struct {
	ProcessId       [31]byte
	L1BlockNumbers  []*big.Int
	VersionedHashes [][32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProcessStateTransitionedL1 is a free log retrieval operation binding the contract event 0x4707050fda6801db89adf05b220e5793cac2e1ce5346f94f5c4d504b5e52519b.
//
// Solidity: event ProcessStateTransitionedL1(bytes31 indexed processId, uint256[] l1BlockNumbers, bytes32[] versionedHashes)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStateTransitionedL1(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryProcessStateTransitionedL1Iterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessStateTransitionedL1", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessStateTransitionedL1Iterator{contract: _ProcessRegistry.contract, event: "ProcessStateTransitionedL1", logs: logs, sub: sub}, nil
}

// WatchProcessStateTransitionedL1 is a free log subscription operation binding the contract event 0x4707050fda6801db89adf05b220e5793cac2e1ce5346f94f5c4d504b5e52519b.
//
// Solidity: event ProcessStateTransitionedL1(bytes31 indexed processId, uint256[] l1BlockNumbers, bytes32[] versionedHashes)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStateTransitionedL1(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStateTransitionedL1, processId [][31]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessStateTransitionedL1", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessStateTransitionedL1)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateTransitionedL1", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseProcessStateTransitionedL1 is a log parse operation binding the contract event 0x4707050fda6801db89adf05b220e5793cac2e1ce5346f94f5c4d504b5e52519b.
//
// Solidity: event ProcessStateTransitionedL1(bytes31 indexed processId, uint256[] l1BlockNumbers, bytes32[] versionedHashes)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStateTransitionedL1(log types.Log) (*ProcessRegistryProcessStateTransitionedL1, error) {
	event := new(ProcessRegistryProcessStateTransitionedL1)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateTransitionedL1", log); err != nil {
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
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ProcessRegistryResultsDecryptionRequestedIterator is returned from FilterResultsDecryptionRequested and is used to iterate over the raw logs and unpacked data for ResultsDecryptionRequested events raised by the ProcessRegistry contract.
type ProcessRegistryResultsDecryptionRequestedIterator struct {
	Event *ProcessRegistryResultsDecryptionRequested // Event containing the contract specifics and raw log

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
func (it *ProcessRegistryResultsDecryptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryResultsDecryptionRequested)
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
		it.Event = new(ProcessRegistryResultsDecryptionRequested)
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
func (it *ProcessRegistryResultsDecryptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryResultsDecryptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryResultsDecryptionRequested represents a ResultsDecryptionRequested event raised by the ProcessRegistry contract.
type ProcessRegistryResultsDecryptionRequested struct {
	ProcessId  [31]byte
	EpochId    [12]byte
	Aid        [32]byte
	FirstIndex uint16
	Count      uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterResultsDecryptionRequested is a free log retrieval operation binding the contract event 0xdca6075f07367349a836825d3ee7c35c204d2e727ae81a5b7a48aca95b9c270b.
//
// Solidity: event ResultsDecryptionRequested(bytes31 indexed processId, bytes12 epochId, bytes32 aid, uint16 firstIndex, uint8 count)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterResultsDecryptionRequested(opts *bind.FilterOpts, processId [][31]byte) (*ProcessRegistryResultsDecryptionRequestedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ResultsDecryptionRequested", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryResultsDecryptionRequestedIterator{contract: _ProcessRegistry.contract, event: "ResultsDecryptionRequested", logs: logs, sub: sub}, nil
}

// WatchResultsDecryptionRequested is a free log subscription operation binding the contract event 0xdca6075f07367349a836825d3ee7c35c204d2e727ae81a5b7a48aca95b9c270b.
//
// Solidity: event ResultsDecryptionRequested(bytes31 indexed processId, bytes12 epochId, bytes32 aid, uint16 firstIndex, uint8 count)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchResultsDecryptionRequested(opts *bind.WatchOpts, sink chan<- *ProcessRegistryResultsDecryptionRequested, processId [][31]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ResultsDecryptionRequested", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryResultsDecryptionRequested)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ResultsDecryptionRequested", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseResultsDecryptionRequested is a log parse operation binding the contract event 0xdca6075f07367349a836825d3ee7c35c204d2e727ae81a5b7a48aca95b9c270b.
//
// Solidity: event ResultsDecryptionRequested(bytes31 indexed processId, bytes12 epochId, bytes32 aid, uint16 firstIndex, uint8 count)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseResultsDecryptionRequested(log types.Log) (*ProcessRegistryResultsDecryptionRequested, error) {
	event := new(ProcessRegistryResultsDecryptionRequested)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ResultsDecryptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
