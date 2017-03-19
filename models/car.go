package models


type Car struct {
	ContractAddress string
	Brand           string
	Model           string
	Year            uint16
	Vehiclenumber   string
	Drives          []*Drive
	BalanceInt         uint16
	Balance 	float32
}

type Drive struct {
	ContractAddress string
	Kilometers      float64
	Avgspeed        float64
	Avgaccel        float64
	Starttime       uint32
	Endtime         uint32
	PriceInt         uint16
	Price 		float32
}