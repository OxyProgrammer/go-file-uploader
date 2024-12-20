package controllers

import (
	"net/http"
)

func (c *MainController) SolutionOne(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution one", sampleFunction)(w, r)
}
