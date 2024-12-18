package models

import (
	"fmt"
	"strconv"
)

type LandRead struct {
	Address1 string
	Address2 string
	Acreage  float64
	Zoning   string
	Price    int
}

type ReadErrorDetails struct {
	LineNumber int
	ErrorText  string
}

func (l LandRead) ToCSVRow() []string {
	return []string{
		l.Address1,
		l.Address2,
		fmt.Sprintf("%.2f", l.Acreage),
		l.Zoning,
		strconv.Itoa(l.Price),
	}
}
