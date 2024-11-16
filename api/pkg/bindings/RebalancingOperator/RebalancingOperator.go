// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RebalancingOperator

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

// RebalancingOperatorMetaData contains all meta data concerning the RebalancingOperator contract.
var RebalancingOperatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_evc\",\"type\":\"address\",\"internalType\":\"contractIEVC\"},{\"name\":\"_asset\",\"type\":\"address\",\"internalType\":\"contractERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEVC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rebalanceOnBehalf\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"onBehalfOfAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// RebalancingOperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RebalancingOperatorMetaData.ABI instead.
var RebalancingOperatorABI = RebalancingOperatorMetaData.ABI

// RebalancingOperator is an auto generated Go binding around an Ethereum contract.
type RebalancingOperator struct {
	RebalancingOperatorCaller     // Read-only binding to the contract
	RebalancingOperatorTransactor // Write-only binding to the contract
	RebalancingOperatorFilterer   // Log filterer for contract events
}

// RebalancingOperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RebalancingOperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RebalancingOperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RebalancingOperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RebalancingOperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RebalancingOperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RebalancingOperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RebalancingOperatorSession struct {
	Contract     *RebalancingOperator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RebalancingOperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RebalancingOperatorCallerSession struct {
	Contract *RebalancingOperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RebalancingOperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RebalancingOperatorTransactorSession struct {
	Contract     *RebalancingOperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RebalancingOperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RebalancingOperatorRaw struct {
	Contract *RebalancingOperator // Generic contract binding to access the raw methods on
}

// RebalancingOperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RebalancingOperatorCallerRaw struct {
	Contract *RebalancingOperatorCaller // Generic read-only contract binding to access the raw methods on
}

// RebalancingOperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RebalancingOperatorTransactorRaw struct {
	Contract *RebalancingOperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRebalancingOperator creates a new instance of RebalancingOperator, bound to a specific deployed contract.
func NewRebalancingOperator(address common.Address, backend bind.ContractBackend) (*RebalancingOperator, error) {
	contract, err := bindRebalancingOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RebalancingOperator{RebalancingOperatorCaller: RebalancingOperatorCaller{contract: contract}, RebalancingOperatorTransactor: RebalancingOperatorTransactor{contract: contract}, RebalancingOperatorFilterer: RebalancingOperatorFilterer{contract: contract}}, nil
}

// NewRebalancingOperatorCaller creates a new read-only instance of RebalancingOperator, bound to a specific deployed contract.
func NewRebalancingOperatorCaller(address common.Address, caller bind.ContractCaller) (*RebalancingOperatorCaller, error) {
	contract, err := bindRebalancingOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RebalancingOperatorCaller{contract: contract}, nil
}

// NewRebalancingOperatorTransactor creates a new write-only instance of RebalancingOperator, bound to a specific deployed contract.
func NewRebalancingOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RebalancingOperatorTransactor, error) {
	contract, err := bindRebalancingOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RebalancingOperatorTransactor{contract: contract}, nil
}

// NewRebalancingOperatorFilterer creates a new log filterer instance of RebalancingOperator, bound to a specific deployed contract.
func NewRebalancingOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RebalancingOperatorFilterer, error) {
	contract, err := bindRebalancingOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RebalancingOperatorFilterer{contract: contract}, nil
}

