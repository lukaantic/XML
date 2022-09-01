package model

type RegularUser struct {
	User
	IsDisabled   bool        `json:"IsDisabled"`
	BlockedUsers []int       `json:"blockedUsers"`
	Followings   []Following `json:"followings"`
	Followers    []Follower  `json:"followers"`
	Username     string      `bson:"username"`
}
