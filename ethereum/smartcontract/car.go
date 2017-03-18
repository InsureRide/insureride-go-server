// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package smartcontract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractCarABI is the input ABI used to generate the binding from.
const ContractCarABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"model\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"driveaddress\",\"type\":\"address\"}],\"name\":\"addDrive\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vehiclenumber\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brand\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drives\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owneraddress\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"year\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodrives\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_model\",\"type\":\"string\"},{\"name\":\"_brand\",\"type\":\"string\"},{\"name\":\"_year\",\"type\":\"uint8\"},{\"name\":\"_owneraddress\",\"type\":\"string\"},{\"name\":\"_vehiclenumber\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ContractCarBin is the compiled bytecode used for deploying new contracts.
const ContractCarBin = `6060604052341561000c57fe5b604051610ab0380380610ab0833981016040528080518201919060200180518201919060200180519060200190919080518201919060200180518201919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b84600390805190602001906100a6929190610153565b5083600290805190602001906100bd929190610153565b5082600460006101000a81548160ff021916908360ff16021790555081600590805190602001906100ef929190610153565b508060069080519060200190610106929190610153565b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50505050506101f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061019457805160ff19168380011785556101c2565b828001600101855582156101c2579182015b828111156101c15782518255916020019190600101906101a6565b5b5090506101cf91906101d3565b5090565b6101f591905b808211156101f15760008160009055506001016101d9565b5090565b90565b6108a9806102076000396000f30060606040523615610097576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630ad9d052146100995780630d890f34146101325780632cc6e051146101685780632d9fff011461020157806341c0e1b51461029a5780634f481a70146102ac578063aaa484251461030c578063f3269716146103a5578063f6ad8bc8146103d1575bfe5b34156100a157fe5b6100a96103f7565b60405180806020018281038252838181518152602001915080519060200190808383600083146100f8575b8051825260208311156100f8576020820191506020810190506020830392506100d4565b505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013a57fe5b610166600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610495565b005b341561017057fe5b610178610565565b60405180806020018281038252838181518152602001915080519060200190808383600083146101c7575b8051825260208311156101c7576020820191506020810190506020830392506101a3565b505050905090810190601f1680156101f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561020957fe5b610211610603565b6040518080602001828103825283818151815260200191508051906020019080838360008314610260575b8051825260208311156102605760208201915060208101905060208303925061023c565b505050905090810190601f16801561028c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102a257fe5b6102aa6106a1565b005b34156102b457fe5b6102ca6004808035906020019091905050610735565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561031457fe5b61031c610775565b604051808060200182810382528381815181526020019150805190602001908083836000831461036b575b80518252602083111561036b57602082019150602081019050602083039250610347565b505050905090810190601f1680156103975780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103ad57fe5b6103b5610813565b604051808260ff1660ff16815260200191505060405180910390f35b34156103d957fe5b6103e1610826565b6040518082815260200191505060405180910390f35b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561048d5780601f106104625761010080835404028352916020019161048d565b820191906000526020600020905b81548152906001019060200180831161047057829003601f168201915b505050505081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156104f157610562565b60088054806001018281610505919061082c565b916000526020600020900160005b83909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506008805490506007819055505b50565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105fb5780601f106105d0576101008083540402835291602001916105fb565b820191906000526020600020905b8154815290600101906020018083116105de57829003601f168201915b505050505081565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106995780601f1061066e57610100808354040283529160200191610699565b820191906000526020600020905b81548152906001019060200180831161067c57829003601f168201915b505050505081565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561073257600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60088181548110151561074457fe5b906000526020600020900160005b915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60058054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561080b5780601f106107e05761010080835404028352916020019161080b565b820191906000526020600020905b8154815290600101906020018083116107ee57829003601f168201915b505050505081565b600460009054906101000a900460ff1681565b60075481565b815481835581811511610853578183600052602060002091820191016108529190610858565b5b505050565b61087a91905b8082111561087657600081600090555060010161085e565b5090565b905600a165627a7a72305820e22f2555dc08ba1bf8c3ae5b1b40e4bfb6e35762363f5285eece466ffd9db21c0029`

// DeployContractCar deploys a new Ethereum contract, binding an instance of ContractCar to it.
func DeployContractCar(auth *bind.TransactOpts, backend bind.ContractBackend, _model string, _brand string, _year uint8, _owneraddress string, _vehiclenumber string) (common.Address, *types.Transaction, *ContractCar, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractCarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractCarBin), backend, _model, _brand, _year, _owneraddress, _vehiclenumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractCar{ContractCarCaller: ContractCarCaller{contract: contract}, ContractCarTransactor: ContractCarTransactor{contract: contract}}, nil
}

