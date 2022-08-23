package main

import (
	"fmt"
	"userService/repository"
	"userService/service"
	"userService/handler"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initUserRepository(database *mongo.Database) *repository.RegularUserRepository {
	return &repository.RegularUserRepository{Database: database}
}

func initUserService(repository *repository.RegularUserRepository) *service.RegularUserService {
	return &service.RegularUserService{RegularUserRepository: repository}
}

func initUserHandler(service *service.RegularUserService) *handler.RegularUserHandler {
	return &handler.RegularUserHandler{RegularUserService: service}
}

func initDatabase() *mongo.Database{
	//clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")
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

	userDatabase := client.Database("user")

	return userDatabase
}
func handleFunc(userHandler *handler.RegularUserHandler) {
	fmt.Println("proba")
}

func main() {
	userDatabase := initDatabase()

	userRepository := initUserRepository(userDatabase)
	userService := initUserService(userRepository)
	userHandler := initUserHandler(userService)

	handleFunc(userHandler)

}