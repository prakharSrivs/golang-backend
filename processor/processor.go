package processor

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prakharsrivs/kirana-club-assignment/database"
)

func generateRandomSleepDuration() int {
	randomInteger := rand.Intn(3000) + 1000
	return randomInteger
}

func ProcessJob(jobId int, visits []database.Visit, jobStore *database.JobStore) {
	var errors []database.JobError

	for i := 0; i < len(visits); i++ {
		imageUrlsList := visits[i].ImageURLs
		storeId := visits[i].StoreID

		for _, imageUrl := range imageUrlsList {
			_, err := http.Get(imageUrl)
			if err != nil {
				errors = append(errors, database.JobError{
					StoreId: storeId,
					Error:   "Failed to Download Image, " + err.Error(),
				})
				continue
			}
			time.Sleep(time.Duration(generateRandomSleepDuration()) * time.Millisecond)
		}

	}

	status := database.JobCompleted
	if len(errors) > 0 {
		status = database.JobFailed
	}
	err := jobStore.UpdateJobStatus(jobId, status, errors)
	if err != nil {
		log.Panic("Failed to Update Job Status", jobId)
	}

}
