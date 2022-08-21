package service

import (
	"authenticationService/repository"

)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}