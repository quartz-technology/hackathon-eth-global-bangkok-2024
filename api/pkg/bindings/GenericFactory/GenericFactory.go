// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericFactory

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

// GenericFactoryProxyConfig is an auto generated low-level Go binding around an user-defined struct.
type GenericFactoryProxyConfig struct {
	Upgradeable    bool
	Implementation common.Address
	TrailingData   []byte
}

// GenericFactortMetaData contains all meta data concerning the GenericFactort contract.
var GenericFactortMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createProxy\",\"inputs\":[{\"name\":\"desiredImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"upgradeable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"trailingData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProxyConfig\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structGenericFactory.ProxyConfig\",\"components\":[{\"name\":\"upgradeable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"trailingData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProxyListLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProxyListSlice\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"}]",
}

// GenericFactortABI is the input ABI used to generate the binding from.
// Deprecated: Use GenericFactortMetaData.ABI instead.
var GenericFactortABI = GenericFactortMetaData.ABI

// GenericFactort is an auto generated Go binding around an Ethereum contract.
type GenericFactort struct {
	GenericFactortCaller     // Read-only binding to the contract
	GenericFactortTransactor // Write-only binding to the contract
	GenericFactortFilterer   // Log filterer for contract events
}

// GenericFactortCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericFactortCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericFactortTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericFactortTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericFactortFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericFactortFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericFactortSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericFactortSession struct {
	Contract     *GenericFactort   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericFactortCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericFactortCallerSession struct {
	Contract *GenericFactortCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GenericFactortTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericFactortTransactorSession struct {
	Contract     *GenericFactortTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GenericFactortRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericFactortRaw struct {
	Contract *GenericFactort // Generic contract binding to access the raw methods on
}

// GenericFactortCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericFactortCallerRaw struct {
	Contract *GenericFactortCaller // Generic read-only contract binding to access the raw methods on
}

// GenericFactortTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericFactortTransactorRaw struct {
	Contract *GenericFactortTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericFactort creates a new instance of GenericFactort, bound to a specific deployed contract.
func NewGenericFactort(address common.Address, backend bind.ContractBackend) (*GenericFactort, error) {
	contract, err := bindGenericFactort(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericFactort{GenericFactortCaller: GenericFactortCaller{contract: contract}, GenericFactortTransactor: GenericFactortTransactor{contract: contract}, GenericFactortFilterer: GenericFactortFilterer{contract: contract}}, nil
}

// NewGenericFactortCaller creates a new read-only instance of GenericFactort, bound to a specific deployed contract.
func NewGenericFactortCaller(address common.Address, caller bind.ContractCaller) (*GenericFactortCaller, error) {
	contract, err := bindGenericFactort(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericFactortCaller{contract: contract}, nil
}

// NewGenericFactortTransactor creates a new write-only instance of GenericFactort, bound to a specific deployed contract.
func NewGenericFactortTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericFactortTransactor, error) {
	contract, err := bindGenericFactort(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericFactortTransactor{contract: contract}, nil
}

// NewGenericFactortFilterer creates a new log filterer instance of GenericFactort, bound to a specific deployed contract.
func NewGenericFactortFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericFactortFilterer, error) {
	contract, err := bindGenericFactort(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericFactortFilterer{contract: contract}, nil
}

// bindGenericFactort binds a generic wrapper to an already deployed contract.
func bindGenericFactort(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GenericFactortMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericFactort *GenericFactortRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericFactort.Contract.GenericFactortCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericFactort *GenericFactortRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericFactort.Contract.GenericFactortTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericFactort *GenericFactortRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericFactort.Contract.GenericFactortTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericFactort *GenericFactortCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericFactort.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericFactort *GenericFactortTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericFactort.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericFactort *GenericFactortTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericFactort.Contract.contract.Transact(opts, method, params...)
}

// GetProxyConfig is a free data retrieval call binding the contract method 0xa20ea5c1.
//
// Solidity: function getProxyConfig(address proxy) view returns((bool,address,bytes) config)
func (_GenericFactort *GenericFactortCaller) GetProxyConfig(opts *bind.CallOpts, proxy common.Address) (GenericFactoryProxyConfig, error) {
	var out []interface{}
	err := _GenericFactort.contract.Call(opts, &out, "getProxyConfig", proxy)

	if err != nil {
		return *new(GenericFactoryProxyConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(GenericFactoryProxyConfig)).(*GenericFactoryProxyConfig)

	return out0, err

}

// GetProxyConfig is a free data retrieval call binding the contract method 0xa20ea5c1.
//
// Solidity: function getProxyConfig(address proxy) view returns((bool,address,bytes) config)
func (_GenericFactort *GenericFactortSession) GetProxyConfig(proxy common.Address) (GenericFactoryProxyConfig, error) {
	return _GenericFactort.Contract.GetProxyConfig(&_GenericFactort.CallOpts, proxy)
}

// GetProxyConfig is a free data retrieval call binding the contract method 0xa20ea5c1.
//
// Solidity: function getProxyConfig(address proxy) view returns((bool,address,bytes) config)
func (_GenericFactort *GenericFactortCallerSession) GetProxyConfig(proxy common.Address) (GenericFactoryProxyConfig, error) {
	return _GenericFactort.Contract.GetProxyConfig(&_GenericFactort.CallOpts, proxy)
}

// GetProxyListLength is a free data retrieval call binding the contract method 0x0a68b7ba.
//
// Solidity: function getProxyListLength() view returns(uint256)
func (_GenericFactort *GenericFactortCaller) GetProxyListLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GenericFactort.contract.Call(opts, &out, "getProxyListLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProxyListLength is a free data retrieval call binding the contract method 0x0a68b7ba.
//
// Solidity: function getProxyListLength() view returns(uint256)
func (_GenericFactort *GenericFactortSession) GetProxyListLength() (*big.Int, error) {
	return _GenericFactort.Contract.GetProxyListLength(&_GenericFactort.CallOpts)
}

// GetProxyListLength is a free data retrieval call binding the contract method 0x0a68b7ba.
//
// Solidity: function getProxyListLength() view returns(uint256)
func (_GenericFactort *GenericFactortCallerSession) GetProxyListLength() (*big.Int, error) {
	return _GenericFactort.Contract.GetProxyListLength(&_GenericFactort.CallOpts)
}

// GetProxyListSlice is a free data retrieval call binding the contract method 0xc0e96df6.
//
// Solidity: function getProxyListSlice(uint256 start, uint256 end) view returns(address[] list)
func (_GenericFactort *GenericFactortCaller) GetProxyListSlice(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _GenericFactort.contract.Call(opts, &out, "getProxyListSlice", start, end)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetProxyListSlice is a free data retrieval call binding the contract method 0xc0e96df6.
//
// Solidity: function getProxyListSlice(uint256 start, uint256 end) view returns(address[] list)
func (_GenericFactort *GenericFactortSession) GetProxyListSlice(start *big.Int, end *big.Int) ([]common.Address, error) {
	return _GenericFactort.Contract.GetProxyListSlice(&_GenericFactort.CallOpts, start, end)
}

// GetProxyListSlice is a free data retrieval call binding the contract method 0xc0e96df6.
//
// Solidity: function getProxyListSlice(uint256 start, uint256 end) view returns(address[] list)
func (_GenericFactort *GenericFactortCallerSession) GetProxyListSlice(start *big.Int, end *big.Int) ([]common.Address, error) {
	return _GenericFactort.Contract.GetProxyListSlice(&_GenericFactort.CallOpts, start, end)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x83e85b27.
//
// Solidity: function createProxy(address desiredImplementation, bool upgradeable, bytes trailingData) returns(address)
func (_GenericFactort *GenericFactortTransactor) CreateProxy(opts *bind.TransactOpts, desiredImplementation common.Address, upgradeable bool, trailingData []byte) (*types.Transaction, error) {
	return _GenericFactort.contract.Transact(opts, "createProxy", desiredImplementation, upgradeable, trailingData)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x83e85b27.
//
// Solidity: function createProxy(address desiredImplementation, bool upgradeable, bytes trailingData) returns(address)
func (_GenericFactort *GenericFactortSession) CreateProxy(desiredImplementation common.Address, upgradeable bool, trailingData []byte) (*types.Transaction, error) {
	return _GenericFactort.Contract.CreateProxy(&_GenericFactort.TransactOpts, desiredImplementation, upgradeable, trailingData)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x83e85b27.
//
// Solidity: function createProxy(address desiredImplementation, bool upgradeable, bytes trailingData) returns(address)
func (_GenericFactort *GenericFactortTransactorSession) CreateProxy(desiredImplementation common.Address, upgradeable bool, trailingData []byte) (*types.Transaction, error) {
	return _GenericFactort.Contract.CreateProxy(&_GenericFactort.TransactOpts, desiredImplementation, upgradeable, trailingData)
}
