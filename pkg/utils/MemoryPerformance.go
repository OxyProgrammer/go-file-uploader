package utils

import (
	"go-file-uploader/internal/dto"
	"go-file-uploader/pkg/db"
	"runtime"
	"time"
)

func MeasurePerformance(database *db.DB, f func(database *db.DB) error) dto.Response {
	runtime.GC() // Run garbage collection before measurement

	var stats dto.Response
	var m runtime.MemStats

	// Measure start time and CPU time
	startTime := time.Now()

	// Run the function
	err := f(database)
	if err != nil {
		stats.Error = err.Error()
		return stats
	}
	// Measure end time and CPU time
	stats.ExecutionTime = time.Since(startTime)

	// Measure memory usage
	runtime.ReadMemStats(&m)
	stats.MemoryUsage = m.Alloc

	return stats
}
