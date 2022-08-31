package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"userService/handler"
	"userService/postserver"
	"userService/repository"
	"userService/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	fmt.Println("cekam zahteve")

	ruter := mux.NewRouter().StrictSlash(true)
	ruter.HandleFunc("/register", userHandler.Register).Methods("POST")
	ruter.HandleFunc("/update", userHandler.UpdatePersonalInformations).Methods("PUT")
	ruter.HandleFunc("/delete", userHandler.DeleteRegularUser).Methods("DELETE")
	ruter.HandleFunc("/update-profile-privacy", userHandler.UpdateProfilePrivacy).Methods("PUT")
	ruter.HandleFunc("/find-user/{username}", userHandler.FindRegularUserByUsername).Methods("GET")
	ruter.HandleFunc("/public-regular-users", userHandler.GetAllPublicRegularUsers).Methods("GET")
	ruter.HandleFunc("/by-username/{username}", userHandler.CreateRegularUserPostDTOByUsername).Methods("GET")

	c := SetupCors()

	http.Handle("/", c.Handler(ruter))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("USER_SERVICE_PORT")), c.Handler(ruter))
	//err := http.ListenAndServe(":8081", ruter)
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

	server, err := postserver.NewPostServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer server.CloseTracer()
	defer server.CloseDB()

	userDatabase := initDatabase()

	userRepository := initUserRepository(userDatabase)
	userService := initUserService(userRepository)
	userHandler := initUserHandler(userService)

	handleFunc(userHandler)

}
