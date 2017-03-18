package ethereum

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"math/big"
	"github.com/scmo/insureride-go-server/ethereum/smartcontract"
	"github.com/scmo/insureride-go-server/models"
)

type EthereumController struct {
	Auth   *bind.TransactOpts
	Client *ethclient.Client
}

var ethereumController EthereumController

func Init() {
	dataDir := beego.AppConfig.String("dataDir")
	if beego.BConfig.RunMode == "dev" {
		dataDir = dataDir + "testnet/"
	}

	ipcfile := dataDir + "geth.ipc"

	// Create an IPC based RPC connection to a remote node
	client, err := ethclient.Dial(ipcfile)
	if err != nil {
		beego.Critical("Failed to connect to the Ethereum client: ", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(beego.AppConfig.String("systemAccount")), beego.AppConfig.String("systemAccountPassword"))
	if err != nil {
		beego.Critical("Failed to create authorized transactor: ", err)
	}
	ethereumController = EthereumController{Auth:auth, Client:client}
}

func getContractCarSession(contractCar *smartcontract.ContractCar) (*smartcontract.ContractCarSession) {
	contractCarSession := &smartcontract.ContractCarSession{
		Contract: contractCar,
		CallOpts: bind.CallOpts{Pending:true},
		TransactOpts: bind.TransactOpts{
			From:ethereumController.Auth.From,
			Signer:ethereumController.Auth.Signer,
			GasLimit:big.NewInt(3141592),
		},
	}
	return contractCarSession
}

func getContractDriveSession(contractDrive *smartcontract.ContractDrive) (*smartcontract.ContractDriveSession) {
	contractDriveSession := &smartcontract.ContractDriveSession{
		Contract: contractDrive,
		CallOpts: bind.CallOpts{Pending:true},
		TransactOpts: bind.TransactOpts{
			From:ethereumController.Auth.From,
			Signer:ethereumController.Auth.Signer,
			GasLimit:big.NewInt(3141592),
		},
	}
	return contractDriveSession
}


func GetCar(car *models.Car) (*models.Car, error) {
	contractcar, err := smartcontract.NewContractCar(common.HexToAddress(car.ContractAddress), ethereumController.Client)
	if err != nil {
		beego.Critical("Failed to instantiate a Token contract: %v", err)
	}
	session := getContractCarSession(contractcar)
	car.Brand, err = session.Brand()
	car.Model, err = session.Model()
	car.Year, err = session.Year()
	car.Vehiclenumber, err = session.Vehiclenumber()

	drivescount, _ := session.Nodrives()
	beego.Debug("Number of drives: ", drivescount)



	bigstr := drivescount.String()
	numberofdrives, err := strconv.Atoi(bigstr)
	car.Drives = make([]*models.Drive,numberofdrives)
	for i := 0; i < numberofdrives; i++ {
		add, err := session.Drives(big.NewInt(int64(i)));
		if(err != nil){
			beego.Critical("Failed getting arraz ", err)
		}
		drive := models.Drive{}
		drive.ContractAddress = add.String()

		GetDrive(&drive)
		car.Drives[i] = &drive
	}
	return car, err
}

func GetDrive(drive *models.Drive) (*models.Drive, error) {
	contractDrive, err := smartcontract.NewContractDrive(common.HexToAddress(drive.ContractAddress), ethereumController.Client)
	if err != nil {
		beego.Critical("Failed to instantiate a Token contract: %v", err)
	}
	session := getContractDriveSession(contractDrive)
	drive.Kilometers, _ = session.Kilometers()
	drive.Avgspeed, _ = session.Avgspeed()
	_ = session
	return drive, err
}

// Deploys a drive contract to the blockchain
func AddDrive(d models.Drive) (models.Drive, error) {
	address, tx, _, err := smartcontract.DeployContractDrive(ethereumController.Auth, ethereumController.Client, d.Kilometers, d.Avgspeed, d.Starttime, d.Endtime)
	if err != nil {
		beego.Critical("Failed to deploy new token contract: ", err)
	}
	beego.Info("Contract pending deploy: ", address.String())
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
	d.ContractAddress = address.String()
	return d, err
}

func AddDriveToCar(carContractAddress string, driveContractAddress string){
	contractcar, err := smartcontract.NewContractCar(common.HexToAddress(carContractAddress), ethereumController.Client)
	if err != nil {
		beego.Critical("Failed to instantiate a Token contract: %v", err)
	}
	session := getContractCarSession(contractcar)

	tx, err := session.AddDrive(common.HexToAddress(driveContractAddress))
	if(err != nil){
		beego.Critical(err)
	}
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
}