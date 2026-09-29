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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_ziskVerifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_batchProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_resultsProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rootCVadcopFinal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_ballotVKHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_dkgManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_defaultGrace\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceFloor\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceCeil\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_graceMaxTotal\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_noticeMin\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlobCountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CircuitFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DKGDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyTransition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GraceOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccumulator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobCommitmentLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidBlobOpening\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobsDigest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDKGParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncryptionKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGrace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInclusionProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKZGProofLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOccupiedBefore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxPossibleResultCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"MissingBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoBlobs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsAlreadyRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtMaxLevelsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"}],\"name\":\"ProcessGraceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"ProcessMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nBlobs\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"firstIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"name\":\"ResultsDecryptionRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"aidFor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotVKHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGrace\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dkgAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"finalizeResultsFromDKG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"}],\"name\":\"genesisRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastVoteAt\",\"type\":\"uint64\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessGraceEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceCeil\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceFloor\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"graceMaxTotal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint256\",\"name\":\"orgPKx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orgPKy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popZ\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.DKGParams\",\"name\":\"dkg\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noticeMin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"lastVoteAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256[64]\",\"name\":\"accumulator\",\"type\":\"uint256[64]\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"requestResultsDecryption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultsProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"sk\",\"type\":\"uint256\"}],\"name\":\"revealProcessKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootCVadcopFinal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint32\",\"name\":\"grace\",\"type\":\"uint32\"}],\"name\":\"setProcessGrace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"setProcessMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"commitments\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"kzgProofs\",\"type\":\"bytes[]\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ziskVerifier\",\"outputs\":[{\"internalType\":\"contractIZiskVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101e0806040523461038a576101808161659f8038038091610021828561038e565b83398101031261038a57610034816103b1565b610040602083016103c2565b916040810151606082015160808301519160a08401519361006360c082016103c2565b9661007060e083016103b1565b9161007e61010082016103b1565b61008b61012083016103b1565b906100a661016061009f61014086016103b1565b94016103b1565b60015f556001600160a01b039094169485158015610382575b801561037a575b8015610372575b801561036a575b61035b5763ffffffff8216801590811561034b575b508015610336575b8015610321575b8015610313575b610304576101405261016052610180526101a0526101c05260805260a05260c05260e0526101005260035467ffffffff000000006bffffffff0000000000000000604051602081019063ffffffff60e01b8660e01b1682523060601b60248201526018815261016f60388261038e565b51902060401b169260201b1690640100000000600160601b031916171760035560018060a01b031680155f146102a457505f5b6101205260405161516090816103d782396080518181816110ae015281816114160152612596015260a0518181816126020152613e06015260c0518181816114860152614165015260e05181818161146501528181611c5901526125e1015261010051818181610bfb01528181612b6a0152613114015261012051818181610595015281816108df015281816116df015281816117b201528181613598015281816136410152614ce2015261014051818181610996015261332d015261016051818181610fa20152611c9a0152610180518181816110310152612b2f01526101a051818181612203015261475701526101c05181818161092c0152610e7a0152f35b604051906110688083016001600160401b038111848210176102f0576020928492615537843981520301905ff080156102e5576001600160a01b03166101a2565b6040513d5f823e3d90fd5b634e487b7160e01b5f52604160045260245ffd5b63795ee5af60e01b5f5260045ffd5b5063ffffffff8516156100ff565b5063ffffffff841663ffffffff8416116100f8565b5063ffffffff831663ffffffff8216116100f1565b905063ffffffff8216105f6100e9565b6306f9b90760e01b5f5260045ffd5b5089156100d4565b5088156100cd565b5087156100c6565b5086156100bf565b5f80fd5b601f909101601f19168101906001600160401b038211908210176102f057604052565b519063ffffffff8216820361038a57565b51906001600160a01b038216820361038a5756fe6101406040526004361015610012575f80fd5b5f60a0525f3560e01c8063026cdee814613c7257806304ed00fa14613a69578063082b642e1461389e57806308c0fdd314612b8d5780630e2ebcf714612b535780631542bbe214612b135780631fdf34491461231d5780633ea4ee41146122e757806346c15da8146122275780634c0acc561461106d578063549d5995146121e557806359d821c414611cbe5780635ff5f98114611c7c5780636211533814611c4057806362fa11fc14611bf757806368141f2c14611b785780636c7aff7f14611861578063702574b31461179757806372c628ef1461177a5780637341770514611696578063766422e01461132d578063784df74a146110dd5780637f64b72f14611097578063848df54014611072578063946544bf1461106d5780639a03778914610f1c5780639b46499414610d46578063aa240221146101b5578063adc879e914610d1f578063bf74291e146109ba578063c8f0582f14610978578063cddf08bc14610950578063d4138a201461090e578063e16d5b7c146108c8578063e965eead146102c5578063f1431097146101ba5763f9aa4499146101b5575f80fd5b61414e565b346102bf5760203660031901126102bf576101d3613d81565b6101db6144d0565b6101e4816146f7565b601881015460ff811660038110156102a7571561029457815460ff1661020981613e29565b60028114908115610280575b5061026d576102238261473c565b421061025a5760901c60ff16156102475761023d91614b81565b60a0516001815580f35b630e0d4dc560e41b60a05152600460a051fd5b63611f72f360e11b60a05152600460a051fd5b6307a92f1960e51b60a05152600460a051fd5b6004915061028d81613e29565b1484610215565b6365b75c3960e01b60a05152600460a051fd5b634e487b7160e01b60a051526021600452602460a051fd5b60a05180fd5b346102bf576108403660031901126102bf576102df613d81565b36610824116102bf57610824356001600160401b0381116102bf57610308903690600401613dbf565b6103106144d0565b610319836146f7565b916018830190815460ff811660038110156102a757156102945760ff8160901c166108b557845460ff81169290919061035184613e29565b6002841480156108a2575b61026d5761036984613e29565b6001841415958680610887575b610874576103838861473c565b421061025a5760a0515b6040811061084e575060206040518181016108006024823761080082526103b661082083613f79565b60405191518091835e81019060a05182528060a05192039060025afa15610617576103e99160a0515160038a0154614abb565b1561083b5760ff60901b1916600160901b178084559361040883613e29565b6107fc575b505060ff600e84015460081c169161042483614295565b926104326040519485613f79565b808452601f1961044182614295565b0160a0515b8181106107da57505060a05191829182905b80821061066957505060a051821595909190861561052a575b50508354607883901b60ff60781b1664ffffffffff60681b19909116606883901b61ffff60681b161717608093841b61ffff60801b16179384905560198601546040805160989690961b6001600160a01b0319168652602086019190915261ffff919091169084015260ff16606083015260ff198516917fdca6075f07367349a836825d3ee7c35c204d2e727ae81a5b7a48aca95b9c270b9190a261051a5760a0516001815580f35b61052391614b81565b808061023d565b838152601988015460405163574bbf6f60e11b815260989390931b6001600160a01b031916600484015260248301526060604483015280516064830181905260a0519293508392608484019260200191905b81811061062457505060a05160209392839003915082907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af19081156106175760a051916105d9575b508261ffff610471565b90506020813d60201161060f575b816105f460209383613f79565b810103126102bf575161ffff811681036102bf5760ff6105cf565b3d91506105e7565b6040513d60a051823e3d90fd5b9180945092909251819060a051915b60048310610653575050506020608060019201940191019184939261057c565b6020806001928451815201920192019190610633565b9092600284901b6001600160fe1b038516850361074b576106898161439c565b35916001820180831161074b5761069f9061439c565b3560a051600284019182851161074b576106b88361439c565b3515806107b7575b8615806107ad575b61078e5761077b5760405195608087018781106001600160401b03821117610763576040528652602086015261074b576107019061439c565b3560408401526003820191821061074b576001926107216107429361439c565b356060820152610731828b61434f565b5261073c818a61434f565b50614287565b935b0190610458565b634e487b7160e01b60a051526011600452602460a051fd5b634e487b7160e01b60a051526041600452602460a051fd5b632a23591560e21b60a05152600460a051fd5b9450505050949591501561077b5760019061ffff82871b161794610744565b50600182146106c8565b5060a05191506003850180861161074b576107d360019161439c565b35146106c0565b6040516020919060806107ed8183613f79565b36823782828901015201610446565b60ff191660011784556040519061081281613e29565b81526001602082015260ff198516905f5160206150eb5f395f51905f5290604090a2848061040d565b6303cd656760e61b60a05152600460a051fd5b5f51602061510b5f395f51905f526108658261439c565b35101561077b5760010161038d565b63e843c5eb60e01b60a05152600460a051fd5b5061089b600589015460068a015490614196565b4210610376565b506108ac84613e29565b6004841461035c565b636c8ae94b60e11b60a05152600460a051fd5b346102bf5760a0513660031901126102bf576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102bf5760a0513660031901126102bf57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bf5760a0513660031901126102bf57602063ffffffff60035460401c16604051908152f35b346102bf5760a0513660031901126102bf57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bf576101803660031901126102bf576109d4613d81565b6101003660231901126102bf576040366101231901126102bf576101643560058110156102bf5760405190610a0882613f27565b61012435825261014435602083019081526040519290610a2960e085613f79565b6006845260c0948536602087013760405195610a4660e088613f79565b6006875236602088013760081c610a5c866142de565b5260a051610a69856142de565b5260643560ff811692908381036102bf576044359360ff8516908186036102bf578110610d0c5760a43565ffffffffffff8111610cf95760c4359165ffffffffffff8311610ce65760e43593677fffffffffffffff8511610cd3576101043597677fffffffffffffff8911610cc0575060243580151581036102bf5715610cb7576001905b6084359260ff841684036102bf5761ff0062ff00006301fe000060209c60b81b9960791b9860491b9760191b9660111b169460101b169260081b1617171717171717610b39876142ff565b526002610b45866142ff565b525190519060405191838301918252604083015260408252610b68606083613f79565b60405191518091835e81019060a05182528060a05192039060025afa156106175760a05151610b968461430f565b526003610ba28361430f565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae278610bcd8461431f565b526004610bd98361431f565b52610be381613e29565b610bec8361432f565b526006610bf88261432f565b527f0000000000000000000000000000000000000000000000000000000000000000610c238361433f565b526007610c2f8261433f565b528051825103610ca457610c4381516142ac565b60a0515b8251811015610c8c5780610c7b6001600160401b03610c686001948761434f565b5116610c74838861434f565b5190614e91565b610c85828561434f565b5201610c47565b6020610c9c848460a05191614f24565b604051908152f35b63d088249360e01b60a05152600460a051fd5b60a05190610aee565b63dd6f54df60e01b60a05152600460a051fd5b63271fb80560e01b60a05152600460a051fd5b63871a7fa360e01b60a05152600460a051fd5b63481eb79f60e01b60a05152600460a051fd5b632cbdc23160e01b60a05152600460a051fd5b346102bf5760a0513660031901126102bf57602063ffffffff600354821c16604051908152f35b346102bf5760403660031901126102bf57610d5f613d81565b60243560ff198216918215610f095763ffffffff808060035460401c16169160401c1603610ef65760a08051839052600160205251604090208054600881901c6001600160a01b03168015610ee3573303610ed15760ff16610dc081613e29565b8015159081610ebc575b5061026d5760066005820154910190610de4825482614196565b90428211156108745783610df791614196565b8315918215610eb1575b8215610ea7575b8215610e5c575b5050610e4957817f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba9260209255604051908152a260a05180f35b637616640160e01b60a05152600460a051fd5b8110915081610e6e575b508480610e0f565b9050610ea063ffffffff7f00000000000000000000000000000000000000000000000000000000000000001642614196565b1184610e66565b8181149250610e08565b428211159250610e01565b60039150610ec981613e29565b141584610dca565b6282b42960e81b60a05152600460a051fd5b634d36eb6960e01b60a05152600460a051fd5b632299770d60e11b60a05152600460a051fd5b63cbf4a64560e01b60a05152600460a051fd5b346102bf5760403660031901126102bf57610f35613d81565b6024359063ffffffff82168092036102bf57610f50816146f7565b805433600882901c6001600160a01b031603610ed15760ff16610f7281613e29565b8015159081611058575b5061026d57610f946005820154600683015490614196565b4210156108745763ffffffff7f00000000000000000000000000000000000000000000000000000000000000001683108015611029575b61101657601a01805463ffffffff19168317905560405191825260ff1916907fdf161c6af27d090672f982bb8002da0e7f2a530ffbd779e02ab9eba9397e51f190602090a260a05180f35b63795ee5af60e01b60a05152600460a051fd5b5063ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168311610fcb565b6003915061106581613e29565b141584610f7c565b613def565b346102bf5760a0513660031901126102bf57602063ffffffff60035416604051908152f35b346102bf5760a0513660031901126102bf576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102bf5760203660031901126102bf5760ff196110f9613d81565b1660a05152600160205260a0516040902080546001820161111990613f9a565b60e0526003820154610120526005820154600683015460078401549060088501546009860154600a870154600b88015491600c890161115790613ff0565b93600d8a015496600e8b0161116b90614090565b9661117860138d016140f2565b9960188d015460c05260198d01549c601a01549b6040516101005260ff81166111a090613e29565b60ff81166101005152600160a01b600190039060081c16610100516020015260e05151610100516040015260e0516020015161010051606001526101205161010051608001526101005160a001526101005160c001526101005160e0015261010051610100015261010051610120015261010051610140015261010051610160015261010051610180016104009052610100516104000161124091613e33565b91610100516101a00152610100516101c00161125b91613e57565b610100518103610100516102c0015261127391613ead565b91610100516102e00160c05160ff169061128c91613efe565b6001600160601b0360a01b60c05160981b1661010051610300015260c05160681c61ffff1661010051610320015260c05160781c60ff1661010051610340015260c05160801c61ffff1661010051610360015260c05160901c60ff161515610100516103800152610100516103a0015263ffffffff8116610100516103c0015260201c6001600160401b0316610100516103e0015261010051900361010051f35b346102bf5760603660031901126102bf57611346613d81565b6024356001600160401b0381116102bf57611365903690600401613d92565b90916044356001600160401b0381116102bf57611386903690600401613d92565b9261138f6144d0565b611398836146f7565b9160ff60188401541660038110156102a75761029457825460ff8116959092906113c187613e29565b600287148015611683575b61026d576113d987613e29565b600187141580611668575b610874576113f18561473c565b421061025a5761140182896147db565b61140a8861483a565b600386015403611655577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156102bf576114ae9389936040519586948593849363be98686160e01b855260a051987f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801614251565b03915afa80156106175761163c575b5060ff600e83015460081c16946114d3866142ac565b9560a0515b8181106115d35750505060ff191660049081178255845191016001600160401b03821161076357600160401b82116107635780548282558083106115b3575b50602085019060a05152602060a0512060a0515b83811061159f57866040875f5160206150eb5f395f51905f528860ff191692839281519061155881613e29565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e60405180611592339582614363565b0390a360a0516001815580f35b60019060208451940193818401550161152b565b6115cd908260a0515283602060a05120918201910161421b565b85611517565b600181901b906001600160ff1b038116810361074b5781600a019182600a1161074b57600b61160a8460031b87013560c01c614e44565b910192831061074b5761162760019360031b86013560c01c614e44565b60201b17611635828b61434f565b52016114d8565b60a05161164891613f79565b60a0516102bf57856114bd565b630b6fac0360e41b60a05152600460a051fd5b5061167c6005860154600687015490614196565b42106113e4565b5061168d87613e29565b600487146113cc565b346102bf5760403660031901126102bf576116bf6116b2613d81565b6116ba6144d0565b6146f7565b60188101549060ff821660038110156102a75760020361029457601901547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156102bf576040519263193f942b60e21b84526001600160601b0360a01b9060981b166004840152602483015260243560448301528160648160a0519360a051905af18015610617576117615760a0516001815580f35b60a05161176d91613f79565b60a0516102bf578061023d565b346102bf5760a0513660031901126102bf57602060405160048152f35b346102bf5760203660031901126102bf576117b0613d81565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690811561184f5760209060246040518094819363702574b360e01b835260ff191660048301525afa80156106175760a0519061181c575b602090604051908152f35b506020813d602011611847575b8161183660209383613f79565b810103126102bf5760209051611811565b3d9150611829565b6203eb8f60e01b60a05152600460a051fd5b346102bf5760403660031901126102bf5761187a613d81565b6024356001600160401b0381116102bf578060040160a060031983360301126102bf576118a6836146f7565b8054919033600884901c6001600160a01b031603610ed157601381015460029060ff166118d281613e29565b03611b6557813560058110156102bf576002906118ee81613e29565b03611b52576001600160a01b03611907604486016141d5565b16611b3f576024840135938415611b2c576064019261192684846141e9565b905015611b195760ff1661193981613e29565b8015159081611b04575b5061026d5761195b6005820154600683015490614196565b42101561087457836014820155601661197484846141e9565b91909201916001600160401b038211610763576119918354613fb8565b601f8111611ac5575b5060a05190601f8311600114611a2f579282611a1b9896937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed29896936119fb9660a05192611a24575b50508160011b915f199060031b1c19161790556141e9565b949060405193849384526040602085015260ff1916956040840191614231565b0390a260a05180f35b013590508a806119e3565b601f198316918460a05152602060a051209260a0515b818110611aad5750937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed29896936119fb969360019383611a1b9d9b9810611a94575b505050811b0190556141e9565b01355f19600384901b60f8161c191690558a8080611a87565b91936020600181928787013581550195019201611a45565b611af4908460a05152602060a05120601f850160051c81019160208610611afa575b601f0160051c019061421b565b8761199a565b9091508190611ae7565b60039150611b1181613e29565b141586611943565b630f8b932160e11b60a05152600460a051fd5b635e32eadd60e01b60a05152600460a051fd5b63562e597160e11b60a05152600460a051fd5b63f37f7b5d60e01b60a05152600460a051fd5b63050b77c760e21b60a05152600460a051fd5b346102bf5760203660031901126102bf576004356001600160a01b038116908181036102bf5760035460a0805193909352600260209081529251604090205466ffffffffffffff1660589290921b600160581b600160f81b0316600891821c63ffffffff60381b161791909117901b60ff19166040519060ff19168152f35b346102bf5760203660031901126102bf576004356001600160a01b038116908190036102bf5760a05152600260205260206001600160401b03604060a051205416604051908152f35b346102bf5760a0513660031901126102bf5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346102bf5760a0513660031901126102bf57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bf5760203660031901126102bf57611cd7613d81565b604051611ce381613f0b565b60a051815260a0516020820152604051611cfc81613f27565b60a051815260a0516020820152604082015260a05160608201526060608082015260a05160a082015260a05160c082015260a05160e082015260a05161010082015260a05161012082015260a05161014082015260a051610160820152606061018082015260a0516101a0820152604051611d7681613f42565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260a05160c082015260a05160e08201526101c0820152604051611dc581613f5e565b60a051815260a051602082015260a051604082015260608082015260a05160808201526101e082015260a05161020082015260a05161022082015260a05161024082015260a05161026082015260a05161028082015260a0516102a082015260a0516102c082015260a0516102e082015261030060a05191015260ff191660a051526001602052604060a0512060405190611e5f82613f0b565b805460ff8116611e6e81613e29565b835260081c6001600160a01b03166020830152611e8d60018201613f9a565b604083015260038101546060830152600481016040518082602082945493848152019060a05152602060a051209260a0515b8181106121cc575050611ed492500382613f79565b6080830152600581015460a0830152600681015460c0830152600781015460e083015260088101546101008301526009810154610120830152600a810154610140830152600b810154610160830152611f2f600c8201613ff0565b610180830152600d8101546101a0830152611f4c600e8201614090565b6101c0830152611f5e601382016140f2565b6101e08301526018810154600360ff821610156102a7576001600160401b039160ff8281601a94166102008701526001600160601b0360a01b8160981b1661022087015261ffff8160681c16610240870152818160781c1661026087015261ffff8160801c1661028087015260901c1615156102a085015260198101546102c0850152015463ffffffff81166102e084015260201c1661030082015260405160208152610440810191805161201281613e29565b602083015260018060a01b036020820151166040830152602060408201518051606085015201516080830152606081015160a083015260808101519261042060c084015283518091526020610460840194019060a0515b8181106121b6575050506001600160401b036103006121246120ee859660a086015160e088015260c086015161010088015260e08601516101208801526101008601516101408801526101208601516101608801526101408601516101808801526101608601516101a0880152610180860151601f19888303016101c0890152613e33565b6101a08501516101e087015261210e6101c0860151610200880190613e57565b6101e0850151868203601f190184880152613ead565b92612139610200820151610320870190613efe565b6102208101516001600160a01b03191661034086015261024081015161ffff90811661036087015261026082015160ff16610380870152610280820151166103a08601526102a081015115156103c08601526102c08101516103e08601526102e081015163ffffffff166104008601520151166104208301520390f35b8251865260209586019590920191600101612069565b8454835260019485019486945060209093019201611ebf565b346102bf5760a0513660031901126102bf57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bf5760603660031901126102bf57612240613d81565b6024356001600160401b0381116102bf5761225f903690600401613d92565b6044359161226c846146f7565b805433600882901c6001600160a01b031603610ed15760ff906122908686866144ee565b1661229a81613e29565b80151590816122d2575b5061026d576122bc6005820154600683015490614196565b421015610874576122cc94614591565b60a05180f35b600391506122df81613e29565b1415866122a4565b346102bf5760203660031901126102bf5760ff19612303613d81565b1660a0515260016020526020610c9c604060a0512061473c565b3461295e5760c036600319011261295e57612336613d81565b6080526024356001600160401b03811161295e57612358903690600401613d92565b906044356001600160401b03811161295e57612378903690600401613d92565b90916064356001600160401b03811161295e57612399903690600401613dbf565b9190946084356001600160401b03811161295e576123bb903690600401613dbf565b9060a4356001600160401b03811161295e576123db903690600401613dbf565b91906123e56144d0565b6123f06080516146f7565b9660ff8854166123ff81613e29565b8015159081612afd575b81612ac6575b50612ab75760058801544210612aa8576124288861473c565b421015612aa85761243986886147db565b6124428761483a565b9a60038901548c03612a99575f5f5b60088110612a78575061246d906124679061486f565b8a6149d5565b600889015495866124856101508b013560c01c614e44565b03612a695761249a60988a013560c01c614e44565b9a6124b48c6124af60908d013560c01c614e44565b6141bb565b986124bf8d8b614196565b15612a5a576124ce8a8a614196565b60078d015410612a4b576124e96101208c013560c01c614e44565b9d8e15612a3c578e80871490811591612a31575b8115612a26575b50612a17578e49612a17578e6050810204605003612a03578e61252960508202614a75565b906125376040519283613f79565b6050810280835261254790614a75565b601f19013660208401375f5b8181106129935750505f60208092604051918183925191829101835e8101838152039060025afa15612953575f8051818e5b6008821061297157505003612962577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561295e575f928d61262a6040519687958694859463be98686160e01b86527f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801614251565b03915afa80156129535761293f575b506126438d61486f565b9560a0515b84811061274157505060a051988996509450505050505b6008821061271e575050601a9161267b91846003870155614196565b9283600882015561269160098201958654614196565b809555600b81016126a28154614287565b90550180546bffffffffffffffff000000004260201b16906bffffffffffffffff000000001916179055604051948552602085015260408401526060830152608082015233907f36c6781d994e030a156d2f6fa11abcc1cd6814482a5d61f90da3f336318b1d2360a060ff196080511692a360a0516001815580f35b909360019085600a0160031b83013560e01c8660051b60e0031b1794019061265f565b80491561292957602060608961279161275b858a8a614a90565b9390846040519586928884019660805160081c8852604085015284840137810160a051838201520301601f198101845283613f79565b60405191518091835e81019060a05182528060a05192039060025afa1561061757868661284586608086612804876127fc818e6127f5828f7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff0000000160a05151069d614aab565b3597614a90565b939097614a90565b828193604051988997602089019b8d498d5260408a015260608901528688013785019184830160a051815237010160a051815203601f198101835282613f79565b60a0519160a051915190600a5afa3d15612921573d9061286482614a75565b916128726040519384613f79565b825260a0513d90602084013e5b158015612915575b6128fe576040818051810103126102bf57611000604060208301519201519114908115916128d3575b506128bd57600101612648565b638308e1e960e01b60a05152600452602460a051fd5b7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001915014155f6128b0565b50638308e1e960e01b60a05152600452602460a051fd5b50604081511415612887565b60609061287f565b634af69d9960e11b60a05152600452602460a051fd5b5f61294991613f79565b5f60a0528d612639565b6040513d5f823e3d90fd5b5f80fd5b630f2e4cb960e31b5f5260045ffd5b928193600192601c0160031b013560e01c8460051b60e0031b1792018e612585565b806129a16030928b8b614a90565b929092036129f45760306129b6828f8e614a90565b9050036129e55760019160506129ce838f8c90614aab565b359160308285028801916020830137015201612553565b6350320ab160e01b5f5260045ffd5b6338b2168b60e21b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b63b8ff08e960e01b5f5260045ffd5b90508914158f612504565b8581141591506124fd565b63fdac229f60e01b5f5260045ffd5b6357d18d5360e11b5f5260045ffd5b633f8cdc5560e11b5f5260045ffd5b63341203dd60e21b5f5260045ffd5b906001908260140160031b8b013560e01c8360051b60e0031b179101612451565b630b6fac0360e41b5f5260045ffd5b63e843c5eb60e01b5f5260045ffd5b6307a92f1960e51b5f5260045ffd5b60039150612ad381613e29565b1480612ae1575b158c61240f565b50612af5600589015460068a015490614196565b421015612ada565b9050612b0881613e29565b600181141590612409565b3461295e575f36600319011261295e57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461295e575f36600319011261295e5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b3461295e5761030036600319011261295e576005600435101561295e5761010036608319011261295e576001600160401b03610184351161295e5760a0610184353603600319011261295e576101a4356001600160401b03811161295e57612bf9903690600401613d92565b6040366101e319011261295e5760e03661022319011261295e57612c1b6144d0565b600354335f81815260026020526040902054602435939266ffffffffffffff90911660589290921b600160581b600160f81b0316600891821c63ffffffff60381b161791909117901b60ff191660ff1981165f9081526001602052604090205490929060081c6001600160a01b0316331461388f5760a43560ff81161415938461295e5760ff60a4351615801561387d575b61386e5760c43560ff81161415948561295e5761295e5760ff60a4351660ff60c435161161351a5761010435610124351161385f57610144356101643511613850576064351561384157612d06610104356064356143ae565b60056101843560040135101561295e57612d266101843560040135613e29565b61018435600401351561383257612d42608461018435016141c8565b6138235760246101843501359384612d606101843560040135613e29565b60046101843501356003036137bf5750612d7f604461018435016141d5565b803b156137b05760405163c1da869160e01b602082015260048152612dae91612da9602483613f79565b614e29565b90156137b057935b612dcc60646101843501610184356004016141e9565b9050156137a157612dde600435613e29565b600460ff813516118015613772575b612ab757612dff6101c43582856144ee565b6024351561376a575b42841061375b5742612e1c60443586614196565b111561374c5760ff1982165f52600160205260405f20604051612e3e81613f27565b6101e435815261020435602082015294600361022435101561295e576102243561358357610244356001600160601b0360a01b811680910361295e5715801590613577575b801561356b575b801561355f575b8015613553575b8015613547575b613538575b612ead86614517565b1561352957612ebe600435836141a3565b6005820155604435600682015560643560078201558054610100600160a81b0319163360081b610100600160a81b0316178155845160018201556020850151600282015560405194612f1160e087613f79565b6006865260c036602088013760405198612f2c60e08b613f79565b60068a5260c03660208c01378460081c612f458b6142de565b525f612f50886142de565b528061295e5760ff60a4351660a4350361295e5760ff60a4351660ff60c435161161351a5765ffffffffffff610104351161350b5765ffffffffffff61012435116134fc57677fffffffffffffff61014435116134ed57677fffffffffffffff61016435116134de5760ff60a4351660a4350361295e5761295e5760843515156084350361295e57608435156134d85760015b60ff60e4351660e4350361295e576020915f9160a43560ff1660c43560081b61ff00161760109190911b62ff0000161760e43560111b6301fe000016176101043560191b176101243560491b176101443560791b176101643560b81b176130498c6142ff565b526002613055896142ff565b528281519101516040519084820192835260408201526040815261307a606082613f79565b604051918291518091835e8101838152039060025afa15612953575f516130a08961430f565b5260036130ac8661430f565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae2786130d78961431f565b5260046130e38661431f565b526130f46101843560040135613e29565b61018435600401356131058961432f565b5260066131118661432f565b527f000000000000000000000000000000000000000000000000000000000000000061313c8961433f565b5260076131488661433f565b5284518851036134c95761315c85516142ac565b955f5b865181101561319e578061318d8b610c74836001600160401b036131856001978e61434f565b51169261434f565b613197828b61434f565b520161315f565b50876131ab5f8989614f24565b6003840155600e830160843515156084350361295e57805460ff191660ff608435151516178155805460ff60e4351660e4350361295e5763ff00000060e43560181b169061ff0060a43560081b169063ffffff0019161762ff000060c43560101b161717905561010435600f840155610124356010840155610144356011840155610164356012840155601383016132496101843560040135613e29565b60ff1981541660ff610184356004013516179055601483019081556015830160018060a01b0361327e604461018435016141d5565b82546001600160a01b0319169116179055601683016132a76101843560648101906004016141e9565b906001600160401b0382116134b55781906132c28454613fb8565b601f8111613485575b505f90601f831160011461341e575f92613413575b50508160011b915f199060031b1c19161790555b61331b613306608461018435016141c8565b601785019060ff801983541691151516179055565b5543600a820155601a810163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff1982541617905560035463ffffffff811663ffffffff8114612a0357600163ffffffff9101169063ffffffff191617600355335f52600260205260405f20918254946001600160401b038616936001600160401b038514612a03576020966001600160401b036001613403970116906001600160401b0319161790553360ff1986167feefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d15f80a36101c4359285614591565b60015f556040519060ff19168152f35b0135905089806132e0565b909150601f19831691845f5260205f20925f5b81811061346d5750908460019594939210613454575b505050811b0190556132f4565b01355f19600384901b60f8161c19169055898080613447565b91936020600181928787013581550195019201613431565b6134af90855f5260205f20601f850160051c81019160208610611afa57601f0160051c019061421b565b8a6132cb565b634e487b7160e01b5f52604160045260245ffd5b63d088249360e01b5f5260045ffd5b5f612fe3565b63dd6f54df60e01b5f5260045ffd5b63271fb80560e01b5f5260045ffd5b63871a7fa360e01b5f5260045ffd5b63481eb79f60e01b5f5260045ffd5b632cbdc23160e01b5f5260045ffd5b63208e53e560e11b5f5260045ffd5b63e4291a1960e01b5f5260045ffd5b506102e4351515612e9f565b506102c4351515612e98565b506102a4351515612e91565b50610284351515612e8a565b50610264351515612e83565b94506101e43515801590613740575b613529577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156137325760405163441d214d60e11b815260ff1984166004820152946135ed6024870161022435613efe565b610244356001600160601b0360a01b811680910361295e5760448701526102643560648701526102843560848701526102a43560a48701526102c43560c48701526102e43560e4870152608086610104815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af18015612953575f965f915f915f916136d7575b506040519161368c83613f27565b825260208201526018840180546cffffffffffffffffffffffffff19166102243560ff161760989990991c6cffffffffffffffffffffffff0016989098179097556019830155612ea4565b98505050506080863d60801161372a575b816136f560809383613f79565b8101031261295e578551956001600160a01b03198716870361295e57602081015160408201516060909201519091908c61367e565b3d91506136e8565b6203eb8f60e01b5f5260045ffd5b50610204351515613592565b637616640160e01b5f5260045ffd5b632ca4094f60e21b5f5260045ffd5b429350612e08565b5061377e600435613e29565b60043515158015612ded5750613795600435613e29565b60036004351415612ded565b630f8b932160e11b5f5260045ffd5b63562e597160e11b5f5260045ffd5b936001600160a01b036137d7610184356044016141d5565b166137b0578415613808576137f26101843560040135613e29565b6004610184358101351480613817575b15612db6575b635e32eadd60e01b5f5260045ffd5b508460a01c1515613802565b63f545b7bf60e01b5f5260045ffd5b63f37f7b5d60e01b5f5260045ffd5b632c45be6f60e01b5f5260045ffd5b636d403ec760e11b5f5260045ffd5b63207ea56d60e01b5f5260045ffd5b63ac38930b60e01b5f5260045ffd5b505f9450601060a43560ff1611612cad565b635040c4d360e11b5f5260045ffd5b3461295e57604036600319011261295e576138b7613d81565b602435600581101561295e5760ff198216918215613a5a5763ffffffff808060035460401c16169160401c1603613a4b576138f181613e29565b600460ff821611612ab7575f8281526001602052604090208054600881901c6001600160a01b03168015613a3c573303613a2e5760ff1661393283826143e5565b15612ab757600582015492600683019261394d845486614196565b9361395783613e29565b60018314958680613a25575b612aa85761397084613e29565b6003841480613a1b575b612aa8575f5160206150eb5f395f51905f529660409661399b868b966141a3565b6139a486613e29565b81613a11575b506139d2575b5050508251916139bf81613e29565b82526139ca81613e29565b6020820152a2005b7f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba91613a00602092426141bb565b8091558651908152a28486806139b0565b90504210896139aa565b508542101561397a565b50804210613963565b6282b42960e81b5f5260045ffd5b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b63cbf4a64560e01b5f5260045ffd5b3461295e57602036600319011261295e5760ff19613a85613d81565b165f52600160205260405f20604051613a9d81613f0b565b815460ff8116613aac81613e29565b825260081c6001600160a01b03166020820152613acb60018301613f9a565b60408201526003820154606082015260048201604051808260208294549384815201905f5260205f20925f5b818110613c59575050613b0c92500382613f79565b6080820152600582015460a0820190815260068301549060c08301918252600784015460e084015260088401546101008401526009840154610120840152600a840154610140840152600b840154610160840152613b6c600c8501613ff0565b610180840152600d8401546101a0840152613b89600e8501614090565b6101c0840152613b9b601385016140f2565b6101e084015260188401549260ff8416946003861015613c45576001600160401b03601a6103009260ff610c9c9860209a6102008801526001600160601b0360a01b8160981b1661022088015261ffff8160681c16610240880152818160781c1661026088015261ffff8160801c1661028088015260901c1615156102a086015260198101546102c0860152015463ffffffff81166102e0850152871c1691015251905190614196565b634e487b7160e01b5f52602160045260245ffd5b8454835260019485019486945060209093019201613af7565b3461295e57604036600319011261295e57613c8b613d81565b60243560ff198216918215613a5a5763ffffffff808060035460401c16169160401c1603613a4b575f8281526001602052604090208054600881901c6001600160a01b03168015613a3c573303613a2e5760ff16613ce881613e29565b8015159081613d6c575b50612ab757613d0a6005820154600683015490614196565b421015612aa85781158015613d5f575b61384157817f36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be92600783613d54600f6020960154846143ae565b0155604051908152a2005b5060088101548210613d1a565b60039150613d7981613e29565b141584613cf2565b6004359060ff198216820361295e57565b9181601f8401121561295e578235916001600160401b03831161295e576020838186019501011161295e57565b9181601f8401121561295e578235916001600160401b03831161295e576020808501948460051b01011161295e57565b3461295e575f36600319011261295e5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b60051115613c4557565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60e0809180511515845260ff602082015116602085015260ff604082015116604085015260ff60608201511660608501526080810151608085015260a081015160a085015260c081015160c08501520151910152565b908151613eb981613e29565b81526020820151602082015260018060a01b036040830151166040820152608080613ef3606085015160a0606086015260a0850190613e33565b930151151591015290565b906003821015613c455752565b61032081019081106001600160401b038211176134b557604052565b604081019081106001600160401b038211176134b557604052565b61010081019081106001600160401b038211176134b557604052565b60a081019081106001600160401b038211176134b557604052565b90601f801991011681019081106001600160401b038211176134b557604052565b90604051613fa781613f27565b602060018294805484520154910152565b90600182811c92168015613fe6575b6020831014613fd257565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613fc7565b9060405191825f82549261400384613fb8565b808452936001811690811561406e575060011461402a575b5061402892500383613f79565b565b90505f9291925260205f20905f915b818310614052575050906020614028928201015f61401b565b6020919350806001915483858901015201910190918492614039565b90506020925061402894915060ff191682840152151560051b8201015f61401b565b9060405161409d81613f42565b60e06004829460ff815481811615158652818160081c166020870152818160101c16604087015260181c16606085015260018101546080850152600281015460a0850152600381015460c08501520154910152565b906040516140ff81613f5e565b608060ff600483958281541661411481613e29565b85526001810154602086015260028101546001600160a01b0316604086015261413f60038201613ff0565b60608601520154161515910152565b3461295e575f36600319011261295e5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b9060018201809211612a0357565b91908201809211612a0357565b906141ad81613e29565b60ff80198354169116179055565b91908203918211612a0357565b35801515810361295e5790565b356001600160a01b038116810361295e5790565b903590601e198136030182121561295e57018035906001600160401b03821161295e5760200191813603831361295e57565b818110614226575050565b5f815560010161421b565b908060209392818452848401375f828201840152601f01601f1916010190565b9492909361427692614284979587526020870152608060408701526080860191614231565b926060818503910152614231565b90565b5f198114612a035760010190565b6001600160401b0381116134b55760051b60200190565b906142b682614295565b6142c36040519182613f79565b82815280926142d4601f1991614295565b0190602036910137565b8051156142eb5760200190565b634e487b7160e01b5f52603260045260245ffd5b8051600110156142eb5760400190565b8051600210156142eb5760600190565b8051600310156142eb5760800190565b8051600410156142eb5760a00190565b8051600510156142eb5760c00190565b80518210156142eb5760209160051b010190565b60206040818301928281528451809452019201905f5b8181106143865750505090565b8251845260209384019390920191600101614379565b60408110156142eb5760051b60240190565b80156143d15764e8d4a5100004106143c257565b63eba5c29b60e01b5f5260045ffd5b634e487b7160e01b5f52601260045260245ffd5b6143ee81613e29565b6143f782613e29565b8082146144a45761440781613e29565b6002811480156144bd575b80156144aa575b6144a45761442681613e29565b80156144855760039061443881613e29565b1461444257505f90565b61444b81613e29565b8015908115614470575b811561445f575090565b6001915061446c81613e29565b1490565b905061447b81613e29565b6002811490614455565b5061448f81613e29565b6003811490811561447057811561445f575090565b50505f90565b506144b481613e29565b60018114614419565b506144c781613e29565b60048114614412565b60025f54146144df5760025f55565b633ee5aeb560e01b5f5260045ffd5b501590811561450e575b506144ff57565b635e765b2560e11b5f5260045ffd5b9050155f6144f8565b60208151910151908015801561457a575b8015614563575b6144a4575f51602061510b5f395f51905f52808281930992800981808080848709620292f80960010893620292fc09081490565b505f51602061510b5f395f51905f5282101561452f565b505f51602061510b5f395f51905f52811015614528565b600c820193916001600160401b0383116134b55785906145b18654613fb8565b601f81116146c7575b505f95601f851160011461463d5790600d91857f77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f0985f91614632575b508660011b905f198860031b1c19161790555b0155614622604051938493604085526040850191614231565b94602083015260ff1916930390a2565b90508701355f6145f6565b601f19851696815f5260205f20975f5b8181106146ac575097600d939291877f77e65e34059d7d8e9b78033507a4bc1fbac6bd614e0703b7ca2c6c7d5fa4e1f09a10614693575b5050600186811b019055614609565b8801355f19600389901b60f8161c191690555f80614684565b888301358a556001909901988a95506020928301920161464d565b6146f190875f5260205f20601f870160051c81019160208810611afa57601f0160051c019061421b565b5f6145ba565b60ff198116908115613a5a5763ffffffff808060035460401c16169160401c1603613a4b575f52600160205260405f209060018060a01b03825460081c1615613a3c57565b61474f6005820154600683015490614196565b9063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016801983116147d3576147af601a6147b59301546001600160401b038160201c168581115f146147c65763ffffffff90915b1690614196565b92614196565b808210156147c1575090565b905090565b5063ffffffff85916147a8565b5050505f1990565b906102000361482b5760016147f3823560c01c614e44565b1490811591614812575b5061480457565b6254aabb60e81b5f5260045ffd5b61482391506008013560c01c614e44565b15155f6147fd565b630f61e7fd60e21b5f5260045ffd5b5f905f905b6008821061484c57505090565b90916001908360020160031b83013560e01c8460051b60e0031b1792019061483f565b77ffffffffffffffff0000000000000000ffffffffffffffff8160081c9160081b917cff000000ff000000ff000000ff000000ff000000ff000000ff000000ff7dff000000ff000000ff000000ff000000ff000000ff000000ff000000ff007fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff0085167eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff84161760101c941691161760101b9179ffff000000000000ffff000000000000ffff000000000000ffff7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff847dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff84161760201c941691161760201b7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff82821673ffffffff000000000000000000000000ffffffff85161760401b93161760401c16178060801b9060801c1790565b90600360ff6013840154166149e981613e29565b03614a6957601582015460405163650e5fcf60e01b6020820152602480820193909352918252614a2891906001600160a01b0316612da9604483613f79565b9015918215614a60575b8215614a56575b8215614a48575b505061380857565b600a01541190505f80614a40565b4382119250614a39565b81159250614a32565b90601401540361380857565b6001600160401b0381116134b557601f01601f191660200190565b908210156142eb57614aa79160051b8101906141e9565b9091565b91908110156142eb5760051b0190565b9291909180158015614b77575b614b6f575f925f5b828110614b435750818414614b3a57614aea906004614e91565b92805b614af8575050501490565b5f19019283906004821c600116614b2557614b1f90614b18838587614aab565b3590614f01565b93614aed565b614b1f90614b34838587614aab565b35614f01565b50505050505f90565b614b4e818486614aab565b35614b5c575b600101614ad0565b935060018401808511612a035793614b54565b505050505f90565b5060408111614ac8565b9060ff600e82015460081c1691614b97836142ac565b92601883015460ff8160781c169081614c97575b50505060048254928160ff85169460ff1916178155018351906001600160401b0382116134b557600160401b82116134b5578054828255808310614c7b575b5060208501905f5260205f205f5b838110614c675750505050905f5160206150eb5f395f51905f5260409260ff1916928392815190614c2881613e29565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e60405180614c62339582614363565b0390a3565b600190602084519401938184015501614bf8565b614c9190825f528360205f20918201910161421b565b5f614bea565b6019850154604051630222162f60e21b8152609883901b6001600160a01b03191660048201526024810191909152606882901c61ffff16604482015260648101929092525f826084817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215612953575f905f93614d89575b5015614d7a575f919060801c61ffff16825b848110614d3c575050614bab565b600182821c1615614d50575b600101614d2e565b92614d7281614d616001938661434f565b51614d6c878c61434f565b52614287565b939050614d48565b630e0d4dc560e41b5f5260045ffd5b9250503d805f843e614d9b8184613f79565b82019160408184031261295e57805190811515820361295e576020810151906001600160401b03821161295e57019280601f8501121561295e578351614de081614295565b94614dee6040519687613f79565b81865260208087019260051b82010192831161295e57602001905b828210614e19575050505f614d1c565b8151815260209182019101614e09565b6020915f91838251920190620186a0fa601f3d1116905f5190565b66ff00ff00ff00ff67ff00ff00ff00ff008260081b169160081c161765ffff0000ffff67ffff0000ffff00008260101b169160101c161767ffffffff000000008160201b169060201c1790565b602091614eb95f92614eaf6001600160401b038060c01b9216614e44565b60c01b169161486f565b604051908482019283526028820152600160f81b604882015260298152614ee1604982613f79565b604051918291518091835e8101838152039060025afa15612953575f5190565b5f9060209260405190848201928352604082015260408152614ee1606082613f79565b9081518015614b6f57600181146150d957604084146150ca575f915f5b8281106150a15750614f5b614f5684846141bb565b6142ac565b91614f69614f5685836141bb565b91614f7c614f76866142ac565b956142ac565b955f9081805b888a8c888710614fc15795505050505050614faa94939250614fa49150614188565b91614f24565b91614fa4614fb89394614188565b61428491614f01565b906150066001614ff98a9594614fe7614fda8c8c61434f565b516001600160401b031690565b906001600160401b03809216901c1690565b166001600160401b031690565b61505f57505050615056600191615051615023614fda888861434f565b61502d888a61434f565b51615038848d61434f565b52615043838d61434f565b906001600160401b03169052565b614287565b935b0192614f82565b61509b92615043868094615051946150958c9a61508e8c60019c9f8f61508891614fda9161434f565b9861434f565b519261434f565b5261434f565b91615058565b926150c36001916150bd83614ff98a614fe7614fda8b8d61434f565b90614196565b9301614f41565b63394fd24160e21b5f5260045ffd5b5090506150e691506142de565b519056fe56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122078f9343de9ddff9596db84f10a74ed8c8c59450aa389595f4e2570a0f5cca2bb64736f6c634300081c003360e0806040523461011157602061002e60049261106880380380916100248285610115565b833981019061014c565b336080526001600160a01b031660a081905260405163ebe86c1360e01b815292839182905afa908115610106575f916100d7575b506001600160a01b031660c052604051610efc908161016c82396080518181816101530152818161037701528181610756015281816107cd0152610bce015260a051818181610186015281816108b701528181610a2f0152610c43015260c05181818160c30152818161047001526107fd0152f35b6100f9915060203d6020116100ff575b6100f18183610115565b81019061014c565b5f610062565b503d6100e7565b6040513d5f823e3d90fd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761013857604052565b634e487b7160e01b5f52604160045260245ffd5b9081602091031261011157516001600160a01b0381168103610111579056fe60806040526004361015610011575f80fd5b5f5f3560e01c8063088858bc146108e6578063481c6a75146108a257806364fe50ac146107b1578063702574b3146107855780637b10399914610740578063883a429a1461034e578063ae977ede146100f2578063ebe86c13146100ad5763f08c9b3c1461007d575f80fd5b346100aa57806003193601126100aa576020610097610c34565b6040516001600160a01b03199091168152f35b80fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760603660031901126100aa5761010c610977565b9060243591604435916001600160401b0383116100aa57366023840112156100aa578260040135936001600160401b03851161034a573660248660071b8601011161034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b579093927f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316928592916001600160a01b031990911690835b8688101561032c578760071b82019060845f516020610ea75f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d2899602485013509925f516020610ea75f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d289960648301350960405194633d98dab360e11b865287600487015288602487015260448601526044820135606486015282850152013560a483015260208260c481898b5af19182156103215786926102e2575b508861029c5750600190975b01966101c1565b979061ffff8916908282018092116102ce5761ffff16036102bf57600190610295565b6357149e2560e01b8552600485fd5b634e487b7160e01b87526011600452602487fd5b9091506020813d8211610319575b816102fd6020938361099f565b810103126103155761030e906109eb565b905f610289565b8580fd5b3d91506102f0565b6040513d88823e3d90fd5b60209061ffff60405191168152f35b633217675b60e21b8252600482fd5b5080fd5b50346100aa576101003660031901126100aa5761036961098e565b60e036602319011261034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b576103ae90610bab565b602435600381101561073c5760021490811561072e576044356001600160a01b03198116810361072a57915b6040908151906103ea838361099f565b6001825260208201601f198401368237825115610716573090521561070e5784935b82519461010086018681106001600160401b038211176106fa57845260028110156106e6578552602085019486865283810192835260608101601081526080820188815260a0830189815260c08401918a835260e08501938b855260018060a01b037f00000000000000000000000000000000000000000000000000000000000000001697883b156106e257895197631bb64f5960e21b89526001600160601b0360a01b169b8c60048a01528b60248a015261010060448a01526102048901975160028110156106ce5791899693918f989593610104899b989b01525115156101248801525194610100610144880152855180915260206102248801960190885b81811061069d57505050916001600160401b03869798818096959461ffff829651166101648b0152511661018489015251166101a487015251166101c485015251166101e48301526064356064830152608435608483015260a43560a483015260c43560c483015260e43560e4830152038183865af180156106935761067a575b5093816044958151968780926303e95d1360e21b82528860048301528760248301525afa92831561066e57608095829461060b575b505f516020610ea75f395f51905f52917f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e0097542918451968752602087015209908301526060820152f35b7f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e00975429194505f516020610ea75f395f51905f52925061065e90843d8611610667575b610656818361099f565b810190610c1e565b949150916105c3565b503d61064c565b509051903d90823e3d90fd5b61068586809261099f565b61068f575f61058e565b8480fd5b83513d88823e3d90fd5b92949750929497509794602080600192838060a01b038c511681520199019101908e9794928a97949299969961050d565b634e487b7160e01b8f52602160045260248ffd5b8c80fd5b634e487b7160e01b87526021600452602487fd5b634e487b7160e01b88526041600452602488fd5b60019361040c565b634e487b7160e01b87526032600452602487fd5b8380fd5b610736610c34565b916103da565b8280fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760203660031901126100aa5760206107a96107a461098e565b610bab565b604051908152f35b503461088f57606036600319011261088f576107cb610977565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610893577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561088f575f9160648392604051948593849263a59b7a4d60e01b84526001600160601b0360a01b166004840152602435602484015260443560448401525af1801561088457610876575080f35b61088291505f9061099f565b005b6040513d5f823e3d90fd5b5f80fd5b633217675b60e21b5f5260045ffd5b3461088f575f36600319011261088f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461088f57608036600319011261088f576108ff610977565b60443561ffff8116810361088f576064359061ffff8216820361088f5761092992602435906109fa565b906040519182916040830190151583526040602084015281518091526020606084019201905f5b81811061095e575050500390f35b8251845285945060209384019390920191600101610950565b600435906001600160a01b03198216820361088f57565b6004359060ff198216820361088f57565b90601f801991011681019081106001600160401b038211176109c057604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b0381116109c05760051b60200190565b519061ffff8216820361088f57565b9261ffff1693610a09856109d4565b610a16604051918261099f565b858152601f19610a25876109d4565b01366020830137937f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925f5b61ffff811688811015610b9e578061ffff88160161ffff8111610b8a57604051639bbada6760e01b81526001600160a01b0319861660048201526024810185905261ffff9190911660448201526060816064818a5afa908115610884575f91610b14575b50602081015115610b0657604001519088511115610af257600582901b621fffe01688016020015260010161ffff16610a5a565b634e487b7160e01b5f52603260045260245ffd5b505f98509395505050505050565b90506060813d8211610b82575b81610b2e6060938361099f565b8101031261088f5760405190606082018281106001600160401b038211176109c057604052610b5c816109eb565b8252602081015190811515820361088f576040916020840152015160408201525f610abe565b3d9150610b21565b634e487b7160e01b5f52601160045260245ffd5b5050505093505050600191565b5f516020610ea75f395f51905f5290604051602081019146835260018060a01b037f000000000000000000000000000000000000000000000000000000000000000016604083015260ff1916606082015260608152610c0b60808261099f565b519020068015610c185790565b50600190565b919082604091031261088f576020825192015190565b60405163a4adcd7f60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316602082600481845afa918215610884575f92610e62575b506040516323488be560e01b8152602081600481855afa8015610884575f90610e1c575b63ffffffff60401b915060401b165f5b6001600160401b0381166008811080610e0a575b15610dfb576001600160401b038516036001600160401b038111610b8a576001600160401b036001600160601b0360a01b9116831760a01b1660405163d397925360e01b8152816004820152602081602481885afa8015610884575f90610dbf575b60ff9150166010811015610daf57604051906356cbb5f360e01b82528260048301526024820152604081604481885afa9081610d92575b50610d8a57506001600160401b03905b166001600160401b038114610b8a57600101610cb6565b935050505090565b610da99060403d811161066757610656818361099f565b50610d63565b50506001600160401b0390610d73565b506020813d8211610df3575b81610dd86020938361099f565b8101031261088f575160ff8116810361088f5760ff90610d2c565b3d9150610dcb565b63081ea97160e31b5f5260045ffd5b50806001600160401b03861611610cca565b506020813d602011610e5a575b81610e366020938361099f565b8101031261088f575163ffffffff8116810361088f5763ffffffff60401b90610ca6565b3d9150610e29565b9091506020813d602011610e9e575b81610e7e6020938361099f565b8101031261088f57516001600160401b038116810361088f57905f610c82565b3d9150610e7156fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220ba4af154f443c35cd7f28236da90badae727334286f58681a4f008089332408d64736f6c634300081c0033",
}

// ProcessRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProcessRegistryMetaData.ABI instead.
var ProcessRegistryABI = ProcessRegistryMetaData.ABI

// ProcessRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProcessRegistryMetaData.Bin instead.
var ProcessRegistryBin = ProcessRegistryMetaData.Bin

// DeployProcessRegistry deploys a new Ethereum contract, binding an instance of ProcessRegistry to it.
func DeployProcessRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID uint32, _ziskVerifier common.Address, _batchProgramVK [32]byte, _resultsProgramVK [32]byte, _rootCVadcopFinal [32]byte, _ballotVKHash [32]byte, _dkgManager common.Address, _defaultGrace uint32, _graceFloor uint32, _graceCeil uint32, _graceMaxTotal uint32, _noticeMin uint32) (common.Address, *types.Transaction, *ProcessRegistry, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProcessRegistryBin), backend, _chainID, _ziskVerifier, _batchProgramVK, _resultsProgramVK, _rootCVadcopFinal, _ballotVKHash, _dkgManager, _defaultGrace, _graceFloor, _graceCeil, _graceMaxTotal, _noticeMin)
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
