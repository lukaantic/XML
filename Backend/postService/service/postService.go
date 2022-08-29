package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"postService/dto"
	"postService/model"
	"postService/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type PostService struct {
	PostRepository *repository.PostRepository
}

func (service *PostService) CreateNewPost(postUploadDto dto.PostUploadDTO) error {
	fmt.Println("Creating new post")

	post, err := createPostFromPostUploadDTO(&postUploadDto)
	if err != nil {
		fmt.Println("ovde sam pao")
		return err
	}

	_, err1 := service.PostRepository.Create(post)
	if err1 != nil {
		fmt.Println("ili ovde")
		return err1
	}

	return nil
}

func createPostFromPostUploadDTO(postUploadDto *dto.PostUploadDTO) (*model.Post, error) {
	regularUser, err := getRegularUserFromUsername(postUploadDto.Username)
	if err != nil {
		fmt.Println("pokvario sam se tu")
		return nil, err
	}
	var post model.Post
	post.Description = postUploadDto.Description
	post.MediaPaths = postUploadDto.MediaPaths
	post.UploadDate = postUploadDto.UploadDate
	post.RegularUser = *regularUser
	post.RegularUser.Username = postUploadDto.Username
	post.Likes = 0
	post.Dislikes = 0
	return &post, nil
}

func getRegularUserFromUsername(username string) (*model.RegularUser, error) {
	//requestUrl := fmt.Sprintf("http://localhost:%s/by-username/%s", os.Getenv("USER_SERVICE_PORT"), username)
	requestUrl := fmt.Sprintf("http://%s:%s/by-username/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), username)

	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println("mozda i tu ")
		fmt.Println(err)
		return nil, err
	}
	var regularUser model.RegularUser
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&regularUser)

	return &regularUser, nil
}

func (service *PostService) GetAllRegularUserPosts(username string) []model.Post {
	regularUserPostDocuments := service.PostRepository.GetAllByUsername(username)

	regularUserPosts := CreatePostsFromDocuments(regularUserPostDocuments)
	return regularUserPosts
}

func CreatePostsFromDocuments(PostsDocuments []bson.D) []model.Post {
	var publicPosts []model.Post
	for i := 0; i < len(PostsDocuments); i++ {
		var post model.Post
		bsonBytes, _ := bson.Marshal(PostsDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &post)
		publicPosts = append(publicPosts, post)
	}
	return publicPosts
}
