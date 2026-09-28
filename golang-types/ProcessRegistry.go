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
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_ziskVerifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_batchProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_resultsProgramVK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rootCVadcopFinal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_ballotVKHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_dkgManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMaxValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueSumTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BallotModeMinValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlobCountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CircuitFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DKGDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccumulator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobCommitmentLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidBlobOpening\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlobsDigest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDKGParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncryptionKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGroupSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInclusionProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKZGProofLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOccupiedBefore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxPossibleResultCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"MissingBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoBlobs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsAlreadyRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultsNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtMaxLevelsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nBlobs\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"indexed\":false,\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"firstIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"name\":\"ResultsDecryptionRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"aidFor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotVKHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dkgAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"finalizeResultsFromDKG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"}],\"name\":\"genesisRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"}],\"internalType\":\"structDAVINCITypes.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"epochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint256\",\"name\":\"orgPKx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orgPKy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popAy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"popZ\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.DKGParams\",\"name\":\"dkg\",\"type\":\"tuple\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"\",\"type\":\"bytes31\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"groupSize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structDAVINCITypes.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"enumDAVINCITypes.KeyMode\",\"name\":\"keyMode\",\"type\":\"uint8\"},{\"internalType\":\"bytes12\",\"name\":\"dkgEpochId\",\"type\":\"bytes12\"},{\"internalType\":\"uint16\",\"name\":\"dkgFirstIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"dkgCount\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"dkgZeroSkipped\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"dkgResultsRequested\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"dkgAid\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256[64]\",\"name\":\"accumulator\",\"type\":\"uint256[64]\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"name\":\"requestResultsDecryption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultsProgramVK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"sk\",\"type\":\"uint256\"}],\"name\":\"revealProcessKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootCVadcopFinal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"components\":[{\"internalType\":\"enumDAVINCITypes.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structDAVINCITypes.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"enumDAVINCITypes.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes31\",\"name\":\"processId\",\"type\":\"bytes31\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"commitments\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"kzgProofs\",\"type\":\"bytes[]\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ziskVerifier\",\"outputs\":[{\"internalType\":\"contractIZiskVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610140806040523461025e5760e081615bf480380380916100208285610262565b83398101031261025e57805163ffffffff8116810361025e5761004560208301610285565b916040810151606082015160808301519161006760c060a08601519501610285565b60015f55956001600160a01b031680158015610256575b801561024e575b8015610246575b801561023e575b61022f5760805260a05260c05260e0526101005260035467ffffffff000000006bffffffff0000000000000000604051602081019063ffffffff60e01b8660e01b1682523060601b6024820152601881526100ef603882610262565b51902060401b169260201b1690640100000000600160601b031916171760035560018060a01b031680155f146101cf57505f5b610120526040516148f2908161029a8239608051818181611bdd01528181611e8c0152612d93015260a051818181612dff01526137a3015260c051818181611efc0152613b16015260e051818181611edb015281816126660152612dde015261010051818181610e0901528181611944015261329c01526101205181818161050b015281816108550152818161139301528181611449015281816121550152818161222801526144740152f35b604051906110688083016001600160401b0381118482101761021b576020928492614b8c843981520301905ff08015610210576001600160a01b0316610122565b6040513d5f823e3d90fd5b634e487b7160e01b5f52604160045260245ffd5b6306f9b90760e01b5f5260045ffd5b508415610093565b50831561008c565b508215610085565b50811561007e565b5f80fd5b601f909101601f19168101906001600160401b0382119082101761021b57604052565b51906001600160a01b038216820361025e5756fe60c06040526004361015610011575f80fd5b5f60a0525f3560e01c8063026cdee81461361a57806304ed00fa1461343f578063082b642e146132bf5780630e2ebcf7146132855780631fdf344914612b355780634c0acc5614611b9c57806359d821c414612689578063621153381461264d57806362fa11fc1461260457806368141f2c146125855780636c7aff7f146122c5578063702574b31461220d57806372c628ef146121f0578063734177051461210c578063766422e014611db2578063784df74a14611c0c5780637f64b72f14611bc6578063848df54014611ba1578063946544bf14611b9c5780639b46499414611a16578063aa2402211461015c578063adc879e9146119ef578063bf74291e14611703578063cddf08bc146116db578063d8738c8114610884578063e16d5b7c1461083e578063e965eead1461024a578063f1431097146101615763f9aa44991461015c575f80fd5b613aff565b346102445760203660031901126102445761017a61371e565b610182613e81565b61018b81613e9f565b601781015460ff8116600381101561022c571561021957815460ff166101b0816137c6565b60028114908115610205575b506101f25760901c60ff16156101df576101d591614313565b60a0516001815580f35b630e0d4dc560e41b60a05152600460a051fd5b6307a92f1960e51b60a05152600460a051fd5b60049150610212816137c6565b14846101bc565b6365b75c3960e01b60a05152600460a051fd5b634e487b7160e01b60a051526021600452602460a051fd5b60a05180fd5b34610244576108403660031901126102445761026461371e565b366108241161024457610824356001600160401b0381116102445761028d90369060040161375c565b610295613e81565b61029e83613e9f565b916017830190815460ff8116600381101561022c57156102195760ff8160901c1661082b57845460ff8116929091906102d6846137c6565b600284148015610818575b6101f2576102ee846137c6565b60018414159586806107fd575b6107ea5760a0515b604081106107c45750602060405181810161080060248237610800825261032c6108208361392a565b60405191518091835e81019060a05182528060a05192039060025afa1561058d5761035f9160a0515160038a015461424d565b156107b15760ff60901b1916600160901b178084559361037e836137c6565b610772575b505060ff600d84015460081c169161039a83613c39565b926103a8604051948561392a565b808452601f196103b782613c39565b0160a0515b81811061075057505060a05191829182905b8082106105df57505060a05182159590919086156104a0575b50508354607883901b60ff60781b1664ffffffffff60681b19909116606883901b61ffff60681b161717608093841b61ffff60801b16179384905560188601546040805160989690961b6001600160a01b0319168652602086019190915261ffff919091169084015260ff16606083015260ff198516917fdca6075f07367349a836825d3ee7c35c204d2e727ae81a5b7a48aca95b9c270b9190a26104905760a0516001815580f35b61049991614313565b80806101d5565b838152601888015460405163574bbf6f60e11b815260989390931b6001600160a01b031916600484015260248301526060604483015280516064830181905260a0519293508392608484019260200191905b81811061059a57505060a05160209392839003915082907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af190811561058d5760a0519161054f575b508261ffff6103e7565b90506020813d602011610585575b8161056a6020938361392a565b81010312610244575161ffff811681036102445760ff610545565b3d915061055d565b6040513d60a051823e3d90fd5b9180945092909251819060a051915b600483106105c957505050602060806001920194019101918493926104f2565b60208060019284518152019201920191906105a9565b9092600284901b6001600160fe1b03851685036106c1576105ff81613d4d565b3591600182018083116106c15761061590613d4d565b3560a05160028401918285116106c15761062e83613d4d565b35158061072d575b861580610723575b610704576106f15760405195608087018781106001600160401b038211176106d957604052865260208601526106c15761067790613d4d565b356040840152600382019182106106c1576001926106976106b893613d4d565b3560608201526106a7828b613cf3565b526106b2818a613cf3565b50613bcf565b935b01906103ce565b634e487b7160e01b60a051526011600452602460a051fd5b634e487b7160e01b60a051526041600452602460a051fd5b632a23591560e21b60a05152600460a051fd5b945050505094959150156106f15760019061ffff82871b1617946106ba565b506001821461063e565b5060a0519150600385018086116106c157610749600191613d4d565b3514610636565b604051602091906080610763818361392a565b368237828289010152016103bc565b60ff1916600117845560405190610788816137c6565b81526001602082015260ff198516905f51602061487d5f395f51905f5290604090a28480610383565b6303cd656760e61b60a05152600460a051fd5b5f51602061489d5f395f51905f526107db82613d4d565b3510156106f157600101610303565b63e843c5eb60e01b60a05152600460a051fd5b50610811600589015460068a015490613b47565b42106102fb565b50610822846137c6565b600484146102e1565b636c8ae94b60e11b60a05152600460a051fd5b346102445760a051366003190112610244576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610244576102e0366003190112610244576005600435101561024457610100366083190112610244576001600160401b0361018435116102445760a06101843536036003190112610244576101a4356001600160401b038111610244576108f090369060040161372f565b906040366101c31901126102445760e03661020319011261024457610913613e81565b60a08051600354339182905260026020908152835160408082205466ffffffffffffff16605886901b600160581b600160f81b0316600895861c63ffffffff60381b161717841b60ff19169182905260019092529351205492949260243592911c6001600160a01b0316146116c85760a43560ff8116141593846102445760ff60a435161580156116b6575b6116a35760c43560ff811614159485610244576102445760ff60a4351660ff60c4351611611309576101043561012435116116905761014435610164351161167d576064351561166a576109f861010435606435613d5f565b60056101843560040135101561024457610a1860046101843501356137c6565b61018435600401351561165757610a3460846101843501613d40565b61164457610184356024810135908190610a5190600401356137c6565b60046101843501356003036115dd5750610a7060446101843501613bdd565b803b156115ca5760405163c1da869160e01b602082015260048152610a9f91610a9a60248361392a565b6145bb565b90156115ca575b610abc6064610184350161018435600401613bf1565b9050156115b757610ace6004356137c6565b600460ff813516118015611588575b6101f25760243515611580575b42841061156d5742610afe60443586613b47565b111561155a5760ff19831660a051526001602052604060a051209160a05150604051610b29816138d8565b6101c43581526101e4356020820152946003610204351015610244576102043561137e57610224356001600160601b0360a01b81168091036102445715801590611372575b8015611366575b801561135a575b801561134e575b8015611342575b61132f575b610b98866141d3565b1561131c57610ba960043585613b54565b6005840155604435600684015560643560078401558254610100600160a81b0319163360081b610100600160a81b0316178355845160018401556020850151600284015560405194610bfc60e08761392a565b6006865260c036602088013760405198610c1760e08b61392a565b60068a5260c03660208c01378560081c610c308b613c82565b5260a051610c3d88613c82565b52806112b05760ff60a4351660a435036112b05760ff60a4351660ff60c43516116113095765ffffffffffff61010435116112f65765ffffffffffff61012435116112e357677fffffffffffffff61014435116112d057677fffffffffffffff61016435116112bd5760ff60a4351660a435036112b0576112b0576084351515608435036112b057608435156112b4576001905b60ff60e4351660e435036112b05760209160a43560ff1660c43560081b61ff00161760109190911b62ff0000161760e43560111b6301fe000016176101043560191b176101243560491b176101443560791b176101643560b81b17610d358b613ca3565b526002610d4188613ca3565b528181519101519060405191838301918252604083015260408252610d6760608361392a565b60405191518091835e81019060a05182528060a05192039060025afa1561058d5760a05151610d9589613cb3565b526003610da186613cb3565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae278610dcc89613cc3565b526004610dd886613cc3565b52610de961018435600401356137c6565b6101843560040135610dfa89613cd3565b526006610e0686613cd3565b527f0000000000000000000000000000000000000000000000000000000000000000610e3189613ce3565b526007610e3d86613ce3565b52845188510361129d57610e518551613c50565b9560a0515b8651811015610e9c5780610e8b8b610e84836001600160401b03610e7c6001978e613cf3565b511692613cf3565b5190614623565b610e95828b613cf3565b5201610e56565b5087610eac878960a051916146b6565b60038601556001600160401b0382116106d9578190610ece600c870154613969565b601f8111611266575b5060a05190601f83116001146111f65760a051926111eb575b50508160011b915f199060031b1c191617600c8401555b600d8301608435801515900361024457805460ff191660ff608435151516178155805460e43560ff811690036102445763ff00000060e43560181b169061ff0060a43560081b169063ffffff0019161762ff000060c43560101b161717905561010435600e84015561012435600f8401556101443560108401556101643560118401556012830160a05150610fa261018435600401356137c6565b60ff1981541660ff610184356004013516179055601383019081556014830160018060a01b03610fd760446101843501613bdd565b82546001600160a01b031916911617905560158301611000610184356064810190600401613bf1565b906001600160401b0382116106d957819061101b8454613969565b601f81116111ac575b5060a05190601f831160011461113f5760a05192611134575b50508160011b915f199060031b1c19161790555b61107861106360846101843501613d40565b601685019060ff801983541691151516179055565b55600a4391015560035463ffffffff811663ffffffff81146106c157600163ffffffff9101169063ffffffff1916176003553360a051526002602052604060a05120908154916001600160401b038316926001600160401b0384146106c1576001600160401b0360016020950116906001600160401b0319161790553360ff1982167feefcd49abfaf7291d2e1c15f581f85a3610d4f103666e075bd536faef609e1d160a05160a051a3600160a051556040519060ff19168152f35b01359050878061103d565b909150601f198316918460a05152602060a051209260a0515b818110611194575090846001959493921061117b575b505050811b019055611051565b01355f19600384901b60f8161c1916905587808061116e565b91936020600181928787013581550195019201611158565b6111db908560a05152602060a05120601f850160051c810191602086106111e1575b601f0160051c0190613c23565b88611024565b90915081906111ce565b013590508680610ef0565b60a08051600c890190525160208120909450915b601f198416851061124e576001945083601f19811610611235575b505050811b01600c840155610f07565b01355f19600384901b60f8161c19169055868080611225565b8181013583556020948501946001909301920161120a565b61129790600c880160a05152602060a05120601f850160051c810191602086106111e157601f0160051c0190613c23565b87610ed7565b63d088249360e01b60a05152600460a051fd5b5f80fd5b60a05190610cd1565b63dd6f54df60e01b60a05152600460a051fd5b63271fb80560e01b60a05152600460a051fd5b63871a7fa360e01b60a05152600460a051fd5b63481eb79f60e01b60a05152600460a051fd5b632cbdc23160e01b60a05152600460a051fd5b63208e53e560e11b60a05152600460a051fd5b63e4291a1960e01b60a05152600460a051fd5b506102c4351515610b8a565b506102a4351515610b83565b50610284351515610b7c565b50610264351515610b75565b50610244351515610b6e565b94506101c4351580159061154e575b61131c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161561153c5760405163441d214d60e11b815260ff1985166004820152946113e8602487016102043561389b565b610224356001600160a01b031981168103610244576001600160a01b03191660448701526102443560648701526102643560848701526102843560a48701526102a43560c48701526102c43560e487015260a05160809087906101049082907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1801561058d5760a051968791829182916114e1575b5060405191611496836138d8565b825260208201526017860180546cffffffffffffffffffffffffff19166102043560ff161760989990991c6cffffffffffffffffffffffff0016989098179097556018850155610b8f565b98505050506080863d608011611534575b816114ff6080938361392a565b81010312610244578551956001600160a01b03198716870361024457602081015160408201516060909201519091908c611488565b3d91506114f2565b6203eb8f60e01b60a05152600460a051fd5b506101e435151561138d565b637616640160e01b60a05152600460a051fd5b632ca4094f60e21b60a05152600460a051fd5b429350610aea565b506115946004356137c6565b60043515158015610add57506115ab6004356137c6565b60036004351415610add565b630f8b932160e11b60a05152600460a051fd5b63562e597160e11b60a05152600460a051fd5b6001600160a01b036115f461018435604401613bdd565b166115ca5780156116255761160f60046101843501356137c6565b6004610184358101351480611638575b15610aa6575b635e32eadd60e01b60a05152600460a051fd5b508060a01c151561161f565b63f545b7bf60e01b60a05152600460a051fd5b63f37f7b5d60e01b60a05152600460a051fd5b632c45be6f60e01b60a05152600460a051fd5b636d403ec760e11b60a05152600460a051fd5b63207ea56d60e01b60a05152600460a051fd5b63ac38930b60e01b60a05152600460a051fd5b505f9450601060a43560ff161161099f565b635040c4d360e11b60a05152600460a051fd5b346102445760a05136600319011261024457602063ffffffff60035460401c16604051908152f35b34610244576101803660031901126102445761171d61371e565b61010036602319011261024457604036610123190112610244576101643560058110156102445760405190611751826138d8565b6101243582526101443560208301908152604051929061177260e08561392a565b6006845260c094853660208701376040519561178f60e08861392a565b6006875236602088013760081c6117a586613c82565b5260a0516117b285613c82565b5260643560ff81169290838103610244576044359360ff8516908186036102445781106113095760a43565ffffffffffff81116112f65760c4359165ffffffffffff83116112e35760e43593677fffffffffffffff85116112d0576101043597677fffffffffffffff89116112bd5750602435801515810361024457156119e6576001905b6084359260ff841684036102445761ff0062ff00006301fe000060209c60b81b9960791b9860491b9760191b9660111b169460101b169260081b161717171717171761188287613ca3565b52600261188e86613ca3565b5251905190604051918383019182526040830152604082526118b160608361392a565b60405191518091835e81019060a05182528060a05192039060025afa1561058d5760a051516118df84613cb3565b5260036118eb83613cb3565b527f1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae27861191684613cc3565b52600461192283613cc3565b5261192c816137c6565b61193583613cd3565b52600661194182613cd3565b527f000000000000000000000000000000000000000000000000000000000000000061196c83613ce3565b52600761197882613ce3565b52805182510361129d5761198c8151613c50565b60a0515b82518110156119ce57806119bd6001600160401b036119b160019487613cf3565b5116610e848388613cf3565b6119c78285613cf3565b5201611990565b60206119de848460a051916146b6565b604051908152f35b60a05190611837565b346102445760a05136600319011261024457602063ffffffff600354821c16604051908152f35b3461024457604036600319011261024457611a2f61371e565b60243560ff198216918215611b895763ffffffff808060035460401c16169160401c1603611b765760a08051839052600160205251604090208054600881901c6001600160a01b03168015611b63573303611b515760ff16611a90816137c6565b8015159081611b3c575b506101f25760066005820154910190815490611ab68282613b47565b4210156107ea578315918215611b27575b8215611b08575b505061155a57817f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba9260209255604051908152a260a05180f35b819250611b1885611b1e93613b47565b92613b47565b10158480611ace565b915042611b348583613b47565b111591611ac7565b60039150611b49816137c6565b141584611a9a565b6282b42960e81b60a05152600460a051fd5b634d36eb6960e01b60a05152600460a051fd5b632299770d60e11b60a05152600460a051fd5b63cbf4a64560e01b60a05152600460a051fd5b61378c565b346102445760a05136600319011261024457602063ffffffff60035416604051908152f35b346102445760a051366003190112610244576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102445760203660031901126102445760ff19611c2861371e565b1660a05152600160205260a0516040902080549060018101611c499061394b565b90600381015460058201546006830154600784015460088501546009860154600a87015490600b88015492600c8901611c81906139a1565b94611c8e600d8b01613a41565b96611c9b60128c01613aa3565b9860178c01549b601801549a6040519e8f9e8f9160ff8116611cbc906137c6565b60ff8116835260081c6001600160a01b031660208084019190915281516040840152015160608201526080015260a08d015260c08c015260e08b01526101008a01526101208901526101408801526101608701526103a06101808701819052611d2891908701906137d0565b906101a08601611d37916137f4565b8481036102a0860152611d499161384a565b91611d5b6102c0850160ff831661389b565b6001600160601b0360a01b8160981b166102e08501528060681c61ffff166103008501528060781c60ff166103208501528060801c61ffff1661034085015260901c60ff1615156103608401526103808301520390f35b3461024457606036600319011261024457611dcb61371e565b6024356001600160401b03811161024457611dea90369060040161372f565b90916044356001600160401b03811161024457611e0b90369060040161372f565b92611e14613e81565b611e1d83613e9f565b9160ff601784015416600381101561022c5761021957825460ff811695909290611e46876137c6565b6002871480156120f9575b6101f257611e5e876137c6565b6001871415806120de575b6107ea57611e778289613ee4565b611e8088613f43565b6003860154036120cb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561024457611f249389936040519586948593849363be98686160e01b855260a051987f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801613b99565b03915afa801561058d576120b2575b5060ff600d83015460081c1694611f4986613c50565b9560a0515b8181106120495750505060ff191660049081178255845191016001600160401b0382116106d957600160401b82116106d9578054828255808310612029575b50602085019060a05152602060a0512060a0515b83811061201557866040875f51602061487d5f395f51905f528860ff1916928392815190611fce816137c6565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e60405180612008339582613d07565b0390a360a0516001815580f35b600190602084519401938184015501611fa1565b612043908260a0515283602060a051209182019101613c23565b85611f8d565b600181901b906001600160ff1b03811681036106c15781600a019182600a116106c157600b6120808460031b87013560c01c6145d6565b91019283106106c15761209d60019360031b86013560c01c6145d6565b60201b176120ab828b613cf3565b5201611f4e565b60a0516120be9161392a565b60a0516102445785611f33565b630b6fac0360e41b60a05152600460a051fd5b506120f26005860154600687015490613b47565b4210611e69565b50612103876137c6565b60048714611e51565b346102445760403660031901126102445761213561212861371e565b612130613e81565b613e9f565b60178101549060ff8216600381101561022c5760020361021957601801547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b15610244576040519263193f942b60e21b84526001600160601b0360a01b9060981b166004840152602483015260243560448301528160648160a0519360a051905af1801561058d576121d75760a0516001815580f35b60a0516121e39161392a565b60a05161024457806101d5565b346102445760a05136600319011261024457602060405160048152f35b346102445760203660031901126102445761222661371e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690811561153c5760209060246040518094819363702574b360e01b835260ff191660048301525afa801561058d5760a05190612292575b602090604051908152f35b506020813d6020116122bd575b816122ac6020938361392a565b810103126102445760209051612287565b3d915061229f565b34610244576040366003190112610244576122de61371e565b6024356001600160401b038111610244578060040160a060031983360301126102445761230a83613e9f565b8054919033600884901c6001600160a01b031603611b5157601281015460029060ff16612336816137c6565b03612572578135600581101561024457600290612352816137c6565b03611657576001600160a01b0361236b60448601613bdd565b166115ca576024840135938415611625576064019261238a8484613bf1565b9050156115b75760ff1661239d816137c6565b801515908161255d575b506101f2576123bf6005820154600683015490613b47565b4210156107ea5783601382015560156123d88484613bf1565b91909201916001600160401b0382116106d9576123f58354613969565b601f8111612529575b5060a05190601f831160011461249357928261247f9896937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed298969361245f9660a05192612488575b50508160011b915f199060031b1c1916179055613bf1565b949060405193849384526040602085015260ff1916956040840191613b79565b0390a260a05180f35b013590508a80612447565b601f198316918460a05152602060a051209260a0515b8181106125115750937f660d494893b9a2e6c617bc5137fb9bac10f3cf87e7c43125657872f2a1959ed298969361245f96936001938361247f9d9b98106124f8575b505050811b019055613bf1565b01355f19600384901b60f8161c191690558a80806124eb565b919360206001819287870135815501950192016124a9565b612557908460a05152602060a05120601f850160051c810191602086106111e157601f0160051c0190613c23565b876123fe565b6003915061256a816137c6565b1415866123a7565b63050b77c760e21b60a05152600460a051fd5b34610244576020366003190112610244576004356001600160a01b038116908181036102445760035460a0805193909352600260209081529251604090205466ffffffffffffff1660589290921b600160581b600160f81b0316600891821c63ffffffff60381b161791909117901b60ff19166040519060ff19168152f35b34610244576020366003190112610244576004356001600160a01b038116908190036102445760a05152600260205260206001600160401b03604060a051205416604051908152f35b346102445760a0513660031901126102445760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b34610244576020366003190112610244576126a261371e565b6040516126ae816138a8565b60a051815260a05160208201526040516126c7816138d8565b60a051815260a0516020820152604082015260a05160608201526060608082015260a05160a082015260a05160c082015260a05160e082015260a05161010082015260a05161012082015260a05161014082015260a0516101608201526060610180820152604051612738816138f3565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260a05160c082015260a05160e08201526101a08201526040516127878161390f565b60a051815260a051602082015260a051604082015260608082015260a05160808201526101c082015260a0516101e082015260a05161020082015260a05161022082015260a05161024082015260a05161026082015260a0516102808201526102a060a05191015260ff191660a051526001602052604060a051206040519061280f826138a8565b805460ff811661281e816137c6565b835260081c6001600160a01b0316602083015261283d6001820161394b565b604083015260038101546060830152600481016040518082602082945493848152019060a05152602060a051209260a0515b818110612b1c5750506128849250038261392a565b6080830152600581015460a0830152600681015460c0830152600781015460e083015260088101546101008301526009810154610120830152600a810154610140830152600b8101546101608301526128df600c82016139a1565b6101808301526128f1600d8201613a41565b6101a083015261290360128201613aa3565b6101c0830152601781015490600360ff8316101561022c5760ff8281601894166101e08601526001600160601b0360a01b8160981b1661020086015261ffff8160681c16610220860152818160781c1661024086015261ffff8160801c1661026086015260901c16151561028084015201546102a0820152604051602081526103e08101918051612993816137c6565b602083015260018060a01b036020820151166040830152602060408201518051606085015201516080830152606081015160a08301526080810151926103c060c084015283518091526020610400840194019060a0515b818110612b06575050506102a0612a93612a67849560a085015160e087015260c085015161010087015260e08501516101208701526101008501516101408701526101208501516101608701526101408501516101808701526101608501516101a0870152610180850151601f19878303016101c08801526137d0565b612a7b6101a08501516101e08701906137f4565b6101c0840151858203601f19016102e087015261384a565b91612aa86101e082015161030086019061389b565b6102008101516001600160a01b03191661032085015261022081015161ffff90811661034086015261024082015160ff166103608601526102608201511661038085015261028081015115156103a085015201516103c08301520390f35b82518652602095860195909201916001016129ea565b845483526001948501948694506020909301920161286f565b346112b05760c03660031901126112b057612b4e61371e565b6080526024356001600160401b0381116112b057612b7090369060040161372f565b906044356001600160401b0381116112b057612b9090369060040161372f565b90916064356001600160401b0381116112b057612bb190369060040161375c565b9190946084356001600160401b0381116112b057612bd390369060040161375c565b9060a4356001600160401b0381116112b057612bf390369060040161375c565b9190612bfd613e81565b612c08608051613e9f565b9660ff885416612c17816137c6565b613276576005880154612c2e60068a015482613b47565b42101561326757421061326757612c458688613ee4565b612c4e87613f43565b9a60038901548c03613258575f5f5b600881106132375750612c7990612c7390613f78565b8a6140de565b60088901549586612c916101508b013560c01c6145d6565b0361322857612ca660988a013560c01c6145d6565b9a612cc08c612cbb60908d013560c01c6145d6565b613b6c565b98612ccb8a8a613b47565b60078d01541061321957612ce66101208c013560c01c6145d6565b9d8e1561320a578e808714908115916131ff575b81156131f4575b506131e5578e496131e5578e60508102046050036131d1578e612d266050820261418d565b90612d34604051928361392a565b60508102808352612d449061418d565b601f19013660208401375f5b8181106131615750505f60208092604051918183925191829101835e8101838152039060025afa15613125575f8051818e5b6008821061313f57505003613130577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156112b0575f928d612e276040519687958694859463be98686160e01b86527f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060048801613b99565b03915afa801561312557613111575b50612e408d613f78565b9560a0515b848110612f1357505060a051988996509450505050505b60088210612ef0575050600b91612e7891846003870155613b47565b92836008820155612e8e60098201958654613b47565b80955501612e9c8154613bcf565b9055604051948552602085015260408401526060830152608082015233907f36c6781d994e030a156d2f6fa11abcc1cd6814482a5d61f90da3f336318b1d2360a060ff196080511692a360a0516001815580f35b909360019085600a0160031b83013560e01c8660051b60e0031b17940190612e5c565b8049156130fb576020606089612f63612f2d858a8a6141a8565b9390846040519586928884019660805160081c8852604085015284840137810160a051838201520301601f19810184528361392a565b60405191518091835e81019060a05182528060a05192039060025afa1561058d57868661301786608086612fd687612fce818e612fc7828f7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff0000000160a05151069d6141c3565b35976141a8565b9390976141a8565b828193604051988997602089019b8d498d5260408a015260608901528688013785019184830160a051815237010160a051815203601f19810183528261392a565b60a0519160a051915190600a5afa3d156130f3573d906130368261418d565b91613044604051938461392a565b825260a0513d90602084013e5b1580156130e7575b6130d05760408180518101031261024457611000604060208301519201519114908115916130a5575b5061308f57600101612e45565b638308e1e960e01b60a05152600452602460a051fd5b7f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001915014155f613082565b50638308e1e960e01b60a05152600452602460a051fd5b50604081511415613059565b606090613051565b634af69d9960e11b60a05152600452602460a051fd5b5f61311b9161392a565b5f60a0528d612e36565b6040513d5f823e3d90fd5b630f2e4cb960e31b5f5260045ffd5b928193600192601c0160031b013560e01c8460051b60e0031b1792018e612d82565b8061316f6030928b8b6141a8565b929092036131c2576030613184828f8e6141a8565b9050036131b357600191605061319c838f8c906141c3565b359160308285028801916020830137015201612d50565b6350320ab160e01b5f5260045ffd5b6338b2168b60e21b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b63b8ff08e960e01b5f5260045ffd5b90508914158f612d01565b858114159150612cfa565b63fdac229f60e01b5f5260045ffd5b6357d18d5360e11b5f5260045ffd5b63341203dd60e21b5f5260045ffd5b906001908260140160031b8b013560e01c8360051b60e0031b179101612c5d565b630b6fac0360e41b5f5260045ffd5b63e843c5eb60e01b5f5260045ffd5b6307a92f1960e51b5f5260045ffd5b346112b0575f3660031901126112b05760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346112b05760403660031901126112b0576132d861371e565b60243560058110156112b05760ff1982169182156134305763ffffffff808060035460401c16169160401c160361342157613312816137c6565b600460ff821611613276575f8281526001602052604090208054600881901c6001600160a01b031680156134125733036134045760ff166133538382613d96565b15613276578284836133755f51602061487d5f395f51905f5296604096613b54565b61337e836137c6565b600183146133a8575b5050825191613395816137c6565b82526133a0816137c6565b6020820152a2005b602060067f45edf61f525089c4937f17d4abc513c0a865c52ff2f704d35bb9a5e207af41ba926005810154804210155f146133fa576133e79042613b6c565b9182915b01558651908152a28486613387565b505f9182916133eb565b6282b42960e81b5f5260045ffd5b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b63cbf4a64560e01b5f5260045ffd5b346112b05760203660031901126112b05760ff1961345b61371e565b165f52600160205260405f2060405190613474826138a8565b805460ff8116613483816137c6565b835260081c6001600160a01b031660208301526134a26001820161394b565b60408301526003810154606083015260048101604051808260208294549384815201905f5260205f20925f5b8181106136015750506134e39250038261392a565b6080830152600581015460a0830190815260068201549060c08401918252600783015460e085015260088301546101008501526009830154610120850152600a830154610140850152600b830154610160850152613543600c84016139a1565b610180850152613555600d8401613a41565b6101a085015261356760128401613aa3565b6101c085015260178301549360ff851660038110156135ed5760186119de9560ff6020986102a0946101e08701526001600160601b0360a01b8160981b1661020087015261ffff8160681c16610220870152818160781c1661024087015261ffff8160801c1661026087015260901c161515610280850152015491015251905190613b47565b634e487b7160e01b5f52602160045260245ffd5b84548352600194850194869450602090930192016134ce565b346112b05760403660031901126112b05761363361371e565b60243560ff1982169182156134305763ffffffff808060035460401c16169160401c1603613421575f8281526001602052604090208054600881901c6001600160a01b031680156134125733036134045760ff16613690816137c6565b8015159081613709575b5061327657811580156136fc575b6136ed57817f36c67c90d9fb754eb7d39c3c925b71dad1c9c05324eaf1fc6976dd7fcff4d2be926007836136e2600e602096015484613d5f565b0155604051908152a2005b632c45be6f60e01b5f5260045ffd5b50600881015482106136a8565b60039150613716816137c6565b14158461369a565b6004359060ff19821682036112b057565b9181601f840112156112b0578235916001600160401b0383116112b057602083818601950101116112b057565b9181601f840112156112b0578235916001600160401b0383116112b0576020808501948460051b0101116112b057565b346112b0575f3660031901126112b05760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b600511156135ed57565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60e0809180511515845260ff602082015116602085015260ff604082015116604085015260ff60608201511660608501526080810151608085015260a081015160a085015260c081015160c08501520151910152565b908151613856816137c6565b81526020820151602082015260018060a01b036040830151166040820152608080613890606085015160a0606086015260a08501906137d0565b930151151591015290565b9060038210156135ed5752565b6102c081019081106001600160401b038211176138c457604052565b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b038211176138c457604052565b61010081019081106001600160401b038211176138c457604052565b60a081019081106001600160401b038211176138c457604052565b90601f801991011681019081106001600160401b038211176138c457604052565b90604051613958816138d8565b602060018294805484520154910152565b90600182811c92168015613997575b602083101461398357565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613978565b9060405191825f8254926139b484613969565b8084529360018116908115613a1f57506001146139db575b506139d99250038361392a565b565b90505f9291925260205f20905f915b818310613a035750509060206139d9928201015f6139cc565b60209193508060019154838589010152019101909184926139ea565b9050602092506139d994915060ff191682840152151560051b8201015f6139cc565b90604051613a4e816138f3565b60e06004829460ff815481811615158652818160081c166020870152818160101c16604087015260181c16606085015260018101546080850152600281015460a0850152600381015460c08501520154910152565b90604051613ab08161390f565b608060ff6004839582815416613ac5816137c6565b85526001810154602086015260028101546001600160a01b03166040860152613af0600382016139a1565b60608601520154161515910152565b346112b0575f3660031901126112b05760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b90600182018092116131d157565b919082018092116131d157565b90613b5e816137c6565b60ff80198354169116179055565b919082039182116131d157565b908060209392818452848401375f828201840152601f01601f1916010190565b94929093613bbe92613bcc979587526020870152608060408701526080860191613b79565b926060818503910152613b79565b90565b5f1981146131d15760010190565b356001600160a01b03811681036112b05790565b903590601e19813603018212156112b057018035906001600160401b0382116112b0576020019181360383136112b057565b818110613c2e575050565b5f8155600101613c23565b6001600160401b0381116138c45760051b60200190565b90613c5a82613c39565b613c67604051918261392a565b8281528092613c78601f1991613c39565b0190602036910137565b805115613c8f5760200190565b634e487b7160e01b5f52603260045260245ffd5b805160011015613c8f5760400190565b805160021015613c8f5760600190565b805160031015613c8f5760800190565b805160041015613c8f5760a00190565b805160051015613c8f5760c00190565b8051821015613c8f5760209160051b010190565b60206040818301928281528451809452019201905f5b818110613d2a5750505090565b8251845260209384019390920191600101613d1d565b3580151581036112b05790565b6040811015613c8f5760051b60240190565b8015613d825764e8d4a510000410613d7357565b63eba5c29b60e01b5f5260045ffd5b634e487b7160e01b5f52601260045260245ffd5b613d9f816137c6565b613da8826137c6565b808214613e5557613db8816137c6565b600281148015613e6e575b8015613e5b575b613e5557613dd7816137c6565b8015613e3657600390613de9816137c6565b14613df357505f90565b613dfc816137c6565b8015908115613e21575b8115613e10575090565b60019150613e1d816137c6565b1490565b9050613e2c816137c6565b6002811490613e06565b50613e40816137c6565b60038114908115613e21578115613e10575090565b50505f90565b50613e65816137c6565b60018114613dca565b50613e78816137c6565b60048114613dc3565b60025f5414613e905760025f55565b633ee5aeb560e01b5f5260045ffd5b60ff1981169081156134305763ffffffff808060035460401c16169160401c1603613421575f52600160205260405f209060018060a01b03825460081c161561341257565b9061020003613f34576001613efc823560c01c6145d6565b1490811591613f1b575b50613f0d57565b6254aabb60e81b5f5260045ffd5b613f2c91506008013560c01c6145d6565b15155f613f06565b630f61e7fd60e21b5f5260045ffd5b5f905f905b60088210613f5557505090565b90916001908360020160031b83013560e01c8460051b60e0031b17920190613f48565b77ffffffffffffffff0000000000000000ffffffffffffffff8160081c9160081b917cff000000ff000000ff000000ff000000ff000000ff000000ff000000ff7dff000000ff000000ff000000ff000000ff000000ff000000ff000000ff007fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff0085167eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff84161760101c941691161760101b9179ffff000000000000ffff000000000000ffff000000000000ffff7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff847dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff84161760201c941691161760201b7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff82821673ffffffff000000000000000000000000ffffffff85161760401b93161760401c16178060801b9060801c1790565b90600360ff6012840154166140f2816137c6565b0361418157601482015460405163650e5fcf60e01b602082015260248082019390935291825261413191906001600160a01b0316610a9a60448361392a565b9015918215614178575b821561416e575b8215614160575b505061415157565b635e32eadd60e01b5f5260045ffd5b600a01541190505f80614149565b4382119250614142565b8115925061413b565b90601301540361415157565b6001600160401b0381116138c457601f01601f191660200190565b90821015613c8f576141bf9160051b810190613bf1565b9091565b9190811015613c8f5760051b0190565b602081519101519080158015614236575b801561421f575b613e55575f51602061489d5f395f51905f52808281930992800981808080848709620292f80960010893620292fc09081490565b505f51602061489d5f395f51905f528210156141eb565b505f51602061489d5f395f51905f528110156141e4565b9291909180158015614309575b614301575f925f5b8281106142d557508184146142cc5761427c906004614623565b92805b61428a575050501490565b5f19019283906004821c6001166142b7576142b1906142aa8385876141c3565b3590614693565b9361427f565b6142b1906142c68385876141c3565b35614693565b50505050505f90565b6142e08184866141c3565b356142ee575b600101614262565b9350600184018085116131d157936142e6565b505050505f90565b506040811161425a565b9060ff600d82015460081c169161432983613c50565b92601783015460ff8160781c169081614429575b50505060048254928160ff85169460ff1916178155018351906001600160401b0382116138c457600160401b82116138c457805482825580831061440d575b5060208501905f5260205f205f5b8381106143f95750505050905f51602061487d5f395f51905f5260409260ff19169283928151906143ba816137c6565b815260046020820152a27fdf1be195647bf0f039490311aa7fd2242eb64a0eb3844c37f174b8d7c25d448e604051806143f4339582613d07565b0390a3565b60019060208451940193818401550161438a565b61442390825f528360205f209182019101613c23565b5f61437c565b6018850154604051630222162f60e21b8152609883901b6001600160a01b03191660048201526024810191909152606882901c61ffff16604482015260648101929092525f826084817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215613125575f905f9361451b575b501561450c575f919060801c61ffff16825b8481106144ce57505061433d565b600182821c16156144e2575b6001016144c0565b92614504816144f360019386613cf3565b516144fe878c613cf3565b52613bcf565b9390506144da565b630e0d4dc560e41b5f5260045ffd5b9250503d805f843e61452d818461392a565b8201916040818403126112b05780519081151582036112b0576020810151906001600160401b0382116112b057019280601f850112156112b057835161457281613c39565b94614580604051968761392a565b81865260208087019260051b8201019283116112b057602001905b8282106145ab575050505f6144ae565b815181526020918201910161459b565b6020915f91838251920190620186a0fa601f3d1116905f5190565b66ff00ff00ff00ff67ff00ff00ff00ff008260081b169160081c161765ffff0000ffff67ffff0000ffff00008260101b169160101c161767ffffffff000000008160201b169060201c1790565b60209161464b5f926146416001600160401b038060c01b92166145d6565b60c01b1691613f78565b604051908482019283526028820152600160f81b60488201526029815261467360498261392a565b604051918291518091835e8101838152039060025afa15613125575f5190565b5f906020926040519084820192835260408201526040815261467360608261392a565b9081518015614301576001811461486b576040841461485c575f915f5b82811061483357506146ed6146e88484613b6c565b613c50565b916146fb6146e88583613b6c565b9161470e61470886613c50565b95613c50565b955f9081805b888a8c888710614753579550505050505061473c949392506147369150613b39565b916146b6565b9161473661474a9394613b39565b613bcc91614693565b90614798600161478b8a959461477961476c8c8c613cf3565b516001600160401b031690565b906001600160401b03809216901c1690565b166001600160401b031690565b6147f1575050506147e86001916147e36147b561476c8888613cf3565b6147bf888a613cf3565b516147ca848d613cf3565b526147d5838d613cf3565b906001600160401b03169052565b613bcf565b935b0192614714565b61482d926147d58680946147e3946148278c9a6148208c60019c9f8f61481a9161476c91613cf3565b98613cf3565b5192613cf3565b52613cf3565b916147ea565b9261485560019161484f8361478b8a61477961476c8b8d613cf3565b90613b47565b93016146d3565b63394fd24160e21b5f5260045ffd5b5090506148789150613c82565b519056fe56f95be551d4235ff95edcee7dca6f56f66968ed1b176f73ffd899721aa19abf30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220f4ee6aca5c46b3e0d2ad3bfe804ae8fd251b0e813daa5c80865ce7737d0c50d864736f6c634300081c003360e0806040523461011157602061002e60049261106880380380916100248285610115565b833981019061014c565b336080526001600160a01b031660a081905260405163ebe86c1360e01b815292839182905afa908115610106575f916100d7575b506001600160a01b031660c052604051610efc908161016c82396080518181816101530152818161037701528181610756015281816107cd0152610bce015260a051818181610186015281816108b701528181610a2f0152610c43015260c05181818160c30152818161047001526107fd0152f35b6100f9915060203d6020116100ff575b6100f18183610115565b81019061014c565b5f610062565b503d6100e7565b6040513d5f823e3d90fd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761013857604052565b634e487b7160e01b5f52604160045260245ffd5b9081602091031261011157516001600160a01b0381168103610111579056fe60806040526004361015610011575f80fd5b5f5f3560e01c8063088858bc146108e6578063481c6a75146108a257806364fe50ac146107b1578063702574b3146107855780637b10399914610740578063883a429a1461034e578063ae977ede146100f2578063ebe86c13146100ad5763f08c9b3c1461007d575f80fd5b346100aa57806003193601126100aa576020610097610c34565b6040516001600160a01b03199091168152f35b80fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760603660031901126100aa5761010c610977565b9060243591604435916001600160401b0383116100aa57366023840112156100aa578260040135936001600160401b03851161034a573660248660071b8601011161034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b579093927f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316928592916001600160a01b031990911690835b8688101561032c578760071b82019060845f516020610ea75f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d2899602485013509925f516020610ea75f395f51905f527f22545b22db5abade8bd584e7fc9b46e5b38df17e479acf79d76612d2174d289960648301350960405194633d98dab360e11b865287600487015288602487015260448601526044820135606486015282850152013560a483015260208260c481898b5af19182156103215786926102e2575b508861029c5750600190975b01966101c1565b979061ffff8916908282018092116102ce5761ffff16036102bf57600190610295565b6357149e2560e01b8552600485fd5b634e487b7160e01b87526011600452602487fd5b9091506020813d8211610319575b816102fd6020938361099f565b810103126103155761030e906109eb565b905f610289565b8580fd5b3d91506102f0565b6040513d88823e3d90fd5b60209061ffff60405191168152f35b633217675b60e21b8252600482fd5b5080fd5b50346100aa576101003660031901126100aa5761036961098e565b60e036602319011261034a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361033b576103ae90610bab565b602435600381101561073c5760021490811561072e576044356001600160a01b03198116810361072a57915b6040908151906103ea838361099f565b6001825260208201601f198401368237825115610716573090521561070e5784935b82519461010086018681106001600160401b038211176106fa57845260028110156106e6578552602085019486865283810192835260608101601081526080820188815260a0830189815260c08401918a835260e08501938b855260018060a01b037f00000000000000000000000000000000000000000000000000000000000000001697883b156106e257895197631bb64f5960e21b89526001600160601b0360a01b169b8c60048a01528b60248a015261010060448a01526102048901975160028110156106ce5791899693918f989593610104899b989b01525115156101248801525194610100610144880152855180915260206102248801960190885b81811061069d57505050916001600160401b03869798818096959461ffff829651166101648b0152511661018489015251166101a487015251166101c485015251166101e48301526064356064830152608435608483015260a43560a483015260c43560c483015260e43560e4830152038183865af180156106935761067a575b5093816044958151968780926303e95d1360e21b82528860048301528760248301525afa92831561066e57608095829461060b575b505f516020610ea75f395f51905f52917f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e0097542918451968752602087015209908301526060820152f35b7f043a24d9a1c954e75f55ef19c539fa43c22fa6626c121e9a21ec5cd3e00975429194505f516020610ea75f395f51905f52925061065e90843d8611610667575b610656818361099f565b810190610c1e565b949150916105c3565b503d61064c565b509051903d90823e3d90fd5b61068586809261099f565b61068f575f61058e565b8480fd5b83513d88823e3d90fd5b92949750929497509794602080600192838060a01b038c511681520199019101908e9794928a97949299969961050d565b634e487b7160e01b8f52602160045260248ffd5b8c80fd5b634e487b7160e01b87526021600452602487fd5b634e487b7160e01b88526041600452602488fd5b60019361040c565b634e487b7160e01b87526032600452602487fd5b8380fd5b610736610c34565b916103da565b8280fd5b50346100aa57806003193601126100aa576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346100aa5760203660031901126100aa5760206107a96107a461098e565b610bab565b604051908152f35b503461088f57606036600319011261088f576107cb610977565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610893577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561088f575f9160648392604051948593849263a59b7a4d60e01b84526001600160601b0360a01b166004840152602435602484015260443560448401525af1801561088457610876575080f35b61088291505f9061099f565b005b6040513d5f823e3d90fd5b5f80fd5b633217675b60e21b5f5260045ffd5b3461088f575f36600319011261088f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461088f57608036600319011261088f576108ff610977565b60443561ffff8116810361088f576064359061ffff8216820361088f5761092992602435906109fa565b906040519182916040830190151583526040602084015281518091526020606084019201905f5b81811061095e575050500390f35b8251845285945060209384019390920191600101610950565b600435906001600160a01b03198216820361088f57565b6004359060ff198216820361088f57565b90601f801991011681019081106001600160401b038211176109c057604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b0381116109c05760051b60200190565b519061ffff8216820361088f57565b9261ffff1693610a09856109d4565b610a16604051918261099f565b858152601f19610a25876109d4565b01366020830137937f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316925f5b61ffff811688811015610b9e578061ffff88160161ffff8111610b8a57604051639bbada6760e01b81526001600160a01b0319861660048201526024810185905261ffff9190911660448201526060816064818a5afa908115610884575f91610b14575b50602081015115610b0657604001519088511115610af257600582901b621fffe01688016020015260010161ffff16610a5a565b634e487b7160e01b5f52603260045260245ffd5b505f98509395505050505050565b90506060813d8211610b82575b81610b2e6060938361099f565b8101031261088f5760405190606082018281106001600160401b038211176109c057604052610b5c816109eb565b8252602081015190811515820361088f576040916020840152015160408201525f610abe565b3d9150610b21565b634e487b7160e01b5f52601160045260245ffd5b5050505093505050600191565b5f516020610ea75f395f51905f5290604051602081019146835260018060a01b037f000000000000000000000000000000000000000000000000000000000000000016604083015260ff1916606082015260608152610c0b60808261099f565b519020068015610c185790565b50600190565b919082604091031261088f576020825192015190565b60405163a4adcd7f60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316602082600481845afa918215610884575f92610e62575b506040516323488be560e01b8152602081600481855afa8015610884575f90610e1c575b63ffffffff60401b915060401b165f5b6001600160401b0381166008811080610e0a575b15610dfb576001600160401b038516036001600160401b038111610b8a576001600160401b036001600160601b0360a01b9116831760a01b1660405163d397925360e01b8152816004820152602081602481885afa8015610884575f90610dbf575b60ff9150166010811015610daf57604051906356cbb5f360e01b82528260048301526024820152604081604481885afa9081610d92575b50610d8a57506001600160401b03905b166001600160401b038114610b8a57600101610cb6565b935050505090565b610da99060403d811161066757610656818361099f565b50610d63565b50506001600160401b0390610d73565b506020813d8211610df3575b81610dd86020938361099f565b8101031261088f575160ff8116810361088f5760ff90610d2c565b3d9150610dcb565b63081ea97160e31b5f5260045ffd5b50806001600160401b03861611610cca565b506020813d602011610e5a575b81610e366020938361099f565b8101031261088f575163ffffffff8116810361088f5763ffffffff60401b90610ca6565b3d9150610e29565b9091506020813d602011610e9e575b81610e7e6020938361099f565b8101031261088f57516001600160401b038116810361088f57905f610c82565b3d9150610e7156fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122083f64bd087e4e1038f23f14b7f0e980336ced0b7f92e88054f09c8a9f6f64ac864736f6c634300081c0033",
}

// ProcessRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProcessRegistryMetaData.ABI instead.
var ProcessRegistryABI = ProcessRegistryMetaData.ABI

// ProcessRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProcessRegistryMetaData.Bin instead.
var ProcessRegistryBin = ProcessRegistryMetaData.Bin

// DeployProcessRegistry deploys a new Ethereum contract, binding an instance of ProcessRegistry to it.
func DeployProcessRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID uint32, _ziskVerifier common.Address, _batchProgramVK [32]byte, _resultsProgramVK [32]byte, _rootCVadcopFinal [32]byte, _ballotVKHash [32]byte, _dkgManager common.Address) (common.Address, *types.Transaction, *ProcessRegistry, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProcessRegistryBin), backend, _chainID, _ziskVerifier, _batchProgramVK, _resultsProgramVK, _rootCVadcopFinal, _ballotVKHash, _dkgManager)
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),bytes32,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),uint8,bytes12,uint16,uint8,uint16,bool,bytes32))
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
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),bytes32,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),uint8,bytes12,uint16,uint8,uint16,bool,bytes32))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [31]byte) (DAVINCITypesProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x59d821c4.
//
// Solidity: function getProcess(bytes31 processId) view returns((uint8,address,(uint256,uint256),bytes32,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,address,string,bool),uint8,bytes12,uint16,uint8,uint16,bool,bytes32))
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
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, bytes32 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, uint8 keyMode, bytes12 dkgEpochId, uint16 dkgFirstIndex, uint8 dkgCount, uint16 dkgZeroSkipped, bool dkgResultsRequested, bytes32 dkgAid)
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
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
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
		BallotMode            DAVINCITypesBallotMode
		Census                DAVINCITypesCensus
		KeyMode               uint8
		DkgEpochId            [12]byte
		DkgFirstIndex         uint16
		DkgCount              uint8
		DkgZeroSkipped        uint16
		DkgResultsRequested   bool
		DkgAid                [32]byte
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
	outstruct.BallotMode = *abi.ConvertType(out[12], new(DAVINCITypesBallotMode)).(*DAVINCITypesBallotMode)
	outstruct.Census = *abi.ConvertType(out[13], new(DAVINCITypesCensus)).(*DAVINCITypesCensus)
	outstruct.KeyMode = *abi.ConvertType(out[14], new(uint8)).(*uint8)
	outstruct.DkgEpochId = *abi.ConvertType(out[15], new([12]byte)).(*[12]byte)
	outstruct.DkgFirstIndex = *abi.ConvertType(out[16], new(uint16)).(*uint16)
	outstruct.DkgCount = *abi.ConvertType(out[17], new(uint8)).(*uint8)
	outstruct.DkgZeroSkipped = *abi.ConvertType(out[18], new(uint16)).(*uint16)
	outstruct.DkgResultsRequested = *abi.ConvertType(out[19], new(bool)).(*bool)
	outstruct.DkgAid = *abi.ConvertType(out[20], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, bytes32 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, uint8 keyMode, bytes12 dkgEpochId, uint16 dkgFirstIndex, uint8 dkgCount, uint16 dkgZeroSkipped, bool dkgResultsRequested, bytes32 dkgAid)
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
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x784df74a.
//
// Solidity: function processes(bytes31 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, bytes32 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, uint8 keyMode, bytes12 dkgEpochId, uint16 dkgFirstIndex, uint8 dkgCount, uint16 dkgZeroSkipped, bool dkgResultsRequested, bytes32 dkgAid)
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
	BallotMode            DAVINCITypesBallotMode
	Census                DAVINCITypesCensus
	KeyMode               uint8
	DkgEpochId            [12]byte
	DkgFirstIndex         uint16
	DkgCount              uint8
	DkgZeroSkipped        uint16
	DkgResultsRequested   bool
	DkgAid                [32]byte
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

// NewProcess is a paid mutator transaction binding the contract method 0xd8738c81.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (uint8,bytes12,uint256,uint256,uint256,uint256,uint256) dkg) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, dkg DAVINCITypesDKGParams) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, dkg)
}

// NewProcess is a paid mutator transaction binding the contract method 0xd8738c81.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (uint8,bytes12,uint256,uint256,uint256,uint256,uint256) dkg) returns(bytes31)
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, dkg DAVINCITypesDKGParams) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, dkg)
}

// NewProcess is a paid mutator transaction binding the contract method 0xd8738c81.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,uint8,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,address,string,bool) census, string metadata, (uint256,uint256) encryptionKey, (uint8,bytes12,uint256,uint256,uint256,uint256,uint256) dkg) returns(bytes31)
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode DAVINCITypesBallotMode, census DAVINCITypesCensus, metadata string, encryptionKey DAVINCITypesEncryptionKey, dkg DAVINCITypesDKGParams) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, dkg)
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
