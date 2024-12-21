package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"go-file-uploader/internal/models"
	"os"
	"strconv"
)

func ReadCSVAll(filename string) ([]models.LandRead, []error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, []error{err}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, []error{err}
	}

	var lands []models.LandRead
	var errors []error

	// Assuming the first row is headers, we start from index 1
	for i, record := range records[1:] {
		land, err := CreateEntityFromRecord(record, i)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		lands = append(lands, *land)
	}
	return lands, errors
}

func CreateEntityFromRecord(record []string, index int) (*models.LandRead, error) {

	if nil == record {
		return nil, errors.New("record is nil")
	}

	if len(record) != 5 {
		return nil, fmt.Errorf("row %d: invalid number of fields", index)
	}

	acreage, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return nil, fmt.Errorf("row %d: invalid acreage: %w", index, err)
	}

	price, err := strconv.Atoi(record[4])
	if err != nil {
		return nil, fmt.Errorf("row %d: invalid price: %w", index, err)
	}

	land := models.LandRead{
		Address1: record[0],
		Address2: record[1],
		Acreage:  acreage,
		Zoning:   record[3],
		Price:    price,
	}
	return &land, nil
}
