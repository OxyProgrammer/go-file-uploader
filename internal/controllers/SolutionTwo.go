package controllers

import "net/http"

/*
Reads the entities one by one and saves them in db one by one.
*/
func (c *MainController) SolutionTwo(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution two", sampleFunction)(w, r)
}
