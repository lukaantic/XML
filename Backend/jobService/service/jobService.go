package service

import (
	"fmt"
	"jobService/model"
	"jobService/repository"
	"jobService/dto"
	"net/http"
	"encoding/json"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type JobService struct {
	JobRepository *repository.JobRepository
}

func (service *JobService) CreateNewJob(jobUploadDto dto.JobUploadDTO) error {
	fmt.Println("Creating new job")

	job, err := createJobFromUploadDTO(&jobUploadDto)
	if err != nil {
		return err
	}
	_, err1 := service.JobRepository.Create(job)
	if err1 != nil {
		return err1
	}
	return nil
}

func createJobFromUploadDTO(jobDTO *dto.JobUploadDTO) (*model.Job, error) {
	regularUser, err := getRegularUserFromUsername(jobDTO.Username)
	if err != nil {
		return nil, err
	}
	var job model.Job
	
	job.Skills = jobDTO.Skills
	job.Description = jobDTO.Description
	job.RegularUser = *regularUser
	job.RegularUser.Username = jobDTO.Username

	return &job, nil
}

func getRegularUserFromUsername(username string) (*model.RegularUser, error) {
	requestUrl := fmt.Sprintf("http://%s:%s/find-user/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), username)

	resp, err := http.Get(requestUrl)
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var regularUser model.RegularUser
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&regularUser)

	return &regularUser, nil
}

func (service *JobService) GetAllRegularUserJobs(username string) []model.Job {
	regularUserJobDocuments := service.JobRepository.GetAllByUsername(username)

	regularUserJobs := CreateJobsFromDocuments(regularUserJobDocuments)
	return regularUserJobs
}

func CreateJobsFromDocuments(JobsDocuments []bson.D) []model.Job {
	var jobs []model.Job
	for i := 0; i < len(JobsDocuments); i++ {
		var job model.Job
		bsonBytes, _ := bson.Marshal(JobsDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &job)
		jobs = append(jobs, job)
	}
	return jobs
}

func (service *JobService) GetAllJobs() []model.Job {
	jobDocuments := service.JobRepository.GetAllJobs()

	publicPosts := CreateJobsFromDocuments(jobDocuments)
	return publicPosts
}

func (service *JobService) DeleteJob(id primitive.ObjectID) error{
	err := service.JobRepository.DeleteJob(id)
	if err != nil {
		return err
	}
	return nil
}