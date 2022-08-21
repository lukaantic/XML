package repository

import (
	"authenticationService/model"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Database *gorm.DB
}

func (repository *AuthRepository) RegisterUser(user *model.User) error {
	res := repository.Database.Create(user)
	if res.RowsAffected == 0 {
		return fmt.Errorf("Korisnik nije registrovan")
	}
	fmt.Println("Korisnik uspesno registrovan")
	return nil
}
