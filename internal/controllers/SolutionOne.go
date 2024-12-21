package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

func (c *MainController) SolutionOne(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("LoadAll And Insert In Batches", c.Database, engines.LoadAllAndInsertInBatches)(w, r)
}
