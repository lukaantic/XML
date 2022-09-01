package service

import (
	"fmt"
	"net/http"
	"userService/dto"
	"userService/model"
	"userService/repository"
	"userService/tracer"

	"bytes"
	"context"
	"encoding/json"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"string"
)

type RegularUserService struct {
	RegularUserRepository *repository.RegularUserRepository
}

func (service *RegularUserService) Register(ctx context.Context, regularUserRegistrationDto dto.RegularUserRegistrationDTO) error {

	span := tracer.StartSpanFromContext(ctx, "Server Register")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	if service.RegularUserRepository.ExistByUsername(ctx, regularUserRegistrationDto.Username) {
		return fmt.Errorf("username is already taken")
	}

	var regularUser = createRegularUserFromRegularUserRegistrationDTO(&regularUserRegistrationDto)
	userId, err := service.RegularUserRepository.Register(ctx, regularUser)
	if err != nil {
		return err
	}

	err2 := service.registerUserInAuthenticationService(regularUserRegistrationDto, userId)
	if err2 != nil {
		return err2
	}
	fmt.Println("User created")
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
	regularUser.Interests = userUpdateDto.Interests
	regularUser.Experience = userUpdateDto.Experience
	regularUser.Education = userUpdateDto.Education
	regularUser.Skills = userUpdateDto.Skills

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
	regularUser.Skills = regularUserDto.Skills
	regularUser.Experience = regularUserDto.Experience
	regularUser.Interests = regularUserDto.Interests
	regularUser.Education = regularUserDto.Education

	return &regularUser
}

func (service *RegularUserService) registerUserInAuthenticationService(regularUserRegistrationDto dto.RegularUserRegistrationDTO, createdUserId string) error {

	fmt.Println("Ovo cu da saljem kao id:", createdUserId)
	postBody, _ := json.Marshal(map[string]string{
		"id":       createdUserId,
		"password": regularUserRegistrationDto.Password,
		"username": regularUserRegistrationDto.Username,
		"name":     regularUserRegistrationDto.Name,
		"surname":  regularUserRegistrationDto.Surname,
	})

	requestUrl := fmt.Sprintf("http://%s:%s/register", os.Getenv("AUTHENTICATION_SERVICE_DOMAIN"), os.Getenv("AUTHENTICATION_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) UpdatePersonalInformations(ctx context.Context, regularUserUpdateDto dto.RegularUserUpdateDTO) error {
	fmt.Println("Updating regular user")

	span := tracer.StartSpanFromContext(ctx, "Register")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if service.RegularUserRepository.ExistByUsername(ctx, regularUserUpdateDto.Username) {
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
	requestUrl := fmt.Sprintf("http://%s:%s/update", os.Getenv("AUTHENTICATION_SERVICE_DOMAIN"), os.Getenv("AUTHENTICATION_SERVICE_PORT"))

	//requestUrl := "http://localhost:1231/update"
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) DeleteRegularUser(ctx context.Context, deleteUserDto dto.DeleteUserDTO) error {

	span := tracer.StartSpanFromContext(ctx, "Delete")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id, err := primitive.ObjectIDFromHex(deleteUserDto.Id)
	if err != nil {
		return err
	}
	err1 := service.RegularUserRepository.DeleteRegularUser(ctx, id)
	if err1 != nil {
		return err1
	}
	return nil
}

func (service *RegularUserService) deleteUserInAuthenticationService(id string) error {
	postBody, _ := json.Marshal(map[string]string{
		"userId": id,
	})
	requestUrl := fmt.Sprintf("http://%s:%s/delete", os.Getenv("AUTHENTICATION_SERVICE_DOMAIN"), os.Getenv("AUTHENTICATION_SERVICE_PORT"))
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
	regularUserProfileDataDto.Skills = regularUser.Skills
	regularUserProfileDataDto.Interests = regularUser.Interests
	regularUserProfileDataDto.Experience = regularUser.Experience
	regularUserProfileDataDto.Education = regularUser.Education

	return &regularUserProfileDataDto
}

func (service *RegularUserService) GetAllPublicRegularUsers() ([]dto.RegularUserDTO, error) {
	allRegularUsers, err := service.RegularUserRepository.GetAllPublicRegularUsers()
	if err != nil {
		fmt.Println("serv greska")
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
		regularUserIteration.Skills = allRegularUsers[i].Skills
		regularUserIteration.Interests = allRegularUsers[i].Interests
		regularUserIteration.Experience = allRegularUsers[i].Experience
		regularUserIteration.Education = allRegularUsers[i].Education
		regularUser = append(regularUser, regularUserIteration)
	}
	return regularUser
}
func (service *RegularUserService) CreateRegularUserPostDTOByUsername(username string) (*dto.RegularUserPostDTO, error) {
	regularUser, err := service.RegularUserRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	regularUserPostDto := createRegularUserPostDTOFromRegularUser(regularUser)
	return regularUserPostDto, nil
}

func createRegularUserPostDTOFromRegularUser(regularUser *model.RegularUser) *dto.RegularUserPostDTO {
	var regularUserPostDto dto.RegularUserPostDTO
	regularUserPostDto.Id = regularUser.Id.Hex()
	regularUserPostDto.PrivacyType = &regularUser.ProfilePrivacy.PrivacyType

	return &regularUserPostDto
}

func (service *RegularUserService) FindUsersByIds(usersIds []string) (*[]dto.UserFollowDTO, error) {
	var users []model.RegularUser
	for i := 0; i < len(usersIds); i++ {
		id, _ := primitive.ObjectIDFromHex(usersIds[i])
		regularUser, err := service.RegularUserRepository.FindUserById(id)
		if err != nil {
			return nil, err
		}
		users = append(users, *regularUser)
	}

	userFollowDTOs := createUserFollowDTOsFromRegularUsers(users)
	return userFollowDTOs, nil
}

func createUserFollowDTOsFromRegularUsers(regularUsers []model.RegularUser) *[]dto.UserFollowDTO {
	var userFollowDTOs []dto.UserFollowDTO
	for i := 0; i < len(regularUsers); i++ {
		var userFollowDto dto.UserFollowDTO
		userFollowDto.Username = regularUsers[i].Username
		userFollowDto.UserId = regularUsers[i].Id.Hex()
		userFollowDTOs = append(userFollowDTOs, userFollowDto)
	}

	return &userFollowDTOs
}

func (service *RegularUserService) FindUserById(userId primitive.ObjectID) (*model.RegularUser, error){
	fmt.Print("Searching for logged user...")
	regularUser, err := service.RegularUserRepository.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	return regularUser, err
}

func (service *RegularUserService) GetUserSearchResults(searchInput string) ([]model.RegularUser, error){
	searchUser,err := service.RegularUserRepository.GetAllRegularUsers()
	if err != nil {
		return nil, err
	}
	searchrUserModel := CreateUserFromDocuments(searchUser)
	searchUserResults := service.RegularUserRepository.GetUserSearchResults(searchInput, searchrUserModel)

	return searchUserResults, nil
}