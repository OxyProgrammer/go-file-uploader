package dto

import (
	"time"
)

type Response struct {
	Error         string        `json:"error"`
	Message       string        `json:"message"`
	MemoryUsage   uint64        `json:"memoryUsage"`
	CPUTime       time.Duration `json:"cpuTime"`
	ExecutionTime time.Duration `json:"elapsed"`
}
