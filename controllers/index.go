package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/prakharsrivs/kirana-club-assignment/database"
)

var JobStore = database.CreateNewJobStore()

// Error Struct for all the Http Errors
type HttpError struct {
	ErrorMsg string `json:"error"`
}

// Incoming Request Struct For JobSubmissionController
type JobRequest struct {
	Count  int              `json:"count"`
	Visits []database.Visit `json:"visits"`
}

// Outgoing Response For JobSubmissionController
type JobResponse struct {
	JobId int `json:"job_id"`
}

func ReplyError(w http.ResponseWriter, err string, errorStatus int) {
	errorResponse := HttpError{
		ErrorMsg: err,
	}
	w.WriteHeader(errorStatus)
	json.NewEncoder(w).Encode(errorResponse)
}

// Submit Job - /api/submit
func JobSubmissionController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Job Submission Request Recieved")
	w.Header().Set("Content-Type", "application/json")

	var jobReq JobRequest
	err := json.NewDecoder(r.Body).Decode(&jobReq)

	if err != nil || jobReq.Count != len(jobReq.Visits) {
		ReplyError(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	jobId := JobStore.CreateNewJob(jobReq.Visits)
	jobRes := &JobResponse{
		JobId: jobId,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(jobRes)
}

// Outgoing JobInfo Response for JobInfoController
type JobInfoResponse struct {
	Status database.JobStatus  `json:"status"`
	JobId  int                 `json:"job_id"`
	Errors []database.JobError `json:"error,omitempty"`
}

// Get Job Info Request Controller - /api/status?jobId=123
func JobInfoController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Job Info Request Recieved")
	w.Header().Set("Content-Type", "application/json")

	jobIdString := r.URL.Query().Get("jobid")

	if jobIdString == "" {
		ReplyError(w, "JobId Missing", http.StatusBadRequest)
		return
	}

	jobId, err := strconv.Atoi(jobIdString)

	if err != nil {
		ReplyError(w, "Invalid JobId", http.StatusBadRequest)
		return
	}

	job, error := JobStore.GetJob(jobId)

	if error != nil {
		ReplyError(w, "JobId does not exist", http.StatusBadRequest)
		return
	}

	jobResponse := JobInfoResponse{
		Status: job.Status,
		JobId:  jobId,
		Errors: job.Errors,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobResponse)
}