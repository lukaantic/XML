package model

type Following struct {
	UserFollowing RegularUser `json:"UserFollowing"`
	Notifying     bool        `json:"notifying"`
}
