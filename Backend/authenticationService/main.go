package main

import (
	"authenticationService/handler"
	"authenticationService/model"
	"authenticationService/repository"
	"authenticationService/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initAuthRepository(database *gorm.DB) *repository.AuthRepository {
	return &repository.AuthRepository{Database: database}
}

func initAuthService(repository *repository.AuthRepository) *service.AuthService {
	return &service.AuthService{AuthRepository: repository}
}

func initAuthHandler(service *service.AuthService) *handler.AuthHandler {
	return &handler.AuthHandler{AuthService: service}
}

func initDatabase() *gorm.DB {
	var database *gorm.DB
	//err := godotenv.Load()

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"postgres", "admin", "auth-service", "5432")
	log.Print("Connecting to PostgreSQL DB...")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Otvorio sam bazu")
	database.AutoMigrate(&model.User{})

	return database
}
func handleFunc(handler *handler.AuthHandler) {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.RegisterUser).Methods("POST")

	http.ListenAndServe(":3344", router)
}

func main() {

	database := initDatabase()
	rep := initAuthRepository(database)
	ser := initAuthService(rep)
	hen := initAuthHandler(ser)

	handleFunc(hen)
}
