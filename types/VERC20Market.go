// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package types

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

// VERC20Order is an auto generated low-level Go binding around an user-defined struct.
type VERC20Order struct {
	Maker          common.Address
	Sell           bool
	ListId         [32]byte
	Tick           string
	Amount         *big.Int
	Price          *big.Int
	ListingTime    uint64
	ExpirationTime uint64
	V              uint8
	R              [32]byte
	S              [32]byte
}

// VERC20MarketMetaData contains all meta data concerning the VERC20Market contract.
var VERC20MarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOrdersMatched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoncesInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trustedVerifier\",\"type\":\"address\"}],\"name\":\"NewTrustedVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sell\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"listId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"VERC20OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sell\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"listId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"VERC20OrderExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"verc20_protocol_TransferVERC20Token\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"verc20_protocol_TransferVERC20TokenForListing\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"sell\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"listId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"listingTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structVERC20Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"sell\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"listId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"listingTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structVERC20Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"}],\"name\":\"executeOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedVerifier\",\"type\":\"address\"}],\"name\":\"updateTrustedVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawUnexpectedERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VERC20MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use VERC20MarketMetaData.ABI instead.
var VERC20MarketABI = VERC20MarketMetaData.ABI

// VERC20Market is an auto generated Go binding around an Ethereum contract.
type VERC20Market struct {
	VERC20MarketCaller     // Read-only binding to the contract
	VERC20MarketTransactor // Write-only binding to the contract
	VERC20MarketFilterer   // Log filterer for contract events
}

