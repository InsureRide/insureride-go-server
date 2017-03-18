package models

import "math/big"

type Car struct {
	ContractAddress string
	Brand           string
	Model           string
	Year            uint8
	Vehiclenumber   string
	Drives          []*Drive
	BalanceWei      *big.Int
}

type Drive struct {
	ContractAddress string
	Kilometers      float64
	Avgspeed        float64
	Avgaccel        float64
	Starttime       uint32
	Endtime         uint32
	PriceWei        *big.Int
}