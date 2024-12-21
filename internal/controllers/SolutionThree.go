package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

func (c *MainController) SolutionThree(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("Miltiprocessing For Reading And Writing", c.Database, engines.MultiprocessingForReadingAndWriting)(w, r)
}
