package utils

import (
	"encoding/csv"
	"fmt"
	"go-file-uploader/internal/models"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func CreateCSVFeedFile(filename string, numRows int64) (bool, error) {

	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return false, nil
	}

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Address1", "Address2", "Acreage", "Zoning", "Price"}
	if err := writer.Write(header); err != nil {
		panic(err)
	}

	// Create a channel to receive property rows
	rowChan := make(chan []string, 1000)

	// Start goroutine to write rows
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for row := range rowChan {
			if err := writer.Write(row); err != nil {
				panic(err)
			}
		}
	}()

	// Generate properties
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	zoningTypes := []string{"Residential", "Commercial", "Agricultural", "Industrial", "Mixed-Use"}

	var start, end int64 = 1, numRows
	for i := start; i <= end; i++ {
		land := models.LandRead{
			Address1: fmt.Sprintf("%d Sector No ", rng.Intn(10000)+1),
			Address2: fmt.Sprintf("%d Land Rd", rng.Intn(10000)+1),
			Acreage:  float64(rand.Intn(1000)+1) + rng.Float64(),
			Zoning:   zoningTypes[rng.Intn(len(zoningTypes))],
			Price:    rng.Intn(10000000) + 10000,
		}
		rowChan <- land.ToCSVRow()

		if i%100000 == 0 {
			log.Printf("Generated %d rows", i)
		}
	}

	close(rowChan)
	wg.Wait()

	return true, nil
}
