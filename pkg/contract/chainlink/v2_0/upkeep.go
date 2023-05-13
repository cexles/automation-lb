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

// OnchainConfig is an auto generated low-level Go binding around an user-defined struct.
type OnchainConfig struct {
	PaymentPremiumPPB    uint32
	FlatFeeMicroLink     uint32
	CheckGasLimit        uint32
	StalenessSeconds     *big.Int
	GasCeilingMultiplier uint16
	MinUpkeepSpend       *big.Int
	MaxPerformGas        uint32
	MaxCheckDataSize     uint32
	MaxPerformDataSize   uint32
	FallbackGasPrice     *big.Int
	FallbackLinkPrice    *big.Int
	Transcoder           common.Address
	Registrar            common.Address
}

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

// State is an auto generated low-level Go binding around an user-defined struct.
type State struct {
	Nonce                   uint32
	OwnerLinkBalance        *big.Int
	ExpectedLinkBalance     *big.Int
	TotalPremium            *big.Int
	NumUpkeeps              *big.Int
	ConfigCount             uint32
	LatestConfigBlockNumber uint32
	LatestConfigDigest      [32]byte
	LatestEpoch             uint32
	Paused                  bool
}

// UpkeepControllerExampleDetailedUpkeep is an auto generated low-level Go binding around an user-defined struct.
type UpkeepControllerExampleDetailedUpkeep struct {
	Id        *big.Int
	MinAmount *big.Int
	Info      UpkeepInfo
}

// UpkeepInfo is an auto generated low-level Go binding around an user-defined struct.
type UpkeepInfo struct {
	Target                 common.Address
	ExecuteGas             uint32
	CheckData              []byte
	Balance                *big.Int
	Admin                  common.Address
	MaxValidBlocknumber    uint64
	LastPerformBlockNumber uint32
	AmountSpent            *big.Int
	Paused                 bool
	OffchainConfig         []byte
}

// UpkeepControllerContractMetaData contains all meta data concerning the UpkeepControllerContract contract.
var UpkeepControllerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"contractKeeperRegistrarInterface\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"contractAutomationRegistryWithMinANeededAmountInterface\",\"name\":\"registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"FundsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"UpkeepGasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"name\":\"UpkeepOffchainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newCheckData\",\"type\":\"bytes\"}],\"name\":\"UpkeepUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"addFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"cancelUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"},{\"internalType\":\"enumUpkeepFailureReason\",\"name\":\"upkeepFailureReason\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastGasWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkNative\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"name\":\"decodePerformDataTest\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getActiveUpkeepIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"upkeeps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getDetailedUpkeeps\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"minAmount\",\"type\":\"uint96\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executeGas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxValidBlocknumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"lastPerformBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"amountSpent\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"internalType\":\"structUpkeepInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structUpkeepControllerExample.DetailedUpkeep[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUpkeepId\",\"type\":\"uint256\"}],\"name\":\"getInfoTest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"performOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastIdTest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"getMinBalanceForUpkeep\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getMinBalancesForUpkeeps\",\"outputs\":[{\"internalType\":\"uint96[]\",\"name\":\"\",\"type\":\"uint96[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAutomationCompatibleWithViewInterface\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"}],\"name\":\"getPerformDataTest\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"performOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"performArray\",\"type\":\"uint256[]\"}],\"name\":\"getResultTest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNeeded\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"newOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"ownerLinkBalance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"expectedLinkBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"totalPremium\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"numUpkeeps\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"latestConfigBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"latestConfigDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"latestEpoch\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"internalType\":\"structState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"paymentPremiumPPB\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flatFeeMicroLink\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"stalenessSeconds\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"gasCeilingMultiplier\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"minUpkeepSpend\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"maxPerformGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxCheckDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerformDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fallbackGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fallbackLinkPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"}],\"internalType\":\"structOnchainConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"getUpkeep\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executeGas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxValidBlocknumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"lastPerformBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"amountSpent\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"internalType\":\"structUpkeepInfo\",\"name\":\"upkeepInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getUpkeeps\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executeGas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxValidBlocknumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"lastPerformBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"amountSpent\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"internalType\":\"structUpkeepInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUpkeepsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_registrar\",\"outputs\":[{\"internalType\":\"contractKeeperRegistrarInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_registry\",\"outputs\":[{\"internalType\":\"contractAutomationRegistryWithMinANeededAmountInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isNewUpkeepNeeded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNeeded\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"newOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"pauseUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"registerAndPredictID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"setUpkeepGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"name\":\"setUpkeepOffchainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"unpauseUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"newCheckData\",\"type\":\"bytes\"}],\"name\":\"updateCheckData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UpkeepControllerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use UpkeepControllerContractMetaData.ABI instead.
var UpkeepControllerContractABI = UpkeepControllerContractMetaData.ABI

// UpkeepControllerContract is an auto generated Go binding around an Ethereum contract.
type UpkeepControllerContract struct {
	UpkeepControllerContractCaller     // Read-only binding to the contract
	UpkeepControllerContractTransactor // Write-only binding to the contract
	UpkeepControllerContractFilterer   // Log filterer for contract events
}

// UpkeepControllerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpkeepControllerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpkeepControllerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpkeepControllerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpkeepControllerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpkeepControllerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpkeepControllerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpkeepControllerContractSession struct {
	Contract     *UpkeepControllerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UpkeepControllerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpkeepControllerContractCallerSession struct {
	Contract *UpkeepControllerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// UpkeepControllerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpkeepControllerContractTransactorSession struct {
	Contract     *UpkeepControllerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// UpkeepControllerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpkeepControllerContractRaw struct {
	Contract *UpkeepControllerContract // Generic contract binding to access the raw methods on
}

// UpkeepControllerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpkeepControllerContractCallerRaw struct {
	Contract *UpkeepControllerContractCaller // Generic read-only contract binding to access the raw methods on
}

// UpkeepControllerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpkeepControllerContractTransactorRaw struct {
	Contract *UpkeepControllerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpkeepControllerContract creates a new instance of UpkeepControllerContract, bound to a specific deployed contract.
func NewUpkeepControllerContract(address common.Address, backend bind.ContractBackend) (*UpkeepControllerContract, error) {
	contract, err := bindUpkeepControllerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContract{UpkeepControllerContractCaller: UpkeepControllerContractCaller{contract: contract}, UpkeepControllerContractTransactor: UpkeepControllerContractTransactor{contract: contract}, UpkeepControllerContractFilterer: UpkeepControllerContractFilterer{contract: contract}}, nil
}

