package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"distlinkt.bab/backend/auth/handler"
	"distlinkt.bab/backend/auth/model"
	"distlinkt.bab/backend/auth/repository"
	"distlinkt.bab/backend/auth/services"
	"gorm.io/driver/postgres"
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
		{Uloga: "admin", Ime: "Milovan",Prezime: "Antic",Lozinka: "2109",Username: "bebi",Email: "bebizr@gmail.com"},
		{Uloga: "admin", Ime: "Luka", Prezime: "Antic",Lozinka: "2810",Username: "luka", Email: "lukalazy@gmail.com"},
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

	http.ListenAndServe(":2109",ruter)

}