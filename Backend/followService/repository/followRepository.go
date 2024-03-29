package repository

import (
	"fmt"
	"followService/dto"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowRepository struct {
	DatabaseSession *neo4j.Session
}

func (repository *FollowRepository) AddFollowing(newFollow dto.NewFollowDTO) bool {
	session := *repository.DatabaseSession

	err1 := repository.addUser(session, newFollow.FollowerId)
	err2 := repository.addUser(session, newFollow.FollowedId)
	if err1 != nil || err2 != nil {
		return false
	}
	result, err3 := session.Run("match (u1:User), (u2:User) where u1.Id = $followerId and "+
		"u2.Id = $followedId merge (u1)-[f:follow{blocked : FALSE, request: $isPrivate}]->(u2) return f",
		map[string]interface{}{"followerId": newFollow.FollowerId, "followedId": newFollow.FollowedId, "isPrivate": newFollow.IsPrivate})
	if err3 != nil {
		return false
	}
	if result.Next() {
		return true
	}
	return false
}

func (repository *FollowRepository) addUser(session neo4j.Session, userId string) error {
	_, err := session.Run("merge (u:User{Id:$followerId}) return u", map[string]interface{}{"followerId": userId})
	if err != nil {
		return err
	}
	return nil
}

func (repository *FollowRepository) SetFollowRequestFalse(loggedUserId string, followerId string) bool {
	session := *repository.DatabaseSession

	fmt.Println("repository ", loggedUserId, " follower", followerId)

	result, err := session.Run("match (u1:User{Id:$followerId})"+
		"-[f:follow {request: TRUE}]->(u2:User{Id:$loggedUserId}) set f.request = false return f;",
		map[string]interface{}{"loggedUserId": loggedUserId, "followerId": followerId})

	fmt.Println(result)
	fmt.Println("error kod prihvatanja ", err)
	if err != nil {
		return false
	}
	if result.Next() {
		fmt.Println("result.Next je true")
		return true
	}
	fmt.Println("result.Next je true")

	return false
}

func (repository *FollowRepository) DeleteUser(id string) bool {
	session := *repository.DatabaseSession
	result, err := session.Run("match (u:User{Id:$id}) detach delete u",
		map[string]interface{}{"id": id})
	if err != nil {
		return false
	}
	if result.Next() {
		return true
	}
	return false
}

func (repository *FollowRepository) BlockUser(loggedUserId string, userId string) bool {
	session := *repository.DatabaseSession
	err1 := repository.addUser(session, loggedUserId)
	err2 := repository.addUser(session, userId)
	if err1 != nil || err2 != nil {
		return false
	}
	result, err := session.Run("match (u1:User), (u2:User) where u1.Id = $loggedUserId and "+
		"u2.Id = $userId merge (u1)-[f:follow{blocked : TRUE, request: FALSE}]->(u2) return f",
		map[string]interface{}{"loggedUserId": loggedUserId, "userId": userId})
	if err != nil {
		return false
	}
	if result.Next() {
		return true
	}
	return false
}

func (repository *FollowRepository) RemoveFollowing(loggedUserId string, followingId string) bool {
	session := *repository.DatabaseSession
	result, err := session.Run(" match (u1:User{Id:$loggedUserId})"+
		"-[f:follow]->(u2:User{Id: $followingId}) detach delete f return u1,u2",
		map[string]interface{}{"loggedUserId": loggedUserId, "followingId": followingId})
	if err != nil {
		return false
	}
	if result.Next() {
		return true
	}
	return false
}

func (repository *FollowRepository) RemoveFollower(loggedUserId string, followerId string) bool {
	session := *repository.DatabaseSession
	result, err := session.Run(" match (u1:User{Id:$followerId})"+
		"-[f:follow]->(u2:User{Id:$loggedUserId}) detach delete f return u1,u2",
		map[string]interface{}{"loggedUserId": loggedUserId, "followerId": followerId})
	if err != nil {
		return false
	}
	if result.Next() {
		return true
	}
	return false
}

func (repository *FollowRepository) FindAllUserFollowersIds(userId string) ([]string, error) {
	session := *repository.DatabaseSession
	var followersIds []string
	result, err := session.Run("match (u)"+
		"-[f:follow{blocked:FALSE,request:FALSE}]->(u1:User{Id:$userId}) return u.Id",
		map[string]interface{}{"userId": userId})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		followersIds = append(followersIds, id)
	}
	if len(followersIds) == 0 {
		return nil, fmt.Errorf("no followers found")
	}
	return followersIds, nil
}

func (repository *FollowRepository) FindAllUserFollowingsIds(userId string) ([]string, error) {
	session := *repository.DatabaseSession
	var followingsIds []string
	result, err := session.Run("match (u1:User{Id:$userId})"+
		"-[f:follow{blocked:FALSE,request:FALSE}]->(u) return u.Id",
		map[string]interface{}{"userId": userId})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		followingsIds = append(followingsIds, id)
	}
	if len(followingsIds) == 0 {
		return nil, fmt.Errorf("no followings found")
	}
	return followingsIds, nil
}

func (repository *FollowRepository) FindAllFeedUsersIds(userId string) ([]string, error) {
	session := *repository.DatabaseSession
	var feedUsersIds []string
	result, err := session.Run("match (u1:User{Id:$userId})"+
		"-[f:follow{blocked:FALSE,request:FALSE}]->(u) return u.Id",
		map[string]interface{}{"userId": userId})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		feedUsersIds = append(feedUsersIds, id)
	}
	if len(feedUsersIds) == 0 {
		return nil, fmt.Errorf("no feed users found")
	}
	return feedUsersIds, nil
}

func (repository *FollowRepository) FindAllUserBlockedUsersIds(userId string) ([]string, error) {
	session := *repository.DatabaseSession
	var blockedUsersIds []string
	result, err := session.Run("match (u1:User{Id:$userId})"+
		"-[f:follow{blocked:TRUE}]->(u:User) return u.Id",
		map[string]interface{}{"userId": userId})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		blockedUsersIds = append(blockedUsersIds, id)
	}
	if len(blockedUsersIds) == 0 {
		return nil, fmt.Errorf("no blocked users found")
	}
	return blockedUsersIds, nil
}

func (repository *FollowRepository) FindAllUserFollowRequestsIds(userId string) ([]string, error) {
	session := *repository.DatabaseSession
	var mutedUsersIds []string
	result, err := session.Run("match (u:User)"+
		"-[f:follow{request:TRUE}]->(u1:User{Id:$userId}) return u.Id",
		map[string]interface{}{"userId": userId})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		mutedUsersIds = append(mutedUsersIds, id)
	}
	if len(mutedUsersIds) == 0 {
		return nil, fmt.Errorf("no follow requests found")
	}
	return mutedUsersIds, nil
}

func (repository *FollowRepository) UserAlreadyFollowed(followerId string, followedId string) error {
	session := *repository.DatabaseSession
	fmt.Println("Da prvo proverimm da li se vec prate Repository")

	result, err := session.Run("match (u1:User{Id:$followerId})-[f:follow]->(u2:User{Id:$followedId}) return f;",
		map[string]interface{}{"followerId": followerId, "followedId": followedId})

	fmt.Println("Result: ", result)
	fmt.Println("result greska:", err)
	if err != nil {
		return nil
	}
	if result.Next() {
		println(result)
		return fmt.Errorf("user is already followed")
	}
	return nil
}