// NewUpkeepControllerContractCaller creates a new read-only instance of UpkeepControllerContract, bound to a specific deployed contract.
func NewUpkeepControllerContractCaller(address common.Address, caller bind.ContractCaller) (*UpkeepControllerContractCaller, error) {
	contract, err := bindUpkeepControllerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractCaller{contract: contract}, nil
}

// NewUpkeepControllerContractTransactor creates a new write-only instance of UpkeepControllerContract, bound to a specific deployed contract.
func NewUpkeepControllerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*UpkeepControllerContractTransactor, error) {
	contract, err := bindUpkeepControllerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractTransactor{contract: contract}, nil
}

// NewUpkeepControllerContractFilterer creates a new log filterer instance of UpkeepControllerContract, bound to a specific deployed contract.
func NewUpkeepControllerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*UpkeepControllerContractFilterer, error) {
	contract, err := bindUpkeepControllerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractFilterer{contract: contract}, nil
}

// bindUpkeepControllerContract binds a generic wrapper to an already deployed contract.
func bindUpkeepControllerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpkeepControllerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpkeepControllerContract *UpkeepControllerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpkeepControllerContract.Contract.UpkeepControllerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpkeepControllerContract *UpkeepControllerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.UpkeepControllerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpkeepControllerContract *UpkeepControllerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.UpkeepControllerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpkeepControllerContract *UpkeepControllerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpkeepControllerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpkeepControllerContract *UpkeepControllerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpkeepControllerContract *UpkeepControllerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.contract.Transact(opts, method, params...)
}

// DecodePerformDataTest is a free data retrieval call binding the contract method 0xf9913dcf.
//
// Solidity: function decodePerformDataTest(bytes performData) pure returns(uint256[])
func (_UpkeepControllerContract *UpkeepControllerContractCaller) DecodePerformDataTest(opts *bind.CallOpts, performData []byte) ([]*big.Int, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "decodePerformDataTest", performData)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// DecodePerformDataTest is a free data retrieval call binding the contract method 0xf9913dcf.
//
// Solidity: function decodePerformDataTest(bytes performData) pure returns(uint256[])
func (_UpkeepControllerContract *UpkeepControllerContractSession) DecodePerformDataTest(performData []byte) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.DecodePerformDataTest(&_UpkeepControllerContract.CallOpts, performData)
}

// DecodePerformDataTest is a free data retrieval call binding the contract method 0xf9913dcf.
//
// Solidity: function decodePerformDataTest(bytes performData) pure returns(uint256[])
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) DecodePerformDataTest(performData []byte) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.DecodePerformDataTest(&_UpkeepControllerContract.CallOpts, performData)
}

// GetActiveUpkeepIDs is a free data retrieval call binding the contract method 0x06e3b632.
//
// Solidity: function getActiveUpkeepIDs(uint256 offset, uint256 limit) view returns(uint256[] upkeeps)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetActiveUpkeepIDs(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getActiveUpkeepIDs", offset, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveUpkeepIDs is a free data retrieval call binding the contract method 0x06e3b632.
//
// Solidity: function getActiveUpkeepIDs(uint256 offset, uint256 limit) view returns(uint256[] upkeeps)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetActiveUpkeepIDs(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetActiveUpkeepIDs(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetActiveUpkeepIDs is a free data retrieval call binding the contract method 0x06e3b632.
//
// Solidity: function getActiveUpkeepIDs(uint256 offset, uint256 limit) view returns(uint256[] upkeeps)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetActiveUpkeepIDs(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetActiveUpkeepIDs(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetDetailedUpkeeps is a free data retrieval call binding the contract method 0xc43322bb.
//
// Solidity: function getDetailedUpkeeps(uint256 offset, uint256 limit) view returns((uint256,uint96,(address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes))[])
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetDetailedUpkeeps(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]UpkeepControllerExampleDetailedUpkeep, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getDetailedUpkeeps", offset, limit)

	if err != nil {
		return *new([]UpkeepControllerExampleDetailedUpkeep), err
	}

	out0 := *abi.ConvertType(out[0], new([]UpkeepControllerExampleDetailedUpkeep)).(*[]UpkeepControllerExampleDetailedUpkeep)

	return out0, err

}

// GetDetailedUpkeeps is a free data retrieval call binding the contract method 0xc43322bb.
//
// Solidity: function getDetailedUpkeeps(uint256 offset, uint256 limit) view returns((uint256,uint96,(address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes))[])
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetDetailedUpkeeps(offset *big.Int, limit *big.Int) ([]UpkeepControllerExampleDetailedUpkeep, error) {
	return _UpkeepControllerContract.Contract.GetDetailedUpkeeps(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetDetailedUpkeeps is a free data retrieval call binding the contract method 0xc43322bb.
//
// Solidity: function getDetailedUpkeeps(uint256 offset, uint256 limit) view returns((uint256,uint96,(address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes))[])
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetDetailedUpkeeps(offset *big.Int, limit *big.Int) ([]UpkeepControllerExampleDetailedUpkeep, error) {
	return _UpkeepControllerContract.Contract.GetDetailedUpkeeps(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetInfoTest is a free data retrieval call binding the contract method 0x7daa8f5a.
//
// Solidity: function getInfoTest(uint256 lastUpkeepId) view returns(uint256 performOffset, uint256 performLimit)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetInfoTest(opts *bind.CallOpts, lastUpkeepId *big.Int) (struct {
	PerformOffset *big.Int
	PerformLimit  *big.Int
}, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getInfoTest", lastUpkeepId)

	outstruct := new(struct {
		PerformOffset *big.Int
		PerformLimit  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerformOffset = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PerformLimit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetInfoTest is a free data retrieval call binding the contract method 0x7daa8f5a.
//
// Solidity: function getInfoTest(uint256 lastUpkeepId) view returns(uint256 performOffset, uint256 performLimit)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetInfoTest(lastUpkeepId *big.Int) (struct {
	PerformOffset *big.Int
	PerformLimit  *big.Int
}, error) {
	return _UpkeepControllerContract.Contract.GetInfoTest(&_UpkeepControllerContract.CallOpts, lastUpkeepId)
}

// GetInfoTest is a free data retrieval call binding the contract method 0x7daa8f5a.
//
// Solidity: function getInfoTest(uint256 lastUpkeepId) view returns(uint256 performOffset, uint256 performLimit)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetInfoTest(lastUpkeepId *big.Int) (struct {
	PerformOffset *big.Int
	PerformLimit  *big.Int
}, error) {
	return _UpkeepControllerContract.Contract.GetInfoTest(&_UpkeepControllerContract.CallOpts, lastUpkeepId)
}

// GetLastIdTest is a free data retrieval call binding the contract method 0x4bcfe7d9.
//
// Solidity: function getLastIdTest() view returns(uint256)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetLastIdTest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getLastIdTest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastIdTest is a free data retrieval call binding the contract method 0x4bcfe7d9.
//
// Solidity: function getLastIdTest() view returns(uint256)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetLastIdTest() (*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetLastIdTest(&_UpkeepControllerContract.CallOpts)
}

