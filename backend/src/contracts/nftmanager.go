// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// NftManagerMetaData contains all meta data concerning the NftManager contract.
var NftManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CollectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"deployCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NftManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NftManagerMetaData.ABI instead.
var NftManagerABI = NftManagerMetaData.ABI

// NftManager is an auto generated Go binding around an Ethereum contract.
type NftManager struct {
	NftManagerCaller     // Read-only binding to the contract
	NftManagerTransactor // Write-only binding to the contract
	NftManagerFilterer   // Log filterer for contract events
}

// NftManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftManagerSession struct {
	Contract     *NftManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftManagerCallerSession struct {
	Contract *NftManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NftManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftManagerTransactorSession struct {
	Contract     *NftManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NftManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftManagerRaw struct {
	Contract *NftManager // Generic contract binding to access the raw methods on
}

// NftManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftManagerCallerRaw struct {
	Contract *NftManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NftManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftManagerTransactorRaw struct {
	Contract *NftManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftManager creates a new instance of NftManager, bound to a specific deployed contract.
func NewNftManager(address common.Address, backend bind.ContractBackend) (*NftManager, error) {
	contract, err := bindNftManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftManager{NftManagerCaller: NftManagerCaller{contract: contract}, NftManagerTransactor: NftManagerTransactor{contract: contract}, NftManagerFilterer: NftManagerFilterer{contract: contract}}, nil
}

// NewNftManagerCaller creates a new read-only instance of NftManager, bound to a specific deployed contract.
func NewNftManagerCaller(address common.Address, caller bind.ContractCaller) (*NftManagerCaller, error) {
	contract, err := bindNftManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftManagerCaller{contract: contract}, nil
}

// NewNftManagerTransactor creates a new write-only instance of NftManager, bound to a specific deployed contract.
func NewNftManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NftManagerTransactor, error) {
	contract, err := bindNftManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftManagerTransactor{contract: contract}, nil
}

// NewNftManagerFilterer creates a new log filterer instance of NftManager, bound to a specific deployed contract.
func NewNftManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NftManagerFilterer, error) {
	contract, err := bindNftManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftManagerFilterer{contract: contract}, nil
}

// bindNftManager binds a generic wrapper to an already deployed contract.
func bindNftManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftManager *NftManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftManager.Contract.NftManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftManager *NftManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftManager.Contract.NftManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftManager *NftManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftManager.Contract.NftManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftManager *NftManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftManager *NftManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftManager *NftManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftManager.Contract.contract.Transact(opts, method, params...)
}

// DeployCollection is a paid mutator transaction binding the contract method 0xd2b6e78d.
//
// Solidity: function deployCollection(string uri, string name, string symbol) returns()
func (_NftManager *NftManagerTransactor) DeployCollection(opts *bind.TransactOpts, uri string, name string, symbol string) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "deployCollection", uri, name, symbol)
}

// DeployCollection is a paid mutator transaction binding the contract method 0xd2b6e78d.
//
// Solidity: function deployCollection(string uri, string name, string symbol) returns()
func (_NftManager *NftManagerSession) DeployCollection(uri string, name string, symbol string) (*types.Transaction, error) {
	return _NftManager.Contract.DeployCollection(&_NftManager.TransactOpts, uri, name, symbol)
}

