package engines

import (
	"encoding/csv"
	"go-file-uploader/internal/models"
	"go-file-uploader/pkg/db"
	"go-file-uploader/pkg/utils"
	"io"
	"os"
)

/*
Reads the entities one by one and saves them in db one by one.
*/
func ReadLineAndAndInsertInBatches(database *db.DB) error {
	file, err := os.Open("data/land_feed.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	// Create a CSV reader
	reader := csv.NewReader(file)

	//Read the headers
	_, _ = reader.Read()
	lineNumber := 1
	var buffer []*models.Land

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// We've reached the end of the file, break out of the loop
				break
			}
			return err
		}

		// Create an entity from the CSV record
		readEntity, err := utils.CreateEntityFromRecord(record, lineNumber)

		if err != nil {
			return err
		}

		dbEntity := models.FromReadModel(*readEntity)
		buffer = append(buffer, dbEntity)

		if len(buffer) == 10000 {
			database.CreateLands(buffer)
			buffer = buffer[:0]
		}

		lineNumber++
	}

	// Flush any remaining entities in the buffer
	if len(buffer) > 0 {
		database.CreateLands(buffer)
	}

	return nil
}
