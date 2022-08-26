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
	"go.mongodb.org/mongo-driver/bson/primitive"

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


func createRegularUserFromRegularUserUpdateDTO(userUpdateDto *dto.RegularUserUpdateDTO) *model.RegularUser{
	id, _ := primitive.ObjectIDFromHex(userUpdateDto.Id)
	var regularUser model.RegularUser
	regularUser.Id = id
	regularUser.Name = userUpdateDto.Name
	regularUser.Surname = userUpdateDto.Surname
	regularUser.Username = userUpdateDto.Username
	regularUser.Email = userUpdateDto.Email
	regularUser.PhoneNumber = userUpdateDto.PhoneNumber
	regularUser.Gender= userUpdateDto.Gender
	regularUser.BirthDate = userUpdateDto.BirthDate
	regularUser.Biography = userUpdateDto.Biography

	return &regularUser
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

func (service *RegularUserService) UpdatePersonalInformations(regularUserUpdateDto dto.RegularUserUpdateDTO) error {
	fmt.Println("Updating regular user")

	if service.RegularUserRepository.ExistByUsername(regularUserUpdateDto.Username) {
		id, _ := primitive.ObjectIDFromHex(regularUserUpdateDto.Id)
		if service.RegularUserRepository.UsernameChanged(regularUserUpdateDto.Username, id) {
			return fmt.Errorf("username is already taken")
		}
	}
	id := regularUserUpdateDto.Id
	var regularUser = createRegularUserFromRegularUserUpdateDTO(&regularUserUpdateDto)
	err := service.RegularUserRepository.UpdatePersonalInformations(regularUser)
	if err != nil {
		return err
	}
	err2 := service.updateUserInAuthenticationService(regularUserUpdateDto, id)
	if err2 != nil {
		return err2
	}
	return nil
}

func (service *RegularUserService) updateUserInAuthenticationService(regularUserUpdateDto dto.RegularUserUpdateDTO, createdUserId string) error {
	postBody, _ := json.Marshal(map[string]string{
		"_id": 		createdUserId,
		"email":    regularUserUpdateDto.Email,
		"username": regularUserUpdateDto.Username,
		"name":     regularUserUpdateDto.Name,
		"surname":  regularUserUpdateDto.Surname,
	})
	requestUrl := "http://localhost:8080/update"
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) DeleteRegularUser(deleteUserDto dto.DeleteUserDTO) error{
	id, err := primitive.ObjectIDFromHex(deleteUserDto.Id)
	if err != nil {
		return err
	}
	err1 := service.RegularUserRepository.DeleteRegularUser(id)
	if err1 != nil {
		return err1
	}
	return nil
}

func (service *RegularUserService) deleteUserInAuthenticationService(id string) error {
	postBody, _ := json.Marshal(map[string]string{
		"userId":   id,
	})
	requestUrl := "http://localhost:8080/delete"
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}
