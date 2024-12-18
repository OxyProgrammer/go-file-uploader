package utils

import (
	"encoding/csv"
	"errors"
	"go-file-uploader/internal/models"
	"log"
	"os"
)

func ReadCSVContinuous(filename string, entityChan chan<- models.LandRead, errorChan chan<- models.ReadErrorDetails) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Create a CSV reader
	reader := csv.NewReader(file)
	lineNumber := 0
	for {
		record, err := reader.Read()
		if err != nil {
			close(entityChan)
			close(errorChan)
			return
		}

		// Create an entity from the CSV record
		entity, err := createEntityFromRecord(record)

		if err != nil {
			errorChan <- models.ReadErrorDetails{
				LineNumber: 0,
				ErrorText:  err.Error(),
			}
		}

		// Send the entity to the channel
		entityChan <- *entity
		lineNumber++
	}
}

func ReadCSVAll(reader *csv.Reader) ([]models.LandRead, []error) {
	return nil, nil
}

func createEntityFromRecord(record []string) (*models.LandRead, error) {
	if nil == record {
		return nil, errors.New("record is nil")
	}
	return &models.LandRead{}, nil
}
