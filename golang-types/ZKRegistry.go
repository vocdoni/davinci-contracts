// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZKRegistry

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

// ZKRegistryMetaData contains all meta data concerning the ZKRegistry contract.
var ZKRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DeleteAccountSIK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sik\",\"type\":\"bytes32\"}],\"name\":\"SetAccountSIK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountSIK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ZKRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKRegistryMetaData.ABI instead.
var ZKRegistryABI = ZKRegistryMetaData.ABI

// ZKRegistry is an auto generated Go binding around an Ethereum contract.
type ZKRegistry struct {
	ZKRegistryCaller     // Read-only binding to the contract
	ZKRegistryTransactor // Write-only binding to the contract
	ZKRegistryFilterer   // Log filterer for contract events
}

// ZKRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKRegistrySession struct {
	Contract     *ZKRegistry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKRegistryCallerSession struct {
	Contract *ZKRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZKRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKRegistryTransactorSession struct {
	Contract     *ZKRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZKRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKRegistryRaw struct {
	Contract *ZKRegistry // Generic contract binding to access the raw methods on
}

// ZKRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKRegistryCallerRaw struct {
	Contract *ZKRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ZKRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKRegistryTransactorRaw struct {
	Contract *ZKRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKRegistry creates a new instance of ZKRegistry, bound to a specific deployed contract.
func NewZKRegistry(address common.Address, backend bind.ContractBackend) (*ZKRegistry, error) {
	contract, err := bindZKRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKRegistry{ZKRegistryCaller: ZKRegistryCaller{contract: contract}, ZKRegistryTransactor: ZKRegistryTransactor{contract: contract}, ZKRegistryFilterer: ZKRegistryFilterer{contract: contract}}, nil
}

// NewZKRegistryCaller creates a new read-only instance of ZKRegistry, bound to a specific deployed contract.
func NewZKRegistryCaller(address common.Address, caller bind.ContractCaller) (*ZKRegistryCaller, error) {
	contract, err := bindZKRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKRegistryCaller{contract: contract}, nil
}

// NewZKRegistryTransactor creates a new write-only instance of ZKRegistry, bound to a specific deployed contract.
func NewZKRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKRegistryTransactor, error) {
	contract, err := bindZKRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKRegistryTransactor{contract: contract}, nil
}

// NewZKRegistryFilterer creates a new log filterer instance of ZKRegistry, bound to a specific deployed contract.
func NewZKRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKRegistryFilterer, error) {
	contract, err := bindZKRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKRegistryFilterer{contract: contract}, nil
}

// bindZKRegistry binds a generic wrapper to an already deployed contract.
func bindZKRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKRegistry *ZKRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKRegistry.Contract.ZKRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKRegistry *ZKRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKRegistry.Contract.ZKRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKRegistry *ZKRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKRegistry.Contract.ZKRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKRegistry *ZKRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKRegistry *ZKRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKRegistry *ZKRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZKRegistry *ZKRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZKRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZKRegistry *ZKRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZKRegistry.Contract.UPGRADEINTERFACEVERSION(&_ZKRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZKRegistry *ZKRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZKRegistry.Contract.UPGRADEINTERFACEVERSION(&_ZKRegistry.CallOpts)
}

// AccountSIK is a free data retrieval call binding the contract method 0x608e1880.
//
// Solidity: function accountSIK(address ) view returns(bytes32)
func (_ZKRegistry *ZKRegistryCaller) AccountSIK(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ZKRegistry.contract.Call(opts, &out, "accountSIK", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccountSIK is a free data retrieval call binding the contract method 0x608e1880.
//
// Solidity: function accountSIK(address ) view returns(bytes32)
func (_ZKRegistry *ZKRegistrySession) AccountSIK(arg0 common.Address) ([32]byte, error) {
	return _ZKRegistry.Contract.AccountSIK(&_ZKRegistry.CallOpts, arg0)
}

// AccountSIK is a free data retrieval call binding the contract method 0x608e1880.
//
// Solidity: function accountSIK(address ) view returns(bytes32)
func (_ZKRegistry *ZKRegistryCallerSession) AccountSIK(arg0 common.Address) ([32]byte, error) {
	return _ZKRegistry.Contract.AccountSIK(&_ZKRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKRegistry *ZKRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKRegistry *ZKRegistrySession) Owner() (common.Address, error) {
	return _ZKRegistry.Contract.Owner(&_ZKRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKRegistry *ZKRegistryCallerSession) Owner() (common.Address, error) {
	return _ZKRegistry.Contract.Owner(&_ZKRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZKRegistry *ZKRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZKRegistry *ZKRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ZKRegistry.Contract.ProxiableUUID(&_ZKRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZKRegistry *ZKRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ZKRegistry.Contract.ProxiableUUID(&_ZKRegistry.CallOpts)
}

// DeleteAccountSIK is a paid mutator transaction binding the contract method 0x40723176.
//
// Solidity: function DeleteAccountSIK() returns()
func (_ZKRegistry *ZKRegistryTransactor) DeleteAccountSIK(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKRegistry.contract.Transact(opts, "DeleteAccountSIK")
}

// DeleteAccountSIK is a paid mutator transaction binding the contract method 0x40723176.
//
// Solidity: function DeleteAccountSIK() returns()
func (_ZKRegistry *ZKRegistrySession) DeleteAccountSIK() (*types.Transaction, error) {
	return _ZKRegistry.Contract.DeleteAccountSIK(&_ZKRegistry.TransactOpts)
}

// DeleteAccountSIK is a paid mutator transaction binding the contract method 0x40723176.
//
// Solidity: function DeleteAccountSIK() returns()
func (_ZKRegistry *ZKRegistryTransactorSession) DeleteAccountSIK() (*types.Transaction, error) {
	return _ZKRegistry.Contract.DeleteAccountSIK(&_ZKRegistry.TransactOpts)
}

// SetAccountSIK is a paid mutator transaction binding the contract method 0xdc994e67.
//
// Solidity: function SetAccountSIK(bytes32 _sik) returns()
func (_ZKRegistry *ZKRegistryTransactor) SetAccountSIK(opts *bind.TransactOpts, _sik [32]byte) (*types.Transaction, error) {
	return _ZKRegistry.contract.Transact(opts, "SetAccountSIK", _sik)
}

// SetAccountSIK is a paid mutator transaction binding the contract method 0xdc994e67.
//
// Solidity: function SetAccountSIK(bytes32 _sik) returns()
func (_ZKRegistry *ZKRegistrySession) SetAccountSIK(_sik [32]byte) (*types.Transaction, error) {
	return _ZKRegistry.Contract.SetAccountSIK(&_ZKRegistry.TransactOpts, _sik)
}

// SetAccountSIK is a paid mutator transaction binding the contract method 0xdc994e67.
//
// Solidity: function SetAccountSIK(bytes32 _sik) returns()
func (_ZKRegistry *ZKRegistryTransactorSession) SetAccountSIK(_sik [32]byte) (*types.Transaction, error) {
	return _ZKRegistry.Contract.SetAccountSIK(&_ZKRegistry.TransactOpts, _sik)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ZKRegistry *ZKRegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKRegistry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ZKRegistry *ZKRegistrySession) Initialize() (*types.Transaction, error) {
	return _ZKRegistry.Contract.Initialize(&_ZKRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ZKRegistry *ZKRegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _ZKRegistry.Contract.Initialize(&_ZKRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZKRegistry *ZKRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZKRegistry *ZKRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKRegistry.Contract.RenounceOwnership(&_ZKRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZKRegistry *ZKRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKRegistry.Contract.RenounceOwnership(&_ZKRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZKRegistry *ZKRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZKRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZKRegistry *ZKRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKRegistry.Contract.TransferOwnership(&_ZKRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZKRegistry *ZKRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKRegistry.Contract.TransferOwnership(&_ZKRegistry.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZKRegistry *ZKRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZKRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZKRegistry *ZKRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZKRegistry.Contract.UpgradeToAndCall(&_ZKRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZKRegistry *ZKRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZKRegistry.Contract.UpgradeToAndCall(&_ZKRegistry.TransactOpts, newImplementation, data)
}

// ZKRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZKRegistry contract.
type ZKRegistryInitializedIterator struct {
	Event *ZKRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ZKRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKRegistryInitialized)
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
		it.Event = new(ZKRegistryInitialized)
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
func (it *ZKRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKRegistryInitialized represents a Initialized event raised by the ZKRegistry contract.
type ZKRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZKRegistry *ZKRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZKRegistryInitializedIterator, error) {

	logs, sub, err := _ZKRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZKRegistryInitializedIterator{contract: _ZKRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZKRegistry *ZKRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZKRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ZKRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKRegistryInitialized)
				if err := _ZKRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ZKRegistry *ZKRegistryFilterer) ParseInitialized(log types.Log) (*ZKRegistryInitialized, error) {
	event := new(ZKRegistryInitialized)
	if err := _ZKRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZKRegistry contract.
type ZKRegistryOwnershipTransferredIterator struct {
	Event *ZKRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZKRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKRegistryOwnershipTransferred)
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
		it.Event = new(ZKRegistryOwnershipTransferred)
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
func (it *ZKRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ZKRegistry contract.
type ZKRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZKRegistry *ZKRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZKRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKRegistryOwnershipTransferredIterator{contract: _ZKRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZKRegistry *ZKRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZKRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKRegistryOwnershipTransferred)
				if err := _ZKRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZKRegistry *ZKRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ZKRegistryOwnershipTransferred, error) {
	event := new(ZKRegistryOwnershipTransferred)
	if err := _ZKRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ZKRegistry contract.
type ZKRegistryUpgradedIterator struct {
	Event *ZKRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ZKRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKRegistryUpgraded)
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
		it.Event = new(ZKRegistryUpgraded)
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
func (it *ZKRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKRegistryUpgraded represents a Upgraded event raised by the ZKRegistry contract.
type ZKRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZKRegistry *ZKRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZKRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZKRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZKRegistryUpgradedIterator{contract: _ZKRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZKRegistry *ZKRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZKRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZKRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKRegistryUpgraded)
				if err := _ZKRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ZKRegistry *ZKRegistryFilterer) ParseUpgraded(log types.Log) (*ZKRegistryUpgraded, error) {
	event := new(ZKRegistryUpgraded)
	if err := _ZKRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
