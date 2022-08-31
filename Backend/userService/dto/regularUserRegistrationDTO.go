package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userService/model"
)

type RegularUserRegistrationDTO struct {
	Id          primitive.ObjectID  `json:"_id"`
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
	Experience  string              `json:"experience"`
	Interests   string              `json:"interests"`
}
