// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pbridge

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

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058206967258b5867a82b39700ce1b3b59658fa0ddaa5f67e7e2f6ebe7d7d9963dcd90029"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820ff5282fdd09a603052ff2ab78ac2a90ffb8a790942de03244d8c59f54c5a72e60029"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MinterABI is the input ABI used to generate the binding from.
const IERC20MinterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"replaceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20MinterFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MinterFuncSigs = map[string]string{
	"42966c68": "burn(uint256)",
	"40c10f19": "mint(address,uint256)",
	"07f1af44": "replaceMinter(address)",
}

// IERC20Minter is an auto generated Go binding around an Ethereum contract.
type IERC20Minter struct {
	IERC20MinterCaller     // Read-only binding to the contract
	IERC20MinterTransactor // Write-only binding to the contract
	IERC20MinterFilterer   // Log filterer for contract events
}

// IERC20MinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MinterSession struct {
	Contract     *IERC20Minter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MinterCallerSession struct {
	Contract *IERC20MinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20MinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MinterTransactorSession struct {
	Contract     *IERC20MinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20MinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MinterRaw struct {
	Contract *IERC20Minter // Generic contract binding to access the raw methods on
}

// IERC20MinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MinterCallerRaw struct {
	Contract *IERC20MinterCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MinterTransactorRaw struct {
	Contract *IERC20MinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Minter creates a new instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20Minter(address common.Address, backend bind.ContractBackend) (*IERC20Minter, error) {
	contract, err := bindIERC20Minter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Minter{IERC20MinterCaller: IERC20MinterCaller{contract: contract}, IERC20MinterTransactor: IERC20MinterTransactor{contract: contract}, IERC20MinterFilterer: IERC20MinterFilterer{contract: contract}}, nil
}

// NewIERC20MinterCaller creates a new read-only instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20MinterCaller(address common.Address, caller bind.ContractCaller) (*IERC20MinterCaller, error) {
	contract, err := bindIERC20Minter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MinterCaller{contract: contract}, nil
}

// NewIERC20MinterTransactor creates a new write-only instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20MinterTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MinterTransactor, error) {
	contract, err := bindIERC20Minter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MinterTransactor{contract: contract}, nil
}

// NewIERC20MinterFilterer creates a new log filterer instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20MinterFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MinterFilterer, error) {
	contract, err := bindIERC20Minter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MinterFilterer{contract: contract}, nil
}

