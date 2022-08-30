package handler

import (
	//"fmt"
	"jobService/service"
	//"jobService/model"
	"encoding/json"
	"jobService/dto"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobHandler struct {
	JobService *service.JobService
}

func (handler *JobHandler) CreateNewJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jobUploadDto dto.JobUploadDTO
	err := json.NewDecoder(r.Body).Decode(&jobUploadDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.JobService.CreateNewJob(jobUploadDto)
	if err != nil {
		if err.Error() == "user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *JobHandler) GetAllRegularUserJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserJobs := handler.JobService.GetAllRegularUserJobs(username)
	regularUserJobsJson, err := json.Marshal(regularUserJobs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(regularUserJobsJson)
	}
}

func (handler *JobHandler) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allJobs := handler.JobService.GetAllJobs()
	jobsJson, err := json.Marshal(allJobs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jobsJson)
	}
}

func (handler *JobHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	jobId,err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.JobService.DeleteJob(jobId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *JobHandler) GetJobSearchResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	searchInput := param["searchInput"]
	searchJobs, err := handler.JobService.GetJobSearchResults(searchInput)
	searchJobsJson, err := json.Marshal(searchJobs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(searchJobsJson)
	}
}