package engines

import (
	"errors"
	"go-file-uploader/internal/models"
	"go-file-uploader/pkg/db"
	"go-file-uploader/pkg/utils"
	"log"
)

/*
Loads all entities in the memory and save them in the db.
*/
func LoadAllAndInsert(database *db.DB) error {
	landReadModels, err := utils.ReadCSVAll("data/land_feed.csv")
	if len(err) > 0 {
		log.Fatal(err)
		return errors.New("Some error happened. Check logs.")
	}

	var dbLandModels []*models.Land

	for _, landReadModel := range landReadModels {
		dbLandModels = append(dbLandModels, models.FromReadModel(landReadModel))
	}

	eror := database.CreateLands(dbLandModels)
	return eror
}
