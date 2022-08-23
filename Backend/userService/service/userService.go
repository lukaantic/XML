package service

import (
	"userService/repository"
)


type RegularUserService struct {
	RegularUserRepository *repository.RegularUserRepository
}
