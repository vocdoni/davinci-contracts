// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProposalRegistry

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

// ProposalRegistryCensus is an auto generated low-level Go binding around an user-defined struct.
type ProposalRegistryCensus struct {
	CensusOrigin  uint8
	MaxCensusSize *big.Int
	CensusRoot    [32]byte
	CensusURI     string
}

// ProposalRegistryProposal is an auto generated low-level Go binding around an user-defined struct.
type ProposalRegistryProposal struct {
	Status          uint8
	OrganizationId  [32]byte
	EncryptionKeys  [2][32]byte
	LatestStateRoot [32]byte
	Result          [][]*big.Int
	StartTime       *big.Int
	Duration        *big.Int
	MetadataURI     string
	Options         ProposalRegistryProposalOptions
	Census          ProposalRegistryCensus
}

// ProposalRegistryProposalOptions is an auto generated low-level Go binding around an user-defined struct.
type ProposalRegistryProposalOptions struct {
	EnvelopeType uint8
	ProposalMode uint8
	StartTime    *big.Int
	Duration     *big.Int
	Status       uint8
	VoteOptions  ProposalRegistryVoteOptions
}

// ProposalRegistryVoteOptions is an auto generated low-level Go binding around an user-defined struct.
type ProposalRegistryVoteOptions struct {
	UniqueValues  bool
	MaxCount      *big.Int
	MaxValue      *big.Int
	MinValue      *big.Int
	MaxOverwrites *big.Int
	MaxTotalCost  *big.Int
	MinTotalCost  *big.Int
	CostExponent  *big.Int
}

// ProposalRegistryMetaData contains all meta data concerning the ProposalRegistry contract.
var ProposalRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxCensusSize\",\"type\":\"uint256\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProposalDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProposalStateRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumProposalRegistry.ProposalStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProposalStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"}],\"name\":\"endProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"enumProposalRegistry.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"organizationId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"encryptionKeys\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[][]\",\"name\":\"result\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"envelopeType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"proposalMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"enumProposalRegistry.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOverwrites\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costExponent\",\"type\":\"uint256\"}],\"internalType\":\"structProposalRegistry.VoteOptions\",\"name\":\"voteOptions\",\"type\":\"tuple\"}],\"internalType\":\"structProposalRegistry.ProposalOptions\",\"name\":\"options\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumProposalRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxCensusSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structProposalRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structProposalRegistry.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chainID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_organizationRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"envelopeType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"proposalMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"enumProposalRegistry.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOverwrites\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costExponent\",\"type\":\"uint256\"}],\"internalType\":\"structProposalRegistry.VoteOptions\",\"name\":\"voteOptions\",\"type\":\"tuple\"}],\"internalType\":\"structProposalRegistry.ProposalOptions\",\"name\":\"_options\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumProposalRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxCensusSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structProposalRegistry.Census\",\"name\":\"_census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_organizationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_encryptionPubKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_initStateRoot\",\"type\":\"bytes32\"}],\"name\":\"newProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"enumProposalRegistry.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"organizationId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"latestStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"envelopeType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"proposalMode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"enumProposalRegistry.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"uniqueValues\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxOverwrites\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costExponent\",\"type\":\"uint256\"}],\"internalType\":\"structProposalRegistry.VoteOptions\",\"name\":\"voteOptions\",\"type\":\"tuple\"}],\"internalType\":\"structProposalRegistry.ProposalOptions\",\"name\":\"options\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumProposalRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxCensusSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structProposalRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumProposalRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxCensusSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structProposalRegistry.Census\",\"name\":\"_census\",\"type\":\"tuple\"}],\"name\":\"setProposalCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProposalDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[][]\",\"name\":\"_result\",\"type\":\"uint256[][]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"setProposalResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"enumProposalRegistry.ProposalStatus\",\"name\":\"_newStatus\",\"type\":\"uint8\"}],\"name\":\"setProposalStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_oldRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ProposalRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalRegistryMetaData.ABI instead.
var ProposalRegistryABI = ProposalRegistryMetaData.ABI

// ProposalRegistry is an auto generated Go binding around an Ethereum contract.
type ProposalRegistry struct {
	ProposalRegistryCaller     // Read-only binding to the contract
	ProposalRegistryTransactor // Write-only binding to the contract
	ProposalRegistryFilterer   // Log filterer for contract events
}

// ProposalRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalRegistrySession struct {
	Contract     *ProposalRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalRegistryCallerSession struct {
	Contract *ProposalRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProposalRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalRegistryTransactorSession struct {
	Contract     *ProposalRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProposalRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalRegistryRaw struct {
	Contract *ProposalRegistry // Generic contract binding to access the raw methods on
}

// ProposalRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalRegistryCallerRaw struct {
	Contract *ProposalRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalRegistryTransactorRaw struct {
	Contract *ProposalRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalRegistry creates a new instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistry(address common.Address, backend bind.ContractBackend) (*ProposalRegistry, error) {
	contract, err := bindProposalRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistry{ProposalRegistryCaller: ProposalRegistryCaller{contract: contract}, ProposalRegistryTransactor: ProposalRegistryTransactor{contract: contract}, ProposalRegistryFilterer: ProposalRegistryFilterer{contract: contract}}, nil
}

// NewProposalRegistryCaller creates a new read-only instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProposalRegistryCaller, error) {
	contract, err := bindProposalRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryCaller{contract: contract}, nil
}

// NewProposalRegistryTransactor creates a new write-only instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalRegistryTransactor, error) {
	contract, err := bindProposalRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryTransactor{contract: contract}, nil
}

// NewProposalRegistryFilterer creates a new log filterer instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalRegistryFilterer, error) {
	contract, err := bindProposalRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryFilterer{contract: contract}, nil
}

