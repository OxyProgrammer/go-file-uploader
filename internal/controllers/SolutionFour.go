package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

func (c *MainController) SolutionFour(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("Multiprocessing For Reading, Transform And Writing", c.Database, engines.MultiProcessingForReadingTransformAndWriting)(w, r)
}
