package main

import (
	"go-file-uploader/internal/controllers"
	"go-file-uploader/pkg/db"
	"go-file-uploader/pkg/utils"
	"log"
	"net/http"
	"os"
)

type Config struct {
	FolderName      string
	DbFileName      string
	FeedFileName    string
	FeedFileRowNums int64
}

func main() {

	c := &Config{
		FolderName:      "data",
		DbFileName:      "Db.sql",
		FeedFileName:    "land_feed.csv",
		FeedFileRowNums: 10000000,
	}

	//Create folder
	c.createFolder()
	//Feed file
	c.createFeedFile()

	//Database
	database, err := db.NewDB(c.FolderName + "/" + c.DbFileName)
	if err != nil {
		log.Panic("Error creating db. Quitting!")
	}
	defer database.Close()

	mainController := controllers.NewMainController(database)

	http.HandleFunc("/solution-one", mainController.SolutionOne)
	http.HandleFunc("/solution-two", mainController.SolutionTwo)
	http.HandleFunc("/solution-three", mainController.SolutionThree)
	http.HandleFunc("/solution-four", mainController.SolutionFour)

	log.Println("Listening on 8080!")
	http.ListenAndServe(":8080", nil)
}

func (c *Config) createFolder() {
	err := os.Mkdir(c.FolderName, 0755)
	if err != nil {
		if os.IsExist(err) {
			// Folder already exists
			log.Println("Folder already exists:", c.FolderName)
		} else {
			// Handle the error
			log.Panic("Error creating folder:", err)
			return
		}
	} else {
		log.Println("Folder created:", c.FolderName)
	}
}

func (c *Config) createFeedFile() {
	value, err := utils.CreateCSVFeedFile(c.FolderName+"/"+c.FeedFileName, c.FeedFileRowNums)
	if err != nil {
		log.Panic("Error creating feed file.")
	}
	if value {
		log.Println("Feed file created!")
	} else {
		log.Println("Feed file already exists!")
	}
}
