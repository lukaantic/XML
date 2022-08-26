package handler

import (
	"encoding/json"
	"net/http"
	"userService/dto"
	"userService/service"

	"github.com/gorilla/mux"
)

type RegularUserHandler struct {
	RegularUserService *service.RegularUserService
}

func (handler *RegularUserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var regularUserRegistrationDto dto.RegularUserRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&regularUserRegistrationDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.Register(regularUserRegistrationDto)
	if err != nil {
		if err.Error() == "username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}

func (handler *RegularUserHandler) UpdatePersonalInformations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userUpdateDto dto.RegularUserUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&userUpdateDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.UpdatePersonalInformations(userUpdateDto)
	if err != nil {
		if err.Error() == "username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *RegularUserHandler) DeleteRegularUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var deleteUserDto dto.DeleteUserDTO
	err1 := json.NewDecoder(r.Body).Decode(&deleteUserDto)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.RegularUserService.DeleteRegularUser(deleteUserDto)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RegularUserHandler) UpdateProfilePrivacy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var profilePrivacyDto dto.ProfilePrivacyDTO
	err := json.NewDecoder(r.Body).Decode(&profilePrivacyDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.UpdateProfilePrivacy(profilePrivacyDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
func (handler *RegularUserHandler) FindRegularUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserPostDto, err := handler.RegularUserService.FindRegularUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(regularUserPostDto)
}

func (handler *RegularUserHandler) GetAllPublicRegularUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	allRegularUsersDto, err := handler.RegularUserService.GetAllPublicRegularUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allRegularUsersDto)
	}
}