// bindProposalRegistry binds a generic wrapper to an already deployed contract.
func bindProposalRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposalRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalRegistry *ProposalRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalRegistry.Contract.ProposalRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalRegistry *ProposalRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.ProposalRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalRegistry *ProposalRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.ProposalRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalRegistry *ProposalRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalRegistry *ProposalRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalRegistry *ProposalRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ProposalRegistry *ProposalRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ProposalRegistry *ProposalRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ProposalRegistry.Contract.UPGRADEINTERFACEVERSION(&_ProposalRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ProposalRegistry *ProposalRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ProposalRegistry.Contract.UPGRADEINTERFACEVERSION(&_ProposalRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProposalRegistry *ProposalRegistryCaller) ChainID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProposalRegistry *ProposalRegistrySession) ChainID() (string, error) {
	return _ProposalRegistry.Contract.ChainID(&_ProposalRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProposalRegistry *ProposalRegistryCallerSession) ChainID() (string, error) {
	return _ProposalRegistry.Contract.ChainID(&_ProposalRegistry.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0x430694cf.
//
// Solidity: function getProposal(bytes32 _proposalID) view returns((uint8,bytes32,bytes32[2],bytes32,uint256[][],uint256,uint256,string,(uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)),(uint8,uint256,bytes32,string)))
func (_ProposalRegistry *ProposalRegistryCaller) GetProposal(opts *bind.CallOpts, _proposalID [32]byte) (ProposalRegistryProposal, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getProposal", _proposalID)

	if err != nil {
		return *new(ProposalRegistryProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(ProposalRegistryProposal)).(*ProposalRegistryProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0x430694cf.
//
// Solidity: function getProposal(bytes32 _proposalID) view returns((uint8,bytes32,bytes32[2],bytes32,uint256[][],uint256,uint256,string,(uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)),(uint8,uint256,bytes32,string)))
func (_ProposalRegistry *ProposalRegistrySession) GetProposal(_proposalID [32]byte) (ProposalRegistryProposal, error) {
	return _ProposalRegistry.Contract.GetProposal(&_ProposalRegistry.CallOpts, _proposalID)
}

// GetProposal is a free data retrieval call binding the contract method 0x430694cf.
//
// Solidity: function getProposal(bytes32 _proposalID) view returns((uint8,bytes32,bytes32[2],bytes32,uint256[][],uint256,uint256,string,(uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)),(uint8,uint256,bytes32,string)))
func (_ProposalRegistry *ProposalRegistryCallerSession) GetProposal(_proposalID [32]byte) (ProposalRegistryProposal, error) {
	return _ProposalRegistry.Contract.GetProposal(&_ProposalRegistry.CallOpts, _proposalID)
}

// OrganizationRegistry is a free data retrieval call binding the contract method 0x8cafab7f.
//
// Solidity: function organizationRegistry() view returns(address)
func (_ProposalRegistry *ProposalRegistryCaller) OrganizationRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "organizationRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrganizationRegistry is a free data retrieval call binding the contract method 0x8cafab7f.
//
// Solidity: function organizationRegistry() view returns(address)
func (_ProposalRegistry *ProposalRegistrySession) OrganizationRegistry() (common.Address, error) {
	return _ProposalRegistry.Contract.OrganizationRegistry(&_ProposalRegistry.CallOpts)
}

// OrganizationRegistry is a free data retrieval call binding the contract method 0x8cafab7f.
//
// Solidity: function organizationRegistry() view returns(address)
func (_ProposalRegistry *ProposalRegistryCallerSession) OrganizationRegistry() (common.Address, error) {
	return _ProposalRegistry.Contract.OrganizationRegistry(&_ProposalRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProposalRegistry *ProposalRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProposalRegistry *ProposalRegistrySession) Owner() (common.Address, error) {
	return _ProposalRegistry.Contract.Owner(&_ProposalRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProposalRegistry *ProposalRegistryCallerSession) Owner() (common.Address, error) {
	return _ProposalRegistry.Contract.Owner(&_ProposalRegistry.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint32)
func (_ProposalRegistry *ProposalRegistryCaller) ProposalCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint32)
func (_ProposalRegistry *ProposalRegistrySession) ProposalCount() (uint32, error) {
	return _ProposalRegistry.Contract.ProposalCount(&_ProposalRegistry.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint32)
func (_ProposalRegistry *ProposalRegistryCallerSession) ProposalCount() (uint32, error) {
	return _ProposalRegistry.Contract.ProposalCount(&_ProposalRegistry.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 status, bytes32 organizationId, bytes32 latestStateRoot, uint256 startTime, uint256 duration, string metadataURI, (uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)) options, (uint8,uint256,bytes32,string) census)
func (_ProposalRegistry *ProposalRegistryCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status          uint8
	OrganizationId  [32]byte
	LatestStateRoot [32]byte
	StartTime       *big.Int
	Duration        *big.Int
	MetadataURI     string
	Options         ProposalRegistryProposalOptions
	Census          ProposalRegistryCensus
}, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Status          uint8
		OrganizationId  [32]byte
		LatestStateRoot [32]byte
		StartTime       *big.Int
		Duration        *big.Int
		MetadataURI     string
		Options         ProposalRegistryProposalOptions
		Census          ProposalRegistryCensus
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.OrganizationId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.LatestStateRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Options = *abi.ConvertType(out[6], new(ProposalRegistryProposalOptions)).(*ProposalRegistryProposalOptions)
	outstruct.Census = *abi.ConvertType(out[7], new(ProposalRegistryCensus)).(*ProposalRegistryCensus)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 status, bytes32 organizationId, bytes32 latestStateRoot, uint256 startTime, uint256 duration, string metadataURI, (uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)) options, (uint8,uint256,bytes32,string) census)
func (_ProposalRegistry *ProposalRegistrySession) Proposals(arg0 [32]byte) (struct {
	Status          uint8
	OrganizationId  [32]byte
	LatestStateRoot [32]byte
	StartTime       *big.Int
	Duration        *big.Int
	MetadataURI     string
	Options         ProposalRegistryProposalOptions
	Census          ProposalRegistryCensus
}, error) {
	return _ProposalRegistry.Contract.Proposals(&_ProposalRegistry.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 status, bytes32 organizationId, bytes32 latestStateRoot, uint256 startTime, uint256 duration, string metadataURI, (uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)) options, (uint8,uint256,bytes32,string) census)
func (_ProposalRegistry *ProposalRegistryCallerSession) Proposals(arg0 [32]byte) (struct {
	Status          uint8
	OrganizationId  [32]byte
	LatestStateRoot [32]byte
	StartTime       *big.Int
	Duration        *big.Int
	MetadataURI     string
	Options         ProposalRegistryProposalOptions
	Census          ProposalRegistryCensus
}, error) {
	return _ProposalRegistry.Contract.Proposals(&_ProposalRegistry.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProposalRegistry *ProposalRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProposalRegistry *ProposalRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ProposalRegistry.Contract.ProxiableUUID(&_ProposalRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProposalRegistry *ProposalRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ProposalRegistry.Contract.ProxiableUUID(&_ProposalRegistry.CallOpts)
}

// EndProposal is a paid mutator transaction binding the contract method 0xb3b47061.
//
// Solidity: function endProposal(bytes32 _proposalID) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) EndProposal(opts *bind.TransactOpts, _proposalID [32]byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "endProposal", _proposalID)
}

// EndProposal is a paid mutator transaction binding the contract method 0xb3b47061.
//
// Solidity: function endProposal(bytes32 _proposalID) returns()
func (_ProposalRegistry *ProposalRegistrySession) EndProposal(_proposalID [32]byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.EndProposal(&_ProposalRegistry.TransactOpts, _proposalID)
}

// EndProposal is a paid mutator transaction binding the contract method 0xb3b47061.
//
// Solidity: function endProposal(bytes32 _proposalID) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) EndProposal(_proposalID [32]byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.EndProposal(&_ProposalRegistry.TransactOpts, _proposalID)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string _chainID, address _organizationRegistry) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) Initialize(opts *bind.TransactOpts, _chainID string, _organizationRegistry common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "initialize", _chainID, _organizationRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string _chainID, address _organizationRegistry) returns()
func (_ProposalRegistry *ProposalRegistrySession) Initialize(_chainID string, _organizationRegistry common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.Initialize(&_ProposalRegistry.TransactOpts, _chainID, _organizationRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string _chainID, address _organizationRegistry) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) Initialize(_chainID string, _organizationRegistry common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.Initialize(&_ProposalRegistry.TransactOpts, _chainID, _organizationRegistry)
}

// NewProposal is a paid mutator transaction binding the contract method 0xb2dc5ded.
//
// Solidity: function newProposal((uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)) _options, (uint8,uint256,bytes32,string) _census, string _metadata, bytes32 _organizationID, bytes32 _proposalID, bytes32 _encryptionPubKey, bytes32 _initStateRoot) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) NewProposal(opts *bind.TransactOpts, _options ProposalRegistryProposalOptions, _census ProposalRegistryCensus, _metadata string, _organizationID [32]byte, _proposalID [32]byte, _encryptionPubKey [32]byte, _initStateRoot [32]byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "newProposal", _options, _census, _metadata, _organizationID, _proposalID, _encryptionPubKey, _initStateRoot)
}

// NewProposal is a paid mutator transaction binding the contract method 0xb2dc5ded.
//
// Solidity: function newProposal((uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)) _options, (uint8,uint256,bytes32,string) _census, string _metadata, bytes32 _organizationID, bytes32 _proposalID, bytes32 _encryptionPubKey, bytes32 _initStateRoot) returns()
func (_ProposalRegistry *ProposalRegistrySession) NewProposal(_options ProposalRegistryProposalOptions, _census ProposalRegistryCensus, _metadata string, _organizationID [32]byte, _proposalID [32]byte, _encryptionPubKey [32]byte, _initStateRoot [32]byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.NewProposal(&_ProposalRegistry.TransactOpts, _options, _census, _metadata, _organizationID, _proposalID, _encryptionPubKey, _initStateRoot)
}

// NewProposal is a paid mutator transaction binding the contract method 0xb2dc5ded.
//
// Solidity: function newProposal((uint8,uint8,uint256,uint256,uint8,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)) _options, (uint8,uint256,bytes32,string) _census, string _metadata, bytes32 _organizationID, bytes32 _proposalID, bytes32 _encryptionPubKey, bytes32 _initStateRoot) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) NewProposal(_options ProposalRegistryProposalOptions, _census ProposalRegistryCensus, _metadata string, _organizationID [32]byte, _proposalID [32]byte, _encryptionPubKey [32]byte, _initStateRoot [32]byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.NewProposal(&_ProposalRegistry.TransactOpts, _options, _census, _metadata, _organizationID, _proposalID, _encryptionPubKey, _initStateRoot)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProposalRegistry *ProposalRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProposalRegistry *ProposalRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ProposalRegistry.Contract.RenounceOwnership(&_ProposalRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProposalRegistry.Contract.RenounceOwnership(&_ProposalRegistry.TransactOpts)
}

// SetProposalCensus is a paid mutator transaction binding the contract method 0x17f02ebe.
//
// Solidity: function setProposalCensus(bytes32 _proposalID, (uint8,uint256,bytes32,string) _census) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) SetProposalCensus(opts *bind.TransactOpts, _proposalID [32]byte, _census ProposalRegistryCensus) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "setProposalCensus", _proposalID, _census)
}

// SetProposalCensus is a paid mutator transaction binding the contract method 0x17f02ebe.
//
// Solidity: function setProposalCensus(bytes32 _proposalID, (uint8,uint256,bytes32,string) _census) returns()
func (_ProposalRegistry *ProposalRegistrySession) SetProposalCensus(_proposalID [32]byte, _census ProposalRegistryCensus) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalCensus(&_ProposalRegistry.TransactOpts, _proposalID, _census)
}

// SetProposalCensus is a paid mutator transaction binding the contract method 0x17f02ebe.
//
// Solidity: function setProposalCensus(bytes32 _proposalID, (uint8,uint256,bytes32,string) _census) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) SetProposalCensus(_proposalID [32]byte, _census ProposalRegistryCensus) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalCensus(&_ProposalRegistry.TransactOpts, _proposalID, _census)
}

// SetProposalDuration is a paid mutator transaction binding the contract method 0xc0bd47a1.
//
// Solidity: function setProposalDuration(bytes32 _proposalID, uint256 _duration) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) SetProposalDuration(opts *bind.TransactOpts, _proposalID [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "setProposalDuration", _proposalID, _duration)
}

// SetProposalDuration is a paid mutator transaction binding the contract method 0xc0bd47a1.
//
// Solidity: function setProposalDuration(bytes32 _proposalID, uint256 _duration) returns()
func (_ProposalRegistry *ProposalRegistrySession) SetProposalDuration(_proposalID [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalDuration(&_ProposalRegistry.TransactOpts, _proposalID, _duration)
}

// SetProposalDuration is a paid mutator transaction binding the contract method 0xc0bd47a1.
//
// Solidity: function setProposalDuration(bytes32 _proposalID, uint256 _duration) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) SetProposalDuration(_proposalID [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalDuration(&_ProposalRegistry.TransactOpts, _proposalID, _duration)
}

// SetProposalResult is a paid mutator transaction binding the contract method 0x09d72ae7.
//
// Solidity: function setProposalResult(bytes32 _proposalID, uint256[][] _result, bytes _proof) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) SetProposalResult(opts *bind.TransactOpts, _proposalID [32]byte, _result [][]*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "setProposalResult", _proposalID, _result, _proof)
}

// SetProposalResult is a paid mutator transaction binding the contract method 0x09d72ae7.
//
// Solidity: function setProposalResult(bytes32 _proposalID, uint256[][] _result, bytes _proof) returns()
func (_ProposalRegistry *ProposalRegistrySession) SetProposalResult(_proposalID [32]byte, _result [][]*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalResult(&_ProposalRegistry.TransactOpts, _proposalID, _result, _proof)
}

// SetProposalResult is a paid mutator transaction binding the contract method 0x09d72ae7.
//
// Solidity: function setProposalResult(bytes32 _proposalID, uint256[][] _result, bytes _proof) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) SetProposalResult(_proposalID [32]byte, _result [][]*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalResult(&_ProposalRegistry.TransactOpts, _proposalID, _result, _proof)
}

