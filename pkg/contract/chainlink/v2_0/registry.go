// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2_0

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

// KeeperRegistrar20RegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type KeeperRegistrar20RegistrationParams struct {
	Name           string
	EncryptedEmail []byte
	UpkeepContract common.Address
	GasLimit       uint32
	AdminAddress   common.Address
	CheckData      []byte
	OffchainConfig []byte
	Amount         *big.Int
}

// RegistryContractMetaData contains all meta data concerning the RegistryContract contract.
var RegistryContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LINKAddress\",\"type\":\"address\"},{\"internalType\":\"enumKeeperRegistrar2_0.AutoApproveType\",\"name\":\"autoApproveConfigType\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"minLINKJuels\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AmountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotPermitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"LinkTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistrationRequestFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AutoApproveAllowedSenderSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumKeeperRegistrar2_0.AutoApproveType\",\"name\":\"autoApproveConfigType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"minLINKJuels\",\"type\":\"uint96\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"RegistrationApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"RegistrationRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"RegistrationRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"getAutoApproveAllowedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getPendingRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationConfig\",\"outputs\":[{\"internalType\":\"enumKeeperRegistrar2_0.AutoApproveType\",\"name\":\"autoApproveConfigType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"approvedCount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minLINKJuels\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structKeeperRegistrar2_0.RegistrationParams\",\"name\":\"requestParams\",\"type\":\"tuple\"}],\"name\":\"registerUpkeep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setAutoApproveAllowedSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumKeeperRegistrar2_0.AutoApproveType\",\"name\":\"autoApproveConfigType\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"minLINKJuels\",\"type\":\"uint96\"}],\"name\":\"setRegistrationConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RegistryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryContractMetaData.ABI instead.
var RegistryContractABI = RegistryContractMetaData.ABI

// RegistryContract is an auto generated Go binding around an Ethereum contract.
type RegistryContract struct {
	RegistryContractCaller     // Read-only binding to the contract
	RegistryContractTransactor // Write-only binding to the contract
	RegistryContractFilterer   // Log filterer for contract events
}

// RegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistryContractSession struct {
	Contract     *RegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryContractCallerSession struct {
	Contract *RegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryContractTransactorSession struct {
	Contract     *RegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryContractRaw struct {
	Contract *RegistryContract // Generic contract binding to access the raw methods on
}

// RegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryContractCallerRaw struct {
	Contract *RegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryContractTransactorRaw struct {
	Contract *RegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistryContract creates a new instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContract(address common.Address, backend bind.ContractBackend) (*RegistryContract, error) {
	contract, err := bindRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryContract{RegistryContractCaller: RegistryContractCaller{contract: contract}, RegistryContractTransactor: RegistryContractTransactor{contract: contract}, RegistryContractFilterer: RegistryContractFilterer{contract: contract}}, nil
}

// NewRegistryContractCaller creates a new read-only instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*RegistryContractCaller, error) {
	contract, err := bindRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryContractCaller{contract: contract}, nil
}

// NewRegistryContractTransactor creates a new write-only instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryContractTransactor, error) {
	contract, err := bindRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryContractTransactor{contract: contract}, nil
}

// NewRegistryContractFilterer creates a new log filterer instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryContractFilterer, error) {
	contract, err := bindRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryContractFilterer{contract: contract}, nil
}

// bindRegistryContract binds a generic wrapper to an already deployed contract.
func bindRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryContract *RegistryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryContract.Contract.RegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryContract *RegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryContract *RegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryContract *RegistryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryContract *RegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryContract *RegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryContract.Contract.contract.Transact(opts, method, params...)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_RegistryContract *RegistryContractCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_RegistryContract *RegistryContractSession) LINK() (common.Address, error) {
	return _RegistryContract.Contract.LINK(&_RegistryContract.CallOpts)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_RegistryContract *RegistryContractCallerSession) LINK() (common.Address, error) {
	return _RegistryContract.Contract.LINK(&_RegistryContract.CallOpts)
}

