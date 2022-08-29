package model
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Post struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Description string `bson:"description,omitempty"`
	MediaPaths []string `bson:"mediaPaths,omitempty"`
	UploadDate *primitive.DateTime `bson:"uploadDate,omitempty"`
	RegularUser RegularUser `bson:"regularUser,omitempty"`
	Likes    int       `bson:"likes"`
	Dislikes int       `bson:"dislikes"`
	Comment  []Comment `bson:"comments,omitempty"`
}
