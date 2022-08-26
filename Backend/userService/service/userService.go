package service

import (
	"fmt"
	"net/http"
	"userService/dto"
	"userService/model"
	"userService/repository"

	"go.mongodb.org/mongo-driver/bson"

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

func createRegularUserFromRegularUserUpdateDTO(userUpdateDto *dto.RegularUserUpdateDTO) *model.RegularUser {
	id, _ := primitive.ObjectIDFromHex(userUpdateDto.Id)
	var regularUser model.RegularUser
	regularUser.Id = id
	regularUser.Name = userUpdateDto.Name
	regularUser.Surname = userUpdateDto.Surname
	regularUser.Username = userUpdateDto.Username
	regularUser.Email = userUpdateDto.Email
	regularUser.PhoneNumber = userUpdateDto.PhoneNumber
	regularUser.Gender = userUpdateDto.Gender
	regularUser.BirthDate = userUpdateDto.BirthDate
	regularUser.Biography = userUpdateDto.Biography

	return &regularUser
}

func createRegularUserFromRegularUserRegistrationDTO(regularUserDto *dto.RegularUserRegistrationDTO) *model.RegularUser {
	profilePrivacy := model.ProfilePrivacy{
		PrivacyType:        model.PrivacyType(0),
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
	requestUrl := "http://localhost:1231/register"
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
		"_id":      createdUserId,
		"email":    regularUserUpdateDto.Email,
		"username": regularUserUpdateDto.Username,
		"name":     regularUserUpdateDto.Name,
		"surname":  regularUserUpdateDto.Surname,
	})
	requestUrl := "http://localhost:1231/update"
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) DeleteRegularUser(deleteUserDto dto.DeleteUserDTO) error {
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
		"userId": id,
	})
	requestUrl := "http://localhost:1231/delete"
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}
func (service *RegularUserService) UpdateProfilePrivacy(profilePrivacyDto dto.ProfilePrivacyDTO) error {
	fmt.Println("Updating regular user")

	var regularUser = createRegularUserFromProfilePrivacyDTO(&profilePrivacyDto)
	err := service.RegularUserRepository.UpdateProfilePrivacy(regularUser)
	if err != nil {
		return err
	}
	return nil
}

func createRegularUserFromProfilePrivacyDTO(profilePrivacyDto *dto.ProfilePrivacyDTO) *model.RegularUser {
	id, _ := primitive.ObjectIDFromHex(profilePrivacyDto.Id)
	var regularUser model.RegularUser
	regularUser.Id = id
	regularUser.ProfilePrivacy.PrivacyType = profilePrivacyDto.PrivacyType
	regularUser.ProfilePrivacy.AllMessageRequests = profilePrivacyDto.AllMessagesRequests
	return &regularUser
}

func (service *RegularUserService) FindRegularUserByUsername(username string) (*dto.RegularUserProfileDataDTO, error) {
	regularUser, err := service.RegularUserRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	regularUserPostDto := createRegularUserProfileDataDto(regularUser)
	return regularUserPostDto, nil
}

func createRegularUserProfileDataDto(regularUser *model.RegularUser) *dto.RegularUserProfileDataDTO {
	var regularUserProfileDataDto dto.RegularUserProfileDataDTO

	regularUserProfileDataDto.Id = regularUser.Id
	regularUserProfileDataDto.Name = regularUser.Name
	regularUserProfileDataDto.Surname = regularUser.Surname
	regularUserProfileDataDto.Username = regularUser.Username
	regularUserProfileDataDto.Biography = regularUser.Biography
	regularUserProfileDataDto.ProfilePrivacy = regularUser.ProfilePrivacy

	return &regularUserProfileDataDto
}

func (service *RegularUserService) GetAllPublicRegularUsers() ([]dto.RegularUserDTO, error) {
	allRegularUsers, err := service.RegularUserRepository.GetAllPublicRegularUsers()
	if err != nil {
		return nil, err
	}

	allRegularUsersModel := CreateUserFromDocuments(allRegularUsers)

	allRegularUsersDto := createRegularUserDtoFromRegularUser(allRegularUsersModel)
	return allRegularUsersDto, nil
}

func CreateUserFromDocuments(UserDocuments []bson.D) []model.RegularUser {
	var users []model.RegularUser
	for i := 0; i < len(UserDocuments); i++ {
		var user model.RegularUser
		bsonBytes, _ := bson.Marshal(UserDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &user)
		users = append(users, user)
	}
	return users
}

func createRegularUserDtoFromRegularUser(allRegularUsers []model.RegularUser) []dto.RegularUserDTO {

	var regularUser []dto.RegularUserDTO
	for i := 0; i < len(allRegularUsers); i++ {
		var regularUserIteration dto.RegularUserDTO
		regularUserIteration.Id = allRegularUsers[i].Id
		regularUserIteration.Username = allRegularUsers[i].Username
		regularUserIteration.Name = allRegularUsers[i].Name
		regularUserIteration.Surname = allRegularUsers[i].Surname
		regularUser = append(regularUser, regularUserIteration)
	}
	return regularUser
}
