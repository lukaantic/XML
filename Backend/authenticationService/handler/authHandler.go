package handler

import (
	"authenticationService/service"
	"net/http"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (hem *AuthHandler) RegisterUser (res http.ResponseWriter, req *http.Request){
	println("stampam")
} 