// SetProposalStatus is a paid mutator transaction binding the contract method 0x39eacf34.
//
// Solidity: function setProposalStatus(bytes32 _proposalID, uint8 _newStatus) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) SetProposalStatus(opts *bind.TransactOpts, _proposalID [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "setProposalStatus", _proposalID, _newStatus)
}

// SetProposalStatus is a paid mutator transaction binding the contract method 0x39eacf34.
//
// Solidity: function setProposalStatus(bytes32 _proposalID, uint8 _newStatus) returns()
func (_ProposalRegistry *ProposalRegistrySession) SetProposalStatus(_proposalID [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalStatus(&_ProposalRegistry.TransactOpts, _proposalID, _newStatus)
}

// SetProposalStatus is a paid mutator transaction binding the contract method 0x39eacf34.
//
// Solidity: function setProposalStatus(bytes32 _proposalID, uint8 _newStatus) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) SetProposalStatus(_proposalID [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SetProposalStatus(&_ProposalRegistry.TransactOpts, _proposalID, _newStatus)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xb66ba6eb.
//
// Solidity: function submitStateTransition(bytes32 _proposalID, bytes32 _oldRoot, bytes32 _newRoot, bytes _proof) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) SubmitStateTransition(opts *bind.TransactOpts, _proposalID [32]byte, _oldRoot [32]byte, _newRoot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "submitStateTransition", _proposalID, _oldRoot, _newRoot, _proof)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xb66ba6eb.
//
// Solidity: function submitStateTransition(bytes32 _proposalID, bytes32 _oldRoot, bytes32 _newRoot, bytes _proof) returns()
func (_ProposalRegistry *ProposalRegistrySession) SubmitStateTransition(_proposalID [32]byte, _oldRoot [32]byte, _newRoot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SubmitStateTransition(&_ProposalRegistry.TransactOpts, _proposalID, _oldRoot, _newRoot, _proof)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xb66ba6eb.
//
// Solidity: function submitStateTransition(bytes32 _proposalID, bytes32 _oldRoot, bytes32 _newRoot, bytes _proof) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) SubmitStateTransition(_proposalID [32]byte, _oldRoot [32]byte, _newRoot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.SubmitStateTransition(&_ProposalRegistry.TransactOpts, _proposalID, _oldRoot, _newRoot, _proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProposalRegistry *ProposalRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.TransferOwnership(&_ProposalRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.TransferOwnership(&_ProposalRegistry.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProposalRegistry *ProposalRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProposalRegistry *ProposalRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.UpgradeToAndCall(&_ProposalRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.UpgradeToAndCall(&_ProposalRegistry.TransactOpts, newImplementation, data)
}

// ProposalRegistryCensusUpdatedIterator is returned from FilterCensusUpdated and is used to iterate over the raw logs and unpacked data for CensusUpdated events raised by the ProposalRegistry contract.
type ProposalRegistryCensusUpdatedIterator struct {
	Event *ProposalRegistryCensusUpdated // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryCensusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryCensusUpdated)
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
		it.Event = new(ProposalRegistryCensusUpdated)
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
func (it *ProposalRegistryCensusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryCensusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryCensusUpdated represents a CensusUpdated event raised by the ProposalRegistry contract.
type ProposalRegistryCensusUpdated struct {
	ProposalID    [32]byte
	CensusRoot    [32]byte
	CensusURI     string
	MaxCensusSize *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCensusUpdated is a free log retrieval operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed proposalID, bytes32 censusRoot, string censusURI, uint256 maxCensusSize)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterCensusUpdated(opts *bind.FilterOpts, proposalID [][32]byte) (*ProposalRegistryCensusUpdatedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "CensusUpdated", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryCensusUpdatedIterator{contract: _ProposalRegistry.contract, event: "CensusUpdated", logs: logs, sub: sub}, nil
}

// WatchCensusUpdated is a free log subscription operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed proposalID, bytes32 censusRoot, string censusURI, uint256 maxCensusSize)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchCensusUpdated(opts *bind.WatchOpts, sink chan<- *ProposalRegistryCensusUpdated, proposalID [][32]byte) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "CensusUpdated", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryCensusUpdated)
				if err := _ProposalRegistry.contract.UnpackLog(event, "CensusUpdated", log); err != nil {
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
// Solidity: event CensusUpdated(bytes32 indexed proposalID, bytes32 censusRoot, string censusURI, uint256 maxCensusSize)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseCensusUpdated(log types.Log) (*ProposalRegistryCensusUpdated, error) {
	event := new(ProposalRegistryCensusUpdated)
	if err := _ProposalRegistry.contract.UnpackLog(event, "CensusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProposalRegistry contract.
type ProposalRegistryInitializedIterator struct {
	Event *ProposalRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryInitialized)
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
		it.Event = new(ProposalRegistryInitialized)
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
func (it *ProposalRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryInitialized represents a Initialized event raised by the ProposalRegistry contract.
type ProposalRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProposalRegistryInitializedIterator, error) {

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryInitializedIterator{contract: _ProposalRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProposalRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryInitialized)
				if err := _ProposalRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseInitialized(log types.Log) (*ProposalRegistryInitialized, error) {
	event := new(ProposalRegistryInitialized)
	if err := _ProposalRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProposalRegistry contract.
type ProposalRegistryOwnershipTransferredIterator struct {
	Event *ProposalRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryOwnershipTransferred)
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
		it.Event = new(ProposalRegistryOwnershipTransferred)
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
func (it *ProposalRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ProposalRegistry contract.
type ProposalRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProposalRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryOwnershipTransferredIterator{contract: _ProposalRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProposalRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryOwnershipTransferred)
				if err := _ProposalRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ProposalRegistryOwnershipTransferred, error) {
	event := new(ProposalRegistryOwnershipTransferred)
	if err := _ProposalRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the ProposalRegistry contract.
type ProposalRegistryProposalCreatedIterator struct {
	Event *ProposalRegistryProposalCreated // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryProposalCreated)
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
		it.Event = new(ProposalRegistryProposalCreated)
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
func (it *ProposalRegistryProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryProposalCreated represents a ProposalCreated event raised by the ProposalRegistry contract.
type ProposalRegistryProposalCreated struct {
	ProposalID [32]byte
	Creator    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xe90bb18a6fc9203e3678a7f41e132910465e0fce1421609662d02c481fbfbcf8.
//
// Solidity: event ProposalCreated(bytes32 indexed proposalID, address indexed creator)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterProposalCreated(opts *bind.FilterOpts, proposalID [][32]byte, creator []common.Address) (*ProposalRegistryProposalCreatedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "ProposalCreated", proposalIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryProposalCreatedIterator{contract: _ProposalRegistry.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xe90bb18a6fc9203e3678a7f41e132910465e0fce1421609662d02c481fbfbcf8.
//
// Solidity: event ProposalCreated(bytes32 indexed proposalID, address indexed creator)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *ProposalRegistryProposalCreated, proposalID [][32]byte, creator []common.Address) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "ProposalCreated", proposalIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryProposalCreated)
				if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xe90bb18a6fc9203e3678a7f41e132910465e0fce1421609662d02c481fbfbcf8.
//
// Solidity: event ProposalCreated(bytes32 indexed proposalID, address indexed creator)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseProposalCreated(log types.Log) (*ProposalRegistryProposalCreated, error) {
	event := new(ProposalRegistryProposalCreated)
	if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryProposalDurationChangedIterator is returned from FilterProposalDurationChanged and is used to iterate over the raw logs and unpacked data for ProposalDurationChanged events raised by the ProposalRegistry contract.
type ProposalRegistryProposalDurationChangedIterator struct {
	Event *ProposalRegistryProposalDurationChanged // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryProposalDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryProposalDurationChanged)
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
		it.Event = new(ProposalRegistryProposalDurationChanged)
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
func (it *ProposalRegistryProposalDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryProposalDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryProposalDurationChanged represents a ProposalDurationChanged event raised by the ProposalRegistry contract.
type ProposalRegistryProposalDurationChanged struct {
	ProposalID [32]byte
	Duration   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalDurationChanged is a free log retrieval operation binding the contract event 0x983165304729e1c4e36ce532740769c8a5c4b56dc61c3c60b3f96322dac48091.
//
// Solidity: event ProposalDurationChanged(bytes32 indexed proposalID, uint256 duration)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterProposalDurationChanged(opts *bind.FilterOpts, proposalID [][32]byte) (*ProposalRegistryProposalDurationChangedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "ProposalDurationChanged", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryProposalDurationChangedIterator{contract: _ProposalRegistry.contract, event: "ProposalDurationChanged", logs: logs, sub: sub}, nil
}

// WatchProposalDurationChanged is a free log subscription operation binding the contract event 0x983165304729e1c4e36ce532740769c8a5c4b56dc61c3c60b3f96322dac48091.
//
// Solidity: event ProposalDurationChanged(bytes32 indexed proposalID, uint256 duration)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchProposalDurationChanged(opts *bind.WatchOpts, sink chan<- *ProposalRegistryProposalDurationChanged, proposalID [][32]byte) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "ProposalDurationChanged", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryProposalDurationChanged)
				if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalDurationChanged", log); err != nil {
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

// ParseProposalDurationChanged is a log parse operation binding the contract event 0x983165304729e1c4e36ce532740769c8a5c4b56dc61c3c60b3f96322dac48091.
//
// Solidity: event ProposalDurationChanged(bytes32 indexed proposalID, uint256 duration)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseProposalDurationChanged(log types.Log) (*ProposalRegistryProposalDurationChanged, error) {
	event := new(ProposalRegistryProposalDurationChanged)
	if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryProposalStateRootUpdatedIterator is returned from FilterProposalStateRootUpdated and is used to iterate over the raw logs and unpacked data for ProposalStateRootUpdated events raised by the ProposalRegistry contract.
type ProposalRegistryProposalStateRootUpdatedIterator struct {
	Event *ProposalRegistryProposalStateRootUpdated // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryProposalStateRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryProposalStateRootUpdated)
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
		it.Event = new(ProposalRegistryProposalStateRootUpdated)
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
func (it *ProposalRegistryProposalStateRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryProposalStateRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryProposalStateRootUpdated represents a ProposalStateRootUpdated event raised by the ProposalRegistry contract.
type ProposalRegistryProposalStateRootUpdated struct {
	ProposalID   [32]byte
	NewStateRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalStateRootUpdated is a free log retrieval operation binding the contract event 0x6fff2d70a00d754c43c8c120be5e18415b3443738a8cff4b907faa8c52e173a2.
//
// Solidity: event ProposalStateRootUpdated(bytes32 indexed proposalID, bytes32 newStateRoot)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterProposalStateRootUpdated(opts *bind.FilterOpts, proposalID [][32]byte) (*ProposalRegistryProposalStateRootUpdatedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "ProposalStateRootUpdated", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryProposalStateRootUpdatedIterator{contract: _ProposalRegistry.contract, event: "ProposalStateRootUpdated", logs: logs, sub: sub}, nil
}

// WatchProposalStateRootUpdated is a free log subscription operation binding the contract event 0x6fff2d70a00d754c43c8c120be5e18415b3443738a8cff4b907faa8c52e173a2.
//
// Solidity: event ProposalStateRootUpdated(bytes32 indexed proposalID, bytes32 newStateRoot)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchProposalStateRootUpdated(opts *bind.WatchOpts, sink chan<- *ProposalRegistryProposalStateRootUpdated, proposalID [][32]byte) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "ProposalStateRootUpdated", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryProposalStateRootUpdated)
				if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalStateRootUpdated", log); err != nil {
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

// ParseProposalStateRootUpdated is a log parse operation binding the contract event 0x6fff2d70a00d754c43c8c120be5e18415b3443738a8cff4b907faa8c52e173a2.
//
// Solidity: event ProposalStateRootUpdated(bytes32 indexed proposalID, bytes32 newStateRoot)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseProposalStateRootUpdated(log types.Log) (*ProposalRegistryProposalStateRootUpdated, error) {
	event := new(ProposalRegistryProposalStateRootUpdated)
	if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalStateRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryProposalStatusChangedIterator is returned from FilterProposalStatusChanged and is used to iterate over the raw logs and unpacked data for ProposalStatusChanged events raised by the ProposalRegistry contract.
type ProposalRegistryProposalStatusChangedIterator struct {
	Event *ProposalRegistryProposalStatusChanged // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryProposalStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryProposalStatusChanged)
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
		it.Event = new(ProposalRegistryProposalStatusChanged)
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
func (it *ProposalRegistryProposalStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryProposalStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryProposalStatusChanged represents a ProposalStatusChanged event raised by the ProposalRegistry contract.
type ProposalRegistryProposalStatusChanged struct {
	ProposalID [32]byte
	NewStatus  uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalStatusChanged is a free log retrieval operation binding the contract event 0x7def22fee322edec5f79e74d0747e712935e3076161cceab960aa0daca48e8d2.
//
// Solidity: event ProposalStatusChanged(bytes32 indexed proposalID, uint8 newStatus)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterProposalStatusChanged(opts *bind.FilterOpts, proposalID [][32]byte) (*ProposalRegistryProposalStatusChangedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "ProposalStatusChanged", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryProposalStatusChangedIterator{contract: _ProposalRegistry.contract, event: "ProposalStatusChanged", logs: logs, sub: sub}, nil
}

// WatchProposalStatusChanged is a free log subscription operation binding the contract event 0x7def22fee322edec5f79e74d0747e712935e3076161cceab960aa0daca48e8d2.
//
// Solidity: event ProposalStatusChanged(bytes32 indexed proposalID, uint8 newStatus)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchProposalStatusChanged(opts *bind.WatchOpts, sink chan<- *ProposalRegistryProposalStatusChanged, proposalID [][32]byte) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "ProposalStatusChanged", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryProposalStatusChanged)
				if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalStatusChanged", log); err != nil {
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

// ParseProposalStatusChanged is a log parse operation binding the contract event 0x7def22fee322edec5f79e74d0747e712935e3076161cceab960aa0daca48e8d2.
//
// Solidity: event ProposalStatusChanged(bytes32 indexed proposalID, uint8 newStatus)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseProposalStatusChanged(log types.Log) (*ProposalRegistryProposalStatusChanged, error) {
	event := new(ProposalRegistryProposalStatusChanged)
	if err := _ProposalRegistry.contract.UnpackLog(event, "ProposalStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ProposalRegistry contract.
type ProposalRegistryUpgradedIterator struct {
	Event *ProposalRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryUpgraded)
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
		it.Event = new(ProposalRegistryUpgraded)
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
func (it *ProposalRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryUpgraded represents a Upgraded event raised by the ProposalRegistry contract.
type ProposalRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ProposalRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryUpgradedIterator{contract: _ProposalRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ProposalRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryUpgraded)
				if err := _ProposalRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseUpgraded(log types.Log) (*ProposalRegistryUpgraded, error) {
	event := new(ProposalRegistryUpgraded)
	if err := _ProposalRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
