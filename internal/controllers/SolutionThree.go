package controllers

import "net/http"

/*
Uses a buffered channel to read
entities and save entitites in chaunks.
*/

func (c *MainController) SolutionThree(w http.ResponseWriter, r *http.Request) {
	c.createSolutionHandler("solution three", sampleFunction)(w, r)
}
func sampleFunction() error {
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	return nil
}
