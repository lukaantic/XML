package model

type RegularUser struct {
	User           `bson:",inline"`
	BlockedUsers   []int          `bson:"regularUser,omitempty"`
	LikedPosts     []string       `bson:"likedPosts,omitempty"`
	DislikedPosts  []string       `bson:"dislikedPosts,omitempty"`
	ProfilePrivacy ProfilePrivacy `bson:",inline,omitempty"`
	Notifications  []Notification `bson:"notifications,omitempty"`
	Followings     []Following    `bson:"followings,omitempty"`
	Followers      []Follower     `bson:"followers,omitempty"`
	Skills         []Skills       `bson:"skills,omitempty"`
	Education      string         `bson:"education"`
	Expirience     string         `bson:"expirience"`
	Interests      string         `bson:"interests"`
}
