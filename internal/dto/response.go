package dto

import (
	"time"
)

type TimeResponse struct {
	Elapsed time.Duration `json:"elapsed"`
	Error   string        `json:"error"`
}