// bindIERC20Minter binds a generic wrapper to an already deployed contract.
func bindIERC20Minter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Minter *IERC20MinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Minter.Contract.IERC20MinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Minter *IERC20MinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Minter.Contract.IERC20MinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Minter *IERC20MinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Minter.Contract.IERC20MinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Minter *IERC20MinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Minter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Minter *IERC20MinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Minter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Minter *IERC20MinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Minter.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IERC20Minter *IERC20MinterSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Burn(&_IERC20Minter.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Burn(&_IERC20Minter.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Minter *IERC20MinterSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Mint(&_IERC20Minter.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Mint(&_IERC20Minter.TransactOpts, to, amount)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC20Minter *IERC20MinterTransactor) ReplaceMinter(opts *bind.TransactOpts, newMinter common.Address) (*types.Transaction, error) {
	return _IERC20Minter.contract.Transact(opts, "replaceMinter", newMinter)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC20Minter *IERC20MinterSession) ReplaceMinter(newMinter common.Address) (*types.Transaction, error) {
	return _IERC20Minter.Contract.ReplaceMinter(&_IERC20Minter.TransactOpts, newMinter)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC20Minter *IERC20MinterTransactorSession) ReplaceMinter(newMinter common.Address) (*types.Transaction, error) {
	return _IERC20Minter.Contract.ReplaceMinter(&_IERC20Minter.TransactOpts, newMinter)
}

// PBridgeABI is the input ABI used to generate the binding from.
const PBridgeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"txKey\",\"type\":\"string\"},{\"name\":\"adds\",\"type\":\"address[]\"},{\"name\":\"removes\",\"type\":\"address[]\"},{\"name\":\"count\",\"type\":\"uint8\"},{\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignManagerChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"crossOut\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signatureLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgradeContractS1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"registerMinterERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txKey\",\"type\":\"string\"},{\"name\":\"upgradeContract\",\"type\":\"address\"},{\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"upgradeContractS2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"isMinterERC20\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"ifManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allManagers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"unregisterMinterERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"isCompletedTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txKey\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"isERC20\",\"type\":\"bool\"},{\"name\":\"ERC20\",\"type\":\"address\"},{\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"min_managers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"closeUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_min_signatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"max_managers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_managers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"CrossOutFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxWithdrawCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxManagerChangeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxUpgradeCompleted\",\"type\":\"event\"}]"

// PBridgeFuncSigs maps the 4-byte function signature to its string representation.
var PBridgeFuncSigs = map[string]string{
	"9c30b35e": "allManagers()",
	"d4cacbaa": "closeUpgrade()",
	"00719226": "createOrSignManagerChange(string,address[],address[],uint8,bytes)",
	"408e8b7a": "createOrSignUpgrade(string,address,bytes)",
	"ab6c2b10": "createOrSignWithdraw(string,address,uint256,bool,address,bytes)",
	"0889d1f0": "crossOut(string,uint256,address)",
	"e079cee9": "current_min_signatures()",
	"75173b70": "ifManager(address)",
	"a5e399b3": "isCompletedTx(string)",
	"6a7142e1": "isMinterERC20(address)",
	"f7f2ff74": "max_managers()",
	"ad4b61a8": "min_managers()",
	"8da5cb5b": "owner()",
	"2c4e722e": "rate()",
	"34774b71": "registerMinterERC20(address)",
	"1b9a9323": "signatureLength()",
	"9dcdc978": "unregisterMinterERC20(address)",
	"d55ec697": "upgrade()",
	"30b2d84d": "upgradeContractAddress()",
	"1dda9c05": "upgradeContractS1()",
	"5bda3fcf": "upgradeContractS2(address)",
}

// PBridgeBin is the compiled bytecode used for deploying new contracts.
var PBridgeBin = "0x6080604052600080546001600160a81b0319169055600f600155600360028190556042905560416004553480156200003657600080fd5b506040516200400038038062004000833981018060405260208110156200005c57600080fd5b8101908080516401000000008111156200007557600080fd5b820160208101848111156200008957600080fd5b8151856020820283011164010000000082111715620000a757600080fd5b50509291905050506001548151111515156200010f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018062003fb36027913960400191505060405180910390fd5b600254815110156200016d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018062003f8c6027913960400191505060405180910390fd5b60058054610100600160a81b0319163361010002179055805162000199906009906020840190620003ca565b5060005b60095460ff82161015620002995760016008600060098460ff16815481101515620001c457fe5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff191660ff9384161790556009805460019360069392919086169081106200021057fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191660ff9283161790556009805460079284169081106200025657fe5b6000918252602080832090910154835460018181018655948452919092200180546001600160a01b0319166001600160a01b03909216919091179055016200019d565b5060055461010090046001600160a01b031660009081526008602052604090205460ff161562000315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018062003fda6026913960400191505060405180910390fd5b6200032b6009805490506200034660201b60201c565b6005805460ff191660ff92909216919091179055506200045e565b6000811515620003b757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f4d616e616765722043616e277420656d7074792e000000000000000000000000604482015290519081900360640190fd5b6003548202606301606481049392505050565b82805482825590600052602060002090810192821562000422579160200282015b828111156200042257825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003eb565b506200043092915062000434565b5090565b6200045b91905b80821115620004305780546001600160a01b03191681556001016200043b565b90565b613b1e806200046e6000396000f3fe6080604052600436106101295760003560e01c806375173b70116100ab578063ab6c2b101161006f578063ab6c2b1014610865578063ad4b61a8146109c1578063d4cacbaa146109d6578063d55ec697146109eb578063e079cee914610a00578063f7f2ff7414610a2b57610129565b806375173b70146106d45780638da5cb5b146107075780639c30b35e1461071c5780639dcdc97814610781578063a5e399b3146107b457610129565b806330b2d84d116100f257806330b2d84d146104c357806334774b71146104f4578063408e8b7a146105275780635bda3fcf1461066e5780636a7142e1146106a157610129565b8062719226146101655780630889d1f0146103ac5780631b9a9323146104725780631dda9c05146104995780632c4e722e146104ae575b6040805133815234602082015281517fd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88929181900390910190a1005b34801561017157600080fd5b506103aa600480360360a081101561018857600080fd5b810190602081018135600160201b8111156101a257600080fd5b8201836020820111156101b457600080fd5b803590602001918460018302840111600160201b831117156101d557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561022757600080fd5b82018360208201111561023957600080fd5b803590602001918460208302840111600160201b8311171561025a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102a957600080fd5b8201836020820111156102bb57600080fd5b803590602001918460208302840111600160201b831117156102dc57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929560ff853516959094909350604081019250602001359050600160201b81111561033657600080fd5b82018360208201111561034857600080fd5b803590602001918460018302840111600160201b8311171561036957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a40945050505050565b005b61045e600480360360608110156103c257600080fd5b810190602081018135600160201b8111156103dc57600080fd5b8201836020820111156103ee57600080fd5b803590602001918460018302840111600160201b8311171561040f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356001600160a01b0316610ea5565b604080519115158252519081900360200190f35b34801561047e57600080fd5b50610487611313565b60408051918252519081900360200190f35b3480156104a557600080fd5b506103aa611319565b3480156104ba57600080fd5b5061048761144c565b3480156104cf57600080fd5b506104d8611452565b604080516001600160a01b039092168252519081900360200190f35b34801561050057600080fd5b506103aa6004803603602081101561051757600080fd5b50356001600160a01b0316611466565b34801561053357600080fd5b506103aa6004803603606081101561054a57600080fd5b810190602081018135600160201b81111561056457600080fd5b82018360208201111561057657600080fd5b803590602001918460018302840111600160201b8311171561059757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092956001600160a01b03853516959094909350604081019250602001359050600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460018302840111600160201b8311171561062d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506115da945050505050565b34801561067a57600080fd5b506103aa6004803603602081101561069157600080fd5b50356001600160a01b0316611a28565b3480156106ad57600080fd5b5061045e600480360360208110156106c457600080fd5b50356001600160a01b0316611ceb565b3480156106e057600080fd5b5061045e600480360360208110156106f757600080fd5b50356001600160a01b0316611d0f565b34801561071357600080fd5b506104d8611d30565b34801561072857600080fd5b50610731611d44565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561076d578181015183820152602001610755565b505050509050019250505060405180910390f35b34801561078d57600080fd5b506103aa600480360360208110156107a457600080fd5b50356001600160a01b0316611da7565b3480156107c057600080fd5b5061045e600480360360208110156107d757600080fd5b810190602081018135600160201b8111156107f157600080fd5b82018360208201111561080357600080fd5b803590602001918460018302840111600160201b8311171561082457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611e7c945050505050565b34801561087157600080fd5b506103aa600480360360c081101561088857600080fd5b810190602081018135600160201b8111156108a257600080fd5b8201836020820111156108b457600080fd5b803590602001918460018302840111600160201b831117156108d557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092956001600160a01b0385358116966020870135966040810135151596506060810135909216945091925060a081019060800135600160201b81111561094d57600080fd5b82018360208201111561095f57600080fd5b803590602001918460018302840111600160201b8311171561098057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611eec945050505050565b3480156109cd57600080fd5b50610487612445565b3480156109e257600080fd5b506103aa61244b565b3480156109f757600080fd5b5061045e6124f4565b348015610a0c57600080fd5b50610a156124fd565b6040805160ff9092168252519081900360200190f35b348015610a3757600080fd5b50610487612506565b3360009081526008602052604090205460ff16600114610aaa5760408051600160e51b62461bcd02815260206004820152601b60248201527f4f6e6c79206d616e616765722063616e20657865637574652069740000000000604482015290519081900360640190fd5b8451604014610b035760408051600160e51b62461bcd02815260206004820152601960248201527f4669786564206c656e677468206f662074784b65793a20363400000000000000604482015290519081900360640190fd5b600084511180610b14575060008351115b1515610b5457604051600160e51b62461bcd028152600401808060200182810382526028815260200180613a0a6028913960400191505060405180910390fd5b600b856040518082805190602001908083835b60208310610b865780518252601f199092019160209182019101610b67565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff16159150610c0e90505760408051600160e51b62461bcd02815260206004820152601e60248201527f5472616e73616374696f6e20686173206265656e20636f6d706c657465640000604482015290519081900360640190fd5b610c18848461250c565b60008585848660026040516020018086805190602001908083835b60208310610c525780518252601f199092019160209182019101610c33565b51815160209384036101000a60001901801990921691161790528851919093019288810192500280838360005b83811015610c97578181015183820152602001610c7f565b505050509050018460ff1660ff1660f81b8152600101838051906020019060200280838360005b83811015610cd6578181015183820152602001610cbe565b5050505060ff94851660f81b92019182525060408051601e198184030181526001909201815281516020928301206000818152600a909352912054909750909116159450610d6b93505050505760408051600160e51b62461bcd0281526020600482015260126024820152600160701b71496e76616c6964207369676e61747572657302604482015290519081900360640190fd5b610d758183612827565b1515610dc65760408051600160e51b62461bcd02815260206004820152601560248201526001605a1b7415985b1a59081cda59db985d1d5c995cc819985a5b02604482015290519081900360640190fd5b610dcf846128a4565b610dd885612a3b565b600954610de490612b08565b6005805460ff191660ff92909216919091179055610e0486826001612b74565b7fac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae866040518080602001828103825283818151815260200191508051906020019080838360005b83811015610e63578181015183820152602001610e4b565b50505050905090810190601f168015610e905780820380516001836020036101000a031916815260200191505b509250505060405180910390a1505050505050565b600033831515610eff5760408051600160e51b62461bcd02815260206004820152601260248201527f4552524f523a205a65726f20616d6f756e740000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038316156111dc573415610f4e57604051600160e51b62461bcd02815260040180806020018281038252602481526020018061386b6024913960400191505060405180910390fd5b610f60836001600160a01b0316612c08565b1515610fa057604051600160e51b62461bcd028152600401808060200182810382526025815260200180613ace6025913960400191505060405180910390fd5b60408051600160e11b636eb1769f0281526001600160a01b038381166004830152306024830152915185926000929084169163dd62ed3e91604480820192602092909190829003018186803b158015610ff857600080fd5b505afa15801561100c573d6000803e3d6000fd5b505050506040513d602081101561102257600080fd5b505190508581101561106857604051600160e51b62461bcd0281526004018080602001828103825260228152602001806139856022913960400191505060405180910390fd5b6000826001600160a01b03166370a08231856040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156110c057600080fd5b505afa1580156110d4573d6000803e3d6000fd5b505050506040513d60208110156110ea57600080fd5b50519050868110156111465760408051600160e51b62461bcd02815260206004820152601e60248201527f4e6f20656e6f7567682062616c616e6365206f662074686520746f6b656e0000604482015290519081900360640190fd5b6111616001600160a01b03841685308a63ffffffff612c4416565b61116a86611ceb565b156111d4576000869050806001600160a01b03166342966c68896040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156111ba57600080fd5b505af11580156111ce573d6000803e3d6000fd5b50505050505b505050611233565b3484146112335760408051600160e51b62461bcd02815260206004820152601d60248201527f496e636f6e73697374656e637920457468657265756d20616d6f756e74000000604482015290519081900360640190fd5b7f5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e306671468186868660405180856001600160a01b03166001600160a01b0316815260200180602001848152602001836001600160a01b03166001600160a01b03168152602001828103825285818151815260200191508051906020019080838360005b838110156112cb5781810151838201526020016112b3565b50505050905090810190601f1680156112f85780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1506001949350505050565b60045481565b60055461010090046001600160a01b0316331461136e5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020613a5e833981519152604482015290519081900360640190fd5b60005460ff1615156113b65760408051600160e51b62461bcd0281526020600482015260066024820152600160d21b6511195b9a595902604482015290519081900360640190fd5b60005461010090046001600160a01b0316151561140757604051600160e51b62461bcd0281526004018080602001828103825260238152602001806139a76023913960400191505060405180910390fd5b600080546040516001600160a01b036101009092049190911691303180156108fc02929091818181858888f19350505050158015611449573d6000803e3d6000fd5b50565b60035481565b60005461010090046001600160a01b031681565b60055461010090046001600160a01b031633146114bb5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020613a5e833981519152604482015290519081900360640190fd5b306001600160a01b038216141561151c5760408051600160e51b62461bcd02815260206004820152601660248201527f446f206e6f7468696e6720627920796f757273656c6600000000000000000000604482015290519081900360640190fd5b61152e816001600160a01b0316612c08565b151561156e57604051600160e51b62461bcd028152600401808060200182810382526025815260200180613ace6025913960400191505060405180910390fd5b61157781611ceb565b156115b657604051600160e51b62461bcd0281526004018080602001828103825260288152602001806137f16028913960400191505060405180910390fd5b6001600160a01b03166000908152600c60205260409020805460ff19166001179055565b3360009081526008602052604090205460ff166001146116445760408051600160e51b62461bcd02815260206004820152601b60248201527f4f6e6c79206d616e616765722063616e20657865637574652069740000000000604482015290519081900360640190fd5b825160401461169d5760408051600160e51b62461bcd02815260206004820152601960248201527f4669786564206c656e677468206f662074784b65793a20363400000000000000604482015290519081900360640190fd5b600b836040518082805190602001908083835b602083106116cf5780518252601f1990920191602091820191016116b0565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1615915061175790505760408051600160e51b62461bcd02815260206004820152601e60248201527f5472616e73616374696f6e20686173206265656e20636f6d706c657465640000604482015290519081900360640190fd5b60005460ff16156117b25760408051600160e51b62461bcd02815260206004820152601460248201527f497420686173206265656e207570677261646564000000000000000000000000604482015290519081900360640190fd5b6117c4826001600160a01b0316612c08565b151561180457604051600160e51b62461bcd028152600401808060200182810382526025815260200180613ace6025913960400191505060405180910390fd5b6000838360026040516020018084805190602001908083835b6020831061183c5780518252601f19909201916020918201910161181d565b51815160209384036101000a60001901801990921691161790526001600160a01b039690961660601b92019182525060ff92831660f81b601482015260408051808303600a19018152601590920181528151918501919091206000818152600a90955293205492945050161590506118f65760408051600160e51b62461bcd0281526020600482015260126024820152600160701b71496e76616c6964207369676e61747572657302604482015290519081900360640190fd5b6119008183612827565b15156119515760408051600160e51b62461bcd02815260206004820152601560248201526001605a1b7415985b1a59081cda59db985d1d5c995cc819985a5b02604482015290519081900360640190fd5b60008054600160ff199091168117610100600160a81b0319166101006001600160a01b03871602179091556119899085908390612b74565b7f5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed0846040518080602001828103825283818151815260200191508051906020019080838360005b838110156119e85781810151838201526020016119d0565b50505050905090810190601f168015611a155780820380516001836020036101000a031916815260200191505b509250505060405180910390a150505050565b60055461010090046001600160a01b03163314611a7d5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020613a5e833981519152604482015290519081900360640190fd5b60005460ff161515611ac55760408051600160e51b62461bcd0281526020600482015260066024820152600160d21b6511195b9a595902604482015290519081900360640190fd5b60005461010090046001600160a01b03161515611b1657604051600160e51b62461bcd0281526004018080602001828103825260238152602001806139a76023913960400191505060405180910390fd5b306001600160a01b0382161415611b775760408051600160e51b62461bcd02815260206004820152601660248201527f446f206e6f7468696e6720627920796f757273656c6600000000000000000000604482015290519081900360640190fd5b611b89816001600160a01b0316612c08565b1515611bc957604051600160e51b62461bcd028152600401808060200182810382526025815260200180613ace6025913960400191505060405180910390fd5b60408051600160e01b6370a08231028152306004820152905182916000916001600160a01b038416916370a08231916024808301926020929190829003018186803b158015611c1757600080fd5b505afa158015611c2b573d6000803e3d6000fd5b505050506040513d6020811015611c4157600080fd5b50519050600054611c6a906001600160a01b03848116916101009004168363ffffffff612ca116565b611c7383611ceb565b15611ce6576000805460408051600160e21b6301fc6bd10281526001600160a01b036101009093048316600482015290518693928416926307f1af44926024808201939182900301818387803b158015611ccc57600080fd5b505af1158015611ce0573d6000803e3d6000fd5b50505050505b505050565b6001600160a01b0381166000908152600c602052604090205460ff1615155b919050565b6001600160a01b031660009081526008602052604090205460ff1660011490565b60055461010090046001600160a01b031681565b60606009805480602002602001604051908101604052809291908181526020018280548015611d9c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611d7e575b505050505090505b90565b60055461010090046001600160a01b03163314611dfc5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020613a5e833981519152604482015290519081900360640190fd5b611e0581611ceb565b1515611e5b5760408051600160e51b62461bcd02815260206004820152601e60248201527f546869732061646472657373206973206e6f7420726567697374657265640000604482015290519081900360640190fd5b6001600160a01b03166000908152600c60205260409020805460ff19169055565b600080600b836040518082805190602001908083835b60208310611eb15780518252601f199092019160209182019101611e92565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1692909211949350505050565b3360009081526008602052604090205460ff16600114611f565760408051600160e51b62461bcd02815260206004820152601b60248201527f4f6e6c79206d616e616765722063616e20657865637574652069740000000000604482015290519081900360640190fd5b8551604014611faf5760408051600160e51b62461bcd02815260206004820152601960248201527f4669786564206c656e677468206f662074784b65793a20363400000000000000604482015290519081900360640190fd5b6001600160a01b0385161515611ff957604051600160e51b62461bcd0281526004018080602001828103825260268152602001806138196026913960400191505060405180910390fd5b83151561203a57604051600160e51b62461bcd02815260040180806020018281038252602881526020018061395d6028913960400191505060405180910390fd5b600b866040518082805190602001908083835b6020831061206c5780518252601f19909201916020918201910161204d565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff161591506120f490505760408051600160e51b62461bcd02815260206004820152601e60248201527f5472616e73616374696f6e20686173206265656e20636f6d706c657465640000604482015290519081900360640190fd5b821561210a57612105828686612cf6565b61214d565b303184111561214d57604051600160e51b62461bcd02815260040180806020018281038252603f81526020018061388f603f913960400191505060405180910390fd5b6000868686868660026040516020018087805190602001908083835b602083106121885780518252601f199092019160209182019101612169565b51815160209384036101000a60001901801990921691161790526001600160a01b03998a16606090811b92909401918252601482019890985295151560f890811b603488015294909716901b60358501525060ff90811690911b604983015260408051808403602a018152604a90930181528251928401929092206000818152600a90945291909220549094501615915061226790505760408051600160e51b62461bcd0281526020600482015260126024820152600160701b71496e76616c6964207369676e61747572657302604482015290519081900360640190fd5b6122718183612827565b15156122c25760408051600160e51b62461bcd02815260206004820152601560248201526001605a1b7415985b1a59081cda59db985d1d5c995cc819985a5b02604482015290519081900360640190fd5b83156122d8576122d3838787612ee1565b612397565b303185111561231b57604051600160e51b62461bcd02815260040180806020018281038252603f81526020018061388f603f913960400191505060405180910390fd5b6040516001600160a01b0387169086156108fc029087906000818181858888f19350505050158015612351573d6000803e3d6000fd5b50604080516001600160a01b03881681526020810187905281517fc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a929181900390910190a15b6123a387826001612b74565b7f8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1876040518080602001828103825283818151815260200191508051906020019080838360005b838110156124025781810151838201526020016123ea565b50505050905090810190601f16801561242f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a150505050505050565b60025481565b60055461010090046001600160a01b031633146124a05760408051600160e51b62461bcd0281526020600482015260196024820152600080516020613a5e833981519152604482015290519081900360640190fd5b60005460ff1615156124e85760408051600160e51b62461bcd0281526020600482015260066024820152600160d21b6511195b9a595902604482015290519081900360640190fd5b6000805460ff19169055565b60005460ff1681565b60055460ff1681565b60015481565b815160005b818110156125e3576000848281518110151561252957fe5b6020908102909101015190506001600160a01b038116151561257f57604051600160e51b62461bcd0281526004018080602001828103825260248152602001806138ce6024913960400191505060405180910390fd5b6001600160a01b03811660009081526008602052604090205460ff16156125da57604051600160e51b62461bcd0281526004018080602001828103825260408152602001806139ca6040913960400191505060405180910390fd5b50600101612511565b506125ed83613050565b151561262d57604051600160e51b62461bcd02815260040180806020018281038252602c81526020018061383f602c913960400191505060405180910390fd5b6005546126489061010090046001600160a01b031684613112565b151561268857604051600160e51b62461bcd028152600401808060200182810382526026815260200180613aa86026913960400191505060405180910390fd5b61269182613050565b15156126d157604051600160e51b62461bcd02815260040180806020018281038252602c815260200180613a32602c913960400191505060405180910390fd5b815160005b818110156127d457600084828151811015156126ee57fe5b60209081029091018101516001600160a01b0381166000908152600690925260409091205490915060ff161561276e5760408051600160e51b62461bcd02815260206004820152601760248201527f43616e277420657869742073656564206d616e61676572000000000000000000604482015290519081900360640190fd5b6001600160a01b03811660009081526008602052604090205460ff166001146127cb57604051600160e51b62461bcd0281526004018080602001828103825260448152602001806139196044913960600191505060405180910390fd5b506001016126d6565b50600154835185516009540103111561282157604051600160e51b62461bcd0281526004018080602001828103825260278152602001806138f26027913960400191505060405180910390fd5b50505050565b60006103cf8251111515156128865760408051600160e51b62461bcd02815260206004820152601d60248201527f4d6178206c656e677468206f66207369676e6174757265733a20393735000000604482015290519081900360640190fd5b60006128928484613187565b60055460ff1611159150505b92915050565b805115156128b157611449565b60005b8151811015612901576008600083838151811015156128cf57fe5b60209081029091018101516001600160a01b03168252810191909152604001600020805460ff191690556001016128b4565b5060005b600954811015612980576008600060098381548110151561292257fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16151561297857600980548290811061295e57fe5b600091825260209091200180546001600160a01b03191690555b600101612905565b50601060005b600954811015612a255760006009828154811015156129a157fe5b6000918252602090912001546001600160a01b031690508015156129d25782601014156129cc578192505b50612a1d565b60108314612a1b57806009848154811015156129ea57fe5b600091825260209091200180546001600160a01b0319166001600160a01b0392909216919091179055600192909201915b505b600101612986565b508151600980549190910390611ce69082613790565b80511515612a4857611449565b60005b8151811015612b045760008282815181101515612a6457fe5b60209081029091018101516001600160a01b0381166000908152600890925260409091205490915060ff161515612afb576001600160a01b0381166000818152600860205260408120805460ff191660019081179091556009805491820181559091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03191690911790555b50600101612a4b565b5050565b6000811515612b615760408051600160e51b62461bcd02815260206004820152601460248201527f4d616e616765722043616e277420656d7074792e000000000000000000000000604482015290519081900360640190fd5b6003548202606301606481049392505050565b80600b846040518082805190602001908083835b60208310612ba75780518252601f199092019160209182019101612b88565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201909420805460ff1990811660ff978816179091556000978852600a9091529290952080549092169390921692909217909155505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590612c3c57508115155b949350505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b0316600160e01b6323b872dd02179052612821908590613349565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0316600160e01b63a9059cbb02179052611ce6908490613349565b6001600160a01b0382161515612d4057604051600160e51b62461bcd0281526004018080602001828103825260238152602001806137ce6023913960400191505060405180910390fd5b306001600160a01b0384161415612da15760408051600160e51b62461bcd02815260206004820152601660248201527f446f206e6f7468696e6720627920796f757273656c6600000000000000000000604482015290519081900360640190fd5b612db3836001600160a01b0316612c08565b1515612df357604051600160e51b62461bcd028152600401808060200182810382526025815260200180613ace6025913960400191505060405180910390fd5b612dfc83611ceb565b15612e0657611ce6565b60408051600160e01b6370a08231028152306004820152905184916000916001600160a01b038416916370a08231916024808301926020929190829003018186803b158015612e5457600080fd5b505afa158015612e68573d6000803e3d6000fd5b505050506040513d6020811015612e7e57600080fd5b5051905082811015612eda5760408051600160e51b62461bcd02815260206004820152601a60248201527f4e6f20656e6f7567682062616c616e6365206f6620746f6b656e000000000000604482015290519081900360640190fd5b5050505050565b612eea83611ceb565b15612f625760408051600160e01b6340c10f190281526001600160a01b03848116600483015260248201849052915185928316916340c10f1991604480830192600092919082900301818387803b158015612f4457600080fd5b505af1158015612f58573d6000803e3d6000fd5b5050505050611ce6565b60408051600160e01b6370a08231028152306004820152905184916000916001600160a01b038416916370a08231916024808301926020929190829003018186803b158015612fb057600080fd5b505afa158015612fc4573d6000803e3d6000fd5b505050506040513d6020811015612fda57600080fd5b50519050828110156130365760408051600160e51b62461bcd02815260206004820152601a60248201527f4e6f20656e6f7567682062616c616e6365206f6620746f6b656e000000000000604482015290519081900360640190fd5b612eda6001600160a01b038316858563ffffffff612ca116565b6000805b8251811015613109576000838281518110151561306d57fe5b6020908102909101015190506001600160a01b038116151561308f5750613109565b600182015b84518110156130ff57600085828151811015156130ad57fe5b6020908102909101015190506001600160a01b03811615156130cf57506130ff565b806001600160a01b0316836001600160a01b031614156130f6576000945050505050611d0a565b50600101613094565b5050600101613054565b50600192915050565b600080805b835181101561317c57838181518110151561312e57fe5b6020908102909101015191506001600160a01b038216151561314f5761317c565b846001600160a01b0316826001600160a01b031614156131745760009250505061289e565b600101613117565b506001949350505050565b6004548151600091829182916131a3919063ffffffff61351016565b90506060816040519080825280602002602001820160405280156131d1578160200160208202803883390190505b509050600080805b848110156132d35760606131fa846004548b6135599092919063ffffffff16565b905060006132088b836135db565b90506001600160a01b038116151561326a5760408051600160e51b62461bcd02815260206004820152601060248201527f5369676e617475726573206572726f7200000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03811660009081526008602052604090205460ff16600114156132c15785516001988901988501948291889160ff169081106132a957fe5b6001600160a01b039092166020928302909101909101525b505060045492909201916001016131d9565b5060006132df84613050565b90506060935080151561333c5760408051600160e51b62461bcd02815260206004820152601460248201527f5369676e617475726573206475706c6963617465000000000000000000000000604482015290519081900360640190fd5b5093979650505050505050565b61335b826001600160a01b0316612c08565b15156133b15760408051600160e51b62461bcd02815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b602083106133ef5780518252601f1990920191602091820191016133d0565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613451576040519150601f19603f3d011682016040523d82523d6000602084013e613456565b606091505b50915091508115156134b25760408051600160e51b62461bcd02815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115612821578080602001905160208110156134ce57600080fd5b5051151561282157604051600160e51b62461bcd02815260040180806020018281038252602a815260200180613a7e602a913960400191505060405180910390fd5b600061355283836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506136e7565b9392505050565b606081830184511015151561356d57600080fd5b606082158015613588576040519150602082016040526135d2565b6040519150601f8416801560200281840101858101878315602002848b0101015b818310156135c15780518352602092830192016135a9565b5050858452601f01601f1916604052505b50949350505050565b60008060008060045485511415156135f9576000935050505061289e565b50505060208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115613642576000935050505061289e565b601b8160ff16101561365257601b015b8060ff16601b1415801561366a57508060ff16601c14155b1561367b576000935050505061289e565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156136d2573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008183151561377857604051600160e51b62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561373d578181015183820152602001613725565b50505050905090810190601f16801561376a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581151561378657fe5b0495945050505050565b815481835581811115611ce657600083815260209020611ce6918101908301611da491905b808211156137c957600081556001016137b5565b509056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735468697320616464726573732068617320616c7265616479206265656e207265676973746572656457697468647261773a207472616e7366657220746f20746865207a65726f20616464726573734475706c696361746520706172616d657465727320666f7220746865206164647265737320746f206a6f696e45524332303a20446f6573206e6f742061636365707420457468657265756d20436f696e5468697320636f6e7472616374206164647265737320646f6573206e6f7420686176652073756666696369656e742062616c616e6365206f662065746865724552524f523a204465746563746564207a65726f206164647265737320696e2061646473457863656564656420746865206d6178696d756d206e756d626572206f66206d616e61676572735468657265206172652061646472657373657320696e207468652065786974696e672061646472657373206c697374207468617420617265206e6f74206d616e616765725769746864726177616c20616d6f756e74206d7573742062652067726561746572207468616e20304e6f20656e6f75676820616d6f756e7420666f7220617574686f72697a6174696f6e4552524f523a207472616e7366657220746f20746865207a65726f20616464726573735468652061646472657373206c6973742074686174206973206265696e6720616464656420616c7265616479206578697374732061732061206d616e61676572546865726520617265206e6f206d616e6167657273206a6f696e696e67206f722065786974696e674475706c696361746520706172616d657465727320666f7220746865206164647265737320746f20657869744f6e6c79206f776e65722063616e2065786563757465206974000000000000005361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564436f6e74726163742063726561746f722063616e6e6f7420616374206173206d616e616765725468652061646472657373206973206e6f74206120636f6e74726163742061646472657373a165627a7a723058205eec7306ed60f78fea389b5b61ab203f6d58e8e545693845587da8ece7977d2700294e6f74207265616368696e6720746865206d696e206e756d626572206f66206d616e6167657273457863656564656420746865206d6178696d756d206e756d626572206f66206d616e6167657273436f6e74726163742063726561746f722063616e6e6f7420616374206173206d616e61676572"

// DeployPBridge deploys a new Ethereum contract, binding an instance of PBridge to it.
func DeployPBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _managers []common.Address) (common.Address, *types.Transaction, *PBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(PBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PBridgeBin), backend, _managers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PBridge{PBridgeCaller: PBridgeCaller{contract: contract}, PBridgeTransactor: PBridgeTransactor{contract: contract}, PBridgeFilterer: PBridgeFilterer{contract: contract}}, nil
}

// PBridge is an auto generated Go binding around an Ethereum contract.
type PBridge struct {
	PBridgeCaller     // Read-only binding to the contract
	PBridgeTransactor // Write-only binding to the contract
	PBridgeFilterer   // Log filterer for contract events
}

// PBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PBridgeSession struct {
	Contract     *PBridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PBridgeCallerSession struct {
	Contract *PBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PBridgeTransactorSession struct {
	Contract     *PBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PBridgeRaw struct {
	Contract *PBridge // Generic contract binding to access the raw methods on
}

// PBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PBridgeCallerRaw struct {
	Contract *PBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// PBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PBridgeTransactorRaw struct {
	Contract *PBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPBridge creates a new instance of PBridge, bound to a specific deployed contract.
func NewPBridge(address common.Address, backend bind.ContractBackend) (*PBridge, error) {
	contract, err := bindPBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PBridge{PBridgeCaller: PBridgeCaller{contract: contract}, PBridgeTransactor: PBridgeTransactor{contract: contract}, PBridgeFilterer: PBridgeFilterer{contract: contract}}, nil
}

// NewPBridgeCaller creates a new read-only instance of PBridge, bound to a specific deployed contract.
func NewPBridgeCaller(address common.Address, caller bind.ContractCaller) (*PBridgeCaller, error) {
	contract, err := bindPBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PBridgeCaller{contract: contract}, nil
}

// NewPBridgeTransactor creates a new write-only instance of PBridge, bound to a specific deployed contract.
func NewPBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PBridgeTransactor, error) {
	contract, err := bindPBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PBridgeTransactor{contract: contract}, nil
}

// NewPBridgeFilterer creates a new log filterer instance of PBridge, bound to a specific deployed contract.
func NewPBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PBridgeFilterer, error) {
	contract, err := bindPBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PBridgeFilterer{contract: contract}, nil
}

// bindPBridge binds a generic wrapper to an already deployed contract.
func bindPBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PBridge *PBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PBridge.Contract.PBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PBridge *PBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PBridge.Contract.PBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PBridge *PBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PBridge.Contract.PBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PBridge *PBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PBridge *PBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PBridge *PBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PBridge.Contract.contract.Transact(opts, method, params...)
}

