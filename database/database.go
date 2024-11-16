package database

import (
	"errors"
	"sync"
)

var StoreIdCache = make(map[string]bool)

type JobStatus string

const (
	JobOngoing   JobStatus = "ongoing"
	JobCompleted JobStatus = "completed"
	JobFailed    JobStatus = "failed"
)

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type JobError struct {
	StoreId string `json:"store_id"`
	Error   string `json:"error"`
}

type Job struct {
	JobId  int        `json:"job_id"`
	Status JobStatus  `json:"status"`
	Visits []Visit    `json:"visits"`
	Errors []JobError `json:"error,omitempty"`
}

type JobStore struct {
	mu     sync.Mutex
	jobs   map[int]Job
	lastId int
}

func CreateNewJobStore() *JobStore {
	return &JobStore{
		jobs:   make(map[int]Job),
		lastId: 0,
	}
}

func (js *JobStore) CreateNewJob(visits []Visit) int {
	js.mu.Lock()
	defer js.mu.Unlock()

	js.lastId++
	jobId := js.lastId
	js.jobs[jobId] = Job{
		JobId:  jobId,
		Visits: visits,
		Status: JobOngoing,
	}

	return jobId
}

func (js *JobStore) GetJob(jobId int) (Job, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	job, exists := js.jobs[jobId]

	if !exists {
		return job, errors.New("JobId does not exists")
	}

	return job, nil
}

func (js *JobStore) UpdateJobStatus(jobId int, status JobStatus, jobErrors []JobError) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	job, error := js.GetJob(jobId)

	if error != nil {
		return error
	}

	job.Status = status
	job.Errors = jobErrors
	return nil
}
