package dto

import (
	"postService/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostDTO struct {
	Id          string              `json:"Id"`
	Description string              `json:"Description"`
	MediaPaths  []string            `json:"MediaPaths"`
	UploadDate  *primitive.DateTime `json:"UploadDate"`
	RegularUser model.RegularUser   `json:"RegularUser"`
	Link        string              `json:"Link"`
	Likes       int                 `json:"Likes"`
	Dislikes    int                 `json:"Dislikes"`
	Comment     []model.Comment     `json:"Comment"`
}
