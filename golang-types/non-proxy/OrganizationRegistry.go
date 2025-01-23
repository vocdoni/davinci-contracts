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

// OrganizationRegistryMetaData contains all meta data concerning the OrganizationRegistry contract.
var OrganizationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"OrganizationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"}],\"name\":\"OrganizationUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"addAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"administrators\",\"type\":\"address[]\"}],\"name\":\"createOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"deleteOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"getOrganization\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdministrator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"organizations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"}],\"name\":\"removeAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506110cc8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80637acbb8af116100665780637acbb8af146100fd578063c1af6e0314610110578063c2a950be1461015e578063d2c30a6d14610171578063f1c621041461018457600080fd5b80631c2e3d82146100985780633c10eee5146100ad5780635a1f7406146100c05780636cca67bf146100ea575b600080fd5b6100ab6100a6366004610b12565b6101a9565b005b6100ab6100bb366004610b12565b610297565b6100d36100ce366004610b45565b610379565b6040516100e1929190610bad565b60405180910390f35b6100ab6100f8366004610c24565b6104a5565b6100ab61010b366004610b45565b61061d565b61014e61011e366004610b12565b6001600160a01b039182166000908152602081815260408083209390941682526002909201909152205460ff1690565b60405190151581526020016100e1565b6100ab61016c366004610caa565b6106ca565b6100d361017f366004610b45565b610959565b6001546101949063ffffffff1681565b60405163ffffffff90911681526020016100e1565b6001600160a01b038216600090815260208181526040808320338452600201909152902054829060ff166101f85760405162461bcd60e51b81526004016101ef90610d96565b60405180910390fd5b6001600160a01b0383166000908152602081905260408120805461021b90610de0565b90501161023a5760405162461bcd60e51b81526004016101ef90610e1a565b6001600160a01b0382166102605760405162461bcd60e51b81526004016101ef90610e6b565b506001600160a01b03918216600090815260208181526040808320939094168252600290920190915220805460ff19166001179055565b6001600160a01b038216600090815260208181526040808320338452600201909152902054829060ff166102dd5760405162461bcd60e51b81526004016101ef90610d96565b6001600160a01b0383166000908152602081905260408120805461030090610de0565b90501161031f5760405162461bcd60e51b81526004016101ef90610e1a565b6001600160a01b0382166103455760405162461bcd60e51b81526004016101ef90610e6b565b506001600160a01b03918216600090815260208181526040808320939094168252600290920190915220805460ff19169055565b60006020819052908152604090208054819061039490610de0565b80601f01602080910402602001604051908101604052809291908181526020018280546103c090610de0565b801561040d5780601f106103e25761010080835404028352916020019161040d565b820191906000526020600020905b8154815290600101906020018083116103f057829003601f168201915b50505050509080600101805461042290610de0565b80601f016020809104026020016040519081016040528092919081815260200182805461044e90610de0565b801561049b5780601f106104705761010080835404028352916020019161049b565b820191906000526020600020905b81548152906001019060200180831161047e57829003601f168201915b5050505050905082565b6001600160a01b038516600090815260208181526040808320338452600201909152902054859060ff166104eb5760405162461bcd60e51b81526004016101ef90610d96565b836105085760405162461bcd60e51b81526004016101ef90610ebe565b816105675760405162461bcd60e51b815260206004820152602960248201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c6964206d6560448201526874616461746155524960b81b60648201526084016101ef565b6001600160a01b0386166000908152602081905260408120805461058a90610de0565b9050116105a95760405162461bcd60e51b81526004016101ef90610e1a565b6001600160a01b0386166000908152602081905260409020806105cd868883610f65565b50600181016105dd848683610f65565b5060405133906001600160a01b038916907fdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f690600090a350505050505050565b6001600160a01b0381166000908152602081905260408120805461064090610de0565b90501161065f5760405162461bcd60e51b81526004016101ef90610e1a565b6001600160a01b0381166000908152602081905260408120906106828282610aa0565b610690600183016000610aa0565b50506001805463ffffffff169060006106a88361103b565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b6001600160a01b0387166107205760405162461bcd60e51b815260206004820181905260248201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c696420696460448201526064016101ef565b8461073d5760405162461bcd60e51b81526004016101ef90610ebe565b6001600160a01b0387166000908152602081905260409020805461076090610de0565b1590506107c95760405162461bcd60e51b815260206004820152603160248201527f4f7267616e697a6174696f6e52656769737472793a206f7267616e697a6174696044820152706f6e20616c72656164792065786973747360781b60648201526084016101ef565b6001600160a01b0387166000908152602081905260409020806107ed878983610f65565b50600181016107fd858783610f65565b5081156108bf5760005b828110156108bd5760008484838181106108235761082361105b565b90506020020160208101906108389190610b45565b6001600160a01b03160361085e5760405162461bcd60e51b81526004016101ef90610e6b565b60018260020160008686858181106108785761087861105b565b905060200201602081019061088d9190610b45565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055600101610807565b505b3360009081526002820160205260408120805460ff19166001908117909155805463ffffffff16916108f083611071565b91906101000a81548163ffffffff021916908363ffffffff16021790555050336001600160a01b0316886001600160a01b03167f2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a60405160405180910390a35050505050505050565b6001600160a01b03811660009081526020819052604090208054606091829181906001820190829061098a90610de0565b80601f01602080910402602001604051908101604052809291908181526020018280546109b690610de0565b8015610a035780601f106109d857610100808354040283529160200191610a03565b820191906000526020600020905b8154815290600101906020018083116109e657829003601f168201915b50505050509150808054610a1690610de0565b80601f0160208091040260200160405190810160405280929190818152602001828054610a4290610de0565b8015610a8f5780601f10610a6457610100808354040283529160200191610a8f565b820191906000526020600020905b815481529060010190602001808311610a7257829003601f168201915b505050505090509250925050915091565b508054610aac90610de0565b6000825580601f10610abc575050565b601f016020900490600052602060002090810190610ada9190610add565b50565b5b80821115610af25760008155600101610ade565b5090565b80356001600160a01b0381168114610b0d57600080fd5b919050565b60008060408385031215610b2557600080fd5b610b2e83610af6565b9150610b3c60208401610af6565b90509250929050565b600060208284031215610b5757600080fd5b610b6082610af6565b9392505050565b6000815180845260005b81811015610b8d57602081850181015186830182015201610b71565b506000602082860101526020601f19601f83011685010191505092915050565b604081526000610bc06040830185610b67565b8281036020840152610bd28185610b67565b95945050505050565b60008083601f840112610bed57600080fd5b50813567ffffffffffffffff811115610c0557600080fd5b602083019150836020828501011115610c1d57600080fd5b9250929050565b600080600080600060608688031215610c3c57600080fd5b610c4586610af6565b9450602086013567ffffffffffffffff811115610c6157600080fd5b610c6d88828901610bdb565b909550935050604086013567ffffffffffffffff811115610c8d57600080fd5b610c9988828901610bdb565b969995985093965092949392505050565b60008060008060008060006080888a031215610cc557600080fd5b610cce88610af6565b9650602088013567ffffffffffffffff811115610cea57600080fd5b610cf68a828b01610bdb565b909750955050604088013567ffffffffffffffff811115610d1657600080fd5b610d228a828b01610bdb565b909550935050606088013567ffffffffffffffff811115610d4257600080fd5b8801601f81018a13610d5357600080fd5b803567ffffffffffffffff811115610d6a57600080fd5b8a60208260051b8401011115610d7f57600080fd5b602082019350809250505092959891949750929550565b6020808252602a908201527f4f7267616e697a6174696f6e52656769737472793a206e6f7420616e2061646d60408201526934b734b9ba3930ba37b960b11b606082015260800190565b600181811c90821680610df457607f821691505b602082108103610e1457634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526031908201527f4f7267616e697a6174696f6e52656769737472793a206f7267616e697a6174696040820152701bdb88191bd95cc81b9bdd08195e1a5cdd607a1b606082015260800190565b60208082526033908201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c69642061646040820152726d696e6973747261746f72206164647265737360681b606082015260800190565b60208082526022908201527f4f7267616e697a6174696f6e52656769737472793a20696e76616c6964206e616040820152616d6560f01b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b601f821115610f6057806000526020600020601f840160051c81016020851015610f3d5750805b601f840160051c820191505b81811015610f5d5760008155600101610f49565b50505b505050565b67ffffffffffffffff831115610f7d57610f7d610f00565b610f9183610f8b8354610de0565b83610f16565b6000601f841160018114610fc55760008515610fad5750838201355b600019600387901b1c1916600186901b178355610f5d565b600083815260209020601f19861690835b82811015610ff65786850135825560209485019460019092019101610fd6565b50868210156110135760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff82168061105157611051611025565b6000190192915050565b634e487b7160e01b600052603260045260246000fd5b600063ffffffff821663ffffffff810361108d5761108d611025565b6001019291505056fea264697066735822122018ec8b7c61b34b841bd77f4f39528684570248481d630500ba9705e45287160d64736f6c634300081c0033",
}

// OrganizationRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OrganizationRegistryMetaData.ABI instead.
var OrganizationRegistryABI = OrganizationRegistryMetaData.ABI

// OrganizationRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrganizationRegistryMetaData.Bin instead.
var OrganizationRegistryBin = OrganizationRegistryMetaData.Bin

// DeployOrganizationRegistry deploys a new Ethereum contract, binding an instance of OrganizationRegistry to it.
func DeployOrganizationRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrganizationRegistry, error) {
	parsed, err := OrganizationRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrganizationRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrganizationRegistry{OrganizationRegistryCaller: OrganizationRegistryCaller{contract: contract}, OrganizationRegistryTransactor: OrganizationRegistryTransactor{contract: contract}, OrganizationRegistryFilterer: OrganizationRegistryFilterer{contract: contract}}, nil
}

// OrganizationRegistry is an auto generated Go binding around an Ethereum contract.
type OrganizationRegistry struct {
	OrganizationRegistryCaller     // Read-only binding to the contract
	OrganizationRegistryTransactor // Write-only binding to the contract
	OrganizationRegistryFilterer   // Log filterer for contract events
}

// OrganizationRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrganizationRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganizationRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganizationRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganizationRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganizationRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganizationRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganizationRegistrySession struct {
	Contract     *OrganizationRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OrganizationRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganizationRegistryCallerSession struct {
	Contract *OrganizationRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OrganizationRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganizationRegistryTransactorSession struct {
	Contract     *OrganizationRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OrganizationRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrganizationRegistryRaw struct {
	Contract *OrganizationRegistry // Generic contract binding to access the raw methods on
}

// OrganizationRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganizationRegistryCallerRaw struct {
	Contract *OrganizationRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// OrganizationRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganizationRegistryTransactorRaw struct {
	Contract *OrganizationRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganizationRegistry creates a new instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistry(address common.Address, backend bind.ContractBackend) (*OrganizationRegistry, error) {
	contract, err := bindOrganizationRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistry{OrganizationRegistryCaller: OrganizationRegistryCaller{contract: contract}, OrganizationRegistryTransactor: OrganizationRegistryTransactor{contract: contract}, OrganizationRegistryFilterer: OrganizationRegistryFilterer{contract: contract}}, nil
}

// NewOrganizationRegistryCaller creates a new read-only instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistryCaller(address common.Address, caller bind.ContractCaller) (*OrganizationRegistryCaller, error) {
	contract, err := bindOrganizationRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryCaller{contract: contract}, nil
}

// NewOrganizationRegistryTransactor creates a new write-only instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OrganizationRegistryTransactor, error) {
	contract, err := bindOrganizationRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryTransactor{contract: contract}, nil
}

