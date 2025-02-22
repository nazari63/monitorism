// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dispute

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

// FaultDisputeGameGameConstructorParams is an auto generated low-level Go binding around an user-defined struct.
type FaultDisputeGameGameConstructorParams struct {
	GameType            uint32
	AbsolutePrestate    [32]byte
	MaxGameDepth        *big.Int
	SplitDepth          *big.Int
	ClockExtension      uint64
	MaxClockDuration    uint64
	Vm                  common.Address
	Weth                common.Address
	AnchorStateRegistry common.Address
	L2ChainId           *big.Int
}

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
}

// DisputeGameMetaData contains all meta data concerning the DisputeGame contract.
var DisputeGameMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structFaultDisputeGame.GameConstructorParams\",\"components\":[{\"name\":\"gameType\",\"type\":\"uint32\",\"internalType\":\"GameType\"},{\"name\":\"absolutePrestate\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"maxGameDepth\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"splitDepth\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clockExtension\",\"type\":\"uint64\",\"internalType\":\"Duration\"},{\"name\":\"maxClockDuration\",\"type\":\"uint64\",\"internalType\":\"Duration\"},{\"name\":\"vm\",\"type\":\"address\",\"internalType\":\"contractIBigStepper\"},{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIDelayedWETH\"},{\"name\":\"anchorStateRegistry\",\"type\":\"address\",\"internalType\":\"contractIAnchorStateRegistry\"},{\"name\":\"l2ChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"absolutePrestate\",\"inputs\":[],\"outputs\":[{\"name\":\"absolutePrestate_\",\"type\":\"bytes32\",\"internalType\":\"Claim\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addLocalData\",\"inputs\":[{\"name\":\"_ident\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_execLeafIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_partOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"anchorStateRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"registry_\",\"type\":\"address\",\"internalType\":\"contractIAnchorStateRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attack\",\"inputs\":[{\"name\":\"_disputed\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"_parentIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_claim\",\"type\":\"bytes32\",\"internalType\":\"Claim\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"challengeRootL2Block\",\"inputs\":[{\"name\":\"_outputRootProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_headerRLP\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimCredit\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimData\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"parentIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"counteredBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bond\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"claim\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"position\",\"type\":\"uint128\",\"internalType\":\"Position\"},{\"name\":\"clock\",\"type\":\"uint128\",\"internalType\":\"Clock\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDataLen\",\"inputs\":[],\"outputs\":[{\"name\":\"len_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claims\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"Hash\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clockExtension\",\"inputs\":[],\"outputs\":[{\"name\":\"clockExtension_\",\"type\":\"uint64\",\"internalType\":\"Duration\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createdAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"Timestamp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"credit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defend\",\"inputs\":[{\"name\":\"_disputed\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"_parentIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_claim\",\"type\":\"bytes32\",\"internalType\":\"Claim\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"extraData\",\"inputs\":[],\"outputs\":[{\"name\":\"extraData_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"gameCreator\",\"inputs\":[],\"outputs\":[{\"name\":\"creator_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"gameData\",\"inputs\":[],\"outputs\":[{\"name\":\"gameType_\",\"type\":\"uint32\",\"internalType\":\"GameType\"},{\"name\":\"rootClaim_\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"extraData_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gameType\",\"inputs\":[],\"outputs\":[{\"name\":\"gameType_\",\"type\":\"uint32\",\"internalType\":\"GameType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallengerDuration\",\"inputs\":[{\"name\":\"_claimIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"duration_\",\"type\":\"uint64\",\"internalType\":\"Duration\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumToResolve\",\"inputs\":[{\"name\":\"_claimIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"numRemainingChildren_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredBond\",\"inputs\":[{\"name\":\"_position\",\"type\":\"uint128\",\"internalType\":\"Position\"}],\"outputs\":[{\"name\":\"requiredBond_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"l1Head\",\"inputs\":[],\"outputs\":[{\"name\":\"l1Head_\",\"type\":\"bytes32\",\"internalType\":\"Hash\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"l2BlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"l2BlockNumber_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"l2BlockNumberChallenged\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2BlockNumberChallenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2ChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"l2ChainId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxClockDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"maxClockDuration_\",\"type\":\"uint64\",\"internalType\":\"Duration\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGameDepth\",\"inputs\":[],\"outputs\":[{\"name\":\"maxGameDepth_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"move\",\"inputs\":[{\"name\":\"_disputed\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"_challengeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_claim\",\"type\":\"bytes32\",\"internalType\":\"Claim\"},{\"name\":\"_isAttack\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"resolutionCheckpoints\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"initialCheckpointComplete\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"subgameIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"leftmostPosition\",\"type\":\"uint128\",\"internalType\":\"Position\"},{\"name\":\"counteredBy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[],\"outputs\":[{\"name\":\"status_\",\"type\":\"uint8\",\"internalType\":\"enumGameStatus\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveClaim\",\"inputs\":[{\"name\":\"_claimIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_numToResolve\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolvedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"Timestamp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolvedSubgames\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rootClaim\",\"inputs\":[],\"outputs\":[{\"name\":\"rootClaim_\",\"type\":\"bytes32\",\"internalType\":\"Claim\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"splitDepth\",\"inputs\":[],\"outputs\":[{\"name\":\"splitDepth_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startingBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"startingBlockNumber_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startingOutputRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"Hash\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startingRootHash\",\"inputs\":[],\"outputs\":[{\"name\":\"startingRootHash_\",\"type\":\"bytes32\",\"internalType\":\"Hash\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"status\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumGameStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"step\",\"inputs\":[{\"name\":\"_claimIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isAttack\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_stateData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"subgames\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vm\",\"inputs\":[],\"outputs\":[{\"name\":\"vm_\",\"type\":\"address\",\"internalType\":\"contractIBigStepper\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"weth_\",\"type\":\"address\",\"internalType\":\"contractIDelayedWETH\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Move\",\"inputs\":[{\"name\":\"parentIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"claim\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"Claim\"},{\"name\":\"claimant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resolved\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumGameStatus\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AnchorRootNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberMatches\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BondTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotDefendRootClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimAboveSplit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimAlreadyResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClockNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClockTimeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContentLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateStep\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyItem\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GameDepthExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GameNotInProgress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectBondAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChallengePeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClockExtension\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDisputedClaimIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidHeader\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidHeaderRLP\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLocalIdent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutputRootProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidParent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrestate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSplitDepth\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2BlockNumberChallenged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxDepthTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoCreditToClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfOrderResolution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedRootClaim\",\"inputs\":[{\"name\":\"rootClaim\",\"type\":\"bytes32\",\"internalType\":\"Claim\"}]},{\"type\":\"error\",\"name\":\"UnexpectedString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidStep\",\"inputs\":[]}]",
}

// DisputeGameABI is the input ABI used to generate the binding from.
// Deprecated: Use DisputeGameMetaData.ABI instead.
var DisputeGameABI = DisputeGameMetaData.ABI