// GetLastIdTest is a free data retrieval call binding the contract method 0x4bcfe7d9.
//
// Solidity: function getLastIdTest() view returns(uint256)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetLastIdTest() (*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetLastIdTest(&_UpkeepControllerContract.CallOpts)
}

// GetMinBalanceForUpkeep is a free data retrieval call binding the contract method 0xb657bc9c.
//
// Solidity: function getMinBalanceForUpkeep(uint256 upkeepId) view returns(uint96)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetMinBalanceForUpkeep(opts *bind.CallOpts, upkeepId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getMinBalanceForUpkeep", upkeepId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBalanceForUpkeep is a free data retrieval call binding the contract method 0xb657bc9c.
//
// Solidity: function getMinBalanceForUpkeep(uint256 upkeepId) view returns(uint96)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetMinBalanceForUpkeep(upkeepId *big.Int) (*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetMinBalanceForUpkeep(&_UpkeepControllerContract.CallOpts, upkeepId)
}

// GetMinBalanceForUpkeep is a free data retrieval call binding the contract method 0xb657bc9c.
//
// Solidity: function getMinBalanceForUpkeep(uint256 upkeepId) view returns(uint96)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetMinBalanceForUpkeep(upkeepId *big.Int) (*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetMinBalanceForUpkeep(&_UpkeepControllerContract.CallOpts, upkeepId)
}