// NewOrganizationRegistryFilterer creates a new log filterer instance of OrganizationRegistry, bound to a specific deployed contract.
func NewOrganizationRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OrganizationRegistryFilterer, error) {
	contract, err := bindOrganizationRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryFilterer{contract: contract}, nil
}

// bindOrganizationRegistry binds a generic wrapper to an already deployed contract.
func bindOrganizationRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrganizationRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganizationRegistry *OrganizationRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganizationRegistry.Contract.OrganizationRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganizationRegistry *OrganizationRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.OrganizationRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganizationRegistry *OrganizationRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.OrganizationRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganizationRegistry *OrganizationRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganizationRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganizationRegistry *OrganizationRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganizationRegistry *OrganizationRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns(string, string)
func (_OrganizationRegistry *OrganizationRegistryCaller) GetOrganization(opts *bind.CallOpts, id common.Address) (string, string, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "getOrganization", id)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns(string, string)
func (_OrganizationRegistry *OrganizationRegistrySession) GetOrganization(id common.Address) (string, string, error) {
	return _OrganizationRegistry.Contract.GetOrganization(&_OrganizationRegistry.CallOpts, id)
}

// GetOrganization is a free data retrieval call binding the contract method 0xd2c30a6d.
//
// Solidity: function getOrganization(address id) view returns(string, string)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) GetOrganization(id common.Address) (string, string, error) {
	return _OrganizationRegistry.Contract.GetOrganization(&_OrganizationRegistry.CallOpts, id)
}

