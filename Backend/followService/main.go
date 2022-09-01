package main

import (
	"fmt"
	"followService/handler"
	"followService/repository"
	"followService/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/cors"
	//"os"
)

func initFollowRepository(databaseSession *neo4j.Session) *repository.FollowRepository {
	return &repository.FollowRepository{DatabaseSession: databaseSession}
}

func initFollowService(repository *repository.FollowRepository) *service.FollowService {
	return &service.FollowService{FollowRepository: repository}
}

func initFollowHandler(service *service.FollowService) *handler.FollowHandler {
	return &handler.FollowHandler{FollowService: service}
}

func handleFunc(handler *handler.FollowHandler) {
	ruter := mux.NewRouter().StrictSlash(true)

	ruter.HandleFunc("/follow", handler.FollowUser).Methods("POST")
	ruter.HandleFunc("/accept-follow/{loggedUserId}/{followerId}", handler.AcceptFollowRequest).Methods("PUT")
	ruter.HandleFunc("/block-user/{loggedUserId}/{userId}", handler.BlockUser).Methods("POST")
	ruter.HandleFunc("/unblock-user/{loggedUserId}/{userId}", handler.UnblockUser).Methods("POST")
	ruter.HandleFunc("/remove-following/{loggedUserId}/{followingId}", handler.RemoveFollowing).Methods("POST")
	ruter.HandleFunc("/remove-follower/{loggedUserId}/{followerId}", handler.RemoveFollower).Methods("POST")
	ruter.HandleFunc("/followers/{loggedUserId}", handler.FindAllUserFollowers).Methods("GET")
	ruter.HandleFunc("/followings/{loggedUserId}", handler.FindAllUserFollowings).Methods("GET")
	ruter.HandleFunc("/blocked-users/{loggedUserId}", handler.FindAllUserBlockedUsers).Methods("GET")
	ruter.HandleFunc("/follow-requests/{loggedUserId}", handler.FindAllUserFollowRequests).Methods("GET")
	ruter.HandleFunc("/users-for-feed/{loggedUserId}", handler.FindAllFeedUsers).Methods("POST")
	ruter.HandleFunc("/delete-user", handler.DeleteUser).Methods("POST")

	c := SetupCors()

	http.Handle("/", c.Handler(ruter))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(ruter))
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)

	}
}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins, for now
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}

func initDatabase() (neo4j.Session, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	if driver, err = neo4j.NewDriver("neo4j://neo4j:7687", neo4j.BasicAuth("neo4j", "12345", "")); err != nil {
		//if driver, err = neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345", "")); err != nil {
		return nil, err
	}

	if session = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}); err != nil {
		return nil, err
	}

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("match (u) return u;", map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], err
		}
		return nil, result.Err()
	})

	if err != nil {
		return nil, err
	}
	return session, nil
}

func main() {
	session, err := initDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	authenticationRepository := initFollowRepository(&session)
	authenticationService := initFollowService(authenticationRepository)
	authenticationHandler := initFollowHandler(authenticationService)

	handleFunc(authenticationHandler)
}