// GetMinBalancesForUpkeeps is a free data retrieval call binding the contract method 0xdbe5dcdf.
//
// Solidity: function getMinBalancesForUpkeeps(uint256 offset, uint256 limit) view returns(uint96[])
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetMinBalancesForUpkeeps(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getMinBalancesForUpkeeps", offset, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMinBalancesForUpkeeps is a free data retrieval call binding the contract method 0xdbe5dcdf.
//
// Solidity: function getMinBalancesForUpkeeps(uint256 offset, uint256 limit) view returns(uint96[])
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetMinBalancesForUpkeeps(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetMinBalancesForUpkeeps(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetMinBalancesForUpkeeps is a free data retrieval call binding the contract method 0xdbe5dcdf.
//
// Solidity: function getMinBalancesForUpkeeps(uint256 offset, uint256 limit) view returns(uint96[])
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetMinBalancesForUpkeeps(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetMinBalancesForUpkeeps(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetPerformDataTest is a free data retrieval call binding the contract method 0x2515fb10.
//
// Solidity: function getPerformDataTest(address target, bytes checkData) view returns(uint256[])
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetPerformDataTest(opts *bind.CallOpts, target common.Address, checkData []byte) ([]*big.Int, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getPerformDataTest", target, checkData)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPerformDataTest is a free data retrieval call binding the contract method 0x2515fb10.
//
// Solidity: function getPerformDataTest(address target, bytes checkData) view returns(uint256[])
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetPerformDataTest(target common.Address, checkData []byte) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetPerformDataTest(&_UpkeepControllerContract.CallOpts, target, checkData)
}

// GetPerformDataTest is a free data retrieval call binding the contract method 0x2515fb10.
//
// Solidity: function getPerformDataTest(address target, bytes checkData) view returns(uint256[])
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetPerformDataTest(target common.Address, checkData []byte) ([]*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetPerformDataTest(&_UpkeepControllerContract.CallOpts, target, checkData)
}

// GetResultTest is a free data retrieval call binding the contract method 0xed1769dd.
//
// Solidity: function getResultTest(uint256 performOffset, uint256 performLimit, uint256[] performArray) pure returns(bool isNeeded, uint256 newOffset, uint256 newLimit)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetResultTest(opts *bind.CallOpts, performOffset *big.Int, performLimit *big.Int, performArray []*big.Int) (struct {
	IsNeeded  bool
	NewOffset *big.Int
	NewLimit  *big.Int
}, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getResultTest", performOffset, performLimit, performArray)

	outstruct := new(struct {
		IsNeeded  bool
		NewOffset *big.Int
		NewLimit  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsNeeded = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.NewOffset = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetResultTest is a free data retrieval call binding the contract method 0xed1769dd.
//
// Solidity: function getResultTest(uint256 performOffset, uint256 performLimit, uint256[] performArray) pure returns(bool isNeeded, uint256 newOffset, uint256 newLimit)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetResultTest(performOffset *big.Int, performLimit *big.Int, performArray []*big.Int) (struct {
	IsNeeded  bool
	NewOffset *big.Int
	NewLimit  *big.Int
}, error) {
	return _UpkeepControllerContract.Contract.GetResultTest(&_UpkeepControllerContract.CallOpts, performOffset, performLimit, performArray)
}

// GetResultTest is a free data retrieval call binding the contract method 0xed1769dd.
//
// Solidity: function getResultTest(uint256 performOffset, uint256 performLimit, uint256[] performArray) pure returns(bool isNeeded, uint256 newOffset, uint256 newLimit)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetResultTest(performOffset *big.Int, performLimit *big.Int, performArray []*big.Int) (struct {
	IsNeeded  bool
	NewOffset *big.Int
	NewLimit  *big.Int
}, error) {
	return _UpkeepControllerContract.Contract.GetResultTest(&_UpkeepControllerContract.CallOpts, performOffset, performLimit, performArray)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((uint32,uint96,uint256,uint96,uint256,uint32,uint32,bytes32,uint32,bool) state, (uint32,uint32,uint32,uint24,uint16,uint96,uint32,uint32,uint32,uint256,uint256,address,address) config, address[] signers, address[] transmitters, uint8 f)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetState(opts *bind.CallOpts) (struct {
	State        State
	Config       OnchainConfig
	Signers      []common.Address
	Transmitters []common.Address
	F            uint8
}, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getState")

	outstruct := new(struct {
		State        State
		Config       OnchainConfig
		Signers      []common.Address
		Transmitters []common.Address
		F            uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(State)).(*State)
	outstruct.Config = *abi.ConvertType(out[1], new(OnchainConfig)).(*OnchainConfig)
	outstruct.Signers = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.Transmitters = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)
	outstruct.F = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((uint32,uint96,uint256,uint96,uint256,uint32,uint32,bytes32,uint32,bool) state, (uint32,uint32,uint32,uint24,uint16,uint96,uint32,uint32,uint32,uint256,uint256,address,address) config, address[] signers, address[] transmitters, uint8 f)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetState() (struct {
	State        State
	Config       OnchainConfig
	Signers      []common.Address
	Transmitters []common.Address
	F            uint8
}, error) {
	return _UpkeepControllerContract.Contract.GetState(&_UpkeepControllerContract.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((uint32,uint96,uint256,uint96,uint256,uint32,uint32,bytes32,uint32,bool) state, (uint32,uint32,uint32,uint24,uint16,uint96,uint32,uint32,uint32,uint256,uint256,address,address) config, address[] signers, address[] transmitters, uint8 f)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetState() (struct {
	State        State
	Config       OnchainConfig
	Signers      []common.Address
	Transmitters []common.Address
	F            uint8
}, error) {
	return _UpkeepControllerContract.Contract.GetState(&_UpkeepControllerContract.CallOpts)
}

// GetUpkeep is a free data retrieval call binding the contract method 0xc7c3a19a.
//
// Solidity: function getUpkeep(uint256 upkeepId) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes) upkeepInfo)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetUpkeep(opts *bind.CallOpts, upkeepId *big.Int) (UpkeepInfo, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getUpkeep", upkeepId)

	if err != nil {
		return *new(UpkeepInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UpkeepInfo)).(*UpkeepInfo)

	return out0, err

}

// GetUpkeep is a free data retrieval call binding the contract method 0xc7c3a19a.
//
// Solidity: function getUpkeep(uint256 upkeepId) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes) upkeepInfo)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetUpkeep(upkeepId *big.Int) (UpkeepInfo, error) {
	return _UpkeepControllerContract.Contract.GetUpkeep(&_UpkeepControllerContract.CallOpts, upkeepId)
}

// GetUpkeep is a free data retrieval call binding the contract method 0xc7c3a19a.
//
// Solidity: function getUpkeep(uint256 upkeepId) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes) upkeepInfo)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetUpkeep(upkeepId *big.Int) (UpkeepInfo, error) {
	return _UpkeepControllerContract.Contract.GetUpkeep(&_UpkeepControllerContract.CallOpts, upkeepId)
}

// GetUpkeeps is a free data retrieval call binding the contract method 0x7f25c83d.
//
// Solidity: function getUpkeeps(uint256 offset, uint256 limit) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes)[])
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetUpkeeps(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]UpkeepInfo, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getUpkeeps", offset, limit)

	if err != nil {
		return *new([]UpkeepInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]UpkeepInfo)).(*[]UpkeepInfo)

	return out0, err

}

// GetUpkeeps is a free data retrieval call binding the contract method 0x7f25c83d.
//
// Solidity: function getUpkeeps(uint256 offset, uint256 limit) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes)[])
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetUpkeeps(offset *big.Int, limit *big.Int) ([]UpkeepInfo, error) {
	return _UpkeepControllerContract.Contract.GetUpkeeps(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetUpkeeps is a free data retrieval call binding the contract method 0x7f25c83d.
//
// Solidity: function getUpkeeps(uint256 offset, uint256 limit) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes)[])
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetUpkeeps(offset *big.Int, limit *big.Int) ([]UpkeepInfo, error) {
	return _UpkeepControllerContract.Contract.GetUpkeeps(&_UpkeepControllerContract.CallOpts, offset, limit)
}

// GetUpkeepsCount is a free data retrieval call binding the contract method 0xe357cd07.
//
// Solidity: function getUpkeepsCount() view returns(uint256)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) GetUpkeepsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "getUpkeepsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpkeepsCount is a free data retrieval call binding the contract method 0xe357cd07.
//
// Solidity: function getUpkeepsCount() view returns(uint256)
func (_UpkeepControllerContract *UpkeepControllerContractSession) GetUpkeepsCount() (*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetUpkeepsCount(&_UpkeepControllerContract.CallOpts)
}

// GetUpkeepsCount is a free data retrieval call binding the contract method 0xe357cd07.
//
// Solidity: function getUpkeepsCount() view returns(uint256)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) GetUpkeepsCount() (*big.Int, error) {
	return _UpkeepControllerContract.Contract.GetUpkeepsCount(&_UpkeepControllerContract.CallOpts)
}

// ILink is a free data retrieval call binding the contract method 0x7d253aff.
//
// Solidity: function i_link() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ILink is a free data retrieval call binding the contract method 0x7d253aff.
//
// Solidity: function i_link() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractSession) ILink() (common.Address, error) {
	return _UpkeepControllerContract.Contract.ILink(&_UpkeepControllerContract.CallOpts)
}

// ILink is a free data retrieval call binding the contract method 0x7d253aff.
//
// Solidity: function i_link() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) ILink() (common.Address, error) {
	return _UpkeepControllerContract.Contract.ILink(&_UpkeepControllerContract.CallOpts)
}

// IRegistrar is a free data retrieval call binding the contract method 0x442b1278.
//
// Solidity: function i_registrar() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) IRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "i_registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IRegistrar is a free data retrieval call binding the contract method 0x442b1278.
//
// Solidity: function i_registrar() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractSession) IRegistrar() (common.Address, error) {
	return _UpkeepControllerContract.Contract.IRegistrar(&_UpkeepControllerContract.CallOpts)
}

// IRegistrar is a free data retrieval call binding the contract method 0x442b1278.
//
// Solidity: function i_registrar() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) IRegistrar() (common.Address, error) {
	return _UpkeepControllerContract.Contract.IRegistrar(&_UpkeepControllerContract.CallOpts)
}

// IRegistry is a free data retrieval call binding the contract method 0x2a663606.
//
// Solidity: function i_registry() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) IRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "i_registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IRegistry is a free data retrieval call binding the contract method 0x2a663606.
//
// Solidity: function i_registry() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractSession) IRegistry() (common.Address, error) {
	return _UpkeepControllerContract.Contract.IRegistry(&_UpkeepControllerContract.CallOpts)
}

// IRegistry is a free data retrieval call binding the contract method 0x2a663606.
//
// Solidity: function i_registry() view returns(address)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) IRegistry() (common.Address, error) {
	return _UpkeepControllerContract.Contract.IRegistry(&_UpkeepControllerContract.CallOpts)
}