// IsAdministrator is a free data retrieval call binding the contract method 0xc1af6e03.
//
// Solidity: function isAdministrator(address id, address account) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistryCaller) IsAdministrator(opts *bind.CallOpts, id common.Address, account common.Address) (bool, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "isAdministrator", id, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdministrator is a free data retrieval call binding the contract method 0xc1af6e03.
//
// Solidity: function isAdministrator(address id, address account) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistrySession) IsAdministrator(id common.Address, account common.Address) (bool, error) {
	return _OrganizationRegistry.Contract.IsAdministrator(&_OrganizationRegistry.CallOpts, id, account)
}

// IsAdministrator is a free data retrieval call binding the contract method 0xc1af6e03.
//
// Solidity: function isAdministrator(address id, address account) view returns(bool)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) IsAdministrator(id common.Address, account common.Address) (bool, error) {
	return _OrganizationRegistry.Contract.IsAdministrator(&_OrganizationRegistry.CallOpts, id, account)
}

// OrganizationCount is a free data retrieval call binding the contract method 0xf1c62104.
//
// Solidity: function organizationCount() view returns(uint32)
func (_OrganizationRegistry *OrganizationRegistryCaller) OrganizationCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "organizationCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OrganizationCount is a free data retrieval call binding the contract method 0xf1c62104.
//
// Solidity: function organizationCount() view returns(uint32)
func (_OrganizationRegistry *OrganizationRegistrySession) OrganizationCount() (uint32, error) {
	return _OrganizationRegistry.Contract.OrganizationCount(&_OrganizationRegistry.CallOpts)
}

// OrganizationCount is a free data retrieval call binding the contract method 0xf1c62104.
//
// Solidity: function organizationCount() view returns(uint32)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) OrganizationCount() (uint32, error) {
	return _OrganizationRegistry.Contract.OrganizationCount(&_OrganizationRegistry.CallOpts)
}

