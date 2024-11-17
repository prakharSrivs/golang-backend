package processor

import (
	"log"
	"math/rand"
	"time"

	"github.com/prakharsrivs/kirana-club-assignment/database"
	"github.com/prakharsrivs/kirana-club-assignment/helpers"
)

func generateRandomSleepDuration() int {
	randomInteger := rand.Intn(300) + 100
	return randomInteger
}

func ProcessJob(jobId int, visits []database.Visit, jobStore *database.JobStore) {
	var errors []database.JobError
	var results []database.Result

	for i := 0; i < len(visits); i++ {
		imageUrlsList := visits[i].ImageURLs
		storeId := visits[i].StoreID

		if !helpers.ValidateStoreId(storeId, database.StoreIdCache) {
			errors = append(errors, database.JobError{
				StoreId: storeId,
				Error:   "StoreId not present in the Provided Store Master CSV File",
			})
			continue
		}

		for _, imageUrl := range imageUrlsList {
			perimeter, err := helpers.CalculatePerimeter(imageUrl)
			if err != nil {
				errors = append(errors, database.JobError{StoreId: storeId, Error: err.Error()})
				continue
			}
			result := database.Result{ImageURL: imageUrl, Perimeter: perimeter}
			results = append(results, result)
			time.Sleep(time.Duration(generateRandomSleepDuration()) * time.Millisecond)
		}

	}

	status := database.JobCompleted
	if len(errors) > 0 {
		status = database.JobFailed
	}
	err := jobStore.UpdateJobStatus(jobId, status, errors, results)
	if err != nil {
		log.Panic("Failed to Update Job Status", jobId)
	}

}
