package models

type Land struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Address string
	Acreage float64
	Zoning  string
	Price   int
}