// bindRebalancingOperator binds a generic wrapper to an already deployed contract.
func bindRebalancingOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RebalancingOperatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RebalancingOperator *RebalancingOperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RebalancingOperator.Contract.RebalancingOperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RebalancingOperator *RebalancingOperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.RebalancingOperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RebalancingOperator *RebalancingOperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.RebalancingOperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RebalancingOperator *RebalancingOperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RebalancingOperator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RebalancingOperator *RebalancingOperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RebalancingOperator *RebalancingOperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.contract.Transact(opts, method, params...)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_RebalancingOperator *RebalancingOperatorCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebalancingOperator.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_RebalancingOperator *RebalancingOperatorSession) Asset() (common.Address, error) {
	return _RebalancingOperator.Contract.Asset(&_RebalancingOperator.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_RebalancingOperator *RebalancingOperatorCallerSession) Asset() (common.Address, error) {
	return _RebalancingOperator.Contract.Asset(&_RebalancingOperator.CallOpts)
}

// Evc is a free data retrieval call binding the contract method 0x0fb411e8.
//
// Solidity: function evc() view returns(address)
func (_RebalancingOperator *RebalancingOperatorCaller) Evc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebalancingOperator.contract.Call(opts, &out, "evc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Evc is a free data retrieval call binding the contract method 0x0fb411e8.
//
// Solidity: function evc() view returns(address)
func (_RebalancingOperator *RebalancingOperatorSession) Evc() (common.Address, error) {
	return _RebalancingOperator.Contract.Evc(&_RebalancingOperator.CallOpts)
}

// Evc is a free data retrieval call binding the contract method 0x0fb411e8.
//
// Solidity: function evc() view returns(address)
func (_RebalancingOperator *RebalancingOperatorCallerSession) Evc() (common.Address, error) {
	return _RebalancingOperator.Contract.Evc(&_RebalancingOperator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RebalancingOperator *RebalancingOperatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebalancingOperator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RebalancingOperator *RebalancingOperatorSession) Owner() (common.Address, error) {
	return _RebalancingOperator.Contract.Owner(&_RebalancingOperator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RebalancingOperator *RebalancingOperatorCallerSession) Owner() (common.Address, error) {
	return _RebalancingOperator.Contract.Owner(&_RebalancingOperator.CallOpts)
}

// RebalanceOnBehalf is a paid mutator transaction binding the contract method 0x56d84ffb.
//
// Solidity: function rebalanceOnBehalf(address from, address to, address onBehalfOfAccount) returns()
func (_RebalancingOperator *RebalancingOperatorTransactor) RebalanceOnBehalf(opts *bind.TransactOpts, from common.Address, to common.Address, onBehalfOfAccount common.Address) (*types.Transaction, error) {
	return _RebalancingOperator.contract.Transact(opts, "rebalanceOnBehalf", from, to, onBehalfOfAccount)
}

// RebalanceOnBehalf is a paid mutator transaction binding the contract method 0x56d84ffb.
//
// Solidity: function rebalanceOnBehalf(address from, address to, address onBehalfOfAccount) returns()
func (_RebalancingOperator *RebalancingOperatorSession) RebalanceOnBehalf(from common.Address, to common.Address, onBehalfOfAccount common.Address) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.RebalanceOnBehalf(&_RebalancingOperator.TransactOpts, from, to, onBehalfOfAccount)
}

// RebalanceOnBehalf is a paid mutator transaction binding the contract method 0x56d84ffb.
//
// Solidity: function rebalanceOnBehalf(address from, address to, address onBehalfOfAccount) returns()
func (_RebalancingOperator *RebalancingOperatorTransactorSession) RebalanceOnBehalf(from common.Address, to common.Address, onBehalfOfAccount common.Address) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.RebalanceOnBehalf(&_RebalancingOperator.TransactOpts, from, to, onBehalfOfAccount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RebalancingOperator *RebalancingOperatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebalancingOperator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RebalancingOperator *RebalancingOperatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RebalancingOperator.Contract.RenounceOwnership(&_RebalancingOperator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RebalancingOperator *RebalancingOperatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RebalancingOperator.Contract.RenounceOwnership(&_RebalancingOperator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RebalancingOperator *RebalancingOperatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RebalancingOperator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RebalancingOperator *RebalancingOperatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.TransferOwnership(&_RebalancingOperator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RebalancingOperator *RebalancingOperatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RebalancingOperator.Contract.TransferOwnership(&_RebalancingOperator.TransactOpts, newOwner)
}

// RebalancingOperatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RebalancingOperator contract.
type RebalancingOperatorOwnershipTransferredIterator struct {
	Event *RebalancingOperatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RebalancingOperatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebalancingOperatorOwnershipTransferred)
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
		it.Event = new(RebalancingOperatorOwnershipTransferred)
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
func (it *RebalancingOperatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebalancingOperatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebalancingOperatorOwnershipTransferred represents a OwnershipTransferred event raised by the RebalancingOperator contract.
type RebalancingOperatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RebalancingOperator *RebalancingOperatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RebalancingOperatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RebalancingOperator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RebalancingOperatorOwnershipTransferredIterator{contract: _RebalancingOperator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RebalancingOperator *RebalancingOperatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RebalancingOperatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RebalancingOperator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebalancingOperatorOwnershipTransferred)
				if err := _RebalancingOperator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RebalancingOperator *RebalancingOperatorFilterer) ParseOwnershipTransferred(log types.Log) (*RebalancingOperatorOwnershipTransferred, error) {
	event := new(RebalancingOperatorOwnershipTransferred)
	if err := _RebalancingOperator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
