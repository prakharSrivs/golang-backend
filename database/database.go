package database

import (
	"errors"
	"log"
	"sync"
)

var StoreIdCache = make(map[string]bool)

type JobStatus string
type Perimeter int

// Constant Grouping for three kinds of Job Statuses
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

type Result struct {
	ImageURL  string `json:"image_url"`
	Perimeter int    `json:"perimeter"`
}

type Job struct {
	JobId   int        `json:"job_id"`
	Status  JobStatus  `json:"status"`
	Visits  []Visit    `json:"visits"`
	Results []Result   `json:"results,omitempty"`
	Errors  []JobError `json:"error,omitempty"`
}

// Model for In Memory - Non Persistent Database
type JobStore struct {
	mu     sync.Mutex // for Mutual Exclusivity
	jobs   map[int]Job
	lastId int
}

// To Create a New Job Store
func CreateNewJobStore() *JobStore {
	return &JobStore{
		jobs:   make(map[int]Job),
		lastId: 0,
	}
}

// To Create a new Job
func (js *JobStore) CreateNewJob(visits []Visit) int {
	js.mu.Lock()
	defer js.mu.Unlock()

	js.lastId++
	jobId := js.lastId
	js.jobs[jobId] = Job{
		JobId:   jobId,
		Visits:  visits,
		Status:  JobOngoing,
		Results: []Result{},
	}

	return jobId
}

// To Fetch a Job
func (js *JobStore) GetJob(jobId int) (Job, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	job, exists := js.jobs[jobId]

	if !exists {
		return job, errors.New("JobId does not exist")
	}

	return job, nil
}

// To Update the Status of a Job
func (js *JobStore) UpdateJobStatus(jobId int, status JobStatus, jobErrors []JobError, jobResults []Result) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	job, exists := js.jobs[jobId]
	if !exists {
		log.Println("Job Id does not exist", jobId)
		return errors.New("JobId does not exist")
	}
	job.Status = status
	job.Errors = jobErrors
	job.Results = jobResults
	js.jobs[jobId] = job
	return nil
}
