package dto

import (
	"userService/model"
)

type RegularUserPostDTO struct {
	Id          string             `bson:"_id"`
	PrivacyType *model.PrivacyType `bson:"privacyType"`
}