// DisputeGame is an auto generated Go binding around an Ethereum contract.
type DisputeGame struct {
	DisputeGameCaller     // Read-only binding to the contract
	DisputeGameTransactor // Write-only binding to the contract
	DisputeGameFilterer   // Log filterer for contract events
}

// DisputeGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type DisputeGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DisputeGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DisputeGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DisputeGameSession struct {
	Contract     *DisputeGame      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DisputeGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DisputeGameCallerSession struct {
	Contract *DisputeGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DisputeGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DisputeGameTransactorSession struct {
	Contract     *DisputeGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DisputeGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type DisputeGameRaw struct {
	Contract *DisputeGame // Generic contract binding to access the raw methods on
}

// DisputeGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DisputeGameCallerRaw struct {
	Contract *DisputeGameCaller // Generic read-only contract binding to access the raw methods on
}

// DisputeGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DisputeGameTransactorRaw struct {
	Contract *DisputeGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDisputeGame creates a new instance of DisputeGame, bound to a specific deployed contract.
func NewDisputeGame(address common.Address, backend bind.ContractBackend) (*DisputeGame, error) {
	contract, err := bindDisputeGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DisputeGame{DisputeGameCaller: DisputeGameCaller{contract: contract}, DisputeGameTransactor: DisputeGameTransactor{contract: contract}, DisputeGameFilterer: DisputeGameFilterer{contract: contract}}, nil
}

// NewDisputeGameCaller creates a new read-only instance of DisputeGame, bound to a specific deployed contract.
func NewDisputeGameCaller(address common.Address, caller bind.ContractCaller) (*DisputeGameCaller, error) {
	contract, err := bindDisputeGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeGameCaller{contract: contract}, nil
}

// NewDisputeGameTransactor creates a new write-only instance of DisputeGame, bound to a specific deployed contract.
func NewDisputeGameTransactor(address common.Address, transactor bind.ContractTransactor) (*DisputeGameTransactor, error) {
	contract, err := bindDisputeGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeGameTransactor{contract: contract}, nil
}

// NewDisputeGameFilterer creates a new log filterer instance of DisputeGame, bound to a specific deployed contract.
func NewDisputeGameFilterer(address common.Address, filterer bind.ContractFilterer) (*DisputeGameFilterer, error) {
	contract, err := bindDisputeGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DisputeGameFilterer{contract: contract}, nil
}

// bindDisputeGame binds a generic wrapper to an already deployed contract.
func bindDisputeGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DisputeGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeGame *DisputeGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeGame.Contract.DisputeGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeGame *DisputeGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeGame.Contract.DisputeGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeGame *DisputeGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeGame.Contract.DisputeGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeGame *DisputeGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeGame *DisputeGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeGame *DisputeGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeGame.Contract.contract.Transact(opts, method, params...)
}

// AbsolutePrestate is a free data retrieval call binding the contract method 0x8d450a95.
//
// Solidity: function absolutePrestate() view returns(bytes32 absolutePrestate_)
func (_DisputeGame *DisputeGameCaller) AbsolutePrestate(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "absolutePrestate")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AbsolutePrestate is a free data retrieval call binding the contract method 0x8d450a95.
//
// Solidity: function absolutePrestate() view returns(bytes32 absolutePrestate_)
func (_DisputeGame *DisputeGameSession) AbsolutePrestate() ([32]byte, error) {
	return _DisputeGame.Contract.AbsolutePrestate(&_DisputeGame.CallOpts)
}

// AbsolutePrestate is a free data retrieval call binding the contract method 0x8d450a95.
//
// Solidity: function absolutePrestate() view returns(bytes32 absolutePrestate_)
func (_DisputeGame *DisputeGameCallerSession) AbsolutePrestate() ([32]byte, error) {
	return _DisputeGame.Contract.AbsolutePrestate(&_DisputeGame.CallOpts)
}

// AnchorStateRegistry is a free data retrieval call binding the contract method 0x5c0cba33.
//
// Solidity: function anchorStateRegistry() view returns(address registry_)
func (_DisputeGame *DisputeGameCaller) AnchorStateRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "anchorStateRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnchorStateRegistry is a free data retrieval call binding the contract method 0x5c0cba33.
//
// Solidity: function anchorStateRegistry() view returns(address registry_)
func (_DisputeGame *DisputeGameSession) AnchorStateRegistry() (common.Address, error) {
	return _DisputeGame.Contract.AnchorStateRegistry(&_DisputeGame.CallOpts)
}

// AnchorStateRegistry is a free data retrieval call binding the contract method 0x5c0cba33.
//
// Solidity: function anchorStateRegistry() view returns(address registry_)
func (_DisputeGame *DisputeGameCallerSession) AnchorStateRegistry() (common.Address, error) {
	return _DisputeGame.Contract.AnchorStateRegistry(&_DisputeGame.CallOpts)
}

