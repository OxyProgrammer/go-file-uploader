package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

func (c *MainController) SolutionTwo(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("Read Line And And Insert In Batches", c.Database, engines.ReadLineAndAndInsertInBatches)(w, r)
}
