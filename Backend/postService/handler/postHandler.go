package handler

import (
	"encoding/json"
	"net/http"
	"postService/dto"
	"postService/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (handler *PostHandler) GetAllPublicPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	publicPosts := handler.PostService.GetAllPublicPosts()
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(publicPostsJson)
	}
}

func (handler *PostHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var commentDTO dto.CommentDTO
	err := json.NewDecoder(r.Body).Decode(&commentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.CommentPost(commentDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postLikeDTO dto.PostLikeDTO
	err := json.NewDecoder(r.Body).Decode(&postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.LikePost(postLikeDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) DislikePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postLikeDTO dto.PostLikeDTO
	err := json.NewDecoder(r.Body).Decode(&postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.DislikePost(postLikeDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	postId, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.PostService.DeletePost(postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
