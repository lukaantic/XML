package main

import (
	"context"
	"fmt"
	"jobService/handler"
	"jobService/repository"
	"jobService/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initJobRepository(database *mongo.Database) *repository.JobRepository {
	return &repository.JobRepository{Database: database}
}

func initJobService(repository *repository.JobRepository) *service.JobService {
	return &service.JobService{JobRepository: repository}
}

func initJobHandler(service *service.JobService) *handler.JobHandler {
	return &handler.JobHandler{JobService: service}
}

func initDatabase() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	jobsDatabase := client.Database("jobs")
	return jobsDatabase
}

func handleFunc(jobHandler *handler.JobHandler) {
	ruter := mux.NewRouter().StrictSlash(true)
	ruter.HandleFunc("/new-job", jobHandler.CreateNewJob).Methods("POST")
	ruter.HandleFunc("/get-all-jobs", jobHandler.GetAllJobs).Methods("GET")
	ruter.HandleFunc("/get-all-users-jobs/{username}", jobHandler.GetAllRegularUserJobs).Methods("GET")
	ruter.HandleFunc("/delete-job/{id}", jobHandler.DeleteJob).Methods("DELETE")
	ruter.HandleFunc("/job-search/{searchInput}", jobHandler.GetJobSearchResults).Methods("GET")

	c := SetupCors()

	http.Handle("/", c.Handler(ruter))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(ruter))
	if err != nil {
		log.Println(err)
	}

}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins, for now
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}

func main() {

	jobsDatabase := initDatabase()

	jobRepository := initJobRepository(jobsDatabase)
	jobService := initJobService(jobRepository)
	jobHandler := initJobHandler(jobService)

	handleFunc(jobHandler)
}
