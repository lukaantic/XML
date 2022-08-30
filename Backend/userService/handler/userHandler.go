package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"userService/dto"
	"userService/service"
	"userService/tracer"
	"userService/helper"
	//"userService/poststore"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"context"
)

type RegularUserHandler struct {
	RegularUserService *service.RegularUserService
}

func (handler *RegularUserHandler) Register(w http.ResponseWriter, r *http.Request) {

	span := tracer.StartSpanFromRequest("Register Regular User",opentracing.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.ContextWithSpan(context.Background(), span)


	w.Header().Set("Content-Type", "application/json")
	var regularUserRegistrationDto dto.RegularUserRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&regularUserRegistrationDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.Register(ctx, regularUserRegistrationDto)
	if err != nil {
		if err.Error() == "username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	ctx = tracer.ContextWithSpan(context.Background(), span)
	
	_, err = helper.DecodeBody(ctx, r.Body)
	if err != nil {
		tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
}

func (handler *RegularUserHandler) UpdatePersonalInformations(w http.ResponseWriter, r *http.Request) {

	span := tracer.StartSpanFromRequest("Update Regular User",opentracing.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.ContextWithSpan(context.Background(), span)


	w.Header().Set("Content-Type", "application/json")
	var userUpdateDto dto.RegularUserUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&userUpdateDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.UpdatePersonalInformations(ctx, userUpdateDto)
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

	span := tracer.StartSpanFromRequest("Delete Regular User",opentracing.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.ContextWithSpan(context.Background(), span)

	w.Header().Set("content-type", "application/json")
	var deleteUserDto dto.DeleteUserDTO
	err1 := json.NewDecoder(r.Body).Decode(&deleteUserDto)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.RegularUserService.DeleteRegularUser(ctx, deleteUserDto)
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
		fmt.Println("hen greska")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allRegularUsersDto)
	}
}

func (handler *RegularUserHandler) CreateRegularUserPostDTOByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserPostDto, err := handler.RegularUserService.CreateRegularUserPostDTOByUsername(username)
	if err != nil {
		fmt.Println("stigao sam ovde da padnem")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(regularUserPostDto)
}

func (handler *RegularUserHandler) FindUsersByIds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var usersIds []string
	err := json.NewDecoder(r.Body).Decode(&usersIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userFollowDtos, err := handler.RegularUserService.FindUsersByIds(usersIds)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(userFollowDtos)
}
}