// IsNewUpkeepNeeded is a free data retrieval call binding the contract method 0x37d10eac.
//
// Solidity: function isNewUpkeepNeeded() view returns(bool isNeeded, uint256 newOffset, uint256 newLimit)
func (_UpkeepControllerContract *UpkeepControllerContractCaller) IsNewUpkeepNeeded(opts *bind.CallOpts) (struct {
	IsNeeded  bool
	NewOffset *big.Int
	NewLimit  *big.Int
}, error) {
	var out []interface{}
	err := _UpkeepControllerContract.contract.Call(opts, &out, "isNewUpkeepNeeded")

	outstruct := new(struct {
		IsNeeded  bool
		NewOffset *big.Int
		NewLimit  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsNeeded = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.NewOffset = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IsNewUpkeepNeeded is a free data retrieval call binding the contract method 0x37d10eac.
//
// Solidity: function isNewUpkeepNeeded() view returns(bool isNeeded, uint256 newOffset, uint256 newLimit)
func (_UpkeepControllerContract *UpkeepControllerContractSession) IsNewUpkeepNeeded() (struct {
	IsNeeded  bool
	NewOffset *big.Int
	NewLimit  *big.Int
}, error) {
	return _UpkeepControllerContract.Contract.IsNewUpkeepNeeded(&_UpkeepControllerContract.CallOpts)
}

// IsNewUpkeepNeeded is a free data retrieval call binding the contract method 0x37d10eac.
//
// Solidity: function isNewUpkeepNeeded() view returns(bool isNeeded, uint256 newOffset, uint256 newLimit)
func (_UpkeepControllerContract *UpkeepControllerContractCallerSession) IsNewUpkeepNeeded() (struct {
	IsNeeded  bool
	NewOffset *big.Int
	NewLimit  *big.Int
}, error) {
	return _UpkeepControllerContract.Contract.IsNewUpkeepNeeded(&_UpkeepControllerContract.CallOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0x948108f7.
//
// Solidity: function addFunds(uint256 upkeepId, uint96 amount) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) AddFunds(opts *bind.TransactOpts, upkeepId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "addFunds", upkeepId, amount)
}

// AddFunds is a paid mutator transaction binding the contract method 0x948108f7.
//
// Solidity: function addFunds(uint256 upkeepId, uint96 amount) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) AddFunds(upkeepId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.AddFunds(&_UpkeepControllerContract.TransactOpts, upkeepId, amount)
}

// AddFunds is a paid mutator transaction binding the contract method 0x948108f7.
//
// Solidity: function addFunds(uint256 upkeepId, uint96 amount) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) AddFunds(upkeepId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.AddFunds(&_UpkeepControllerContract.TransactOpts, upkeepId, amount)
}

// CancelUpkeep is a paid mutator transaction binding the contract method 0xc8048022.
//
// Solidity: function cancelUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) CancelUpkeep(opts *bind.TransactOpts, upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "cancelUpkeep", upkeepId)
}

// CancelUpkeep is a paid mutator transaction binding the contract method 0xc8048022.
//
// Solidity: function cancelUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) CancelUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.CancelUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// CancelUpkeep is a paid mutator transaction binding the contract method 0xc8048022.
//
// Solidity: function cancelUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) CancelUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.CancelUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0xf7d334ba.
//
// Solidity: function checkUpkeep(uint256 upkeepId) returns(bool upkeepNeeded, bytes performData, uint8 upkeepFailureReason, uint256 gasUsed, uint256 fastGasWei, uint256 linkNative)
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) CheckUpkeep(opts *bind.TransactOpts, upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "checkUpkeep", upkeepId)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0xf7d334ba.
//
// Solidity: function checkUpkeep(uint256 upkeepId) returns(bool upkeepNeeded, bytes performData, uint8 upkeepFailureReason, uint256 gasUsed, uint256 fastGasWei, uint256 linkNative)
func (_UpkeepControllerContract *UpkeepControllerContractSession) CheckUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.CheckUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0xf7d334ba.
//
// Solidity: function checkUpkeep(uint256 upkeepId) returns(bool upkeepNeeded, bytes performData, uint8 upkeepFailureReason, uint256 gasUsed, uint256 fastGasWei, uint256 linkNative)
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) CheckUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.CheckUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// PauseUpkeep is a paid mutator transaction binding the contract method 0x8765ecbe.
//
// Solidity: function pauseUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) PauseUpkeep(opts *bind.TransactOpts, upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "pauseUpkeep", upkeepId)
}

// PauseUpkeep is a paid mutator transaction binding the contract method 0x8765ecbe.
//
// Solidity: function pauseUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) PauseUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.PauseUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// PauseUpkeep is a paid mutator transaction binding the contract method 0x8765ecbe.
//
// Solidity: function pauseUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) PauseUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.PauseUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// RegisterAndPredictID is a paid mutator transaction binding the contract method 0xfd111cca.
//
// Solidity: function registerAndPredictID((string,bytes,address,uint32,address,bytes,bytes,uint96) params) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) RegisterAndPredictID(opts *bind.TransactOpts, params RegistrationParams) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "registerAndPredictID", params)
}

// RegisterAndPredictID is a paid mutator transaction binding the contract method 0xfd111cca.
//
// Solidity: function registerAndPredictID((string,bytes,address,uint32,address,bytes,bytes,uint96) params) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) RegisterAndPredictID(params RegistrationParams) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.RegisterAndPredictID(&_UpkeepControllerContract.TransactOpts, params)
}

// RegisterAndPredictID is a paid mutator transaction binding the contract method 0xfd111cca.
//
// Solidity: function registerAndPredictID((string,bytes,address,uint32,address,bytes,bytes,uint96) params) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) RegisterAndPredictID(params RegistrationParams) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.RegisterAndPredictID(&_UpkeepControllerContract.TransactOpts, params)
}

// SetUpkeepGasLimit is a paid mutator transaction binding the contract method 0xa72aa27e.
//
// Solidity: function setUpkeepGasLimit(uint256 upkeepId, uint32 gasLimit) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) SetUpkeepGasLimit(opts *bind.TransactOpts, upkeepId *big.Int, gasLimit uint32) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "setUpkeepGasLimit", upkeepId, gasLimit)
}

// SetUpkeepGasLimit is a paid mutator transaction binding the contract method 0xa72aa27e.
//
// Solidity: function setUpkeepGasLimit(uint256 upkeepId, uint32 gasLimit) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) SetUpkeepGasLimit(upkeepId *big.Int, gasLimit uint32) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.SetUpkeepGasLimit(&_UpkeepControllerContract.TransactOpts, upkeepId, gasLimit)
}