// Organizations is a free data retrieval call binding the contract method 0x5a1f7406.
//
// Solidity: function organizations(address ) view returns(string name, string metadataURI)
func (_OrganizationRegistry *OrganizationRegistryCaller) Organizations(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name        string
	MetadataURI string
}, error) {
	var out []interface{}
	err := _OrganizationRegistry.contract.Call(opts, &out, "organizations", arg0)

	outstruct := new(struct {
		Name        string
		MetadataURI string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.MetadataURI = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Organizations is a free data retrieval call binding the contract method 0x5a1f7406.
//
// Solidity: function organizations(address ) view returns(string name, string metadataURI)
func (_OrganizationRegistry *OrganizationRegistrySession) Organizations(arg0 common.Address) (struct {
	Name        string
	MetadataURI string
}, error) {
	return _OrganizationRegistry.Contract.Organizations(&_OrganizationRegistry.CallOpts, arg0)
}

// Organizations is a free data retrieval call binding the contract method 0x5a1f7406.
//
// Solidity: function organizations(address ) view returns(string name, string metadataURI)
func (_OrganizationRegistry *OrganizationRegistryCallerSession) Organizations(arg0 common.Address) (struct {
	Name        string
	MetadataURI string
}, error) {
	return _OrganizationRegistry.Contract.Organizations(&_OrganizationRegistry.CallOpts, arg0)
}

// AddAdministrator is a paid mutator transaction binding the contract method 0x1c2e3d82.
//
// Solidity: function addAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) AddAdministrator(opts *bind.TransactOpts, id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "addAdministrator", id, administrator)
}

// AddAdministrator is a paid mutator transaction binding the contract method 0x1c2e3d82.
//
// Solidity: function addAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) AddAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.AddAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// AddAdministrator is a paid mutator transaction binding the contract method 0x1c2e3d82.
//
// Solidity: function addAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) AddAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.AddAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xc2a950be.
//
// Solidity: function createOrganization(address id, string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) CreateOrganization(opts *bind.TransactOpts, id common.Address, name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "createOrganization", id, name, metadataURI, administrators)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xc2a950be.
//
// Solidity: function createOrganization(address id, string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) CreateOrganization(id common.Address, name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.CreateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI, administrators)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0xc2a950be.
//
// Solidity: function createOrganization(address id, string name, string metadataURI, address[] administrators) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) CreateOrganization(id common.Address, name string, metadataURI string, administrators []common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.CreateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI, administrators)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x7acbb8af.
//
// Solidity: function deleteOrganization(address id) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) DeleteOrganization(opts *bind.TransactOpts, id common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "deleteOrganization", id)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x7acbb8af.
//
// Solidity: function deleteOrganization(address id) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) DeleteOrganization(id common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.DeleteOrganization(&_OrganizationRegistry.TransactOpts, id)
}

// DeleteOrganization is a paid mutator transaction binding the contract method 0x7acbb8af.
//
// Solidity: function deleteOrganization(address id) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) DeleteOrganization(id common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.DeleteOrganization(&_OrganizationRegistry.TransactOpts, id)
}

// RemoveAdministrator is a paid mutator transaction binding the contract method 0x3c10eee5.
//
// Solidity: function removeAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) RemoveAdministrator(opts *bind.TransactOpts, id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "removeAdministrator", id, administrator)
}

// RemoveAdministrator is a paid mutator transaction binding the contract method 0x3c10eee5.
//
// Solidity: function removeAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) RemoveAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.RemoveAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// RemoveAdministrator is a paid mutator transaction binding the contract method 0x3c10eee5.
//
// Solidity: function removeAdministrator(address id, address administrator) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) RemoveAdministrator(id common.Address, administrator common.Address) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.RemoveAdministrator(&_OrganizationRegistry.TransactOpts, id, administrator)
}

// UpdateOrganization is a paid mutator transaction binding the contract method 0x6cca67bf.
//
// Solidity: function updateOrganization(address id, string name, string metadataURI) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactor) UpdateOrganization(opts *bind.TransactOpts, id common.Address, name string, metadataURI string) (*types.Transaction, error) {
	return _OrganizationRegistry.contract.Transact(opts, "updateOrganization", id, name, metadataURI)
}

// UpdateOrganization is a paid mutator transaction binding the contract method 0x6cca67bf.
//
// Solidity: function updateOrganization(address id, string name, string metadataURI) returns()
func (_OrganizationRegistry *OrganizationRegistrySession) UpdateOrganization(id common.Address, name string, metadataURI string) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.UpdateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI)
}

// UpdateOrganization is a paid mutator transaction binding the contract method 0x6cca67bf.
//
// Solidity: function updateOrganization(address id, string name, string metadataURI) returns()
func (_OrganizationRegistry *OrganizationRegistryTransactorSession) UpdateOrganization(id common.Address, name string, metadataURI string) (*types.Transaction, error) {
	return _OrganizationRegistry.Contract.UpdateOrganization(&_OrganizationRegistry.TransactOpts, id, name, metadataURI)
}

