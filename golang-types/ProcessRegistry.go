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
	CostFromWeight bool
	UniqueValues   bool
	NumFields      uint8
	CostExponent   uint8
	MaxValue       *big.Int
	MinValue       *big.Int
	MaxValueSum    *big.Int
	MinValueSum    *big.Int
}

// IProcessRegistryCensus is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryCensus struct {
	CensusOrigin             uint8
	CensusRoot               [32]byte
	CensusURI                string
	OnchainAllowAnyValidRoot bool
}

// IProcessRegistryEncryptionKey is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryEncryptionKey struct {
	X *big.Int
	Y *big.Int
}

// IProcessRegistryProcess is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryProcess struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         IProcessRegistryEncryptionKey
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
	BallotMode            IProcessRegistryBallotMode
	Census                IProcessRegistryCensus
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_stVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rVerifier\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_blobsDA\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BlobNotFoundInTx\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CensusNotUpdatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"limbIndex\",\"type\":\"uint8\"}],\"name\":\"InvalidBlobCommitmentLimb\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCensusURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxMinValueBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValueSum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxVoters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinTotalCost\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProcessId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniqueValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValueSumBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxVotersReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownProcessIdPrefix\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"}],\"name\":\"ProcessMaxVotersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotersCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOverwrittenVotesCount\",\"type\":\"uint256\"}],\"name\":\"ProcessStateTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOB_INDEX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CENSUS_ORIGIN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobsDA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"}],\"name\":\"getNextProcessId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structIProcessRegistry.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"}],\"name\":\"getProcessEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"initStateRoot\",\"type\":\"uint256\"}],\"name\":\"newProcess\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pidPrefix\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVoters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overwrittenVotesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"numFields\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValueSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValueSum\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"onchainAllowAnyValidRoot\",\"type\":\"bool\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxVoters\",\"type\":\"uint256\"}],\"name\":\"setProcessMaxVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461011e576080816126d7803803809161001f8285610122565b83398101031261011e57805163ffffffff8116810361011e5761004460208301610159565b606061005260408501610159565b9301519182151580930361011e5760028054600354600160201b600160e01b0319909116604094851b600160401b600160e01b031617602084811b67ffffffff000000001691909117909255925160e09290921b6001600160e01b0319169082019081523060601b6024830152601882526100ce603883610122565b9051902063ffffffff60a01b60a09190911b166001600160a01b03939093166001600160c81b0319919091161760ff60c01b60c09290921b919091161717600355604051612569908161016e8239f35b5f80fd5b601f909101601f19168101906001600160401b0382119082101761014557604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361011e5756fe60a0806040526004361015610012575f80fd5b5f905f3560e01c90816302cc919a14611ec3575080630535fece14611da45780630970639214611d7a57806326795bc314611d605780632c645ddb1461168e57806344368408146113895780634c0acc561461133957806356a6f1e2146111e857806362fa11fc146111a357806368141f2c1461113557806372c628ef1461111a57806374e0782614611026578063848df54014611003578063992bc45b14610d275780639c18deaf14610cff578063a5b110de14610986578063a5f01e1914610835578063adc879e91461080f578063c718c01f146106a3578063cddf08bc1461067c578063cdf497c814610656578063d71d072e146101b05763f9aa44991461011b575f80fd5b346101ad57806003193601126101ad5760035460405163233ace1160e01b81529190602090839060049082906001600160a01b03165afa9081156101a1579061016a575b602090604051908152f35b506020813d602011610199575b8161018460209383611f4a565b81010312610195576020905161015f565b5f80fd5b3d9150610177565b604051903d90823e3d90fd5b80fd5b50346101ad576101bf3661225a565b939291908315610647576003549463ffffffff808760a01c161663ffffffff8660401c160361063857848752866020526040872093845460018060a01b038160081c16156106295760ff16610213816120c5565b61061a5761022a60058601546006870154906122af565b42101561060b57610239612366565b5081830192610100818503126106075783601f8201121561060757604051936101006102658187611f4a565b859083019182116106035782905b8282106105f3575050508351946020850151946040810151966060820151968c8a600360ff601260808801519360a08901519660e060c08b01519a0151986040516080526102c2608051611f2e565b60805152602060805101526040608051019e8f526060608051019d8e52608080510194855260a06080510196875260c06080510198895260e06080510197885201541661030e816120c5565b14610533575b5050608051519860038b01998a54036105245751885161033391612359565b9b8c151580610512575b610503579060ff8e9897969594939260c01c16610450575b505060025460401c6001600160a01b0316919050813b1561044157859361039260405196879586948594635c73957b60e11b8652600486016123a2565b03915afa80156104455761042c575b5050600b916020608051015190556103be600884019586546122af565b80955551916103d2600982019384546122af565b809355016103e081546123cc565b90556080515192602060805101516040519485526020850152604084015260608301527f43b1a882c73bc84ae667608d3a61baececd99134224962eba9d13eb63c595ae160803393a380f35b8161043691611f4a565b61044157855f6103a1565b8580fd5b6040513d84823e3d90fd5b90919293949596505191519051906040519261046d606085611f4a565b60308452602084019060403683378060801c6104ef5760801b90528060801c6104db5760801b60308301528060801c6104c7576104ba8c96959493926104bf9260801b60408201526124c5565b612502565b5f8080610355565b6316416a3d60e01b8c52600260045260248cfd5b6316416a3d60e01b8d52600160045260248dfd5b6316416a3d60e01b8f5260048f905260248ffd5b6357d18d5360e11b8e5260048efd5b5060088b015460078c0154111561033d565b630b6fac0360e41b8e5260048efd5b60138c0154905160405163650e5fcf60e01b8152600481019190915290602090829060249082906001600160a01b03165afa9182156101a157916105bd575b5060ff60158c0154161590816105af575b81156105a5575b50610596578c5f610314565b635e32eadd60e01b8d5260048dfd5b905043105f61058a565b600a8c015481109150610583565b90506020813d6020116105eb575b816105d860209383611f4a565b810103126105e757515f610572565b8d80fd5b3d91506105cb565b8135815260209182019101610273565b8a80fd5b8880fd5b63e843c5eb60e01b8852600488fd5b6307a92f1960e51b8852600488fd5b634d36eb6960e01b8952600489fd5b632299770d60e11b8752600487fd5b63cbf4a64560e01b8652600486fd5b50346101ad57806003193601126101ad57602060ff60035460c01c166040519015158152f35b50346101ad57806003193601126101ad57602063ffffffff60035460a01c16604051908152f35b50346101ad576106b236612211565b81156108005763ffffffff8060035460a01c161663ffffffff8360401c16036107f1578183526020839052604083208054600881901c6001600160a01b031680156107e25733036107d45760ff16610709816120c5565b80151590816107bf575b506107b05760066005820154910190815490831591821561079b575b821561077c575b505061076d57817f0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f469260209255604051908152a280f35b637616640160e01b8452600484fd5b81925061078c85610792936122af565b926122af565b10155f80610736565b91506107a784826122af565b4210159161072f565b6307a92f1960e51b8452600484fd5b600391506107cc816120c5565b14155f610713565b6282b42960e81b8552600485fd5b634d36eb6960e01b8652600486fd5b632299770d60e11b8352600483fd5b63cbf4a64560e01b8352600483fd5b50346101ad57806003193601126101ad57602063ffffffff600254821c16604051908152f35b50346101ad5760203660031901126101ad5760043581528060205260408120906040519061086282611edc565b825460ff8116610871816120c5565b835260081c6001600160a01b0316602083015261089060018401611f6b565b60408301526003830154606083015260048301604051918260208354918281520192825260208220915b8181106109705760206109688888886108d5818a0382611f4a565b608082015260058201549060a081019182526101c061095c601260068601549560c08501968752600781015460e086015260088101546101008601526009810154610120860152600a810154610140860152600b81015461016086015261093e600c8201611fc1565b610180860152610950600d8201612061565b6101a0860152016120e3565b910152519051906122af565b604051908152f35b82548452602090930192600192830192016108ba565b5034610195576109953661225a565b8491929495939515610cf0576003549063ffffffff808360a01c161663ffffffff8460401c1603610ce157825f525f60205260405f209687549660018060a01b038860081c1615610cd25760ff8816936109ee856120c5565b600285148015610cbf575b610cb057610a06856120c5565b600185141580610c95575b610c86576001600160a01b031691823b156101955783925f928892610a4c60405196879586948594635c73957b60e11b8652600486016123a2565b03915afa8015610c7b57610c66575b50830161012084820312610c625780601f85011215610c62576101209060405194610a868387611f4a565b8590838101928311610c5e57905b828210610c4e575050508351600388015403610c3f5760405196610ab88289611f4a565b600888526020880191601f190136833760015b6009811015610b28578060051b8601515f198201828111610b14578a51811015610b005760051b8a0160200152600101610acb565b634e487b7160e01b89526032600452602489fd5b634e487b7160e01b89526011600452602489fd5b50879450600490818860ff1916178155018451916001600160401b038311610c2b57680100000000000000008311610c2b578154838355808410610c10575b50908493929190875260208720875b838110610bf9575050505060407f2f25f350b681441b5bb45430fbe61e78d0170cff4ad02cb0aa48a6427909aa8c91815190610bb1816120c5565b815260046020820152a27f5da8009aeb11de7ef64598b0de897a03d2f93efabda8507a70e29cca4530cf5d6040516020815280610bf333956020830190612227565b0390a380f35b825181830155869550602090920191600101610b76565b82885260208820610c259181019085016122d4565b87610b67565b634e487b7160e01b87526041600452602487fd5b630b6fac0360e41b8552600485fd5b8135815260209182019101610a94565b8780fd5b8480fd5b610c739195505f90611f4a565b5f935f610a5b565b6040513d5f823e3d90fd5b63e843c5eb60e01b5f5260045ffd5b50610ca960058b015460068c0154906122af565b4210610a11565b6307a92f1960e51b5f5260045ffd5b50610cc9856120c5565b600485146109f9565b634d36eb6960e01b5f5260045ffd5b632299770d60e11b5f5260045ffd5b63cbf4a64560e01b5f5260045ffd5b34610195575f366003190112610195576003546040516001600160a01b039091168152602090f35b3461019557602036600319011261019557604051610d4481611edc565b5f81525f6020820152604051610d5981611ef8565b5f81525f602082015260408201525f6060820152606060808201525f60a08201525f60c08201525f60e08201525f6101008201525f6101208201525f6101408201525f6101608201526060610180820152610db2612366565b6101a08201526101c060405191610dc883611f13565b5f83525f6020840152606060408401525f606084015201526004355f525f60205260405f2060405190610dfa82611edc565b80549060ff8216610e0a816120c5565b835260089190911c6001600160a01b03166020830190815290610e2f60018201611f6b565b9260408101938452600382015460608201908152600483019060405180928391602082549182815201915f5260205f20905f5b818110610fea5750505003610e779083611f4a565b6080830191825260058401549160a08401928352600685015460c08501908152600786015460e08601908152600887015461010087019081526009880154916101208801928352600a890154936101408901948552600b8a0154956101608a01968752600c8b01610ee790611fc1565b6101808b0190815297610efc600d8d01612061565b9b6101a08c019c8d52601201610f11906120e3565b9a6101c081019b8c526040519e8f9e8f926020845251610f30816120c5565b6020840152600160a01b60019003905116604083015251805160608301526020015190608001525160a08d01525160c08c016102e090526103008c01610f7591612227565b975160e08c0152516101008b0152516101208a0152516101408901525161016088015251610180870152516101a086015251848203601f19016101c0860152610fbe919061212c565b91516101e08401610fce91612150565b51828203601f19016102e0840152610fe691906121a5565b0390f35b8254845286945060209093019260019283019201610e62565b34610195575f36600319011261019557602063ffffffff60025416604051908152f35b346101955761103436612211565b8115610cf05763ffffffff8060035460a01c161663ffffffff8360401c1603610ce1575f8281526020819052604090208054600881901c6001600160a01b03168015610cd257330361110c5760ff1661108c816120c5565b80151590816110f7575b50610cb057811580156110ea575b6110db57817fdbed30289c8edb9c8e40449ce45792cee5b9eba7c798afbadbcb4345699f1e249260076020930155604051908152a2005b632c45be6f60e01b5f5260045ffd5b50600881015482106110a4565b60039150611104816120c5565b141584611096565b6282b42960e81b5f5260045ffd5b34610195575f36600319011261019557602060405160048152f35b34610195576020366003190112610195576004356001600160a01b03811690818103610195576003545f92835260016020908152604090932054606092831b6bffffffffffffffffffffffff19169190921c63ffffffff60401b16176001600160401b039190911617610968565b34610195576020366003190112610195576004356001600160a01b03811690819003610195575f52600160205260206001600160401b0360405f205416604051908152f35b34610195576040366003190112610195576004356024356005811015610195578115610cf05763ffffffff8060035460a01c161663ffffffff8360401c1603610ce157611234816120c5565b600460ff821611610cb057815f525f60205260405f2090815460018060a01b038160081c168015610cd257330361110c5760ff169061127381836123da565b15610cb05783836112a5837f2f25f350b681441b5bb45430fbe61e78d0170cff4ad02cb0aa48a6427909aa8c966122bc565b6112ae836120c5565b600183146112dc575b5050604051916112c6816120c5565b82526112d1816120c5565b6020820152604090a2005b602060067f0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46926005810154804210155f1461132f5761131b9042612359565b9182915b0155604051908152a283856112b7565b505f91829161131f565b34610195575f366003190112610195576002546040805163233ace1160e01b81529160209183916004918391901c6001600160a01b03165afa8015610c7b575f9061016a57602090604051908152f35b34610195576040366003190112610195576004356024356001600160401b038111610195578060040160806003198336030112610195578215610cf05763ffffffff8060035460a01c161663ffffffff8460401c1603610ce1575f83815260208190526040902080549190600883901c6001600160a01b03168015610cd257330361110c5760ff60128201541661141f816120c5565b6002810361167f5782359060058210156101955761143c826120c5565b611445816120c5565b0361167057602484013593841561166157604401926114648484612307565b9050156116525760ff16611477816120c5565b801515908161163d575b50610cb05761149960058201546006830154906122af565b421015610c865783601382015560146114b28484612307565b91909201916001600160401b038211611629576114cf8354611f89565b601f81116115ee575b505f90601f831160011461155f57928261154f96937f87e1c82bc035c552e997ec2f1cf023790ced0cc01098d516dc6e7200835ffb27989693611534965f92611554575b50508160011b915f199060031b1c1916179055612307565b92906040519384938452604060208501526040840191612339565b0390a2005b013590508a8061151c565b601f19831691845f5260205f20925f5b8181106115d65750937f87e1c82bc035c552e997ec2f1cf023790ced0cc01098d516dc6e7200835ffb2798969361153496936001938361154f9b98106115bd575b505050811b019055612307565b01355f19600384901b60f8161c191690558a80806115b0565b9193602060018192878701358155019501920161156f565b61161990845f5260205f20601f850160051c8101916020861061161f575b601f0160051c01906122d4565b876114d8565b909150819061160c565b634e487b7160e01b5f52604160045260245ffd5b6003915061164a816120c5565b141586611481565b630f8b932160e11b5f5260045ffd5b635e32eadd60e01b5f5260045ffd5b63f37f7b5d60e01b5f5260045ffd5b63050b77c760e21b5f5260045ffd5b3461019557610220366003190112610195576005600435101561019557602435610100366083190112610195576001600160401b0361018435116101955760806101843536036003190112610195576101a4356001600160401b038111610195576116fd9036906004016121e4565b90916040366101c319011261019557600354335f81815260016020526040902054929384939092606092831b6bffffffffffffffffffffffff1916921c63ffffffff60401b16919091176001600160401b03919091161791825f525f6020523360018060a01b0360405f205460081c1614611d515760ff61177c6122f7565b16158015611d3d575b611d2e5761ffff6101443511611d1f57610104356101243511611d1057610144356101643511611d0157606435156110db57600561018435600401351015610195576117d761018435600401356120c5565b600560ff61018435600401351611611670576117f961018435600401356120c5565b60036101843560040135141580611ce7575b611cd857602461018435013515611661576118326044610184350161018435600401612307565b905015611652576118446004356120c5565b600460ff813516118015611ca9575b610cb05715611ca1575b428310611c925742611871604435856122af565b1115611c8357815f525f60205260405f209261188f600435856122bc565b6005840155604435600684015560643560078401558254610100600160a81b0319163360081b610100600160a81b03161783556101c43560018401556101e4356002840155610204356003840155600c8301906001600160401b038111611629576118fa8254611f89565b601f8111611c53575b505f601f8211600114611bf05781929394955f92611be5575b50508160011b915f199060031b1c19161790555b600d8201608435801515810361019557815490151560ff1660ff199190911617815560a435801515908181036101955750815462ff000061196f6122f7565b60101b169060e43560ff811681036101955763ff0000008161ff00925060181b169360081b169063ffffff001916171717905561010435600e83015561012435600f830155610144356010830155610164356011830155601282016119da61018435600401356120c5565b60ff61018435600401351660ff19825416179055602461018435013560138301556014820191611a166044610184350161018435600401612307565b6001600160401b03819592951161162957611a318254611f89565b601f8111611bb5575b505f601f8211600114611b525781929394955f92611b47575b50508160011b915f199060031b1c19161790555b611a8e611a79606461018435016122ea565b601583019060ff801983541691151516179055565b600a4391015560025463ffffffff811663ffffffff8114611b3357600163ffffffff9101169063ffffffff191617600255335f52600160205260405f20908154916001600160401b038316926001600160401b038414611b33576001600160401b0360016020950116906001600160401b0319161790556040519033817fada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa5f80a38152f35b634e487b7160e01b5f52601160045260245ffd5b013590508580611a53565b601f19821695835f5260205f20915f5b888110611b9d57508360019596979810611b84575b505050811b019055611a67565b01355f19600384901b60f8161c19169055858080611b77565b90926020600181928686013581550194019101611b62565b611bdf90835f5260205f20601f840160051c8101916020851061161f57601f0160051c01906122d4565b85611a3a565b01359050858061191c565b601f19821695835f5260205f20915f5b888110611c3b57508360019596979810611c22575b505050811b019055611930565b01355f19600384901b60f8161c19169055858080611c15565b90926020600181928686013581550194019101611c00565b611c7d90835f5260205f20601f840160051c8101916020851061161f57601f0160051c01906122d4565b85611903565b637616640160e01b5f5260045ffd5b632ca4094f60e21b5f5260045ffd5b42925061185d565b50611cb56004356120c5565b600435151580156118535750611ccc6004356120c5565b60036004351415611853565b63f545b7bf60e01b5f5260045ffd5b506001611cf9606461018435016122ea565b15151461180b565b636d403ec760e11b5f5260045ffd5b63207ea56d60e01b5f5260045ffd5b634dba795160e01b5f5260045ffd5b63ac38930b60e01b5f5260045ffd5b50600860ff611d4a6122f7565b1611611785565b635040c4d360e11b5f5260045ffd5b34610195575f3660031901126101955760206040515f8152f35b34610195575f366003190112610195576002546040805191901c6001600160a01b03168152602090f35b34610195576020366003190112610195576004355f525f60205260405f2080549060ff82169060018101611dd790611f6b565b9060038101546005820154600683015460078401546008850154600986015491600a87015493600b88015495600c8901611e1090611fc1565b97611e1d600d8b01612061565b99601201611e2a906120e3565b9a604051809e8e829f611e3c906120c5565b825260081c6001600160a01b0316602091820152815160408f0152015160608d015260808c015260a08b015260c08a015260e08901526101008801526101208701526101408601526101608501526102c06101808501819052611ea2919085019061212c565b906101a08401611eb191612150565b8281036102a0840152610fe6916121a5565b34610195575f3660031901126101955780600560209252f35b6101e081019081106001600160401b0382111761162957604052565b604081019081106001600160401b0382111761162957604052565b608081019081106001600160401b0382111761162957604052565b61010081019081106001600160401b0382111761162957604052565b90601f801991011681019081106001600160401b0382111761162957604052565b90604051611f7881611ef8565b602060018294805484520154910152565b90600182811c92168015611fb7575b6020831014611fa357565b634e487b7160e01b5f52602260045260245ffd5b91607f1691611f98565b9060405191825f825492611fd484611f89565b808452936001811690811561203f5750600114611ffb575b50611ff992500383611f4a565b565b90505f9291925260205f20905f915b818310612023575050906020611ff9928201015f611fec565b602091935080600191548385890101520191019091849261200a565b905060209250611ff994915060ff191682840152151560051b8201015f611fec565b9060405161206e81611f2e565b60e06004829460ff815481811615158652818160081c1615156020870152818160101c16604087015260181c16606085015260018101546080850152600281015460a0850152600381015460c08501520154910152565b600511156120cf57565b634e487b7160e01b5f52602160045260245ffd5b906040516120f081611f13565b606060ff6003839582815416612105816120c5565b85526001810154602086015261211d60028201611fc1565b60408601520154161515910152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60e0809180511515845260208101511515602085015260ff604082015116604085015260ff60608201511660608501526080810151608085015260a081015160a085015260c081015160c08501520151910152565b9081516121b1816120c5565b8152602082015160208201526060806121d9604085015160806040860152608085019061212c565b930151151591015290565b9181601f84011215610195578235916001600160401b038311610195576020838186019501011161019557565b6040906003190112610195576004359060243590565b90602080835192838152019201905f5b8181106122445750505090565b8251845260209384019390920191600101612237565b90606060031983011261019557600435916024356001600160401b0381116101955781612289916004016121e4565b92909291604435906001600160401b038211610195576122ab916004016121e4565b9091565b91908201809211611b3357565b906122c6816120c5565b60ff80198354169116179055565b8181106122df575050565b5f81556001016122d4565b3580151581036101955790565b60c43560ff811681036101955790565b903590601e198136030182121561019557018035906001600160401b0382116101955760200191813603831361019557565b908060209392818452848401375f828201840152601f01601f1916010190565b91908203918211611b3357565b6040519061237382611f2e565b5f60e0838281528260208201528260408201528260608201528260808201528260a08201528260c08201520152565b92906123bb906123c99593604086526040860191612339565b926020818503910152612339565b90565b5f198114611b335760010190565b6123e3816120c5565b6123ec826120c5565b808214612499576123fc816120c5565b6002811480156124b2575b801561249f575b6124995761241b816120c5565b801561247a5760039061242d816120c5565b1461243757505f90565b612440816120c5565b8015908115612465575b8115612454575090565b60019150612461816120c5565b1490565b9050612470816120c5565b600281149061244a565b50612484816120c5565b60038114908115612465578115612454575090565b50505f90565b506124a9816120c5565b6001811461240e565b506124bc816120c5565b60048114612407565b80519081156124995760205f9183829460405193849301835e8101838152039060025afa15610c7b575f516001600160f81b0316600160f81b1790565b5f5b804980156125245782146125205761251b906123cc565b612504565b5050565b630110b90960e61b5f5260045ffdfea26469706673582212205e3b9c7ddf726a735950f502cac4bf3b80f8671e1381dafaf6516f14ca9fd5d164736f6c634300081c0033",
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
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,string,bool)))
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
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,string,bool)))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [32]byte) (IProcessRegistryProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,bytes32,string,bool)))
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcess(processId [32]byte) (IProcessRegistryProcess, error) {
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
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,string,bool) census)
func (_ProcessRegistry *ProcessRegistryCaller) Processes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         IProcessRegistryEncryptionKey
	LatestStateRoot       *big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	BallotMode            IProcessRegistryBallotMode
	Census                IProcessRegistryCensus
}, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processes", arg0)

	outstruct := new(struct {
		Status                uint8
		OrganizationId        common.Address
		EncryptionKey         IProcessRegistryEncryptionKey
		LatestStateRoot       *big.Int
		StartTime             *big.Int
		Duration              *big.Int
		MaxVoters             *big.Int
		VotersCount           *big.Int
		OverwrittenVotesCount *big.Int
		CreationBlock         *big.Int
		BatchNumber           *big.Int
		MetadataURI           string
		BallotMode            IProcessRegistryBallotMode
		Census                IProcessRegistryCensus
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
	outstruct.MaxVoters = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VotersCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.OverwrittenVotesCount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.CreationBlock = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.BatchNumber = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.BallotMode = *abi.ConvertType(out[12], new(IProcessRegistryBallotMode)).(*IProcessRegistryBallotMode)
	outstruct.Census = *abi.ConvertType(out[13], new(IProcessRegistryCensus)).(*IProcessRegistryCensus)

	return *outstruct, err

}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,string,bool) census)
func (_ProcessRegistry *ProcessRegistrySession) Processes(arg0 [32]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         IProcessRegistryEncryptionKey
	LatestStateRoot       *big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	BallotMode            IProcessRegistryBallotMode
	Census                IProcessRegistryCensus
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 maxVoters, uint256 votersCount, uint256 overwrittenVotesCount, uint256 creationBlock, uint256 batchNumber, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,string,bool) census)
func (_ProcessRegistry *ProcessRegistryCallerSession) Processes(arg0 [32]byte) (struct {
	Status                uint8
	OrganizationId        common.Address
	EncryptionKey         IProcessRegistryEncryptionKey
	LatestStateRoot       *big.Int
	StartTime             *big.Int
	Duration              *big.Int
	MaxVoters             *big.Int
	VotersCount           *big.Int
	OverwrittenVotesCount *big.Int
	CreationBlock         *big.Int
	BatchNumber           *big.Int
	MetadataURI           string
	BallotMode            IProcessRegistryBallotMode
	Census                IProcessRegistryCensus
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

// NewProcess is a paid mutator transaction binding the contract method 0x2c645ddb.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,string,bool) census, string metadata, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns(bytes32)
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, initStateRoot)
}

// NewProcess is a paid mutator transaction binding the contract method 0x2c645ddb.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,string,bool) census, string metadata, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, initStateRoot)
}

// NewProcess is a paid mutator transaction binding the contract method 0x2c645ddb.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, uint256 maxVoters, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,bytes32,string,bool) census, string metadata, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns(bytes32)
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, maxVoters *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, maxVoters, ballotMode, census, metadata, encryptionKey, initStateRoot)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x44368408.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,bytes32,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessCensus(opts *bind.TransactOpts, processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessCensus", processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x44368408.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,bytes32,string,bool) census) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessCensus(processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x44368408.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,bytes32,string,bool) census) returns()
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