// AllManagers is a free data retrieval call binding the contract method 0x9c30b35e.
//
// Solidity: function allManagers() view returns(address[])
func (_PBridge *PBridgeCaller) AllManagers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "allManagers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllManagers is a free data retrieval call binding the contract method 0x9c30b35e.
//
// Solidity: function allManagers() view returns(address[])
func (_PBridge *PBridgeSession) AllManagers() ([]common.Address, error) {
	return _PBridge.Contract.AllManagers(&_PBridge.CallOpts)
}

// AllManagers is a free data retrieval call binding the contract method 0x9c30b35e.
//
// Solidity: function allManagers() view returns(address[])
func (_PBridge *PBridgeCallerSession) AllManagers() ([]common.Address, error) {
	return _PBridge.Contract.AllManagers(&_PBridge.CallOpts)
}

// CurrentMinSignatures is a free data retrieval call binding the contract method 0xe079cee9.
//
// Solidity: function current_min_signatures() view returns(uint8)
func (_PBridge *PBridgeCaller) CurrentMinSignatures(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "current_min_signatures")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrentMinSignatures is a free data retrieval call binding the contract method 0xe079cee9.
//
// Solidity: function current_min_signatures() view returns(uint8)
func (_PBridge *PBridgeSession) CurrentMinSignatures() (uint8, error) {
	return _PBridge.Contract.CurrentMinSignatures(&_PBridge.CallOpts)
}

