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
const ContractCarABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"model\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"driveaddress\",\"type\":\"address\"}],\"name\":\"addDrive\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vehiclenumber\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brand\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drives\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payInsurance\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owneraddress\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"year\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodrives\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_model\",\"type\":\"string\"},{\"name\":\"_brand\",\"type\":\"string\"},{\"name\":\"_year\",\"type\":\"uint8\"},{\"name\":\"_owneraddress\",\"type\":\"string\"},{\"name\":\"_vehiclenumber\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ContractCarBin is the compiled bytecode used for deploying new contracts.
const ContractCarBin = `6060604052341561000c57fe5b604051610b30380380610b30833981016040528080518201919060200180518201919060200180519060200190919080518201919060200180518201919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b84600390805190602001906100a692919061015d565b5083600290805190602001906100bd92919061015d565b5082600460006101000a81548160ff021916908360ff16021790555081600590805190602001906100ef92919061015d565b50806006908051906020019061010692919061015d565b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550629896806009819055505b5050505050610202565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061019e57805160ff19168380011785556101cc565b828001600101855582156101cc579182015b828111156101cb5782518255916020019190600101906101b0565b5b5090506101d991906101dd565b5090565b6101ff91905b808211156101fb5760008160009055506001016101e3565b5090565b90565b61091f806102116000396000f300606060405236156100ad576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630ad9d052146100af5780630d890f34146101485780632cc6e0511461017e5780632d9fff011461021757806341c0e1b5146102b05780634f481a70146102c2578063508e264e14610322578063aaa4842514610342578063b69ef8a8146103db578063f326971614610401578063f6ad8bc81461042d575bfe5b34156100b757fe5b6100bf610453565b604051808060200182810382528381815181526020019150805190602001908083836000831461010e575b80518252602083111561010e576020820191506020810190506020830392506100ea565b505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015057fe5b61017c600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104f1565b005b341561018657fe5b61018e6105c1565b60405180806020018281038252838181518152602001915080519060200190808383600083146101dd575b8051825260208311156101dd576020820191506020810190506020830392506101b9565b505050905090810190601f1680156102095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561021f57fe5b61022761065f565b6040518080602001828103825283818151815260200191508051906020019080838360008314610276575b80518252602083111561027657602082019150602081019050602083039250610252565b505050905090810190601f1680156102a25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102b857fe5b6102c06106fd565b005b34156102ca57fe5b6102e06004808035906020019091905050610791565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561032a57fe5b61034060048080359060200190919050506107d1565b005b341561034a57fe5b6103526107e5565b60405180806020018281038252838181518152602001915080519060200190808383600083146103a1575b8051825260208311156103a15760208201915060208101905060208303925061037d565b505050905090810190601f1680156103cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103e357fe5b6103eb610883565b6040518082815260200191505060405180910390f35b341561040957fe5b610411610889565b604051808260ff1660ff16815260200191505060405180910390f35b341561043557fe5b61043d61089c565b6040518082815260200191505060405180910390f35b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104e95780601f106104be576101008083540402835291602001916104e9565b820191906000526020600020905b8154815290600101906020018083116104cc57829003601f168201915b505050505081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561054d576105be565b6008805480600101828161056191906108a2565b916000526020600020900160005b83909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506008805490506007819055505b50565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106575780601f1061062c57610100808354040283529160200191610657565b820191906000526020600020905b81548152906001019060200180831161063a57829003601f168201915b505050505081565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106f55780601f106106ca576101008083540402835291602001916106f5565b820191906000526020600020905b8154815290600101906020018083116106d857829003601f168201915b505050505081565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561078e57600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b6008818154811015156107a057fe5b906000526020600020900160005b915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b806009600082825403925050819055505b50565b60058054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561087b5780601f106108505761010080835404028352916020019161087b565b820191906000526020600020905b81548152906001019060200180831161085e57829003601f168201915b505050505081565b60095481565b600460009054906101000a900460ff1681565b60075481565b8154818355818115116108c9578183600052602060002091820191016108c891906108ce565b5b505050565b6108f091905b808211156108ec5760008160009055506001016108d4565b5090565b905600a165627a7a7230582009d0b60c5aeb221f1b530016775b21342bf4b1bab38e1dd66c3a3a5c57780a740029`

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

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_ContractCar *ContractCarCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractCar.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_ContractCar *ContractCarSession) Balance() (*big.Int, error) {
	return _ContractCar.Contract.Balance(&_ContractCar.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_ContractCar *ContractCarCallerSession) Balance() (*big.Int, error) {
	return _ContractCar.Contract.Balance(&_ContractCar.CallOpts)
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

// PayInsurance is a paid mutator transaction binding the contract method 0x508e264e.
//
// Solidity: function payInsurance(amount uint256) returns()
func (_ContractCar *ContractCarTransactor) PayInsurance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ContractCar.contract.Transact(opts, "payInsurance", amount)
}

// PayInsurance is a paid mutator transaction binding the contract method 0x508e264e.
//
// Solidity: function payInsurance(amount uint256) returns()
func (_ContractCar *ContractCarSession) PayInsurance(amount *big.Int) (*types.Transaction, error) {
	return _ContractCar.Contract.PayInsurance(&_ContractCar.TransactOpts, amount)
}

// PayInsurance is a paid mutator transaction binding the contract method 0x508e264e.
//
// Solidity: function payInsurance(amount uint256) returns()
func (_ContractCar *ContractCarTransactorSession) PayInsurance(amount *big.Int) (*types.Transaction, error) {
	return _ContractCar.Contract.PayInsurance(&_ContractCar.TransactOpts, amount)
}
