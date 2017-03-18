package models

import (
	"errors"
)



type Car struct {
	ContractAddress    string
	Brand string
	Model string
	Year uint8
	Vehiclenumber string;
	Drives []*Drive;
}


var (
	DriveListe map[string]*Drive
)

func init() {
	DriveListe = make(map[string]*Drive)
	u := Drive{"asdf", 222, 90, 1489826177, 1489826377}
	DriveListe[u.ContractAddress] = &u
}

type Drive struct {
	ContractAddress       string
	Kilometers uint64
	Avgspeed uint16
	Starttime uint32
	Endtime uint32
}


func AddDrive(d Drive) string {

	return d.ContractAddress
}

func GetDrive(dId string) (u *Drive, err error) {
	if drive, ok := DriveListe[dId]; ok {
		return drive, nil
	}
	return nil, errors.New("Drive not exists")
}

func GetAllDrives() map[string]*Drive {
	return DriveListe
}