// ClaimData is a free data retrieval call binding the contract method 0xc6f0308c.
//
// Solidity: function claimData(uint256 ) view returns(uint32 parentIndex, address counteredBy, address claimant, uint128 bond, bytes32 claim, uint128 position, uint128 clock)
func (_DisputeGame *DisputeGameCaller) ClaimData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Claimant    common.Address
	Bond        *big.Int
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "claimData", arg0)

	outstruct := new(struct {
		ParentIndex uint32
		CounteredBy common.Address
		Claimant    common.Address
		Bond        *big.Int
		Claim       [32]byte
		Position    *big.Int
		Clock       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParentIndex = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.CounteredBy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Claimant = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Bond = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Claim = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.Position = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Clock = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ClaimData is a free data retrieval call binding the contract method 0xc6f0308c.
//
// Solidity: function claimData(uint256 ) view returns(uint32 parentIndex, address counteredBy, address claimant, uint128 bond, bytes32 claim, uint128 position, uint128 clock)
func (_DisputeGame *DisputeGameSession) ClaimData(arg0 *big.Int) (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Claimant    common.Address
	Bond        *big.Int
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error) {
	return _DisputeGame.Contract.ClaimData(&_DisputeGame.CallOpts, arg0)
}

// ClaimData is a free data retrieval call binding the contract method 0xc6f0308c.
//
// Solidity: function claimData(uint256 ) view returns(uint32 parentIndex, address counteredBy, address claimant, uint128 bond, bytes32 claim, uint128 position, uint128 clock)
func (_DisputeGame *DisputeGameCallerSession) ClaimData(arg0 *big.Int) (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Claimant    common.Address
	Bond        *big.Int
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error) {
	return _DisputeGame.Contract.ClaimData(&_DisputeGame.CallOpts, arg0)
}

// ClaimDataLen is a free data retrieval call binding the contract method 0x8980e0cc.
//
// Solidity: function claimDataLen() view returns(uint256 len_)
func (_DisputeGame *DisputeGameCaller) ClaimDataLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "claimDataLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimDataLen is a free data retrieval call binding the contract method 0x8980e0cc.
//
// Solidity: function claimDataLen() view returns(uint256 len_)
func (_DisputeGame *DisputeGameSession) ClaimDataLen() (*big.Int, error) {
	return _DisputeGame.Contract.ClaimDataLen(&_DisputeGame.CallOpts)
}

// ClaimDataLen is a free data retrieval call binding the contract method 0x8980e0cc.
//
// Solidity: function claimDataLen() view returns(uint256 len_)
func (_DisputeGame *DisputeGameCallerSession) ClaimDataLen() (*big.Int, error) {
	return _DisputeGame.Contract.ClaimDataLen(&_DisputeGame.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) view returns(bool)
func (_DisputeGame *DisputeGameCaller) Claims(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "claims", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) view returns(bool)
func (_DisputeGame *DisputeGameSession) Claims(arg0 [32]byte) (bool, error) {
	return _DisputeGame.Contract.Claims(&_DisputeGame.CallOpts, arg0)
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) view returns(bool)
func (_DisputeGame *DisputeGameCallerSession) Claims(arg0 [32]byte) (bool, error) {
	return _DisputeGame.Contract.Claims(&_DisputeGame.CallOpts, arg0)
}

// ClockExtension is a free data retrieval call binding the contract method 0x6b6716c0.
//
// Solidity: function clockExtension() view returns(uint64 clockExtension_)
func (_DisputeGame *DisputeGameCaller) ClockExtension(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "clockExtension")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ClockExtension is a free data retrieval call binding the contract method 0x6b6716c0.
//
// Solidity: function clockExtension() view returns(uint64 clockExtension_)
func (_DisputeGame *DisputeGameSession) ClockExtension() (uint64, error) {
	return _DisputeGame.Contract.ClockExtension(&_DisputeGame.CallOpts)
}

// ClockExtension is a free data retrieval call binding the contract method 0x6b6716c0.
//
// Solidity: function clockExtension() view returns(uint64 clockExtension_)
func (_DisputeGame *DisputeGameCallerSession) ClockExtension() (uint64, error) {
	return _DisputeGame.Contract.ClockExtension(&_DisputeGame.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64)
func (_DisputeGame *DisputeGameCaller) CreatedAt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64)
func (_DisputeGame *DisputeGameSession) CreatedAt() (uint64, error) {
	return _DisputeGame.Contract.CreatedAt(&_DisputeGame.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64)
func (_DisputeGame *DisputeGameCallerSession) CreatedAt() (uint64, error) {
	return _DisputeGame.Contract.CreatedAt(&_DisputeGame.CallOpts)
}

// Credit is a free data retrieval call binding the contract method 0xd5d44d80.
//
// Solidity: function credit(address ) view returns(uint256)
func (_DisputeGame *DisputeGameCaller) Credit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "credit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Credit is a free data retrieval call binding the contract method 0xd5d44d80.
//
// Solidity: function credit(address ) view returns(uint256)
func (_DisputeGame *DisputeGameSession) Credit(arg0 common.Address) (*big.Int, error) {
	return _DisputeGame.Contract.Credit(&_DisputeGame.CallOpts, arg0)
}

// Credit is a free data retrieval call binding the contract method 0xd5d44d80.
//
// Solidity: function credit(address ) view returns(uint256)
func (_DisputeGame *DisputeGameCallerSession) Credit(arg0 common.Address) (*big.Int, error) {
	return _DisputeGame.Contract.Credit(&_DisputeGame.CallOpts, arg0)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_DisputeGame *DisputeGameCaller) ExtraData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "extraData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_DisputeGame *DisputeGameSession) ExtraData() ([]byte, error) {
	return _DisputeGame.Contract.ExtraData(&_DisputeGame.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_DisputeGame *DisputeGameCallerSession) ExtraData() ([]byte, error) {
	return _DisputeGame.Contract.ExtraData(&_DisputeGame.CallOpts)
}

// GameCreator is a free data retrieval call binding the contract method 0x37b1b229.
//
// Solidity: function gameCreator() pure returns(address creator_)
func (_DisputeGame *DisputeGameCaller) GameCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "gameCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameCreator is a free data retrieval call binding the contract method 0x37b1b229.
//
// Solidity: function gameCreator() pure returns(address creator_)
func (_DisputeGame *DisputeGameSession) GameCreator() (common.Address, error) {
	return _DisputeGame.Contract.GameCreator(&_DisputeGame.CallOpts)
}

// GameCreator is a free data retrieval call binding the contract method 0x37b1b229.
//
// Solidity: function gameCreator() pure returns(address creator_)
func (_DisputeGame *DisputeGameCallerSession) GameCreator() (common.Address, error) {
	return _DisputeGame.Contract.GameCreator(&_DisputeGame.CallOpts)
}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() view returns(uint32 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_DisputeGame *DisputeGameCaller) GameData(opts *bind.CallOpts) (struct {
	GameType  uint32
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "gameData")

	outstruct := new(struct {
		GameType  uint32
		RootClaim [32]byte
		ExtraData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameType = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.RootClaim = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ExtraData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() view returns(uint32 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_DisputeGame *DisputeGameSession) GameData() (struct {
	GameType  uint32
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	return _DisputeGame.Contract.GameData(&_DisputeGame.CallOpts)
}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() view returns(uint32 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_DisputeGame *DisputeGameCallerSession) GameData() (struct {
	GameType  uint32
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	return _DisputeGame.Contract.GameData(&_DisputeGame.CallOpts)
}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() view returns(uint32 gameType_)
func (_DisputeGame *DisputeGameCaller) GameType(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "gameType")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() view returns(uint32 gameType_)
func (_DisputeGame *DisputeGameSession) GameType() (uint32, error) {
	return _DisputeGame.Contract.GameType(&_DisputeGame.CallOpts)
}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() view returns(uint32 gameType_)
func (_DisputeGame *DisputeGameCallerSession) GameType() (uint32, error) {
	return _DisputeGame.Contract.GameType(&_DisputeGame.CallOpts)
}

// GetChallengerDuration is a free data retrieval call binding the contract method 0xbd8da956.
//
// Solidity: function getChallengerDuration(uint256 _claimIndex) view returns(uint64 duration_)
func (_DisputeGame *DisputeGameCaller) GetChallengerDuration(opts *bind.CallOpts, _claimIndex *big.Int) (uint64, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "getChallengerDuration", _claimIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetChallengerDuration is a free data retrieval call binding the contract method 0xbd8da956.
//
// Solidity: function getChallengerDuration(uint256 _claimIndex) view returns(uint64 duration_)
func (_DisputeGame *DisputeGameSession) GetChallengerDuration(_claimIndex *big.Int) (uint64, error) {
	return _DisputeGame.Contract.GetChallengerDuration(&_DisputeGame.CallOpts, _claimIndex)
}

// GetChallengerDuration is a free data retrieval call binding the contract method 0xbd8da956.
//
// Solidity: function getChallengerDuration(uint256 _claimIndex) view returns(uint64 duration_)
func (_DisputeGame *DisputeGameCallerSession) GetChallengerDuration(_claimIndex *big.Int) (uint64, error) {
	return _DisputeGame.Contract.GetChallengerDuration(&_DisputeGame.CallOpts, _claimIndex)
}

// GetNumToResolve is a free data retrieval call binding the contract method 0x5a5fa2d9.
//
// Solidity: function getNumToResolve(uint256 _claimIndex) view returns(uint256 numRemainingChildren_)
func (_DisputeGame *DisputeGameCaller) GetNumToResolve(opts *bind.CallOpts, _claimIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "getNumToResolve", _claimIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumToResolve is a free data retrieval call binding the contract method 0x5a5fa2d9.
//
// Solidity: function getNumToResolve(uint256 _claimIndex) view returns(uint256 numRemainingChildren_)
func (_DisputeGame *DisputeGameSession) GetNumToResolve(_claimIndex *big.Int) (*big.Int, error) {
	return _DisputeGame.Contract.GetNumToResolve(&_DisputeGame.CallOpts, _claimIndex)
}

// GetNumToResolve is a free data retrieval call binding the contract method 0x5a5fa2d9.
//
// Solidity: function getNumToResolve(uint256 _claimIndex) view returns(uint256 numRemainingChildren_)
func (_DisputeGame *DisputeGameCallerSession) GetNumToResolve(_claimIndex *big.Int) (*big.Int, error) {
	return _DisputeGame.Contract.GetNumToResolve(&_DisputeGame.CallOpts, _claimIndex)
}

// GetRequiredBond is a free data retrieval call binding the contract method 0xc395e1ca.
//
// Solidity: function getRequiredBond(uint128 _position) view returns(uint256 requiredBond_)
func (_DisputeGame *DisputeGameCaller) GetRequiredBond(opts *bind.CallOpts, _position *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "getRequiredBond", _position)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredBond is a free data retrieval call binding the contract method 0xc395e1ca.
//
// Solidity: function getRequiredBond(uint128 _position) view returns(uint256 requiredBond_)
func (_DisputeGame *DisputeGameSession) GetRequiredBond(_position *big.Int) (*big.Int, error) {
	return _DisputeGame.Contract.GetRequiredBond(&_DisputeGame.CallOpts, _position)
}

// GetRequiredBond is a free data retrieval call binding the contract method 0xc395e1ca.
//
// Solidity: function getRequiredBond(uint128 _position) view returns(uint256 requiredBond_)
func (_DisputeGame *DisputeGameCallerSession) GetRequiredBond(_position *big.Int) (*big.Int, error) {
	return _DisputeGame.Contract.GetRequiredBond(&_DisputeGame.CallOpts, _position)
}

// L1Head is a free data retrieval call binding the contract method 0x6361506d.
//
// Solidity: function l1Head() pure returns(bytes32 l1Head_)
func (_DisputeGame *DisputeGameCaller) L1Head(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "l1Head")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1Head is a free data retrieval call binding the contract method 0x6361506d.
//
// Solidity: function l1Head() pure returns(bytes32 l1Head_)
func (_DisputeGame *DisputeGameSession) L1Head() ([32]byte, error) {
	return _DisputeGame.Contract.L1Head(&_DisputeGame.CallOpts)
}

// L1Head is a free data retrieval call binding the contract method 0x6361506d.
//
// Solidity: function l1Head() pure returns(bytes32 l1Head_)
func (_DisputeGame *DisputeGameCallerSession) L1Head() ([32]byte, error) {
	return _DisputeGame.Contract.L1Head(&_DisputeGame.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_DisputeGame *DisputeGameCaller) L2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "l2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_DisputeGame *DisputeGameSession) L2BlockNumber() (*big.Int, error) {
	return _DisputeGame.Contract.L2BlockNumber(&_DisputeGame.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_DisputeGame *DisputeGameCallerSession) L2BlockNumber() (*big.Int, error) {
	return _DisputeGame.Contract.L2BlockNumber(&_DisputeGame.CallOpts)
}

// L2BlockNumberChallenged is a free data retrieval call binding the contract method 0x3e3ac912.
//
// Solidity: function l2BlockNumberChallenged() view returns(bool)
func (_DisputeGame *DisputeGameCaller) L2BlockNumberChallenged(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "l2BlockNumberChallenged")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// L2BlockNumberChallenged is a free data retrieval call binding the contract method 0x3e3ac912.
//
// Solidity: function l2BlockNumberChallenged() view returns(bool)
func (_DisputeGame *DisputeGameSession) L2BlockNumberChallenged() (bool, error) {
	return _DisputeGame.Contract.L2BlockNumberChallenged(&_DisputeGame.CallOpts)
}

// L2BlockNumberChallenged is a free data retrieval call binding the contract method 0x3e3ac912.
//
// Solidity: function l2BlockNumberChallenged() view returns(bool)
func (_DisputeGame *DisputeGameCallerSession) L2BlockNumberChallenged() (bool, error) {
	return _DisputeGame.Contract.L2BlockNumberChallenged(&_DisputeGame.CallOpts)
}

// L2BlockNumberChallenger is a free data retrieval call binding the contract method 0x30dbe570.
//
// Solidity: function l2BlockNumberChallenger() view returns(address)
func (_DisputeGame *DisputeGameCaller) L2BlockNumberChallenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "l2BlockNumberChallenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2BlockNumberChallenger is a free data retrieval call binding the contract method 0x30dbe570.
//
// Solidity: function l2BlockNumberChallenger() view returns(address)
func (_DisputeGame *DisputeGameSession) L2BlockNumberChallenger() (common.Address, error) {
	return _DisputeGame.Contract.L2BlockNumberChallenger(&_DisputeGame.CallOpts)
}

// L2BlockNumberChallenger is a free data retrieval call binding the contract method 0x30dbe570.
//
// Solidity: function l2BlockNumberChallenger() view returns(address)
func (_DisputeGame *DisputeGameCallerSession) L2BlockNumberChallenger() (common.Address, error) {
	return _DisputeGame.Contract.L2BlockNumberChallenger(&_DisputeGame.CallOpts)
}

// L2ChainId is a free data retrieval call binding the contract method 0xd6ae3cd5.
//
// Solidity: function l2ChainId() view returns(uint256 l2ChainId_)
func (_DisputeGame *DisputeGameCaller) L2ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "l2ChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ChainId is a free data retrieval call binding the contract method 0xd6ae3cd5.
//
// Solidity: function l2ChainId() view returns(uint256 l2ChainId_)
func (_DisputeGame *DisputeGameSession) L2ChainId() (*big.Int, error) {
	return _DisputeGame.Contract.L2ChainId(&_DisputeGame.CallOpts)
}

// L2ChainId is a free data retrieval call binding the contract method 0xd6ae3cd5.
//
// Solidity: function l2ChainId() view returns(uint256 l2ChainId_)
func (_DisputeGame *DisputeGameCallerSession) L2ChainId() (*big.Int, error) {
	return _DisputeGame.Contract.L2ChainId(&_DisputeGame.CallOpts)
}

// MaxClockDuration is a free data retrieval call binding the contract method 0xdabd396d.
//
// Solidity: function maxClockDuration() view returns(uint64 maxClockDuration_)
func (_DisputeGame *DisputeGameCaller) MaxClockDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "maxClockDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxClockDuration is a free data retrieval call binding the contract method 0xdabd396d.
//
// Solidity: function maxClockDuration() view returns(uint64 maxClockDuration_)
func (_DisputeGame *DisputeGameSession) MaxClockDuration() (uint64, error) {
	return _DisputeGame.Contract.MaxClockDuration(&_DisputeGame.CallOpts)
}

// MaxClockDuration is a free data retrieval call binding the contract method 0xdabd396d.
//
// Solidity: function maxClockDuration() view returns(uint64 maxClockDuration_)
func (_DisputeGame *DisputeGameCallerSession) MaxClockDuration() (uint64, error) {
	return _DisputeGame.Contract.MaxClockDuration(&_DisputeGame.CallOpts)
}

// MaxGameDepth is a free data retrieval call binding the contract method 0xfa315aa9.
//
// Solidity: function maxGameDepth() view returns(uint256 maxGameDepth_)
func (_DisputeGame *DisputeGameCaller) MaxGameDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "maxGameDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGameDepth is a free data retrieval call binding the contract method 0xfa315aa9.
//
// Solidity: function maxGameDepth() view returns(uint256 maxGameDepth_)
func (_DisputeGame *DisputeGameSession) MaxGameDepth() (*big.Int, error) {
	return _DisputeGame.Contract.MaxGameDepth(&_DisputeGame.CallOpts)
}

// MaxGameDepth is a free data retrieval call binding the contract method 0xfa315aa9.
//
// Solidity: function maxGameDepth() view returns(uint256 maxGameDepth_)
func (_DisputeGame *DisputeGameCallerSession) MaxGameDepth() (*big.Int, error) {
	return _DisputeGame.Contract.MaxGameDepth(&_DisputeGame.CallOpts)
}

// ResolutionCheckpoints is a free data retrieval call binding the contract method 0xa445ece6.
//
// Solidity: function resolutionCheckpoints(uint256 ) view returns(bool initialCheckpointComplete, uint32 subgameIndex, uint128 leftmostPosition, address counteredBy)
func (_DisputeGame *DisputeGameCaller) ResolutionCheckpoints(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InitialCheckpointComplete bool
	SubgameIndex              uint32
	LeftmostPosition          *big.Int
	CounteredBy               common.Address
}, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "resolutionCheckpoints", arg0)

	outstruct := new(struct {
		InitialCheckpointComplete bool
		SubgameIndex              uint32
		LeftmostPosition          *big.Int
		CounteredBy               common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialCheckpointComplete = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SubgameIndex = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LeftmostPosition = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CounteredBy = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ResolutionCheckpoints is a free data retrieval call binding the contract method 0xa445ece6.
//
// Solidity: function resolutionCheckpoints(uint256 ) view returns(bool initialCheckpointComplete, uint32 subgameIndex, uint128 leftmostPosition, address counteredBy)
func (_DisputeGame *DisputeGameSession) ResolutionCheckpoints(arg0 *big.Int) (struct {
	InitialCheckpointComplete bool
	SubgameIndex              uint32
	LeftmostPosition          *big.Int
	CounteredBy               common.Address
}, error) {
	return _DisputeGame.Contract.ResolutionCheckpoints(&_DisputeGame.CallOpts, arg0)
}

// ResolutionCheckpoints is a free data retrieval call binding the contract method 0xa445ece6.
//
// Solidity: function resolutionCheckpoints(uint256 ) view returns(bool initialCheckpointComplete, uint32 subgameIndex, uint128 leftmostPosition, address counteredBy)
func (_DisputeGame *DisputeGameCallerSession) ResolutionCheckpoints(arg0 *big.Int) (struct {
	InitialCheckpointComplete bool
	SubgameIndex              uint32
	LeftmostPosition          *big.Int
	CounteredBy               common.Address
}, error) {
	return _DisputeGame.Contract.ResolutionCheckpoints(&_DisputeGame.CallOpts, arg0)
}

// ResolvedAt is a free data retrieval call binding the contract method 0x19effeb4.
//
// Solidity: function resolvedAt() view returns(uint64)
func (_DisputeGame *DisputeGameCaller) ResolvedAt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "resolvedAt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ResolvedAt is a free data retrieval call binding the contract method 0x19effeb4.
//
// Solidity: function resolvedAt() view returns(uint64)
func (_DisputeGame *DisputeGameSession) ResolvedAt() (uint64, error) {
	return _DisputeGame.Contract.ResolvedAt(&_DisputeGame.CallOpts)
}

// ResolvedAt is a free data retrieval call binding the contract method 0x19effeb4.
//
// Solidity: function resolvedAt() view returns(uint64)
func (_DisputeGame *DisputeGameCallerSession) ResolvedAt() (uint64, error) {
	return _DisputeGame.Contract.ResolvedAt(&_DisputeGame.CallOpts)
}

// ResolvedSubgames is a free data retrieval call binding the contract method 0xfe2bbeb2.
//
// Solidity: function resolvedSubgames(uint256 ) view returns(bool)
func (_DisputeGame *DisputeGameCaller) ResolvedSubgames(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "resolvedSubgames", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ResolvedSubgames is a free data retrieval call binding the contract method 0xfe2bbeb2.
//
// Solidity: function resolvedSubgames(uint256 ) view returns(bool)
func (_DisputeGame *DisputeGameSession) ResolvedSubgames(arg0 *big.Int) (bool, error) {
	return _DisputeGame.Contract.ResolvedSubgames(&_DisputeGame.CallOpts, arg0)
}

// ResolvedSubgames is a free data retrieval call binding the contract method 0xfe2bbeb2.
//
// Solidity: function resolvedSubgames(uint256 ) view returns(bool)
func (_DisputeGame *DisputeGameCallerSession) ResolvedSubgames(arg0 *big.Int) (bool, error) {
	return _DisputeGame.Contract.ResolvedSubgames(&_DisputeGame.CallOpts, arg0)
}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_DisputeGame *DisputeGameCaller) RootClaim(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "rootClaim")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_DisputeGame *DisputeGameSession) RootClaim() ([32]byte, error) {
	return _DisputeGame.Contract.RootClaim(&_DisputeGame.CallOpts)
}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_DisputeGame *DisputeGameCallerSession) RootClaim() ([32]byte, error) {
	return _DisputeGame.Contract.RootClaim(&_DisputeGame.CallOpts)
}

// SplitDepth is a free data retrieval call binding the contract method 0xec5e6308.
//
// Solidity: function splitDepth() view returns(uint256 splitDepth_)
func (_DisputeGame *DisputeGameCaller) SplitDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "splitDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SplitDepth is a free data retrieval call binding the contract method 0xec5e6308.
//
// Solidity: function splitDepth() view returns(uint256 splitDepth_)
func (_DisputeGame *DisputeGameSession) SplitDepth() (*big.Int, error) {
	return _DisputeGame.Contract.SplitDepth(&_DisputeGame.CallOpts)
}

// SplitDepth is a free data retrieval call binding the contract method 0xec5e6308.
//
// Solidity: function splitDepth() view returns(uint256 splitDepth_)
func (_DisputeGame *DisputeGameCallerSession) SplitDepth() (*big.Int, error) {
	return _DisputeGame.Contract.SplitDepth(&_DisputeGame.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256 startingBlockNumber_)
func (_DisputeGame *DisputeGameCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256 startingBlockNumber_)
func (_DisputeGame *DisputeGameSession) StartingBlockNumber() (*big.Int, error) {
	return _DisputeGame.Contract.StartingBlockNumber(&_DisputeGame.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256 startingBlockNumber_)
func (_DisputeGame *DisputeGameCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _DisputeGame.Contract.StartingBlockNumber(&_DisputeGame.CallOpts)
}

// StartingOutputRoot is a free data retrieval call binding the contract method 0x57da950e.
//
// Solidity: function startingOutputRoot() view returns(bytes32 root, uint256 l2BlockNumber)
func (_DisputeGame *DisputeGameCaller) StartingOutputRoot(opts *bind.CallOpts) (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "startingOutputRoot")

	outstruct := new(struct {
		Root          [32]byte
		L2BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.L2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StartingOutputRoot is a free data retrieval call binding the contract method 0x57da950e.
//
// Solidity: function startingOutputRoot() view returns(bytes32 root, uint256 l2BlockNumber)
func (_DisputeGame *DisputeGameSession) StartingOutputRoot() (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _DisputeGame.Contract.StartingOutputRoot(&_DisputeGame.CallOpts)
}

// StartingOutputRoot is a free data retrieval call binding the contract method 0x57da950e.
//
// Solidity: function startingOutputRoot() view returns(bytes32 root, uint256 l2BlockNumber)
func (_DisputeGame *DisputeGameCallerSession) StartingOutputRoot() (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _DisputeGame.Contract.StartingOutputRoot(&_DisputeGame.CallOpts)
}

// StartingRootHash is a free data retrieval call binding the contract method 0x25fc2ace.
//
// Solidity: function startingRootHash() view returns(bytes32 startingRootHash_)
func (_DisputeGame *DisputeGameCaller) StartingRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "startingRootHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StartingRootHash is a free data retrieval call binding the contract method 0x25fc2ace.
//
// Solidity: function startingRootHash() view returns(bytes32 startingRootHash_)
func (_DisputeGame *DisputeGameSession) StartingRootHash() ([32]byte, error) {
	return _DisputeGame.Contract.StartingRootHash(&_DisputeGame.CallOpts)
}

// StartingRootHash is a free data retrieval call binding the contract method 0x25fc2ace.
//
// Solidity: function startingRootHash() view returns(bytes32 startingRootHash_)
func (_DisputeGame *DisputeGameCallerSession) StartingRootHash() ([32]byte, error) {
	return _DisputeGame.Contract.StartingRootHash(&_DisputeGame.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_DisputeGame *DisputeGameCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_DisputeGame *DisputeGameSession) Status() (uint8, error) {
	return _DisputeGame.Contract.Status(&_DisputeGame.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_DisputeGame *DisputeGameCallerSession) Status() (uint8, error) {
	return _DisputeGame.Contract.Status(&_DisputeGame.CallOpts)
}

// Subgames is a free data retrieval call binding the contract method 0x2ad69aeb.
//
// Solidity: function subgames(uint256 , uint256 ) view returns(uint256)
func (_DisputeGame *DisputeGameCaller) Subgames(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "subgames", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Subgames is a free data retrieval call binding the contract method 0x2ad69aeb.
//
// Solidity: function subgames(uint256 , uint256 ) view returns(uint256)
func (_DisputeGame *DisputeGameSession) Subgames(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DisputeGame.Contract.Subgames(&_DisputeGame.CallOpts, arg0, arg1)
}

// Subgames is a free data retrieval call binding the contract method 0x2ad69aeb.
//
// Solidity: function subgames(uint256 , uint256 ) view returns(uint256)
func (_DisputeGame *DisputeGameCallerSession) Subgames(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DisputeGame.Contract.Subgames(&_DisputeGame.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DisputeGame *DisputeGameCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DisputeGame *DisputeGameSession) Version() (string, error) {
	return _DisputeGame.Contract.Version(&_DisputeGame.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DisputeGame *DisputeGameCallerSession) Version() (string, error) {
	return _DisputeGame.Contract.Version(&_DisputeGame.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() view returns(address vm_)
func (_DisputeGame *DisputeGameCaller) Vm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "vm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() view returns(address vm_)
func (_DisputeGame *DisputeGameSession) Vm() (common.Address, error) {
	return _DisputeGame.Contract.Vm(&_DisputeGame.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() view returns(address vm_)
func (_DisputeGame *DisputeGameCallerSession) Vm() (common.Address, error) {
	return _DisputeGame.Contract.Vm(&_DisputeGame.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address weth_)
func (_DisputeGame *DisputeGameCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeGame.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address weth_)
func (_DisputeGame *DisputeGameSession) Weth() (common.Address, error) {
	return _DisputeGame.Contract.Weth(&_DisputeGame.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address weth_)
func (_DisputeGame *DisputeGameCallerSession) Weth() (common.Address, error) {
	return _DisputeGame.Contract.Weth(&_DisputeGame.CallOpts)
}

// AddLocalData is a paid mutator transaction binding the contract method 0xf8f43ff6.
//
// Solidity: function addLocalData(uint256 _ident, uint256 _execLeafIdx, uint256 _partOffset) returns()
func (_DisputeGame *DisputeGameTransactor) AddLocalData(opts *bind.TransactOpts, _ident *big.Int, _execLeafIdx *big.Int, _partOffset *big.Int) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "addLocalData", _ident, _execLeafIdx, _partOffset)
}

// AddLocalData is a paid mutator transaction binding the contract method 0xf8f43ff6.
//
// Solidity: function addLocalData(uint256 _ident, uint256 _execLeafIdx, uint256 _partOffset) returns()
func (_DisputeGame *DisputeGameSession) AddLocalData(_ident *big.Int, _execLeafIdx *big.Int, _partOffset *big.Int) (*types.Transaction, error) {
	return _DisputeGame.Contract.AddLocalData(&_DisputeGame.TransactOpts, _ident, _execLeafIdx, _partOffset)
}

// AddLocalData is a paid mutator transaction binding the contract method 0xf8f43ff6.
//
// Solidity: function addLocalData(uint256 _ident, uint256 _execLeafIdx, uint256 _partOffset) returns()
func (_DisputeGame *DisputeGameTransactorSession) AddLocalData(_ident *big.Int, _execLeafIdx *big.Int, _partOffset *big.Int) (*types.Transaction, error) {
	return _DisputeGame.Contract.AddLocalData(&_DisputeGame.TransactOpts, _ident, _execLeafIdx, _partOffset)
}

// Attack is a paid mutator transaction binding the contract method 0x472777c6.
//
// Solidity: function attack(bytes32 _disputed, uint256 _parentIndex, bytes32 _claim) payable returns()
func (_DisputeGame *DisputeGameTransactor) Attack(opts *bind.TransactOpts, _disputed [32]byte, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "attack", _disputed, _parentIndex, _claim)
}

// Attack is a paid mutator transaction binding the contract method 0x472777c6.
//
// Solidity: function attack(bytes32 _disputed, uint256 _parentIndex, bytes32 _claim) payable returns()
func (_DisputeGame *DisputeGameSession) Attack(_disputed [32]byte, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.Attack(&_DisputeGame.TransactOpts, _disputed, _parentIndex, _claim)
}

// Attack is a paid mutator transaction binding the contract method 0x472777c6.
//
// Solidity: function attack(bytes32 _disputed, uint256 _parentIndex, bytes32 _claim) payable returns()
func (_DisputeGame *DisputeGameTransactorSession) Attack(_disputed [32]byte, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.Attack(&_DisputeGame.TransactOpts, _disputed, _parentIndex, _claim)
}

// ChallengeRootL2Block is a paid mutator transaction binding the contract method 0x01935130.
//
// Solidity: function challengeRootL2Block((bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes _headerRLP) returns()
func (_DisputeGame *DisputeGameTransactor) ChallengeRootL2Block(opts *bind.TransactOpts, _outputRootProof TypesOutputRootProof, _headerRLP []byte) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "challengeRootL2Block", _outputRootProof, _headerRLP)
}

// ChallengeRootL2Block is a paid mutator transaction binding the contract method 0x01935130.
//
// Solidity: function challengeRootL2Block((bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes _headerRLP) returns()
func (_DisputeGame *DisputeGameSession) ChallengeRootL2Block(_outputRootProof TypesOutputRootProof, _headerRLP []byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.ChallengeRootL2Block(&_DisputeGame.TransactOpts, _outputRootProof, _headerRLP)
}

// ChallengeRootL2Block is a paid mutator transaction binding the contract method 0x01935130.
//
// Solidity: function challengeRootL2Block((bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes _headerRLP) returns()
func (_DisputeGame *DisputeGameTransactorSession) ChallengeRootL2Block(_outputRootProof TypesOutputRootProof, _headerRLP []byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.ChallengeRootL2Block(&_DisputeGame.TransactOpts, _outputRootProof, _headerRLP)
}

// ClaimCredit is a paid mutator transaction binding the contract method 0x60e27464.
//
// Solidity: function claimCredit(address _recipient) returns()
func (_DisputeGame *DisputeGameTransactor) ClaimCredit(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "claimCredit", _recipient)
}

// ClaimCredit is a paid mutator transaction binding the contract method 0x60e27464.
//
// Solidity: function claimCredit(address _recipient) returns()
func (_DisputeGame *DisputeGameSession) ClaimCredit(_recipient common.Address) (*types.Transaction, error) {
	return _DisputeGame.Contract.ClaimCredit(&_DisputeGame.TransactOpts, _recipient)
}

// ClaimCredit is a paid mutator transaction binding the contract method 0x60e27464.
//
// Solidity: function claimCredit(address _recipient) returns()
func (_DisputeGame *DisputeGameTransactorSession) ClaimCredit(_recipient common.Address) (*types.Transaction, error) {
	return _DisputeGame.Contract.ClaimCredit(&_DisputeGame.TransactOpts, _recipient)
}

// Defend is a paid mutator transaction binding the contract method 0x7b0f0adc.
//
// Solidity: function defend(bytes32 _disputed, uint256 _parentIndex, bytes32 _claim) payable returns()
func (_DisputeGame *DisputeGameTransactor) Defend(opts *bind.TransactOpts, _disputed [32]byte, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "defend", _disputed, _parentIndex, _claim)
}

// Defend is a paid mutator transaction binding the contract method 0x7b0f0adc.
//
// Solidity: function defend(bytes32 _disputed, uint256 _parentIndex, bytes32 _claim) payable returns()
func (_DisputeGame *DisputeGameSession) Defend(_disputed [32]byte, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.Defend(&_DisputeGame.TransactOpts, _disputed, _parentIndex, _claim)
}

// Defend is a paid mutator transaction binding the contract method 0x7b0f0adc.
//
// Solidity: function defend(bytes32 _disputed, uint256 _parentIndex, bytes32 _claim) payable returns()
func (_DisputeGame *DisputeGameTransactorSession) Defend(_disputed [32]byte, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.Defend(&_DisputeGame.TransactOpts, _disputed, _parentIndex, _claim)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_DisputeGame *DisputeGameTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_DisputeGame *DisputeGameSession) Initialize() (*types.Transaction, error) {
	return _DisputeGame.Contract.Initialize(&_DisputeGame.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_DisputeGame *DisputeGameTransactorSession) Initialize() (*types.Transaction, error) {
	return _DisputeGame.Contract.Initialize(&_DisputeGame.TransactOpts)
}

// Move is a paid mutator transaction binding the contract method 0x6f034409.
//
// Solidity: function move(bytes32 _disputed, uint256 _challengeIndex, bytes32 _claim, bool _isAttack) payable returns()
func (_DisputeGame *DisputeGameTransactor) Move(opts *bind.TransactOpts, _disputed [32]byte, _challengeIndex *big.Int, _claim [32]byte, _isAttack bool) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "move", _disputed, _challengeIndex, _claim, _isAttack)
}

// Move is a paid mutator transaction binding the contract method 0x6f034409.
//
// Solidity: function move(bytes32 _disputed, uint256 _challengeIndex, bytes32 _claim, bool _isAttack) payable returns()
func (_DisputeGame *DisputeGameSession) Move(_disputed [32]byte, _challengeIndex *big.Int, _claim [32]byte, _isAttack bool) (*types.Transaction, error) {
	return _DisputeGame.Contract.Move(&_DisputeGame.TransactOpts, _disputed, _challengeIndex, _claim, _isAttack)
}

// Move is a paid mutator transaction binding the contract method 0x6f034409.
//
// Solidity: function move(bytes32 _disputed, uint256 _challengeIndex, bytes32 _claim, bool _isAttack) payable returns()
func (_DisputeGame *DisputeGameTransactorSession) Move(_disputed [32]byte, _challengeIndex *big.Int, _claim [32]byte, _isAttack bool) (*types.Transaction, error) {
	return _DisputeGame.Contract.Move(&_DisputeGame.TransactOpts, _disputed, _challengeIndex, _claim, _isAttack)
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8 status_)
func (_DisputeGame *DisputeGameTransactor) Resolve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "resolve")
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8 status_)
func (_DisputeGame *DisputeGameSession) Resolve() (*types.Transaction, error) {
	return _DisputeGame.Contract.Resolve(&_DisputeGame.TransactOpts)
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8 status_)
func (_DisputeGame *DisputeGameTransactorSession) Resolve() (*types.Transaction, error) {
	return _DisputeGame.Contract.Resolve(&_DisputeGame.TransactOpts)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0x03c2924d.
//
// Solidity: function resolveClaim(uint256 _claimIndex, uint256 _numToResolve) returns()
func (_DisputeGame *DisputeGameTransactor) ResolveClaim(opts *bind.TransactOpts, _claimIndex *big.Int, _numToResolve *big.Int) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "resolveClaim", _claimIndex, _numToResolve)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0x03c2924d.
//
// Solidity: function resolveClaim(uint256 _claimIndex, uint256 _numToResolve) returns()
func (_DisputeGame *DisputeGameSession) ResolveClaim(_claimIndex *big.Int, _numToResolve *big.Int) (*types.Transaction, error) {
	return _DisputeGame.Contract.ResolveClaim(&_DisputeGame.TransactOpts, _claimIndex, _numToResolve)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0x03c2924d.
//
// Solidity: function resolveClaim(uint256 _claimIndex, uint256 _numToResolve) returns()
func (_DisputeGame *DisputeGameTransactorSession) ResolveClaim(_claimIndex *big.Int, _numToResolve *big.Int) (*types.Transaction, error) {
	return _DisputeGame.Contract.ResolveClaim(&_DisputeGame.TransactOpts, _claimIndex, _numToResolve)
}

// Step is a paid mutator transaction binding the contract method 0xd8cc1a3c.
//
// Solidity: function step(uint256 _claimIndex, bool _isAttack, bytes _stateData, bytes _proof) returns()
func (_DisputeGame *DisputeGameTransactor) Step(opts *bind.TransactOpts, _claimIndex *big.Int, _isAttack bool, _stateData []byte, _proof []byte) (*types.Transaction, error) {
	return _DisputeGame.contract.Transact(opts, "step", _claimIndex, _isAttack, _stateData, _proof)
}

// Step is a paid mutator transaction binding the contract method 0xd8cc1a3c.
//
// Solidity: function step(uint256 _claimIndex, bool _isAttack, bytes _stateData, bytes _proof) returns()
func (_DisputeGame *DisputeGameSession) Step(_claimIndex *big.Int, _isAttack bool, _stateData []byte, _proof []byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.Step(&_DisputeGame.TransactOpts, _claimIndex, _isAttack, _stateData, _proof)
}

// Step is a paid mutator transaction binding the contract method 0xd8cc1a3c.
//
// Solidity: function step(uint256 _claimIndex, bool _isAttack, bytes _stateData, bytes _proof) returns()
func (_DisputeGame *DisputeGameTransactorSession) Step(_claimIndex *big.Int, _isAttack bool, _stateData []byte, _proof []byte) (*types.Transaction, error) {
	return _DisputeGame.Contract.Step(&_DisputeGame.TransactOpts, _claimIndex, _isAttack, _stateData, _proof)
}

// DisputeGameMoveIterator is returned from FilterMove and is used to iterate over the raw logs and unpacked data for Move events raised by the DisputeGame contract.
type DisputeGameMoveIterator struct {
	Event *DisputeGameMove // Event containing the contract specifics and raw log

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
func (it *DisputeGameMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeGameMove)
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
		it.Event = new(DisputeGameMove)
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
func (it *DisputeGameMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeGameMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeGameMove represents a Move event raised by the DisputeGame contract.
type DisputeGameMove struct {
	ParentIndex *big.Int
	Claim       [32]byte
	Claimant    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMove is a free log retrieval operation binding the contract event 0x9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be.
//
// Solidity: event Move(uint256 indexed parentIndex, bytes32 indexed claim, address indexed claimant)
func (_DisputeGame *DisputeGameFilterer) FilterMove(opts *bind.FilterOpts, parentIndex []*big.Int, claim [][32]byte, claimant []common.Address) (*DisputeGameMoveIterator, error) {

	var parentIndexRule []interface{}
	for _, parentIndexItem := range parentIndex {
		parentIndexRule = append(parentIndexRule, parentIndexItem)
	}
	var claimRule []interface{}
	for _, claimItem := range claim {
		claimRule = append(claimRule, claimItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _DisputeGame.contract.FilterLogs(opts, "Move", parentIndexRule, claimRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return &DisputeGameMoveIterator{contract: _DisputeGame.contract, event: "Move", logs: logs, sub: sub}, nil
}

// WatchMove is a free log subscription operation binding the contract event 0x9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be.
//
// Solidity: event Move(uint256 indexed parentIndex, bytes32 indexed claim, address indexed claimant)
func (_DisputeGame *DisputeGameFilterer) WatchMove(opts *bind.WatchOpts, sink chan<- *DisputeGameMove, parentIndex []*big.Int, claim [][32]byte, claimant []common.Address) (event.Subscription, error) {

	var parentIndexRule []interface{}
	for _, parentIndexItem := range parentIndex {
		parentIndexRule = append(parentIndexRule, parentIndexItem)
	}
	var claimRule []interface{}
	for _, claimItem := range claim {
		claimRule = append(claimRule, claimItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _DisputeGame.contract.WatchLogs(opts, "Move", parentIndexRule, claimRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeGameMove)
				if err := _DisputeGame.contract.UnpackLog(event, "Move", log); err != nil {
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

// ParseMove is a log parse operation binding the contract event 0x9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be.
//
// Solidity: event Move(uint256 indexed parentIndex, bytes32 indexed claim, address indexed claimant)
func (_DisputeGame *DisputeGameFilterer) ParseMove(log types.Log) (*DisputeGameMove, error) {
	event := new(DisputeGameMove)
	if err := _DisputeGame.contract.UnpackLog(event, "Move", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeGameResolvedIterator is returned from FilterResolved and is used to iterate over the raw logs and unpacked data for Resolved events raised by the DisputeGame contract.
type DisputeGameResolvedIterator struct {
	Event *DisputeGameResolved // Event containing the contract specifics and raw log

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
func (it *DisputeGameResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeGameResolved)
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
		it.Event = new(DisputeGameResolved)
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
func (it *DisputeGameResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeGameResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeGameResolved represents a Resolved event raised by the DisputeGame contract.
type DisputeGameResolved struct {
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResolved is a free log retrieval operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_DisputeGame *DisputeGameFilterer) FilterResolved(opts *bind.FilterOpts, status []uint8) (*DisputeGameResolvedIterator, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _DisputeGame.contract.FilterLogs(opts, "Resolved", statusRule)
	if err != nil {
		return nil, err
	}
	return &DisputeGameResolvedIterator{contract: _DisputeGame.contract, event: "Resolved", logs: logs, sub: sub}, nil
}

// WatchResolved is a free log subscription operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_DisputeGame *DisputeGameFilterer) WatchResolved(opts *bind.WatchOpts, sink chan<- *DisputeGameResolved, status []uint8) (event.Subscription, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _DisputeGame.contract.WatchLogs(opts, "Resolved", statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeGameResolved)
				if err := _DisputeGame.contract.UnpackLog(event, "Resolved", log); err != nil {
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

// ParseResolved is a log parse operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_DisputeGame *DisputeGameFilterer) ParseResolved(log types.Log) (*DisputeGameResolved, error) {
	event := new(DisputeGameResolved)
	if err := _DisputeGame.contract.UnpackLog(event, "Resolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
