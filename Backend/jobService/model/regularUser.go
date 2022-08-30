package model

type RegularUser struct {
	Id string `bson:"_id,omitempty"`
	Username 	string `bson:"username"`
}
