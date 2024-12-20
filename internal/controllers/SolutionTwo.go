package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

/*
Reads the entities one by one and saves them in db one by one.
*/
func (c *MainController) SolutionTwo(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution two", c.Database, engines.SolutionTwo)(w, r)
}
