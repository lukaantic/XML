package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	)

type User struct{
	ID 		uuid.UUID	`json:"id"`
	Ime 	string		`json:"ime"`
	Prezime	string		`json:"prezime"`
	Uloga	string		`json:"uloga"`
	Username string		`json:"username"`
	Lozinka string		`json:"lozinka"`
	Email   string		`json:"email"`
	UserRole UserRole 	`json:"role"`
}	

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}