// ContractCar is an auto generated Go binding around an Ethereum contract.
type ContractCar struct {
	ContractCarCaller     // Read-only binding to the contract
	ContractCarTransactor // Write-only binding to the contract
}

// ContractCarCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractCarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractCarSession struct {
	Contract     *ContractCar      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCarCallerSession struct {
	Contract *ContractCarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContractCarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractCarTransactorSession struct {
	Contract     *ContractCarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractCarRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractCarRaw struct {
	Contract *ContractCar // Generic contract binding to access the raw methods on
}

// ContractCarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCarCallerRaw struct {
	Contract *ContractCarCaller // Generic read-only contract binding to access the raw methods on
}

// ContractCarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractCarTransactorRaw struct {
	Contract *ContractCarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractCar creates a new instance of ContractCar, bound to a specific deployed contract.
func NewContractCar(address common.Address, backend bind.ContractBackend) (*ContractCar, error) {
	contract, err := bindContractCar(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractCar{ContractCarCaller: ContractCarCaller{contract: contract}, ContractCarTransactor: ContractCarTransactor{contract: contract}}, nil
}

// NewContractCarCaller creates a new read-only instance of ContractCar, bound to a specific deployed contract.
func NewContractCarCaller(address common.Address, caller bind.ContractCaller) (*ContractCarCaller, error) {
	contract, err := bindContractCar(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCarCaller{contract: contract}, nil
}

// NewContractCarTransactor creates a new write-only instance of ContractCar, bound to a specific deployed contract.
func NewContractCarTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractCarTransactor, error) {
	contract, err := bindContractCar(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ContractCarTransactor{contract: contract}, nil
}

// bindContractCar binds a generic wrapper to an already deployed contract.
func bindContractCar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractCarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCar *ContractCarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractCar.Contract.ContractCarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCar *ContractCarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCar.Contract.ContractCarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCar *ContractCarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCar.Contract.ContractCarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCar *ContractCarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractCar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCar *ContractCarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCar *ContractCarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCar.Contract.contract.Transact(opts, method, params...)
}

// Brand is a free data retrieval call binding the contract method 0x2d9fff01.
//
// Solidity: function brand() constant returns(string)
func (_ContractCar *ContractCarCaller) Brand(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "brand")
	return *ret0, err
}

// Brand is a free data retrieval call binding the contract method 0x2d9fff01.
//
// Solidity: function brand() constant returns(string)
func (_ContractCar *ContractCarSession) Brand() (string, error) {
	return _ContractCar.Contract.Brand(&_ContractCar.CallOpts)
}

// Brand is a free data retrieval call binding the contract method 0x2d9fff01.
//
// Solidity: function brand() constant returns(string)
func (_ContractCar *ContractCarCallerSession) Brand() (string, error) {
	return _ContractCar.Contract.Brand(&_ContractCar.CallOpts)
}

// Drives is a free data retrieval call binding the contract method 0x4f481a70.
//
// Solidity: function drives( uint256) constant returns(address)
func (_ContractCar *ContractCarCaller) Drives(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "drives", arg0)
	return *ret0, err
}

// Drives is a free data retrieval call binding the contract method 0x4f481a70.
//
// Solidity: function drives( uint256) constant returns(address)
func (_ContractCar *ContractCarSession) Drives(arg0 *big.Int) (common.Address, error) {
	return _ContractCar.Contract.Drives(&_ContractCar.CallOpts, arg0)
}

// Drives is a free data retrieval call binding the contract method 0x4f481a70.
//
// Solidity: function drives( uint256) constant returns(address)
func (_ContractCar *ContractCarCallerSession) Drives(arg0 *big.Int) (common.Address, error) {
	return _ContractCar.Contract.Drives(&_ContractCar.CallOpts, arg0)
}

// Model is a free data retrieval call binding the contract method 0x0ad9d052.
//
// Solidity: function model() constant returns(string)
func (_ContractCar *ContractCarCaller) Model(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "model")
	return *ret0, err
}

// Model is a free data retrieval call binding the contract method 0x0ad9d052.
//
// Solidity: function model() constant returns(string)
func (_ContractCar *ContractCarSession) Model() (string, error) {
	return _ContractCar.Contract.Model(&_ContractCar.CallOpts)
}

// Model is a free data retrieval call binding the contract method 0x0ad9d052.
//
// Solidity: function model() constant returns(string)
func (_ContractCar *ContractCarCallerSession) Model() (string, error) {
	return _ContractCar.Contract.Model(&_ContractCar.CallOpts)
}

// Nodrives is a free data retrieval call binding the contract method 0xf6ad8bc8.
//
// Solidity: function nodrives() constant returns(uint256)
func (_ContractCar *ContractCarCaller) Nodrives(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "nodrives")
	return *ret0, err
}

