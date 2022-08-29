package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostUploadDTO struct {
	Description string              `json:"description"`
	MediaPaths  []string            `json:"mediaPaths"`
	UploadDate  *primitive.DateTime `json:"uploadDate"`
	Username    string              `json:"username"`
}
