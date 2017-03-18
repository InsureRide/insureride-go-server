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
const ContractDriveABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"kilometers\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endtime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"avgspeed\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"starttime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"avgaccel\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kilometers\",\"type\":\"string\"},{\"name\":\"_avgspeed\",\"type\":\"string\"},{\"name\":\"_avgaccel\",\"type\":\"string\"},{\"name\":\"_starttime\",\"type\":\"uint32\"},{\"name\":\"_endtime\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ContractDriveBin is the compiled bytecode used for deploying new contracts.
const ContractDriveBin = `6060604052341561000c57fe5b60405161074d38038061074d833981016040528080518201919060200180518201919060200180518201919060200180519060200190919080519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b84600190805190602001906100a992919061011a565b5083600290805190602001906100c092919061011a565b5082600390805190602001906100d792919061011a565b5081600460006101000a81548163ffffffff021916908363ffffffff16021790555080600460049054906101000a900463ffffffff1650505b50505050506101bf565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061015b57805160ff1916838001178555610189565b82800160010185558215610189579182015b8281111561018857825182559160200191906001019061016d565b5b509050610196919061019a565b5090565b6101bc91905b808211156101b85760008160009055506001016101a0565b5090565b90565b61057f806101ce6000396000f30060606040523615610076576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806321d6dc90146100785780632ca126f81461011157806341c0e1b51461014357806348b1ad20146101555780638da58897146101ee578063ed0ad9c514610220575bfe5b341561008057fe5b6100886102b9565b60405180806020018281038252838181518152602001915080519060200190808383600083146100d7575b8051825260208311156100d7576020820191506020810190506020830392506100b3565b505050905090810190601f1680156101035780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561011957fe5b610121610357565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b341561014b57fe5b61015361036d565b005b341561015d57fe5b610165610401565b60405180806020018281038252838181518152602001915080519060200190808383600083146101b4575b8051825260208311156101b457602082019150602081019050602083039250610190565b505050905090810190601f1680156101e05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101f657fe5b6101fe61049f565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b341561022857fe5b6102306104b5565b604051808060200182810382528381815181526020019150805190602001908083836000831461027f575b80518252602083111561027f5760208201915060208101905060208303925061025b565b505050905090810190601f1680156102ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561034f5780601f106103245761010080835404028352916020019161034f565b820191906000526020600020905b81548152906001019060200180831161033257829003601f168201915b505050505081565b600460049054906101000a900463ffffffff1681565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103fe57600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104975780601f1061046c57610100808354040283529160200191610497565b820191906000526020600020905b81548152906001019060200180831161047a57829003601f168201915b505050505081565b600460009054906101000a900463ffffffff1681565b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561054b5780601f106105205761010080835404028352916020019161054b565b820191906000526020600020905b81548152906001019060200180831161052e57829003601f168201915b5050505050815600a165627a7a723058209880ebcdee911274c1800f4e2aa869f837c59521cf6346269b9684e7075318a20029`

// DeployContractDrive deploys a new Ethereum contract, binding an instance of ContractDrive to it.
func DeployContractDrive(auth *bind.TransactOpts, backend bind.ContractBackend, _kilometers string, _avgspeed string, _avgaccel string, _starttime uint32, _endtime uint32) (common.Address, *types.Transaction, *ContractDrive, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractDriveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractDriveBin), backend, _kilometers, _avgspeed, _avgaccel, _starttime, _endtime)
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