// GetAutoApproveAllowedSender is a free data retrieval call binding the contract method 0x7e776f7f.
//
// Solidity: function getAutoApproveAllowedSender(address senderAddress) view returns(bool)
func (_RegistryContract *RegistryContractCaller) GetAutoApproveAllowedSender(opts *bind.CallOpts, senderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getAutoApproveAllowedSender", senderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAutoApproveAllowedSender is a free data retrieval call binding the contract method 0x7e776f7f.
//
// Solidity: function getAutoApproveAllowedSender(address senderAddress) view returns(bool)
func (_RegistryContract *RegistryContractSession) GetAutoApproveAllowedSender(senderAddress common.Address) (bool, error) {
	return _RegistryContract.Contract.GetAutoApproveAllowedSender(&_RegistryContract.CallOpts, senderAddress)
}

// GetAutoApproveAllowedSender is a free data retrieval call binding the contract method 0x7e776f7f.
//
// Solidity: function getAutoApproveAllowedSender(address senderAddress) view returns(bool)
func (_RegistryContract *RegistryContractCallerSession) GetAutoApproveAllowedSender(senderAddress common.Address) (bool, error) {
	return _RegistryContract.Contract.GetAutoApproveAllowedSender(&_RegistryContract.CallOpts, senderAddress)
}

// GetPendingRequest is a free data retrieval call binding the contract method 0x88b12d55.
//
// Solidity: function getPendingRequest(bytes32 hash) view returns(address, uint96)
func (_RegistryContract *RegistryContractCaller) GetPendingRequest(opts *bind.CallOpts, hash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getPendingRequest", hash)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPendingRequest is a free data retrieval call binding the contract method 0x88b12d55.
//
// Solidity: function getPendingRequest(bytes32 hash) view returns(address, uint96)
func (_RegistryContract *RegistryContractSession) GetPendingRequest(hash [32]byte) (common.Address, *big.Int, error) {
	return _RegistryContract.Contract.GetPendingRequest(&_RegistryContract.CallOpts, hash)
}

// GetPendingRequest is a free data retrieval call binding the contract method 0x88b12d55.
//
// Solidity: function getPendingRequest(bytes32 hash) view returns(address, uint96)
func (_RegistryContract *RegistryContractCallerSession) GetPendingRequest(hash [32]byte) (common.Address, *big.Int, error) {
	return _RegistryContract.Contract.GetPendingRequest(&_RegistryContract.CallOpts, hash)
}

// GetRegistrationConfig is a free data retrieval call binding the contract method 0x850af0cb.
//
// Solidity: function getRegistrationConfig() view returns(uint8 autoApproveConfigType, uint32 autoApproveMaxAllowed, uint32 approvedCount, address keeperRegistry, uint256 minLINKJuels)
func (_RegistryContract *RegistryContractCaller) GetRegistrationConfig(opts *bind.CallOpts) (struct {
	AutoApproveConfigType uint8
	AutoApproveMaxAllowed uint32
	ApprovedCount         uint32
	KeeperRegistry        common.Address
	MinLINKJuels          *big.Int
}, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getRegistrationConfig")

	outstruct := new(struct {
		AutoApproveConfigType uint8
		AutoApproveMaxAllowed uint32
		ApprovedCount         uint32
		KeeperRegistry        common.Address
		MinLINKJuels          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AutoApproveConfigType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.AutoApproveMaxAllowed = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ApprovedCount = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.KeeperRegistry = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.MinLINKJuels = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRegistrationConfig is a free data retrieval call binding the contract method 0x850af0cb.
//
// Solidity: function getRegistrationConfig() view returns(uint8 autoApproveConfigType, uint32 autoApproveMaxAllowed, uint32 approvedCount, address keeperRegistry, uint256 minLINKJuels)
func (_RegistryContract *RegistryContractSession) GetRegistrationConfig() (struct {
	AutoApproveConfigType uint8
	AutoApproveMaxAllowed uint32
	ApprovedCount         uint32
	KeeperRegistry        common.Address
	MinLINKJuels          *big.Int
}, error) {
	return _RegistryContract.Contract.GetRegistrationConfig(&_RegistryContract.CallOpts)
}

// GetRegistrationConfig is a free data retrieval call binding the contract method 0x850af0cb.
//
// Solidity: function getRegistrationConfig() view returns(uint8 autoApproveConfigType, uint32 autoApproveMaxAllowed, uint32 approvedCount, address keeperRegistry, uint256 minLINKJuels)
func (_RegistryContract *RegistryContractCallerSession) GetRegistrationConfig() (struct {
	AutoApproveConfigType uint8
	AutoApproveMaxAllowed uint32
	ApprovedCount         uint32
	KeeperRegistry        common.Address
	MinLINKJuels          *big.Int
}, error) {
	return _RegistryContract.Contract.GetRegistrationConfig(&_RegistryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryContract *RegistryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryContract *RegistryContractSession) Owner() (common.Address, error) {
	return _RegistryContract.Contract.Owner(&_RegistryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryContract *RegistryContractCallerSession) Owner() (common.Address, error) {
	return _RegistryContract.Contract.Owner(&_RegistryContract.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_RegistryContract *RegistryContractCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_RegistryContract *RegistryContractSession) TypeAndVersion() (string, error) {
	return _RegistryContract.Contract.TypeAndVersion(&_RegistryContract.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_RegistryContract *RegistryContractCallerSession) TypeAndVersion() (string, error) {
	return _RegistryContract.Contract.TypeAndVersion(&_RegistryContract.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RegistryContract *RegistryContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RegistryContract *RegistryContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _RegistryContract.Contract.AcceptOwnership(&_RegistryContract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RegistryContract *RegistryContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RegistryContract.Contract.AcceptOwnership(&_RegistryContract.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x62105854.
//
// Solidity: function approve(string name, address upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, bytes offchainConfig, bytes32 hash) returns()
func (_RegistryContract *RegistryContractTransactor) Approve(opts *bind.TransactOpts, name string, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, checkData []byte, offchainConfig []byte, hash [32]byte) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "approve", name, upkeepContract, gasLimit, adminAddress, checkData, offchainConfig, hash)
}

// Approve is a paid mutator transaction binding the contract method 0x62105854.
//
// Solidity: function approve(string name, address upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, bytes offchainConfig, bytes32 hash) returns()
func (_RegistryContract *RegistryContractSession) Approve(name string, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, checkData []byte, offchainConfig []byte, hash [32]byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.Approve(&_RegistryContract.TransactOpts, name, upkeepContract, gasLimit, adminAddress, checkData, offchainConfig, hash)
}

// Approve is a paid mutator transaction binding the contract method 0x62105854.
//
// Solidity: function approve(string name, address upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, bytes offchainConfig, bytes32 hash) returns()
func (_RegistryContract *RegistryContractTransactorSession) Approve(name string, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, checkData []byte, offchainConfig []byte, hash [32]byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.Approve(&_RegistryContract.TransactOpts, name, upkeepContract, gasLimit, adminAddress, checkData, offchainConfig, hash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 hash) returns()
func (_RegistryContract *RegistryContractTransactor) Cancel(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "cancel", hash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 hash) returns()
func (_RegistryContract *RegistryContractSession) Cancel(hash [32]byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.Cancel(&_RegistryContract.TransactOpts, hash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 hash) returns()
func (_RegistryContract *RegistryContractTransactorSession) Cancel(hash [32]byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.Cancel(&_RegistryContract.TransactOpts, hash)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_RegistryContract *RegistryContractTransactor) OnTokenTransfer(opts *bind.TransactOpts, sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "onTokenTransfer", sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_RegistryContract *RegistryContractSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.OnTokenTransfer(&_RegistryContract.TransactOpts, sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_RegistryContract *RegistryContractTransactorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.OnTokenTransfer(&_RegistryContract.TransactOpts, sender, amount, data)
}

// Register is a paid mutator transaction binding the contract method 0xa611ea56.
//
// Solidity: function register(string name, bytes encryptedEmail, address upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, bytes offchainConfig, uint96 amount, address sender) returns()
func (_RegistryContract *RegistryContractTransactor) Register(opts *bind.TransactOpts, name string, encryptedEmail []byte, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, checkData []byte, offchainConfig []byte, amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "register", name, encryptedEmail, upkeepContract, gasLimit, adminAddress, checkData, offchainConfig, amount, sender)
}

// Register is a paid mutator transaction binding the contract method 0xa611ea56.
//
// Solidity: function register(string name, bytes encryptedEmail, address upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, bytes offchainConfig, uint96 amount, address sender) returns()
func (_RegistryContract *RegistryContractSession) Register(name string, encryptedEmail []byte, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, checkData []byte, offchainConfig []byte, amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _RegistryContract.Contract.Register(&_RegistryContract.TransactOpts, name, encryptedEmail, upkeepContract, gasLimit, adminAddress, checkData, offchainConfig, amount, sender)
}

// Register is a paid mutator transaction binding the contract method 0xa611ea56.
//
// Solidity: function register(string name, bytes encryptedEmail, address upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, bytes offchainConfig, uint96 amount, address sender) returns()
func (_RegistryContract *RegistryContractTransactorSession) Register(name string, encryptedEmail []byte, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, checkData []byte, offchainConfig []byte, amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _RegistryContract.Contract.Register(&_RegistryContract.TransactOpts, name, encryptedEmail, upkeepContract, gasLimit, adminAddress, checkData, offchainConfig, amount, sender)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x08b79da4.
//
// Solidity: function registerUpkeep((string,bytes,address,uint32,address,bytes,bytes,uint96) requestParams) returns(uint256)
func (_RegistryContract *RegistryContractTransactor) RegisterUpkeep(opts *bind.TransactOpts, requestParams KeeperRegistrar20RegistrationParams) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "registerUpkeep", requestParams)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x08b79da4.
//
// Solidity: function registerUpkeep((string,bytes,address,uint32,address,bytes,bytes,uint96) requestParams) returns(uint256)
func (_RegistryContract *RegistryContractSession) RegisterUpkeep(requestParams KeeperRegistrar20RegistrationParams) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegisterUpkeep(&_RegistryContract.TransactOpts, requestParams)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x08b79da4.
//
// Solidity: function registerUpkeep((string,bytes,address,uint32,address,bytes,bytes,uint96) requestParams) returns(uint256)
func (_RegistryContract *RegistryContractTransactorSession) RegisterUpkeep(requestParams KeeperRegistrar20RegistrationParams) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegisterUpkeep(&_RegistryContract.TransactOpts, requestParams)
}

// SetAutoApproveAllowedSender is a paid mutator transaction binding the contract method 0x367b9b4f.
//
// Solidity: function setAutoApproveAllowedSender(address senderAddress, bool allowed) returns()
func (_RegistryContract *RegistryContractTransactor) SetAutoApproveAllowedSender(opts *bind.TransactOpts, senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "setAutoApproveAllowedSender", senderAddress, allowed)
}

// SetAutoApproveAllowedSender is a paid mutator transaction binding the contract method 0x367b9b4f.
//
// Solidity: function setAutoApproveAllowedSender(address senderAddress, bool allowed) returns()
func (_RegistryContract *RegistryContractSession) SetAutoApproveAllowedSender(senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _RegistryContract.Contract.SetAutoApproveAllowedSender(&_RegistryContract.TransactOpts, senderAddress, allowed)
}

// SetAutoApproveAllowedSender is a paid mutator transaction binding the contract method 0x367b9b4f.
//
// Solidity: function setAutoApproveAllowedSender(address senderAddress, bool allowed) returns()
func (_RegistryContract *RegistryContractTransactorSession) SetAutoApproveAllowedSender(senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _RegistryContract.Contract.SetAutoApproveAllowedSender(&_RegistryContract.TransactOpts, senderAddress, allowed)
}

// SetRegistrationConfig is a paid mutator transaction binding the contract method 0xa793ab8b.
//
// Solidity: function setRegistrationConfig(uint8 autoApproveConfigType, uint16 autoApproveMaxAllowed, address keeperRegistry, uint96 minLINKJuels) returns()
func (_RegistryContract *RegistryContractTransactor) SetRegistrationConfig(opts *bind.TransactOpts, autoApproveConfigType uint8, autoApproveMaxAllowed uint16, keeperRegistry common.Address, minLINKJuels *big.Int) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "setRegistrationConfig", autoApproveConfigType, autoApproveMaxAllowed, keeperRegistry, minLINKJuels)
}

// SetRegistrationConfig is a paid mutator transaction binding the contract method 0xa793ab8b.
//
// Solidity: function setRegistrationConfig(uint8 autoApproveConfigType, uint16 autoApproveMaxAllowed, address keeperRegistry, uint96 minLINKJuels) returns()
func (_RegistryContract *RegistryContractSession) SetRegistrationConfig(autoApproveConfigType uint8, autoApproveMaxAllowed uint16, keeperRegistry common.Address, minLINKJuels *big.Int) (*types.Transaction, error) {
	return _RegistryContract.Contract.SetRegistrationConfig(&_RegistryContract.TransactOpts, autoApproveConfigType, autoApproveMaxAllowed, keeperRegistry, minLINKJuels)
}

// SetRegistrationConfig is a paid mutator transaction binding the contract method 0xa793ab8b.
//
// Solidity: function setRegistrationConfig(uint8 autoApproveConfigType, uint16 autoApproveMaxAllowed, address keeperRegistry, uint96 minLINKJuels) returns()
func (_RegistryContract *RegistryContractTransactorSession) SetRegistrationConfig(autoApproveConfigType uint8, autoApproveMaxAllowed uint16, keeperRegistry common.Address, minLINKJuels *big.Int) (*types.Transaction, error) {
	return _RegistryContract.Contract.SetRegistrationConfig(&_RegistryContract.TransactOpts, autoApproveConfigType, autoApproveMaxAllowed, keeperRegistry, minLINKJuels)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_RegistryContract *RegistryContractTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_RegistryContract *RegistryContractSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RegistryContract.Contract.TransferOwnership(&_RegistryContract.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_RegistryContract *RegistryContractTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RegistryContract.Contract.TransferOwnership(&_RegistryContract.TransactOpts, to)
}

// RegistryContractAutoApproveAllowedSenderSetIterator is returned from FilterAutoApproveAllowedSenderSet and is used to iterate over the raw logs and unpacked data for AutoApproveAllowedSenderSet events raised by the RegistryContract contract.
type RegistryContractAutoApproveAllowedSenderSetIterator struct {
	Event *RegistryContractAutoApproveAllowedSenderSet // Event containing the contract specifics and raw log

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
func (it *RegistryContractAutoApproveAllowedSenderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractAutoApproveAllowedSenderSet)
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
		it.Event = new(RegistryContractAutoApproveAllowedSenderSet)
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
func (it *RegistryContractAutoApproveAllowedSenderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractAutoApproveAllowedSenderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractAutoApproveAllowedSenderSet represents a AutoApproveAllowedSenderSet event raised by the RegistryContract contract.
type RegistryContractAutoApproveAllowedSenderSet struct {
	SenderAddress common.Address
	Allowed       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAutoApproveAllowedSenderSet is a free log retrieval operation binding the contract event 0x20c6237dac83526a849285a9f79d08a483291bdd3a056a0ef9ae94ecee1ad356.
//
// Solidity: event AutoApproveAllowedSenderSet(address indexed senderAddress, bool allowed)
func (_RegistryContract *RegistryContractFilterer) FilterAutoApproveAllowedSenderSet(opts *bind.FilterOpts, senderAddress []common.Address) (*RegistryContractAutoApproveAllowedSenderSetIterator, error) {

	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "AutoApproveAllowedSenderSet", senderAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractAutoApproveAllowedSenderSetIterator{contract: _RegistryContract.contract, event: "AutoApproveAllowedSenderSet", logs: logs, sub: sub}, nil
}

// WatchAutoApproveAllowedSenderSet is a free log subscription operation binding the contract event 0x20c6237dac83526a849285a9f79d08a483291bdd3a056a0ef9ae94ecee1ad356.
//
// Solidity: event AutoApproveAllowedSenderSet(address indexed senderAddress, bool allowed)
func (_RegistryContract *RegistryContractFilterer) WatchAutoApproveAllowedSenderSet(opts *bind.WatchOpts, sink chan<- *RegistryContractAutoApproveAllowedSenderSet, senderAddress []common.Address) (event.Subscription, error) {

	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "AutoApproveAllowedSenderSet", senderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractAutoApproveAllowedSenderSet)
				if err := _RegistryContract.contract.UnpackLog(event, "AutoApproveAllowedSenderSet", log); err != nil {
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

// ParseAutoApproveAllowedSenderSet is a log parse operation binding the contract event 0x20c6237dac83526a849285a9f79d08a483291bdd3a056a0ef9ae94ecee1ad356.
//
// Solidity: event AutoApproveAllowedSenderSet(address indexed senderAddress, bool allowed)
func (_RegistryContract *RegistryContractFilterer) ParseAutoApproveAllowedSenderSet(log types.Log) (*RegistryContractAutoApproveAllowedSenderSet, error) {
	event := new(RegistryContractAutoApproveAllowedSenderSet)
	if err := _RegistryContract.contract.UnpackLog(event, "AutoApproveAllowedSenderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractConfigChangedIterator is returned from FilterConfigChanged and is used to iterate over the raw logs and unpacked data for ConfigChanged events raised by the RegistryContract contract.
type RegistryContractConfigChangedIterator struct {
	Event *RegistryContractConfigChanged // Event containing the contract specifics and raw log

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
func (it *RegistryContractConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractConfigChanged)
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
		it.Event = new(RegistryContractConfigChanged)
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
func (it *RegistryContractConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractConfigChanged represents a ConfigChanged event raised by the RegistryContract contract.
type RegistryContractConfigChanged struct {
	AutoApproveConfigType uint8
	AutoApproveMaxAllowed uint32
	KeeperRegistry        common.Address
	MinLINKJuels          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterConfigChanged is a free log retrieval operation binding the contract event 0x6293a703ec7145dfa23c5cde2e627d6a02e153fc2e9c03b14d1e22cbb4a7e9cd.
//
// Solidity: event ConfigChanged(uint8 autoApproveConfigType, uint32 autoApproveMaxAllowed, address keeperRegistry, uint96 minLINKJuels)
func (_RegistryContract *RegistryContractFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*RegistryContractConfigChangedIterator, error) {

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryContractConfigChangedIterator{contract: _RegistryContract.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

// WatchConfigChanged is a free log subscription operation binding the contract event 0x6293a703ec7145dfa23c5cde2e627d6a02e153fc2e9c03b14d1e22cbb4a7e9cd.
//
// Solidity: event ConfigChanged(uint8 autoApproveConfigType, uint32 autoApproveMaxAllowed, address keeperRegistry, uint96 minLINKJuels)
func (_RegistryContract *RegistryContractFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *RegistryContractConfigChanged) (event.Subscription, error) {

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractConfigChanged)
				if err := _RegistryContract.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

// ParseConfigChanged is a log parse operation binding the contract event 0x6293a703ec7145dfa23c5cde2e627d6a02e153fc2e9c03b14d1e22cbb4a7e9cd.
//
// Solidity: event ConfigChanged(uint8 autoApproveConfigType, uint32 autoApproveMaxAllowed, address keeperRegistry, uint96 minLINKJuels)
func (_RegistryContract *RegistryContractFilterer) ParseConfigChanged(log types.Log) (*RegistryContractConfigChanged, error) {
	event := new(RegistryContractConfigChanged)
	if err := _RegistryContract.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the RegistryContract contract.
type RegistryContractOwnershipTransferRequestedIterator struct {
	Event *RegistryContractOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *RegistryContractOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractOwnershipTransferRequested)
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
		it.Event = new(RegistryContractOwnershipTransferRequested)
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
func (it *RegistryContractOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the RegistryContract contract.
type RegistryContractOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_RegistryContract *RegistryContractFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RegistryContractOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractOwnershipTransferRequestedIterator{contract: _RegistryContract.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_RegistryContract *RegistryContractFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RegistryContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractOwnershipTransferRequested)
				if err := _RegistryContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_RegistryContract *RegistryContractFilterer) ParseOwnershipTransferRequested(log types.Log) (*RegistryContractOwnershipTransferRequested, error) {
	event := new(RegistryContractOwnershipTransferRequested)
	if err := _RegistryContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegistryContract contract.
type RegistryContractOwnershipTransferredIterator struct {
	Event *RegistryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractOwnershipTransferred)
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
		it.Event = new(RegistryContractOwnershipTransferred)
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
func (it *RegistryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractOwnershipTransferred represents a OwnershipTransferred event raised by the RegistryContract contract.
type RegistryContractOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_RegistryContract *RegistryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RegistryContractOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractOwnershipTransferredIterator{contract: _RegistryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_RegistryContract *RegistryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractOwnershipTransferred)
				if err := _RegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_RegistryContract *RegistryContractFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryContractOwnershipTransferred, error) {
	event := new(RegistryContractOwnershipTransferred)
	if err := _RegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractRegistrationApprovedIterator is returned from FilterRegistrationApproved and is used to iterate over the raw logs and unpacked data for RegistrationApproved events raised by the RegistryContract contract.
type RegistryContractRegistrationApprovedIterator struct {
	Event *RegistryContractRegistrationApproved // Event containing the contract specifics and raw log

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
func (it *RegistryContractRegistrationApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractRegistrationApproved)
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
		it.Event = new(RegistryContractRegistrationApproved)
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
func (it *RegistryContractRegistrationApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractRegistrationApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractRegistrationApproved represents a RegistrationApproved event raised by the RegistryContract contract.
type RegistryContractRegistrationApproved struct {
	Hash        [32]byte
	DisplayName string
	UpkeepId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistrationApproved is a free log retrieval operation binding the contract event 0xb9a292fb7e3edd920cd2d2829a3615a640c43fd7de0a0820aa0668feb4c37d4b.
//
// Solidity: event RegistrationApproved(bytes32 indexed hash, string displayName, uint256 indexed upkeepId)
func (_RegistryContract *RegistryContractFilterer) FilterRegistrationApproved(opts *bind.FilterOpts, hash [][32]byte, upkeepId []*big.Int) (*RegistryContractRegistrationApprovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepIdRule []interface{}
	for _, upkeepIdItem := range upkeepId {
		upkeepIdRule = append(upkeepIdRule, upkeepIdItem)
	}

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "RegistrationApproved", hashRule, upkeepIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractRegistrationApprovedIterator{contract: _RegistryContract.contract, event: "RegistrationApproved", logs: logs, sub: sub}, nil
}

// WatchRegistrationApproved is a free log subscription operation binding the contract event 0xb9a292fb7e3edd920cd2d2829a3615a640c43fd7de0a0820aa0668feb4c37d4b.
//
// Solidity: event RegistrationApproved(bytes32 indexed hash, string displayName, uint256 indexed upkeepId)
func (_RegistryContract *RegistryContractFilterer) WatchRegistrationApproved(opts *bind.WatchOpts, sink chan<- *RegistryContractRegistrationApproved, hash [][32]byte, upkeepId []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepIdRule []interface{}
	for _, upkeepIdItem := range upkeepId {
		upkeepIdRule = append(upkeepIdRule, upkeepIdItem)
	}

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "RegistrationApproved", hashRule, upkeepIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractRegistrationApproved)
				if err := _RegistryContract.contract.UnpackLog(event, "RegistrationApproved", log); err != nil {
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

// ParseRegistrationApproved is a log parse operation binding the contract event 0xb9a292fb7e3edd920cd2d2829a3615a640c43fd7de0a0820aa0668feb4c37d4b.
//
// Solidity: event RegistrationApproved(bytes32 indexed hash, string displayName, uint256 indexed upkeepId)
func (_RegistryContract *RegistryContractFilterer) ParseRegistrationApproved(log types.Log) (*RegistryContractRegistrationApproved, error) {
	event := new(RegistryContractRegistrationApproved)
	if err := _RegistryContract.contract.UnpackLog(event, "RegistrationApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractRegistrationRejectedIterator is returned from FilterRegistrationRejected and is used to iterate over the raw logs and unpacked data for RegistrationRejected events raised by the RegistryContract contract.
type RegistryContractRegistrationRejectedIterator struct {
	Event *RegistryContractRegistrationRejected // Event containing the contract specifics and raw log

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
func (it *RegistryContractRegistrationRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractRegistrationRejected)
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
		it.Event = new(RegistryContractRegistrationRejected)
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
func (it *RegistryContractRegistrationRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractRegistrationRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractRegistrationRejected represents a RegistrationRejected event raised by the RegistryContract contract.
type RegistryContractRegistrationRejected struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRejected is a free log retrieval operation binding the contract event 0x3663fb28ebc87645eb972c9dad8521bf665c623f287e79f1c56f1eb374b82a22.
//
// Solidity: event RegistrationRejected(bytes32 indexed hash)
func (_RegistryContract *RegistryContractFilterer) FilterRegistrationRejected(opts *bind.FilterOpts, hash [][32]byte) (*RegistryContractRegistrationRejectedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "RegistrationRejected", hashRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractRegistrationRejectedIterator{contract: _RegistryContract.contract, event: "RegistrationRejected", logs: logs, sub: sub}, nil
}

// WatchRegistrationRejected is a free log subscription operation binding the contract event 0x3663fb28ebc87645eb972c9dad8521bf665c623f287e79f1c56f1eb374b82a22.
//
// Solidity: event RegistrationRejected(bytes32 indexed hash)
func (_RegistryContract *RegistryContractFilterer) WatchRegistrationRejected(opts *bind.WatchOpts, sink chan<- *RegistryContractRegistrationRejected, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "RegistrationRejected", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractRegistrationRejected)
				if err := _RegistryContract.contract.UnpackLog(event, "RegistrationRejected", log); err != nil {
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

// ParseRegistrationRejected is a log parse operation binding the contract event 0x3663fb28ebc87645eb972c9dad8521bf665c623f287e79f1c56f1eb374b82a22.
//
// Solidity: event RegistrationRejected(bytes32 indexed hash)
func (_RegistryContract *RegistryContractFilterer) ParseRegistrationRejected(log types.Log) (*RegistryContractRegistrationRejected, error) {
	event := new(RegistryContractRegistrationRejected)
	if err := _RegistryContract.contract.UnpackLog(event, "RegistrationRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractRegistrationRequestedIterator is returned from FilterRegistrationRequested and is used to iterate over the raw logs and unpacked data for RegistrationRequested events raised by the RegistryContract contract.
type RegistryContractRegistrationRequestedIterator struct {
	Event *RegistryContractRegistrationRequested // Event containing the contract specifics and raw log

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
func (it *RegistryContractRegistrationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractRegistrationRequested)
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
		it.Event = new(RegistryContractRegistrationRequested)
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
func (it *RegistryContractRegistrationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractRegistrationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractRegistrationRequested represents a RegistrationRequested event raised by the RegistryContract contract.
type RegistryContractRegistrationRequested struct {
	Hash           [32]byte
	Name           string
	EncryptedEmail []byte
	UpkeepContract common.Address
	GasLimit       uint32
	AdminAddress   common.Address
	CheckData      []byte
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRequested is a free log retrieval operation binding the contract event 0x9b8456f925542af2c5fb15ff4be32cc8f209dda96c544766e301367df40f4998.
//
// Solidity: event RegistrationRequested(bytes32 indexed hash, string name, bytes encryptedEmail, address indexed upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, uint96 amount)
func (_RegistryContract *RegistryContractFilterer) FilterRegistrationRequested(opts *bind.FilterOpts, hash [][32]byte, upkeepContract []common.Address) (*RegistryContractRegistrationRequestedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepContractRule []interface{}
	for _, upkeepContractItem := range upkeepContract {
		upkeepContractRule = append(upkeepContractRule, upkeepContractItem)
	}

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "RegistrationRequested", hashRule, upkeepContractRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractRegistrationRequestedIterator{contract: _RegistryContract.contract, event: "RegistrationRequested", logs: logs, sub: sub}, nil
}

// WatchRegistrationRequested is a free log subscription operation binding the contract event 0x9b8456f925542af2c5fb15ff4be32cc8f209dda96c544766e301367df40f4998.
//
// Solidity: event RegistrationRequested(bytes32 indexed hash, string name, bytes encryptedEmail, address indexed upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, uint96 amount)
func (_RegistryContract *RegistryContractFilterer) WatchRegistrationRequested(opts *bind.WatchOpts, sink chan<- *RegistryContractRegistrationRequested, hash [][32]byte, upkeepContract []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepContractRule []interface{}
	for _, upkeepContractItem := range upkeepContract {
		upkeepContractRule = append(upkeepContractRule, upkeepContractItem)
	}

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "RegistrationRequested", hashRule, upkeepContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractRegistrationRequested)
				if err := _RegistryContract.contract.UnpackLog(event, "RegistrationRequested", log); err != nil {
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

// ParseRegistrationRequested is a log parse operation binding the contract event 0x9b8456f925542af2c5fb15ff4be32cc8f209dda96c544766e301367df40f4998.
//
// Solidity: event RegistrationRequested(bytes32 indexed hash, string name, bytes encryptedEmail, address indexed upkeepContract, uint32 gasLimit, address adminAddress, bytes checkData, uint96 amount)
func (_RegistryContract *RegistryContractFilterer) ParseRegistrationRequested(log types.Log) (*RegistryContractRegistrationRequested, error) {
	event := new(RegistryContractRegistrationRequested)
	if err := _RegistryContract.contract.UnpackLog(event, "RegistrationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