// VERC20MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type VERC20MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VERC20MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VERC20MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VERC20MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VERC20MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VERC20MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VERC20MarketSession struct {
	Contract     *VERC20Market     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VERC20MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VERC20MarketCallerSession struct {
	Contract *VERC20MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VERC20MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VERC20MarketTransactorSession struct {
	Contract     *VERC20MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VERC20MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type VERC20MarketRaw struct {
	Contract *VERC20Market // Generic contract binding to access the raw methods on
}

// VERC20MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VERC20MarketCallerRaw struct {
	Contract *VERC20MarketCaller // Generic read-only contract binding to access the raw methods on
}

// VERC20MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VERC20MarketTransactorRaw struct {
	Contract *VERC20MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVERC20Market creates a new instance of VERC20Market, bound to a specific deployed contract.
func NewVERC20Market(address common.Address, backend bind.ContractBackend) (*VERC20Market, error) {
	contract, err := bindVERC20Market(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VERC20Market{VERC20MarketCaller: VERC20MarketCaller{contract: contract}, VERC20MarketTransactor: VERC20MarketTransactor{contract: contract}, VERC20MarketFilterer: VERC20MarketFilterer{contract: contract}}, nil
}

// NewVERC20MarketCaller creates a new read-only instance of VERC20Market, bound to a specific deployed contract.
func NewVERC20MarketCaller(address common.Address, caller bind.ContractCaller) (*VERC20MarketCaller, error) {
	contract, err := bindVERC20Market(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketCaller{contract: contract}, nil
}

// NewVERC20MarketTransactor creates a new write-only instance of VERC20Market, bound to a specific deployed contract.
func NewVERC20MarketTransactor(address common.Address, transactor bind.ContractTransactor) (*VERC20MarketTransactor, error) {
	contract, err := bindVERC20Market(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketTransactor{contract: contract}, nil
}

// NewVERC20MarketFilterer creates a new log filterer instance of VERC20Market, bound to a specific deployed contract.
func NewVERC20MarketFilterer(address common.Address, filterer bind.ContractFilterer) (*VERC20MarketFilterer, error) {
	contract, err := bindVERC20Market(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketFilterer{contract: contract}, nil
}

// bindVERC20Market binds a generic wrapper to an already deployed contract.
func bindVERC20Market(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VERC20MarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VERC20Market *VERC20MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VERC20Market.Contract.VERC20MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VERC20Market *VERC20MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VERC20Market.Contract.VERC20MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VERC20Market *VERC20MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VERC20Market.Contract.VERC20MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VERC20Market *VERC20MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VERC20Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VERC20Market *VERC20MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VERC20Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VERC20Market *VERC20MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VERC20Market.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VERC20Market *VERC20MarketCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _VERC20Market.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VERC20Market *VERC20MarketSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VERC20Market.Contract.Eip712Domain(&_VERC20Market.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VERC20Market *VERC20MarketCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VERC20Market.Contract.Eip712Domain(&_VERC20Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VERC20Market *VERC20MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VERC20Market.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VERC20Market *VERC20MarketSession) Owner() (common.Address, error) {
	return _VERC20Market.Contract.Owner(&_VERC20Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VERC20Market *VERC20MarketCallerSession) Owner() (common.Address, error) {
	return _VERC20Market.Contract.Owner(&_VERC20Market.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VERC20Market *VERC20MarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VERC20Market.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VERC20Market *VERC20MarketSession) Paused() (bool, error) {
	return _VERC20Market.Contract.Paused(&_VERC20Market.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VERC20Market *VERC20MarketCallerSession) Paused() (bool, error) {
	return _VERC20Market.Contract.Paused(&_VERC20Market.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VERC20Market *VERC20MarketCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VERC20Market.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VERC20Market *VERC20MarketSession) ProxiableUUID() ([32]byte, error) {
	return _VERC20Market.Contract.ProxiableUUID(&_VERC20Market.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VERC20Market *VERC20MarketCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VERC20Market.Contract.ProxiableUUID(&_VERC20Market.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x52327b7d.
//
// Solidity: function cancelOrder((address,bool,bytes32,string,uint256,uint256,uint64,uint64,uint8,bytes32,bytes32) order) returns()
func (_VERC20Market *VERC20MarketTransactor) CancelOrder(opts *bind.TransactOpts, order VERC20Order) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x52327b7d.
//
// Solidity: function cancelOrder((address,bool,bytes32,string,uint256,uint256,uint64,uint64,uint8,bytes32,bytes32) order) returns()
func (_VERC20Market *VERC20MarketSession) CancelOrder(order VERC20Order) (*types.Transaction, error) {
	return _VERC20Market.Contract.CancelOrder(&_VERC20Market.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x52327b7d.
//
// Solidity: function cancelOrder((address,bool,bytes32,string,uint256,uint256,uint64,uint64,uint8,bytes32,bytes32) order) returns()
func (_VERC20Market *VERC20MarketTransactorSession) CancelOrder(order VERC20Order) (*types.Transaction, error) {
	return _VERC20Market.Contract.CancelOrder(&_VERC20Market.TransactOpts, order)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x87649e4f.
//
// Solidity: function executeOrder((address,bool,bytes32,string,uint256,uint256,uint64,uint64,uint8,bytes32,bytes32) order, address taker) payable returns()
func (_VERC20Market *VERC20MarketTransactor) ExecuteOrder(opts *bind.TransactOpts, order VERC20Order, taker common.Address) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "executeOrder", order, taker)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x87649e4f.
//
// Solidity: function executeOrder((address,bool,bytes32,string,uint256,uint256,uint64,uint64,uint8,bytes32,bytes32) order, address taker) payable returns()
func (_VERC20Market *VERC20MarketSession) ExecuteOrder(order VERC20Order, taker common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.ExecuteOrder(&_VERC20Market.TransactOpts, order, taker)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x87649e4f.
//
// Solidity: function executeOrder((address,bool,bytes32,string,uint256,uint256,uint64,uint64,uint8,bytes32,bytes32) order, address taker) payable returns()
func (_VERC20Market *VERC20MarketTransactorSession) ExecuteOrder(order VERC20Order, taker common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.ExecuteOrder(&_VERC20Market.TransactOpts, order, taker)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VERC20Market *VERC20MarketTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VERC20Market *VERC20MarketSession) Initialize() (*types.Transaction, error) {
	return _VERC20Market.Contract.Initialize(&_VERC20Market.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_VERC20Market *VERC20MarketTransactorSession) Initialize() (*types.Transaction, error) {
	return _VERC20Market.Contract.Initialize(&_VERC20Market.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VERC20Market *VERC20MarketTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VERC20Market *VERC20MarketSession) Pause() (*types.Transaction, error) {
	return _VERC20Market.Contract.Pause(&_VERC20Market.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VERC20Market *VERC20MarketTransactorSession) Pause() (*types.Transaction, error) {
	return _VERC20Market.Contract.Pause(&_VERC20Market.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VERC20Market *VERC20MarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VERC20Market *VERC20MarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _VERC20Market.Contract.RenounceOwnership(&_VERC20Market.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VERC20Market *VERC20MarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VERC20Market.Contract.RenounceOwnership(&_VERC20Market.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VERC20Market *VERC20MarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VERC20Market *VERC20MarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.TransferOwnership(&_VERC20Market.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VERC20Market *VERC20MarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.TransferOwnership(&_VERC20Market.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VERC20Market *VERC20MarketTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VERC20Market *VERC20MarketSession) Unpause() (*types.Transaction, error) {
	return _VERC20Market.Contract.Unpause(&_VERC20Market.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VERC20Market *VERC20MarketTransactorSession) Unpause() (*types.Transaction, error) {
	return _VERC20Market.Contract.Unpause(&_VERC20Market.TransactOpts)
}

// UpdateTrustedVerifier is a paid mutator transaction binding the contract method 0x4fa54246.
//
// Solidity: function updateTrustedVerifier(address _trustedVerifier) returns()
func (_VERC20Market *VERC20MarketTransactor) UpdateTrustedVerifier(opts *bind.TransactOpts, _trustedVerifier common.Address) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "updateTrustedVerifier", _trustedVerifier)
}

// UpdateTrustedVerifier is a paid mutator transaction binding the contract method 0x4fa54246.
//
// Solidity: function updateTrustedVerifier(address _trustedVerifier) returns()
func (_VERC20Market *VERC20MarketSession) UpdateTrustedVerifier(_trustedVerifier common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.UpdateTrustedVerifier(&_VERC20Market.TransactOpts, _trustedVerifier)
}

// UpdateTrustedVerifier is a paid mutator transaction binding the contract method 0x4fa54246.
//
// Solidity: function updateTrustedVerifier(address _trustedVerifier) returns()
func (_VERC20Market *VERC20MarketTransactorSession) UpdateTrustedVerifier(_trustedVerifier common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.UpdateTrustedVerifier(&_VERC20Market.TransactOpts, _trustedVerifier)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VERC20Market *VERC20MarketTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VERC20Market *VERC20MarketSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.UpgradeTo(&_VERC20Market.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VERC20Market *VERC20MarketTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VERC20Market.Contract.UpgradeTo(&_VERC20Market.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VERC20Market *VERC20MarketTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VERC20Market *VERC20MarketSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VERC20Market.Contract.UpgradeToAndCall(&_VERC20Market.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VERC20Market *VERC20MarketTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VERC20Market.Contract.UpgradeToAndCall(&_VERC20Market.TransactOpts, newImplementation, data)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_VERC20Market *VERC20MarketTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_VERC20Market *VERC20MarketSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VERC20Market.Contract.WithdrawETH(&_VERC20Market.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_VERC20Market *VERC20MarketTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VERC20Market.Contract.WithdrawETH(&_VERC20Market.TransactOpts, to, amount)
}

// WithdrawUnexpectedERC20 is a paid mutator transaction binding the contract method 0x19bc691a.
//
// Solidity: function withdrawUnexpectedERC20(address token, address to, uint256 amount) returns()
func (_VERC20Market *VERC20MarketTransactor) WithdrawUnexpectedERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VERC20Market.contract.Transact(opts, "withdrawUnexpectedERC20", token, to, amount)
}

// WithdrawUnexpectedERC20 is a paid mutator transaction binding the contract method 0x19bc691a.
//
// Solidity: function withdrawUnexpectedERC20(address token, address to, uint256 amount) returns()
func (_VERC20Market *VERC20MarketSession) WithdrawUnexpectedERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VERC20Market.Contract.WithdrawUnexpectedERC20(&_VERC20Market.TransactOpts, token, to, amount)
}

// WithdrawUnexpectedERC20 is a paid mutator transaction binding the contract method 0x19bc691a.
//
// Solidity: function withdrawUnexpectedERC20(address token, address to, uint256 amount) returns()
func (_VERC20Market *VERC20MarketTransactorSession) WithdrawUnexpectedERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VERC20Market.Contract.WithdrawUnexpectedERC20(&_VERC20Market.TransactOpts, token, to, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VERC20Market *VERC20MarketTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _VERC20Market.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VERC20Market *VERC20MarketSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VERC20Market.Contract.Fallback(&_VERC20Market.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VERC20Market *VERC20MarketTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VERC20Market.Contract.Fallback(&_VERC20Market.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VERC20Market *VERC20MarketTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VERC20Market.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VERC20Market *VERC20MarketSession) Receive() (*types.Transaction, error) {
	return _VERC20Market.Contract.Receive(&_VERC20Market.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VERC20Market *VERC20MarketTransactorSession) Receive() (*types.Transaction, error) {
	return _VERC20Market.Contract.Receive(&_VERC20Market.TransactOpts)
}

// VERC20MarketAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the VERC20Market contract.
type VERC20MarketAdminChangedIterator struct {
	Event *VERC20MarketAdminChanged // Event containing the contract specifics and raw log

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
func (it *VERC20MarketAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketAdminChanged)
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
		it.Event = new(VERC20MarketAdminChanged)
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
func (it *VERC20MarketAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketAdminChanged represents a AdminChanged event raised by the VERC20Market contract.
type VERC20MarketAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VERC20Market *VERC20MarketFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VERC20MarketAdminChangedIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketAdminChangedIterator{contract: _VERC20Market.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VERC20Market *VERC20MarketFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VERC20MarketAdminChanged) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketAdminChanged)
				if err := _VERC20Market.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VERC20Market *VERC20MarketFilterer) ParseAdminChanged(log types.Log) (*VERC20MarketAdminChanged, error) {
	event := new(VERC20MarketAdminChanged)
	if err := _VERC20Market.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the VERC20Market contract.
type VERC20MarketBeaconUpgradedIterator struct {
	Event *VERC20MarketBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VERC20MarketBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketBeaconUpgraded)
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
		it.Event = new(VERC20MarketBeaconUpgraded)
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
func (it *VERC20MarketBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketBeaconUpgraded represents a BeaconUpgraded event raised by the VERC20Market contract.
type VERC20MarketBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VERC20Market *VERC20MarketFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VERC20MarketBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketBeaconUpgradedIterator{contract: _VERC20Market.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VERC20Market *VERC20MarketFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VERC20MarketBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketBeaconUpgraded)
				if err := _VERC20Market.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VERC20Market *VERC20MarketFilterer) ParseBeaconUpgraded(log types.Log) (*VERC20MarketBeaconUpgraded, error) {
	event := new(VERC20MarketBeaconUpgraded)
	if err := _VERC20Market.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the VERC20Market contract.
type VERC20MarketEIP712DomainChangedIterator struct {
	Event *VERC20MarketEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *VERC20MarketEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketEIP712DomainChanged)
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
		it.Event = new(VERC20MarketEIP712DomainChanged)
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
func (it *VERC20MarketEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketEIP712DomainChanged represents a EIP712DomainChanged event raised by the VERC20Market contract.
type VERC20MarketEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VERC20Market *VERC20MarketFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*VERC20MarketEIP712DomainChangedIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketEIP712DomainChangedIterator{contract: _VERC20Market.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VERC20Market *VERC20MarketFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *VERC20MarketEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketEIP712DomainChanged)
				if err := _VERC20Market.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VERC20Market *VERC20MarketFilterer) ParseEIP712DomainChanged(log types.Log) (*VERC20MarketEIP712DomainChanged, error) {
	event := new(VERC20MarketEIP712DomainChanged)
	if err := _VERC20Market.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VERC20Market contract.
type VERC20MarketInitializedIterator struct {
	Event *VERC20MarketInitialized // Event containing the contract specifics and raw log

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
func (it *VERC20MarketInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketInitialized)
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
		it.Event = new(VERC20MarketInitialized)
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
func (it *VERC20MarketInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketInitialized represents a Initialized event raised by the VERC20Market contract.
type VERC20MarketInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VERC20Market *VERC20MarketFilterer) FilterInitialized(opts *bind.FilterOpts) (*VERC20MarketInitializedIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketInitializedIterator{contract: _VERC20Market.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VERC20Market *VERC20MarketFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VERC20MarketInitialized) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketInitialized)
				if err := _VERC20Market.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VERC20Market *VERC20MarketFilterer) ParseInitialized(log types.Log) (*VERC20MarketInitialized, error) {
	event := new(VERC20MarketInitialized)
	if err := _VERC20Market.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketNewTrustedVerifierIterator is returned from FilterNewTrustedVerifier and is used to iterate over the raw logs and unpacked data for NewTrustedVerifier events raised by the VERC20Market contract.
type VERC20MarketNewTrustedVerifierIterator struct {
	Event *VERC20MarketNewTrustedVerifier // Event containing the contract specifics and raw log

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
func (it *VERC20MarketNewTrustedVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketNewTrustedVerifier)
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
		it.Event = new(VERC20MarketNewTrustedVerifier)
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
func (it *VERC20MarketNewTrustedVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketNewTrustedVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketNewTrustedVerifier represents a NewTrustedVerifier event raised by the VERC20Market contract.
type VERC20MarketNewTrustedVerifier struct {
	TrustedVerifier common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewTrustedVerifier is a free log retrieval operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_VERC20Market *VERC20MarketFilterer) FilterNewTrustedVerifier(opts *bind.FilterOpts) (*VERC20MarketNewTrustedVerifierIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "NewTrustedVerifier")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketNewTrustedVerifierIterator{contract: _VERC20Market.contract, event: "NewTrustedVerifier", logs: logs, sub: sub}, nil
}

// WatchNewTrustedVerifier is a free log subscription operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_VERC20Market *VERC20MarketFilterer) WatchNewTrustedVerifier(opts *bind.WatchOpts, sink chan<- *VERC20MarketNewTrustedVerifier) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "NewTrustedVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketNewTrustedVerifier)
				if err := _VERC20Market.contract.UnpackLog(event, "NewTrustedVerifier", log); err != nil {
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

// ParseNewTrustedVerifier is a log parse operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_VERC20Market *VERC20MarketFilterer) ParseNewTrustedVerifier(log types.Log) (*VERC20MarketNewTrustedVerifier, error) {
	event := new(VERC20MarketNewTrustedVerifier)
	if err := _VERC20Market.contract.UnpackLog(event, "NewTrustedVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VERC20Market contract.
type VERC20MarketOwnershipTransferredIterator struct {
	Event *VERC20MarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VERC20MarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketOwnershipTransferred)
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
		it.Event = new(VERC20MarketOwnershipTransferred)
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
func (it *VERC20MarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketOwnershipTransferred represents a OwnershipTransferred event raised by the VERC20Market contract.
type VERC20MarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VERC20Market *VERC20MarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VERC20MarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketOwnershipTransferredIterator{contract: _VERC20Market.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VERC20Market *VERC20MarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VERC20MarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketOwnershipTransferred)
				if err := _VERC20Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VERC20Market *VERC20MarketFilterer) ParseOwnershipTransferred(log types.Log) (*VERC20MarketOwnershipTransferred, error) {
	event := new(VERC20MarketOwnershipTransferred)
	if err := _VERC20Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VERC20Market contract.
type VERC20MarketPausedIterator struct {
	Event *VERC20MarketPaused // Event containing the contract specifics and raw log

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
func (it *VERC20MarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketPaused)
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
		it.Event = new(VERC20MarketPaused)
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
func (it *VERC20MarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketPaused represents a Paused event raised by the VERC20Market contract.
type VERC20MarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VERC20Market *VERC20MarketFilterer) FilterPaused(opts *bind.FilterOpts) (*VERC20MarketPausedIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketPausedIterator{contract: _VERC20Market.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VERC20Market *VERC20MarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VERC20MarketPaused) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketPaused)
				if err := _VERC20Market.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VERC20Market *VERC20MarketFilterer) ParsePaused(log types.Log) (*VERC20MarketPaused, error) {
	event := new(VERC20MarketPaused)
	if err := _VERC20Market.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VERC20Market contract.
type VERC20MarketUnpausedIterator struct {
	Event *VERC20MarketUnpaused // Event containing the contract specifics and raw log

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
func (it *VERC20MarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketUnpaused)
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
		it.Event = new(VERC20MarketUnpaused)
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
func (it *VERC20MarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketUnpaused represents a Unpaused event raised by the VERC20Market contract.
type VERC20MarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VERC20Market *VERC20MarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VERC20MarketUnpausedIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketUnpausedIterator{contract: _VERC20Market.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VERC20Market *VERC20MarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VERC20MarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketUnpaused)
				if err := _VERC20Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VERC20Market *VERC20MarketFilterer) ParseUnpaused(log types.Log) (*VERC20MarketUnpaused, error) {
	event := new(VERC20MarketUnpaused)
	if err := _VERC20Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VERC20Market contract.
type VERC20MarketUpgradedIterator struct {
	Event *VERC20MarketUpgraded // Event containing the contract specifics and raw log

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
func (it *VERC20MarketUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketUpgraded)
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
		it.Event = new(VERC20MarketUpgraded)
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
func (it *VERC20MarketUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketUpgraded represents a Upgraded event raised by the VERC20Market contract.
type VERC20MarketUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VERC20Market *VERC20MarketFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VERC20MarketUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketUpgradedIterator{contract: _VERC20Market.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VERC20Market *VERC20MarketFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VERC20MarketUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketUpgraded)
				if err := _VERC20Market.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VERC20Market *VERC20MarketFilterer) ParseUpgraded(log types.Log) (*VERC20MarketUpgraded, error) {
	event := new(VERC20MarketUpgraded)
	if err := _VERC20Market.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketVERC20OrderCanceledIterator is returned from FilterVERC20OrderCanceled and is used to iterate over the raw logs and unpacked data for VERC20OrderCanceled events raised by the VERC20Market contract.
type VERC20MarketVERC20OrderCanceledIterator struct {
	Event *VERC20MarketVERC20OrderCanceled // Event containing the contract specifics and raw log

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
func (it *VERC20MarketVERC20OrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketVERC20OrderCanceled)
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
		it.Event = new(VERC20MarketVERC20OrderCanceled)
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
func (it *VERC20MarketVERC20OrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketVERC20OrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketVERC20OrderCanceled represents a VERC20OrderCanceled event raised by the VERC20Market contract.
type VERC20MarketVERC20OrderCanceled struct {
	Maker     common.Address
	Sell      bool
	ListId    [32]byte
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVERC20OrderCanceled is a free log retrieval operation binding the contract event 0x374c9bc2fad2e993accd6456fa2a40f1d10d7deb18d59f0d0748f309abe77005.
//
// Solidity: event VERC20OrderCanceled(address maker, bool sell, bytes32 listId, uint64 timestamp)
func (_VERC20Market *VERC20MarketFilterer) FilterVERC20OrderCanceled(opts *bind.FilterOpts) (*VERC20MarketVERC20OrderCanceledIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "VERC20OrderCanceled")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketVERC20OrderCanceledIterator{contract: _VERC20Market.contract, event: "VERC20OrderCanceled", logs: logs, sub: sub}, nil
}

// WatchVERC20OrderCanceled is a free log subscription operation binding the contract event 0x374c9bc2fad2e993accd6456fa2a40f1d10d7deb18d59f0d0748f309abe77005.
//
// Solidity: event VERC20OrderCanceled(address maker, bool sell, bytes32 listId, uint64 timestamp)
func (_VERC20Market *VERC20MarketFilterer) WatchVERC20OrderCanceled(opts *bind.WatchOpts, sink chan<- *VERC20MarketVERC20OrderCanceled) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "VERC20OrderCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketVERC20OrderCanceled)
				if err := _VERC20Market.contract.UnpackLog(event, "VERC20OrderCanceled", log); err != nil {
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

// ParseVERC20OrderCanceled is a log parse operation binding the contract event 0x374c9bc2fad2e993accd6456fa2a40f1d10d7deb18d59f0d0748f309abe77005.
//
// Solidity: event VERC20OrderCanceled(address maker, bool sell, bytes32 listId, uint64 timestamp)
func (_VERC20Market *VERC20MarketFilterer) ParseVERC20OrderCanceled(log types.Log) (*VERC20MarketVERC20OrderCanceled, error) {
	event := new(VERC20MarketVERC20OrderCanceled)
	if err := _VERC20Market.contract.UnpackLog(event, "VERC20OrderCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketVERC20OrderExecutedIterator is returned from FilterVERC20OrderExecuted and is used to iterate over the raw logs and unpacked data for VERC20OrderExecuted events raised by the VERC20Market contract.
type VERC20MarketVERC20OrderExecutedIterator struct {
	Event *VERC20MarketVERC20OrderExecuted // Event containing the contract specifics and raw log

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
func (it *VERC20MarketVERC20OrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketVERC20OrderExecuted)
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
		it.Event = new(VERC20MarketVERC20OrderExecuted)
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
func (it *VERC20MarketVERC20OrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketVERC20OrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketVERC20OrderExecuted represents a VERC20OrderExecuted event raised by the VERC20Market contract.
type VERC20MarketVERC20OrderExecuted struct {
	Maker     common.Address
	Sell      bool
	Taker     common.Address
	ListId    [32]byte
	Tick      string
	Amount    *big.Int
	Price     *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVERC20OrderExecuted is a free log retrieval operation binding the contract event 0x2e7b94f4cfd0f01fb59679cca0548850c180ac26c51412d95cff274ee38ba00d.
//
// Solidity: event VERC20OrderExecuted(address maker, bool sell, address taker, bytes32 listId, string tick, uint256 amount, uint256 price, uint64 timestamp)
func (_VERC20Market *VERC20MarketFilterer) FilterVERC20OrderExecuted(opts *bind.FilterOpts) (*VERC20MarketVERC20OrderExecutedIterator, error) {

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "VERC20OrderExecuted")
	if err != nil {
		return nil, err
	}
	return &VERC20MarketVERC20OrderExecutedIterator{contract: _VERC20Market.contract, event: "VERC20OrderExecuted", logs: logs, sub: sub}, nil
}

// WatchVERC20OrderExecuted is a free log subscription operation binding the contract event 0x2e7b94f4cfd0f01fb59679cca0548850c180ac26c51412d95cff274ee38ba00d.
//
// Solidity: event VERC20OrderExecuted(address maker, bool sell, address taker, bytes32 listId, string tick, uint256 amount, uint256 price, uint64 timestamp)
func (_VERC20Market *VERC20MarketFilterer) WatchVERC20OrderExecuted(opts *bind.WatchOpts, sink chan<- *VERC20MarketVERC20OrderExecuted) (event.Subscription, error) {

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "VERC20OrderExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketVERC20OrderExecuted)
				if err := _VERC20Market.contract.UnpackLog(event, "VERC20OrderExecuted", log); err != nil {
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

// ParseVERC20OrderExecuted is a log parse operation binding the contract event 0x2e7b94f4cfd0f01fb59679cca0548850c180ac26c51412d95cff274ee38ba00d.
//
// Solidity: event VERC20OrderExecuted(address maker, bool sell, address taker, bytes32 listId, string tick, uint256 amount, uint256 price, uint64 timestamp)
func (_VERC20Market *VERC20MarketFilterer) ParseVERC20OrderExecuted(log types.Log) (*VERC20MarketVERC20OrderExecuted, error) {
	event := new(VERC20MarketVERC20OrderExecuted)
	if err := _VERC20Market.contract.UnpackLog(event, "VERC20OrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketVerc20ProtocolTransferVERC20TokenIterator is returned from FilterVerc20ProtocolTransferVERC20Token and is used to iterate over the raw logs and unpacked data for Verc20ProtocolTransferVERC20Token events raised by the VERC20Market contract.
type VERC20MarketVerc20ProtocolTransferVERC20TokenIterator struct {
	Event *VERC20MarketVerc20ProtocolTransferVERC20Token // Event containing the contract specifics and raw log

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
func (it *VERC20MarketVerc20ProtocolTransferVERC20TokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketVerc20ProtocolTransferVERC20Token)
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
		it.Event = new(VERC20MarketVerc20ProtocolTransferVERC20Token)
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
func (it *VERC20MarketVerc20ProtocolTransferVERC20TokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketVerc20ProtocolTransferVERC20TokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketVerc20ProtocolTransferVERC20Token represents a Verc20ProtocolTransferVERC20Token event raised by the VERC20Market contract.
type VERC20MarketVerc20ProtocolTransferVERC20Token struct {
	From   common.Address
	To     common.Address
	Ticker common.Hash
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVerc20ProtocolTransferVERC20Token is a free log retrieval operation binding the contract event 0xe5b2cc322ab41152b0763a85456f7e62ab3d424fb17418e8f6f00f3f7f40e5d0.
//
// Solidity: event verc20_protocol_TransferVERC20Token(address indexed from, address indexed to, string indexed ticker, uint256 amount)
func (_VERC20Market *VERC20MarketFilterer) FilterVerc20ProtocolTransferVERC20Token(opts *bind.FilterOpts, from []common.Address, to []common.Address, ticker []string) (*VERC20MarketVerc20ProtocolTransferVERC20TokenIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "verc20_protocol_TransferVERC20Token", fromRule, toRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketVerc20ProtocolTransferVERC20TokenIterator{contract: _VERC20Market.contract, event: "verc20_protocol_TransferVERC20Token", logs: logs, sub: sub}, nil
}

// WatchVerc20ProtocolTransferVERC20Token is a free log subscription operation binding the contract event 0xe5b2cc322ab41152b0763a85456f7e62ab3d424fb17418e8f6f00f3f7f40e5d0.
//
// Solidity: event verc20_protocol_TransferVERC20Token(address indexed from, address indexed to, string indexed ticker, uint256 amount)
func (_VERC20Market *VERC20MarketFilterer) WatchVerc20ProtocolTransferVERC20Token(opts *bind.WatchOpts, sink chan<- *VERC20MarketVerc20ProtocolTransferVERC20Token, from []common.Address, to []common.Address, ticker []string) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "verc20_protocol_TransferVERC20Token", fromRule, toRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketVerc20ProtocolTransferVERC20Token)
				if err := _VERC20Market.contract.UnpackLog(event, "verc20_protocol_TransferVERC20Token", log); err != nil {
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

// ParseVerc20ProtocolTransferVERC20Token is a log parse operation binding the contract event 0xe5b2cc322ab41152b0763a85456f7e62ab3d424fb17418e8f6f00f3f7f40e5d0.
//
// Solidity: event verc20_protocol_TransferVERC20Token(address indexed from, address indexed to, string indexed ticker, uint256 amount)
func (_VERC20Market *VERC20MarketFilterer) ParseVerc20ProtocolTransferVERC20Token(log types.Log) (*VERC20MarketVerc20ProtocolTransferVERC20Token, error) {
	event := new(VERC20MarketVerc20ProtocolTransferVERC20Token)
	if err := _VERC20Market.contract.UnpackLog(event, "verc20_protocol_TransferVERC20Token", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VERC20MarketVerc20ProtocolTransferVERC20TokenForListingIterator is returned from FilterVerc20ProtocolTransferVERC20TokenForListing and is used to iterate over the raw logs and unpacked data for Verc20ProtocolTransferVERC20TokenForListing events raised by the VERC20Market contract.
type VERC20MarketVerc20ProtocolTransferVERC20TokenForListingIterator struct {
	Event *VERC20MarketVerc20ProtocolTransferVERC20TokenForListing // Event containing the contract specifics and raw log

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
func (it *VERC20MarketVerc20ProtocolTransferVERC20TokenForListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VERC20MarketVerc20ProtocolTransferVERC20TokenForListing)
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
		it.Event = new(VERC20MarketVerc20ProtocolTransferVERC20TokenForListing)
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
func (it *VERC20MarketVerc20ProtocolTransferVERC20TokenForListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VERC20MarketVerc20ProtocolTransferVERC20TokenForListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VERC20MarketVerc20ProtocolTransferVERC20TokenForListing represents a Verc20ProtocolTransferVERC20TokenForListing event raised by the VERC20Market contract.
type VERC20MarketVerc20ProtocolTransferVERC20TokenForListing struct {
	From common.Address
	To   common.Address
	Id   [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterVerc20ProtocolTransferVERC20TokenForListing is a free log retrieval operation binding the contract event 0xd88ec1d577aa21de902dbe19564239291fca3b0630aa716cccb858da95a5caec.
//
// Solidity: event verc20_protocol_TransferVERC20TokenForListing(address indexed from, address indexed to, bytes32 id)
func (_VERC20Market *VERC20MarketFilterer) FilterVerc20ProtocolTransferVERC20TokenForListing(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VERC20MarketVerc20ProtocolTransferVERC20TokenForListingIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VERC20Market.contract.FilterLogs(opts, "verc20_protocol_TransferVERC20TokenForListing", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VERC20MarketVerc20ProtocolTransferVERC20TokenForListingIterator{contract: _VERC20Market.contract, event: "verc20_protocol_TransferVERC20TokenForListing", logs: logs, sub: sub}, nil
}

// WatchVerc20ProtocolTransferVERC20TokenForListing is a free log subscription operation binding the contract event 0xd88ec1d577aa21de902dbe19564239291fca3b0630aa716cccb858da95a5caec.
//
// Solidity: event verc20_protocol_TransferVERC20TokenForListing(address indexed from, address indexed to, bytes32 id)
func (_VERC20Market *VERC20MarketFilterer) WatchVerc20ProtocolTransferVERC20TokenForListing(opts *bind.WatchOpts, sink chan<- *VERC20MarketVerc20ProtocolTransferVERC20TokenForListing, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VERC20Market.contract.WatchLogs(opts, "verc20_protocol_TransferVERC20TokenForListing", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VERC20MarketVerc20ProtocolTransferVERC20TokenForListing)
				if err := _VERC20Market.contract.UnpackLog(event, "verc20_protocol_TransferVERC20TokenForListing", log); err != nil {
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

// ParseVerc20ProtocolTransferVERC20TokenForListing is a log parse operation binding the contract event 0xd88ec1d577aa21de902dbe19564239291fca3b0630aa716cccb858da95a5caec.
//
// Solidity: event verc20_protocol_TransferVERC20TokenForListing(address indexed from, address indexed to, bytes32 id)
func (_VERC20Market *VERC20MarketFilterer) ParseVerc20ProtocolTransferVERC20TokenForListing(log types.Log) (*VERC20MarketVerc20ProtocolTransferVERC20TokenForListing, error) {
	event := new(VERC20MarketVerc20ProtocolTransferVERC20TokenForListing)
	if err := _VERC20Market.contract.UnpackLog(event, "verc20_protocol_TransferVERC20TokenForListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
