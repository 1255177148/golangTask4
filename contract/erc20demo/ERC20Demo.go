// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20demo

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

// Erc20demoMetaData contains all meta data concerning the Erc20demo contract.
var Erc20demoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"threeParty\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"threeParty\",\"type\":\"address\"}],\"name\":\"balanceOfApprove\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc20demoABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20demoMetaData.ABI instead.
var Erc20demoABI = Erc20demoMetaData.ABI

// Erc20demo is an auto generated Go binding around an Ethereum contract.
type Erc20demo struct {
	Erc20demoCaller     // Read-only binding to the contract
	Erc20demoTransactor // Write-only binding to the contract
	Erc20demoFilterer   // Log filterer for contract events
}

// Erc20demoCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20demoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20demoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20demoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20demoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20demoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20demoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20demoSession struct {
	Contract     *Erc20demo        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20demoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20demoCallerSession struct {
	Contract *Erc20demoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Erc20demoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20demoTransactorSession struct {
	Contract     *Erc20demoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Erc20demoRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20demoRaw struct {
	Contract *Erc20demo // Generic contract binding to access the raw methods on
}

// Erc20demoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20demoCallerRaw struct {
	Contract *Erc20demoCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20demoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20demoTransactorRaw struct {
	Contract *Erc20demoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20demo creates a new instance of Erc20demo, bound to a specific deployed contract.
func NewErc20demo(address common.Address, backend bind.ContractBackend) (*Erc20demo, error) {
	contract, err := bindErc20demo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20demo{Erc20demoCaller: Erc20demoCaller{contract: contract}, Erc20demoTransactor: Erc20demoTransactor{contract: contract}, Erc20demoFilterer: Erc20demoFilterer{contract: contract}}, nil
}

// NewErc20demoCaller creates a new read-only instance of Erc20demo, bound to a specific deployed contract.
func NewErc20demoCaller(address common.Address, caller bind.ContractCaller) (*Erc20demoCaller, error) {
	contract, err := bindErc20demo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20demoCaller{contract: contract}, nil
}

// NewErc20demoTransactor creates a new write-only instance of Erc20demo, bound to a specific deployed contract.
func NewErc20demoTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20demoTransactor, error) {
	contract, err := bindErc20demo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20demoTransactor{contract: contract}, nil
}

// NewErc20demoFilterer creates a new log filterer instance of Erc20demo, bound to a specific deployed contract.
func NewErc20demoFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20demoFilterer, error) {
	contract, err := bindErc20demo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20demoFilterer{contract: contract}, nil
}

// bindErc20demo binds a generic wrapper to an already deployed contract.
func bindErc20demo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20demoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20demo *Erc20demoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20demo.Contract.Erc20demoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20demo *Erc20demoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20demo.Contract.Erc20demoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20demo *Erc20demoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20demo.Contract.Erc20demoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20demo *Erc20demoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20demo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20demo *Erc20demoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20demo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20demo *Erc20demoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20demo.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20demo *Erc20demoCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20demo.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20demo *Erc20demoSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20demo.Contract.BalanceOf(&_Erc20demo.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20demo *Erc20demoCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20demo.Contract.BalanceOf(&_Erc20demo.CallOpts, account)
}

// BalanceOfApprove is a free data retrieval call binding the contract method 0xde53131f.
//
// Solidity: function balanceOfApprove(address threeParty) view returns(uint256)
func (_Erc20demo *Erc20demoCaller) BalanceOfApprove(opts *bind.CallOpts, threeParty common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20demo.contract.Call(opts, &out, "balanceOfApprove", threeParty)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfApprove is a free data retrieval call binding the contract method 0xde53131f.
//
// Solidity: function balanceOfApprove(address threeParty) view returns(uint256)
func (_Erc20demo *Erc20demoSession) BalanceOfApprove(threeParty common.Address) (*big.Int, error) {
	return _Erc20demo.Contract.BalanceOfApprove(&_Erc20demo.CallOpts, threeParty)
}

// BalanceOfApprove is a free data retrieval call binding the contract method 0xde53131f.
//
// Solidity: function balanceOfApprove(address threeParty) view returns(uint256)
func (_Erc20demo *Erc20demoCallerSession) BalanceOfApprove(threeParty common.Address) (*big.Int, error) {
	return _Erc20demo.Contract.BalanceOfApprove(&_Erc20demo.CallOpts, threeParty)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address threeParty, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactor) Approve(opts *bind.TransactOpts, threeParty common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.contract.Transact(opts, "approve", threeParty, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address threeParty, uint256 amount) returns()
func (_Erc20demo *Erc20demoSession) Approve(threeParty common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.Approve(&_Erc20demo.TransactOpts, threeParty, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address threeParty, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactorSession) Approve(threeParty common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.Approve(&_Erc20demo.TransactOpts, threeParty, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20demo *Erc20demoSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.Mint(&_Erc20demo.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.Mint(&_Erc20demo.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Erc20demo *Erc20demoSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.Transfer(&_Erc20demo.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.Transfer(&_Erc20demo.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns()
func (_Erc20demo *Erc20demoSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.TransferFrom(&_Erc20demo.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns()
func (_Erc20demo *Erc20demoTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20demo.Contract.TransferFrom(&_Erc20demo.TransactOpts, from, to, amount)
}

// Erc20demoApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the Erc20demo contract.
type Erc20demoApproveIterator struct {
	Event *Erc20demoApprove // Event containing the contract specifics and raw log

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
func (it *Erc20demoApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20demoApprove)
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
		it.Event = new(Erc20demoApprove)
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
func (it *Erc20demoApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20demoApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20demoApprove represents a Approve event raised by the Erc20demo contract.
type Erc20demoApprove struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address owner, address spender, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) FilterApprove(opts *bind.FilterOpts) (*Erc20demoApproveIterator, error) {

	logs, sub, err := _Erc20demo.contract.FilterLogs(opts, "Approve")
	if err != nil {
		return nil, err
	}
	return &Erc20demoApproveIterator{contract: _Erc20demo.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address owner, address spender, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *Erc20demoApprove) (event.Subscription, error) {

	logs, sub, err := _Erc20demo.contract.WatchLogs(opts, "Approve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20demoApprove)
				if err := _Erc20demo.contract.UnpackLog(event, "Approve", log); err != nil {
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

// ParseApprove is a log parse operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address owner, address spender, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) ParseApprove(log types.Log) (*Erc20demoApprove, error) {
	event := new(Erc20demoApprove)
	if err := _Erc20demo.contract.UnpackLog(event, "Approve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20demoMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Erc20demo contract.
type Erc20demoMintIterator struct {
	Event *Erc20demoMint // Event containing the contract specifics and raw log

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
func (it *Erc20demoMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20demoMint)
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
		it.Event = new(Erc20demoMint)
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
func (it *Erc20demoMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20demoMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20demoMint represents a Mint event raised by the Erc20demo contract.
type Erc20demoMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address to, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) FilterMint(opts *bind.FilterOpts) (*Erc20demoMintIterator, error) {

	logs, sub, err := _Erc20demo.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &Erc20demoMintIterator{contract: _Erc20demo.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address to, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Erc20demoMint) (event.Subscription, error) {

	logs, sub, err := _Erc20demo.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20demoMint)
				if err := _Erc20demo.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address to, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) ParseMint(log types.Log) (*Erc20demoMint, error) {
	event := new(Erc20demoMint)
	if err := _Erc20demo.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20demoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20demo contract.
type Erc20demoTransferIterator struct {
	Event *Erc20demoTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20demoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20demoTransfer)
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
		it.Event = new(Erc20demoTransfer)
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
func (it *Erc20demoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20demoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20demoTransfer represents a Transfer event raised by the Erc20demo contract.
type Erc20demoTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) FilterTransfer(opts *bind.FilterOpts) (*Erc20demoTransferIterator, error) {

	logs, sub, err := _Erc20demo.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &Erc20demoTransferIterator{contract: _Erc20demo.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20demoTransfer) (event.Subscription, error) {

	logs, sub, err := _Erc20demo.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20demoTransfer)
				if err := _Erc20demo.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address from, address to, uint256 amount)
func (_Erc20demo *Erc20demoFilterer) ParseTransfer(log types.Log) (*Erc20demoTransfer, error) {
	event := new(Erc20demoTransfer)
	if err := _Erc20demo.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
