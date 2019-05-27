package controllers

import (
	"errors"
	"net/http"
	"razer/referral/models"

	"github.com/pborman/uuid"

	"github.com/gin-gonic/gin"
)

var errParticipationValidation = errors.New("provide promo_code,campaign_id,referee_uuid values")
var errInvalidPromoCode = errors.New("invalid promo code")
var errInvalidCampaign = errors.New("invalid campaign id")
var errReferrerError = errors.New("referee id is same as referrer id")

//Participation - controller
type Participation struct {
}

//Update - updates the referee participation
func (Participation) Update(g *gin.Context) {
	promoCode := g.Query("promo_code")
	campaignID := g.Query("campaign_id")
	refereeID := g.Query("refree_uuid")

	if promoCode == "" || campaignID == "" || refereeID == "" {
		g.JSON(400, map[string]string{"message": errParticipationValidation.Error()})
		return
	}

	promo := models.Promo{}
	promo.Code = promoCode
	err := promo.Load()
	if err != nil {
		switch err {
		case models.ErrRecordNotFound:
			g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": errInvalidPromoCode.Error()})
			return
		default:
			g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
	}

	if refereeID == promo.UUID {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": errReferrerError.Error()})
		return
	}

	campaign := models.Campaign{}
	campaign.ID = campaignID
	err = campaign.Get(campaignID)
	if err != nil {
		switch err {
		case models.ErrRecordNotFound:
			g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": errEmptyCampignID.Error()})
			return
		default:
			g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
	}

	participation := models.Participation{}
	participation.CampaignID = campaignID
	participation.PromoCode = promoCode
	participation.Referee = refereeID
	participation.Referrer = promo.UUID
	participation.ID = uuid.New()
	err = participation.Create()
	if err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(200, map[string]string{"message": "notified"})
}

//List - list the referess by campaign id
func (Participation) List(g *gin.Context) {
	campaignID := g.Query("campaign_id")
	webhookFailure := false
	if g.Query("webhook_failure") == "true" {
		webhookFailure = true
	}

	if campaignID == "" {
		g.JSON(400, map[string]string{"message": "provide campaign id value"})
		return
	}
	participations, err := models.GetParticipations(campaignID, webhookFailure)
	if err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(200, participations)
}
