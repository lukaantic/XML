package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	//Skills      []Skills           `bson:"skills,omitempty"`
	Skills       string           `bson:"skills,omitempty"`
	Description string             `bson:"description,omitempty"`
	RegularUser						`bson:"regularuser,omitempty"`
	Name        string             `bson:"name,omitempty"`
}