// SetUpkeepGasLimit is a paid mutator transaction binding the contract method 0xa72aa27e.
//
// Solidity: function setUpkeepGasLimit(uint256 upkeepId, uint32 gasLimit) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) SetUpkeepGasLimit(upkeepId *big.Int, gasLimit uint32) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.SetUpkeepGasLimit(&_UpkeepControllerContract.TransactOpts, upkeepId, gasLimit)
}

// SetUpkeepOffchainConfig is a paid mutator transaction binding the contract method 0x8dcf0fe7.
//
// Solidity: function setUpkeepOffchainConfig(uint256 upkeepId, bytes config) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) SetUpkeepOffchainConfig(opts *bind.TransactOpts, upkeepId *big.Int, config []byte) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "setUpkeepOffchainConfig", upkeepId, config)
}

// SetUpkeepOffchainConfig is a paid mutator transaction binding the contract method 0x8dcf0fe7.
//
// Solidity: function setUpkeepOffchainConfig(uint256 upkeepId, bytes config) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) SetUpkeepOffchainConfig(upkeepId *big.Int, config []byte) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.SetUpkeepOffchainConfig(&_UpkeepControllerContract.TransactOpts, upkeepId, config)
}

// SetUpkeepOffchainConfig is a paid mutator transaction binding the contract method 0x8dcf0fe7.
//
// Solidity: function setUpkeepOffchainConfig(uint256 upkeepId, bytes config) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) SetUpkeepOffchainConfig(upkeepId *big.Int, config []byte) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.SetUpkeepOffchainConfig(&_UpkeepControllerContract.TransactOpts, upkeepId, config)
}

// UnpauseUpkeep is a paid mutator transaction binding the contract method 0x5165f2f5.
//
// Solidity: function unpauseUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) UnpauseUpkeep(opts *bind.TransactOpts, upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "unpauseUpkeep", upkeepId)
}

// UnpauseUpkeep is a paid mutator transaction binding the contract method 0x5165f2f5.
//
// Solidity: function unpauseUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) UnpauseUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.UnpauseUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// UnpauseUpkeep is a paid mutator transaction binding the contract method 0x5165f2f5.
//
// Solidity: function unpauseUpkeep(uint256 upkeepId) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) UnpauseUpkeep(upkeepId *big.Int) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.UnpauseUpkeep(&_UpkeepControllerContract.TransactOpts, upkeepId)
}

// UpdateCheckData is a paid mutator transaction binding the contract method 0x9fab4386.
//
// Solidity: function updateCheckData(uint256 upkeepId, bytes newCheckData) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactor) UpdateCheckData(opts *bind.TransactOpts, upkeepId *big.Int, newCheckData []byte) (*types.Transaction, error) {
	return _UpkeepControllerContract.contract.Transact(opts, "updateCheckData", upkeepId, newCheckData)
}

// UpdateCheckData is a paid mutator transaction binding the contract method 0x9fab4386.
//
// Solidity: function updateCheckData(uint256 upkeepId, bytes newCheckData) returns()
func (_UpkeepControllerContract *UpkeepControllerContractSession) UpdateCheckData(upkeepId *big.Int, newCheckData []byte) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.UpdateCheckData(&_UpkeepControllerContract.TransactOpts, upkeepId, newCheckData)
}

// UpdateCheckData is a paid mutator transaction binding the contract method 0x9fab4386.
//
// Solidity: function updateCheckData(uint256 upkeepId, bytes newCheckData) returns()
func (_UpkeepControllerContract *UpkeepControllerContractTransactorSession) UpdateCheckData(upkeepId *big.Int, newCheckData []byte) (*types.Transaction, error) {
	return _UpkeepControllerContract.Contract.UpdateCheckData(&_UpkeepControllerContract.TransactOpts, upkeepId, newCheckData)
}

// UpkeepControllerContractFundsAddedIterator is returned from FilterFundsAdded and is used to iterate over the raw logs and unpacked data for FundsAdded events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractFundsAddedIterator struct {
	Event *UpkeepControllerContractFundsAdded // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractFundsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractFundsAdded)
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
		it.Event = new(UpkeepControllerContractFundsAdded)
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
func (it *UpkeepControllerContractFundsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractFundsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractFundsAdded represents a FundsAdded event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractFundsAdded struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsAdded is a free log retrieval operation binding the contract event 0x8137dc366612bf502338bd8951f835ad8ceba421c4eb3d79c7f9b3ce0ac4762e.
//
// Solidity: event FundsAdded(uint256 indexed id, uint96 amount)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterFundsAdded(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractFundsAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "FundsAdded", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractFundsAddedIterator{contract: _UpkeepControllerContract.contract, event: "FundsAdded", logs: logs, sub: sub}, nil
}

// WatchFundsAdded is a free log subscription operation binding the contract event 0x8137dc366612bf502338bd8951f835ad8ceba421c4eb3d79c7f9b3ce0ac4762e.
//
// Solidity: event FundsAdded(uint256 indexed id, uint96 amount)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractFundsAdded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "FundsAdded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractFundsAdded)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "FundsAdded", log); err != nil {
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

// ParseFundsAdded is a log parse operation binding the contract event 0x8137dc366612bf502338bd8951f835ad8ceba421c4eb3d79c7f9b3ce0ac4762e.
//
// Solidity: event FundsAdded(uint256 indexed id, uint96 amount)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseFundsAdded(log types.Log) (*UpkeepControllerContractFundsAdded, error) {
	event := new(UpkeepControllerContractFundsAdded)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "FundsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpkeepControllerContractUpkeepCanceledIterator is returned from FilterUpkeepCanceled and is used to iterate over the raw logs and unpacked data for UpkeepCanceled events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepCanceledIterator struct {
	Event *UpkeepControllerContractUpkeepCanceled // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractUpkeepCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractUpkeepCanceled)
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
		it.Event = new(UpkeepControllerContractUpkeepCanceled)
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
func (it *UpkeepControllerContractUpkeepCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractUpkeepCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractUpkeepCanceled represents a UpkeepCanceled event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepCanceled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpkeepCanceled is a free log retrieval operation binding the contract event 0xa08e5e152c9438f8441ddfe294ad3b09457b826764a94e320a40b228cf2d410a.
//
// Solidity: event UpkeepCanceled(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterUpkeepCanceled(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractUpkeepCanceledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "UpkeepCanceled", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractUpkeepCanceledIterator{contract: _UpkeepControllerContract.contract, event: "UpkeepCanceled", logs: logs, sub: sub}, nil
}

