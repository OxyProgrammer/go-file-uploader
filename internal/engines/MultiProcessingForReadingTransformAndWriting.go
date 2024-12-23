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

type Record struct {
	LineNo int
	Data   []string
}

/*
Uses three buffered channels to read, transform, and write
entities in a pipeline fashion.
*/
func MultiProcessingForReadingTransformAndWriting(database *db.DB) error {
	recordCh := make(chan *Record, 1000)
	transformCh := make(chan *models.Land, 1000)
	doneCh := make(chan struct{})
	errCh := make(chan error)
	const numTransformers = 5 // Number of transformation goroutines
	const numWriters = 5      // Number of writing goroutines

	go readAndProduceRecords(recordCh, errCh)

	var wg sync.WaitGroup
	for i := 0; i < numTransformers; i++ {
		wg.Add(1)
		go transformAndProduceDbModel(recordCh, transformCh, errCh, &wg)
	}

	var writeWg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < numWriters; i++ {
		writeWg.Add(1)
		go writeAndConsumeDbModel(database, transformCh, errCh, &writeWg, &mutex)
	}

	go func() {
		wg.Wait()
		close(transformCh) // Close transformCh once all transform goroutines finish
	}()

	go func() {
		writeWg.Wait()
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

/*
Reads file and sends the records as string array to the channel
*/
func readAndProduceRecords(recordCh chan<- *Record, errCh chan<- error) {
	defer close(recordCh)
	file, err := os.Open("data/land_feed.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	_, _ = reader.Read() // Read the headers
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
		recordCh <- &Record{
			LineNo: lineNumber,
			Data:   record,
		}
		lineNumber++
	}
}

/*
Reads from record channel, converts to read model and then subsequently to db model.
*/
func transformAndProduceDbModel(recordCh <-chan *Record, transformCh chan<- *models.Land, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range recordCh {
		readEntity, err := utils.CreateEntityFromRecord(record.Data, record.LineNo)
		if err != nil {
			errCh <- err
			return
		}
		dbEntity := models.FromReadModel(*readEntity)
		transformCh <- dbEntity
	}
}

/*
Takes feed from a channel that has db models and writes into db.
*/
func writeAndConsumeDbModel(database *db.DB, transformCh <-chan *models.Land, errCh chan<- error, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	var buffer []*models.Land
	batchSize := 10000
	for dbEntity := range transformCh {
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
