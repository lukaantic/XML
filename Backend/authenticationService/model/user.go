package model

import (

)

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Username string   `json:"username"`
	Password  string   `json:"password"`
	UserRole UserRole `json:"role"`
}