// CurrentMinSignatures is a free data retrieval call binding the contract method 0xe079cee9.
//
// Solidity: function current_min_signatures() view returns(uint8)
func (_PBridge *PBridgeCallerSession) CurrentMinSignatures() (uint8, error) {
	return _PBridge.Contract.CurrentMinSignatures(&_PBridge.CallOpts)
}

// IfManager is a free data retrieval call binding the contract method 0x75173b70.
//
// Solidity: function ifManager(address _manager) view returns(bool)
func (_PBridge *PBridgeCaller) IfManager(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "ifManager", _manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IfManager is a free data retrieval call binding the contract method 0x75173b70.
//
// Solidity: function ifManager(address _manager) view returns(bool)
func (_PBridge *PBridgeSession) IfManager(_manager common.Address) (bool, error) {
	return _PBridge.Contract.IfManager(&_PBridge.CallOpts, _manager)
}

// IfManager is a free data retrieval call binding the contract method 0x75173b70.
//
// Solidity: function ifManager(address _manager) view returns(bool)
func (_PBridge *PBridgeCallerSession) IfManager(_manager common.Address) (bool, error) {
	return _PBridge.Contract.IfManager(&_PBridge.CallOpts, _manager)
}

// IsCompletedTx is a free data retrieval call binding the contract method 0xa5e399b3.
//
// Solidity: function isCompletedTx(string txKey) view returns(bool)
func (_PBridge *PBridgeCaller) IsCompletedTx(opts *bind.CallOpts, txKey string) (bool, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "isCompletedTx", txKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCompletedTx is a free data retrieval call binding the contract method 0xa5e399b3.
//
// Solidity: function isCompletedTx(string txKey) view returns(bool)
func (_PBridge *PBridgeSession) IsCompletedTx(txKey string) (bool, error) {
	return _PBridge.Contract.IsCompletedTx(&_PBridge.CallOpts, txKey)
}

// IsCompletedTx is a free data retrieval call binding the contract method 0xa5e399b3.
//
// Solidity: function isCompletedTx(string txKey) view returns(bool)
func (_PBridge *PBridgeCallerSession) IsCompletedTx(txKey string) (bool, error) {
	return _PBridge.Contract.IsCompletedTx(&_PBridge.CallOpts, txKey)
}

// IsMinterERC20 is a free data retrieval call binding the contract method 0x6a7142e1.
//
// Solidity: function isMinterERC20(address ERC20) view returns(bool)
func (_PBridge *PBridgeCaller) IsMinterERC20(opts *bind.CallOpts, ERC20 common.Address) (bool, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "isMinterERC20", ERC20)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinterERC20 is a free data retrieval call binding the contract method 0x6a7142e1.
//
// Solidity: function isMinterERC20(address ERC20) view returns(bool)
func (_PBridge *PBridgeSession) IsMinterERC20(ERC20 common.Address) (bool, error) {
	return _PBridge.Contract.IsMinterERC20(&_PBridge.CallOpts, ERC20)
}

// IsMinterERC20 is a free data retrieval call binding the contract method 0x6a7142e1.
//
// Solidity: function isMinterERC20(address ERC20) view returns(bool)
func (_PBridge *PBridgeCallerSession) IsMinterERC20(ERC20 common.Address) (bool, error) {
	return _PBridge.Contract.IsMinterERC20(&_PBridge.CallOpts, ERC20)
}

// MaxManagers is a free data retrieval call binding the contract method 0xf7f2ff74.
//
// Solidity: function max_managers() view returns(uint256)
func (_PBridge *PBridgeCaller) MaxManagers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "max_managers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxManagers is a free data retrieval call binding the contract method 0xf7f2ff74.
//
// Solidity: function max_managers() view returns(uint256)
func (_PBridge *PBridgeSession) MaxManagers() (*big.Int, error) {
	return _PBridge.Contract.MaxManagers(&_PBridge.CallOpts)
}

// MaxManagers is a free data retrieval call binding the contract method 0xf7f2ff74.
//
// Solidity: function max_managers() view returns(uint256)
func (_PBridge *PBridgeCallerSession) MaxManagers() (*big.Int, error) {
	return _PBridge.Contract.MaxManagers(&_PBridge.CallOpts)
}

// MinManagers is a free data retrieval call binding the contract method 0xad4b61a8.
//
// Solidity: function min_managers() view returns(uint256)
func (_PBridge *PBridgeCaller) MinManagers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "min_managers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinManagers is a free data retrieval call binding the contract method 0xad4b61a8.
//
// Solidity: function min_managers() view returns(uint256)
func (_PBridge *PBridgeSession) MinManagers() (*big.Int, error) {
	return _PBridge.Contract.MinManagers(&_PBridge.CallOpts)
}

// MinManagers is a free data retrieval call binding the contract method 0xad4b61a8.
//
// Solidity: function min_managers() view returns(uint256)
func (_PBridge *PBridgeCallerSession) MinManagers() (*big.Int, error) {
	return _PBridge.Contract.MinManagers(&_PBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PBridge *PBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PBridge *PBridgeSession) Owner() (common.Address, error) {
	return _PBridge.Contract.Owner(&_PBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PBridge *PBridgeCallerSession) Owner() (common.Address, error) {
	return _PBridge.Contract.Owner(&_PBridge.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_PBridge *PBridgeCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_PBridge *PBridgeSession) Rate() (*big.Int, error) {
	return _PBridge.Contract.Rate(&_PBridge.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_PBridge *PBridgeCallerSession) Rate() (*big.Int, error) {
	return _PBridge.Contract.Rate(&_PBridge.CallOpts)
}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_PBridge *PBridgeCaller) SignatureLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "signatureLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_PBridge *PBridgeSession) SignatureLength() (*big.Int, error) {
	return _PBridge.Contract.SignatureLength(&_PBridge.CallOpts)
}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_PBridge *PBridgeCallerSession) SignatureLength() (*big.Int, error) {
	return _PBridge.Contract.SignatureLength(&_PBridge.CallOpts)
}

// Upgrade is a free data retrieval call binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() view returns(bool)
func (_PBridge *PBridgeCaller) Upgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "upgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Upgrade is a free data retrieval call binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() view returns(bool)
func (_PBridge *PBridgeSession) Upgrade() (bool, error) {
	return _PBridge.Contract.Upgrade(&_PBridge.CallOpts)
}

// Upgrade is a free data retrieval call binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() view returns(bool)
func (_PBridge *PBridgeCallerSession) Upgrade() (bool, error) {
	return _PBridge.Contract.Upgrade(&_PBridge.CallOpts)
}

// UpgradeContractAddress is a free data retrieval call binding the contract method 0x30b2d84d.
//
// Solidity: function upgradeContractAddress() view returns(address)
func (_PBridge *PBridgeCaller) UpgradeContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PBridge.contract.Call(opts, &out, "upgradeContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeContractAddress is a free data retrieval call binding the contract method 0x30b2d84d.
//
// Solidity: function upgradeContractAddress() view returns(address)
func (_PBridge *PBridgeSession) UpgradeContractAddress() (common.Address, error) {
	return _PBridge.Contract.UpgradeContractAddress(&_PBridge.CallOpts)
}

// UpgradeContractAddress is a free data retrieval call binding the contract method 0x30b2d84d.
//
// Solidity: function upgradeContractAddress() view returns(address)
func (_PBridge *PBridgeCallerSession) UpgradeContractAddress() (common.Address, error) {
	return _PBridge.Contract.UpgradeContractAddress(&_PBridge.CallOpts)
}

// CloseUpgrade is a paid mutator transaction binding the contract method 0xd4cacbaa.
//
// Solidity: function closeUpgrade() returns()
func (_PBridge *PBridgeTransactor) CloseUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "closeUpgrade")
}

// CloseUpgrade is a paid mutator transaction binding the contract method 0xd4cacbaa.
//
// Solidity: function closeUpgrade() returns()
func (_PBridge *PBridgeSession) CloseUpgrade() (*types.Transaction, error) {
	return _PBridge.Contract.CloseUpgrade(&_PBridge.TransactOpts)
}

// CloseUpgrade is a paid mutator transaction binding the contract method 0xd4cacbaa.
//
// Solidity: function closeUpgrade() returns()
func (_PBridge *PBridgeTransactorSession) CloseUpgrade() (*types.Transaction, error) {
	return _PBridge.Contract.CloseUpgrade(&_PBridge.TransactOpts)
}

// CreateOrSignManagerChange is a paid mutator transaction binding the contract method 0x00719226.
//
// Solidity: function createOrSignManagerChange(string txKey, address[] adds, address[] removes, uint8 count, bytes signatures) returns()
func (_PBridge *PBridgeTransactor) CreateOrSignManagerChange(opts *bind.TransactOpts, txKey string, adds []common.Address, removes []common.Address, count uint8, signatures []byte) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "createOrSignManagerChange", txKey, adds, removes, count, signatures)
}

