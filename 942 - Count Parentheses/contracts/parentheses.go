// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ParenthesisABI is the input ABI used to generate the binding from.
const ParenthesisABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"countWrong\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ParenthesisFuncSigs maps the 4-byte function signature to its string representation.
var ParenthesisFuncSigs = map[string]string{
	"247de6d4": "countWrong(string)",
}

// ParenthesisBin is the compiled bytecode used for deploying new contracts.
var ParenthesisBin = "0x608060405234801561001057600080fd5b5061030b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063247de6d414610030575b600080fd5b61004361003e366004610182565b610055565b60405190815260200160405180910390f35b600081600560fb1b602960f81b8380805b855181101561015657846001600160f81b03191686828151811061008c5761008c610233565b01602001516001600160f81b031916036100b257826100aa8161025f565b935050610144565b836001600160f81b0319168682815181106100cf576100cf610233565b01602001516001600160f81b0319161480156100eb5750600083135b156100fa57826100aa8161027e565b836001600160f81b03191686828151811061011757610117610233565b01602001516001600160f81b031916148015610131575082155b1561014457816101408161025f565b9250505b8061014e8161029b565b915050610066565b5061016182826102ad565b979650505050505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561019457600080fd5b813567ffffffffffffffff808211156101ac57600080fd5b818401915084601f8301126101c057600080fd5b8135818111156101d2576101d261016c565b604051601f8201601f19908116603f011681019083821181831017156101fa576101fa61016c565b8160405282815287602084870101111561021357600080fd5b826020860160208301376000928101602001929092525095945050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160ff1b01820161027757610277610249565b5060010190565b6000600160ff1b820161029357610293610249565b506000190190565b60006001820161027757610277610249565b80820182811260008312801582168215821617156102cd576102cd610249565b50509291505056fea26469706673582212209d851abf23fe037c031136b00eba2be188aaab753eb8175797d7bae5ce43f99a64736f6c63430008100033"

// DeployParenthesis deploys a new Ethereum contract, binding an instance of Parenthesis to it.
func DeployParenthesis(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Parenthesis, error) {
	parsed, err := abi.JSON(strings.NewReader(ParenthesisABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParenthesisBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Parenthesis{ParenthesisCaller: ParenthesisCaller{contract: contract}, ParenthesisTransactor: ParenthesisTransactor{contract: contract}, ParenthesisFilterer: ParenthesisFilterer{contract: contract}}, nil
}

// Parenthesis is an auto generated Go binding around an Ethereum contract.
type Parenthesis struct {
	ParenthesisCaller     // Read-only binding to the contract
	ParenthesisTransactor // Write-only binding to the contract
	ParenthesisFilterer   // Log filterer for contract events
}

// ParenthesisCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParenthesisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParenthesisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParenthesisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParenthesisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParenthesisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParenthesisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParenthesisSession struct {
	Contract     *Parenthesis      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParenthesisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParenthesisCallerSession struct {
	Contract *ParenthesisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ParenthesisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParenthesisTransactorSession struct {
	Contract     *ParenthesisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ParenthesisRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParenthesisRaw struct {
	Contract *Parenthesis // Generic contract binding to access the raw methods on
}

// ParenthesisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParenthesisCallerRaw struct {
	Contract *ParenthesisCaller // Generic read-only contract binding to access the raw methods on
}

// ParenthesisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParenthesisTransactorRaw struct {
	Contract *ParenthesisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParenthesis creates a new instance of Parenthesis, bound to a specific deployed contract.
func NewParenthesis(address common.Address, backend bind.ContractBackend) (*Parenthesis, error) {
	contract, err := bindParenthesis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parenthesis{ParenthesisCaller: ParenthesisCaller{contract: contract}, ParenthesisTransactor: ParenthesisTransactor{contract: contract}, ParenthesisFilterer: ParenthesisFilterer{contract: contract}}, nil
}

// NewParenthesisCaller creates a new read-only instance of Parenthesis, bound to a specific deployed contract.
func NewParenthesisCaller(address common.Address, caller bind.ContractCaller) (*ParenthesisCaller, error) {
	contract, err := bindParenthesis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParenthesisCaller{contract: contract}, nil
}

// NewParenthesisTransactor creates a new write-only instance of Parenthesis, bound to a specific deployed contract.
func NewParenthesisTransactor(address common.Address, transactor bind.ContractTransactor) (*ParenthesisTransactor, error) {
	contract, err := bindParenthesis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParenthesisTransactor{contract: contract}, nil
}

// NewParenthesisFilterer creates a new log filterer instance of Parenthesis, bound to a specific deployed contract.
func NewParenthesisFilterer(address common.Address, filterer bind.ContractFilterer) (*ParenthesisFilterer, error) {
	contract, err := bindParenthesis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParenthesisFilterer{contract: contract}, nil
}

// bindParenthesis binds a generic wrapper to an already deployed contract.
func bindParenthesis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParenthesisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parenthesis *ParenthesisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Parenthesis.Contract.ParenthesisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parenthesis *ParenthesisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parenthesis.Contract.ParenthesisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parenthesis *ParenthesisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parenthesis.Contract.ParenthesisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parenthesis *ParenthesisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Parenthesis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parenthesis *ParenthesisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parenthesis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parenthesis *ParenthesisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parenthesis.Contract.contract.Transact(opts, method, params...)
}

// CountWrong is a free data retrieval call binding the contract method 0x247de6d4.
//
// Solidity: function countWrong(string text) pure returns(int256)
func (_Parenthesis *ParenthesisCaller) CountWrong(opts *bind.CallOpts, text string) (*big.Int, error) {
	var out []interface{}
	err := _Parenthesis.contract.Call(opts, &out, "countWrong", text)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountWrong is a free data retrieval call binding the contract method 0x247de6d4.
//
// Solidity: function countWrong(string text) pure returns(int256)
func (_Parenthesis *ParenthesisSession) CountWrong(text string) (*big.Int, error) {
	return _Parenthesis.Contract.CountWrong(&_Parenthesis.CallOpts, text)
}

// CountWrong is a free data retrieval call binding the contract method 0x247de6d4.
//
// Solidity: function countWrong(string text) pure returns(int256)
func (_Parenthesis *ParenthesisCallerSession) CountWrong(text string) (*big.Int, error) {
	return _Parenthesis.Contract.CountWrong(&_Parenthesis.CallOpts, text)
}
