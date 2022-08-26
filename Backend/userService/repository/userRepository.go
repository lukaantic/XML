package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"userService/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularUserRepository struct {
	Database *mongo.Database
}

func (repository *RegularUserRepository) Register(user *model.RegularUser) (string, error){
	regularUserCollection := repository.Database.Collection("regularUsers")
	res, err := regularUserCollection.InsertOne(context.TODO(), &user)
	if err != nil {
		return "", fmt.Errorf("regular user is NOT created")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repository *RegularUserRepository) ExistByUsername(username string) bool{
	regularUserCollection := repository.Database.Collection("regularUsers")
	filterCursor, err := regularUserCollection.Find(context.TODO(), bson.M{"username": username})
	if err != nil {
		log.Fatal(err)
	}

	var userFiltered bson.D
	if err = filterCursor.All(context.TODO(), &userFiltered); err != nil {
		log.Fatal(err)
	}
	if len(userFiltered) != 0 {
		return true
	}
	return false
}