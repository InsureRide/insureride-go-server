// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package smartcontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractDriveABI is the input ABI used to generate the binding from.
const ContractDriveABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"kilometers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endtime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"avgspeed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"starttime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kilometers\",\"type\":\"uint64\"},{\"name\":\"_avgspeed\",\"type\":\"uint16\"},{\"name\":\"_starttime\",\"type\":\"uint32\"},{\"name\":\"_endtime\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ContractDriveBin is the compiled bytecode used for deploying new contracts.
const ContractDriveBin = `6060604052341561000c57fe5b604051608080610378833981016040528080519060200190919080519060200190919080519060200190919080519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b83600060146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000601c6101000a81548161ffff021916908361ffff16021790555081600160006101000a81548163ffffffff021916908363ffffffff16021790555080600160049054906101000a900463ffffffff1650505b505050505b61025f806101196000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806321d6dc90146100675780632ca126f8146100a157806341c0e1b5146100d357806348b1ad20146100e55780638da5889714610113575bfe5b341561006f57fe5b610077610145565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b34156100a957fe5b6100b161015f565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34156100db57fe5b6100e3610175565b005b34156100ed57fe5b6100f5610209565b604051808261ffff1661ffff16815260200191505060405180910390f35b341561011b57fe5b61012361021d565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b600060149054906101000a900467ffffffffffffffff1681565b600160049054906101000a900463ffffffff1681565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561020657600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b6000601c9054906101000a900461ffff1681565b600160009054906101000a900463ffffffff16815600a165627a7a723058209ca3a04d021732490733a253b31446c3047b13582f411f5c297198891dd767fa0029`

// DeployContractDrive deploys a new Ethereum contract, binding an instance of ContractDrive to it.
func DeployContractDrive(auth *bind.TransactOpts, backend bind.ContractBackend, _kilometers uint64, _avgspeed uint16, _starttime uint32, _endtime uint32) (common.Address, *types.Transaction, *ContractDrive, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractDriveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractDriveBin), backend, _kilometers, _avgspeed, _starttime, _endtime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractDrive{ContractDriveCaller: ContractDriveCaller{contract: contract}, ContractDriveTransactor: ContractDriveTransactor{contract: contract}}, nil
}

// ContractDrive is an auto generated Go binding around an Ethereum contract.
type ContractDrive struct {
	ContractDriveCaller     // Read-only binding to the contract
	ContractDriveTransactor // Write-only binding to the contract
}

// ContractDriveCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractDriveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDriveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractDriveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDriveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractDriveSession struct {
	Contract     *ContractDrive    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractDriveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractDriveCallerSession struct {
	Contract *ContractDriveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ContractDriveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractDriveTransactorSession struct {
	Contract     *ContractDriveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractDriveRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractDriveRaw struct {
	Contract *ContractDrive // Generic contract binding to access the raw methods on
}

// ContractDriveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractDriveCallerRaw struct {
	Contract *ContractDriveCaller // Generic read-only contract binding to access the raw methods on
}

// ContractDriveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractDriveTransactorRaw struct {
	Contract *ContractDriveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractDrive creates a new instance of ContractDrive, bound to a specific deployed contract.
func NewContractDrive(address common.Address, backend bind.ContractBackend) (*ContractDrive, error) {
	contract, err := bindContractDrive(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractDrive{ContractDriveCaller: ContractDriveCaller{contract: contract}, ContractDriveTransactor: ContractDriveTransactor{contract: contract}}, nil
}

// NewContractDriveCaller creates a new read-only instance of ContractDrive, bound to a specific deployed contract.
func NewContractDriveCaller(address common.Address, caller bind.ContractCaller) (*ContractDriveCaller, error) {
	contract, err := bindContractDrive(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDriveCaller{contract: contract}, nil
}

// NewContractDriveTransactor creates a new write-only instance of ContractDrive, bound to a specific deployed contract.
func NewContractDriveTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractDriveTransactor, error) {
	contract, err := bindContractDrive(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ContractDriveTransactor{contract: contract}, nil
}

// bindContractDrive binds a generic wrapper to an already deployed contract.
func bindContractDrive(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractDriveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDrive *ContractDriveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractDrive.Contract.ContractDriveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDrive *ContractDriveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDrive.Contract.ContractDriveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDrive *ContractDriveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDrive.Contract.ContractDriveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDrive *ContractDriveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractDrive.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDrive *ContractDriveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDrive.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDrive *ContractDriveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDrive.Contract.contract.Transact(opts, method, params...)
}

// Avgspeed is a free data retrieval call binding the contract method 0x48b1ad20.
//
// Solidity: function avgspeed() constant returns(uint16)
func (_ContractDrive *ContractDriveCaller) Avgspeed(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "avgspeed")
	return *ret0, err
}

// Avgspeed is a free data retrieval call binding the contract method 0x48b1ad20.
//
// Solidity: function avgspeed() constant returns(uint16)
func (_ContractDrive *ContractDriveSession) Avgspeed() (uint16, error) {
	return _ContractDrive.Contract.Avgspeed(&_ContractDrive.CallOpts)
}

// Avgspeed is a free data retrieval call binding the contract method 0x48b1ad20.
//
// Solidity: function avgspeed() constant returns(uint16)
func (_ContractDrive *ContractDriveCallerSession) Avgspeed() (uint16, error) {
	return _ContractDrive.Contract.Avgspeed(&_ContractDrive.CallOpts)
}

// Endtime is a free data retrieval call binding the contract method 0x2ca126f8.
//
// Solidity: function endtime() constant returns(uint32)
func (_ContractDrive *ContractDriveCaller) Endtime(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "endtime")
	return *ret0, err
}

// Endtime is a free data retrieval call binding the contract method 0x2ca126f8.
//
// Solidity: function endtime() constant returns(uint32)
func (_ContractDrive *ContractDriveSession) Endtime() (uint32, error) {
	return _ContractDrive.Contract.Endtime(&_ContractDrive.CallOpts)
}

// Endtime is a free data retrieval call binding the contract method 0x2ca126f8.
//
// Solidity: function endtime() constant returns(uint32)
func (_ContractDrive *ContractDriveCallerSession) Endtime() (uint32, error) {
	return _ContractDrive.Contract.Endtime(&_ContractDrive.CallOpts)
}

// Kilometers is a free data retrieval call binding the contract method 0x21d6dc90.
//
// Solidity: function kilometers() constant returns(uint64)
func (_ContractDrive *ContractDriveCaller) Kilometers(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "kilometers")
	return *ret0, err
}

// Kilometers is a free data retrieval call binding the contract method 0x21d6dc90.
//
// Solidity: function kilometers() constant returns(uint64)
func (_ContractDrive *ContractDriveSession) Kilometers() (uint64, error) {
	return _ContractDrive.Contract.Kilometers(&_ContractDrive.CallOpts)
}

// Kilometers is a free data retrieval call binding the contract method 0x21d6dc90.
//
// Solidity: function kilometers() constant returns(uint64)
func (_ContractDrive *ContractDriveCallerSession) Kilometers() (uint64, error) {
	return _ContractDrive.Contract.Kilometers(&_ContractDrive.CallOpts)
}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() constant returns(uint32)
func (_ContractDrive *ContractDriveCaller) Starttime(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "starttime")
	return *ret0, err
}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() constant returns(uint32)
func (_ContractDrive *ContractDriveSession) Starttime() (uint32, error) {
	return _ContractDrive.Contract.Starttime(&_ContractDrive.CallOpts)
}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() constant returns(uint32)
func (_ContractDrive *ContractDriveCallerSession) Starttime() (uint32, error) {
	return _ContractDrive.Contract.Starttime(&_ContractDrive.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ContractDrive *ContractDriveTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDrive.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ContractDrive *ContractDriveSession) Kill() (*types.Transaction, error) {
	return _ContractDrive.Contract.Kill(&_ContractDrive.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ContractDrive *ContractDriveTransactorSession) Kill() (*types.Transaction, error) {
	return _ContractDrive.Contract.Kill(&_ContractDrive.TransactOpts)
}
