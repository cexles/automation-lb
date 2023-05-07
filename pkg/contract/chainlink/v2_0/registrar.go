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

// RegistrarContractMetaData contains all meta data concerning the RegistrarContract contract.
var RegistrarContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractKeeperRegistryBase2_0\",\"name\":\"keeperRegistryLogic\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ArrayHasNoEntries\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CheckDataExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateEntry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasLimitCanOnlyIncrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasLimitOutsideRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfFaultyOracles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCheckDataSizeCanOnlyIncrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxPerformDataSizeCanOnlyIncrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MigrationNotPermitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnchainConfigNonEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyActiveSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyActiveTransmitters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByLINKToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrRegistrar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByPayee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByProposedAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByProposedPayee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPausedUpkeep\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySimulatedBackend\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpausedUpkeep\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParameterLengthError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentGreaterThanAllLINK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepeatedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepeatedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"TargetCheckReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyOracles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TranscoderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepNotCanceled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepNotNeeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotChanged\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CancelledUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"FundsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"InsufficientFundsUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"OwnerFundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"PayeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"PaymentWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ReorgedUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"StaleUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"UpkeepAdminTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"UpkeepAdminTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"atBlockHeight\",\"type\":\"uint64\"}],\"name\":\"UpkeepCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newCheckData\",\"type\":\"bytes\"}],\"name\":\"UpkeepCheckDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"gasLimit\",\"type\":\"uint96\"}],\"name\":\"UpkeepGasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"UpkeepMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"UpkeepOffchainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"checkBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"totalPayment\",\"type\":\"uint96\"}],\"name\":\"UpkeepPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"importedFrom\",\"type\":\"address\"}],\"name\":\"UpkeepReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"executeGas\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"UpkeepRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepUnpaused\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"acceptUpkeepAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"addFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancelUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"},{\"internalType\":\"enumUpkeepFailureReason\",\"name\":\"upkeepFailureReason\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastGasWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkNative\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"getActiveUpkeepIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFastGasFeedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKeeperRegistryLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkNativeFeedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"getMaxPaymentForGas\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"maxPayment\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getMinBalanceForUpkeep\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"minBalance\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentModel\",\"outputs\":[{\"internalType\":\"enumKeeperRegistryBase2_0.PaymentModel\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"peer\",\"type\":\"address\"}],\"name\":\"getPeerRegistryMigrationPermission\",\"outputs\":[{\"internalType\":\"enumKeeperRegistryBase2_0.MigrationPermission\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"query\",\"type\":\"address\"}],\"name\":\"getSignerInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"ownerLinkBalance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"expectedLinkBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"totalPremium\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"numUpkeeps\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"latestConfigBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"latestConfigDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"latestEpoch\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"internalType\":\"structState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"paymentPremiumPPB\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flatFeeMicroLink\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"checkGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"stalenessSeconds\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"gasCeilingMultiplier\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"minUpkeepSpend\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"maxPerformGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxCheckDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerformDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fallbackGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fallbackLinkPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"}],\"internalType\":\"structOnchainConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"query\",\"type\":\"address\"}],\"name\":\"getTransmitterInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"lastCollected\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getUpkeep\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executeGas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxValidBlocknumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"lastPerformBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"amountSpent\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"internalType\":\"structUpkeepInfo\",\"name\":\"upkeepInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"migrateUpkeeps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pauseUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedUpkeeps\",\"type\":\"bytes\"}],\"name\":\"receiveUpkeeps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"registerUpkeep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"peer\",\"type\":\"address\"},{\"internalType\":\"enumKeeperRegistryBase2_0.MigrationPermission\",\"name\":\"permission\",\"type\":\"uint8\"}],\"name\":\"setPeerRegistryMigrationPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"setUpkeepGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"name\":\"setUpkeepOffchainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"name\":\"simulatePerformUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferUpkeepAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"rawReport\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"unpauseUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"newCheckData\",\"type\":\"bytes\"}],\"name\":\"updateCheckData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upkeepTranscoderVersion\",\"outputs\":[{\"internalType\":\"enumUpkeepFormat\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawOwnerFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RegistrarContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrarContractMetaData.ABI instead.
var RegistrarContractABI = RegistrarContractMetaData.ABI

// RegistrarContract is an auto generated Go binding around an Ethereum contract.
type RegistrarContract struct {
	RegistrarContractCaller     // Read-only binding to the contract
	RegistrarContractTransactor // Write-only binding to the contract
	RegistrarContractFilterer   // Log filterer for contract events
}

// RegistrarContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrarContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarContractSession struct {
	Contract     *RegistrarContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RegistrarContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarContractCallerSession struct {
	Contract *RegistrarContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RegistrarContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarContractTransactorSession struct {
	Contract     *RegistrarContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RegistrarContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarContractRaw struct {
	Contract *RegistrarContract // Generic contract binding to access the raw methods on
}

// RegistrarContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarContractCallerRaw struct {
	Contract *RegistrarContractCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarContractTransactorRaw struct {
	Contract *RegistrarContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrarContract creates a new instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContract(address common.Address, backend bind.ContractBackend) (*RegistrarContract, error) {
	contract, err := bindRegistrarContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistrarContract{RegistrarContractCaller: RegistrarContractCaller{contract: contract}, RegistrarContractTransactor: RegistrarContractTransactor{contract: contract}, RegistrarContractFilterer: RegistrarContractFilterer{contract: contract}}, nil
}

// NewRegistrarContractCaller creates a new read-only instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContractCaller(address common.Address, caller bind.ContractCaller) (*RegistrarContractCaller, error) {
	contract, err := bindRegistrarContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractCaller{contract: contract}, nil
}

// NewRegistrarContractTransactor creates a new write-only instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarContractTransactor, error) {
	contract, err := bindRegistrarContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractTransactor{contract: contract}, nil
}

// NewRegistrarContractFilterer creates a new log filterer instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrarContractFilterer, error) {
	contract, err := bindRegistrarContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractFilterer{contract: contract}, nil
}

// bindRegistrarContract binds a generic wrapper to an already deployed contract.
func bindRegistrarContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrarContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrarContract *RegistrarContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrarContract.Contract.RegistrarContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrarContract *RegistrarContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.Contract.RegistrarContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrarContract *RegistrarContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrarContract.Contract.RegistrarContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrarContract *RegistrarContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrarContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrarContract *RegistrarContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrarContract *RegistrarContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrarContract.Contract.contract.Transact(opts, method, params...)
}

// GetActiveUpkeepIDs is a free data retrieval call binding the contract method 0x06e3b632.
//
// Solidity: function getActiveUpkeepIDs(uint256 startIndex, uint256 maxCount) view returns(uint256[])
func (_RegistrarContract *RegistrarContractCaller) GetActiveUpkeepIDs(opts *bind.CallOpts, startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getActiveUpkeepIDs", startIndex, maxCount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveUpkeepIDs is a free data retrieval call binding the contract method 0x06e3b632.
//
// Solidity: function getActiveUpkeepIDs(uint256 startIndex, uint256 maxCount) view returns(uint256[])
func (_RegistrarContract *RegistrarContractSession) GetActiveUpkeepIDs(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _RegistrarContract.Contract.GetActiveUpkeepIDs(&_RegistrarContract.CallOpts, startIndex, maxCount)
}

// GetActiveUpkeepIDs is a free data retrieval call binding the contract method 0x06e3b632.
//
// Solidity: function getActiveUpkeepIDs(uint256 startIndex, uint256 maxCount) view returns(uint256[])
func (_RegistrarContract *RegistrarContractCallerSession) GetActiveUpkeepIDs(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _RegistrarContract.Contract.GetActiveUpkeepIDs(&_RegistrarContract.CallOpts, startIndex, maxCount)
}

// GetFastGasFeedAddress is a free data retrieval call binding the contract method 0x6709d0e5.
//
// Solidity: function getFastGasFeedAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCaller) GetFastGasFeedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getFastGasFeedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFastGasFeedAddress is a free data retrieval call binding the contract method 0x6709d0e5.
//
// Solidity: function getFastGasFeedAddress() view returns(address)
func (_RegistrarContract *RegistrarContractSession) GetFastGasFeedAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetFastGasFeedAddress(&_RegistrarContract.CallOpts)
}

// GetFastGasFeedAddress is a free data retrieval call binding the contract method 0x6709d0e5.
//
// Solidity: function getFastGasFeedAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCallerSession) GetFastGasFeedAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetFastGasFeedAddress(&_RegistrarContract.CallOpts)
}

// GetKeeperRegistryLogicAddress is a free data retrieval call binding the contract method 0x572e05e1.
//
// Solidity: function getKeeperRegistryLogicAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCaller) GetKeeperRegistryLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getKeeperRegistryLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeeperRegistryLogicAddress is a free data retrieval call binding the contract method 0x572e05e1.
//
// Solidity: function getKeeperRegistryLogicAddress() view returns(address)
func (_RegistrarContract *RegistrarContractSession) GetKeeperRegistryLogicAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetKeeperRegistryLogicAddress(&_RegistrarContract.CallOpts)
}

// GetKeeperRegistryLogicAddress is a free data retrieval call binding the contract method 0x572e05e1.
//
// Solidity: function getKeeperRegistryLogicAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCallerSession) GetKeeperRegistryLogicAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetKeeperRegistryLogicAddress(&_RegistrarContract.CallOpts)
}

// GetLinkAddress is a free data retrieval call binding the contract method 0xca30e603.
//
// Solidity: function getLinkAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCaller) GetLinkAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getLinkAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkAddress is a free data retrieval call binding the contract method 0xca30e603.
//
// Solidity: function getLinkAddress() view returns(address)
func (_RegistrarContract *RegistrarContractSession) GetLinkAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetLinkAddress(&_RegistrarContract.CallOpts)
}

// GetLinkAddress is a free data retrieval call binding the contract method 0xca30e603.
//
// Solidity: function getLinkAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCallerSession) GetLinkAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetLinkAddress(&_RegistrarContract.CallOpts)
}

// GetLinkNativeFeedAddress is a free data retrieval call binding the contract method 0xb10b673c.
//
// Solidity: function getLinkNativeFeedAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCaller) GetLinkNativeFeedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getLinkNativeFeedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkNativeFeedAddress is a free data retrieval call binding the contract method 0xb10b673c.
//
// Solidity: function getLinkNativeFeedAddress() view returns(address)
func (_RegistrarContract *RegistrarContractSession) GetLinkNativeFeedAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetLinkNativeFeedAddress(&_RegistrarContract.CallOpts)
}

// GetLinkNativeFeedAddress is a free data retrieval call binding the contract method 0xb10b673c.
//
// Solidity: function getLinkNativeFeedAddress() view returns(address)
func (_RegistrarContract *RegistrarContractCallerSession) GetLinkNativeFeedAddress() (common.Address, error) {
	return _RegistrarContract.Contract.GetLinkNativeFeedAddress(&_RegistrarContract.CallOpts)
}

