package controllers

import (
	"encoding/json"
	"go-file-uploader/pkg/db"
	"go-file-uploader/pkg/utils"
	"net/http"
)

func (c *MainController) createSolutionHandler(solutionName string, database *db.DB, f func(database *db.DB) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := utils.MeasurePerformance(database, f)
		response.Message = "Successfully completed request for " + solutionName + "."

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
