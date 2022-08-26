package service

import (
	"userService/repository"
	"fmt"
	"userService/dto"
	"userService/model"
	"net/http"
	//"os"
	"bytes"
	"encoding/json"
)


type RegularUserService struct {
	RegularUserRepository *repository.RegularUserRepository
}

func (service *RegularUserService) Register(regularUserRegistrationDto dto.RegularUserRegistrationDTO) error {
	fmt.Println("Creating regular user")

	if service.RegularUserRepository.ExistByUsername(regularUserRegistrationDto.Username) {
		return fmt.Errorf("username is already taken")
	}

	var regularUser = createRegularUserFromRegularUserRegistrationDTO(&regularUserRegistrationDto)
	createdUserId, err := service.RegularUserRepository.Register(regularUser)
	if err != nil {
		return err
	}
	err2 := service.registerUserInAuthenticationService(regularUserRegistrationDto, createdUserId)
	if err2 != nil {
		return err2
	}
	return nil
}

func createRegularUserFromRegularUserRegistrationDTO(regularUserDto *dto.RegularUserRegistrationDTO) *model.RegularUser{
	profilePrivacy := model.ProfilePrivacy{
		PrivacyType: model.PrivacyType(0),
		AllMessageRequests: true,
	}
	var regularUser model.RegularUser
	regularUser.Name = regularUserDto.Name
	regularUser.Surname = regularUserDto.Surname
	regularUser.Username = regularUserDto.Username
	regularUser.Password = regularUserDto.Password
	regularUser.Email = regularUserDto.Email
	regularUser.PhoneNumber = regularUserDto.PhoneNumber
	regularUser.BirthDate = regularUserDto.BirthDate
	regularUser.Biography = regularUserDto.Biography
	regularUser.ProfilePrivacy = profilePrivacy
	regularUser.UserRole = model.UserRole(0)
	regularUser.Gender = regularUserDto.Gender

	return &regularUser
}

func (service *RegularUserService) registerUserInAuthenticationService(regularUserRegistrationDto dto.RegularUserRegistrationDTO, createdUserId string) error {
	postBody, _ := json.Marshal(map[string]string{
		"userId":   createdUserId,
		"email":    regularUserRegistrationDto.Email,
		"password": regularUserRegistrationDto.Password,
		"username": regularUserRegistrationDto.Username,
		"name":     regularUserRegistrationDto.Name,
		"surname":  regularUserRegistrationDto.Surname,
	})
	requestUrl := "http://localhost:8080/register"
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}