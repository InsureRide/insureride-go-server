#!/bin/bash

# Navigate to Project directory
cd $GOPATH/src/github.com/scmo/starthack

# Remove old files
rm ./ethereum/smartcontract/*.go

# Generate new Bindings
abigen --abi ./ethereum/smartcontract/car.abi --pkg smartcontract --type ContractCar --out ./ethereum/smartcontract/car.go --bin ./ethereum/smartcontract/car.bin
abigen --abi ./ethereum/smartcontract/drive.abi --pkg smartcontract --type ContractDrive --out ./ethereum/smartcontract/drive.go --bin ./ethereum/smartcontract/drive.bin