// CreateOrSignManagerChange is a paid mutator transaction binding the contract method 0x00719226.
//
// Solidity: function createOrSignManagerChange(string txKey, address[] adds, address[] removes, uint8 count, bytes signatures) returns()
func (_PBridge *PBridgeSession) CreateOrSignManagerChange(txKey string, adds []common.Address, removes []common.Address, count uint8, signatures []byte) (*types.Transaction, error) {
	return _PBridge.Contract.CreateOrSignManagerChange(&_PBridge.TransactOpts, txKey, adds, removes, count, signatures)
}

// CreateOrSignManagerChange is a paid mutator transaction binding the contract method 0x00719226.
//
// Solidity: function createOrSignManagerChange(string txKey, address[] adds, address[] removes, uint8 count, bytes signatures) returns()
func (_PBridge *PBridgeTransactorSession) CreateOrSignManagerChange(txKey string, adds []common.Address, removes []common.Address, count uint8, signatures []byte) (*types.Transaction, error) {
	return _PBridge.Contract.CreateOrSignManagerChange(&_PBridge.TransactOpts, txKey, adds, removes, count, signatures)
}

// CreateOrSignUpgrade is a paid mutator transaction binding the contract method 0x408e8b7a.
//
// Solidity: function createOrSignUpgrade(string txKey, address upgradeContract, bytes signatures) returns()
func (_PBridge *PBridgeTransactor) CreateOrSignUpgrade(opts *bind.TransactOpts, txKey string, upgradeContract common.Address, signatures []byte) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "createOrSignUpgrade", txKey, upgradeContract, signatures)
}