// Nodrives is a free data retrieval call binding the contract method 0xf6ad8bc8.
//
// Solidity: function nodrives() constant returns(uint256)
func (_ContractCar *ContractCarSession) Nodrives() (*big.Int, error) {
	return _ContractCar.Contract.Nodrives(&_ContractCar.CallOpts)
}

// Nodrives is a free data retrieval call binding the contract method 0xf6ad8bc8.
//
// Solidity: function nodrives() constant returns(uint256)
func (_ContractCar *ContractCarCallerSession) Nodrives() (*big.Int, error) {
	return _ContractCar.Contract.Nodrives(&_ContractCar.CallOpts)
}

// Owneraddress is a free data retrieval call binding the contract method 0xaaa48425.
//
// Solidity: function owneraddress() constant returns(string)
func (_ContractCar *ContractCarCaller) Owneraddress(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "owneraddress")
	return *ret0, err
}

// Owneraddress is a free data retrieval call binding the contract method 0xaaa48425.
//
// Solidity: function owneraddress() constant returns(string)
func (_ContractCar *ContractCarSession) Owneraddress() (string, error) {
	return _ContractCar.Contract.Owneraddress(&_ContractCar.CallOpts)
}

// Owneraddress is a free data retrieval call binding the contract method 0xaaa48425.
//
// Solidity: function owneraddress() constant returns(string)
func (_ContractCar *ContractCarCallerSession) Owneraddress() (string, error) {
	return _ContractCar.Contract.Owneraddress(&_ContractCar.CallOpts)
}

// Vehiclenumber is a free data retrieval call binding the contract method 0x2cc6e051.
//
// Solidity: function vehiclenumber() constant returns(string)
func (_ContractCar *ContractCarCaller) Vehiclenumber(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "vehiclenumber")
	return *ret0, err
}

// Vehiclenumber is a free data retrieval call binding the contract method 0x2cc6e051.
//
// Solidity: function vehiclenumber() constant returns(string)
func (_ContractCar *ContractCarSession) Vehiclenumber() (string, error) {
	return _ContractCar.Contract.Vehiclenumber(&_ContractCar.CallOpts)
}

// Vehiclenumber is a free data retrieval call binding the contract method 0x2cc6e051.
//
// Solidity: function vehiclenumber() constant returns(string)
func (_ContractCar *ContractCarCallerSession) Vehiclenumber() (string, error) {
	return _ContractCar.Contract.Vehiclenumber(&_ContractCar.CallOpts)
}

// Year is a free data retrieval call binding the contract method 0xf3269716.
//
// Solidity: function year() constant returns(uint8)
func (_ContractCar *ContractCarCaller) Year(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "year")
	return *ret0, err
}

// Year is a free data retrieval call binding the contract method 0xf3269716.
//
// Solidity: function year() constant returns(uint8)
func (_ContractCar *ContractCarSession) Year() (uint8, error) {
	return _ContractCar.Contract.Year(&_ContractCar.CallOpts)
}

// Year is a free data retrieval call binding the contract method 0xf3269716.
//
// Solidity: function year() constant returns(uint8)
func (_ContractCar *ContractCarCallerSession) Year() (uint8, error) {
	return _ContractCar.Contract.Year(&_ContractCar.CallOpts)
}

// AddDrive is a paid mutator transaction binding the contract method 0x0d890f34.
//
// Solidity: function addDrive(driveaddress address) returns()
func (_ContractCar *ContractCarTransactor) AddDrive(opts *bind.TransactOpts, driveaddress common.Address) (*types.Transaction, error) {
	return _ContractCar.contract.Transact(opts, "addDrive", driveaddress)
}

// AddDrive is a paid mutator transaction binding the contract method 0x0d890f34.
//
// Solidity: function addDrive(driveaddress address) returns()
func (_ContractCar *ContractCarSession) AddDrive(driveaddress common.Address) (*types.Transaction, error) {
	return _ContractCar.Contract.AddDrive(&_ContractCar.TransactOpts, driveaddress)
}

// AddDrive is a paid mutator transaction binding the contract method 0x0d890f34.
//
// Solidity: function addDrive(driveaddress address) returns()
func (_ContractCar *ContractCarTransactorSession) AddDrive(driveaddress common.Address) (*types.Transaction, error) {
	return _ContractCar.Contract.AddDrive(&_ContractCar.TransactOpts, driveaddress)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ContractCar *ContractCarTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCar.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ContractCar *ContractCarSession) Kill() (*types.Transaction, error) {
	return _ContractCar.Contract.Kill(&_ContractCar.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ContractCar *ContractCarTransactorSession) Kill() (*types.Transaction, error) {
	return _ContractCar.Contract.Kill(&_ContractCar.TransactOpts)
}
