package main

import (
	"authenticationService/handler"
	"authenticationService/model"
	"authenticationService/repository"
	service "authenticationService/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "auth_baza"
	//host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "korisnici"
  )

func main() {

	baza := initBaza()
	repo := initRepo(baza)
	serv := initServis(repo)
	hend := initHandler(serv)

	fmt.Println("Proba")

	handlerFunkcija(hend)


}

func initBaza() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil{
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})
	
	
	useri:= []model.User{
		{UserRole: 3, Name: "Luka", Surname: "Antic",Password: "2810",Username: "luka"},
		{UserRole: 3, Name: "Vladan",Surname: "Lalic",Password: "2302",Username: "lala"},	
	}

	for _, user := range useri {
		db.Create(&user)
	}
	return db
}

func initRepo(database *gorm.DB) *repository.AuthRepository {
	return &repository.AuthRepository{Database: database}
}

func initServis(repository *repository.AuthRepository) *service.AuthService{
	return &service.AuthService{AuthRepository: repository}
}

func initHandler(service *service.AuthService) *handler.AuthHandler{
	return &handler.AuthHandler{Handler: service }
}

func handlerFunkcija (handler *handler.AuthHandler){

	ruter := mux.NewRouter().StrictSlash(true)

	ruter.HandleFunc("/register",handler.RegisterUser).Methods("POST")
	// za update je potrebno da polje 'username' bude popunjeno
	ruter.HandleFunc("/update",handler.UpdateUser).Methods("POST")
	// za delete je potreban 'id'
	ruter.HandleFunc("/delete",handler.DeleteUser).Methods("POST")
	// za login treba 'username' i 'lozinka'
	ruter.HandleFunc("/login",handler.Login).Methods("POST")	
	fmt.Println("sad cu da slusam")

	http.ListenAndServe(":8081",ruter)

}