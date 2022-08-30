package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"jobService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"
)

type JobRepository struct {
	Database *mongo.Database

}

func (repository *JobRepository) Create(job *model.Job) (string, error) {
	jobsCollection := repository.Database.Collection("jobs")
	res, err := jobsCollection.InsertOne(context.TODO(), &job)
	if err != nil {
		return "", fmt.Errorf("post is NOT created")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repository *JobRepository) GetAllByUsername(username string) []bson.D{
	jobsCollection := repository.Database.Collection("jobs")
	filterCursor, err := jobsCollection.Find(context.TODO(), bson.M{"regularuser.username": username})
	if err != nil {
		log.Fatal(err)
	}
	
	var jobsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &jobsFiltered); err != nil {
		log.Fatal(err)
	}
	return jobsFiltered
}

func (repository *JobRepository) GetAllJobs() []bson.D{
	jobsCollection := repository.Database.Collection("jobs")
	cursor, err := jobsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var jobs []bson.D
	if err = cursor.All(context.TODO(), &jobs); err != nil {
		log.Fatal(err)
	}
	return jobs
}

func (repository *JobRepository) DeleteJob(id primitive.ObjectID) error{

	jobsCollection := repository.Database.Collection("jobs")
	_, err := jobsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (repository *JobRepository) GetJobSearchResults(searchInput string, allJobs []model.Job) []model.Job{
	var searchResults []model.Job
	for i := 0; i<len(allJobs); i++{
		if(strings.Contains(strings.ToLower(allJobs[i].Name), strings.ToLower(searchInput))){
			searchResults = append(searchResults, allJobs[i])
		}
	}
	return searchResults
}