// CreateOrSignUpgrade is a paid mutator transaction binding the contract method 0x408e8b7a.
//
// Solidity: function createOrSignUpgrade(string txKey, address upgradeContract, bytes signatures) returns()
func (_PBridge *PBridgeSession) CreateOrSignUpgrade(txKey string, upgradeContract common.Address, signatures []byte) (*types.Transaction, error) {
	return _PBridge.Contract.CreateOrSignUpgrade(&_PBridge.TransactOpts, txKey, upgradeContract, signatures)
}

// CreateOrSignUpgrade is a paid mutator transaction binding the contract method 0x408e8b7a.
//
// Solidity: function createOrSignUpgrade(string txKey, address upgradeContract, bytes signatures) returns()
func (_PBridge *PBridgeTransactorSession) CreateOrSignUpgrade(txKey string, upgradeContract common.Address, signatures []byte) (*types.Transaction, error) {
	return _PBridge.Contract.CreateOrSignUpgrade(&_PBridge.TransactOpts, txKey, upgradeContract, signatures)
}

// CreateOrSignWithdraw is a paid mutator transaction binding the contract method 0xab6c2b10.
//
// Solidity: function createOrSignWithdraw(string txKey, address to, uint256 amount, bool isERC20, address ERC20, bytes signatures) returns()
func (_PBridge *PBridgeTransactor) CreateOrSignWithdraw(opts *bind.TransactOpts, txKey string, to common.Address, amount *big.Int, isERC20 bool, ERC20 common.Address, signatures []byte) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "createOrSignWithdraw", txKey, to, amount, isERC20, ERC20, signatures)
}

// CreateOrSignWithdraw is a paid mutator transaction binding the contract method 0xab6c2b10.
//
// Solidity: function createOrSignWithdraw(string txKey, address to, uint256 amount, bool isERC20, address ERC20, bytes signatures) returns()
func (_PBridge *PBridgeSession) CreateOrSignWithdraw(txKey string, to common.Address, amount *big.Int, isERC20 bool, ERC20 common.Address, signatures []byte) (*types.Transaction, error) {
	return _PBridge.Contract.CreateOrSignWithdraw(&_PBridge.TransactOpts, txKey, to, amount, isERC20, ERC20, signatures)
}

// CreateOrSignWithdraw is a paid mutator transaction binding the contract method 0xab6c2b10.
//
// Solidity: function createOrSignWithdraw(string txKey, address to, uint256 amount, bool isERC20, address ERC20, bytes signatures) returns()
func (_PBridge *PBridgeTransactorSession) CreateOrSignWithdraw(txKey string, to common.Address, amount *big.Int, isERC20 bool, ERC20 common.Address, signatures []byte) (*types.Transaction, error) {
	return _PBridge.Contract.CreateOrSignWithdraw(&_PBridge.TransactOpts, txKey, to, amount, isERC20, ERC20, signatures)
}

// CrossOut is a paid mutator transaction binding the contract method 0x0889d1f0.
//
// Solidity: function crossOut(string to, uint256 amount, address ERC20) payable returns(bool)
func (_PBridge *PBridgeTransactor) CrossOut(opts *bind.TransactOpts, to string, amount *big.Int, ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "crossOut", to, amount, ERC20)
}

// CrossOut is a paid mutator transaction binding the contract method 0x0889d1f0.
//
// Solidity: function crossOut(string to, uint256 amount, address ERC20) payable returns(bool)
func (_PBridge *PBridgeSession) CrossOut(to string, amount *big.Int, ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.CrossOut(&_PBridge.TransactOpts, to, amount, ERC20)
}

// CrossOut is a paid mutator transaction binding the contract method 0x0889d1f0.
//
// Solidity: function crossOut(string to, uint256 amount, address ERC20) payable returns(bool)
func (_PBridge *PBridgeTransactorSession) CrossOut(to string, amount *big.Int, ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.CrossOut(&_PBridge.TransactOpts, to, amount, ERC20)
}

// RegisterMinterERC20 is a paid mutator transaction binding the contract method 0x34774b71.
//
// Solidity: function registerMinterERC20(address ERC20) returns()
func (_PBridge *PBridgeTransactor) RegisterMinterERC20(opts *bind.TransactOpts, ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "registerMinterERC20", ERC20)
}

// RegisterMinterERC20 is a paid mutator transaction binding the contract method 0x34774b71.
//
// Solidity: function registerMinterERC20(address ERC20) returns()
func (_PBridge *PBridgeSession) RegisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.RegisterMinterERC20(&_PBridge.TransactOpts, ERC20)
}