// DeployCollection is a paid mutator transaction binding the contract method 0xd2b6e78d.
//
// Solidity: function deployCollection(string uri, string name, string symbol) returns()
func (_NftManager *NftManagerTransactorSession) DeployCollection(uri string, name string, symbol string) (*types.Transaction, error) {
	return _NftManager.Contract.DeployCollection(&_NftManager.TransactOpts, uri, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0xc6c3bbe6.
//
// Solidity: function mint(address collection, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactor) Mint(opts *bind.TransactOpts, collection common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "mint", collection, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xc6c3bbe6.
//
// Solidity: function mint(address collection, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerSession) Mint(collection common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.Mint(&_NftManager.TransactOpts, collection, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xc6c3bbe6.
//
// Solidity: function mint(address collection, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactorSession) Mint(collection common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.Mint(&_NftManager.TransactOpts, collection, to, tokenId)
}

// NftManagerCollectionCreatedIterator is returned from FilterCollectionCreated and is used to iterate over the raw logs and unpacked data for CollectionCreated events raised by the NftManager contract.
type NftManagerCollectionCreatedIterator struct {
	Event *NftManagerCollectionCreated // Event containing the contract specifics and raw log

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
func (it *NftManagerCollectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerCollectionCreated)
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
		it.Event = new(NftManagerCollectionCreated)
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
func (it *NftManagerCollectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerCollectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerCollectionCreated represents a CollectionCreated event raised by the NftManager contract.
type NftManagerCollectionCreated struct {
	Collection common.Address
	Name       string
	Symbol     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectionCreated is a free log retrieval operation binding the contract event 0x3454b57f2dca4f5a54e8358d096ac9d1a0d2dab98991ddb89ff9ea1746260617.
//
// Solidity: event CollectionCreated(address collection, string name, string symbol)
func (_NftManager *NftManagerFilterer) FilterCollectionCreated(opts *bind.FilterOpts) (*NftManagerCollectionCreatedIterator, error) {

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "CollectionCreated")
	if err != nil {
		return nil, err
	}
	return &NftManagerCollectionCreatedIterator{contract: _NftManager.contract, event: "CollectionCreated", logs: logs, sub: sub}, nil
}

// WatchCollectionCreated is a free log subscription operation binding the contract event 0x3454b57f2dca4f5a54e8358d096ac9d1a0d2dab98991ddb89ff9ea1746260617.
//
// Solidity: event CollectionCreated(address collection, string name, string symbol)
func (_NftManager *NftManagerFilterer) WatchCollectionCreated(opts *bind.WatchOpts, sink chan<- *NftManagerCollectionCreated) (event.Subscription, error) {

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "CollectionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerCollectionCreated)
				if err := _NftManager.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
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

// ParseCollectionCreated is a log parse operation binding the contract event 0x3454b57f2dca4f5a54e8358d096ac9d1a0d2dab98991ddb89ff9ea1746260617.
//
// Solidity: event CollectionCreated(address collection, string name, string symbol)
func (_NftManager *NftManagerFilterer) ParseCollectionCreated(log types.Log) (*NftManagerCollectionCreated, error) {
	event := new(NftManagerCollectionCreated)
	if err := _NftManager.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the NftManager contract.
type NftManagerTokenMintedIterator struct {
	Event *NftManagerTokenMinted // Event containing the contract specifics and raw log

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
func (it *NftManagerTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerTokenMinted)
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
		it.Event = new(NftManagerTokenMinted)
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
func (it *NftManagerTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerTokenMinted represents a TokenMinted event raised by the NftManager contract.
type NftManagerTokenMinted struct {
	Collection common.Address
	Recipient  common.Address
	TokenId    *big.Int
	TokenURI   string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0xc9fee7cd4889f66f10ff8117316524260a5242e88e25e0656dfb3f4196a21917.
//
// Solidity: event TokenMinted(address collection, address recipient, uint256 tokenId, string tokenURI)
func (_NftManager *NftManagerFilterer) FilterTokenMinted(opts *bind.FilterOpts) (*NftManagerTokenMintedIterator, error) {

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "TokenMinted")
	if err != nil {
		return nil, err
	}
	return &NftManagerTokenMintedIterator{contract: _NftManager.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0xc9fee7cd4889f66f10ff8117316524260a5242e88e25e0656dfb3f4196a21917.
//
// Solidity: event TokenMinted(address collection, address recipient, uint256 tokenId, string tokenURI)
func (_NftManager *NftManagerFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *NftManagerTokenMinted) (event.Subscription, error) {

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "TokenMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerTokenMinted)
				if err := _NftManager.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0xc9fee7cd4889f66f10ff8117316524260a5242e88e25e0656dfb3f4196a21917.
//
// Solidity: event TokenMinted(address collection, address recipient, uint256 tokenId, string tokenURI)
func (_NftManager *NftManagerFilterer) ParseTokenMinted(log types.Log) (*NftManagerTokenMinted, error) {
	event := new(NftManagerTokenMinted)
	if err := _NftManager.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
