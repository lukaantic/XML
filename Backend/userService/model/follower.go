package model

type Follower struct {
	UserFollower RegularUser `json:"userFollower"`
	IsAccepted bool `json:"isAccepted"`
}