// RegisterMinterERC20 is a paid mutator transaction binding the contract method 0x34774b71.
//
// Solidity: function registerMinterERC20(address ERC20) returns()
func (_PBridge *PBridgeTransactorSession) RegisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.RegisterMinterERC20(&_PBridge.TransactOpts, ERC20)
}

// UnregisterMinterERC20 is a paid mutator transaction binding the contract method 0x9dcdc978.
//
// Solidity: function unregisterMinterERC20(address ERC20) returns()
func (_PBridge *PBridgeTransactor) UnregisterMinterERC20(opts *bind.TransactOpts, ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "unregisterMinterERC20", ERC20)
}

// UnregisterMinterERC20 is a paid mutator transaction binding the contract method 0x9dcdc978.
//
// Solidity: function unregisterMinterERC20(address ERC20) returns()
func (_PBridge *PBridgeSession) UnregisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.UnregisterMinterERC20(&_PBridge.TransactOpts, ERC20)
}

// UnregisterMinterERC20 is a paid mutator transaction binding the contract method 0x9dcdc978.
//
// Solidity: function unregisterMinterERC20(address ERC20) returns()
func (_PBridge *PBridgeTransactorSession) UnregisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.UnregisterMinterERC20(&_PBridge.TransactOpts, ERC20)
}

// UpgradeContractS1 is a paid mutator transaction binding the contract method 0x1dda9c05.
//
// Solidity: function upgradeContractS1() returns()
func (_PBridge *PBridgeTransactor) UpgradeContractS1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "upgradeContractS1")
}

// UpgradeContractS1 is a paid mutator transaction binding the contract method 0x1dda9c05.
//
// Solidity: function upgradeContractS1() returns()
func (_PBridge *PBridgeSession) UpgradeContractS1() (*types.Transaction, error) {
	return _PBridge.Contract.UpgradeContractS1(&_PBridge.TransactOpts)
}

// UpgradeContractS1 is a paid mutator transaction binding the contract method 0x1dda9c05.
//
// Solidity: function upgradeContractS1() returns()
func (_PBridge *PBridgeTransactorSession) UpgradeContractS1() (*types.Transaction, error) {
	return _PBridge.Contract.UpgradeContractS1(&_PBridge.TransactOpts)
}

// UpgradeContractS2 is a paid mutator transaction binding the contract method 0x5bda3fcf.
//
// Solidity: function upgradeContractS2(address ERC20) returns()
func (_PBridge *PBridgeTransactor) UpgradeContractS2(opts *bind.TransactOpts, ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.contract.Transact(opts, "upgradeContractS2", ERC20)
}

// UpgradeContractS2 is a paid mutator transaction binding the contract method 0x5bda3fcf.
//
// Solidity: function upgradeContractS2(address ERC20) returns()
func (_PBridge *PBridgeSession) UpgradeContractS2(ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.UpgradeContractS2(&_PBridge.TransactOpts, ERC20)
}

