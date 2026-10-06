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

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_ziskVerifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_batchProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_resultsProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rootCVadcopFinal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_ballotVKHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_dkgManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_councilManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_defaultGrace\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceFloor\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceCeil\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceMaxTotal\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_noticeMin\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlobCountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CircuitFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CouncilDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DKGDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DecryptionNotOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyTransition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GraceOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccumulator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobCommitmentLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidBlobOpening\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobsDigest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDKGParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncryptionKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGrace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInclusionProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKZGProofLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOccupiedBefore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxPossibleResultCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"MissingBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoBlobs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsAlreadyRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtMaxLevelsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"}],\"name\":\"ProcessGraceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"ProcessMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nBlobs\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"firstIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"name\":\"ResultsDecryptionRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"aidFor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotVKHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"councilAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGrace\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dkgAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"finalizeResultsFromDKG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"}],\"name\":\"genesisRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastVoteAt\",\"type\":\"uint64\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessGraceEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceCeil\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceFloor\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceMaxTotal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint256\",\"name\":\"orgPKx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orgPKy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popZ\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.DKGParams\",\"name\":\"dkg\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noticeMin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastVoteAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256[64]\",\"name\":\"accumulator\",\"type\":\"uint256[64]\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"requestResultsDecryption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultsProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"sk\",\"type\":\"uint256\"}],\"name\":\"revealProcessKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootCVadcopFinal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"}],\"name\":\"setProcessGrace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"setProcessMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"commitments\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"kzgProofs\",\"type\":\"bytes[]\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ziskVerifier\",\"outputs\":[{\"internalType\":\"contractIZiskVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610200806040523461040e576101a08161733980380380916100218285610412565b83398101031261040e5761003481610435565b9061004160208201610446565b916040820151606083015160808401519060a08501519261006460c08701610446565b9561007160e08201610446565b9761007f6101008301610435565b9161008d6101208201610435565b61009a6101408301610435565b906100b56101806100ae6101608601610435565b9401610435565b60015f556001600160a01b039094169485158015610406575b80156103fe575b80156103f6575b80156103ee575b6103df5763ffffffff821680159081156103cf575b5080156103ba575b80156103a5575b8015610397575b6103885761016052610180526101a0526101c0526101e05260805260a05260c05260e0526101005260035467ffffffff000000006bffffffff0000000000000000604051602081019063ffffffff60e01b8660e01b1682523060601b60248201526018815261017e603882610412565b51902060401b169260201b1690640100000000600160601b031916171760035560018060a01b031680155f1461034757505f5b610120526001600160a01b0316806102e757505f5b61014052604051615432908161045b823960805181818161114b015281816114b30152612633015260a05181818161269f0152613f2b015260c051818181611523015261428a015260e05181818161150201528181611cf6015261267e015261010051818181610c5201528181612c0701526131b10152610120518181816109360152818161177c0152818161184f01528181613789015281816137e70152614d80015261014051818181610db401528181613640015281816136a401528181614d520152614dfd0152610160518181816109ed01526133ca01526101805181818161103f0152611d3701526101a0518181816110ce0152612bcc01526101c0518181816122a0015261491601526101e0518181816109830152610f170152f35b60405190610a448083016001600160401b038111848210176103335760209284926168f5843981520301905ff08015610328576001600160a01b03166101c6565b6040513d5f823e3d90fd5b634e487b7160e01b5f52604160045260245ffd5b604051906110688083016001600160401b0381118482101761033357602092849261588d843981520301905ff08015610328576001600160a01b03166101b1565b63795ee5af60e01b5f5260045ffd5b5063ffffffff85161561010e565b5063ffffffff841663ffffffff841611610107565b5063ffffffff831663ffffffff821611610100565b905063ffffffff8216105f6100f8565b6306f9b90760e01b5f5260045ffd5b5089156100e3565b5088156100dc565b5087156100d5565b5086156100ce565b5f80fd5b601f909101601f19168101906001600160401b0382119082101761033357604052565b519063ffffffff8216820361040e57565b51906001600160a01b038216820361040e5756fe6101406040526004361015610012575f80fd5b5f60a0525f3560e01c8063026cdee814613d9757806304ed00fa14613b8e578063082b642e146139c357806308c0fdd314612c2a5780630e2ebcf714612bf05780631542bbe214612bb05780631fdf3449146123ba5780633ea4ee411461238457806346c15da8146122c45780634c0acc561461110a578063549d59951461228257806359d821c414611d5b5780635ff5f98114611d195780636211533814611cdd57806362fa11fc14611c9457806368141f2c14611c155780636c7aff7f146118fe578063702574b31461183457806372c628ef146118175780637341770514611733578063766422e0146113ca578063784df74a1461117a5780637f64b72f14611134578063848df5401461110f578063946544bf1461110a5780639a03778914610fb95780639b46499414610de3578063aa240221146101c0578063abaab13a14610d9d578063adc879e914610d76578063bf74291e14610a11578063c8f0582f146109cf578063cddf08bc146109a7578063d4138a2014610965578063e16d5b7c1461091f578063e965eead146102f1578063f1431097146101c55763f9aa4499146101c0575f80fd5b614273565b346102eb5760203660031901126102eb576101de613ea6565b6101e661468f565b6101ef816148b6565b601881015460ff811660048110156102d357156102c057815460ff1661021481613f4e565b600281149081156102ac575b506102995761022e826148fb565b42106102865760901c60ff16156102735761024881614db1565b156102605761025691614e6c565b60a0516001815580f35b6314badd7360e31b60a05152600460a051fd5b630e0d4dc560e41b60a05152600460a051fd5b63611f72f360e11b60a05152600460a051fd5b6307a92f1960e51b60a05152600460a051fd5b600491506102b981613f4e565b1484610220565b6365b75c3960e01b60a05152600460a051fd5b634e487b7160e01b60a051526021600452602460a051fd5b60a05180fd5b346102eb576108403660031901126102eb5761030b613ea6565b36610824116102eb57610824356001600160401b0381116102eb57610334903690600401613ee4565b61033c61468f565b610345836148b6565b91601883015491600460ff841610156102d35760ff8316156102c05760ff8360901c1661090c57835460ff81169290919061037f84613f4e565b6002841480156108f9575b6102995761039784613f4e565b60018414159182806108de575b6108cb576103b1876148fb565b42106102865760a0515b604081106108a5575060206040518181016108006024823761080082526103e46108208361409e565b60405191518091835e81019060a05182528060a05192039060025afa1561065a576104179160a051516003890154614c7a565b156108925760ff60901b198416600160901b17601886015561043883613f4e565b610853575b505060ff600e83015460081c169061045482614454565b91610462604051938461409e565b808352601f1961047182614454565b0160a0515b81811061083157505060a05190819081905b8082106106ac57505060a051821594909390851561056b575b50509061ffff60ff926018870154908260801b9060801b16908260681b8660681b169064ffffffffff60681b1916178460781b8460781b161717938460188801556019870154604051956001600160601b0360a01b9060981b16865260208601521660408401521660608201527fdca6075f07367349a836825d3ee7c35c204d2e727ae81a5b7a48aca95b9c270b608060ff19861692a28061055c575b61054c5760a0516001815580f35b61055591614e6c565b8080610256565b5061056681614db1565b61053e565b838152919350906001600160a01b0361058660ff8416614d40565b601988015460405163574bbf6f60e11b815260ff60901b19909516600160901b1760981b6001600160a01b031916600486015260248501526060604485015282516064850181905260a051929091169284926084840192602090920191905b81811061066757505050918180602094039160a051905af190811561065a5760a05191610619575b509161ffff60ff6104a1565b90506020813d602011610652575b816106346020938361409e565b810103126102eb57519061ffff821682036102eb579061ffff61060d565b3d9150610627565b6040513d60a051823e3d90fd5b9180945092909251819060a051915b6004831061069657505050602060806001920194019101918593926105e5565b6020806001928451815201920192019190610676565b90926001600160fe1b038416840361079c576106ca8460021b61455b565b3590600285901b6001810190811061079c576106e59061455b565b3560a051600287811b0191828860021b1161079c576107038361455b565b351580610807575b8515806107fd575b6107df576107cc5760405194608086018681106001600160401b038211176107b4576040528552602085015261079c5761074c9061455b565b356040830152600285901b600381019190821061079c576001926107726107939361455b565b356060820152610782828a61450e565b5261078d818961450e565b50614446565b935b0190610488565b634e487b7160e01b60a051526011600452602460a051fd5b634e487b7160e01b60a051526041600452602460a051fd5b632a23591560e21b60a05152600460a051fd5b9397969450505050156107cc5760019061ffff82851b161792610795565b5060018214610713565b50905060a0519060038860021b01808960021b1161079c5761082a60019161455b565b351461070b565b604051602091906080610844818361409e565b36823782828801015201610476565b60ff191660011783556040519061086981613f4e565b81526001602082015260ff198416905f5160206153bd5f395f51905f5290604090a2838061043d565b6303cd656760e61b60a05152600460a051fd5b5f5160206153dd5f395f51905f526108bc8261455b565b3510156107cc576001016103bb565b63e843c5eb60e01b60a05152600460a051fd5b506108f260058801546006890154906142bb565b42106103a4565b5061090384613f4e565b6004841461038a565b636c8ae94b60e11b60a05152600460a051fd5b346102eb5760a0513660031901126102eb576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102eb5760a0513660031901126102eb57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102eb5760a0513660031901126102eb57602063ffffffff60035460401c16604051908152f35b346102eb5760a0513660031901126102eb57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102eb576101803660031901126102eb57610a2b613ea6565b6101003660231901126102eb576040366101231901126102eb576101643560058110156102eb5760405190610a5f8261404c565b61012435825261014435602083019081526040519290610a8060e08561409e565b6006845260c0948536602087013760405195610a9d60e08861409e565b6006875236602088013760081c610ab38661449d565b5260a051610ac08561449d565b5260643560ff811692908381036102eb576044359360ff8516908186036102eb578110610d635760a43565ffffffffffff8111610d505760c4359165ffffffffffff8311610d3d5760e43593677fffffffffffffff8511610d2a576101043597677fffffffffffffff8911610d17575060243580151581036102eb5715610d0e576001905b6084359260ff841684036102eb5761ff0062ff00006301fe000060209c60b81b9960791b9860491b9760191b9660111b169460101b169260081b1617171717171717610b90876144be565b526002610b9c866144be565b525190519060405191838301918252604083015260408252610bbf60608361409e565b60405191518091835e81019060a05182528060a05192039060025afa1561065a5760a05151610bed846144ce565b526003610bf9836144ce565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae278610c24846144de565b526004610c30836144de565b52610c3a81613f4e565b610c43836144ee565b526006610c4f826144ee565b527f0000000000000000000000000000000000000000000000000000000000000000610c7a836144fe565b526007610c86826144fe565b528051825103610cfb57610c9a815161446b565b60a0515b8251811015610ce35780610cd26001600160401b03610cbf6001948761450e565b5116610ccb838861450e565b5190615163565b610cdc828561450e565b5201610c9e565b6020610cf3848460a051916151f6565b604051908152f35b63d088249360e01b60a05152600460a051fd5b60a05190610b45565b63dd6f54df60e01b60a05152600460a051fd5b63271fb80560e01b60a05152600460a051fd5b63871a7fa360e01b60a05152600460a051fd5b63481eb79f60e01b60a05152600460a051fd5b632cbdc23160e01b60a05152600460a051fd5b346102eb5760a0513660031901126102eb57602063ffffffff600354821c16604051908152f35b346102eb5760a0513660031901126102eb576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102eb5760403660031901126102eb57610dfc613ea6565b60243560ff198216918215610fa65763ffffffff808060035460401c16169160401c1603610f935760a08051839052600160205251604090208054600881901c6001600160a01b03168015610f80573303610f6e5760ff16610e5d81613f4e565b8015159081610f59575b506102995760066005820154910190610e818254826142bb565b90428211156108cb5783610e94916142bb565b8315918215610f4e575b8215610f44575b8215610ef9575b5050610ee657817f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba9260209255604051908152a260a05180f35b637616640160e01b60a05152600460a051fd5b8110915081610f0b575b508480610eac565b9050610f3d63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016426142bb565b1184610f03565b8181149250610ea5565b428211159250610e9e565b60039150610f6681613f4e565b141584610e67565b6282b42960e81b60a05152600460a051fd5b634d36eb6960e01b60a05152600460a051fd5b632299770d60e11b60a05152600460a051fd5b63cbf4a64560e01b60a05152600460a051fd5b346102eb5760403660031901126102eb57610fd2613ea6565b6024359063ffffffff82168092036102eb57610fed816148b6565b805433600882901c6001600160a01b031603610f6e5760ff1661100f81613f4e565b80151590816110f5575b506102995761103160058201546006830154906142bb565b4210156108cb5763ffffffff7f000000000000000000000000000000000000000000000000000000000000000016831080156110c6575b6110b357601a01805463ffffffff19168317905560405191825260ff1916907fdf161c6af27d090672f982bb8002da0e7f2a530ffbd779e02ab9eba9397e51f190602090a260a05180f35b63795ee5af60e01b60a05152600460a051fd5b5063ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168311611068565b6003915061110281613f4e565b141584611019565b613f14565b346102eb5760a0513660031901126102eb57602063ffffffff60035416604051908152f35b346102eb5760a0513660031901126102eb576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102eb5760203660031901126102eb5760ff19611196613ea6565b1660a05152600160205260a051604090208054600182016111b6906140bf565b60e0526003820154610120526005820154600683015460078401549060088501546009860154600a870154600b88015491600c89016111f490614115565b93600d8a015496600e8b01611208906141b5565b9661121560138d01614217565b9960188d015460c05260198d01549c601a01549b6040516101005260ff811661123d90613f4e565b60ff81166101005152600160a01b600190039060081c16610100516020015260e05151610100516040015260e0516020015161010051606001526101205161010051608001526101005160a001526101005160c001526101005160e001526101005161010001526101005161012001526101005161014001526101005161016001526101005161018001610400905261010051610400016112dd91613f58565b91610100516101a00152610100516101c0016112f891613f7c565b610100518103610100516102c0015261131091613fd2565b91610100516102e00160c05160ff169061132991614023565b6001600160601b0360a01b60c05160981b1661010051610300015260c05160681c61ffff1661010051610320015260c05160781c60ff1661010051610340015260c05160801c61ffff1661010051610360015260c05160901c60ff161515610100516103800152610100516103a0015263ffffffff8116610100516103c0015260201c6001600160401b0316610100516103e0015261010051900361010051f35b346102eb5760603660031901126102eb576113e3613ea6565b6024356001600160401b0381116102eb57611402903690600401613eb7565b90916044356001600160401b0381116102eb57611423903690600401613eb7565b9261142c61468f565b611435836148b6565b9160ff60188401541660048110156102d3576102c057825460ff81169590929061145e87613f4e565b600287148015611720575b6102995761147687613f4e565b600187141580611705575b6108cb5761148e856148fb565b42106102865761149e828961499a565b6114a7886149f9565b6003860154036116f2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156102eb5761154b9389936040519586948593849363be98686160e01b855260a051987f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801614410565b03915afa801561065a576116d9575b5060ff600e83015460081c16946115708661446b565b9560a0515b8181106116705750505060ff191660049081178255845191016001600160401b0382116107b457600160401b82116107b4578054828255808310611650575b50602085019060a05152602060a0512060a0515b83811061163c57866040875f5160206153bd5f395f51905f528860ff19169283928151906115f581613f4e565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e6040518061162f339582614522565b0390a360a0516001815580f35b6001906020845194019381840155016115c8565b61166a908260a0515283602060a0512091820191016143da565b856115b4565b600181901b906001600160ff1b038116810361079c5781600a019182600a1161079c57600b6116a78460031b87013560c01c615116565b910192831061079c576116c460019360031b86013560c01c615116565b60201b176116d2828b61450e565b5201611575565b60a0516116e59161409e565b60a0516102eb578561155a565b630b6fac0360e41b60a05152600460a051fd5b5061171960058601546006870154906142bb565b4210611481565b5061172a87613f4e565b60048714611469565b346102eb5760403660031901126102eb5761175c61174f613ea6565b61175761468f565b6148b6565b60188101549060ff821660048110156102d3576002036102c057601901547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156102eb576040519263193f942b60e21b84526001600160601b0360a01b9060981b166004840152602483015260243560448301528160648160a0519360a051905af1801561065a576117fe5760a0516001815580f35b60a05161180a9161409e565b60a0516102eb5780610256565b346102eb5760a0513660031901126102eb57602060405160048152f35b346102eb5760203660031901126102eb5761184d613ea6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169081156118ec5760209060246040518094819363702574b360e01b835260ff191660048301525afa801561065a5760a051906118b9575b602090604051908152f35b506020813d6020116118e4575b816118d36020938361409e565b810103126102eb57602090516118ae565b3d91506118c6565b6203eb8f60e01b60a05152600460a051fd5b346102eb5760403660031901126102eb57611917613ea6565b6024356001600160401b0381116102eb578060040160a060031983360301126102eb57611943836148b6565b8054919033600884901c6001600160a01b031603610f6e57601381015460029060ff1661196f81613f4e565b03611c0257813560058110156102eb5760029061198b81613f4e565b03611bef576001600160a01b036119a460448601614394565b16611bdc576024840135938415611bc957606401926119c384846143a8565b905015611bb65760ff166119d681613f4e565b8015159081611ba1575b50610299576119f860058201546006830154906142bb565b4210156108cb578360148201556016611a1184846143a8565b91909201916001600160401b0382116107b457611a2e83546140dd565b601f8111611b62575b5060a05190601f8311600114611acc579282611ab89896937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2989693611a989660a05192611ac1575b50508160011b915f199060031b1c19161790556143a8565b949060405193849384526040602085015260ff19169560408401916143f0565b0390a260a05180f35b013590508a80611a80565b601f198316918460a05152602060a051209260a0515b818110611b4a5750937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed2989693611a98969360019383611ab89d9b9810611b31575b505050811b0190556143a8565b01355f19600384901b60f8161c191690558a8080611b24565b91936020600181928787013581550195019201611ae2565b611b91908460a05152602060a05120601f850160051c81019160208610611b97575b601f0160051c01906143da565b87611a37565b9091508190611b84565b60039150611bae81613f4e565b1415866119e0565b630f8b932160e11b60a05152600460a051fd5b635e32eadd60e01b60a05152600460a051fd5b63562e597160e11b60a05152600460a051fd5b63f37f7b5d60e01b60a05152600460a051fd5b63050b77c760e21b60a05152600460a051fd5b346102eb5760203660031901126102eb576004356001600160a01b038116908181036102eb5760035460a0805193909352600260209081529251604090205466ffffffffffffff1660589290921b600160581b600160f81b0316600891821c63ffffffff60381b161791909117901b60ff19166040519060ff19168152f35b346102eb5760203660031901126102eb576004356001600160a01b038116908190036102eb5760a05152600260205260206001600160401b03604060a051205416604051908152f35b346102eb5760a0513660031901126102eb5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346102eb5760a0513660031901126102eb57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102eb5760203660031901126102eb57611d74613ea6565b604051611d8081614030565b60a051815260a0516020820152604051611d998161404c565b60a051815260a0516020820152604082015260a05160608201526060608082015260a05160a082015260a05160c082015260a05160e082015260a05161010082015260a05161012082015260a05161014082015260a051610160820152606061018082015260a0516101a0820152604051611e1381614067565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260a05160c082015260a05160e08201526101c0820152604051611e6281614083565b60a051815260a051602082015260a051604082015260608082015260a05160808201526101e082015260a05161020082015260a05161022082015260a05161024082015260a05161026082015260a05161028082015260a0516102a082015260a0516102c082015260a0516102e082015261030060a05191015260ff191660a051526001602052604060a0512060405190611efc82614030565b805460ff8116611f0b81613f4e565b835260081c6001600160a01b03166020830152611f2a600182016140bf565b604083015260038101546060830152600481016040518082602082945493848152019060a05152602060a051209260a0515b818110612269575050611f719250038261409e565b6080830152600581015460a0830152600681015460c0830152600781015460e083015260088101546101008301526009810154610120830152600a810154610140830152600b810154610160830152611fcc600c8201614115565b610180830152600d8101546101a0830152611fe9600e82016141b5565b6101c0830152611ffb60138201614217565b6101e08301526018810154600460ff821610156102d3576001600160401b039160ff8281601a94166102008701526001600160601b0360a01b8160981b1661022087015261ffff8160681c16610240870152818160781c1661026087015261ffff8160801c1661028087015260901c1615156102a085015260198101546102c0850152015463ffffffff81166102e084015260201c166103008201526040516020815261044081019180516120af81613f4e565b602083015260018060a01b036020820151166040830152602060408201518051606085015201516080830152606081015160a083015260808101519261042060c084015283518091526020610460840194019060a0515b818110612253575050506001600160401b036103006121c161218b859660a086015160e088015260c086015161010088015260e08601516101208801526101008601516101408801526101208601516101608801526101408601516101808801526101608601516101a0880152610180860151601f19888303016101c0890152613f58565b6101a08501516101e08701526121ab6101c0860151610200880190613f7c565b6101e0850151868203601f190184880152613fd2565b926121d6610200820151610320870190614023565b6102208101516001600160a01b03191661034086015261024081015161ffff90811661036087015261026082015160ff16610380870152610280820151166103a08601526102a081015115156103c08601526102c08101516103e08601526102e081015163ffffffff166104008601520151166104208301520390f35b8251865260209586019590920191600101612106565b8454835260019485019486945060209093019201611f5c565b346102eb5760a0513660031901126102eb57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102eb5760603660031901126102eb576122dd613ea6565b6024356001600160401b0381116102eb576122fc903690600401613eb7565b60443591612309846148b6565b805433600882901c6001600160a01b031603610f6e5760ff9061232d8686866146ad565b1661233781613f4e565b801515908161236f575b506102995761235960058201546006830154906142bb565b4210156108cb5761236994614750565b60a05180f35b6003915061237c81613f4e565b141586612341565b346102eb5760203660031901126102eb5760ff196123a0613ea6565b1660a0515260016020526020610cf3604060a051206148fb565b346129fb5760c03660031901126129fb576123d3613ea6565b6080526024356001600160401b0381116129fb576123f5903690600401613eb7565b906044356001600160401b0381116129fb57612415903690600401613eb7565b90916064356001600160401b0381116129fb57612436903690600401613ee4565b9190946084356001600160401b0381116129fb57612458903690600401613ee4565b9060a4356001600160401b0381116129fb57612478903690600401613ee4565b919061248261468f565b61248d6080516148b6565b9660ff88541661249c81613f4e565b8015159081612b9a575b81612b63575b50612b545760058801544210612b45576124c5886148fb565b421015612b45576124d6868861499a565b6124df876149f9565b9a60038901548c03612b36575f5f5b60088110612b15575061250a9061250490614a2e565b8a614b94565b600889015495866125226101508b013560c01c615116565b03612b065761253760988a013560c01c615116565b9a6125518c61254c60908d013560c01c615116565b6142e0565b9861255c8d8b6142bb565b15612af75761256b8a8a6142bb565b60078d015410612ae8576125866101208c013560c01c615116565b9d8e15612ad9578e80871490811591612ace575b8115612ac3575b50612ab4578e49612ab4578e6050810204605003612aa0578e6125c660508202614c34565b906125d4604051928361409e565b605081028083526125e490614c34565b601f19013660208401375f5b818110612a305750505f60208092604051918183925191829101835e8101838152039060025afa156129f0575f8051818e5b60088210612a0e575050036129ff577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156129fb575f928d6126c76040519687958694859463be98686160e01b86527f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801614410565b03915afa80156129f0576129dc575b506126e08d614a2e565b9560a0515b8481106127de57505060a051988996509450505050505b600882106127bb575050601a91612718918460038701556142bb565b9283600882015561272e600982019586546142bb565b809555600b810161273f8154614446565b90550180546bffffffffffffffff000000004260201b16906bffffffffffffffff000000001916179055604051948552602085015260408401526060830152608082015233907f36c6781d994e030a156d2f6fa11abcc1cd6814482a5d61f90da3f336318b1d2360a060ff196080511692a360a0516001815580f35b909360019085600a0160031b83013560e01c8660051b60e0031b179401906126fc565b8049156129c657602060608961282e6127f8858a8a614c4f565b9390846040519586928884019660805160081c8852604085015284840137810160a051838201520301601f19810184528361409e565b60405191518091835e81019060a05182528060a05192039060025afa1561065a5786866128e2866080866128a187612899818e612892828f7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff0000000160a05151069d614c6a565b3597614c4f565b939097614c4f565b828193604051988997602089019b8d498d5260408a015260608901528688013785019184830160a051815237010160a051815203601f19810183528261409e565b60a0519160a051915190600a5afa3d156129be573d9061290182614c34565b9161290f604051938461409e565b825260a0513d90602084013e5b1580156129b2575b61299b576040818051810103126102eb5761100060406020830151920151911490811591612970575b5061295a576001016126e5565b638308e1e960e01b60a05152600452602460a051fd5b7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001915014155f61294d565b50638308e1e960e01b60a05152600452602460a051fd5b50604081511415612924565b60609061291c565b634af69d9960e11b60a05152600452602460a051fd5b5f6129e69161409e565b5f60a0528d6126d6565b6040513d5f823e3d90fd5b5f80fd5b630f2e4cb960e31b5f5260045ffd5b928193600192601c0160031b013560e01c8460051b60e0031b1792018e612622565b80612a3e6030928b8b614c4f565b92909203612a91576030612a53828f8e614c4f565b905003612a82576001916050612a6b838f8c90614c6a565b3591603082850288019160208301370152016125f0565b6350320ab160e01b5f5260045ffd5b6338b2168b60e21b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b63b8ff08e960e01b5f5260045ffd5b90508914158f6125a1565b85811415915061259a565b63fdac229f60e01b5f5260045ffd5b6357d18d5360e11b5f5260045ffd5b633f8cdc5560e11b5f5260045ffd5b63341203dd60e21b5f5260045ffd5b906001908260140160031b8b013560e01c8360051b60e0031b1791016124ee565b630b6fac0360e41b5f5260045ffd5b63e843c5eb60e01b5f5260045ffd5b6307a92f1960e51b5f5260045ffd5b60039150612b7081613f4e565b1480612b7e575b158c6124ac565b50612b92600589015460068a0154906142bb565b421015612b77565b9050612ba581613f4e565b6001811415906124a6565b346129fb575f3660031901126129fb57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346129fb575f3660031901126129fb5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346129fb576103003660031901126129fb57600560043510156129fb576101003660831901126129fb576001600160401b0361018435116129fb5760a061018435360360031901126129fb576101a4356001600160401b0381116129fb57612c96903690600401613eb7565b6040366101e31901126129fb5760e0366102231901126129fb57612cb861468f565b600354335f81815260026020526040902054602435939266ffffffffffffff90911660589290921b600160581b600160f81b0316600891821c63ffffffff60381b161791909117901b60ff191660ff1981165f9081526001602052604090205490929060081c6001600160a01b031633146139b45760a43560ff8116141593846129fb5760ff60a435161580156139a2575b6139935760c43560ff8116141594856129fb576129fb5760ff60a4351660ff60c43516116135b75761010435610124351161398457610144356101643511613975576064351561396657612da36101043560643561456d565b6005610184356004013510156129fb57612dc36101843560040135613f4e565b61018435600401351561395757612ddf60846101843501614387565b6139485760246101843501359384612dfd6101843560040135613f4e565b60046101843501356003036138e45750612e1c60446101843501614394565b803b156138d55760405163c1da869160e01b602082015260048152612e4b91612e4660248361409e565b6150fb565b90156138d557935b612e6960646101843501610184356004016143a8565b9050156138c657612e7b600435613f4e565b600460ff813516118015613897575b612b5457612e9c6101c43582856146ad565b6024351561388f575b4284106138805742612eb9604435866142bb565b11156138715760ff1982165f52600160205260405f20604051612edb8161404c565b6101e43581526102043560208201529460046102243510156129fb576102243561362057610244356001600160601b0360a01b81168091036129fb5715801590613614575b8015613608575b80156135fc575b80156135f0575b80156135e4575b6135d5575b612f4a866146d6565b156135c657612f5b600435836142c8565b6005820155604435600682015560643560078201558054610100600160a81b0319163360081b610100600160a81b0316178155845160018201556020850151600282015560405194612fae60e08761409e565b6006865260c036602088013760405198612fc960e08b61409e565b60068a5260c03660208c01378460081c612fe28b61449d565b525f612fed8861449d565b52806129fb5760ff60a4351660a435036129fb5760ff60a4351660ff60c43516116135b75765ffffffffffff61010435116135a85765ffffffffffff610124351161359957677fffffffffffffff610144351161358a57677fffffffffffffff610164351161357b5760ff60a4351660a435036129fb576129fb576084351515608435036129fb57608435156135755760015b60ff60e4351660e435036129fb576020915f9160a43560ff1660c43560081b61ff00161760109190911b62ff0000161760e43560111b6301fe000016176101043560191b176101243560491b176101443560791b176101643560b81b176130e68c6144be565b5260026130f2896144be565b528281519101516040519084820192835260408201526040815261311760608261409e565b604051918291518091835e8101838152039060025afa156129f0575f5161313d896144ce565b526003613149866144ce565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae278613174896144de565b526004613180866144de565b526131916101843560040135613f4e565b61018435600401356131a2896144ee565b5260066131ae866144ee565b527f00000000000000000000000000000000000000000000000000000000000000006131d9896144fe565b5260076131e5866144fe565b528451885103613566576131f9855161446b565b955f5b865181101561323b578061322a8b610ccb836001600160401b036132226001978e61450e565b51169261450e565b613234828b61450e565b52016131fc565b50876132485f89896151f6565b6003840155600e83016084351515608435036129fb57805460ff191660ff608435151516178155805460ff60e4351660e435036129fb5763ff00000060e43560181b169061ff0060a43560081b169063ffffff0019161762ff000060c43560101b161717905561010435600f840155610124356010840155610144356011840155610164356012840155601383016132e66101843560040135613f4e565b60ff1981541660ff610184356004013516179055601483019081556015830160018060a01b0361331b60446101843501614394565b82546001600160a01b0319169116179055601683016133446101843560648101906004016143a8565b906001600160401b03821161355257819061335f84546140dd565b601f8111613522575b505f90601f83116001146134bb575f926134b0575b50508160011b915f199060031b1c19161790555b6133b86133a360846101843501614387565b601785019060ff801983541691151516179055565b5543600a820155601a810163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff1982541617905560035463ffffffff811663ffffffff8114612aa057600163ffffffff9101169063ffffffff191617600355335f52600260205260405f20918254946001600160401b038616936001600160401b038514612aa0576020966001600160401b0360016134a0970116906001600160401b0319161790553360ff1986167feefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d15f80a36101c4359285614750565b60015f556040519060ff19168152f35b01359050898061337d565b909150601f19831691845f5260205f20925f5b81811061350a57509084600195949392106134f1575b505050811b019055613391565b01355f19600384901b60f8161c191690558980806134e4565b919360206001819287870135815501950192016134ce565b61354c90855f5260205f20601f850160051c81019160208610611b9757601f0160051c01906143da565b8a613368565b634e487b7160e01b5f52604160045260245ffd5b63d088249360e01b5f5260045ffd5b5f613080565b63dd6f54df60e01b5f5260045ffd5b63271fb80560e01b5f5260045ffd5b63871a7fa360e01b5f5260045ffd5b63481eb79f60e01b5f5260045ffd5b632cbdc23160e01b5f5260045ffd5b63208e53e560e11b5f5260045ffd5b63e4291a1960e01b5f5260045ffd5b506102e4351515612f3c565b506102c4351515612f35565b506102a4351515612f2e565b50610284351515612f27565b50610264351515612f20565b94506101e43515801590613865575b6135c65761022435600303613787577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156137785760405163ecd8e20560e01b815260ff19841660048201523360248201529461369a60448701610224614320565b608086610124815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af180156129f0575f905f905f985f9161373f575b50979091975b604051916136f48361404c565b825260208201526018840180546cffffffffffffffffffffffffff19166102243560ff161760989990991c6cffffffffffffffffffffffff0016989098179097556019830155612f41565b9250505061376691965060803d608011613771575b61375e818361409e565b8101906142ed565b97929190978c6136e1565b503d613754565b63e00ffde960e01b5f5260045ffd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156138575760405163441d214d60e11b815260ff1984166004820152946137dd60248701610224614320565b608086610104815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af180156129f0575f905f905f985f9161382e575b50979091976136e7565b9250505061384c91965060803d6080116137715761375e818361409e565b97929190978c613824565b6203eb8f60e01b5f5260045ffd5b5061020435151561362f565b637616640160e01b5f5260045ffd5b632ca4094f60e21b5f5260045ffd5b429350612ea5565b506138a3600435613f4e565b60043515158015612e8a57506138ba600435613f4e565b60036004351415612e8a565b630f8b932160e11b5f5260045ffd5b63562e597160e11b5f5260045ffd5b936001600160a01b036138fc61018435604401614394565b166138d557841561392d576139176101843560040135613f4e565b600461018435810135148061393c575b15612e53575b635e32eadd60e01b5f5260045ffd5b508460a01c1515613927565b63f545b7bf60e01b5f5260045ffd5b63f37f7b5d60e01b5f5260045ffd5b632c45be6f60e01b5f5260045ffd5b636d403ec760e11b5f5260045ffd5b63207ea56d60e01b5f5260045ffd5b63ac38930b60e01b5f5260045ffd5b505f9450601060a43560ff1611612d4a565b635040c4d360e11b5f5260045ffd5b346129fb5760403660031901126129fb576139dc613ea6565b60243560058110156129fb5760ff198216918215613b7f5763ffffffff808060035460401c16169160401c1603613b7057613a1681613f4e565b600460ff821611612b54575f8281526001602052604090208054600881901c6001600160a01b03168015613b61573303613b535760ff16613a5783826145a4565b15612b54576005820154926006830192613a728454866142bb565b93613a7c83613f4e565b60018314958680613b4a575b612b4557613a9584613f4e565b6003841480613b40575b612b45575f5160206153bd5f395f51905f5296604096613ac0868b966142c8565b613ac986613f4e565b81613b36575b50613af7575b505050825191613ae481613f4e565b8252613aef81613f4e565b6020820152a2005b7f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba91613b25602092426142e0565b8091558651908152a2848680613ad5565b9050421089613acf565b5085421015613a9f565b50804210613a88565b6282b42960e81b5f5260045ffd5b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b63cbf4a64560e01b5f5260045ffd5b346129fb5760203660031901126129fb5760ff19613baa613ea6565b165f52600160205260405f20604051613bc281614030565b815460ff8116613bd181613f4e565b825260081c6001600160a01b03166020820152613bf0600183016140bf565b60408201526003820154606082015260048201604051808260208294549384815201905f5260205f20925f5b818110613d7e575050613c319250038261409e565b6080820152600582015460a0820190815260068301549060c08301918252600784015460e084015260088401546101008401526009840154610120840152600a840154610140840152600b840154610160840152613c91600c8501614115565b610180840152600d8401546101a0840152613cae600e85016141b5565b6101c0840152613cc060138501614217565b6101e084015260188401549260ff8416946004861015613d6a576001600160401b03601a6103009260ff610cf39860209a6102008801526001600160601b0360a01b8160981b1661022088015261ffff8160681c16610240880152818160781c1661026088015261ffff8160801c1661028088015260901c1615156102a086015260198101546102c0860152015463ffffffff81166102e0850152871c16910152519051906142bb565b634e487b7160e01b5f52602160045260245ffd5b8454835260019485019486945060209093019201613c1c565b346129fb5760403660031901126129fb57613db0613ea6565b60243560ff198216918215613b7f5763ffffffff808060035460401c16169160401c1603613b70575f8281526001602052604090208054600881901c6001600160a01b03168015613b61573303613b535760ff16613e0d81613f4e565b8015159081613e91575b50612b5457613e2f60058201546006830154906142bb565b421015612b455781158015613e84575b61396657817f36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be92600783613e79600f60209601548461456d565b0155604051908152a2005b5060088101548210613e3f565b60039150613e9e81613f4e565b141584613e17565b6004359060ff19821682036129fb57565b9181601f840112156129fb578235916001600160401b0383116129fb57602083818601950101116129fb57565b9181601f840112156129fb578235916001600160401b0383116129fb576020808501948460051b0101116129fb57565b346129fb575f3660031901126129fb5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b60051115613d6a57565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60e0809180511515845260ff602082015116602085015260ff604082015116604085015260ff60608201511660608501526080810151608085015260a081015160a085015260c081015160c08501520151910152565b908151613fde81613f4e565b81526020820151602082015260018060a01b036040830151166040820152608080614018606085015160a0606086015260a0850190613f58565b930151151591015290565b906004821015613d6a5752565b61032081019081106001600160401b0382111761355257604052565b604081019081106001600160401b0382111761355257604052565b61010081019081106001600160401b0382111761355257604052565b60a081019081106001600160401b0382111761355257604052565b90601f801991011681019081106001600160401b0382111761355257604052565b906040516140cc8161404c565b602060018294805484520154910152565b90600182811c9216801561410b575b60208310146140f757565b634e487b7160e01b5f52602260045260245ffd5b91607f16916140ec565b9060405191825f825492614128846140dd565b8084529360018116908115614193575060011461414f575b5061414d9250038361409e565b565b90505f9291925260205f20905f915b81831061417757505090602061414d928201015f614140565b602091935080600191548385890101520191019091849261415e565b90506020925061414d94915060ff191682840152151560051b8201015f614140565b906040516141c281614067565b60e06004829460ff815481811615158652818160081c166020870152818160101c16604087015260181c16606085015260018101546080850152600281015460a0850152600381015460c08501520154910152565b9060405161422481614083565b608060ff600483958281541661423981613f4e565b85526001810154602086015260028101546001600160a01b0316604086015261426460038201614115565b60608601520154161515910152565b346129fb575f3660031901126129fb5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b9060018201809211612aa057565b91908201809211612aa057565b906142d281613f4e565b60ff80198354169116179055565b91908203918211612aa057565b91908260809103126129fb5781516001600160a01b0319811681036129fb57916020810151916060604083015192015190565b803560048110156129fb578261433591614023565b60208101356001600160601b0360a01b81168091036129fb5760c0918291602085015260408101356040850152606081013560608501526080810135608085015260a081013560a08501520135910152565b3580151581036129fb5790565b356001600160a01b03811681036129fb5790565b903590601e19813603018212156129fb57018035906001600160401b0382116129fb576020019181360383136129fb57565b8181106143e5575050565b5f81556001016143da565b908060209392818452848401375f828201840152601f01601f1916010190565b94929093614435926144439795875260208701526080604087015260808601916143f0565b9260608185039101526143f0565b90565b5f198114612aa05760010190565b6001600160401b0381116135525760051b60200190565b9061447582614454565b614482604051918261409e565b8281528092614493601f1991614454565b0190602036910137565b8051156144aa5760200190565b634e487b7160e01b5f52603260045260245ffd5b8051600110156144aa5760400190565b8051600210156144aa5760600190565b8051600310156144aa5760800190565b8051600410156144aa5760a00190565b8051600510156144aa5760c00190565b80518210156144aa5760209160051b010190565b60206040818301928281528451809452019201905f5b8181106145455750505090565b8251845260209384019390920191600101614538565b60408110156144aa5760051b60240190565b80156145905764e8d4a51000041061458157565b63eba5c29b60e01b5f5260045ffd5b634e487b7160e01b5f52601260045260245ffd5b6145ad81613f4e565b6145b682613f4e565b808214614663576145c681613f4e565b60028114801561467c575b8015614669575b614663576145e581613f4e565b8015614644576003906145f781613f4e565b1461460157505f90565b61460a81613f4e565b801590811561462f575b811561461e575090565b6001915061462b81613f4e565b1490565b905061463a81613f4e565b6002811490614614565b5061464e81613f4e565b6003811490811561462f57811561461e575090565b50505f90565b5061467381613f4e565b600181146145d8565b5061468681613f4e565b600481146145d1565b60025f541461469e5760025f55565b633ee5aeb560e01b5f5260045ffd5b50159081156146cd575b506146be57565b635e765b2560e11b5f5260045ffd5b9050155f6146b7565b602081519101519080158015614739575b8015614722575b614663575f5160206153dd5f395f51905f52808281930992800981808080848709620292f80960010893620292fc09081490565b505f5160206153dd5f395f51905f528210156146ee565b505f5160206153dd5f395f51905f528110156146e7565b600c820193916001600160401b03831161355257859061477086546140dd565b601f8111614886575b505f95601f85116001146147fc5790600d91857f77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f0985f916147f1575b508660011b905f198860031b1c19161790555b01556147e16040519384936040855260408501916143f0565b94602083015260ff1916930390a2565b90508701355f6147b5565b601f19851696815f5260205f20975f5b81811061486b575097600d939291877f77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f09a10614852575b5050600186811b0190556147c8565b8801355f19600389901b60f8161c191690555f80614843565b888301358a556001909901988a95506020928301920161480c565b6148b090875f5260205f20601f870160051c81019160208810611b9757601f0160051c01906143da565b5f614779565b60ff198116908115613b7f5763ffffffff808060035460401c16169160401c1603613b70575f52600160205260405f209060018060a01b03825460081c1615613b6157565b61490e60058201546006830154906142bb565b9063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016801983116149925761496e601a6149749301546001600160401b038160201c168581115f146149855763ffffffff90915b16906142bb565b926142bb565b80821015614980575090565b905090565b5063ffffffff8591614967565b5050505f1990565b90610200036149ea5760016149b2823560c01c615116565b14908115916149d1575b506149c357565b6254aabb60e81b5f5260045ffd5b6149e291506008013560c01c615116565b15155f6149bc565b630f61e7fd60e21b5f5260045ffd5b5f905f905b60088210614a0b57505090565b90916001908360020160031b83013560e01c8460051b60e0031b179201906149fe565b77ffffffffffffffff0000000000000000ffffffffffffffff8160081c9160081b917cff000000ff000000ff000000ff000000ff000000ff000000ff000000ff7dff000000ff000000ff000000ff000000ff000000ff000000ff000000ff007fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff0085167eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff84161760101c941691161760101b9179ffff000000000000ffff000000000000ffff000000000000ffff7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff847dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff84161760201c941691161760201b7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff82821673ffffffff000000000000000000000000ffffffff85161760401b93161760401c16178060801b9060801c1790565b90600360ff601384015416614ba881613f4e565b03614c2857601582015460405163650e5fcf60e01b6020820152602480820193909352918252614be791906001600160a01b0316612e4660448361409e565b9015918215614c1f575b8215614c15575b8215614c07575b505061392d57565b600a01541190505f80614bff565b4382119250614bf8565b81159250614bf1565b90601401540361392d57565b6001600160401b03811161355257601f01601f191660200190565b908210156144aa57614c669160051b8101906143a8565b9091565b91908110156144aa5760051b0190565b9291909180158015614d36575b614d2e575f925f5b828110614d025750818414614cf957614ca9906004615163565b92805b614cb7575050501490565b5f19019283906004821c600116614ce457614cde90614cd7838587614c6a565b35906151d3565b93614cac565b614cde90614cf3838587614c6a565b356151d3565b50505050505f90565b614d0d818486614c6a565b35614d1b575b600101614c8f565b935060018401808511612aa05793614d13565b505050505f90565b5060408111614c87565b6004811015613d6a57600303614d7e577f00000000000000000000000000000000000000000000000000000000000000005b6001600160a01b031690565b7f0000000000000000000000000000000000000000000000000000000000000000614d72565b519081151582036129fb57565b6018015460ff81166004811015613d6a5760031490811591614dd1575090565b604051635f0ddacd60e01b815260989190911b6001600160a01b031916600482015290506020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156129f0575f91614e38575090565b90506020813d602011614e64575b81614e536020938361409e565b810103126129fb5761444390614da4565b3d9150614e46565b9060ff600e82015460081c1691614e828361446b565b92601883015460ff8160781c169081614f82575b50505060048254928160ff85169460ff1916178155018351906001600160401b03821161355257600160401b8211613552578054828255808310614f66575b5060208501905f5260205f205f5b838110614f525750505050905f5160206153bd5f395f51905f5260409260ff1916928392815190614f1381613f4e565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e60405180614f4d339582614522565b0390a3565b600190602084519401938184015501614ee3565b614f7c90825f528360205f2091820191016143da565b5f614ed5565b5f6001600160a01b03614f9760ff8416614d40565b1692608460198801546040519586938492630222162f60e21b84526001600160601b0360a01b8860981b166004850152602484015261ffff8760681c16604484015260648301525afa9182156129f0575f905f9361505d575b501561504e575f919060801c61ffff16825b848110615010575050614e96565b600182821c1615615024575b600101615002565b92615046816150356001938661450e565b51615040878c61450e565b52614446565b93905061501c565b630e0d4dc560e41b5f5260045ffd5b9250503d805f843e61506f818461409e565b8201916040818403126129fb5761508581614da4565b906020810151906001600160401b0382116129fb57019280601f850112156129fb5783516150b281614454565b946150c0604051968761409e565b81865260208087019260051b8201019283116129fb57602001905b8282106150eb575050505f614ff0565b81518152602091820191016150db565b6020915f91838251920190620186a0fa601f3d1116905f5190565b66ff00ff00ff00ff67ff00ff00ff00ff008260081b169160081c161765ffff0000ffff67ffff0000ffff00008260101b169160101c161767ffffffff000000008160201b169060201c1790565b60209161518b5f926151816001600160401b038060c01b9216615116565b60c01b1691614a2e565b604051908482019283526028820152600160f81b6048820152602981526151b360498261409e565b604051918291518091835e8101838152039060025afa156129f0575f5190565b5f90602092604051908482019283526040820152604081526151b360608261409e565b9081518015614d2e57600181146153ab576040841461539c575f915f5b828110615373575061522d61522884846142e0565b61446b565b9161523b61522885836142e0565b9161524e6152488661446b565b9561446b565b955f9081805b888a8c888710615293579550505050505061527c9493925061527691506142ad565b916151f6565b9161527661528a93946142ad565b614443916151d3565b906152d860016152cb8a95946152b96152ac8c8c61450e565b516001600160401b031690565b906001600160401b03809216901c1690565b166001600160401b031690565b615331575050506153286001916153236152f56152ac888861450e565b6152ff888a61450e565b5161530a848d61450e565b52615315838d61450e565b906001600160401b03169052565b614446565b935b0192615254565b61536d92615315868094615323946153678c9a6153608c60019c9f8f61535a916152ac9161450e565b9861450e565b519261450e565b5261450e565b9161532a565b9261539560019161538f836152cb8a6152b96152ac8b8d61450e565b906142bb565b9301615213565b63394fd24160e21b5f5260045ffd5b5090506153b8915061449d565b519056fe56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212205fae99d926456dc3d3d096770da58d62f22b1d2785a7b00f3fc262de1b73a00364736f6c634300081c003360e0806040523461011157602061002e60049261106880380380916100248285610115565b833981019061014c565b336080526001600160a01b031660a081905260405163ebe86c1360e01b815292839182905afa908115610106575f916100d7575b506001600160a01b031660c052604051610efc908161016c82396080518181816101530152818161037701528181610756015281816107cd0152610bce015260a051818181610186015281816108b701528181610a2f0152610c43015260c05181818160c30152818161047001526107fd0152f35b6100f9915060203d6020116100ff575b6100f18183610115565b81019061014c565b5f610062565b503d6100e7565b6040513d5f823e3d90fd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761013857604052565b634e487b7160e01b5f52604160045260245ffd5b9081602091031261011157516001600160a01b0381168103610111579056fe60806040526004361015610011575f80fd5b5f5f3560e01c8063088858bc146108e6578063481c6a75146108a257806364fe50ac146107b1578063702574b3146107855780637b10399914610740578063883a429a1461034e578063ae977ede146100f2578063ebe86c13146100ad5763f08c9b3c1461007d575f80fd5b346100aa57806003193601126100aa576020610097610c34565b6040516001600160a01b03199091168152f35b80fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760603660031901126100aa5761010c610977565b9060243591604435916001600160401b0383116100aa57366023840112156100aa578260040135936001600160401b03851161034a573660248660071b8601011161034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b579093927f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316928592916001600160a01b031990911690835b8688101561032c578760071b82019060845f516020610ea75f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d2899602485013509925f516020610ea75f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d289960648301350960405194633d98dab360e11b865287600487015288602487015260448601526044820135606486015282850152013560a483015260208260c481898b5af19182156103215786926102e2575b508861029c5750600190975b01966101c1565b979061ffff8916908282018092116102ce5761ffff16036102bf57600190610295565b6357149e2560e01b8552600485fd5b634e487b7160e01b87526011600452602487fd5b9091506020813d8211610319575b816102fd6020938361099f565b810103126103155761030e906109eb565b905f610289565b8580fd5b3d91506102f0565b6040513d88823e3d90fd5b60209061ffff60405191168152f35b633217675b60e21b8252600482fd5b5080fd5b50346100aa576101003660031901126100aa5761036961098e565b60e036602319011261034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b576103ae90610bab565b602435600481101561073c5760021490811561072e576044356001600160a01b03198116810361072a57915b6040908151906103ea838361099f565b6001825260208201601f198401368237825115610716573090521561070e5784935b82519461010086018681106001600160401b038211176106fa57845260028110156106e6578552602085019486865283810192835260608101601081526080820188815260a0830189815260c08401918a835260e08501938b855260018060a01b037f00000000000000000000000000000000000000000000000000000000000000001697883b156106e257895197631bb64f5960e21b89526001600160601b0360a01b169b8c60048a01528b60248a015261010060448a01526102048901975160028110156106ce5791899693918f989593610104899b989b01525115156101248801525194610100610144880152855180915260206102248801960190885b81811061069d57505050916001600160401b03869798818096959461ffff829651166101648b0152511661018489015251166101a487015251166101c485015251166101e48301526064356064830152608435608483015260a43560a483015260c43560c483015260e43560e4830152038183865af180156106935761067a575b5093816044958151968780926303e95d1360e21b82528860048301528760248301525afa92831561066e57608095829461060b575b505f516020610ea75f395f51905f52917f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e0097542918451968752602087015209908301526060820152f35b7f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e00975429194505f516020610ea75f395f51905f52925061065e90843d8611610667575b610656818361099f565b810190610c1e565b949150916105c3565b503d61064c565b509051903d90823e3d90fd5b61068586809261099f565b61068f575f61058e565b8480fd5b83513d88823e3d90fd5b92949750929497509794602080600192838060a01b038c511681520199019101908e9794928a97949299969961050d565b634e487b7160e01b8f52602160045260248ffd5b8c80fd5b634e487b7160e01b87526021600452602487fd5b634e487b7160e01b88526041600452602488fd5b60019361040c565b634e487b7160e01b87526032600452602487fd5b8380fd5b610736610c34565b916103da565b8280fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760203660031901126100aa5760206107a96107a461098e565b610bab565b604051908152f35b503461088f57606036600319011261088f576107cb610977565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610893577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561088f575f9160648392604051948593849263a59b7a4d60e01b84526001600160601b0360a01b166004840152602435602484015260443560448401525af1801561088457610876575080f35b61088291505f9061099f565b005b6040513d5f823e3d90fd5b5f80fd5b633217675b60e21b5f5260045ffd5b3461088f575f36600319011261088f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461088f57608036600319011261088f576108ff610977565b60443561ffff8116810361088f576064359061ffff8216820361088f5761092992602435906109fa565b906040519182916040830190151583526040602084015281518091526020606084019201905f5b81811061095e575050500390f35b8251845285945060209384019390920191600101610950565b600435906001600160a01b03198216820361088f57565b6004359060ff198216820361088f57565b90601f801991011681019081106001600160401b038211176109c057604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b0381116109c05760051b60200190565b519061ffff8216820361088f57565b9261ffff1693610a09856109d4565b610a16604051918261099f565b858152601f19610a25876109d4565b01366020830137937f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925f5b61ffff811688811015610b9e578061ffff88160161ffff8111610b8a57604051639bbada6760e01b81526001600160a01b0319861660048201526024810185905261ffff9190911660448201526060816064818a5afa908115610884575f91610b14575b50602081015115610b0657604001519088511115610af257600582901b621fffe01688016020015260010161ffff16610a5a565b634e487b7160e01b5f52603260045260245ffd5b505f98509395505050505050565b90506060813d8211610b82575b81610b2e6060938361099f565b8101031261088f5760405190606082018281106001600160401b038211176109c057604052610b5c816109eb565b8252602081015190811515820361088f576040916020840152015160408201525f610abe565b3d9150610b21565b634e487b7160e01b5f52601160045260245ffd5b5050505093505050600191565b5f516020610ea75f395f51905f5290604051602081019146835260018060a01b037f000000000000000000000000000000000000000000000000000000000000000016604083015260ff1916606082015260608152610c0b60808261099f565b519020068015610c185790565b50600190565b919082604091031261088f576020825192015190565b60405163a4adcd7f60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316602082600481845afa918215610884575f92610e62575b506040516323488be560e01b8152602081600481855afa8015610884575f90610e1c575b63ffffffff60401b915060401b165f5b6001600160401b0381166008811080610e0a575b15610dfb576001600160401b038516036001600160401b038111610b8a576001600160401b036001600160601b0360a01b9116831760a01b1660405163d397925360e01b8152816004820152602081602481885afa8015610884575f90610dbf575b60ff9150166010811015610daf57604051906356cbb5f360e01b82528260048301526024820152604081604481885afa9081610d92575b50610d8a57506001600160401b03905b166001600160401b038114610b8a57600101610cb6565b935050505090565b610da99060403d811161066757610656818361099f565b50610d63565b50506001600160401b0390610d73565b506020813d8211610df3575b81610dd86020938361099f565b8101031261088f575160ff8116810361088f5760ff90610d2c565b3d9150610dcb565b63081ea97160e31b5f5260045ffd5b50806001600160401b03861611610cca565b506020813d602011610e5a575b81610e366020938361099f565b8101031261088f575163ffffffff8116810361088f5763ffffffff60401b90610ca6565b3d9150610e29565b9091506020813d602011610e9e575b81610e7e6020938361099f565b8101031261088f57516001600160401b038116810361088f57905f610c82565b3d9150610e7156fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220211a5b4b4c4129b6c3d0e7e9959a307a0521674b2edd8ecaccfc2f92ca9c394464736f6c634300081c003360c0346100a257601f610a4438819003918201601f19168301916001600160401b038311848410176100a6578084926020946040528339810103126100a257516001600160a01b038116908190036100a2573360805260a05260405161098990816100bb823960805181818160b90152818161034101526104ea015260a051818181610168015281816104050152818161058101528181610614015261075a0152f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c908163088858bc14610698575080632680c7d814610643578063481c6a75146105ff5780635f0ddacd1461054257806364fe50ac146105195780637b103999146104d5578063ae977ede146102e25763ecd8e20514610074575f80fd5b34610266576101203660031901126102665760043560ff198116809103610266576024356001600160a01b03811691908290036102665760e0366043190112610266577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036102d357604435600481101561026657600219016102c457606435906001600160601b0360a01b821680920361026657811580156102b9575b80156102ae575b80156102a3575b8015610298575b801561028c575b61027d576040516358d6692360e01b8152600481018390526024810182905260448101939093526060836064815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610272575f5f915f93610228575b6080955060016040516101b581610908565b8681525f602080830182815260408085019687528784529183905291209151825491516cffffffffffffffffffffffffff1990921660a09190911c1760609190911b60ff60601b1617815501905160081c60ff60f81b825416179055604051938452602084015260408301526060820152f35b925050506060833d60601161026a575b8161024560609383610924565b810103126102665782608093519060406020820151910151919091926101a3565b5f80fd5b3d9150610238565b6040513d5f823e3d90fd5b63e4291a1960e01b5f5260045ffd5b50610104351515610138565b5060e4351515610131565b5060c435151561012a565b5060a4351515610123565b50608435151561011c565b6365b75c3960e01b5f5260045ffd5b633217675b60e21b5f5260045ffd5b34610266576060366003190112610266576102fb6108f1565b6024359060443567ffffffffffffffff8111610266573660238201121561026657806004013567ffffffffffffffff8111610266573660248260071b84010111610266577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036102d357835f525f60205260405f2060ff19600182015460081b1693841580156104b9575b6104aa57815460ff60601b1916606084901b60ff60601b16179091559291908060405194859463548efb7160e01b865260648601916001600160601b0360a01b166004870152602486015260606044860152526024608484019201905f5b81811061048e57506020939283900391508290505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610272575f9161045c575b500361044d5760206040515f8152f35b63ba343de560e01b5f5260045ffd5b90506020813d602011610486575b8161047760209383610924565b8101031261026657518261043d565b3d915061046a565b91935091608080828187600195370194019101918493926103ee565b636d08029760e01b5f5260045ffd5b50815460a01b6001600160a01b03199081169082161415610390565b34610266575f366003190112610266576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610266576060366003190112610266576105326108f1565b50633254484760e01b5f5260045ffd5b346102665760203660031901126102665761055b6108f1565b604051635f0ddacd60e01b81526001600160a01b031990911660048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610272575f906105c5575b6020906040519015158152f35b506020813d6020116105f7575b816105df60209383610924565b81010312610266576105f2602091610946565b6105b8565b3d91506105d2565b34610266575f366003190112610266576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610266576020366003190112610266576004355f525f602052606060405f206001815491015460081b60ff604051926001600160601b0360a01b8160a01b168452841c16602083015260ff19166040820152f35b34610266576080366003190112610266576106b16108f1565b6024359160443561ffff8116809103610266576064359261ffff841680940361026657845f525f60205260405f20926106e981610908565b83549060ff6001600160601b0360a01b8360a01b169283835260601c169060406020820196838852600160ff1991015460081b16910152159182156108dc575b50506104aa5715908115916108cd575b506108be57604051635584afe560e11b815260048101929092525f826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610272575f925f926107fc575b508290816107f0575b5061044d57906040519182916040830190151583526040602084015281518091526020606084019201905f5b8181106107d7575050500390f35b82518452859450602093840193909201916001016107c9565b9050815114158361079d565b925090503d805f843e61080f8184610924565b8201916040818403126102665761082581610946565b9060208101519067ffffffffffffffff821161026657019280601f850112156102665783519367ffffffffffffffff85116108aa578460051b90604051956108706020840188610924565b865260208087019282010192831161026657602001905b82821061089a5750505080929190610794565b8151815260209182019101610887565b634e487b7160e01b5f52604160045260245ffd5b6347269e1960e11b5f5260045ffd5b60ff9150511681141583610739565b6001600160a01b031916141590508580610729565b600435906001600160a01b03198216820361026657565b6060810190811067ffffffffffffffff8211176108aa57604052565b90601f8019910116810190811067ffffffffffffffff8211176108aa57604052565b519081151582036102665756fea26469706673582212207bcf340fb397ff11359e7faf726169f40ad1ba49fe64c983f09185991b89c9d364736f6c634300081c0033",
}

// ProcessRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProcessRegistryMetaData.ABI instead.
var ProcessRegistryABI = ProcessRegistryMetaData.ABI

// ProcessRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProcessRegistryMetaData.Bin instead.
var ProcessRegistryBin = ProcessRegistryMetaData.Bin

// DeployProcessRegistry deploys a new Ethereum contract, binding an instance of ProcessRegistry to it.
func DeployProcessRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID uint32, _ziskVerifier common.Address, _batchProgramVK [32]byte, _resultsProgramVK [32]byte, _rootCVadcopFinal [32]byte, _ballotVKHash [32]byte, _dkgManager common.Address, _councilManager common.Address, _defaultGrace uint32, _graceFloor uint32, _graceCeil uint32, _graceMaxTotal uint32, _noticeMin uint32) (common.Address, *types.Transaction, *ProcessRegistry, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProcessRegistryBin), backend, _chainID, _ziskVerifier, _batchProgramVK, _resultsProgramVK, _rootCVadcopFinal, _ballotVKHash, _dkgManager, _councilManager, _defaultGrace, _graceFloor, _graceCeil, _graceMaxTotal, _noticeMin)
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
