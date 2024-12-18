package models

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
