package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

/*
Uses a buffered channel to read
entities and save entitites in chaunks.
*/

func (c *MainController) SolutionThree(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution three", engines.SolutionThree)(w, r)
}
