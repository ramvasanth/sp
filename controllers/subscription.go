package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramvasanth/sp/models"
)

type Subscription struct {
}

func (Subscription) Create(g *gin.Context) {
	subscribe := Subscribe{}
	err := g.BindJSON(&subscribe)
	if err != nil || subscribe.Target == "" || subscribe.Requestor == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "check the requestor and target"})
		return
	}

	subscribtion := &models.Subscription{Requestor: subscribe.Requestor, Target: subscribe.Target, Active: true}
	err = subscribtion.Upsert()

	if err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(http.StatusOK, map[string]bool{"success": true})
}

func (Subscription) Block(g *gin.Context) {
	subscribe := Subscribe{}
	err := g.BindJSON(&subscribe)
	if err != nil || subscribe.Target == "" || subscribe.Requestor == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "check the requestor and target"})
		return
	}

	subscribtion := &models.Subscription{Requestor: subscribe.Requestor, Target: subscribe.Target, Active: false}
	err = subscribtion.Update()
	if err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(http.StatusOK, map[string]bool{"success": true})
}
