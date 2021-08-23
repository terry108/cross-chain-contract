// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pbridge_abi

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
var AddressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c39467608b72d62d89c95a8ecc1560b3a6bc693ef8cf102afcc79f832174aa4c64736f6c63430008000033"

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
var BytesLibBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206712bce07d4dd7d8d90fe447e771e047afc90e17d76438f87693f703d5bfd55064736f6c63430008000033"

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
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
const IERC20MinterABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"replaceMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
const PBridgeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_managers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"CrossOutFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxManagerChangeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxUpgradeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxWithdrawCompleted\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"allManagers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignManagerChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"upgradeContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isERC20\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"crossOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_min_signatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"ifManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"isCompletedTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"isMinterERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_managers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"min_managers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"registerMinterERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"unregisterMinterERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContractS1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"upgradeContractS2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

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
var PBridgeBin = "0x6080604052600080546001600160a81b0319169055600f600155600360028190556042905560416004553480156200003657600080fd5b5060405162003a4938038062003a49833981016040819052620000599162000391565b60015481511115620000885760405162461bcd60e51b81526004016200007f90620004dd565b60405180910390fd5b60025481511015620000ae5760405162461bcd60e51b81526004016200007f906200045f565b60058054610100600160a81b03191633610100021790558051620000da906009906020840190620002f8565b5060005b60095460ff82161015620002215760016008600060098460ff16815481106200011757634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff191660ff9384161790556009805460019360069392919086169081106200017757634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191660ff928316179055600980546007928416908110620001d157634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b03909216919091179055806200021881620005e2565b915050620000de565b5060055461010090046001600160a01b031660009081526008602052604090205460ff1615620002655760405162461bcd60e51b81526004016200007f9062000524565b60095462000273906200028e565b6005805460ff191660ff929092169190911790555062000631565b6000808211620002b25760405162461bcd60e51b81526004016200007f90620004a6565b60006001606484600354620002c89190620005a6565b620002d491906200056a565b620002e09190620005c8565b9050620002ef60648262000585565b9150505b919050565b82805482825590600052602060002090810192821562000350579160200282015b828111156200035057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000319565b506200035e92915062000362565b5090565b5b808211156200035e576000815560010162000363565b80516001600160a01b0381168114620002f357600080fd5b60006020808385031215620003a4578182fd5b82516001600160401b0380821115620003bb578384fd5b818501915085601f830112620003cf578384fd5b815181811115620003e457620003e46200061b565b838102604051858282010181811085821117156200040657620004066200061b565b604052828152858101935084860182860187018a101562000425578788fd5b8795505b8386101562000452576200043d8162000379565b85526001959095019493860193860162000429565b5098975050505050505050565b60208082526027908201527f4e6f74207265616368696e6720746865206d696e206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526014908201527f4d616e616765722043616e277420656d7074792e000000000000000000000000604082015260600190565b60208082526027908201527f457863656564656420746865206d6178696d756d206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526026908201527f436f6e74726163742063726561746f722063616e6e6f7420616374206173206d60408201526530b730b3b2b960d11b606082015260800190565b6000821982111562000580576200058062000605565b500190565b600082620005a157634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615620005c357620005c362000605565b500290565b600082821015620005dd57620005dd62000605565b500390565b600060ff821660ff811415620005fc57620005fc62000605565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b61340880620006416000396000f3fe60806040526004361061012d5760003560e01c806375173b70116100ab578063ab6c2b101161006f578063ab6c2b101461033d578063ad4b61a81461035d578063d4cacbaa14610372578063d55ec69714610387578063e079cee91461039c578063f7f2ff74146103be5761016d565b806375173b70146102a65780638da5cb5b146102c65780639c30b35e146102db5780639dcdc978146102fd578063a5e399b31461031d5761016d565b806330b2d84d116100f257806330b2d84d1461020457806334774b7114610226578063408e8b7a146102465780635bda3fcf146102665780636a7142e1146102865761016d565b80627192261461016f5780630889d1f01461018f5780631b9a9323146101b85780631dda9c05146101da5780632c4e722e146101ef5761016d565b3661016d577fd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b883334604051610163929190612765565b60405180910390a1005b005b34801561017b57600080fd5b5061016d61018a366004612491565b6103d3565b6101a261019d366004612553565b6105ad565b6040516101af9190612844565b60405180910390f35b3480156101c457600080fd5b506101cd610864565b6040516101af9190613248565b3480156101e657600080fd5b5061016d61086a565b3480156101fb57600080fd5b506101cd61092c565b34801561021057600080fd5b50610219610932565b6040516101af9190612751565b34801561023257600080fd5b5061016d61024136600461230a565b610946565b34801561025257600080fd5b5061016d61026136600461241c565b610a16565b34801561027257600080fd5b5061016d61028136600461230a565b610bf5565b34801561029257600080fd5b506101a26102a136600461230a565b610deb565b3480156102b257600080fd5b506101a26102c136600461230a565b610e0f565b3480156102d257600080fd5b50610219610e30565b3480156102e757600080fd5b506102f0610e44565b6040516101af91906127f7565b34801561030957600080fd5b5061016d61031836600461230a565b610ea6565b34801561032957600080fd5b506101a2610338366004612342565b610f1b565b34801561034957600080fd5b5061016d610358366004612375565b610f49565b34801561036957600080fd5b506101cd6111d7565b34801561037e57600080fd5b5061016d6111dd565b34801561039357600080fd5b506101a261123a565b3480156103a857600080fd5b506103b1611243565b6040516101af9190613251565b3480156103ca57600080fd5b506101cd61124c565b3360009081526008602052604090205460ff1660011461040e5760405162461bcd60e51b815260040161040590612b45565b60405180910390fd5b845160401461042f5760405162461bcd60e51b815260040161040590613020565b600084511180610440575060008351115b61045c5760405162461bcd60e51b815260040161040590612f73565b600b8560405161046c919061262b565b9081526040519081900360200190205460ff161561049c5760405162461bcd60e51b815260040161040590612a25565b6104a68484611252565b60008585848660026040516020016104c29594939291906126fe565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff161561050a5760405162461bcd60e51b815260040161040590612951565b6105148183611495565b6105305760405162461bcd60e51b815260040161040590612f44565b610539846114d8565b610542856116d3565b60095461054e906117b9565b6005805460ff191660ff9290921691909117905561056e8682600161180f565b7fac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae8660405161059d919061286d565b60405180910390a1505050505050565b600033836105cd5760405162461bcd60e51b815260040161040590612c84565b6001600160a01b038316156107fb5734156105fa5760405162461bcd60e51b815260040161040590612ad9565b61060c836001600160a01b031661185c565b6106285760405162461bcd60e51b815260040161040590613203565b604051636eb1769f60e11b815283906000906001600160a01b0383169063dd62ed3e9061065b908690309060040161277e565b60206040518083038186803b15801561067357600080fd5b505afa158015610687573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ab91906125ac565b9050858110156106cd5760405162461bcd60e51b815260040161040590612df3565b6040516370a0823160e01b81526000906001600160a01b038416906370a08231906106fc908790600401612751565b60206040518083038186803b15801561071457600080fd5b505afa158015610728573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074c91906125ac565b90508681101561076e5760405162461bcd60e51b815260040161040590612f0d565b6107836001600160a01b03841685308a611862565b61078c86610deb565b156107f357604051630852cd8d60e31b815286906001600160a01b038216906342966c68906107bf908b90600401613248565b600060405180830381600087803b1580156107d957600080fd5b505af11580156107ed573d6000803e3d6000fd5b50505050505b50505061081a565b83341461081a5760405162461bcd60e51b815260040161040590612fbb565b7f5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e306671468186868660405161084f94939291906127bc565b60405180910390a160019150505b9392505050565b60045481565b60055461010090046001600160a01b031633146108995760405162461bcd60e51b8152600401610405906130b9565b60005460ff166108bb5760405162461bcd60e51b815260040161040590612d21565b60005461010090046001600160a01b03166108e85760405162461bcd60e51b815260040161040590612e35565b600080546040516001600160a01b0361010090920491909116914780156108fc02929091818181858888f19350505050158015610929573d6000803e3d6000fd5b50565b60035481565b60005461010090046001600160a01b031681565b60055461010090046001600160a01b031633146109755760405162461bcd60e51b8152600401610405906130b9565b306001600160a01b038216141561099e5760405162461bcd60e51b815260040161040590612c54565b6109b0816001600160a01b031661185c565b6109cc5760405162461bcd60e51b815260040161040590613203565b6109d581610deb565b156109f25760405162461bcd60e51b8152600401610405906128c3565b6001600160a01b03166000908152600c60205260409020805460ff19166001179055565b3360009081526008602052604090205460ff16600114610a485760405162461bcd60e51b815260040161040590612b45565b8251604014610a695760405162461bcd60e51b815260040161040590613020565b600b83604051610a79919061262b565b9081526040519081900360200190205460ff1615610aa95760405162461bcd60e51b815260040161040590612a25565b60005460ff1615610acc5760405162461bcd60e51b815260040161040590612ff2565b610ade826001600160a01b031661185c565b610afa5760405162461bcd60e51b815260040161040590613203565b600083836002604051602001610b12939291906126b2565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff1615610b5a5760405162461bcd60e51b815260040161040590612951565b610b648183611495565b610b805760405162461bcd60e51b815260040161040590612f44565b60008054600160ff199091168117610100600160a81b0319166101006001600160a01b0387160217909155610bb8908590839061180f565b7f5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed084604051610be7919061286d565b60405180910390a150505050565b60055461010090046001600160a01b03163314610c245760405162461bcd60e51b8152600401610405906130b9565b60005460ff16610c465760405162461bcd60e51b815260040161040590612d21565b60005461010090046001600160a01b0316610c735760405162461bcd60e51b815260040161040590612e35565b306001600160a01b0382161415610c9c5760405162461bcd60e51b815260040161040590612c54565b610cae816001600160a01b031661185c565b610cca5760405162461bcd60e51b815260040161040590613203565b6040516370a0823160e01b815281906000906001600160a01b038316906370a0823190610cfb903090600401612751565b60206040518083038186803b158015610d1357600080fd5b505afa158015610d27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4b91906125ac565b9050600054610d6c906001600160a01b0384811691610100900416836118ba565b610d7583610deb565b15610de6576000546040516301fc6bd160e21b815284916001600160a01b03808416926307f1af4492610db2926101009091041690600401612751565b600060405180830381600087803b158015610dcc57600080fd5b505af1158015610de0573d6000803e3d6000fd5b50505050505b505050565b6001600160a01b0381166000908152600c602052604090205460ff1615155b919050565b6001600160a01b031660009081526008602052604090205460ff1660011490565b60055461010090046001600160a01b031681565b60606009805480602002602001604051908101604052809291908181526020018280548015610e9c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e7e575b5050505050905090565b60055461010090046001600160a01b03163314610ed55760405162461bcd60e51b8152600401610405906130b9565b610ede81610deb565b610efa5760405162461bcd60e51b815260040161040590612ed6565b6001600160a01b03166000908152600c60205260409020805460ff19169055565b600080600b83604051610f2e919061262b565b9081526040519081900360200190205460ff16119050919050565b3360009081526008602052604090205460ff16600114610f7b5760405162461bcd60e51b815260040161040590612b45565b8551604014610f9c5760405162461bcd60e51b815260040161040590613020565b6001600160a01b038516610fc25760405162461bcd60e51b81526004016104059061290b565b60008411610fe25760405162461bcd60e51b815260040161040590612dab565b600b86604051610ff2919061262b565b9081526040519081900360200190205460ff16156110225760405162461bcd60e51b815260040161040590612a25565b8215611038576110338286866118d9565b611058565b834710156110585760405162461bcd60e51b815260040161040590612b7c565b60008686868686600260405160200161107696959493929190612647565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff16156110be5760405162461bcd60e51b815260040161040590612951565b6110c88183611495565b6110e45760405162461bcd60e51b815260040161040590612f44565b83156110fa576110f5838787611a13565b61118b565b8447101561111a5760405162461bcd60e51b815260040161040590612b7c565b6040516001600160a01b0387169086156108fc029087906000818181858888f19350505050158015611150573d6000803e3d6000fd5b507fc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a8686604051611182929190612765565b60405180910390a15b6111978782600161180f565b7f8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1876040516111c6919061286d565b60405180910390a150505050505050565b60025481565b60055461010090046001600160a01b0316331461120c5760405162461bcd60e51b8152600401610405906130b9565b60005460ff1661122e5760405162461bcd60e51b815260040161040590612d21565b6000805460ff19169055565b60005460ff1681565b60055460ff1681565b60015481565b815160005b8181101561130a57600084828151811061128157634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b031614156112be5760405162461bcd60e51b815260040161040590612c10565b6001600160a01b03811660009081526008602052604090205460ff16156112f75760405162461bcd60e51b815260040161040590612e78565b508061130281613348565b915050611257565b5061131483611b40565b6113305760405162461bcd60e51b81526004016104059061297d565b60055461134b9061010090046001600160a01b031684611c58565b6113675760405162461bcd60e51b8152600401610405906131bd565b61137082611b40565b61138c5760405162461bcd60e51b8152600401610405906130f0565b815160005b818110156114525760008482815181106113bb57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600690925260409091205490915060ff16156114045760405162461bcd60e51b81526004016104059061313c565b6001600160a01b03811660009081526008602052604090205460ff1660011461143f5760405162461bcd60e51b815260040161040590612d41565b508061144a81613348565b915050611391565b50600154835185516009546114679190613289565b6114719190613305565b111561148f5760405162461bcd60e51b815260040161040590612cda565b50505050565b60006103cf825111156114ba5760405162461bcd60e51b815260040161040590612bd9565b60006114c68484611cf3565b60055460ff1611159150505b92915050565b80516114e357610929565b60005b8151811015611550576008600083838151811061151357634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191690558061154881613348565b9150506114e6565b5060005b6009548110156115fc57600860006009838154811061158357634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff166115ea57600981815481106115d057634e487b7160e01b600052603260045260246000fd5b600091825260209091200180546001600160a01b03191690555b806115f481613348565b915050611554565b50601060005b600954811015610de65760006009828154811061162f57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b031690508061165e578260101415611658578192505b506116c1565b826010146116bf57806009848154811061168857634e487b7160e01b600052603260045260246000fd5b600091825260209091200180546001600160a01b0319166001600160a01b0392909216919091179055826116bb81613348565b9350505b505b806116cb81613348565b915050611602565b80516116de57610929565b60005b81518110156117b557600082828151811061170c57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600890925260409091205490915060ff166117a2576001600160a01b0381166000818152600860205260408120805460ff191660019081179091556009805491820181559091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03191690911790555b50806117ad81613348565b9150506116e1565b5050565b60008082116117da5760405162461bcd60e51b8152600401610405906129c9565b600060016064846003546117ee91906132e6565b6117f89190613289565b6118029190613305565b905061085d6064826132c6565b80600b84604051611820919061262b565b9081526040805160209281900383019020805460ff1990811660ff958616179091556000958652600a9092529093208054909316911617905550565b3b151590565b61148f846323b872dd60e01b85858560405160240161188393929190612798565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611e9f565b610de68363a9059cbb60e01b8484604051602401611883929190612765565b6001600160a01b0382166118ff5760405162461bcd60e51b815260040161040590612880565b306001600160a01b03841614156119285760405162461bcd60e51b815260040161040590612c54565b61193a836001600160a01b031661185c565b6119565760405162461bcd60e51b815260040161040590613203565b61195f83610deb565b1561196957610de6565b6040516370a0823160e01b815283906000906001600160a01b038316906370a082319061199a903090600401612751565b60206040518083038186803b1580156119b257600080fd5b505afa1580156119c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119ea91906125ac565b905082811015611a0c5760405162461bcd60e51b815260040161040590612a5c565b5050505050565b611a1c83610deb565b15611a89576040516340c10f1960e01b815283906001600160a01b038216906340c10f1990611a519086908690600401612765565b600060405180830381600087803b158015611a6b57600080fd5b505af1158015611a7f573d6000803e3d6000fd5b5050505050610de6565b6040516370a0823160e01b815283906000906001600160a01b038316906370a0823190611aba903090600401612751565b60206040518083038186803b158015611ad257600080fd5b505afa158015611ae6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b0a91906125ac565b905082811015611b2c5760405162461bcd60e51b815260040161040590612a5c565b611a0c6001600160a01b03831685856118ba565b6000805b8251811015611c4f576000838281518110611b6f57634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b03161415611b9a5750611c4f565b6000611ba7836001613289565b90505b8451811015611c3a576000858281518110611bd557634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b03161415611c005750611c3a565b806001600160a01b0316836001600160a01b03161415611c27576000945050505050610e0a565b5080611c3281613348565b915050611baa565b50508080611c4790613348565b915050611b44565b50600192915050565b60008060005b8351811015611ce857838181518110611c8757634e487b7160e01b600052603260045260246000fd5b6020026020010151915060006001600160a01b0316826001600160a01b03161415611cb157611ce8565b846001600160a01b0316826001600160a01b03161415611cd6576000925050506114d2565b80611ce081613348565b915050611c5e565b506001949350505050565b600454815160009182918291611d099190611f2e565b905060008167ffffffffffffffff811115611d3457634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611d5d578160200160208202803683370190505b50905060008060005b84811015611e63576000611d87846004548b611f3a9092919063ffffffff16565b90506000611d958b83612002565b90506001600160a01b038116611dbd5760405162461bcd60e51b815260040161040590612cb0565b6001600160a01b03811660009081526008602052604090205460ff1660011415611e3f5787611deb81613348565b98505080868580611dfb90613363565b965060ff1681518110611e1e57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b600454611e4c9086613289565b945050508080611e5b90613348565b915050611d66565b506000611e6f84611b40565b90506060935080611e925760405162461bcd60e51b8152600401610405906129f7565b5093979650505050505050565b6000611ef4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166121049092919063ffffffff16565b805190915015610de65780806020019051810190611f129190612326565b610de65760405162461bcd60e51b815260040161040590613173565b600061085d82846132c6565b606081611f4881601f613289565b1015611f665760405162461bcd60e51b815260040161040590612b1d565b611f708284613289565b84511015611f905760405162461bcd60e51b81526004016104059061308e565b606082158015611faf5760405191506000825260208201604052611ff9565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015611fe8578051835260209283019201611fd0565b5050858452601f01601f1916604052505b50949350505050565b60008060008060045485511461201e57600093505050506114d2565b50505060208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561206757600093505050506114d2565b601b8160ff1610156120815761207e601b826132a1565b90505b8060ff16601b1415801561209957508060ff16601c14155b156120aa57600093505050506114d2565b600186828585604051600081526020016040526040516120cd949392919061284f565b6020604051602081039080840390855afa1580156120ef573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6060612113848460008561211b565b949350505050565b60608247101561213d5760405162461bcd60e51b815260040161040590612a93565b6121468561185c565b6121625760405162461bcd60e51b815260040161040590613057565b600080866001600160a01b0316858760405161217e919061262b565b60006040518083038185875af1925050503d80600081146121bb576040519150601f19603f3d011682016040523d82523d6000602084013e6121c0565b606091505b50915091506121d08282866121db565b979650505050505050565b606083156121ea57508161085d565b8251156121fa5782518084602001fd5b8160405162461bcd60e51b8152600401610405919061286d565b600082601f830112612224578081fd5b8135602067ffffffffffffffff82111561224057612240613399565b80820261224e82820161325f565b838152828101908684018388018501891015612268578687fd5b8693505b8584101561229357803561227f816133af565b83526001939093019291840191840161226c565b50979650505050505050565b600082601f8301126122af578081fd5b813567ffffffffffffffff8111156122c9576122c9613399565b6122dc601f8201601f191660200161325f565b8181528460208386010111156122f0578283fd5b816020850160208301379081016020019190915292915050565b60006020828403121561231b578081fd5b813561085d816133af565b600060208284031215612337578081fd5b815161085d816133c4565b600060208284031215612353578081fd5b813567ffffffffffffffff811115612369578182fd5b6121138482850161229f565b60008060008060008060c0878903121561238d578182fd5b863567ffffffffffffffff808211156123a4578384fd5b6123b08a838b0161229f565b9750602089013591506123c2826133af565b909550604088013594506060880135906123db826133c4565b9093506080880135906123ed826133af565b90925060a08801359080821115612402578283fd5b5061240f89828a0161229f565b9150509295509295509295565b600080600060608486031215612430578283fd5b833567ffffffffffffffff80821115612447578485fd5b6124538783880161229f565b945060208601359150612465826133af565b9092506040850135908082111561247a578283fd5b506124878682870161229f565b9150509250925092565b600080600080600060a086880312156124a8578081fd5b853567ffffffffffffffff808211156124bf578283fd5b6124cb89838a0161229f565b965060208801359150808211156124e0578283fd5b6124ec89838a01612214565b95506040880135915080821115612501578283fd5b61250d89838a01612214565b94506060880135915060ff82168214612524578283fd5b90925060808701359080821115612539578283fd5b506125468882890161229f565b9150509295509295909350565b600080600060608486031215612567578283fd5b833567ffffffffffffffff81111561257d578384fd5b6125898682870161229f565b9350506020840135915060408401356125a1816133af565b809150509250925092565b6000602082840312156125bd578081fd5b5051919050565b600081516020808401835b838110156125f45781516001600160a01b0316875295820195908201906001016125cf565b509495945050505050565b6000815180845261261781602086016020860161331c565b601f01601f19169290920160200192915050565b6000825161263d81846020870161331c565b9190910192915050565b60008751612659818460208c0161331c565b606097881b6bffffffffffffffffffffffff199081169390910192835260148301969096525092151560f890811b60348501529190941b909216603582015291901b6001600160f81b0319166049820152604a01919050565b600084516126c481846020890161331c565b60609490941b6bffffffffffffffffffffffff19169190930190815260f89190911b6001600160f81b031916601482015260150192915050565b60008651612710818460208b0161331c565b61271c818401886125c4565b905060ff60f81b808760f81b16825261273860018301876125c4565b60f89590951b1684525050600190910195945050505050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600060018060a01b038087168352608060208401526127de60808401876125ff565b6040840195909552929092166060909101525092915050565b6020808252825182820181905260009190848201906040850190845b818110156128385783516001600160a01b031683529284019291840191600101612813565b50909695505050505050565b901515815260200190565b93845260ff9290921660208401526040830152606082015260800190565b60006020825261085d60208301846125ff565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526028908201527f5468697320616464726573732068617320616c7265616479206265656e20726560408201526719da5cdd195c995960c21b606082015260800190565b60208082526026908201527f57697468647261773a207472616e7366657220746f20746865207a65726f206160408201526564647265737360d01b606082015260800190565b602080825260129082015271496e76616c6964207369676e61747572657360701b604082015260600190565b6020808252602c908201527f4475706c696361746520706172616d657465727320666f72207468652061646460408201526b3932b9b9903a37903537b4b760a11b606082015260800190565b60208082526014908201527326b0b730b3b2b91021b0b713ba1032b6b83a3c9760611b604082015260600190565b6020808252601490820152735369676e617475726573206475706c696361746560601b604082015260600190565b6020808252601e908201527f5472616e73616374696f6e20686173206265656e20636f6d706c657465640000604082015260600190565b6020808252601a908201527f4e6f20656e6f7567682062616c616e6365206f6620746f6b656e000000000000604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b60208082526024908201527f45524332303a20446f6573206e6f742061636365707420457468657265756d2060408201526321b7b4b760e11b606082015260800190565b6020808252600e908201526d736c6963655f6f766572666c6f7760901b604082015260600190565b6020808252601b908201527f4f6e6c79206d616e616765722063616e20657865637574652069740000000000604082015260600190565b6020808252603f908201527f5468697320636f6e7472616374206164647265737320646f6573206e6f74206860408201527f6176652073756666696369656e742062616c616e6365206f6620657468657200606082015260800190565b6020808252601d908201527f4d6178206c656e677468206f66207369676e6174757265733a20393735000000604082015260600190565b60208082526024908201527f4552524f523a204465746563746564207a65726f206164647265737320696e206040820152636164647360e01b606082015260800190565b6020808252601690820152752237903737ba3434b73390313c903cb7bab939b2b63360511b604082015260600190565b60208082526012908201527111549493d48e8816995c9bc8185b5bdd5b9d60721b604082015260600190565b60208082526010908201526f29b4b3b730ba3ab932b99032b93937b960811b604082015260600190565b60208082526027908201527f457863656564656420746865206d6178696d756d206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526006908201526511195b9a595960d21b604082015260600190565b60208082526044908201527f5468657265206172652061646472657373657320696e2074686520657869746960408201527f6e672061646472657373206c697374207468617420617265206e6f74206d616e60608201526330b3b2b960e11b608082015260a00190565b60208082526028908201527f5769746864726177616c20616d6f756e74206d75737420626520677265617465604082015267072207468616e20360c41b606082015260800190565b60208082526022908201527f4e6f20656e6f75676820616d6f756e7420666f7220617574686f72697a61746960408201526137b760f11b606082015260800190565b60208082526023908201527f4552524f523a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b602080825260409082018190527f5468652061646472657373206c6973742074686174206973206265696e672061908201527f6464656420616c7265616479206578697374732061732061206d616e61676572606082015260800190565b6020808252601e908201527f546869732061646472657373206973206e6f7420726567697374657265640000604082015260600190565b6020808252601e908201527f4e6f20656e6f7567682062616c616e6365206f662074686520746f6b656e0000604082015260600190565b60208082526015908201527415985b1a59081cda59db985d1d5c995cc819985a5b605a1b604082015260600190565b60208082526028908201527f546865726520617265206e6f206d616e6167657273206a6f696e696e67206f726040820152672065786974696e6760c01b606082015260800190565b6020808252601d908201527f496e636f6e73697374656e637920457468657265756d20616d6f756e74000000604082015260600190565b602080825260149082015273125d081a185cc81899595b881d5c19dc9859195960621b604082015260600190565b60208082526019908201527f4669786564206c656e677468206f662074784b65793a20363400000000000000604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b602080825260119082015270736c6963655f6f75744f66426f756e647360781b604082015260600190565b60208082526019908201527f4f6e6c79206f776e65722063616e206578656375746520697400000000000000604082015260600190565b6020808252602c908201527f4475706c696361746520706172616d657465727320666f72207468652061646460408201526b1c995cdcc81d1bc8195e1a5d60a21b606082015260800190565b60208082526017908201527f43616e277420657869742073656564206d616e61676572000000000000000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b60208082526026908201527f436f6e74726163742063726561746f722063616e6e6f7420616374206173206d60408201526530b730b3b2b960d11b606082015260800190565b60208082526025908201527f5468652061646472657373206973206e6f74206120636f6e7472616374206164604082015264647265737360d81b606082015260800190565b90815260200190565b60ff91909116815260200190565b60405181810167ffffffffffffffff8111828210171561328157613281613399565b604052919050565b6000821982111561329c5761329c613383565b500190565b600060ff821660ff84168060ff038211156132be576132be613383565b019392505050565b6000826132e157634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561330057613300613383565b500290565b60008282101561331757613317613383565b500390565b60005b8381101561333757818101518382015260200161331f565b8381111561148f5750506000910152565b600060001982141561335c5761335c613383565b5060010190565b600060ff821660ff81141561337a5761337a613383565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461092957600080fd5b801515811461092957600080fdfea2646970667358221220169b05408837aff5e022cd91393054d2049e1632d2bd0488c7a678612d70e05464736f6c63430008000033"

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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PBridge *PBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PBridge *PBridgeSession) Receive() (*types.Transaction, error) {
	return _PBridge.Contract.Receive(&_PBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PBridge *PBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _PBridge.Contract.Receive(&_PBridge.TransactOpts)
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
var SafeERC20Bin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bfe0e17946b2631ade7138f7be59baf2d1bdacab8b9d06ae3a842c573a8d296e64736f6c63430008000033"

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
var SafeMathBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e43b08315b2b41ed63057bea9187ab76aa5a56bc0de31b198a676f0badffb67264736f6c63430008000033"

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

