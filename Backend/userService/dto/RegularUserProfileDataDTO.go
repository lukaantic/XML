package dto

import (
	"userService/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularUserProfileDataDTO struct {
	Id             primitive.ObjectID   `json:"_id"`
	Name           string               `json:"name"`
	Surname        string               `json:"surname"`
	Username       string               `json:"username"`
	Biography      string               `json:"biography"`
	ProfilePrivacy model.ProfilePrivacy `json:"profilePrivacy"`
	Skills      string               `json:"skills"`
	Interests      string               `json:"interests"`
	Experience      string               `json:"experience"`
	Education      string               `json:"education"`
}
