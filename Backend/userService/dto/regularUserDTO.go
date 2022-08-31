package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularUserDTO struct {
	Id       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Surname  string             `json:"surname"`
	Username string             `json:"username"`
	Skills      string               `json:"skills"`
	Interests      string               `json:"interests"`
	Experience      string               `json:"experience"`
	Education      string               `json:"education"`
}
