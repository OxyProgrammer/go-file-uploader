package engines

import (
	"encoding/csv"
	"go-file-uploader/internal/models"
	"go-file-uploader/pkg/db"
	"go-file-uploader/pkg/utils"
	"io"
	"log"
	"os"
	"sync"
)

/*
Uses a buffered channel to read
entities and save entitites in chunks.
*/
func MiltiprocessingForReadingAndWriting(database *db.DB) error {

	readCh := make(chan *models.LandRead, 1000)
	doneCh := make(chan struct{})
	errCh := make(chan error)

	go readAndProduceAsync(readCh, errCh)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeAndConsumeAsync(database, readCh, errCh, &wg, &mutex)
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case <-doneCh:
			return nil
		}
	}
}

func readAndProduceAsync(readCh chan<- *models.LandRead, errCh chan<- error) {

	defer close(readCh)

	file, err := os.Open("data/land_feed.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Create a CSV reader
	reader := csv.NewReader(file)
	//Read the headers
	_, _ = reader.Read()

	lineNumber := 1

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return
			}
			errCh <- err
			return
		}

		readEntity, err := utils.CreateEntityFromRecord(record, lineNumber)
		if err != nil {
			errCh <- err
			return
		}

		readCh <- readEntity
		lineNumber++
	}
}

func writeAndConsumeAsync(database *db.DB, readCh <-chan *models.LandRead, errCh chan<- error, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	var buffer []*models.Land
	batchSize := 10000

	for readEntity := range readCh {
		dbEntity := models.FromReadModel(*readEntity)
		buffer = append(buffer, dbEntity)

		if len(buffer) == batchSize {
			mutex.Lock()
			err := database.CreateLands(buffer)
			mutex.Unlock()
			if err != nil {
				errCh <- err
				return
			}
			buffer = buffer[:0]
		}
	}

	// Flush any remaining entities in the buffer
	if len(buffer) > 0 {
		mutex.Lock()
		err := database.CreateLands(buffer)
		mutex.Unlock()
		if err != nil {
			errCh <- err
			return
		}
	}
}
