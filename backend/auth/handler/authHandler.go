package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"distlinkt.bab/backend/auth/utils"
	"distlinkt.bab/backend/auth/model"
	service "distlinkt.bab/backend/auth/services"

)

type AuthHandler struct {
	Handler *service.AuthService
}

func (handler *AuthHandler) RegisterUser(w http.ResponseWriter, req *http.Request) {
	var korisnik model.User
	err := json.NewDecoder(req.Body).Decode(&korisnik)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Handler.AuthRepository.CreateUser(&korisnik)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("Registrovao sam Korisnika")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(korisnik)
}

func (handler *AuthHandler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	var korisnik model.User
	err := json.NewDecoder(req.Body).Decode(&korisnik)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(err)
	err = handler.Handler.UpdateUser(korisnik)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("Urspesno uradjen Update")
	w.WriteHeader(http.StatusOK)
}

func (handler *AuthHandler) DeleteUser(w http.ResponseWriter, req *http.Request) {
	var korisnik model.User
	err := json.NewDecoder(req.Body).Decode(&korisnik)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	ajdi := korisnik.ID.String()
	err = handler.Handler.DeleteUser(ajdi)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Uspelo je brisanje!!!")

	w.WriteHeader(http.StatusOK)
}

func (handler *AuthHandler) Login(w http.ResponseWriter, req *http.Request) {
	var logKorisnik model.User
	err := json.NewDecoder(req.Body).Decode(&logKorisnik)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	var user *model.User
	user, err = handler.Handler.FindByUsername(logKorisnik)
	if err != nil {
		fmt.Println("Nije pronadjen korisnik")
		w.WriteHeader(http.StatusBadRequest)
		return		
	}
	if (user.Lozinka != logKorisnik.Lozinka){
		fmt.Println("Pogresna lozinka")	
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	ajdi := logKorisnik.ID.String()
	token, err := util.CreateJWT(ajdi, &user.UserRole, user.Username)
	response := model.LoginResponse{
		Token: token,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Uspesno logovanje!!!")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
	w.Header().Set("Content-Type", "application/json")
}
