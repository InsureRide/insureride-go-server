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
const ContractDriveABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"kilometers\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endtime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"avgspeed\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"starttime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"avgaccel\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kilometers\",\"type\":\"string\"},{\"name\":\"_avgspeed\",\"type\":\"string\"},{\"name\":\"_avgaccel\",\"type\":\"string\"},{\"name\":\"_starttime\",\"type\":\"uint32\"},{\"name\":\"_endtime\",\"type\":\"uint32\"},{\"name\":\"_price\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ContractDriveBin is the compiled bytecode used for deploying new contracts.
const ContractDriveBin = `6060604052341561000c57fe5b6040516107cc3803806107cc833981016040528080518201919060200180518201919060200180518201919060200180519060200190919080519060200190919080519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b85600190805190602001906100b292919061014c565b5084600290805190602001906100c992919061014c565b5083600390805190602001906100e092919061014c565b5082600460006101000a81548163ffffffff021916908363ffffffff16021790555081600460046101000a81548163ffffffff021916908363ffffffff16021790555080600460086101000a81548161ffff021916908361ffff1602179055505b5050505050506101f1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061018d57805160ff19168380011785556101bb565b828001600101855582156101bb579182015b828111156101ba57825182559160200191906001019061019f565b5b5090506101c891906101cc565b5090565b6101ee91905b808211156101ea5760008160009055506001016101d2565b5090565b90565b6105cc806102006000396000f30060606040523615610081576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806321d6dc90146100835780632ca126f81461011c57806341c0e1b51461014e57806348b1ad20146101605780638da58897146101f9578063a035b1fe1461022b578063ed0ad9c514610259575bfe5b341561008b57fe5b6100936102f2565b60405180806020018281038252838181518152602001915080519060200190808383600083146100e2575b8051825260208311156100e2576020820191506020810190506020830392506100be565b505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561012457fe5b61012c610390565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b341561015657fe5b61015e6103a6565b005b341561016857fe5b61017061043a565b60405180806020018281038252838181518152602001915080519060200190808383600083146101bf575b8051825260208311156101bf5760208201915060208101905060208303925061019b565b505050905090810190601f1680156101eb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561020157fe5b6102096104d8565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b341561023357fe5b61023b6104ee565b604051808261ffff1661ffff16815260200191505060405180910390f35b341561026157fe5b610269610502565b60405180806020018281038252838181518152602001915080519060200190808383600083146102b8575b8051825260208311156102b857602082019150602081019050602083039250610294565b505050905090810190601f1680156102e45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103885780601f1061035d57610100808354040283529160200191610388565b820191906000526020600020905b81548152906001019060200180831161036b57829003601f168201915b505050505081565b600460049054906101000a900463ffffffff1681565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561043757600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104d05780601f106104a5576101008083540402835291602001916104d0565b820191906000526020600020905b8154815290600101906020018083116104b357829003601f168201915b505050505081565b600460009054906101000a900463ffffffff1681565b600460089054906101000a900461ffff1681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105985780601f1061056d57610100808354040283529160200191610598565b820191906000526020600020905b81548152906001019060200180831161057b57829003601f168201915b5050505050815600a165627a7a7230582039f879f8005d87200d41dd2ca99714b0816cf78ae8b0b2e9ab241278a2d656b40029`

// DeployContractDrive deploys a new Ethereum contract, binding an instance of ContractDrive to it.
func DeployContractDrive(auth *bind.TransactOpts, backend bind.ContractBackend, _kilometers string, _avgspeed string, _avgaccel string, _starttime uint32, _endtime uint32, _price uint16) (common.Address, *types.Transaction, *ContractDrive, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractDriveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractDriveBin), backend, _kilometers, _avgspeed, _avgaccel, _starttime, _endtime, _price)
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

// Avgaccel is a free data retrieval call binding the contract method 0xed0ad9c5.
//
// Solidity: function avgaccel() constant returns(string)
func (_ContractDrive *ContractDriveCaller) Avgaccel(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "avgaccel")
	return *ret0, err
}

// Avgaccel is a free data retrieval call binding the contract method 0xed0ad9c5.
//
// Solidity: function avgaccel() constant returns(string)
func (_ContractDrive *ContractDriveSession) Avgaccel() (string, error) {
	return _ContractDrive.Contract.Avgaccel(&_ContractDrive.CallOpts)
}

// Avgaccel is a free data retrieval call binding the contract method 0xed0ad9c5.
//
// Solidity: function avgaccel() constant returns(string)
func (_ContractDrive *ContractDriveCallerSession) Avgaccel() (string, error) {
	return _ContractDrive.Contract.Avgaccel(&_ContractDrive.CallOpts)
}

// Avgspeed is a free data retrieval call binding the contract method 0x48b1ad20.
//
// Solidity: function avgspeed() constant returns(string)
func (_ContractDrive *ContractDriveCaller) Avgspeed(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "avgspeed")
	return *ret0, err
}

// Avgspeed is a free data retrieval call binding the contract method 0x48b1ad20.
//
// Solidity: function avgspeed() constant returns(string)
func (_ContractDrive *ContractDriveSession) Avgspeed() (string, error) {
	return _ContractDrive.Contract.Avgspeed(&_ContractDrive.CallOpts)
}

// Avgspeed is a free data retrieval call binding the contract method 0x48b1ad20.
//
// Solidity: function avgspeed() constant returns(string)
func (_ContractDrive *ContractDriveCallerSession) Avgspeed() (string, error) {
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
// Solidity: function kilometers() constant returns(string)
func (_ContractDrive *ContractDriveCaller) Kilometers(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "kilometers")
	return *ret0, err
}

// Kilometers is a free data retrieval call binding the contract method 0x21d6dc90.
//
// Solidity: function kilometers() constant returns(string)
func (_ContractDrive *ContractDriveSession) Kilometers() (string, error) {
	return _ContractDrive.Contract.Kilometers(&_ContractDrive.CallOpts)
}

// Kilometers is a free data retrieval call binding the contract method 0x21d6dc90.
//
// Solidity: function kilometers() constant returns(string)
func (_ContractDrive *ContractDriveCallerSession) Kilometers() (string, error) {
	return _ContractDrive.Contract.Kilometers(&_ContractDrive.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint16)
func (_ContractDrive *ContractDriveCaller) Price(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _ContractDrive.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint16)
func (_ContractDrive *ContractDriveSession) Price() (uint16, error) {
	return _ContractDrive.Contract.Price(&_ContractDrive.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint16)
func (_ContractDrive *ContractDriveCallerSession) Price() (uint16, error) {
	return _ContractDrive.Contract.Price(&_ContractDrive.CallOpts)
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
