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

// ThreeThatAddABI is the input ABI used to generate the binding from.
const ThreeThatAddABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"item\",\"type\":\"uint256\"}],\"name\":\"addItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllItems\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sum\",\"type\":\"uint256\"}],\"name\":\"threeThatAdd\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ThreeThatAddFuncSigs maps the 4-byte function signature to its string representation.
var ThreeThatAddFuncSigs = map[string]string{
	"be634045": "addItem(uint256)",
	"4ba1d6aa": "getAllItems()",
	"bfb231d2": "items(uint256)",
	"acec599e": "resetItems()",
	"c734d4c5": "threeThatAdd(uint256)",
}

// ThreeThatAddBin is the compiled bytecode used for deploying new contracts.
var ThreeThatAddBin = "0x608060405234801561001057600080fd5b506106a7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80634ba1d6aa1461005c578063acec599e1461007a578063be63404514610084578063bfb231d2146100c4578063c734d4c5146100e5575b600080fd5b610064610105565b604051610071919061047b565b60405180910390f35b61008261015d565b005b6100826100923660046104bf565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630155565b6100d76100d23660046104bf565b61016a565b604051908152602001610071565b6100f86100f33660046104bf565b61018b565b60405161007191906104fc565b6060600080548060200260200160405190810160405280929190818152602001828054801561015357602002820191906000526020600020905b81548152602001906001019080831161013f575b5050505050905090565b610168600080610441565b565b6000818154811061017a57600080fd5b600091825260209091200154905081565b606060005b6000548110156103095760006101a7826001610545565b90505b6000548110156102f65760006101c1826001610545565b90505b6000548110156102e35784600082815481106101e2576101e261055e565b9060005260206000200154600084815481106102005761020061055e565b90600052602060002001546000868154811061021e5761021e61055e565b90600052602060002001546102339190610545565b61023d9190610545565b036102d157600061026a6000858154811061025a5761025a61055e565b9060005260206000200154610335565b905060006102846000858154811061025a5761025a61055e565b9050600061029e6000858154811061025a5761025a61055e565b90508282826040516020016102b593929190610574565b6040516020818303038152906040529650505050505050919050565b806102db816105ed565b9150506101c4565b50806102ee816105ed565b9150506101aa565b5080610301816105ed565b915050610190565b505060408051808201909152600e81526d1b9bc81a5d195b5cc8199bdd5b9960921b6020820152919050565b60608160000361035c5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156103865780610370816105ed565b915061037f9050600a8361061c565b9150610360565b60008167ffffffffffffffff8111156103a1576103a1610630565b6040519080825280601f01601f1916602001820160405280156103cb576020820181803683370190505b508593509050815b8315610438576103e4600a85610646565b6103ef906030610545565b60f81b826103fc8361065a565b9250828151811061040f5761040f61055e565b60200101906001600160f81b031916908160001a905350610431600a8561061c565b93506103d3565b50949350505050565b508054600082559060005260206000209081019061045f9190610462565b50565b5b808211156104775760008155600101610463565b5090565b6020808252825182820181905260009190848201906040850190845b818110156104b357835183529284019291840191600101610497565b50909695505050505050565b6000602082840312156104d157600080fd5b5035919050565b60005b838110156104f35781810151838201526020016104db565b50506000910152565b602081526000825180602084015261051b8160408501602087016104d8565b601f01601f19169190910160400192915050565b634e487b7160e01b600052601160045260246000fd5b808201808211156105585761055861052f565b92915050565b634e487b7160e01b600052603260045260246000fd5b7103337bab73210199032b632b6b2b73a399d160751b8152600084516105a18160128501602089016104d8565b808301905061016160f51b80601283015285516105c5816014850160208a016104d8565b601492019182015283516105e08160168401602088016104d8565b0160160195945050505050565b6000600182016105ff576105ff61052f565b5060010190565b634e487b7160e01b600052601260045260246000fd5b60008261062b5761062b610606565b500490565b634e487b7160e01b600052604160045260246000fd5b60008261065557610655610606565b500690565b6000816106695761066961052f565b50600019019056fea26469706673582212204a1969bbba6952355f0dfde2f7a2d45fc62e423e67f3748420a79cc4613b184e64736f6c63430008100033"

