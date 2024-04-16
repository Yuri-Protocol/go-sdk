// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc_6551

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC6551AccountABI is the input ABI used to generate the binding from.
const ERC6551AccountABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isValidSigner\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC6551Account is an auto generated Go binding around an Ethereum contract.
type ERC6551Account struct {
	ERC6551AccountCaller     // Read-only binding to the contract
	ERC6551AccountTransactor // Write-only binding to the contract
	ERC6551AccountFilterer   // Log filterer for contract events
}

// ERC6551AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC6551AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC6551AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC6551AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC6551AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC6551AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC6551AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC6551AccountSession struct {
	Contract     *ERC6551Account   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC6551AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC6551AccountCallerSession struct {
	Contract *ERC6551AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC6551AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC6551AccountTransactorSession struct {
	Contract     *ERC6551AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC6551AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC6551AccountRaw struct {
	Contract *ERC6551Account // Generic contract binding to access the raw methods on
}

// ERC6551AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC6551AccountCallerRaw struct {
	Contract *ERC6551AccountCaller // Generic read-only contract binding to access the raw methods on
}

// ERC6551AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC6551AccountTransactorRaw struct {
	Contract *ERC6551AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC6551Account creates a new instance of ERC6551Account, bound to a specific deployed contract.
func NewERC6551Account(address common.Address, backend bind.ContractBackend) (*ERC6551Account, error) {
	contract, err := bindERC6551Account(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC6551Account{ERC6551AccountCaller: ERC6551AccountCaller{contract: contract}, ERC6551AccountTransactor: ERC6551AccountTransactor{contract: contract}, ERC6551AccountFilterer: ERC6551AccountFilterer{contract: contract}}, nil
}

// NewERC6551AccountCaller creates a new read-only instance of ERC6551Account, bound to a specific deployed contract.
func NewERC6551AccountCaller(address common.Address, caller bind.ContractCaller) (*ERC6551AccountCaller, error) {
	contract, err := bindERC6551Account(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC6551AccountCaller{contract: contract}, nil
}

// NewERC6551AccountTransactor creates a new write-only instance of ERC6551Account, bound to a specific deployed contract.
func NewERC6551AccountTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC6551AccountTransactor, error) {
	contract, err := bindERC6551Account(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC6551AccountTransactor{contract: contract}, nil
}

// NewERC6551AccountFilterer creates a new log filterer instance of ERC6551Account, bound to a specific deployed contract.
func NewERC6551AccountFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC6551AccountFilterer, error) {
	contract, err := bindERC6551Account(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC6551AccountFilterer{contract: contract}, nil
}

// bindERC6551Account binds a generic wrapper to an already deployed contract.
func bindERC6551Account(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC6551AccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC6551Account *ERC6551AccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC6551Account.Contract.ERC6551AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC6551Account *ERC6551AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC6551Account.Contract.ERC6551AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC6551Account *ERC6551AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC6551Account.Contract.ERC6551AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC6551Account *ERC6551AccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC6551Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC6551Account *ERC6551AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC6551Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC6551Account *ERC6551AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC6551Account.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ERC6551Account *ERC6551AccountCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ERC6551Account.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ERC6551Account *ERC6551AccountSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _ERC6551Account.Contract.IsValidSignature(&_ERC6551Account.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ERC6551Account *ERC6551AccountCallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _ERC6551Account.Contract.IsValidSignature(&_ERC6551Account.CallOpts, hash, signature)
}

// IsValidSigner is a free data retrieval call binding the contract method 0x523e3260.
//
// Solidity: function isValidSigner(address signer, bytes ) view returns(bytes4)
func (_ERC6551Account *ERC6551AccountCaller) IsValidSigner(opts *bind.CallOpts, signer common.Address, arg1 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ERC6551Account.contract.Call(opts, &out, "isValidSigner", signer, arg1)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSigner is a free data retrieval call binding the contract method 0x523e3260.
//
// Solidity: function isValidSigner(address signer, bytes ) view returns(bytes4)
func (_ERC6551Account *ERC6551AccountSession) IsValidSigner(signer common.Address, arg1 []byte) ([4]byte, error) {
	return _ERC6551Account.Contract.IsValidSigner(&_ERC6551Account.CallOpts, signer, arg1)
}

// IsValidSigner is a free data retrieval call binding the contract method 0x523e3260.
//
// Solidity: function isValidSigner(address signer, bytes ) view returns(bytes4)
func (_ERC6551Account *ERC6551AccountCallerSession) IsValidSigner(signer common.Address, arg1 []byte) ([4]byte, error) {
	return _ERC6551Account.Contract.IsValidSigner(&_ERC6551Account.CallOpts, signer, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC6551Account *ERC6551AccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC6551Account.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC6551Account *ERC6551AccountSession) Owner() (common.Address, error) {
	return _ERC6551Account.Contract.Owner(&_ERC6551Account.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC6551Account *ERC6551AccountCallerSession) Owner() (common.Address, error) {
	return _ERC6551Account.Contract.Owner(&_ERC6551Account.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256)
func (_ERC6551Account *ERC6551AccountCaller) State(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC6551Account.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256)
func (_ERC6551Account *ERC6551AccountSession) State() (*big.Int, error) {
	return _ERC6551Account.Contract.State(&_ERC6551Account.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256)
func (_ERC6551Account *ERC6551AccountCallerSession) State() (*big.Int, error) {
	return _ERC6551Account.Contract.State(&_ERC6551Account.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ERC6551Account *ERC6551AccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC6551Account.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ERC6551Account *ERC6551AccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC6551Account.Contract.SupportsInterface(&_ERC6551Account.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ERC6551Account *ERC6551AccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC6551Account.Contract.SupportsInterface(&_ERC6551Account.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(uint256, address, uint256)
func (_ERC6551Account *ERC6551AccountCaller) Token(opts *bind.CallOpts) (*big.Int, common.Address, *big.Int, error) {
	var out []interface{}
	err := _ERC6551Account.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(uint256, address, uint256)
func (_ERC6551Account *ERC6551AccountSession) Token() (*big.Int, common.Address, *big.Int, error) {
	return _ERC6551Account.Contract.Token(&_ERC6551Account.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(uint256, address, uint256)
func (_ERC6551Account *ERC6551AccountCallerSession) Token() (*big.Int, common.Address, *big.Int, error) {
	return _ERC6551Account.Contract.Token(&_ERC6551Account.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x51945447.
//
// Solidity: function execute(address to, uint256 value, bytes data, uint8 operation) payable returns(bytes result)
func (_ERC6551Account *ERC6551AccountTransactor) Execute(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ERC6551Account.contract.Transact(opts, "execute", to, value, data, operation)
}

// Execute is a paid mutator transaction binding the contract method 0x51945447.
//
// Solidity: function execute(address to, uint256 value, bytes data, uint8 operation) payable returns(bytes result)
func (_ERC6551Account *ERC6551AccountSession) Execute(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ERC6551Account.Contract.Execute(&_ERC6551Account.TransactOpts, to, value, data, operation)
}

// Execute is a paid mutator transaction binding the contract method 0x51945447.
//
// Solidity: function execute(address to, uint256 value, bytes data, uint8 operation) payable returns(bytes result)
func (_ERC6551Account *ERC6551AccountTransactorSession) Execute(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ERC6551Account.Contract.Execute(&_ERC6551Account.TransactOpts, to, value, data, operation)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address token, address to, uint256 amount) returns()
func (_ERC6551Account *ERC6551AccountTransactor) TransferToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC6551Account.contract.Transact(opts, "transferToken", token, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address token, address to, uint256 amount) returns()
func (_ERC6551Account *ERC6551AccountSession) TransferToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC6551Account.Contract.TransferToken(&_ERC6551Account.TransactOpts, token, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf5537ede.
//
// Solidity: function transferToken(address token, address to, uint256 amount) returns()
func (_ERC6551Account *ERC6551AccountTransactorSession) TransferToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC6551Account.Contract.TransferToken(&_ERC6551Account.TransactOpts, token, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC6551Account *ERC6551AccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC6551Account.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC6551Account *ERC6551AccountSession) Receive() (*types.Transaction, error) {
	return _ERC6551Account.Contract.Receive(&_ERC6551Account.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC6551Account *ERC6551AccountTransactorSession) Receive() (*types.Transaction, error) {
	return _ERC6551Account.Contract.Receive(&_ERC6551Account.TransactOpts)
}
