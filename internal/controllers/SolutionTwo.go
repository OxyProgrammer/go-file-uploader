package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

func (c *MainController) SolutionTwo(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution two", c.Database, engines.ReadTheFileAndInsertInBatches)(w, r)
}
