package engines

import (
	"errors"
	"go-file-uploader/internal/models"
	"go-file-uploader/pkg/db"
	"go-file-uploader/pkg/utils"
	"log"
)

func SolutionOne(database *db.DB) error {
	landReadModels, err := utils.ReadCSVAll("data/land_feed.csv")
	if err != nil && len(err) > 0 {
		log.Fatal(err)
		return errors.New("Some error happened. Check logs.")
	}

	var dbLandModels []*models.Land

	for _, landReadModel := range landReadModels {
		dbLandModels = append(dbLandModels, models.FromReadModel(landReadModel))
	}
	errr := database.CreateLands(dbLandModels)
	return errr
}
