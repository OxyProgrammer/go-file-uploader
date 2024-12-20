package controllers

import "go-file-uploader/pkg/db"

type MainController struct {
	Database *db.DB
}

func NewMainController(database *db.DB) *MainController {
	return &MainController{
		Database: database,
	}
}
