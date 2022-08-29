package model


type RegularUser struct {
	User  `bson:",inline"`
	BlockedUsers []int `bson:"regularUser,omitempty"`
	LikedPosts []string `bson:"likedPosts,omitempty"`
	DislikedPosts []string `bson:"dislikedPosts,omitempty"`
	ProfilePrivacy ProfilePrivacy `bson:",inline,omitempty"`
	Notifications []Notification `bson:"notifications,omitempty"`
	MediaContents []MediaContent `bson:"mediaContents,omitempty"`
	Followings []Following `bson:"followings,omitempty"`
	Followers []Follower `bson:"followers,omitempty"`
}