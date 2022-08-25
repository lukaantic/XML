package model

type UserRole int

const(
	Registered UserRole = iota
	Unregistered
	Admin
)