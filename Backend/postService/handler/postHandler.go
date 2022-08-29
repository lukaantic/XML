package handler

import (
	"encoding/json"
	"net/http"
	"postService/dto"
	"postService/service"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postUploadDto dto.PostUploadDTO
	err := json.NewDecoder(r.Body).Decode(&postUploadDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.CreateNewPost(postUploadDto)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) GetAllRegularUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserPosts := handler.PostService.GetAllRegularUserPosts(username)
	regularUserPostsJson, err := json.Marshal(regularUserPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(regularUserPostsJson)
	}
}
