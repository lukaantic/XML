package service

import (
	//"fmt"

	"authenticationService/model"
	"authenticationService/repository"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) RegisterUser(user model.User) error {
	korisnik := model.User{}
	err := service.AuthRepository.CreateUser(&korisnik)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) UpdateUser(user model.User) error {
	korisnik, err := service.AuthRepository.FindUserByUsername(user.Username)
	if err != nil {
		return err
	}
	korisnik.Name 		= user.Name
	korisnik.Surname 	= user.Surname
	korisnik.Username 	= user.Username
	korisnik.Password 	= user.Password
	err = service.AuthRepository.UpdateUser(korisnik)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) DeleteUser(id string) error {

	err := service.AuthRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) FindByUsername(user model.User) (*model.User, error) {
	korisnik, err := service.AuthRepository.FindUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	return korisnik, nil
}
