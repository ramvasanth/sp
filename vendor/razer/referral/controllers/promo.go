package controllers

import (
	"fmt"
	"net/http"
	"razer/referral/models"

	"github.com/dchest/uniuri"

	"github.com/gin-gonic/gin"
)

const (
	promoMaxLength = 25
	promoMinLength = 10
)

var errPromoLength = fmt.Errorf("acceptable promo length is between %d and %d", promoMinLength, promoMaxLength)

//Promo - controller
type Promo struct {
}

//Create - create a n ew promo
func (Promo) Create(g *gin.Context) {
	promo := &models.Promo{}

	if err := g.BindQuery(promo); err != nil {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}
	if promo.Length > promoMaxLength || promo.Length < promoMinLength {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": errPromoLength.Error()})
		return
	}

	promo.Code = uniuri.NewLen(promo.Length)

	if err := promo.Create(); err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(200, promo)
}

//Delete - delete the promo code for the user
func (Promo) Delete(g *gin.Context) {
	code := g.Query("code")
	uuid := g.Query("uuid")

	if code == "" || uuid == "" {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "either promo code or uuid is empty"})
		return
	}

	promo := &models.Promo{}
	if err := promo.Get(code, uuid); err != nil {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}

	if err := promo.Delete(); err != nil {
		g.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}

	g.JSON(200, map[string]string{"message": "promo code is deleted successfully"})
}
