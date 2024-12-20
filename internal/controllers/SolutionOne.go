package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

/*
Loads all entities in the memory and save them in the db.
*/
func (c *MainController) SolutionOne(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution one", c.Database, engines.SolutionOne)(w, r)
}
