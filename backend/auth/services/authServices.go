package service

import (
	//"fmt"

	"distlinkt.bab/backend/auth/model"
	"distlinkt.bab/backend/auth/repository"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) RegisterUser (user model.User) error {
	korisnik := model.User{}
	err := service.AuthRepository.CreateUser(&korisnik)
	if err != nil {
		return err 
	}
	return nil 
}

func (service *AuthService) UpdateUser (user model.User) error {
	korisnik, err := service.AuthRepository.FindUserByUsername(user.Username)
	if err != nil{
		return err
	}
	korisnik.Ime		= user.Ime
	korisnik.Prezime 	= user.Prezime
	korisnik.Username 	= user.Username
	korisnik.Lozinka	= user.Lozinka 	
	korisnik.Email		= user.Email
	err = service.AuthRepository.UpdateUser(korisnik)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) DeleteUser (id string) error {
	
	err := service.AuthRepository.DeleteUser(id)
	if err != nil{
		return err
	}
	return nil
}

func (service *AuthService) FindByUsername (user model.User) (*model.User, error){
	korisnik, err := service.AuthRepository.FindUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	return korisnik, nil
}