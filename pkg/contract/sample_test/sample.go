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

// RegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type RegistrationParams struct {
	Name           string
	EncryptedEmail []byte
	UpkeepContract common.Address
	GasLimit       uint32
	AdminAddress   common.Address
	CheckData      []byte
	OffchainConfig []byte
	Amount         *big.Int
}

// SampleTestContractMetaData contains all meta data concerning the SampleTestContract contract.
var SampleTestContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"contractKeeperRegistrarInterface\",\"name\":\"registrar\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_registrar\",\"outputs\":[{\"internalType\":\"contractKeeperRegistrarInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"registerAndPredictID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ILink is a free data retrieval call binding the contract method 0x7d253aff.
//
// Solidity: function i_link() view returns(address)
func (_SampleTestContract *SampleTestContractCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ILink is a free data retrieval call binding the contract method 0x7d253aff.
//
// Solidity: function i_link() view returns(address)
func (_SampleTestContract *SampleTestContractSession) ILink() (common.Address, error) {
	return _SampleTestContract.Contract.ILink(&_SampleTestContract.CallOpts)
}

// ILink is a free data retrieval call binding the contract method 0x7d253aff.
//
// Solidity: function i_link() view returns(address)
func (_SampleTestContract *SampleTestContractCallerSession) ILink() (common.Address, error) {
	return _SampleTestContract.Contract.ILink(&_SampleTestContract.CallOpts)
}

// IRegistrar is a free data retrieval call binding the contract method 0x442b1278.
//
// Solidity: function i_registrar() view returns(address)
func (_SampleTestContract *SampleTestContractCaller) IRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SampleTestContract.contract.Call(opts, &out, "i_registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IRegistrar is a free data retrieval call binding the contract method 0x442b1278.
//
// Solidity: function i_registrar() view returns(address)
func (_SampleTestContract *SampleTestContractSession) IRegistrar() (common.Address, error) {
	return _SampleTestContract.Contract.IRegistrar(&_SampleTestContract.CallOpts)
}

// IRegistrar is a free data retrieval call binding the contract method 0x442b1278.
//
// Solidity: function i_registrar() view returns(address)
func (_SampleTestContract *SampleTestContractCallerSession) IRegistrar() (common.Address, error) {
	return _SampleTestContract.Contract.IRegistrar(&_SampleTestContract.CallOpts)
}

// RegisterAndPredictID is a paid mutator transaction binding the contract method 0xfd111cca.
//
// Solidity: function registerAndPredictID((string,bytes,address,uint32,address,bytes,bytes,uint96) params) returns()
func (_SampleTestContract *SampleTestContractTransactor) RegisterAndPredictID(opts *bind.TransactOpts, params RegistrationParams) (*types.Transaction, error) {
	return _SampleTestContract.contract.Transact(opts, "registerAndPredictID", params)
}

// RegisterAndPredictID is a paid mutator transaction binding the contract method 0xfd111cca.
//
// Solidity: function registerAndPredictID((string,bytes,address,uint32,address,bytes,bytes,uint96) params) returns()
func (_SampleTestContract *SampleTestContractSession) RegisterAndPredictID(params RegistrationParams) (*types.Transaction, error) {
	return _SampleTestContract.Contract.RegisterAndPredictID(&_SampleTestContract.TransactOpts, params)
}

// RegisterAndPredictID is a paid mutator transaction binding the contract method 0xfd111cca.
//
// Solidity: function registerAndPredictID((string,bytes,address,uint32,address,bytes,bytes,uint96) params) returns()
func (_SampleTestContract *SampleTestContractTransactorSession) RegisterAndPredictID(params RegistrationParams) (*types.Transaction, error) {
	return _SampleTestContract.Contract.RegisterAndPredictID(&_SampleTestContract.TransactOpts, params)
}

// SampleTestContractUpkeepCreatedIterator is returned from FilterUpkeepCreated and is used to iterate over the raw logs and unpacked data for UpkeepCreated events raised by the SampleTestContract contract.
type SampleTestContractUpkeepCreatedIterator struct {
	Event *SampleTestContractUpkeepCreated // Event containing the contract specifics and raw log

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
func (it *SampleTestContractUpkeepCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleTestContractUpkeepCreated)
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
		it.Event = new(SampleTestContractUpkeepCreated)
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
func (it *SampleTestContractUpkeepCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleTestContractUpkeepCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleTestContractUpkeepCreated represents a UpkeepCreated event raised by the SampleTestContract contract.
type SampleTestContractUpkeepCreated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpkeepCreated is a free log retrieval operation binding the contract event 0x73827404e930da1450aa92d0348dd0939840f4850415c6b42a1bcf9de4162d7b.
//
// Solidity: event UpkeepCreated(uint256 id)
func (_SampleTestContract *SampleTestContractFilterer) FilterUpkeepCreated(opts *bind.FilterOpts) (*SampleTestContractUpkeepCreatedIterator, error) {

	logs, sub, err := _SampleTestContract.contract.FilterLogs(opts, "UpkeepCreated")
	if err != nil {
		return nil, err
	}
	return &SampleTestContractUpkeepCreatedIterator{contract: _SampleTestContract.contract, event: "UpkeepCreated", logs: logs, sub: sub}, nil
}

// WatchUpkeepCreated is a free log subscription operation binding the contract event 0x73827404e930da1450aa92d0348dd0939840f4850415c6b42a1bcf9de4162d7b.
//
// Solidity: event UpkeepCreated(uint256 id)
func (_SampleTestContract *SampleTestContractFilterer) WatchUpkeepCreated(opts *bind.WatchOpts, sink chan<- *SampleTestContractUpkeepCreated) (event.Subscription, error) {

	logs, sub, err := _SampleTestContract.contract.WatchLogs(opts, "UpkeepCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleTestContractUpkeepCreated)
				if err := _SampleTestContract.contract.UnpackLog(event, "UpkeepCreated", log); err != nil {
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

// ParseUpkeepCreated is a log parse operation binding the contract event 0x73827404e930da1450aa92d0348dd0939840f4850415c6b42a1bcf9de4162d7b.
//
// Solidity: event UpkeepCreated(uint256 id)
func (_SampleTestContract *SampleTestContractFilterer) ParseUpkeepCreated(log types.Log) (*SampleTestContractUpkeepCreated, error) {
	event := new(SampleTestContractUpkeepCreated)
	if err := _SampleTestContract.contract.UnpackLog(event, "UpkeepCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
