package models

import (
	"math"
	"strings"
)

type Land struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Address string
	Acreage float64
	Zoning  string
	Price   int
}

func FromReadModel(readModel LandRead) *Land {
	address := strings.TrimSpace(readModel.Address1)
	if readModel.Address2 != "" {
		address += ", " + strings.TrimSpace(readModel.Address2)
	}

	return &Land{
		Address: address,
		Acreage: math.Round(readModel.Acreage*100) / 100, // Round to 2 decimal places
		Zoning:  strings.TrimSpace(readModel.Zoning),
		Price:   readModel.Price,
	}
}