// DeployThreeThatAdd deploys a new Ethereum contract, binding an instance of ThreeThatAdd to it.
func DeployThreeThatAdd(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ThreeThatAdd, error) {
	parsed, err := abi.JSON(strings.NewReader(ThreeThatAddABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ThreeThatAddBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ThreeThatAdd{ThreeThatAddCaller: ThreeThatAddCaller{contract: contract}, ThreeThatAddTransactor: ThreeThatAddTransactor{contract: contract}, ThreeThatAddFilterer: ThreeThatAddFilterer{contract: contract}}, nil
}

// ThreeThatAdd is an auto generated Go binding around an Ethereum contract.
type ThreeThatAdd struct {
	ThreeThatAddCaller     // Read-only binding to the contract
	ThreeThatAddTransactor // Write-only binding to the contract
	ThreeThatAddFilterer   // Log filterer for contract events
}

// ThreeThatAddCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThreeThatAddCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeThatAddTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThreeThatAddTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeThatAddFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThreeThatAddFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeThatAddSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThreeThatAddSession struct {
	Contract     *ThreeThatAdd     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThreeThatAddCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThreeThatAddCallerSession struct {
	Contract *ThreeThatAddCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ThreeThatAddTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThreeThatAddTransactorSession struct {
	Contract     *ThreeThatAddTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ThreeThatAddRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThreeThatAddRaw struct {
	Contract *ThreeThatAdd // Generic contract binding to access the raw methods on
}

// ThreeThatAddCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThreeThatAddCallerRaw struct {
	Contract *ThreeThatAddCaller // Generic read-only contract binding to access the raw methods on
}

// ThreeThatAddTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThreeThatAddTransactorRaw struct {
	Contract *ThreeThatAddTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThreeThatAdd creates a new instance of ThreeThatAdd, bound to a specific deployed contract.
func NewThreeThatAdd(address common.Address, backend bind.ContractBackend) (*ThreeThatAdd, error) {
	contract, err := bindThreeThatAdd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThreeThatAdd{ThreeThatAddCaller: ThreeThatAddCaller{contract: contract}, ThreeThatAddTransactor: ThreeThatAddTransactor{contract: contract}, ThreeThatAddFilterer: ThreeThatAddFilterer{contract: contract}}, nil
}

// NewThreeThatAddCaller creates a new read-only instance of ThreeThatAdd, bound to a specific deployed contract.
func NewThreeThatAddCaller(address common.Address, caller bind.ContractCaller) (*ThreeThatAddCaller, error) {
	contract, err := bindThreeThatAdd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeThatAddCaller{contract: contract}, nil
}

// NewThreeThatAddTransactor creates a new write-only instance of ThreeThatAdd, bound to a specific deployed contract.
func NewThreeThatAddTransactor(address common.Address, transactor bind.ContractTransactor) (*ThreeThatAddTransactor, error) {
	contract, err := bindThreeThatAdd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeThatAddTransactor{contract: contract}, nil
}

// NewThreeThatAddFilterer creates a new log filterer instance of ThreeThatAdd, bound to a specific deployed contract.
func NewThreeThatAddFilterer(address common.Address, filterer bind.ContractFilterer) (*ThreeThatAddFilterer, error) {
	contract, err := bindThreeThatAdd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThreeThatAddFilterer{contract: contract}, nil
}

// bindThreeThatAdd binds a generic wrapper to an already deployed contract.
func bindThreeThatAdd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThreeThatAddABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeThatAdd *ThreeThatAddRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreeThatAdd.Contract.ThreeThatAddCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeThatAdd *ThreeThatAddRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.ThreeThatAddTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeThatAdd *ThreeThatAddRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.ThreeThatAddTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeThatAdd *ThreeThatAddCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreeThatAdd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeThatAdd *ThreeThatAddTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeThatAdd *ThreeThatAddTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.contract.Transact(opts, method, params...)
}

// GetAllItems is a free data retrieval call binding the contract method 0x4ba1d6aa.
//
// Solidity: function getAllItems() view returns(uint256[])
func (_ThreeThatAdd *ThreeThatAddCaller) GetAllItems(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _ThreeThatAdd.contract.Call(opts, &out, "getAllItems")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllItems is a free data retrieval call binding the contract method 0x4ba1d6aa.
//
// Solidity: function getAllItems() view returns(uint256[])
func (_ThreeThatAdd *ThreeThatAddSession) GetAllItems() ([]*big.Int, error) {
	return _ThreeThatAdd.Contract.GetAllItems(&_ThreeThatAdd.CallOpts)
}

// GetAllItems is a free data retrieval call binding the contract method 0x4ba1d6aa.
//
// Solidity: function getAllItems() view returns(uint256[])
func (_ThreeThatAdd *ThreeThatAddCallerSession) GetAllItems() ([]*big.Int, error) {
	return _ThreeThatAdd.Contract.GetAllItems(&_ThreeThatAdd.CallOpts)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256)
func (_ThreeThatAdd *ThreeThatAddCaller) Items(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThreeThatAdd.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256)
func (_ThreeThatAdd *ThreeThatAddSession) Items(arg0 *big.Int) (*big.Int, error) {
	return _ThreeThatAdd.Contract.Items(&_ThreeThatAdd.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256)
func (_ThreeThatAdd *ThreeThatAddCallerSession) Items(arg0 *big.Int) (*big.Int, error) {
	return _ThreeThatAdd.Contract.Items(&_ThreeThatAdd.CallOpts, arg0)
}

// ThreeThatAdd is a free data retrieval call binding the contract method 0xc734d4c5.
//
// Solidity: function threeThatAdd(uint256 sum) view returns(string)
func (_ThreeThatAdd *ThreeThatAddCaller) ThreeThatAdd(opts *bind.CallOpts, sum *big.Int) (string, error) {
	var out []interface{}
	err := _ThreeThatAdd.contract.Call(opts, &out, "threeThatAdd", sum)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ThreeThatAdd is a free data retrieval call binding the contract method 0xc734d4c5.
//
// Solidity: function threeThatAdd(uint256 sum) view returns(string)
func (_ThreeThatAdd *ThreeThatAddSession) ThreeThatAdd(sum *big.Int) (string, error) {
	return _ThreeThatAdd.Contract.ThreeThatAdd(&_ThreeThatAdd.CallOpts, sum)
}

// ThreeThatAdd is a free data retrieval call binding the contract method 0xc734d4c5.
//
// Solidity: function threeThatAdd(uint256 sum) view returns(string)
func (_ThreeThatAdd *ThreeThatAddCallerSession) ThreeThatAdd(sum *big.Int) (string, error) {
	return _ThreeThatAdd.Contract.ThreeThatAdd(&_ThreeThatAdd.CallOpts, sum)
}

// AddItem is a paid mutator transaction binding the contract method 0xbe634045.
//
// Solidity: function addItem(uint256 item) returns()
func (_ThreeThatAdd *ThreeThatAddTransactor) AddItem(opts *bind.TransactOpts, item *big.Int) (*types.Transaction, error) {
	return _ThreeThatAdd.contract.Transact(opts, "addItem", item)
}

// AddItem is a paid mutator transaction binding the contract method 0xbe634045.
//
// Solidity: function addItem(uint256 item) returns()
func (_ThreeThatAdd *ThreeThatAddSession) AddItem(item *big.Int) (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.AddItem(&_ThreeThatAdd.TransactOpts, item)
}

// AddItem is a paid mutator transaction binding the contract method 0xbe634045.
//
// Solidity: function addItem(uint256 item) returns()
func (_ThreeThatAdd *ThreeThatAddTransactorSession) AddItem(item *big.Int) (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.AddItem(&_ThreeThatAdd.TransactOpts, item)
}

// ResetItems is a paid mutator transaction binding the contract method 0xacec599e.
//
// Solidity: function resetItems() returns()
func (_ThreeThatAdd *ThreeThatAddTransactor) ResetItems(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeThatAdd.contract.Transact(opts, "resetItems")
}

// ResetItems is a paid mutator transaction binding the contract method 0xacec599e.
//
// Solidity: function resetItems() returns()
func (_ThreeThatAdd *ThreeThatAddSession) ResetItems() (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.ResetItems(&_ThreeThatAdd.TransactOpts)
}

// ResetItems is a paid mutator transaction binding the contract method 0xacec599e.
//
// Solidity: function resetItems() returns()
func (_ThreeThatAdd *ThreeThatAddTransactorSession) ResetItems() (*types.Transaction, error) {
	return _ThreeThatAdd.Contract.ResetItems(&_ThreeThatAdd.TransactOpts)
}