// UpgradeContractS2 is a paid mutator transaction binding the contract method 0x5bda3fcf.
//
// Solidity: function upgradeContractS2(address ERC20) returns()
func (_PBridge *PBridgeTransactorSession) UpgradeContractS2(ERC20 common.Address) (*types.Transaction, error) {
	return _PBridge.Contract.UpgradeContractS2(&_PBridge.TransactOpts, ERC20)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PBridge *PBridgeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PBridge.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PBridge *PBridgeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PBridge.Contract.Fallback(&_PBridge.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PBridge *PBridgeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PBridge.Contract.Fallback(&_PBridge.TransactOpts, calldata)
}

// PBridgeCrossOutFundsIterator is returned from FilterCrossOutFunds and is used to iterate over the raw logs and unpacked data for CrossOutFunds events raised by the PBridge contract.
type PBridgeCrossOutFundsIterator struct {
	Event *PBridgeCrossOutFunds // Event containing the contract specifics and raw log

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
func (it *PBridgeCrossOutFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PBridgeCrossOutFunds)
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
		it.Event = new(PBridgeCrossOutFunds)
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
func (it *PBridgeCrossOutFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PBridgeCrossOutFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PBridgeCrossOutFunds represents a CrossOutFunds event raised by the PBridge contract.
type PBridgeCrossOutFunds struct {
	From   common.Address
	To     string
	Amount *big.Int
	ERC20  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCrossOutFunds is a free log retrieval operation binding the contract event 0x5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e30667146.
//
// Solidity: event CrossOutFunds(address from, string to, uint256 amount, address ERC20)
func (_PBridge *PBridgeFilterer) FilterCrossOutFunds(opts *bind.FilterOpts) (*PBridgeCrossOutFundsIterator, error) {

	logs, sub, err := _PBridge.contract.FilterLogs(opts, "CrossOutFunds")
	if err != nil {
		return nil, err
	}
	return &PBridgeCrossOutFundsIterator{contract: _PBridge.contract, event: "CrossOutFunds", logs: logs, sub: sub}, nil
}

// WatchCrossOutFunds is a free log subscription operation binding the contract event 0x5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e30667146.
//
// Solidity: event CrossOutFunds(address from, string to, uint256 amount, address ERC20)
func (_PBridge *PBridgeFilterer) WatchCrossOutFunds(opts *bind.WatchOpts, sink chan<- *PBridgeCrossOutFunds) (event.Subscription, error) {

	logs, sub, err := _PBridge.contract.WatchLogs(opts, "CrossOutFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PBridgeCrossOutFunds)
				if err := _PBridge.contract.UnpackLog(event, "CrossOutFunds", log); err != nil {
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

// ParseCrossOutFunds is a log parse operation binding the contract event 0x5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e30667146.
//
// Solidity: event CrossOutFunds(address from, string to, uint256 amount, address ERC20)
func (_PBridge *PBridgeFilterer) ParseCrossOutFunds(log types.Log) (*PBridgeCrossOutFunds, error) {
	event := new(PBridgeCrossOutFunds)
	if err := _PBridge.contract.UnpackLog(event, "CrossOutFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PBridgeDepositFundsIterator is returned from FilterDepositFunds and is used to iterate over the raw logs and unpacked data for DepositFunds events raised by the PBridge contract.
type PBridgeDepositFundsIterator struct {
	Event *PBridgeDepositFunds // Event containing the contract specifics and raw log

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
func (it *PBridgeDepositFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PBridgeDepositFunds)
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
		it.Event = new(PBridgeDepositFunds)
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
func (it *PBridgeDepositFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PBridgeDepositFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PBridgeDepositFunds represents a DepositFunds event raised by the PBridge contract.
type PBridgeDepositFunds struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositFunds is a free log retrieval operation binding the contract event 0xd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88.
//
// Solidity: event DepositFunds(address from, uint256 amount)
func (_PBridge *PBridgeFilterer) FilterDepositFunds(opts *bind.FilterOpts) (*PBridgeDepositFundsIterator, error) {

	logs, sub, err := _PBridge.contract.FilterLogs(opts, "DepositFunds")
	if err != nil {
		return nil, err
	}
	return &PBridgeDepositFundsIterator{contract: _PBridge.contract, event: "DepositFunds", logs: logs, sub: sub}, nil
}

// WatchDepositFunds is a free log subscription operation binding the contract event 0xd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88.
//
// Solidity: event DepositFunds(address from, uint256 amount)
func (_PBridge *PBridgeFilterer) WatchDepositFunds(opts *bind.WatchOpts, sink chan<- *PBridgeDepositFunds) (event.Subscription, error) {

	logs, sub, err := _PBridge.contract.WatchLogs(opts, "DepositFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PBridgeDepositFunds)
				if err := _PBridge.contract.UnpackLog(event, "DepositFunds", log); err != nil {
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

// ParseDepositFunds is a log parse operation binding the contract event 0xd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88.
//
// Solidity: event DepositFunds(address from, uint256 amount)
func (_PBridge *PBridgeFilterer) ParseDepositFunds(log types.Log) (*PBridgeDepositFunds, error) {
	event := new(PBridgeDepositFunds)
	if err := _PBridge.contract.UnpackLog(event, "DepositFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PBridgeTransferFundsIterator is returned from FilterTransferFunds and is used to iterate over the raw logs and unpacked data for TransferFunds events raised by the PBridge contract.
type PBridgeTransferFundsIterator struct {
	Event *PBridgeTransferFunds // Event containing the contract specifics and raw log

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
func (it *PBridgeTransferFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PBridgeTransferFunds)
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
		it.Event = new(PBridgeTransferFunds)
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
func (it *PBridgeTransferFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PBridgeTransferFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PBridgeTransferFunds represents a TransferFunds event raised by the PBridge contract.
type PBridgeTransferFunds struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferFunds is a free log retrieval operation binding the contract event 0xc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a.
//
// Solidity: event TransferFunds(address to, uint256 amount)
func (_PBridge *PBridgeFilterer) FilterTransferFunds(opts *bind.FilterOpts) (*PBridgeTransferFundsIterator, error) {

	logs, sub, err := _PBridge.contract.FilterLogs(opts, "TransferFunds")
	if err != nil {
		return nil, err
	}
	return &PBridgeTransferFundsIterator{contract: _PBridge.contract, event: "TransferFunds", logs: logs, sub: sub}, nil
}

// WatchTransferFunds is a free log subscription operation binding the contract event 0xc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a.
//
// Solidity: event TransferFunds(address to, uint256 amount)
func (_PBridge *PBridgeFilterer) WatchTransferFunds(opts *bind.WatchOpts, sink chan<- *PBridgeTransferFunds) (event.Subscription, error) {

	logs, sub, err := _PBridge.contract.WatchLogs(opts, "TransferFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PBridgeTransferFunds)
				if err := _PBridge.contract.UnpackLog(event, "TransferFunds", log); err != nil {
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

// ParseTransferFunds is a log parse operation binding the contract event 0xc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a.
//
// Solidity: event TransferFunds(address to, uint256 amount)
func (_PBridge *PBridgeFilterer) ParseTransferFunds(log types.Log) (*PBridgeTransferFunds, error) {
	event := new(PBridgeTransferFunds)
	if err := _PBridge.contract.UnpackLog(event, "TransferFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PBridgeTxManagerChangeCompletedIterator is returned from FilterTxManagerChangeCompleted and is used to iterate over the raw logs and unpacked data for TxManagerChangeCompleted events raised by the PBridge contract.
type PBridgeTxManagerChangeCompletedIterator struct {
	Event *PBridgeTxManagerChangeCompleted // Event containing the contract specifics and raw log

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
func (it *PBridgeTxManagerChangeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PBridgeTxManagerChangeCompleted)
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
		it.Event = new(PBridgeTxManagerChangeCompleted)
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
func (it *PBridgeTxManagerChangeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PBridgeTxManagerChangeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PBridgeTxManagerChangeCompleted represents a TxManagerChangeCompleted event raised by the PBridge contract.
type PBridgeTxManagerChangeCompleted struct {
	TxKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTxManagerChangeCompleted is a free log retrieval operation binding the contract event 0xac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae.
//
// Solidity: event TxManagerChangeCompleted(string txKey)
func (_PBridge *PBridgeFilterer) FilterTxManagerChangeCompleted(opts *bind.FilterOpts) (*PBridgeTxManagerChangeCompletedIterator, error) {

	logs, sub, err := _PBridge.contract.FilterLogs(opts, "TxManagerChangeCompleted")
	if err != nil {
		return nil, err
	}
	return &PBridgeTxManagerChangeCompletedIterator{contract: _PBridge.contract, event: "TxManagerChangeCompleted", logs: logs, sub: sub}, nil
}

// WatchTxManagerChangeCompleted is a free log subscription operation binding the contract event 0xac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae.
//
// Solidity: event TxManagerChangeCompleted(string txKey)
func (_PBridge *PBridgeFilterer) WatchTxManagerChangeCompleted(opts *bind.WatchOpts, sink chan<- *PBridgeTxManagerChangeCompleted) (event.Subscription, error) {

	logs, sub, err := _PBridge.contract.WatchLogs(opts, "TxManagerChangeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PBridgeTxManagerChangeCompleted)
				if err := _PBridge.contract.UnpackLog(event, "TxManagerChangeCompleted", log); err != nil {
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

// ParseTxManagerChangeCompleted is a log parse operation binding the contract event 0xac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae.
//
// Solidity: event TxManagerChangeCompleted(string txKey)
func (_PBridge *PBridgeFilterer) ParseTxManagerChangeCompleted(log types.Log) (*PBridgeTxManagerChangeCompleted, error) {
	event := new(PBridgeTxManagerChangeCompleted)
	if err := _PBridge.contract.UnpackLog(event, "TxManagerChangeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PBridgeTxUpgradeCompletedIterator is returned from FilterTxUpgradeCompleted and is used to iterate over the raw logs and unpacked data for TxUpgradeCompleted events raised by the PBridge contract.
type PBridgeTxUpgradeCompletedIterator struct {
	Event *PBridgeTxUpgradeCompleted // Event containing the contract specifics and raw log

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
func (it *PBridgeTxUpgradeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PBridgeTxUpgradeCompleted)
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
		it.Event = new(PBridgeTxUpgradeCompleted)
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
func (it *PBridgeTxUpgradeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PBridgeTxUpgradeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PBridgeTxUpgradeCompleted represents a TxUpgradeCompleted event raised by the PBridge contract.
type PBridgeTxUpgradeCompleted struct {
	TxKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTxUpgradeCompleted is a free log retrieval operation binding the contract event 0x5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed0.
//
// Solidity: event TxUpgradeCompleted(string txKey)
func (_PBridge *PBridgeFilterer) FilterTxUpgradeCompleted(opts *bind.FilterOpts) (*PBridgeTxUpgradeCompletedIterator, error) {

	logs, sub, err := _PBridge.contract.FilterLogs(opts, "TxUpgradeCompleted")
	if err != nil {
		return nil, err
	}
	return &PBridgeTxUpgradeCompletedIterator{contract: _PBridge.contract, event: "TxUpgradeCompleted", logs: logs, sub: sub}, nil
}

// WatchTxUpgradeCompleted is a free log subscription operation binding the contract event 0x5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed0.
//
// Solidity: event TxUpgradeCompleted(string txKey)
func (_PBridge *PBridgeFilterer) WatchTxUpgradeCompleted(opts *bind.WatchOpts, sink chan<- *PBridgeTxUpgradeCompleted) (event.Subscription, error) {

	logs, sub, err := _PBridge.contract.WatchLogs(opts, "TxUpgradeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PBridgeTxUpgradeCompleted)
				if err := _PBridge.contract.UnpackLog(event, "TxUpgradeCompleted", log); err != nil {
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

// ParseTxUpgradeCompleted is a log parse operation binding the contract event 0x5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed0.
//
// Solidity: event TxUpgradeCompleted(string txKey)
func (_PBridge *PBridgeFilterer) ParseTxUpgradeCompleted(log types.Log) (*PBridgeTxUpgradeCompleted, error) {
	event := new(PBridgeTxUpgradeCompleted)
	if err := _PBridge.contract.UnpackLog(event, "TxUpgradeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PBridgeTxWithdrawCompletedIterator is returned from FilterTxWithdrawCompleted and is used to iterate over the raw logs and unpacked data for TxWithdrawCompleted events raised by the PBridge contract.
type PBridgeTxWithdrawCompletedIterator struct {
	Event *PBridgeTxWithdrawCompleted // Event containing the contract specifics and raw log

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
func (it *PBridgeTxWithdrawCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PBridgeTxWithdrawCompleted)
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
		it.Event = new(PBridgeTxWithdrawCompleted)
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
func (it *PBridgeTxWithdrawCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PBridgeTxWithdrawCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PBridgeTxWithdrawCompleted represents a TxWithdrawCompleted event raised by the PBridge contract.
type PBridgeTxWithdrawCompleted struct {
	TxKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTxWithdrawCompleted is a free log retrieval operation binding the contract event 0x8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1.
//
// Solidity: event TxWithdrawCompleted(string txKey)
func (_PBridge *PBridgeFilterer) FilterTxWithdrawCompleted(opts *bind.FilterOpts) (*PBridgeTxWithdrawCompletedIterator, error) {

	logs, sub, err := _PBridge.contract.FilterLogs(opts, "TxWithdrawCompleted")
	if err != nil {
		return nil, err
	}
	return &PBridgeTxWithdrawCompletedIterator{contract: _PBridge.contract, event: "TxWithdrawCompleted", logs: logs, sub: sub}, nil
}

// WatchTxWithdrawCompleted is a free log subscription operation binding the contract event 0x8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1.
//
// Solidity: event TxWithdrawCompleted(string txKey)
func (_PBridge *PBridgeFilterer) WatchTxWithdrawCompleted(opts *bind.WatchOpts, sink chan<- *PBridgeTxWithdrawCompleted) (event.Subscription, error) {

	logs, sub, err := _PBridge.contract.WatchLogs(opts, "TxWithdrawCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PBridgeTxWithdrawCompleted)
				if err := _PBridge.contract.UnpackLog(event, "TxWithdrawCompleted", log); err != nil {
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

// ParseTxWithdrawCompleted is a log parse operation binding the contract event 0x8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1.
//
// Solidity: event TxWithdrawCompleted(string txKey)
func (_PBridge *PBridgeFilterer) ParseTxWithdrawCompleted(log types.Log) (*PBridgeTxWithdrawCompleted, error) {
	event := new(PBridgeTxWithdrawCompleted)
	if err := _PBridge.contract.UnpackLog(event, "TxWithdrawCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820f5c15d745565617b7e7ed5c372322b0c653a4f4d35b6fcbe484e5972cef24cef0029"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058201e5b79afe55893115ffbc9d61fa8473effbf3610e9abe12d68dfcb7071550eb80029"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

