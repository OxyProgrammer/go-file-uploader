package controllers

import (
	"go-file-uploader/internal/engines"
	"net/http"
)

func (c *MainController) SolutionOne(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution one", engines.SolutionOne)(w, r)
}
