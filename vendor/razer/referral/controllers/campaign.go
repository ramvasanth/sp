package controllers

import (
	"errors"
	"net/http"
	"razer/referral/models"
	"strconv"

	"github.com/pborman/uuid"

	"github.com/gin-gonic/gin"
)

var errEmptyCampignID = errors.New("campaign id is empty")

//Campaign - controller
type Campaign struct{}

//Create -
func (Campaign) Create(g *gin.Context) {
	campaign := &models.Campaign{}
	if err := g.BindJSON(campaign); err != nil {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}

	campaign.ID = uuid.New()
	if err := campaign.Create(); err != nil {
		switch err {
		case models.ErrDuplicateCampaign:
			g.JSON(409, map[string]string{"message": err.Error()})
			return
		default:
			g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
	}

	g.JSON(200, campaign)
}

//Update -
func (Campaign) Update(g *gin.Context) {
	campaign := &models.Campaign{}
	if err := g.BindJSON(campaign); err != nil {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}
	if campaign.ID == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": errEmptyCampignID.Error()})
		return
	}

	if err := campaign.Update(); err != nil {
		switch err {
		case models.ErrDuplicateCampaign:
			g.JSON(409, map[string]string{"message": err.Error()})
			return
		default:
			g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
	}

	g.JSON(200, campaign)
}

//Delete -
func (Campaign) Delete(g *gin.Context) {
	id := g.Param("id")

	if id == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "campaign id is empty"})
		return
	}

	campaign := &models.Campaign{}
	if err := campaign.Get(id); err != nil {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}

	if err := campaign.Delete(); err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(200, map[string]string{"message": "campaign is deleted successfully"})
}

//Delete -
func (Campaign) Get(g *gin.Context) {
	id := g.Param("id")

	if id == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "campaign id is empty"})
		return
	}

	campaign := &models.Campaign{}
	if err := campaign.Get(id); err != nil {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(200, campaign)
}

//List -
func (Campaign) List(g *gin.Context) {
	startDate := int64(0)
	endDate := int64(0)
	name := g.Query("name")

	if pEDate, err := strconv.ParseInt(g.Query("end_date"), 10, 64); err == nil {
		endDate = pEDate
	}

	if pSDate, err := strconv.ParseInt(g.Query("start_date"), 10, 64); err == nil {
		startDate = pSDate
	}

	campaigns, err := models.GetCampaigns(startDate, endDate, name)
	if err != nil {
		switch err {
		case models.ErrCampaignSearch:
			g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
			return
		default:
			g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
	}

	g.JSON(200, campaigns)
}
