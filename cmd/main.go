package main

import (
	"go-file-uploader/internal/controllers"
	"go-file-uploader/pkg/db"
	"log"
	"net/http"
)

func main() {

	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Created DB!")
	defer database.Close()

	mainController := controllers.NewMainController()

	http.HandleFunc("/", mainController.SolutionOne)

	log.Println("Listening on 8080!")
	http.ListenAndServe(":8080", nil)
}
