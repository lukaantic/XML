package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type RegularUserRepository struct {
	Database *mongo.Database
}