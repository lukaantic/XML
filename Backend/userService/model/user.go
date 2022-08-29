package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id             primitive.ObjectID  `bson:"_id,omitempty"`
	Name           string              `bson:"name,omitempty"`
	Surname        string              `bson:"surname,omitempty"`
	Username       string              `bson:"username,omitempty"`
	Password       string              `bson:"password,omitempty"`
	Email          string              `bson:"email,omitempty"`
	PhoneNumber    string              `bson:"phoneNumber,omitempty"`
	Gender         *Gender             `bson:"gender,omitempty"`
	BirthDate      *primitive.DateTime `bson:"birthDate,omitempty"`
	UserRole       UserRole            `bson:"userRole,omitempty"`
	Biography      string              `bson:"biography,omitempty"`
	Skills         string              `bson:"skills"`
	Education      string              `bson:"education"`
	Expirience     string              `bson:"expirience"`
	Interests      string              `bson:"interests"`
	LikedPosts     []string            `bson:"likedPosts,omitempty"`
	DislikedPosts  []string            `bson:"dislikedPosts,omitempty"`
	ProfilePrivacy ProfilePrivacy      `bson:",inline,omitempty"`
	Notifications  []Notification      `bson:"notifications,omitempty"`
	MediaContents  []MediaContent      `bson:"mediaContents,omitempty"`
	Followings     []Following         `bson:"followings,omitempty"`
	Followers      []Follower          `bson:"followers,omitempty"`
}
