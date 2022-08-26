package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"userService/handler"
	"userService/repository"
	"userService/service"

	"github.com/gorilla/mux"
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

func initDatabase() *mongo.Database {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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

	ruter := mux.NewRouter().StrictSlash(true)
	ruter.HandleFunc("/register", userHandler.Register).Methods("POST")
	ruter.HandleFunc("/update", userHandler.UpdatePersonalInformations).Methods("PUT")
	ruter.HandleFunc("/delete", userHandler.DeleteRegularUser).Methods("DELETE")
	ruter.HandleFunc("/update-profile-privacy", userHandler.UpdateProfilePrivacy).Methods("PUT")
	ruter.HandleFunc("/find-user/{username}", userHandler.FindRegularUserByUsername).Methods("GET")
	ruter.HandleFunc("/public-regular-users", userHandler.GetAllPublicRegularUsers).Methods("GET")




	http.ListenAndServe(":1231", ruter)
}

func main() {
	userDatabase := initDatabase()

	userRepository := initUserRepository(userDatabase)
	userService := initUserService(userRepository)
	userHandler := initUserHandler(userService)

	handleFunc(userHandler)

}
