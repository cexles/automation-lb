// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sample_test

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

// SampleTestContractMetaData contains all meta data concerning the SampleTestContract contract.
var SampleTestContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newElements\",\"type\":\"uint256\"}],\"name\":\"addElementsToArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"array\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkMaxIterations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"checkOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performLimit\",\"type\":\"uint256\"}],\"name\":\"encodeParamsForCheck\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"checkArray\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"performOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performLimit\",\"type\":\"uint256\"}],\"name\":\"encodeParamsForPerform\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUpkeepNeeded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performMaxIterations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"name\":\"performUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLength\",\"type\":\"uint256\"}],\"name\":\"resizeArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_checkMaxIterations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_performMaxIterations\",\"type\":\"uint256\"}],\"name\":\"setMaxIterations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SampleTestContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SampleTestContractMetaData.ABI instead.
var SampleTestContractABI = SampleTestContractMetaData.ABI

// SampleTestContract is an auto generated Go binding around an Ethereum contract.
type SampleTestContract struct {
	SampleTestContractCaller     // Read-only binding to the contract
	SampleTestContractTransactor // Write-only binding to the contract
	SampleTestContractFilterer   // Log filterer for contract events
}

// SampleTestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleTestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleTestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleTestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleTestContractSession struct {
	Contract     *SampleTestContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SampleTestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleTestContractCallerSession struct {
	Contract *SampleTestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SampleTestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleTestContractTransactorSession struct {
	Contract     *SampleTestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SampleTestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleTestContractRaw struct {
	Contract *SampleTestContract // Generic contract binding to access the raw methods on
}

// SampleTestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleTestContractCallerRaw struct {
	Contract *SampleTestContractCaller // Generic read-only contract binding to access the raw methods on
}

// SampleTestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleTestContractTransactorRaw struct {
	Contract *SampleTestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleTestContract creates a new instance of SampleTestContract, bound to a specific deployed contract.
func NewSampleTestContract(address common.Address, backend bind.ContractBackend) (*SampleTestContract, error) {
	contract, err := bindSampleTestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleTestContract{SampleTestContractCaller: SampleTestContractCaller{contract: contract}, SampleTestContractTransactor: SampleTestContractTransactor{contract: contract}, SampleTestContractFilterer: SampleTestContractFilterer{contract: contract}}, nil
}

// NewSampleTestContractCaller creates a new read-only instance of SampleTestContract, bound to a specific deployed contract.
func NewSampleTestContractCaller(address common.Address, caller bind.ContractCaller) (*SampleTestContractCaller, error) {
	contract, err := bindSampleTestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleTestContractCaller{contract: contract}, nil
}

// NewSampleTestContractTransactor creates a new write-only instance of SampleTestContract, bound to a specific deployed contract.
func NewSampleTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleTestContractTransactor, error) {
	contract, err := bindSampleTestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleTestContractTransactor{contract: contract}, nil
}

// NewSampleTestContractFilterer creates a new log filterer instance of SampleTestContract, bound to a specific deployed contract.
func NewSampleTestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleTestContractFilterer, error) {
	contract, err := bindSampleTestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleTestContractFilterer{contract: contract}, nil
}

// bindSampleTestContract binds a generic wrapper to an already deployed contract.
func bindSampleTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SampleTestContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleTestContract *SampleTestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SampleTestContract.Contract.SampleTestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleTestContract *SampleTestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleTestContract.Contract.SampleTestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleTestContract *SampleTestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleTestContract.Contract.SampleTestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleTestContract *SampleTestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SampleTestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleTestContract *SampleTestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleTestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleTestContract *SampleTestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleTestContract.Contract.contract.Transact(opts, method, params...)
}

