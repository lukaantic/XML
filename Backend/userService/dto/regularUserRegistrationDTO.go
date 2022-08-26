package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userService/model"
)

type RegularUserRegistrationDTO struct {
	Name        string              `json:"name"`
	Surname     string              `json:"surname"`
	Username    string              `json:"username"`
	Password    string              `json:"password"`
	Email       string              `json:"email"`
	PhoneNumber string              `json:"phoneNumber"`
	Gender      *model.Gender       `json:"gender"`
	BirthDate   *primitive.DateTime `json:"birthDate"`
	Biography   string              `json:"biography"`
	WebSite     string              `json:"webSite"`
	Skills      string              `json:skills`
	Education   string              `json:"education"`
	Expirience  string              `json:"expirience"`
	Interests   string              `json:"interests"`
}
