package model

type Post struct {
	RegularUser `json:"user"`
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}
