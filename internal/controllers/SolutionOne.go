package controllers

import (
	"encoding/json"
	"go-file-uploader/internal/dto"
	"net/http"
	"time"
)

func (c *MainController) SolutionOne(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//Some code here

	start := time.Now()
	// Simulate some processing time
	time.Sleep(2 * time.Second)
	elapsed := time.Since(start)

	resp := dto.TimeResponse{
		Elapsed: elapsed,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
