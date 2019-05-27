package controllers

import (
	"net/http"

	"github.com/ramvasanth/sp/models"

	"github.com/gin-gonic/gin"
)

//Relation - controller
type Relation struct {
}

//Create - new relation
func (Relation) Create(g *gin.Context) {
	friends := Friends{}
	err := g.BindJSON(&friends)
	if err != nil || len(friends.Friends) < 2 || friends.Friends[0] == "" || friends.Friends[1] == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "check the friends list , two friends are required."})
		return
	}
	relation := models.Relation{Person: friends.Friends[0], Friend: friends.Friends[1]}
	err = relation.Upsert()

	if err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(http.StatusOK, map[string]bool{"success": true})
}

//GetFriend - get existing relation
func (Relation) GetFriend(g *gin.Context) {
	friend := Friend{}
	err := g.BindJSON(&friend)
	if err != nil || friend.Email == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "check the email"})
		return
	}

	relations, err := models.GetFriends(friend.Email)
	if err != nil {
		switch err {
		case models.ErrRecordNotFound:
			g.JSON(http.StatusNotFound, map[string]string{"message": "oops , no friends"})
			return
		default:
			g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
	}
	if len(relations) == 0 {
		g.JSON(http.StatusNotFound, map[string]string{"message": "oops , no friends"})
		return
	}
	friends := FriendsList{}
	friends.Success = true
	friends.Count = len(relations)
	friends.Friends = relations

	g.JSON(http.StatusOK, friends)
}

//GetCommonFriends - get common relations
func (Relation) GetCommonFriends(g *gin.Context) {
	friends := Friends{}
	err := g.BindJSON(&friends)
	if err != nil || len(friends.Friends) < 2 {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "check the friends list , two friends are required."})
		return
	}

	relations, err := models.CommonFriends(friends.Friends[0], friends.Friends[1])
	if err != nil {
		switch err {
		case models.ErrRecordNotFound:
			g.JSON(http.StatusNotFound, map[string]string{"message": "oops , no friends"})
			return
		default:
			g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
	}
	if len(relations) == 0 {
		g.JSON(http.StatusNotFound, map[string]string{"message": "oops , no friends"})
		return
	}

	friendList := FriendsList{}
	friendList.Success = true
	friendList.Count = len(relations)
	friendList.Friends = relations

	g.JSON(http.StatusOK, friendList)
}