// OrganizationRegistryOrganizationCreatedIterator is returned from FilterOrganizationCreated and is used to iterate over the raw logs and unpacked data for OrganizationCreated events raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationCreatedIterator struct {
	Event *OrganizationRegistryOrganizationCreated // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryOrganizationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryOrganizationCreated)
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
		it.Event = new(OrganizationRegistryOrganizationCreated)
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
func (it *OrganizationRegistryOrganizationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryOrganizationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryOrganizationCreated represents a OrganizationCreated event raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationCreated struct {
	Id      common.Address
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrganizationCreated is a free log retrieval operation binding the contract event 0x2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a.
//
// Solidity: event OrganizationCreated(address indexed id, address indexed creator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterOrganizationCreated(opts *bind.FilterOpts, id []common.Address, creator []common.Address) (*OrganizationRegistryOrganizationCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "OrganizationCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryOrganizationCreatedIterator{contract: _OrganizationRegistry.contract, event: "OrganizationCreated", logs: logs, sub: sub}, nil
}

// WatchOrganizationCreated is a free log subscription operation binding the contract event 0x2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a.
//
// Solidity: event OrganizationCreated(address indexed id, address indexed creator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchOrganizationCreated(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryOrganizationCreated, id []common.Address, creator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "OrganizationCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryOrganizationCreated)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationCreated", log); err != nil {
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

// ParseOrganizationCreated is a log parse operation binding the contract event 0x2725ca0bb6f842da395a595005373aaa8e052b21133359b3c75f59a1247e6e7a.
//
// Solidity: event OrganizationCreated(address indexed id, address indexed creator)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseOrganizationCreated(log types.Log) (*OrganizationRegistryOrganizationCreated, error) {
	event := new(OrganizationRegistryOrganizationCreated)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganizationRegistryOrganizationUpdatedIterator is returned from FilterOrganizationUpdated and is used to iterate over the raw logs and unpacked data for OrganizationUpdated events raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationUpdatedIterator struct {
	Event *OrganizationRegistryOrganizationUpdated // Event containing the contract specifics and raw log

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
func (it *OrganizationRegistryOrganizationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganizationRegistryOrganizationUpdated)
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
		it.Event = new(OrganizationRegistryOrganizationUpdated)
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
func (it *OrganizationRegistryOrganizationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganizationRegistryOrganizationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganizationRegistryOrganizationUpdated represents a OrganizationUpdated event raised by the OrganizationRegistry contract.
type OrganizationRegistryOrganizationUpdated struct {
	Id      common.Address
	Updater common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrganizationUpdated is a free log retrieval operation binding the contract event 0xdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f6.
//
// Solidity: event OrganizationUpdated(address indexed id, address indexed updater)
func (_OrganizationRegistry *OrganizationRegistryFilterer) FilterOrganizationUpdated(opts *bind.FilterOpts, id []common.Address, updater []common.Address) (*OrganizationRegistryOrganizationUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.FilterLogs(opts, "OrganizationUpdated", idRule, updaterRule)
	if err != nil {
		return nil, err
	}
	return &OrganizationRegistryOrganizationUpdatedIterator{contract: _OrganizationRegistry.contract, event: "OrganizationUpdated", logs: logs, sub: sub}, nil
}

// WatchOrganizationUpdated is a free log subscription operation binding the contract event 0xdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f6.
//
// Solidity: event OrganizationUpdated(address indexed id, address indexed updater)
func (_OrganizationRegistry *OrganizationRegistryFilterer) WatchOrganizationUpdated(opts *bind.WatchOpts, sink chan<- *OrganizationRegistryOrganizationUpdated, id []common.Address, updater []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _OrganizationRegistry.contract.WatchLogs(opts, "OrganizationUpdated", idRule, updaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganizationRegistryOrganizationUpdated)
				if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationUpdated", log); err != nil {
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

// ParseOrganizationUpdated is a log parse operation binding the contract event 0xdcd663553eb7f5f57b83637c17d95d22a764affd6dbcc98f8ce9dcbac3e239f6.
//
// Solidity: event OrganizationUpdated(address indexed id, address indexed updater)
func (_OrganizationRegistry *OrganizationRegistryFilterer) ParseOrganizationUpdated(log types.Log) (*OrganizationRegistryOrganizationUpdated, error) {
	event := new(OrganizationRegistryOrganizationUpdated)
	if err := _OrganizationRegistry.contract.UnpackLog(event, "OrganizationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
