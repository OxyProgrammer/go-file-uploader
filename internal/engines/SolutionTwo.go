package engines

import "go-file-uploader/pkg/db"

func SolutionTwo(database *db.DB) error {
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	return nil
}
