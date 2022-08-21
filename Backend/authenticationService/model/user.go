package model

type User struct {
	Id int
	Name string
	Surname string
	Email string
	Password string
	Role UserRole
	Gender Gender
}