// WatchUpkeepCanceled is a free log subscription operation binding the contract event 0xa08e5e152c9438f8441ddfe294ad3b09457b826764a94e320a40b228cf2d410a.
//
// Solidity: event UpkeepCanceled(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchUpkeepCanceled(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractUpkeepCanceled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "UpkeepCanceled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractUpkeepCanceled)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepCanceled", log); err != nil {
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

// ParseUpkeepCanceled is a log parse operation binding the contract event 0xa08e5e152c9438f8441ddfe294ad3b09457b826764a94e320a40b228cf2d410a.
//
// Solidity: event UpkeepCanceled(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseUpkeepCanceled(log types.Log) (*UpkeepControllerContractUpkeepCanceled, error) {
	event := new(UpkeepControllerContractUpkeepCanceled)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpkeepControllerContractUpkeepCreatedIterator is returned from FilterUpkeepCreated and is used to iterate over the raw logs and unpacked data for UpkeepCreated events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepCreatedIterator struct {
	Event *UpkeepControllerContractUpkeepCreated // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractUpkeepCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractUpkeepCreated)
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
		it.Event = new(UpkeepControllerContractUpkeepCreated)
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
func (it *UpkeepControllerContractUpkeepCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractUpkeepCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractUpkeepCreated represents a UpkeepCreated event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepCreated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpkeepCreated is a free log retrieval operation binding the contract event 0x73827404e930da1450aa92d0348dd0939840f4850415c6b42a1bcf9de4162d7b.
//
// Solidity: event UpkeepCreated(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterUpkeepCreated(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractUpkeepCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "UpkeepCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractUpkeepCreatedIterator{contract: _UpkeepControllerContract.contract, event: "UpkeepCreated", logs: logs, sub: sub}, nil
}

// WatchUpkeepCreated is a free log subscription operation binding the contract event 0x73827404e930da1450aa92d0348dd0939840f4850415c6b42a1bcf9de4162d7b.
//
// Solidity: event UpkeepCreated(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchUpkeepCreated(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractUpkeepCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "UpkeepCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractUpkeepCreated)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepCreated", log); err != nil {
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
// Solidity: event UpkeepCreated(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseUpkeepCreated(log types.Log) (*UpkeepControllerContractUpkeepCreated, error) {
	event := new(UpkeepControllerContractUpkeepCreated)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpkeepControllerContractUpkeepGasLimitSetIterator is returned from FilterUpkeepGasLimitSet and is used to iterate over the raw logs and unpacked data for UpkeepGasLimitSet events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepGasLimitSetIterator struct {
	Event *UpkeepControllerContractUpkeepGasLimitSet // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractUpkeepGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractUpkeepGasLimitSet)
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
		it.Event = new(UpkeepControllerContractUpkeepGasLimitSet)
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
func (it *UpkeepControllerContractUpkeepGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractUpkeepGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractUpkeepGasLimitSet represents a UpkeepGasLimitSet event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepGasLimitSet struct {
	Id     *big.Int
	Amount uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpkeepGasLimitSet is a free log retrieval operation binding the contract event 0x6e922e3c48062403590addb8e090d556398efe9371d97f454fc4b10b562ae590.
//
// Solidity: event UpkeepGasLimitSet(uint256 indexed id, uint32 amount)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterUpkeepGasLimitSet(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractUpkeepGasLimitSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "UpkeepGasLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractUpkeepGasLimitSetIterator{contract: _UpkeepControllerContract.contract, event: "UpkeepGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchUpkeepGasLimitSet is a free log subscription operation binding the contract event 0x6e922e3c48062403590addb8e090d556398efe9371d97f454fc4b10b562ae590.
//
// Solidity: event UpkeepGasLimitSet(uint256 indexed id, uint32 amount)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchUpkeepGasLimitSet(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractUpkeepGasLimitSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "UpkeepGasLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractUpkeepGasLimitSet)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepGasLimitSet", log); err != nil {
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

// ParseUpkeepGasLimitSet is a log parse operation binding the contract event 0x6e922e3c48062403590addb8e090d556398efe9371d97f454fc4b10b562ae590.
//
// Solidity: event UpkeepGasLimitSet(uint256 indexed id, uint32 amount)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseUpkeepGasLimitSet(log types.Log) (*UpkeepControllerContractUpkeepGasLimitSet, error) {
	event := new(UpkeepControllerContractUpkeepGasLimitSet)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpkeepControllerContractUpkeepOffchainConfigSetIterator is returned from FilterUpkeepOffchainConfigSet and is used to iterate over the raw logs and unpacked data for UpkeepOffchainConfigSet events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepOffchainConfigSetIterator struct {
	Event *UpkeepControllerContractUpkeepOffchainConfigSet // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractUpkeepOffchainConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractUpkeepOffchainConfigSet)
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
		it.Event = new(UpkeepControllerContractUpkeepOffchainConfigSet)
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
func (it *UpkeepControllerContractUpkeepOffchainConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractUpkeepOffchainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractUpkeepOffchainConfigSet represents a UpkeepOffchainConfigSet event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepOffchainConfigSet struct {
	Id     *big.Int
	Config []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpkeepOffchainConfigSet is a free log retrieval operation binding the contract event 0x3e8740446213c8a77d40e08f79136ce3f347d13ed270a6ebdf57159e0faf4850.
//
// Solidity: event UpkeepOffchainConfigSet(uint256 indexed id, bytes config)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterUpkeepOffchainConfigSet(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractUpkeepOffchainConfigSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "UpkeepOffchainConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractUpkeepOffchainConfigSetIterator{contract: _UpkeepControllerContract.contract, event: "UpkeepOffchainConfigSet", logs: logs, sub: sub}, nil
}

// WatchUpkeepOffchainConfigSet is a free log subscription operation binding the contract event 0x3e8740446213c8a77d40e08f79136ce3f347d13ed270a6ebdf57159e0faf4850.
//
// Solidity: event UpkeepOffchainConfigSet(uint256 indexed id, bytes config)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchUpkeepOffchainConfigSet(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractUpkeepOffchainConfigSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "UpkeepOffchainConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractUpkeepOffchainConfigSet)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepOffchainConfigSet", log); err != nil {
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

// ParseUpkeepOffchainConfigSet is a log parse operation binding the contract event 0x3e8740446213c8a77d40e08f79136ce3f347d13ed270a6ebdf57159e0faf4850.
//
// Solidity: event UpkeepOffchainConfigSet(uint256 indexed id, bytes config)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseUpkeepOffchainConfigSet(log types.Log) (*UpkeepControllerContractUpkeepOffchainConfigSet, error) {
	event := new(UpkeepControllerContractUpkeepOffchainConfigSet)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepOffchainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpkeepControllerContractUpkeepPausedIterator is returned from FilterUpkeepPaused and is used to iterate over the raw logs and unpacked data for UpkeepPaused events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepPausedIterator struct {
	Event *UpkeepControllerContractUpkeepPaused // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractUpkeepPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractUpkeepPaused)
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
		it.Event = new(UpkeepControllerContractUpkeepPaused)
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
func (it *UpkeepControllerContractUpkeepPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractUpkeepPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractUpkeepPaused represents a UpkeepPaused event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepPaused struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpkeepPaused is a free log retrieval operation binding the contract event 0x8ab10247ce168c27748e656ecf852b951fcaac790c18106b19aa0ae57a8b741f.
//
// Solidity: event UpkeepPaused(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterUpkeepPaused(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractUpkeepPausedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "UpkeepPaused", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractUpkeepPausedIterator{contract: _UpkeepControllerContract.contract, event: "UpkeepPaused", logs: logs, sub: sub}, nil
}

// WatchUpkeepPaused is a free log subscription operation binding the contract event 0x8ab10247ce168c27748e656ecf852b951fcaac790c18106b19aa0ae57a8b741f.
//
// Solidity: event UpkeepPaused(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchUpkeepPaused(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractUpkeepPaused, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "UpkeepPaused", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractUpkeepPaused)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepPaused", log); err != nil {
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

// ParseUpkeepPaused is a log parse operation binding the contract event 0x8ab10247ce168c27748e656ecf852b951fcaac790c18106b19aa0ae57a8b741f.
//
// Solidity: event UpkeepPaused(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseUpkeepPaused(log types.Log) (*UpkeepControllerContractUpkeepPaused, error) {
	event := new(UpkeepControllerContractUpkeepPaused)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpkeepControllerContractUpkeepUnpausedIterator is returned from FilterUpkeepUnpaused and is used to iterate over the raw logs and unpacked data for UpkeepUnpaused events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepUnpausedIterator struct {
	Event *UpkeepControllerContractUpkeepUnpaused // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractUpkeepUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractUpkeepUnpaused)
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
		it.Event = new(UpkeepControllerContractUpkeepUnpaused)
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
func (it *UpkeepControllerContractUpkeepUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractUpkeepUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractUpkeepUnpaused represents a UpkeepUnpaused event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepUnpaused struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpkeepUnpaused is a free log retrieval operation binding the contract event 0x7bada562044eb163f6b4003c4553e4e62825344c0418eea087bed5ee05a47456.
//
// Solidity: event UpkeepUnpaused(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterUpkeepUnpaused(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractUpkeepUnpausedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "UpkeepUnpaused", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractUpkeepUnpausedIterator{contract: _UpkeepControllerContract.contract, event: "UpkeepUnpaused", logs: logs, sub: sub}, nil
}

// WatchUpkeepUnpaused is a free log subscription operation binding the contract event 0x7bada562044eb163f6b4003c4553e4e62825344c0418eea087bed5ee05a47456.
//
// Solidity: event UpkeepUnpaused(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchUpkeepUnpaused(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractUpkeepUnpaused, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "UpkeepUnpaused", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractUpkeepUnpaused)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepUnpaused", log); err != nil {
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

// ParseUpkeepUnpaused is a log parse operation binding the contract event 0x7bada562044eb163f6b4003c4553e4e62825344c0418eea087bed5ee05a47456.
//
// Solidity: event UpkeepUnpaused(uint256 indexed id)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseUpkeepUnpaused(log types.Log) (*UpkeepControllerContractUpkeepUnpaused, error) {
	event := new(UpkeepControllerContractUpkeepUnpaused)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpkeepControllerContractUpkeepUpdatedIterator is returned from FilterUpkeepUpdated and is used to iterate over the raw logs and unpacked data for UpkeepUpdated events raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepUpdatedIterator struct {
	Event *UpkeepControllerContractUpkeepUpdated // Event containing the contract specifics and raw log

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
func (it *UpkeepControllerContractUpkeepUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpkeepControllerContractUpkeepUpdated)
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
		it.Event = new(UpkeepControllerContractUpkeepUpdated)
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
func (it *UpkeepControllerContractUpkeepUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpkeepControllerContractUpkeepUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpkeepControllerContractUpkeepUpdated represents a UpkeepUpdated event raised by the UpkeepControllerContract contract.
type UpkeepControllerContractUpkeepUpdated struct {
	Id           *big.Int
	NewCheckData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpkeepUpdated is a free log retrieval operation binding the contract event 0x86aa972e9fc0c705fea6a70adc998117ea6a33853d380b06d3a4e777f4e5203a.
//
// Solidity: event UpkeepUpdated(uint256 indexed id, bytes newCheckData)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) FilterUpkeepUpdated(opts *bind.FilterOpts, id []*big.Int) (*UpkeepControllerContractUpkeepUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.FilterLogs(opts, "UpkeepUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &UpkeepControllerContractUpkeepUpdatedIterator{contract: _UpkeepControllerContract.contract, event: "UpkeepUpdated", logs: logs, sub: sub}, nil
}

// WatchUpkeepUpdated is a free log subscription operation binding the contract event 0x86aa972e9fc0c705fea6a70adc998117ea6a33853d380b06d3a4e777f4e5203a.
//
// Solidity: event UpkeepUpdated(uint256 indexed id, bytes newCheckData)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) WatchUpkeepUpdated(opts *bind.WatchOpts, sink chan<- *UpkeepControllerContractUpkeepUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _UpkeepControllerContract.contract.WatchLogs(opts, "UpkeepUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpkeepControllerContractUpkeepUpdated)
				if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepUpdated", log); err != nil {
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

// ParseUpkeepUpdated is a log parse operation binding the contract event 0x86aa972e9fc0c705fea6a70adc998117ea6a33853d380b06d3a4e777f4e5203a.
//
// Solidity: event UpkeepUpdated(uint256 indexed id, bytes newCheckData)
func (_UpkeepControllerContract *UpkeepControllerContractFilterer) ParseUpkeepUpdated(log types.Log) (*UpkeepControllerContractUpkeepUpdated, error) {
	event := new(UpkeepControllerContractUpkeepUpdated)
	if err := _UpkeepControllerContract.contract.UnpackLog(event, "UpkeepUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