// GetMaxPaymentForGas is a free data retrieval call binding the contract method 0x0e08ae84.
//
// Solidity: function getMaxPaymentForGas(uint32 gasLimit) view returns(uint96 maxPayment)
func (_RegistrarContract *RegistrarContractCaller) GetMaxPaymentForGas(opts *bind.CallOpts, gasLimit uint32) (*big.Int, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getMaxPaymentForGas", gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPaymentForGas is a free data retrieval call binding the contract method 0x0e08ae84.
//
// Solidity: function getMaxPaymentForGas(uint32 gasLimit) view returns(uint96 maxPayment)
func (_RegistrarContract *RegistrarContractSession) GetMaxPaymentForGas(gasLimit uint32) (*big.Int, error) {
	return _RegistrarContract.Contract.GetMaxPaymentForGas(&_RegistrarContract.CallOpts, gasLimit)
}

// GetMaxPaymentForGas is a free data retrieval call binding the contract method 0x0e08ae84.
//
// Solidity: function getMaxPaymentForGas(uint32 gasLimit) view returns(uint96 maxPayment)
func (_RegistrarContract *RegistrarContractCallerSession) GetMaxPaymentForGas(gasLimit uint32) (*big.Int, error) {
	return _RegistrarContract.Contract.GetMaxPaymentForGas(&_RegistrarContract.CallOpts, gasLimit)
}

// GetMinBalanceForUpkeep is a free data retrieval call binding the contract method 0xb657bc9c.
//
// Solidity: function getMinBalanceForUpkeep(uint256 id) view returns(uint96 minBalance)
func (_RegistrarContract *RegistrarContractCaller) GetMinBalanceForUpkeep(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getMinBalanceForUpkeep", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBalanceForUpkeep is a free data retrieval call binding the contract method 0xb657bc9c.
//
// Solidity: function getMinBalanceForUpkeep(uint256 id) view returns(uint96 minBalance)
func (_RegistrarContract *RegistrarContractSession) GetMinBalanceForUpkeep(id *big.Int) (*big.Int, error) {
	return _RegistrarContract.Contract.GetMinBalanceForUpkeep(&_RegistrarContract.CallOpts, id)
}

// GetMinBalanceForUpkeep is a free data retrieval call binding the contract method 0xb657bc9c.
//
// Solidity: function getMinBalanceForUpkeep(uint256 id) view returns(uint96 minBalance)
func (_RegistrarContract *RegistrarContractCallerSession) GetMinBalanceForUpkeep(id *big.Int) (*big.Int, error) {
	return _RegistrarContract.Contract.GetMinBalanceForUpkeep(&_RegistrarContract.CallOpts, id)
}

// GetPaymentModel is a free data retrieval call binding the contract method 0xf1570141.
//
// Solidity: function getPaymentModel() view returns(uint8)
func (_RegistrarContract *RegistrarContractCaller) GetPaymentModel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getPaymentModel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPaymentModel is a free data retrieval call binding the contract method 0xf1570141.
//
// Solidity: function getPaymentModel() view returns(uint8)
func (_RegistrarContract *RegistrarContractSession) GetPaymentModel() (uint8, error) {
	return _RegistrarContract.Contract.GetPaymentModel(&_RegistrarContract.CallOpts)
}

// GetPaymentModel is a free data retrieval call binding the contract method 0xf1570141.
//
// Solidity: function getPaymentModel() view returns(uint8)
func (_RegistrarContract *RegistrarContractCallerSession) GetPaymentModel() (uint8, error) {
	return _RegistrarContract.Contract.GetPaymentModel(&_RegistrarContract.CallOpts)
}

// GetPeerRegistryMigrationPermission is a free data retrieval call binding the contract method 0xfaa3e996.
//
// Solidity: function getPeerRegistryMigrationPermission(address peer) view returns(uint8)
func (_RegistrarContract *RegistrarContractCaller) GetPeerRegistryMigrationPermission(opts *bind.CallOpts, peer common.Address) (uint8, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getPeerRegistryMigrationPermission", peer)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPeerRegistryMigrationPermission is a free data retrieval call binding the contract method 0xfaa3e996.
//
// Solidity: function getPeerRegistryMigrationPermission(address peer) view returns(uint8)
func (_RegistrarContract *RegistrarContractSession) GetPeerRegistryMigrationPermission(peer common.Address) (uint8, error) {
	return _RegistrarContract.Contract.GetPeerRegistryMigrationPermission(&_RegistrarContract.CallOpts, peer)
}

// GetPeerRegistryMigrationPermission is a free data retrieval call binding the contract method 0xfaa3e996.
//
// Solidity: function getPeerRegistryMigrationPermission(address peer) view returns(uint8)
func (_RegistrarContract *RegistrarContractCallerSession) GetPeerRegistryMigrationPermission(peer common.Address) (uint8, error) {
	return _RegistrarContract.Contract.GetPeerRegistryMigrationPermission(&_RegistrarContract.CallOpts, peer)
}

// GetSignerInfo is a free data retrieval call binding the contract method 0xed56b3e1.
//
// Solidity: function getSignerInfo(address query) view returns(bool active, uint8 index)
func (_RegistrarContract *RegistrarContractCaller) GetSignerInfo(opts *bind.CallOpts, query common.Address) (struct {
	Active bool
	Index  uint8
}, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getSignerInfo", query)

	outstruct := new(struct {
		Active bool
		Index  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Index = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetSignerInfo is a free data retrieval call binding the contract method 0xed56b3e1.
//
// Solidity: function getSignerInfo(address query) view returns(bool active, uint8 index)
func (_RegistrarContract *RegistrarContractSession) GetSignerInfo(query common.Address) (struct {
	Active bool
	Index  uint8
}, error) {
	return _RegistrarContract.Contract.GetSignerInfo(&_RegistrarContract.CallOpts, query)
}

// GetSignerInfo is a free data retrieval call binding the contract method 0xed56b3e1.
//
// Solidity: function getSignerInfo(address query) view returns(bool active, uint8 index)
func (_RegistrarContract *RegistrarContractCallerSession) GetSignerInfo(query common.Address) (struct {
	Active bool
	Index  uint8
}, error) {
	return _RegistrarContract.Contract.GetSignerInfo(&_RegistrarContract.CallOpts, query)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((uint32,uint96,uint256,uint96,uint256,uint32,uint32,bytes32,uint32,bool) state, (uint32,uint32,uint32,uint24,uint16,uint96,uint32,uint32,uint32,uint256,uint256,address,address) config, address[] signers, address[] transmitters, uint8 f)
func (_RegistrarContract *RegistrarContractCaller) GetState(opts *bind.CallOpts) (struct {
	State        State
	Config       OnchainConfig
	Signers      []common.Address
	Transmitters []common.Address
	F            uint8
}, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getState")

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
func (_RegistrarContract *RegistrarContractSession) GetState() (struct {
	State        State
	Config       OnchainConfig
	Signers      []common.Address
	Transmitters []common.Address
	F            uint8
}, error) {
	return _RegistrarContract.Contract.GetState(&_RegistrarContract.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((uint32,uint96,uint256,uint96,uint256,uint32,uint32,bytes32,uint32,bool) state, (uint32,uint32,uint32,uint24,uint16,uint96,uint32,uint32,uint32,uint256,uint256,address,address) config, address[] signers, address[] transmitters, uint8 f)
func (_RegistrarContract *RegistrarContractCallerSession) GetState() (struct {
	State        State
	Config       OnchainConfig
	Signers      []common.Address
	Transmitters []common.Address
	F            uint8
}, error) {
	return _RegistrarContract.Contract.GetState(&_RegistrarContract.CallOpts)
}

// GetTransmitterInfo is a free data retrieval call binding the contract method 0x421d183b.
//
// Solidity: function getTransmitterInfo(address query) view returns(bool active, uint8 index, uint96 balance, uint96 lastCollected, address payee)
func (_RegistrarContract *RegistrarContractCaller) GetTransmitterInfo(opts *bind.CallOpts, query common.Address) (struct {
	Active        bool
	Index         uint8
	Balance       *big.Int
	LastCollected *big.Int
	Payee         common.Address
}, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getTransmitterInfo", query)

	outstruct := new(struct {
		Active        bool
		Index         uint8
		Balance       *big.Int
		LastCollected *big.Int
		Payee         common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Index = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Balance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastCollected = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Payee = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetTransmitterInfo is a free data retrieval call binding the contract method 0x421d183b.
//
// Solidity: function getTransmitterInfo(address query) view returns(bool active, uint8 index, uint96 balance, uint96 lastCollected, address payee)
func (_RegistrarContract *RegistrarContractSession) GetTransmitterInfo(query common.Address) (struct {
	Active        bool
	Index         uint8
	Balance       *big.Int
	LastCollected *big.Int
	Payee         common.Address
}, error) {
	return _RegistrarContract.Contract.GetTransmitterInfo(&_RegistrarContract.CallOpts, query)
}

// GetTransmitterInfo is a free data retrieval call binding the contract method 0x421d183b.
//
// Solidity: function getTransmitterInfo(address query) view returns(bool active, uint8 index, uint96 balance, uint96 lastCollected, address payee)
func (_RegistrarContract *RegistrarContractCallerSession) GetTransmitterInfo(query common.Address) (struct {
	Active        bool
	Index         uint8
	Balance       *big.Int
	LastCollected *big.Int
	Payee         common.Address
}, error) {
	return _RegistrarContract.Contract.GetTransmitterInfo(&_RegistrarContract.CallOpts, query)
}

// GetUpkeep is a free data retrieval call binding the contract method 0xc7c3a19a.
//
// Solidity: function getUpkeep(uint256 id) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes) upkeepInfo)
func (_RegistrarContract *RegistrarContractCaller) GetUpkeep(opts *bind.CallOpts, id *big.Int) (UpkeepInfo, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "getUpkeep", id)

	if err != nil {
		return *new(UpkeepInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UpkeepInfo)).(*UpkeepInfo)

	return out0, err

}

// GetUpkeep is a free data retrieval call binding the contract method 0xc7c3a19a.
//
// Solidity: function getUpkeep(uint256 id) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes) upkeepInfo)
func (_RegistrarContract *RegistrarContractSession) GetUpkeep(id *big.Int) (UpkeepInfo, error) {
	return _RegistrarContract.Contract.GetUpkeep(&_RegistrarContract.CallOpts, id)
}

// GetUpkeep is a free data retrieval call binding the contract method 0xc7c3a19a.
//
// Solidity: function getUpkeep(uint256 id) view returns((address,uint32,bytes,uint96,address,uint64,uint32,uint96,bool,bytes) upkeepInfo)
func (_RegistrarContract *RegistrarContractCallerSession) GetUpkeep(id *big.Int) (UpkeepInfo, error) {
	return _RegistrarContract.Contract.GetUpkeep(&_RegistrarContract.CallOpts, id)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_RegistrarContract *RegistrarContractCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(struct {
		ConfigCount  uint32
		BlockNumber  uint32
		ConfigDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_RegistrarContract *RegistrarContractSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _RegistrarContract.Contract.LatestConfigDetails(&_RegistrarContract.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_RegistrarContract *RegistrarContractCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _RegistrarContract.Contract.LatestConfigDetails(&_RegistrarContract.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_RegistrarContract *RegistrarContractCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(struct {
		ScanLogs     bool
		ConfigDigest [32]byte
		Epoch        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_RegistrarContract *RegistrarContractSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _RegistrarContract.Contract.LatestConfigDigestAndEpoch(&_RegistrarContract.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_RegistrarContract *RegistrarContractCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _RegistrarContract.Contract.LatestConfigDigestAndEpoch(&_RegistrarContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrarContract *RegistrarContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrarContract *RegistrarContractSession) Owner() (common.Address, error) {
	return _RegistrarContract.Contract.Owner(&_RegistrarContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrarContract *RegistrarContractCallerSession) Owner() (common.Address, error) {
	return _RegistrarContract.Contract.Owner(&_RegistrarContract.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_RegistrarContract *RegistrarContractCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_RegistrarContract *RegistrarContractSession) TypeAndVersion() (string, error) {
	return _RegistrarContract.Contract.TypeAndVersion(&_RegistrarContract.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_RegistrarContract *RegistrarContractCallerSession) TypeAndVersion() (string, error) {
	return _RegistrarContract.Contract.TypeAndVersion(&_RegistrarContract.CallOpts)
}

// UpkeepTranscoderVersion is a free data retrieval call binding the contract method 0x48013d7b.
//
// Solidity: function upkeepTranscoderVersion() view returns(uint8)
func (_RegistrarContract *RegistrarContractCaller) UpkeepTranscoderVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RegistrarContract.contract.Call(opts, &out, "upkeepTranscoderVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UpkeepTranscoderVersion is a free data retrieval call binding the contract method 0x48013d7b.
//
// Solidity: function upkeepTranscoderVersion() view returns(uint8)
func (_RegistrarContract *RegistrarContractSession) UpkeepTranscoderVersion() (uint8, error) {
	return _RegistrarContract.Contract.UpkeepTranscoderVersion(&_RegistrarContract.CallOpts)
}

// UpkeepTranscoderVersion is a free data retrieval call binding the contract method 0x48013d7b.
//
// Solidity: function upkeepTranscoderVersion() view returns(uint8)
func (_RegistrarContract *RegistrarContractCallerSession) UpkeepTranscoderVersion() (uint8, error) {
	return _RegistrarContract.Contract.UpkeepTranscoderVersion(&_RegistrarContract.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RegistrarContract *RegistrarContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RegistrarContract *RegistrarContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptOwnership(&_RegistrarContract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RegistrarContract *RegistrarContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptOwnership(&_RegistrarContract.TransactOpts)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address transmitter) returns()
func (_RegistrarContract *RegistrarContractTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "acceptPayeeship", transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address transmitter) returns()
func (_RegistrarContract *RegistrarContractSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptPayeeship(&_RegistrarContract.TransactOpts, transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address transmitter) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptPayeeship(&_RegistrarContract.TransactOpts, transmitter)
}

// AcceptUpkeepAdmin is a paid mutator transaction binding the contract method 0xb148ab6b.
//
// Solidity: function acceptUpkeepAdmin(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactor) AcceptUpkeepAdmin(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "acceptUpkeepAdmin", id)
}

// AcceptUpkeepAdmin is a paid mutator transaction binding the contract method 0xb148ab6b.
//
// Solidity: function acceptUpkeepAdmin(uint256 id) returns()
func (_RegistrarContract *RegistrarContractSession) AcceptUpkeepAdmin(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptUpkeepAdmin(&_RegistrarContract.TransactOpts, id)
}

// AcceptUpkeepAdmin is a paid mutator transaction binding the contract method 0xb148ab6b.
//
// Solidity: function acceptUpkeepAdmin(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) AcceptUpkeepAdmin(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptUpkeepAdmin(&_RegistrarContract.TransactOpts, id)
}

// AddFunds is a paid mutator transaction binding the contract method 0x948108f7.
//
// Solidity: function addFunds(uint256 id, uint96 amount) returns()
func (_RegistrarContract *RegistrarContractTransactor) AddFunds(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "addFunds", id, amount)
}

// AddFunds is a paid mutator transaction binding the contract method 0x948108f7.
//
// Solidity: function addFunds(uint256 id, uint96 amount) returns()
func (_RegistrarContract *RegistrarContractSession) AddFunds(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AddFunds(&_RegistrarContract.TransactOpts, id, amount)
}

// AddFunds is a paid mutator transaction binding the contract method 0x948108f7.
//
// Solidity: function addFunds(uint256 id, uint96 amount) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) AddFunds(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AddFunds(&_RegistrarContract.TransactOpts, id, amount)
}

// CancelUpkeep is a paid mutator transaction binding the contract method 0xc8048022.
//
// Solidity: function cancelUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactor) CancelUpkeep(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "cancelUpkeep", id)
}

// CancelUpkeep is a paid mutator transaction binding the contract method 0xc8048022.
//
// Solidity: function cancelUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractSession) CancelUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.CancelUpkeep(&_RegistrarContract.TransactOpts, id)
}

// CancelUpkeep is a paid mutator transaction binding the contract method 0xc8048022.
//
// Solidity: function cancelUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) CancelUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.CancelUpkeep(&_RegistrarContract.TransactOpts, id)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0xf7d334ba.
//
// Solidity: function checkUpkeep(uint256 id) returns(bool upkeepNeeded, bytes performData, uint8 upkeepFailureReason, uint256 gasUsed, uint256 fastGasWei, uint256 linkNative)
func (_RegistrarContract *RegistrarContractTransactor) CheckUpkeep(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "checkUpkeep", id)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0xf7d334ba.
//
// Solidity: function checkUpkeep(uint256 id) returns(bool upkeepNeeded, bytes performData, uint8 upkeepFailureReason, uint256 gasUsed, uint256 fastGasWei, uint256 linkNative)
func (_RegistrarContract *RegistrarContractSession) CheckUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.CheckUpkeep(&_RegistrarContract.TransactOpts, id)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0xf7d334ba.
//
// Solidity: function checkUpkeep(uint256 id) returns(bool upkeepNeeded, bytes performData, uint8 upkeepFailureReason, uint256 gasUsed, uint256 fastGasWei, uint256 linkNative)
func (_RegistrarContract *RegistrarContractTransactorSession) CheckUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.CheckUpkeep(&_RegistrarContract.TransactOpts, id)
}

// MigrateUpkeeps is a paid mutator transaction binding the contract method 0x85c1b0ba.
//
// Solidity: function migrateUpkeeps(uint256[] ids, address destination) returns()
func (_RegistrarContract *RegistrarContractTransactor) MigrateUpkeeps(opts *bind.TransactOpts, ids []*big.Int, destination common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "migrateUpkeeps", ids, destination)
}

// MigrateUpkeeps is a paid mutator transaction binding the contract method 0x85c1b0ba.
//
// Solidity: function migrateUpkeeps(uint256[] ids, address destination) returns()
func (_RegistrarContract *RegistrarContractSession) MigrateUpkeeps(ids []*big.Int, destination common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.MigrateUpkeeps(&_RegistrarContract.TransactOpts, ids, destination)
}

// MigrateUpkeeps is a paid mutator transaction binding the contract method 0x85c1b0ba.
//
// Solidity: function migrateUpkeeps(uint256[] ids, address destination) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) MigrateUpkeeps(ids []*big.Int, destination common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.MigrateUpkeeps(&_RegistrarContract.TransactOpts, ids, destination)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_RegistrarContract *RegistrarContractTransactor) OnTokenTransfer(opts *bind.TransactOpts, sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "onTokenTransfer", sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_RegistrarContract *RegistrarContractSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.OnTokenTransfer(&_RegistrarContract.TransactOpts, sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.OnTokenTransfer(&_RegistrarContract.TransactOpts, sender, amount, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RegistrarContract *RegistrarContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RegistrarContract *RegistrarContractSession) Pause() (*types.Transaction, error) {
	return _RegistrarContract.Contract.Pause(&_RegistrarContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RegistrarContract *RegistrarContractTransactorSession) Pause() (*types.Transaction, error) {
	return _RegistrarContract.Contract.Pause(&_RegistrarContract.TransactOpts)
}

// PauseUpkeep is a paid mutator transaction binding the contract method 0x8765ecbe.
//
// Solidity: function pauseUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactor) PauseUpkeep(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "pauseUpkeep", id)
}

// PauseUpkeep is a paid mutator transaction binding the contract method 0x8765ecbe.
//
// Solidity: function pauseUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractSession) PauseUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.PauseUpkeep(&_RegistrarContract.TransactOpts, id)
}

// PauseUpkeep is a paid mutator transaction binding the contract method 0x8765ecbe.
//
// Solidity: function pauseUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) PauseUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.PauseUpkeep(&_RegistrarContract.TransactOpts, id)
}

// ReceiveUpkeeps is a paid mutator transaction binding the contract method 0x8e86139b.
//
// Solidity: function receiveUpkeeps(bytes encodedUpkeeps) returns()
func (_RegistrarContract *RegistrarContractTransactor) ReceiveUpkeeps(opts *bind.TransactOpts, encodedUpkeeps []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "receiveUpkeeps", encodedUpkeeps)
}

// ReceiveUpkeeps is a paid mutator transaction binding the contract method 0x8e86139b.
//
// Solidity: function receiveUpkeeps(bytes encodedUpkeeps) returns()
func (_RegistrarContract *RegistrarContractSession) ReceiveUpkeeps(encodedUpkeeps []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.ReceiveUpkeeps(&_RegistrarContract.TransactOpts, encodedUpkeeps)
}

// ReceiveUpkeeps is a paid mutator transaction binding the contract method 0x8e86139b.
//
// Solidity: function receiveUpkeeps(bytes encodedUpkeeps) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) ReceiveUpkeeps(encodedUpkeeps []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.ReceiveUpkeeps(&_RegistrarContract.TransactOpts, encodedUpkeeps)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_RegistrarContract *RegistrarContractTransactor) RecoverFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "recoverFunds")
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_RegistrarContract *RegistrarContractSession) RecoverFunds() (*types.Transaction, error) {
	return _RegistrarContract.Contract.RecoverFunds(&_RegistrarContract.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_RegistrarContract *RegistrarContractTransactorSession) RecoverFunds() (*types.Transaction, error) {
	return _RegistrarContract.Contract.RecoverFunds(&_RegistrarContract.TransactOpts)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x6ded9eae.
//
// Solidity: function registerUpkeep(address target, uint32 gasLimit, address admin, bytes checkData, bytes offchainConfig) returns(uint256 id)
func (_RegistrarContract *RegistrarContractTransactor) RegisterUpkeep(opts *bind.TransactOpts, target common.Address, gasLimit uint32, admin common.Address, checkData []byte, offchainConfig []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "registerUpkeep", target, gasLimit, admin, checkData, offchainConfig)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x6ded9eae.
//
// Solidity: function registerUpkeep(address target, uint32 gasLimit, address admin, bytes checkData, bytes offchainConfig) returns(uint256 id)
func (_RegistrarContract *RegistrarContractSession) RegisterUpkeep(target common.Address, gasLimit uint32, admin common.Address, checkData []byte, offchainConfig []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.RegisterUpkeep(&_RegistrarContract.TransactOpts, target, gasLimit, admin, checkData, offchainConfig)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x6ded9eae.
//
// Solidity: function registerUpkeep(address target, uint32 gasLimit, address admin, bytes checkData, bytes offchainConfig) returns(uint256 id)
func (_RegistrarContract *RegistrarContractTransactorSession) RegisterUpkeep(target common.Address, gasLimit uint32, admin common.Address, checkData []byte, offchainConfig []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.RegisterUpkeep(&_RegistrarContract.TransactOpts, target, gasLimit, admin, checkData, offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig) returns()
func (_RegistrarContract *RegistrarContractTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig) returns()
func (_RegistrarContract *RegistrarContractSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetConfig(&_RegistrarContract.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetConfig(&_RegistrarContract.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

// SetPayees is a paid mutator transaction binding the contract method 0x3b9cce59.
//
// Solidity: function setPayees(address[] payees) returns()
func (_RegistrarContract *RegistrarContractTransactor) SetPayees(opts *bind.TransactOpts, payees []common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "setPayees", payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x3b9cce59.
//
// Solidity: function setPayees(address[] payees) returns()
func (_RegistrarContract *RegistrarContractSession) SetPayees(payees []common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetPayees(&_RegistrarContract.TransactOpts, payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x3b9cce59.
//
// Solidity: function setPayees(address[] payees) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) SetPayees(payees []common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetPayees(&_RegistrarContract.TransactOpts, payees)
}

// SetPeerRegistryMigrationPermission is a paid mutator transaction binding the contract method 0x187256e8.
//
// Solidity: function setPeerRegistryMigrationPermission(address peer, uint8 permission) returns()
func (_RegistrarContract *RegistrarContractTransactor) SetPeerRegistryMigrationPermission(opts *bind.TransactOpts, peer common.Address, permission uint8) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "setPeerRegistryMigrationPermission", peer, permission)
}

// SetPeerRegistryMigrationPermission is a paid mutator transaction binding the contract method 0x187256e8.
//
// Solidity: function setPeerRegistryMigrationPermission(address peer, uint8 permission) returns()
func (_RegistrarContract *RegistrarContractSession) SetPeerRegistryMigrationPermission(peer common.Address, permission uint8) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetPeerRegistryMigrationPermission(&_RegistrarContract.TransactOpts, peer, permission)
}

// SetPeerRegistryMigrationPermission is a paid mutator transaction binding the contract method 0x187256e8.
//
// Solidity: function setPeerRegistryMigrationPermission(address peer, uint8 permission) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) SetPeerRegistryMigrationPermission(peer common.Address, permission uint8) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetPeerRegistryMigrationPermission(&_RegistrarContract.TransactOpts, peer, permission)
}

// SetUpkeepGasLimit is a paid mutator transaction binding the contract method 0xa72aa27e.
//
// Solidity: function setUpkeepGasLimit(uint256 id, uint32 gasLimit) returns()
func (_RegistrarContract *RegistrarContractTransactor) SetUpkeepGasLimit(opts *bind.TransactOpts, id *big.Int, gasLimit uint32) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "setUpkeepGasLimit", id, gasLimit)
}

// SetUpkeepGasLimit is a paid mutator transaction binding the contract method 0xa72aa27e.
//
// Solidity: function setUpkeepGasLimit(uint256 id, uint32 gasLimit) returns()
func (_RegistrarContract *RegistrarContractSession) SetUpkeepGasLimit(id *big.Int, gasLimit uint32) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetUpkeepGasLimit(&_RegistrarContract.TransactOpts, id, gasLimit)
}

// SetUpkeepGasLimit is a paid mutator transaction binding the contract method 0xa72aa27e.
//
// Solidity: function setUpkeepGasLimit(uint256 id, uint32 gasLimit) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) SetUpkeepGasLimit(id *big.Int, gasLimit uint32) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetUpkeepGasLimit(&_RegistrarContract.TransactOpts, id, gasLimit)
}

// SetUpkeepOffchainConfig is a paid mutator transaction binding the contract method 0x8dcf0fe7.
//
// Solidity: function setUpkeepOffchainConfig(uint256 id, bytes config) returns()
func (_RegistrarContract *RegistrarContractTransactor) SetUpkeepOffchainConfig(opts *bind.TransactOpts, id *big.Int, config []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "setUpkeepOffchainConfig", id, config)
}

// SetUpkeepOffchainConfig is a paid mutator transaction binding the contract method 0x8dcf0fe7.
//
// Solidity: function setUpkeepOffchainConfig(uint256 id, bytes config) returns()
func (_RegistrarContract *RegistrarContractSession) SetUpkeepOffchainConfig(id *big.Int, config []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetUpkeepOffchainConfig(&_RegistrarContract.TransactOpts, id, config)
}

// SetUpkeepOffchainConfig is a paid mutator transaction binding the contract method 0x8dcf0fe7.
//
// Solidity: function setUpkeepOffchainConfig(uint256 id, bytes config) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) SetUpkeepOffchainConfig(id *big.Int, config []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SetUpkeepOffchainConfig(&_RegistrarContract.TransactOpts, id, config)
}

// SimulatePerformUpkeep is a paid mutator transaction binding the contract method 0xaed2e929.
//
// Solidity: function simulatePerformUpkeep(uint256 id, bytes performData) returns(bool success, uint256 gasUsed)
func (_RegistrarContract *RegistrarContractTransactor) SimulatePerformUpkeep(opts *bind.TransactOpts, id *big.Int, performData []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "simulatePerformUpkeep", id, performData)
}

// SimulatePerformUpkeep is a paid mutator transaction binding the contract method 0xaed2e929.
//
// Solidity: function simulatePerformUpkeep(uint256 id, bytes performData) returns(bool success, uint256 gasUsed)
func (_RegistrarContract *RegistrarContractSession) SimulatePerformUpkeep(id *big.Int, performData []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SimulatePerformUpkeep(&_RegistrarContract.TransactOpts, id, performData)
}

// SimulatePerformUpkeep is a paid mutator transaction binding the contract method 0xaed2e929.
//
// Solidity: function simulatePerformUpkeep(uint256 id, bytes performData) returns(bool success, uint256 gasUsed)
func (_RegistrarContract *RegistrarContractTransactorSession) SimulatePerformUpkeep(id *big.Int, performData []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.SimulatePerformUpkeep(&_RegistrarContract.TransactOpts, id, performData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_RegistrarContract *RegistrarContractTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_RegistrarContract *RegistrarContractSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferOwnership(&_RegistrarContract.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferOwnership(&_RegistrarContract.TransactOpts, to)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address transmitter, address proposed) returns()
func (_RegistrarContract *RegistrarContractTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address transmitter, address proposed) returns()
func (_RegistrarContract *RegistrarContractSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferPayeeship(&_RegistrarContract.TransactOpts, transmitter, proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address transmitter, address proposed) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferPayeeship(&_RegistrarContract.TransactOpts, transmitter, proposed)
}

// TransferUpkeepAdmin is a paid mutator transaction binding the contract method 0x1a2af011.
//
// Solidity: function transferUpkeepAdmin(uint256 id, address proposed) returns()
func (_RegistrarContract *RegistrarContractTransactor) TransferUpkeepAdmin(opts *bind.TransactOpts, id *big.Int, proposed common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "transferUpkeepAdmin", id, proposed)
}

// TransferUpkeepAdmin is a paid mutator transaction binding the contract method 0x1a2af011.
//
// Solidity: function transferUpkeepAdmin(uint256 id, address proposed) returns()
func (_RegistrarContract *RegistrarContractSession) TransferUpkeepAdmin(id *big.Int, proposed common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferUpkeepAdmin(&_RegistrarContract.TransactOpts, id, proposed)
}

// TransferUpkeepAdmin is a paid mutator transaction binding the contract method 0x1a2af011.
//
// Solidity: function transferUpkeepAdmin(uint256 id, address proposed) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) TransferUpkeepAdmin(id *big.Int, proposed common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferUpkeepAdmin(&_RegistrarContract.TransactOpts, id, proposed)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes rawReport, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_RegistrarContract *RegistrarContractTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, rawReport []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "transmit", reportContext, rawReport, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes rawReport, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_RegistrarContract *RegistrarContractSession) Transmit(reportContext [3][32]byte, rawReport []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.Transmit(&_RegistrarContract.TransactOpts, reportContext, rawReport, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes rawReport, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) Transmit(reportContext [3][32]byte, rawReport []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.Transmit(&_RegistrarContract.TransactOpts, reportContext, rawReport, rs, ss, rawVs)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RegistrarContract *RegistrarContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RegistrarContract *RegistrarContractSession) Unpause() (*types.Transaction, error) {
	return _RegistrarContract.Contract.Unpause(&_RegistrarContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RegistrarContract *RegistrarContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _RegistrarContract.Contract.Unpause(&_RegistrarContract.TransactOpts)
}

// UnpauseUpkeep is a paid mutator transaction binding the contract method 0x5165f2f5.
//
// Solidity: function unpauseUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactor) UnpauseUpkeep(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "unpauseUpkeep", id)
}

// UnpauseUpkeep is a paid mutator transaction binding the contract method 0x5165f2f5.
//
// Solidity: function unpauseUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractSession) UnpauseUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.UnpauseUpkeep(&_RegistrarContract.TransactOpts, id)
}

// UnpauseUpkeep is a paid mutator transaction binding the contract method 0x5165f2f5.
//
// Solidity: function unpauseUpkeep(uint256 id) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) UnpauseUpkeep(id *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.UnpauseUpkeep(&_RegistrarContract.TransactOpts, id)
}

// UpdateCheckData is a paid mutator transaction binding the contract method 0x9fab4386.
//
// Solidity: function updateCheckData(uint256 id, bytes newCheckData) returns()
func (_RegistrarContract *RegistrarContractTransactor) UpdateCheckData(opts *bind.TransactOpts, id *big.Int, newCheckData []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "updateCheckData", id, newCheckData)
}

// UpdateCheckData is a paid mutator transaction binding the contract method 0x9fab4386.
//
// Solidity: function updateCheckData(uint256 id, bytes newCheckData) returns()
func (_RegistrarContract *RegistrarContractSession) UpdateCheckData(id *big.Int, newCheckData []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.UpdateCheckData(&_RegistrarContract.TransactOpts, id, newCheckData)
}

// UpdateCheckData is a paid mutator transaction binding the contract method 0x9fab4386.
//
// Solidity: function updateCheckData(uint256 id, bytes newCheckData) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) UpdateCheckData(id *big.Int, newCheckData []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.UpdateCheckData(&_RegistrarContract.TransactOpts, id, newCheckData)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x744bfe61.
//
// Solidity: function withdrawFunds(uint256 id, address to) returns()
func (_RegistrarContract *RegistrarContractTransactor) WithdrawFunds(opts *bind.TransactOpts, id *big.Int, to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "withdrawFunds", id, to)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x744bfe61.
//
// Solidity: function withdrawFunds(uint256 id, address to) returns()
func (_RegistrarContract *RegistrarContractSession) WithdrawFunds(id *big.Int, to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.WithdrawFunds(&_RegistrarContract.TransactOpts, id, to)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x744bfe61.
//
// Solidity: function withdrawFunds(uint256 id, address to) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) WithdrawFunds(id *big.Int, to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.WithdrawFunds(&_RegistrarContract.TransactOpts, id, to)
}

// WithdrawOwnerFunds is a paid mutator transaction binding the contract method 0x7d9b97e0.
//
// Solidity: function withdrawOwnerFunds() returns()
func (_RegistrarContract *RegistrarContractTransactor) WithdrawOwnerFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "withdrawOwnerFunds")
}

// WithdrawOwnerFunds is a paid mutator transaction binding the contract method 0x7d9b97e0.
//
// Solidity: function withdrawOwnerFunds() returns()
func (_RegistrarContract *RegistrarContractSession) WithdrawOwnerFunds() (*types.Transaction, error) {
	return _RegistrarContract.Contract.WithdrawOwnerFunds(&_RegistrarContract.TransactOpts)
}

// WithdrawOwnerFunds is a paid mutator transaction binding the contract method 0x7d9b97e0.
//
// Solidity: function withdrawOwnerFunds() returns()
func (_RegistrarContract *RegistrarContractTransactorSession) WithdrawOwnerFunds() (*types.Transaction, error) {
	return _RegistrarContract.Contract.WithdrawOwnerFunds(&_RegistrarContract.TransactOpts)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xa710b221.
//
// Solidity: function withdrawPayment(address from, address to) returns()
func (_RegistrarContract *RegistrarContractTransactor) WithdrawPayment(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "withdrawPayment", from, to)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xa710b221.
//
// Solidity: function withdrawPayment(address from, address to) returns()
func (_RegistrarContract *RegistrarContractSession) WithdrawPayment(from common.Address, to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.WithdrawPayment(&_RegistrarContract.TransactOpts, from, to)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xa710b221.
//
// Solidity: function withdrawPayment(address from, address to) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) WithdrawPayment(from common.Address, to common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.WithdrawPayment(&_RegistrarContract.TransactOpts, from, to)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RegistrarContract *RegistrarContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RegistrarContract *RegistrarContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.Fallback(&_RegistrarContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RegistrarContract *RegistrarContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.Fallback(&_RegistrarContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RegistrarContract *RegistrarContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RegistrarContract *RegistrarContractSession) Receive() (*types.Transaction, error) {
	return _RegistrarContract.Contract.Receive(&_RegistrarContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RegistrarContract *RegistrarContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RegistrarContract.Contract.Receive(&_RegistrarContract.TransactOpts)
}

// RegistrarContractCancelledUpkeepReportIterator is returned from FilterCancelledUpkeepReport and is used to iterate over the raw logs and unpacked data for CancelledUpkeepReport events raised by the RegistrarContract contract.
type RegistrarContractCancelledUpkeepReportIterator struct {
	Event *RegistrarContractCancelledUpkeepReport // Event containing the contract specifics and raw log

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
func (it *RegistrarContractCancelledUpkeepReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractCancelledUpkeepReport)
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
		it.Event = new(RegistrarContractCancelledUpkeepReport)
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
func (it *RegistrarContractCancelledUpkeepReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractCancelledUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractCancelledUpkeepReport represents a CancelledUpkeepReport event raised by the RegistrarContract contract.
type RegistrarContractCancelledUpkeepReport struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelledUpkeepReport is a free log retrieval operation binding the contract event 0xd84831b6a3a7fbd333f42fe7f9104a139da6cca4cc1507aef4ddad79b31d017f.
//
// Solidity: event CancelledUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) FilterCancelledUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractCancelledUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "CancelledUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractCancelledUpkeepReportIterator{contract: _RegistrarContract.contract, event: "CancelledUpkeepReport", logs: logs, sub: sub}, nil
}

// WatchCancelledUpkeepReport is a free log subscription operation binding the contract event 0xd84831b6a3a7fbd333f42fe7f9104a139da6cca4cc1507aef4ddad79b31d017f.
//
// Solidity: event CancelledUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) WatchCancelledUpkeepReport(opts *bind.WatchOpts, sink chan<- *RegistrarContractCancelledUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "CancelledUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractCancelledUpkeepReport)
				if err := _RegistrarContract.contract.UnpackLog(event, "CancelledUpkeepReport", log); err != nil {
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

// ParseCancelledUpkeepReport is a log parse operation binding the contract event 0xd84831b6a3a7fbd333f42fe7f9104a139da6cca4cc1507aef4ddad79b31d017f.
//
// Solidity: event CancelledUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) ParseCancelledUpkeepReport(log types.Log) (*RegistrarContractCancelledUpkeepReport, error) {
	event := new(RegistrarContractCancelledUpkeepReport)
	if err := _RegistrarContract.contract.UnpackLog(event, "CancelledUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the RegistrarContract contract.
type RegistrarContractConfigSetIterator struct {
	Event *RegistrarContractConfigSet // Event containing the contract specifics and raw log

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
func (it *RegistrarContractConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractConfigSet)
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
		it.Event = new(RegistrarContractConfigSet)
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
func (it *RegistrarContractConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractConfigSet represents a ConfigSet event raised by the RegistrarContract contract.
type RegistrarContractConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_RegistrarContract *RegistrarContractFilterer) FilterConfigSet(opts *bind.FilterOpts) (*RegistrarContractConfigSetIterator, error) {

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &RegistrarContractConfigSetIterator{contract: _RegistrarContract.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_RegistrarContract *RegistrarContractFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RegistrarContractConfigSet) (event.Subscription, error) {

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractConfigSet)
				if err := _RegistrarContract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

// ParseConfigSet is a log parse operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_RegistrarContract *RegistrarContractFilterer) ParseConfigSet(log types.Log) (*RegistrarContractConfigSet, error) {
	event := new(RegistrarContractConfigSet)
	if err := _RegistrarContract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractFundsAddedIterator is returned from FilterFundsAdded and is used to iterate over the raw logs and unpacked data for FundsAdded events raised by the RegistrarContract contract.
type RegistrarContractFundsAddedIterator struct {
	Event *RegistrarContractFundsAdded // Event containing the contract specifics and raw log

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
func (it *RegistrarContractFundsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractFundsAdded)
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
		it.Event = new(RegistrarContractFundsAdded)
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
func (it *RegistrarContractFundsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractFundsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractFundsAdded represents a FundsAdded event raised by the RegistrarContract contract.
type RegistrarContractFundsAdded struct {
	Id     *big.Int
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsAdded is a free log retrieval operation binding the contract event 0xafd24114486da8ebfc32f3626dada8863652e187461aa74d4bfa734891506203.
//
// Solidity: event FundsAdded(uint256 indexed id, address indexed from, uint96 amount)
func (_RegistrarContract *RegistrarContractFilterer) FilterFundsAdded(opts *bind.FilterOpts, id []*big.Int, from []common.Address) (*RegistrarContractFundsAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "FundsAdded", idRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractFundsAddedIterator{contract: _RegistrarContract.contract, event: "FundsAdded", logs: logs, sub: sub}, nil
}

// WatchFundsAdded is a free log subscription operation binding the contract event 0xafd24114486da8ebfc32f3626dada8863652e187461aa74d4bfa734891506203.
//
// Solidity: event FundsAdded(uint256 indexed id, address indexed from, uint96 amount)
func (_RegistrarContract *RegistrarContractFilterer) WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *RegistrarContractFundsAdded, id []*big.Int, from []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "FundsAdded", idRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractFundsAdded)
				if err := _RegistrarContract.contract.UnpackLog(event, "FundsAdded", log); err != nil {
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

// ParseFundsAdded is a log parse operation binding the contract event 0xafd24114486da8ebfc32f3626dada8863652e187461aa74d4bfa734891506203.
//
// Solidity: event FundsAdded(uint256 indexed id, address indexed from, uint96 amount)
func (_RegistrarContract *RegistrarContractFilterer) ParseFundsAdded(log types.Log) (*RegistrarContractFundsAdded, error) {
	event := new(RegistrarContractFundsAdded)
	if err := _RegistrarContract.contract.UnpackLog(event, "FundsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the RegistrarContract contract.
type RegistrarContractFundsWithdrawnIterator struct {
	Event *RegistrarContractFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *RegistrarContractFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractFundsWithdrawn)
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
		it.Event = new(RegistrarContractFundsWithdrawn)
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
func (it *RegistrarContractFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractFundsWithdrawn represents a FundsWithdrawn event raised by the RegistrarContract contract.
type RegistrarContractFundsWithdrawn struct {
	Id     *big.Int
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xf3b5906e5672f3e524854103bcafbbdba80dbdfeca2c35e116127b1060a68318.
//
// Solidity: event FundsWithdrawn(uint256 indexed id, uint256 amount, address to)
func (_RegistrarContract *RegistrarContractFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractFundsWithdrawnIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "FundsWithdrawn", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractFundsWithdrawnIterator{contract: _RegistrarContract.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xf3b5906e5672f3e524854103bcafbbdba80dbdfeca2c35e116127b1060a68318.
//
// Solidity: event FundsWithdrawn(uint256 indexed id, uint256 amount, address to)
func (_RegistrarContract *RegistrarContractFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *RegistrarContractFundsWithdrawn, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "FundsWithdrawn", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractFundsWithdrawn)
				if err := _RegistrarContract.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xf3b5906e5672f3e524854103bcafbbdba80dbdfeca2c35e116127b1060a68318.
//
// Solidity: event FundsWithdrawn(uint256 indexed id, uint256 amount, address to)
func (_RegistrarContract *RegistrarContractFilterer) ParseFundsWithdrawn(log types.Log) (*RegistrarContractFundsWithdrawn, error) {
	event := new(RegistrarContractFundsWithdrawn)
	if err := _RegistrarContract.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractInsufficientFundsUpkeepReportIterator is returned from FilterInsufficientFundsUpkeepReport and is used to iterate over the raw logs and unpacked data for InsufficientFundsUpkeepReport events raised by the RegistrarContract contract.
type RegistrarContractInsufficientFundsUpkeepReportIterator struct {
	Event *RegistrarContractInsufficientFundsUpkeepReport // Event containing the contract specifics and raw log

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
func (it *RegistrarContractInsufficientFundsUpkeepReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractInsufficientFundsUpkeepReport)
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
		it.Event = new(RegistrarContractInsufficientFundsUpkeepReport)
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
func (it *RegistrarContractInsufficientFundsUpkeepReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractInsufficientFundsUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractInsufficientFundsUpkeepReport represents a InsufficientFundsUpkeepReport event raised by the RegistrarContract contract.
type RegistrarContractInsufficientFundsUpkeepReport struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInsufficientFundsUpkeepReport is a free log retrieval operation binding the contract event 0x7895fdfe292beab0842d5beccd078e85296b9e17a30eaee4c261a2696b84eb96.
//
// Solidity: event InsufficientFundsUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) FilterInsufficientFundsUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractInsufficientFundsUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "InsufficientFundsUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractInsufficientFundsUpkeepReportIterator{contract: _RegistrarContract.contract, event: "InsufficientFundsUpkeepReport", logs: logs, sub: sub}, nil
}

// WatchInsufficientFundsUpkeepReport is a free log subscription operation binding the contract event 0x7895fdfe292beab0842d5beccd078e85296b9e17a30eaee4c261a2696b84eb96.
//
// Solidity: event InsufficientFundsUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) WatchInsufficientFundsUpkeepReport(opts *bind.WatchOpts, sink chan<- *RegistrarContractInsufficientFundsUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "InsufficientFundsUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractInsufficientFundsUpkeepReport)
				if err := _RegistrarContract.contract.UnpackLog(event, "InsufficientFundsUpkeepReport", log); err != nil {
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

// ParseInsufficientFundsUpkeepReport is a log parse operation binding the contract event 0x7895fdfe292beab0842d5beccd078e85296b9e17a30eaee4c261a2696b84eb96.
//
// Solidity: event InsufficientFundsUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) ParseInsufficientFundsUpkeepReport(log types.Log) (*RegistrarContractInsufficientFundsUpkeepReport, error) {
	event := new(RegistrarContractInsufficientFundsUpkeepReport)
	if err := _RegistrarContract.contract.UnpackLog(event, "InsufficientFundsUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractOwnerFundsWithdrawnIterator is returned from FilterOwnerFundsWithdrawn and is used to iterate over the raw logs and unpacked data for OwnerFundsWithdrawn events raised by the RegistrarContract contract.
type RegistrarContractOwnerFundsWithdrawnIterator struct {
	Event *RegistrarContractOwnerFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *RegistrarContractOwnerFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractOwnerFundsWithdrawn)
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
		it.Event = new(RegistrarContractOwnerFundsWithdrawn)
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
func (it *RegistrarContractOwnerFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractOwnerFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractOwnerFundsWithdrawn represents a OwnerFundsWithdrawn event raised by the RegistrarContract contract.
type RegistrarContractOwnerFundsWithdrawn struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOwnerFundsWithdrawn is a free log retrieval operation binding the contract event 0x1d07d0b0be43d3e5fee41a80b579af370affee03fa595bf56d5d4c19328162f1.
//
// Solidity: event OwnerFundsWithdrawn(uint96 amount)
func (_RegistrarContract *RegistrarContractFilterer) FilterOwnerFundsWithdrawn(opts *bind.FilterOpts) (*RegistrarContractOwnerFundsWithdrawnIterator, error) {

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "OwnerFundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &RegistrarContractOwnerFundsWithdrawnIterator{contract: _RegistrarContract.contract, event: "OwnerFundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchOwnerFundsWithdrawn is a free log subscription operation binding the contract event 0x1d07d0b0be43d3e5fee41a80b579af370affee03fa595bf56d5d4c19328162f1.
//
// Solidity: event OwnerFundsWithdrawn(uint96 amount)
func (_RegistrarContract *RegistrarContractFilterer) WatchOwnerFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *RegistrarContractOwnerFundsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "OwnerFundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractOwnerFundsWithdrawn)
				if err := _RegistrarContract.contract.UnpackLog(event, "OwnerFundsWithdrawn", log); err != nil {
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

// ParseOwnerFundsWithdrawn is a log parse operation binding the contract event 0x1d07d0b0be43d3e5fee41a80b579af370affee03fa595bf56d5d4c19328162f1.
//
// Solidity: event OwnerFundsWithdrawn(uint96 amount)
func (_RegistrarContract *RegistrarContractFilterer) ParseOwnerFundsWithdrawn(log types.Log) (*RegistrarContractOwnerFundsWithdrawn, error) {
	event := new(RegistrarContractOwnerFundsWithdrawn)
	if err := _RegistrarContract.contract.UnpackLog(event, "OwnerFundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the RegistrarContract contract.
type RegistrarContractOwnershipTransferRequestedIterator struct {
	Event *RegistrarContractOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *RegistrarContractOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractOwnershipTransferRequested)
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
		it.Event = new(RegistrarContractOwnershipTransferRequested)
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
func (it *RegistrarContractOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the RegistrarContract contract.
type RegistrarContractOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RegistrarContractOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractOwnershipTransferRequestedIterator{contract: _RegistrarContract.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RegistrarContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractOwnershipTransferRequested)
				if err := _RegistrarContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_RegistrarContract *RegistrarContractFilterer) ParseOwnershipTransferRequested(log types.Log) (*RegistrarContractOwnershipTransferRequested, error) {
	event := new(RegistrarContractOwnershipTransferRequested)
	if err := _RegistrarContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegistrarContract contract.
type RegistrarContractOwnershipTransferredIterator struct {
	Event *RegistrarContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrarContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractOwnershipTransferred)
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
		it.Event = new(RegistrarContractOwnershipTransferred)
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
func (it *RegistrarContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractOwnershipTransferred represents a OwnershipTransferred event raised by the RegistrarContract contract.
type RegistrarContractOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RegistrarContractOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractOwnershipTransferredIterator{contract: _RegistrarContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrarContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractOwnershipTransferred)
				if err := _RegistrarContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RegistrarContract *RegistrarContractFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrarContractOwnershipTransferred, error) {
	event := new(RegistrarContractOwnershipTransferred)
	if err := _RegistrarContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RegistrarContract contract.
type RegistrarContractPausedIterator struct {
	Event *RegistrarContractPaused // Event containing the contract specifics and raw log

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
func (it *RegistrarContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractPaused)
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
		it.Event = new(RegistrarContractPaused)
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
func (it *RegistrarContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractPaused represents a Paused event raised by the RegistrarContract contract.
type RegistrarContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RegistrarContract *RegistrarContractFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistrarContractPausedIterator, error) {

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistrarContractPausedIterator{contract: _RegistrarContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RegistrarContract *RegistrarContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistrarContractPaused) (event.Subscription, error) {

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractPaused)
				if err := _RegistrarContract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RegistrarContract *RegistrarContractFilterer) ParsePaused(log types.Log) (*RegistrarContractPaused, error) {
	event := new(RegistrarContractPaused)
	if err := _RegistrarContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractPayeesUpdatedIterator is returned from FilterPayeesUpdated and is used to iterate over the raw logs and unpacked data for PayeesUpdated events raised by the RegistrarContract contract.
type RegistrarContractPayeesUpdatedIterator struct {
	Event *RegistrarContractPayeesUpdated // Event containing the contract specifics and raw log

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
func (it *RegistrarContractPayeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractPayeesUpdated)
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
		it.Event = new(RegistrarContractPayeesUpdated)
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
func (it *RegistrarContractPayeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractPayeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractPayeesUpdated represents a PayeesUpdated event raised by the RegistrarContract contract.
type RegistrarContractPayeesUpdated struct {
	Transmitters []common.Address
	Payees       []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPayeesUpdated is a free log retrieval operation binding the contract event 0xa46de38886467c59be07a0675f14781206a5477d871628af46c2443822fcb725.
//
// Solidity: event PayeesUpdated(address[] transmitters, address[] payees)
func (_RegistrarContract *RegistrarContractFilterer) FilterPayeesUpdated(opts *bind.FilterOpts) (*RegistrarContractPayeesUpdatedIterator, error) {

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "PayeesUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistrarContractPayeesUpdatedIterator{contract: _RegistrarContract.contract, event: "PayeesUpdated", logs: logs, sub: sub}, nil
}

// WatchPayeesUpdated is a free log subscription operation binding the contract event 0xa46de38886467c59be07a0675f14781206a5477d871628af46c2443822fcb725.
//
// Solidity: event PayeesUpdated(address[] transmitters, address[] payees)
func (_RegistrarContract *RegistrarContractFilterer) WatchPayeesUpdated(opts *bind.WatchOpts, sink chan<- *RegistrarContractPayeesUpdated) (event.Subscription, error) {

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "PayeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractPayeesUpdated)
				if err := _RegistrarContract.contract.UnpackLog(event, "PayeesUpdated", log); err != nil {
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

// ParsePayeesUpdated is a log parse operation binding the contract event 0xa46de38886467c59be07a0675f14781206a5477d871628af46c2443822fcb725.
//
// Solidity: event PayeesUpdated(address[] transmitters, address[] payees)
func (_RegistrarContract *RegistrarContractFilterer) ParsePayeesUpdated(log types.Log) (*RegistrarContractPayeesUpdated, error) {
	event := new(RegistrarContractPayeesUpdated)
	if err := _RegistrarContract.contract.UnpackLog(event, "PayeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractPayeeshipTransferRequestedIterator is returned from FilterPayeeshipTransferRequested and is used to iterate over the raw logs and unpacked data for PayeeshipTransferRequested events raised by the RegistrarContract contract.
type RegistrarContractPayeeshipTransferRequestedIterator struct {
	Event *RegistrarContractPayeeshipTransferRequested // Event containing the contract specifics and raw log

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
func (it *RegistrarContractPayeeshipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractPayeeshipTransferRequested)
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
		it.Event = new(RegistrarContractPayeeshipTransferRequested)
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
func (it *RegistrarContractPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractPayeeshipTransferRequested represents a PayeeshipTransferRequested event raised by the RegistrarContract contract.
type RegistrarContractPayeeshipTransferRequested struct {
	Transmitter common.Address
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferRequested is a free log retrieval operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, from []common.Address, to []common.Address) (*RegistrarContractPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractPayeeshipTransferRequestedIterator{contract: _RegistrarContract.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferRequested is a free log subscription operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *RegistrarContractPayeeshipTransferRequested, transmitter []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractPayeeshipTransferRequested)
				if err := _RegistrarContract.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

// ParsePayeeshipTransferRequested is a log parse operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) ParsePayeeshipTransferRequested(log types.Log) (*RegistrarContractPayeeshipTransferRequested, error) {
	event := new(RegistrarContractPayeeshipTransferRequested)
	if err := _RegistrarContract.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractPayeeshipTransferredIterator is returned from FilterPayeeshipTransferred and is used to iterate over the raw logs and unpacked data for PayeeshipTransferred events raised by the RegistrarContract contract.
type RegistrarContractPayeeshipTransferredIterator struct {
	Event *RegistrarContractPayeeshipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrarContractPayeeshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractPayeeshipTransferred)
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
		it.Event = new(RegistrarContractPayeeshipTransferred)
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
func (it *RegistrarContractPayeeshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractPayeeshipTransferred represents a PayeeshipTransferred event raised by the RegistrarContract contract.
type RegistrarContractPayeeshipTransferred struct {
	Transmitter common.Address
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferred is a free log retrieval operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, from []common.Address, to []common.Address) (*RegistrarContractPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractPayeeshipTransferredIterator{contract: _RegistrarContract.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferred is a free log subscription operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrarContractPayeeshipTransferred, transmitter []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractPayeeshipTransferred)
				if err := _RegistrarContract.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

// ParsePayeeshipTransferred is a log parse operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) ParsePayeeshipTransferred(log types.Log) (*RegistrarContractPayeeshipTransferred, error) {
	event := new(RegistrarContractPayeeshipTransferred)
	if err := _RegistrarContract.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractPaymentWithdrawnIterator is returned from FilterPaymentWithdrawn and is used to iterate over the raw logs and unpacked data for PaymentWithdrawn events raised by the RegistrarContract contract.
type RegistrarContractPaymentWithdrawnIterator struct {
	Event *RegistrarContractPaymentWithdrawn // Event containing the contract specifics and raw log

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
func (it *RegistrarContractPaymentWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractPaymentWithdrawn)
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
		it.Event = new(RegistrarContractPaymentWithdrawn)
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
func (it *RegistrarContractPaymentWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractPaymentWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractPaymentWithdrawn represents a PaymentWithdrawn event raised by the RegistrarContract contract.
type RegistrarContractPaymentWithdrawn struct {
	Transmitter common.Address
	Amount      *big.Int
	To          common.Address
	Payee       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaymentWithdrawn is a free log retrieval operation binding the contract event 0x9819093176a1851202c7bcfa46845809b4e47c261866550e94ed3775d2f40698.
//
// Solidity: event PaymentWithdrawn(address indexed transmitter, uint256 indexed amount, address indexed to, address payee)
func (_RegistrarContract *RegistrarContractFilterer) FilterPaymentWithdrawn(opts *bind.FilterOpts, transmitter []common.Address, amount []*big.Int, to []common.Address) (*RegistrarContractPaymentWithdrawnIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "PaymentWithdrawn", transmitterRule, amountRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractPaymentWithdrawnIterator{contract: _RegistrarContract.contract, event: "PaymentWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaymentWithdrawn is a free log subscription operation binding the contract event 0x9819093176a1851202c7bcfa46845809b4e47c261866550e94ed3775d2f40698.
//
// Solidity: event PaymentWithdrawn(address indexed transmitter, uint256 indexed amount, address indexed to, address payee)
func (_RegistrarContract *RegistrarContractFilterer) WatchPaymentWithdrawn(opts *bind.WatchOpts, sink chan<- *RegistrarContractPaymentWithdrawn, transmitter []common.Address, amount []*big.Int, to []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "PaymentWithdrawn", transmitterRule, amountRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractPaymentWithdrawn)
				if err := _RegistrarContract.contract.UnpackLog(event, "PaymentWithdrawn", log); err != nil {
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

// ParsePaymentWithdrawn is a log parse operation binding the contract event 0x9819093176a1851202c7bcfa46845809b4e47c261866550e94ed3775d2f40698.
//
// Solidity: event PaymentWithdrawn(address indexed transmitter, uint256 indexed amount, address indexed to, address payee)
func (_RegistrarContract *RegistrarContractFilterer) ParsePaymentWithdrawn(log types.Log) (*RegistrarContractPaymentWithdrawn, error) {
	event := new(RegistrarContractPaymentWithdrawn)
	if err := _RegistrarContract.contract.UnpackLog(event, "PaymentWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractReorgedUpkeepReportIterator is returned from FilterReorgedUpkeepReport and is used to iterate over the raw logs and unpacked data for ReorgedUpkeepReport events raised by the RegistrarContract contract.
type RegistrarContractReorgedUpkeepReportIterator struct {
	Event *RegistrarContractReorgedUpkeepReport // Event containing the contract specifics and raw log

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
func (it *RegistrarContractReorgedUpkeepReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractReorgedUpkeepReport)
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
		it.Event = new(RegistrarContractReorgedUpkeepReport)
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
func (it *RegistrarContractReorgedUpkeepReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractReorgedUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractReorgedUpkeepReport represents a ReorgedUpkeepReport event raised by the RegistrarContract contract.
type RegistrarContractReorgedUpkeepReport struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReorgedUpkeepReport is a free log retrieval operation binding the contract event 0x561ff77e59394941a01a456497a9418dea82e2a39abb3ecebfb1cef7e0bfdc13.
//
// Solidity: event ReorgedUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) FilterReorgedUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractReorgedUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "ReorgedUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractReorgedUpkeepReportIterator{contract: _RegistrarContract.contract, event: "ReorgedUpkeepReport", logs: logs, sub: sub}, nil
}

// WatchReorgedUpkeepReport is a free log subscription operation binding the contract event 0x561ff77e59394941a01a456497a9418dea82e2a39abb3ecebfb1cef7e0bfdc13.
//
// Solidity: event ReorgedUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) WatchReorgedUpkeepReport(opts *bind.WatchOpts, sink chan<- *RegistrarContractReorgedUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "ReorgedUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractReorgedUpkeepReport)
				if err := _RegistrarContract.contract.UnpackLog(event, "ReorgedUpkeepReport", log); err != nil {
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

// ParseReorgedUpkeepReport is a log parse operation binding the contract event 0x561ff77e59394941a01a456497a9418dea82e2a39abb3ecebfb1cef7e0bfdc13.
//
// Solidity: event ReorgedUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) ParseReorgedUpkeepReport(log types.Log) (*RegistrarContractReorgedUpkeepReport, error) {
	event := new(RegistrarContractReorgedUpkeepReport)
	if err := _RegistrarContract.contract.UnpackLog(event, "ReorgedUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractStaleUpkeepReportIterator is returned from FilterStaleUpkeepReport and is used to iterate over the raw logs and unpacked data for StaleUpkeepReport events raised by the RegistrarContract contract.
type RegistrarContractStaleUpkeepReportIterator struct {
	Event *RegistrarContractStaleUpkeepReport // Event containing the contract specifics and raw log

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
func (it *RegistrarContractStaleUpkeepReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractStaleUpkeepReport)
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
		it.Event = new(RegistrarContractStaleUpkeepReport)
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
func (it *RegistrarContractStaleUpkeepReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractStaleUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractStaleUpkeepReport represents a StaleUpkeepReport event raised by the RegistrarContract contract.
type RegistrarContractStaleUpkeepReport struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStaleUpkeepReport is a free log retrieval operation binding the contract event 0x5aa44821f7938098502bff537fbbdc9aaaa2fa655c10740646fce27e54987a89.
//
// Solidity: event StaleUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) FilterStaleUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractStaleUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "StaleUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractStaleUpkeepReportIterator{contract: _RegistrarContract.contract, event: "StaleUpkeepReport", logs: logs, sub: sub}, nil
}

// WatchStaleUpkeepReport is a free log subscription operation binding the contract event 0x5aa44821f7938098502bff537fbbdc9aaaa2fa655c10740646fce27e54987a89.
//
// Solidity: event StaleUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) WatchStaleUpkeepReport(opts *bind.WatchOpts, sink chan<- *RegistrarContractStaleUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "StaleUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractStaleUpkeepReport)
				if err := _RegistrarContract.contract.UnpackLog(event, "StaleUpkeepReport", log); err != nil {
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

// ParseStaleUpkeepReport is a log parse operation binding the contract event 0x5aa44821f7938098502bff537fbbdc9aaaa2fa655c10740646fce27e54987a89.
//
// Solidity: event StaleUpkeepReport(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) ParseStaleUpkeepReport(log types.Log) (*RegistrarContractStaleUpkeepReport, error) {
	event := new(RegistrarContractStaleUpkeepReport)
	if err := _RegistrarContract.contract.UnpackLog(event, "StaleUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractTransmittedIterator is returned from FilterTransmitted and is used to iterate over the raw logs and unpacked data for Transmitted events raised by the RegistrarContract contract.
type RegistrarContractTransmittedIterator struct {
	Event *RegistrarContractTransmitted // Event containing the contract specifics and raw log

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
func (it *RegistrarContractTransmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractTransmitted)
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
		it.Event = new(RegistrarContractTransmitted)
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
func (it *RegistrarContractTransmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractTransmitted represents a Transmitted event raised by the RegistrarContract contract.
type RegistrarContractTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransmitted is a free log retrieval operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_RegistrarContract *RegistrarContractFilterer) FilterTransmitted(opts *bind.FilterOpts) (*RegistrarContractTransmittedIterator, error) {

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &RegistrarContractTransmittedIterator{contract: _RegistrarContract.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

// WatchTransmitted is a free log subscription operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_RegistrarContract *RegistrarContractFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *RegistrarContractTransmitted) (event.Subscription, error) {

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractTransmitted)
				if err := _RegistrarContract.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

// ParseTransmitted is a log parse operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_RegistrarContract *RegistrarContractFilterer) ParseTransmitted(log types.Log) (*RegistrarContractTransmitted, error) {
	event := new(RegistrarContractTransmitted)
	if err := _RegistrarContract.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RegistrarContract contract.
type RegistrarContractUnpausedIterator struct {
	Event *RegistrarContractUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUnpaused)
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
		it.Event = new(RegistrarContractUnpaused)
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
func (it *RegistrarContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUnpaused represents a Unpaused event raised by the RegistrarContract contract.
type RegistrarContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RegistrarContract *RegistrarContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistrarContractUnpausedIterator, error) {

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUnpausedIterator{contract: _RegistrarContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RegistrarContract *RegistrarContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistrarContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUnpaused)
				if err := _RegistrarContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RegistrarContract *RegistrarContractFilterer) ParseUnpaused(log types.Log) (*RegistrarContractUnpaused, error) {
	event := new(RegistrarContractUnpaused)
	if err := _RegistrarContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepAdminTransferRequestedIterator is returned from FilterUpkeepAdminTransferRequested and is used to iterate over the raw logs and unpacked data for UpkeepAdminTransferRequested events raised by the RegistrarContract contract.
type RegistrarContractUpkeepAdminTransferRequestedIterator struct {
	Event *RegistrarContractUpkeepAdminTransferRequested // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepAdminTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepAdminTransferRequested)
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
		it.Event = new(RegistrarContractUpkeepAdminTransferRequested)
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
func (it *RegistrarContractUpkeepAdminTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepAdminTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepAdminTransferRequested represents a UpkeepAdminTransferRequested event raised by the RegistrarContract contract.
type RegistrarContractUpkeepAdminTransferRequested struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpkeepAdminTransferRequested is a free log retrieval operation binding the contract event 0xb1cbb2c4b8480034c27e06da5f096b8233a8fd4497028593a41ff6df79726b35.
//
// Solidity: event UpkeepAdminTransferRequested(uint256 indexed id, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepAdminTransferRequested(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*RegistrarContractUpkeepAdminTransferRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepAdminTransferRequested", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepAdminTransferRequestedIterator{contract: _RegistrarContract.contract, event: "UpkeepAdminTransferRequested", logs: logs, sub: sub}, nil
}

// WatchUpkeepAdminTransferRequested is a free log subscription operation binding the contract event 0xb1cbb2c4b8480034c27e06da5f096b8233a8fd4497028593a41ff6df79726b35.
//
// Solidity: event UpkeepAdminTransferRequested(uint256 indexed id, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepAdminTransferRequested(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepAdminTransferRequested, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepAdminTransferRequested", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepAdminTransferRequested)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepAdminTransferRequested", log); err != nil {
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

// ParseUpkeepAdminTransferRequested is a log parse operation binding the contract event 0xb1cbb2c4b8480034c27e06da5f096b8233a8fd4497028593a41ff6df79726b35.
//
// Solidity: event UpkeepAdminTransferRequested(uint256 indexed id, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepAdminTransferRequested(log types.Log) (*RegistrarContractUpkeepAdminTransferRequested, error) {
	event := new(RegistrarContractUpkeepAdminTransferRequested)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepAdminTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepAdminTransferredIterator is returned from FilterUpkeepAdminTransferred and is used to iterate over the raw logs and unpacked data for UpkeepAdminTransferred events raised by the RegistrarContract contract.
type RegistrarContractUpkeepAdminTransferredIterator struct {
	Event *RegistrarContractUpkeepAdminTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepAdminTransferred)
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
		it.Event = new(RegistrarContractUpkeepAdminTransferred)
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
func (it *RegistrarContractUpkeepAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepAdminTransferred represents a UpkeepAdminTransferred event raised by the RegistrarContract contract.
type RegistrarContractUpkeepAdminTransferred struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpkeepAdminTransferred is a free log retrieval operation binding the contract event 0x5cff4db96bef051785e999f44bfcd21c18823e034fb92dd376e3db4ce0feeb2c.
//
// Solidity: event UpkeepAdminTransferred(uint256 indexed id, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepAdminTransferred(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*RegistrarContractUpkeepAdminTransferredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepAdminTransferred", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepAdminTransferredIterator{contract: _RegistrarContract.contract, event: "UpkeepAdminTransferred", logs: logs, sub: sub}, nil
}

// WatchUpkeepAdminTransferred is a free log subscription operation binding the contract event 0x5cff4db96bef051785e999f44bfcd21c18823e034fb92dd376e3db4ce0feeb2c.
//
// Solidity: event UpkeepAdminTransferred(uint256 indexed id, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepAdminTransferred(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepAdminTransferred, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepAdminTransferred", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepAdminTransferred)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepAdminTransferred", log); err != nil {
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

// ParseUpkeepAdminTransferred is a log parse operation binding the contract event 0x5cff4db96bef051785e999f44bfcd21c18823e034fb92dd376e3db4ce0feeb2c.
//
// Solidity: event UpkeepAdminTransferred(uint256 indexed id, address indexed from, address indexed to)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepAdminTransferred(log types.Log) (*RegistrarContractUpkeepAdminTransferred, error) {
	event := new(RegistrarContractUpkeepAdminTransferred)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepCanceledIterator is returned from FilterUpkeepCanceled and is used to iterate over the raw logs and unpacked data for UpkeepCanceled events raised by the RegistrarContract contract.
type RegistrarContractUpkeepCanceledIterator struct {
	Event *RegistrarContractUpkeepCanceled // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepCanceled)
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
		it.Event = new(RegistrarContractUpkeepCanceled)
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
func (it *RegistrarContractUpkeepCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepCanceled represents a UpkeepCanceled event raised by the RegistrarContract contract.
type RegistrarContractUpkeepCanceled struct {
	Id            *big.Int
	AtBlockHeight uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpkeepCanceled is a free log retrieval operation binding the contract event 0x91cb3bb75cfbd718bbfccc56b7f53d92d7048ef4ca39a3b7b7c6d4af1f791181.
//
// Solidity: event UpkeepCanceled(uint256 indexed id, uint64 indexed atBlockHeight)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepCanceled(opts *bind.FilterOpts, id []*big.Int, atBlockHeight []uint64) (*RegistrarContractUpkeepCanceledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var atBlockHeightRule []interface{}
	for _, atBlockHeightItem := range atBlockHeight {
		atBlockHeightRule = append(atBlockHeightRule, atBlockHeightItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepCanceled", idRule, atBlockHeightRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepCanceledIterator{contract: _RegistrarContract.contract, event: "UpkeepCanceled", logs: logs, sub: sub}, nil
}

// WatchUpkeepCanceled is a free log subscription operation binding the contract event 0x91cb3bb75cfbd718bbfccc56b7f53d92d7048ef4ca39a3b7b7c6d4af1f791181.
//
// Solidity: event UpkeepCanceled(uint256 indexed id, uint64 indexed atBlockHeight)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepCanceled(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepCanceled, id []*big.Int, atBlockHeight []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var atBlockHeightRule []interface{}
	for _, atBlockHeightItem := range atBlockHeight {
		atBlockHeightRule = append(atBlockHeightRule, atBlockHeightItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepCanceled", idRule, atBlockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepCanceled)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepCanceled", log); err != nil {
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

// ParseUpkeepCanceled is a log parse operation binding the contract event 0x91cb3bb75cfbd718bbfccc56b7f53d92d7048ef4ca39a3b7b7c6d4af1f791181.
//
// Solidity: event UpkeepCanceled(uint256 indexed id, uint64 indexed atBlockHeight)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepCanceled(log types.Log) (*RegistrarContractUpkeepCanceled, error) {
	event := new(RegistrarContractUpkeepCanceled)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepCheckDataUpdatedIterator is returned from FilterUpkeepCheckDataUpdated and is used to iterate over the raw logs and unpacked data for UpkeepCheckDataUpdated events raised by the RegistrarContract contract.
type RegistrarContractUpkeepCheckDataUpdatedIterator struct {
	Event *RegistrarContractUpkeepCheckDataUpdated // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepCheckDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepCheckDataUpdated)
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
		it.Event = new(RegistrarContractUpkeepCheckDataUpdated)
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
func (it *RegistrarContractUpkeepCheckDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepCheckDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepCheckDataUpdated represents a UpkeepCheckDataUpdated event raised by the RegistrarContract contract.
type RegistrarContractUpkeepCheckDataUpdated struct {
	Id           *big.Int
	NewCheckData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpkeepCheckDataUpdated is a free log retrieval operation binding the contract event 0x7b778136e5211932b51a145badd01959415e79e051a933604b3d323f862dcabf.
//
// Solidity: event UpkeepCheckDataUpdated(uint256 indexed id, bytes newCheckData)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepCheckDataUpdated(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepCheckDataUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepCheckDataUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepCheckDataUpdatedIterator{contract: _RegistrarContract.contract, event: "UpkeepCheckDataUpdated", logs: logs, sub: sub}, nil
}

// WatchUpkeepCheckDataUpdated is a free log subscription operation binding the contract event 0x7b778136e5211932b51a145badd01959415e79e051a933604b3d323f862dcabf.
//
// Solidity: event UpkeepCheckDataUpdated(uint256 indexed id, bytes newCheckData)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepCheckDataUpdated(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepCheckDataUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepCheckDataUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepCheckDataUpdated)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepCheckDataUpdated", log); err != nil {
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

// ParseUpkeepCheckDataUpdated is a log parse operation binding the contract event 0x7b778136e5211932b51a145badd01959415e79e051a933604b3d323f862dcabf.
//
// Solidity: event UpkeepCheckDataUpdated(uint256 indexed id, bytes newCheckData)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepCheckDataUpdated(log types.Log) (*RegistrarContractUpkeepCheckDataUpdated, error) {
	event := new(RegistrarContractUpkeepCheckDataUpdated)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepCheckDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepGasLimitSetIterator is returned from FilterUpkeepGasLimitSet and is used to iterate over the raw logs and unpacked data for UpkeepGasLimitSet events raised by the RegistrarContract contract.
type RegistrarContractUpkeepGasLimitSetIterator struct {
	Event *RegistrarContractUpkeepGasLimitSet // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepGasLimitSet)
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
		it.Event = new(RegistrarContractUpkeepGasLimitSet)
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
func (it *RegistrarContractUpkeepGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepGasLimitSet represents a UpkeepGasLimitSet event raised by the RegistrarContract contract.
type RegistrarContractUpkeepGasLimitSet struct {
	Id       *big.Int
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpkeepGasLimitSet is a free log retrieval operation binding the contract event 0xc24c07e655ce79fba8a589778987d3c015bc6af1632bb20cf9182e02a65d972c.
//
// Solidity: event UpkeepGasLimitSet(uint256 indexed id, uint96 gasLimit)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepGasLimitSet(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepGasLimitSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepGasLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepGasLimitSetIterator{contract: _RegistrarContract.contract, event: "UpkeepGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchUpkeepGasLimitSet is a free log subscription operation binding the contract event 0xc24c07e655ce79fba8a589778987d3c015bc6af1632bb20cf9182e02a65d972c.
//
// Solidity: event UpkeepGasLimitSet(uint256 indexed id, uint96 gasLimit)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepGasLimitSet(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepGasLimitSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepGasLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepGasLimitSet)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepGasLimitSet", log); err != nil {
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

// ParseUpkeepGasLimitSet is a log parse operation binding the contract event 0xc24c07e655ce79fba8a589778987d3c015bc6af1632bb20cf9182e02a65d972c.
//
// Solidity: event UpkeepGasLimitSet(uint256 indexed id, uint96 gasLimit)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepGasLimitSet(log types.Log) (*RegistrarContractUpkeepGasLimitSet, error) {
	event := new(RegistrarContractUpkeepGasLimitSet)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepMigratedIterator is returned from FilterUpkeepMigrated and is used to iterate over the raw logs and unpacked data for UpkeepMigrated events raised by the RegistrarContract contract.
type RegistrarContractUpkeepMigratedIterator struct {
	Event *RegistrarContractUpkeepMigrated // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepMigrated)
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
		it.Event = new(RegistrarContractUpkeepMigrated)
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
func (it *RegistrarContractUpkeepMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepMigrated represents a UpkeepMigrated event raised by the RegistrarContract contract.
type RegistrarContractUpkeepMigrated struct {
	Id               *big.Int
	RemainingBalance *big.Int
	Destination      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpkeepMigrated is a free log retrieval operation binding the contract event 0xb38647142fbb1ea4c000fc4569b37a4e9a9f6313317b84ee3e5326c1a6cd06ff.
//
// Solidity: event UpkeepMigrated(uint256 indexed id, uint256 remainingBalance, address destination)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepMigrated(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepMigratedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepMigrated", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepMigratedIterator{contract: _RegistrarContract.contract, event: "UpkeepMigrated", logs: logs, sub: sub}, nil
}

// WatchUpkeepMigrated is a free log subscription operation binding the contract event 0xb38647142fbb1ea4c000fc4569b37a4e9a9f6313317b84ee3e5326c1a6cd06ff.
//
// Solidity: event UpkeepMigrated(uint256 indexed id, uint256 remainingBalance, address destination)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepMigrated(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepMigrated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepMigrated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepMigrated)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepMigrated", log); err != nil {
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

// ParseUpkeepMigrated is a log parse operation binding the contract event 0xb38647142fbb1ea4c000fc4569b37a4e9a9f6313317b84ee3e5326c1a6cd06ff.
//
// Solidity: event UpkeepMigrated(uint256 indexed id, uint256 remainingBalance, address destination)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepMigrated(log types.Log) (*RegistrarContractUpkeepMigrated, error) {
	event := new(RegistrarContractUpkeepMigrated)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepOffchainConfigSetIterator is returned from FilterUpkeepOffchainConfigSet and is used to iterate over the raw logs and unpacked data for UpkeepOffchainConfigSet events raised by the RegistrarContract contract.
type RegistrarContractUpkeepOffchainConfigSetIterator struct {
	Event *RegistrarContractUpkeepOffchainConfigSet // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepOffchainConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepOffchainConfigSet)
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
		it.Event = new(RegistrarContractUpkeepOffchainConfigSet)
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
func (it *RegistrarContractUpkeepOffchainConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepOffchainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepOffchainConfigSet represents a UpkeepOffchainConfigSet event raised by the RegistrarContract contract.
type RegistrarContractUpkeepOffchainConfigSet struct {
	Id             *big.Int
	OffchainConfig []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpkeepOffchainConfigSet is a free log retrieval operation binding the contract event 0x3e8740446213c8a77d40e08f79136ce3f347d13ed270a6ebdf57159e0faf4850.
//
// Solidity: event UpkeepOffchainConfigSet(uint256 indexed id, bytes offchainConfig)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepOffchainConfigSet(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepOffchainConfigSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepOffchainConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepOffchainConfigSetIterator{contract: _RegistrarContract.contract, event: "UpkeepOffchainConfigSet", logs: logs, sub: sub}, nil
}

// WatchUpkeepOffchainConfigSet is a free log subscription operation binding the contract event 0x3e8740446213c8a77d40e08f79136ce3f347d13ed270a6ebdf57159e0faf4850.
//
// Solidity: event UpkeepOffchainConfigSet(uint256 indexed id, bytes offchainConfig)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepOffchainConfigSet(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepOffchainConfigSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepOffchainConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepOffchainConfigSet)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepOffchainConfigSet", log); err != nil {
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
// Solidity: event UpkeepOffchainConfigSet(uint256 indexed id, bytes offchainConfig)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepOffchainConfigSet(log types.Log) (*RegistrarContractUpkeepOffchainConfigSet, error) {
	event := new(RegistrarContractUpkeepOffchainConfigSet)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepOffchainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepPausedIterator is returned from FilterUpkeepPaused and is used to iterate over the raw logs and unpacked data for UpkeepPaused events raised by the RegistrarContract contract.
type RegistrarContractUpkeepPausedIterator struct {
	Event *RegistrarContractUpkeepPaused // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepPaused)
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
		it.Event = new(RegistrarContractUpkeepPaused)
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
func (it *RegistrarContractUpkeepPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepPaused represents a UpkeepPaused event raised by the RegistrarContract contract.
type RegistrarContractUpkeepPaused struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpkeepPaused is a free log retrieval operation binding the contract event 0x8ab10247ce168c27748e656ecf852b951fcaac790c18106b19aa0ae57a8b741f.
//
// Solidity: event UpkeepPaused(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepPaused(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepPausedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepPaused", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepPausedIterator{contract: _RegistrarContract.contract, event: "UpkeepPaused", logs: logs, sub: sub}, nil
}

// WatchUpkeepPaused is a free log subscription operation binding the contract event 0x8ab10247ce168c27748e656ecf852b951fcaac790c18106b19aa0ae57a8b741f.
//
// Solidity: event UpkeepPaused(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepPaused(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepPaused, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepPaused", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepPaused)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepPaused", log); err != nil {
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
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepPaused(log types.Log) (*RegistrarContractUpkeepPaused, error) {
	event := new(RegistrarContractUpkeepPaused)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepPerformedIterator is returned from FilterUpkeepPerformed and is used to iterate over the raw logs and unpacked data for UpkeepPerformed events raised by the RegistrarContract contract.
type RegistrarContractUpkeepPerformedIterator struct {
	Event *RegistrarContractUpkeepPerformed // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepPerformed)
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
		it.Event = new(RegistrarContractUpkeepPerformed)
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
func (it *RegistrarContractUpkeepPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepPerformed represents a UpkeepPerformed event raised by the RegistrarContract contract.
type RegistrarContractUpkeepPerformed struct {
	Id               *big.Int
	Success          bool
	CheckBlockNumber uint32
	GasUsed          *big.Int
	GasOverhead      *big.Int
	TotalPayment     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpkeepPerformed is a free log retrieval operation binding the contract event 0x29233ba1d7b302b8fe230ad0b81423aba5371b2a6f6b821228212385ee6a4420.
//
// Solidity: event UpkeepPerformed(uint256 indexed id, bool indexed success, uint32 checkBlockNumber, uint256 gasUsed, uint256 gasOverhead, uint96 totalPayment)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepPerformed(opts *bind.FilterOpts, id []*big.Int, success []bool) (*RegistrarContractUpkeepPerformedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepPerformed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepPerformedIterator{contract: _RegistrarContract.contract, event: "UpkeepPerformed", logs: logs, sub: sub}, nil
}

// WatchUpkeepPerformed is a free log subscription operation binding the contract event 0x29233ba1d7b302b8fe230ad0b81423aba5371b2a6f6b821228212385ee6a4420.
//
// Solidity: event UpkeepPerformed(uint256 indexed id, bool indexed success, uint32 checkBlockNumber, uint256 gasUsed, uint256 gasOverhead, uint96 totalPayment)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepPerformed(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepPerformed, id []*big.Int, success []bool) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepPerformed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepPerformed)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepPerformed", log); err != nil {
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

// ParseUpkeepPerformed is a log parse operation binding the contract event 0x29233ba1d7b302b8fe230ad0b81423aba5371b2a6f6b821228212385ee6a4420.
//
// Solidity: event UpkeepPerformed(uint256 indexed id, bool indexed success, uint32 checkBlockNumber, uint256 gasUsed, uint256 gasOverhead, uint96 totalPayment)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepPerformed(log types.Log) (*RegistrarContractUpkeepPerformed, error) {
	event := new(RegistrarContractUpkeepPerformed)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepReceivedIterator is returned from FilterUpkeepReceived and is used to iterate over the raw logs and unpacked data for UpkeepReceived events raised by the RegistrarContract contract.
type RegistrarContractUpkeepReceivedIterator struct {
	Event *RegistrarContractUpkeepReceived // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepReceived)
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
		it.Event = new(RegistrarContractUpkeepReceived)
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
func (it *RegistrarContractUpkeepReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepReceived represents a UpkeepReceived event raised by the RegistrarContract contract.
type RegistrarContractUpkeepReceived struct {
	Id              *big.Int
	StartingBalance *big.Int
	ImportedFrom    common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpkeepReceived is a free log retrieval operation binding the contract event 0x74931a144e43a50694897f241d973aecb5024c0e910f9bb80a163ea3c1cf5a71.
//
// Solidity: event UpkeepReceived(uint256 indexed id, uint256 startingBalance, address importedFrom)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepReceived(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepReceivedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepReceived", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepReceivedIterator{contract: _RegistrarContract.contract, event: "UpkeepReceived", logs: logs, sub: sub}, nil
}

// WatchUpkeepReceived is a free log subscription operation binding the contract event 0x74931a144e43a50694897f241d973aecb5024c0e910f9bb80a163ea3c1cf5a71.
//
// Solidity: event UpkeepReceived(uint256 indexed id, uint256 startingBalance, address importedFrom)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepReceived(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepReceived, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepReceived", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepReceived)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepReceived", log); err != nil {
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

// ParseUpkeepReceived is a log parse operation binding the contract event 0x74931a144e43a50694897f241d973aecb5024c0e910f9bb80a163ea3c1cf5a71.
//
// Solidity: event UpkeepReceived(uint256 indexed id, uint256 startingBalance, address importedFrom)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepReceived(log types.Log) (*RegistrarContractUpkeepReceived, error) {
	event := new(RegistrarContractUpkeepReceived)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepRegisteredIterator is returned from FilterUpkeepRegistered and is used to iterate over the raw logs and unpacked data for UpkeepRegistered events raised by the RegistrarContract contract.
type RegistrarContractUpkeepRegisteredIterator struct {
	Event *RegistrarContractUpkeepRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepRegistered)
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
		it.Event = new(RegistrarContractUpkeepRegistered)
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
func (it *RegistrarContractUpkeepRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepRegistered represents a UpkeepRegistered event raised by the RegistrarContract contract.
type RegistrarContractUpkeepRegistered struct {
	Id         *big.Int
	ExecuteGas uint32
	Admin      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpkeepRegistered is a free log retrieval operation binding the contract event 0xbae366358c023f887e791d7a62f2e4316f1026bd77f6fb49501a917b3bc5d012.
//
// Solidity: event UpkeepRegistered(uint256 indexed id, uint32 executeGas, address admin)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepRegistered(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepRegisteredIterator{contract: _RegistrarContract.contract, event: "UpkeepRegistered", logs: logs, sub: sub}, nil
}

// WatchUpkeepRegistered is a free log subscription operation binding the contract event 0xbae366358c023f887e791d7a62f2e4316f1026bd77f6fb49501a917b3bc5d012.
//
// Solidity: event UpkeepRegistered(uint256 indexed id, uint32 executeGas, address admin)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepRegistered(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepRegistered, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepRegistered)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepRegistered", log); err != nil {
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

// ParseUpkeepRegistered is a log parse operation binding the contract event 0xbae366358c023f887e791d7a62f2e4316f1026bd77f6fb49501a917b3bc5d012.
//
// Solidity: event UpkeepRegistered(uint256 indexed id, uint32 executeGas, address admin)
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepRegistered(log types.Log) (*RegistrarContractUpkeepRegistered, error) {
	event := new(RegistrarContractUpkeepRegistered)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarContractUpkeepUnpausedIterator is returned from FilterUpkeepUnpaused and is used to iterate over the raw logs and unpacked data for UpkeepUnpaused events raised by the RegistrarContract contract.
type RegistrarContractUpkeepUnpausedIterator struct {
	Event *RegistrarContractUpkeepUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistrarContractUpkeepUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractUpkeepUnpaused)
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
		it.Event = new(RegistrarContractUpkeepUnpaused)
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
func (it *RegistrarContractUpkeepUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractUpkeepUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractUpkeepUnpaused represents a UpkeepUnpaused event raised by the RegistrarContract contract.
type RegistrarContractUpkeepUnpaused struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpkeepUnpaused is a free log retrieval operation binding the contract event 0x7bada562044eb163f6b4003c4553e4e62825344c0418eea087bed5ee05a47456.
//
// Solidity: event UpkeepUnpaused(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) FilterUpkeepUnpaused(opts *bind.FilterOpts, id []*big.Int) (*RegistrarContractUpkeepUnpausedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "UpkeepUnpaused", idRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractUpkeepUnpausedIterator{contract: _RegistrarContract.contract, event: "UpkeepUnpaused", logs: logs, sub: sub}, nil
}

// WatchUpkeepUnpaused is a free log subscription operation binding the contract event 0x7bada562044eb163f6b4003c4553e4e62825344c0418eea087bed5ee05a47456.
//
// Solidity: event UpkeepUnpaused(uint256 indexed id)
func (_RegistrarContract *RegistrarContractFilterer) WatchUpkeepUnpaused(opts *bind.WatchOpts, sink chan<- *RegistrarContractUpkeepUnpaused, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "UpkeepUnpaused", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractUpkeepUnpaused)
				if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepUnpaused", log); err != nil {
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
func (_RegistrarContract *RegistrarContractFilterer) ParseUpkeepUnpaused(log types.Log) (*RegistrarContractUpkeepUnpaused, error) {
	event := new(RegistrarContractUpkeepUnpaused)
	if err := _RegistrarContract.contract.UnpackLog(event, "UpkeepUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
