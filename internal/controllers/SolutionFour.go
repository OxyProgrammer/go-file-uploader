package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

func (c *MainController) SolutionFour(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("Advanced Multiprocessing For Reading And Writing", c.Database, engines.MultiProcessingForReadingTransformAndWriting)(w, r)
}
