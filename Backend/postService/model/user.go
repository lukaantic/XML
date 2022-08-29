package model

type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Gender      Gender    `json:"gender"`
	UserRole    UserRole  `json:"userRole"`
}
