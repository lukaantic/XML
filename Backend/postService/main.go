package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"postService/handler"
	"postService/repository"
	"postService/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initPostRepository(database *mongo.Database) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

func initPostService(repository1 *repository.PostRepository) *service.PostService {
	return &service.PostService{PostRepository: repository1}
}

func initPostHandler(service *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{PostService: service}
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

	postDatabase := client.Database("post")

	return postDatabase
}

func handleFunc(handlerPost *handler.PostHandler) {
	ruter := mux.NewRouter().StrictSlash(true)
	ruter.HandleFunc("/new-post", handlerPost.CreateNewPost).Methods("POST")
	ruter.HandleFunc("/regular-user-posts/{username}", handlerPost.GetAllRegularUserPosts).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), ruter)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	postDatabase := initDatabase()
	postRepository := initPostRepository(postDatabase)
	postService := initPostService(postRepository)
	postHandler := initPostHandler(postService)

	handleFunc(postHandler)

}
