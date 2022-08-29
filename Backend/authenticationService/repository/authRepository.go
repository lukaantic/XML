package repository

import (
	"authenticationService/model"
	"fmt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Database *gorm.DB
}

func (repo *AuthRepository) CreateUser (user *model.User) error{
	uradjeno := repo.Database.Create(user)
	if uradjeno.RowsAffected == 0 {

		return fmt.Errorf("Korisnik nije kreiran!!")
	}
	fmt.Println("Korisnik je uspesno kreiran")
	return  nil
}

func (repo *AuthRepository) UpdateUser (user *model.User) error {
	uradjeno := repo.Database.Updates(user)
	if uradjeno.RowsAffected == 0 {
		return fmt.Errorf("Promena podataka nije uradjena!!!")
	}
	fmt.Println("Uspesno izmenjeni podaci")
	return nil
}

func (repo *AuthRepository) DeleteUser (id string) error {
	err := repo.Database.Exec("DELETE FROM users WHERE id = ? ",id).Error
	if err != nil {
		return err
	}
	return nil
}
/*
func (repo *AuthRepository) FindUserByUserNameOld (username string) (*model.User, error){
	korisnik := &model.User{}
	err := repo.Database.Table("users").First(&korisnik, "username = ?", username).Error
	if err != nil {
		fmt.Println("Ne postoji korisnik sa tim nalogom")
		return nil, err
	}
	return korisnik, nil
}
*/

func (repo *AuthRepository) FindUserByUsername(username string) (*model.User, error){
	user := &model.User{}
	err := repo.Database.Table("users").First(&user, "username = ?", username).Error
	if  err != nil {
		return nil, err
	}
	return user, nil
}


func (repo *AuthRepository) FindById (id string) (*model.User, error) {
	korisnik := &model.User{}
	err := repo.Database.Table("users").First(&korisnik, "id = ?", id).Error
	if err != nil {
		fmt.Println("Ne postoji korisnik sa tim nalogom")
		return nil, err 
	}
	return korisnik, nil
}