// Array is a free data retrieval call binding the contract method 0x38d94193.
//
// Solidity: function array(uint256 ) view returns(uint256)
func (_SampleTestContract *SampleTestContractCaller) Array(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "array", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Array is a free data retrieval call binding the contract method 0x38d94193.
//
// Solidity: function array(uint256 ) view returns(uint256)
func (_SampleTestContract *SampleTestContractSession) Array(arg0 *big.Int) (*big.Int, error) {
	return _SampleTestContract.Contract.Array(&_SampleTestContract.CallOpts, arg0)
}

// Array is a free data retrieval call binding the contract method 0x38d94193.
//
// Solidity: function array(uint256 ) view returns(uint256)
func (_SampleTestContract *SampleTestContractCallerSession) Array(arg0 *big.Int) (*big.Int, error) {
	return _SampleTestContract.Contract.Array(&_SampleTestContract.CallOpts, arg0)
}

// CheckMaxIterations is a free data retrieval call binding the contract method 0xf9f9f7ef.
//
// Solidity: function checkMaxIterations() view returns(uint256)
func (_SampleTestContract *SampleTestContractCaller) CheckMaxIterations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "checkMaxIterations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckMaxIterations is a free data retrieval call binding the contract method 0xf9f9f7ef.
//
// Solidity: function checkMaxIterations() view returns(uint256)
func (_SampleTestContract *SampleTestContractSession) CheckMaxIterations() (*big.Int, error) {
	return _SampleTestContract.Contract.CheckMaxIterations(&_SampleTestContract.CallOpts)
}

// CheckMaxIterations is a free data retrieval call binding the contract method 0xf9f9f7ef.
//
// Solidity: function checkMaxIterations() view returns(uint256)
func (_SampleTestContract *SampleTestContractCallerSession) CheckMaxIterations() (*big.Int, error) {
	return _SampleTestContract.Contract.CheckMaxIterations(&_SampleTestContract.CallOpts)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) view returns(bool upkeepNeeded, bytes performData)
func (_SampleTestContract *SampleTestContractCaller) CheckUpkeep(opts *bind.CallOpts, checkData []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "checkUpkeep", checkData)

	outstruct := new(struct {
		UpkeepNeeded bool
		PerformData  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpkeepNeeded = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PerformData = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) view returns(bool upkeepNeeded, bytes performData)
func (_SampleTestContract *SampleTestContractSession) CheckUpkeep(checkData []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _SampleTestContract.Contract.CheckUpkeep(&_SampleTestContract.CallOpts, checkData)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) view returns(bool upkeepNeeded, bytes performData)
func (_SampleTestContract *SampleTestContractCallerSession) CheckUpkeep(checkData []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _SampleTestContract.Contract.CheckUpkeep(&_SampleTestContract.CallOpts, checkData)
}

// EncodeParamsForCheck is a free data retrieval call binding the contract method 0xb047ce2d.
//
// Solidity: function encodeParamsForCheck(uint256 checkOffset, uint256 checkLimit, uint256 performOffset, uint256 performLimit) pure returns(bytes)
func (_SampleTestContract *SampleTestContractCaller) EncodeParamsForCheck(opts *bind.CallOpts, checkOffset *big.Int, checkLimit *big.Int, performOffset *big.Int, performLimit *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "encodeParamsForCheck", checkOffset, checkLimit, performOffset, performLimit)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeParamsForCheck is a free data retrieval call binding the contract method 0xb047ce2d.
//
// Solidity: function encodeParamsForCheck(uint256 checkOffset, uint256 checkLimit, uint256 performOffset, uint256 performLimit) pure returns(bytes)
func (_SampleTestContract *SampleTestContractSession) EncodeParamsForCheck(checkOffset *big.Int, checkLimit *big.Int, performOffset *big.Int, performLimit *big.Int) ([]byte, error) {
	return _SampleTestContract.Contract.EncodeParamsForCheck(&_SampleTestContract.CallOpts, checkOffset, checkLimit, performOffset, performLimit)
}

// EncodeParamsForCheck is a free data retrieval call binding the contract method 0xb047ce2d.
//
// Solidity: function encodeParamsForCheck(uint256 checkOffset, uint256 checkLimit, uint256 performOffset, uint256 performLimit) pure returns(bytes)
func (_SampleTestContract *SampleTestContractCallerSession) EncodeParamsForCheck(checkOffset *big.Int, checkLimit *big.Int, performOffset *big.Int, performLimit *big.Int) ([]byte, error) {
	return _SampleTestContract.Contract.EncodeParamsForCheck(&_SampleTestContract.CallOpts, checkOffset, checkLimit, performOffset, performLimit)
}

// EncodeParamsForPerform is a free data retrieval call binding the contract method 0xa25c54ea.
//
// Solidity: function encodeParamsForPerform(uint256[] checkArray, uint256 performOffset, uint256 performLimit) pure returns(bytes)
func (_SampleTestContract *SampleTestContractCaller) EncodeParamsForPerform(opts *bind.CallOpts, checkArray []*big.Int, performOffset *big.Int, performLimit *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "encodeParamsForPerform", checkArray, performOffset, performLimit)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeParamsForPerform is a free data retrieval call binding the contract method 0xa25c54ea.
//
// Solidity: function encodeParamsForPerform(uint256[] checkArray, uint256 performOffset, uint256 performLimit) pure returns(bytes)
func (_SampleTestContract *SampleTestContractSession) EncodeParamsForPerform(checkArray []*big.Int, performOffset *big.Int, performLimit *big.Int) ([]byte, error) {
	return _SampleTestContract.Contract.EncodeParamsForPerform(&_SampleTestContract.CallOpts, checkArray, performOffset, performLimit)
}

// EncodeParamsForPerform is a free data retrieval call binding the contract method 0xa25c54ea.
//
// Solidity: function encodeParamsForPerform(uint256[] checkArray, uint256 performOffset, uint256 performLimit) pure returns(bytes)
func (_SampleTestContract *SampleTestContractCallerSession) EncodeParamsForPerform(checkArray []*big.Int, performOffset *big.Int, performLimit *big.Int) ([]byte, error) {
	return _SampleTestContract.Contract.EncodeParamsForPerform(&_SampleTestContract.CallOpts, checkArray, performOffset, performLimit)
}

// GetArray is a free data retrieval call binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() view returns(uint256[])
func (_SampleTestContract *SampleTestContractCaller) GetArray(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "getArray")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetArray is a free data retrieval call binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() view returns(uint256[])
func (_SampleTestContract *SampleTestContractSession) GetArray() ([]*big.Int, error) {
	return _SampleTestContract.Contract.GetArray(&_SampleTestContract.CallOpts)
}

// GetArray is a free data retrieval call binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() view returns(uint256[])
func (_SampleTestContract *SampleTestContractCallerSession) GetArray() ([]*big.Int, error) {
	return _SampleTestContract.Contract.GetArray(&_SampleTestContract.CallOpts)
}

// IsUpkeepNeeded is a free data retrieval call binding the contract method 0x9ed7b2bc.
//
// Solidity: function isUpkeepNeeded() view returns(bool)
func (_SampleTestContract *SampleTestContractCaller) IsUpkeepNeeded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "isUpkeepNeeded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpkeepNeeded is a free data retrieval call binding the contract method 0x9ed7b2bc.
//
// Solidity: function isUpkeepNeeded() view returns(bool)
func (_SampleTestContract *SampleTestContractSession) IsUpkeepNeeded() (bool, error) {
	return _SampleTestContract.Contract.IsUpkeepNeeded(&_SampleTestContract.CallOpts)
}

// IsUpkeepNeeded is a free data retrieval call binding the contract method 0x9ed7b2bc.
//
// Solidity: function isUpkeepNeeded() view returns(bool)
func (_SampleTestContract *SampleTestContractCallerSession) IsUpkeepNeeded() (bool, error) {
	return _SampleTestContract.Contract.IsUpkeepNeeded(&_SampleTestContract.CallOpts)
}

// PerformMaxIterations is a free data retrieval call binding the contract method 0xcd61796a.
//
// Solidity: function performMaxIterations() view returns(uint256)
func (_SampleTestContract *SampleTestContractCaller) PerformMaxIterations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "performMaxIterations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformMaxIterations is a free data retrieval call binding the contract method 0xcd61796a.
//
// Solidity: function performMaxIterations() view returns(uint256)
func (_SampleTestContract *SampleTestContractSession) PerformMaxIterations() (*big.Int, error) {
	return _SampleTestContract.Contract.PerformMaxIterations(&_SampleTestContract.CallOpts)
}

// PerformMaxIterations is a free data retrieval call binding the contract method 0xcd61796a.
//
// Solidity: function performMaxIterations() view returns(uint256)
func (_SampleTestContract *SampleTestContractCallerSession) PerformMaxIterations() (*big.Int, error) {
	return _SampleTestContract.Contract.PerformMaxIterations(&_SampleTestContract.CallOpts)
}

// AddElementsToArray is a paid mutator transaction binding the contract method 0xe53213b3.
//
// Solidity: function addElementsToArray(uint256 newElements) returns()
func (_SampleTestContract *SampleTestContractTransactor) AddElementsToArray(opts *bind.TransactOpts, newElements *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.contract.Transact(opts, "addElementsToArray", newElements)
}

// AddElementsToArray is a paid mutator transaction binding the contract method 0xe53213b3.
//
// Solidity: function addElementsToArray(uint256 newElements) returns()
func (_SampleTestContract *SampleTestContractSession) AddElementsToArray(newElements *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.Contract.AddElementsToArray(&_SampleTestContract.TransactOpts, newElements)
}

// AddElementsToArray is a paid mutator transaction binding the contract method 0xe53213b3.
//
// Solidity: function addElementsToArray(uint256 newElements) returns()
func (_SampleTestContract *SampleTestContractTransactorSession) AddElementsToArray(newElements *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.Contract.AddElementsToArray(&_SampleTestContract.TransactOpts, newElements)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_SampleTestContract *SampleTestContractTransactor) PerformUpkeep(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error) {
	return _SampleTestContract.contract.Transact(opts, "performUpkeep", performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_SampleTestContract *SampleTestContractSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _SampleTestContract.Contract.PerformUpkeep(&_SampleTestContract.TransactOpts, performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_SampleTestContract *SampleTestContractTransactorSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _SampleTestContract.Contract.PerformUpkeep(&_SampleTestContract.TransactOpts, performData)
}

// ResizeArray is a paid mutator transaction binding the contract method 0xa00d6cd0.
//
// Solidity: function resizeArray(uint256 newLength) returns()
func (_SampleTestContract *SampleTestContractTransactor) ResizeArray(opts *bind.TransactOpts, newLength *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.contract.Transact(opts, "resizeArray", newLength)
}

// ResizeArray is a paid mutator transaction binding the contract method 0xa00d6cd0.
//
// Solidity: function resizeArray(uint256 newLength) returns()
func (_SampleTestContract *SampleTestContractSession) ResizeArray(newLength *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.Contract.ResizeArray(&_SampleTestContract.TransactOpts, newLength)
}

// ResizeArray is a paid mutator transaction binding the contract method 0xa00d6cd0.
//
// Solidity: function resizeArray(uint256 newLength) returns()
func (_SampleTestContract *SampleTestContractTransactorSession) ResizeArray(newLength *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.Contract.ResizeArray(&_SampleTestContract.TransactOpts, newLength)
}

// SetMaxIterations is a paid mutator transaction binding the contract method 0x8d8987b2.
//
// Solidity: function setMaxIterations(uint256 _checkMaxIterations, uint256 _performMaxIterations) returns()
func (_SampleTestContract *SampleTestContractTransactor) SetMaxIterations(opts *bind.TransactOpts, _checkMaxIterations *big.Int, _performMaxIterations *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.contract.Transact(opts, "setMaxIterations", _checkMaxIterations, _performMaxIterations)
}

// SetMaxIterations is a paid mutator transaction binding the contract method 0x8d8987b2.
//
// Solidity: function setMaxIterations(uint256 _checkMaxIterations, uint256 _performMaxIterations) returns()
func (_SampleTestContract *SampleTestContractSession) SetMaxIterations(_checkMaxIterations *big.Int, _performMaxIterations *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.Contract.SetMaxIterations(&_SampleTestContract.TransactOpts, _checkMaxIterations, _performMaxIterations)
}

// SetMaxIterations is a paid mutator transaction binding the contract method 0x8d8987b2.
//
// Solidity: function setMaxIterations(uint256 _checkMaxIterations, uint256 _performMaxIterations) returns()
func (_SampleTestContract *SampleTestContractTransactorSession) SetMaxIterations(_checkMaxIterations *big.Int, _performMaxIterations *big.Int) (*types.Transaction, error) {
	return _SampleTestContract.Contract.SetMaxIterations(&_SampleTestContract.TransactOpts, _checkMaxIterations, _performMaxIterations)
}

// ToggleFlag is a paid mutator transaction binding the contract method 0x00c0096e.
//
// Solidity: function toggleFlag() returns()
func (_SampleTestContract *SampleTestContractTransactor) ToggleFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleTestContract.contract.Transact(opts, "toggleFlag")
}

// ToggleFlag is a paid mutator transaction binding the contract method 0x00c0096e.
//
// Solidity: function toggleFlag() returns()
func (_SampleTestContract *SampleTestContractSession) ToggleFlag() (*types.Transaction, error) {
	return _SampleTestContract.Contract.ToggleFlag(&_SampleTestContract.TransactOpts)
}

// ToggleFlag is a paid mutator transaction binding the contract method 0x00c0096e.
//
// Solidity: function toggleFlag() returns()
func (_SampleTestContract *SampleTestContractTransactorSession) ToggleFlag() (*types.Transaction, error) {
	return _SampleTestContract.Contract.ToggleFlag(&_SampleTestContract.TransactOpts)
}
