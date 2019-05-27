package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/ramvasanth/sp/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscriptionCreate(t *testing.T) {
	g := gin.Default()
	LoadRoutes(g)
	Convey("Subscription CREATE", t, func() {
		Convey("when request fullfils the validation", func() {
			Convey("it should create a new subscription", func() {
				s := getANewSubscription()
				defer deleteSubs(s)
				reqData, _ := json.Marshal(s)
				req, _ := http.NewRequest("POST", "/subscription/create", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, 200)
				dS, _ := models.GetSubscription(s.Requestor, s.Target)
				So(dS.Target, ShouldEqual, s.Target)
				So(dS.Requestor, ShouldEqual, s.Requestor)
				So(dS.Active, ShouldBeTrue)
			})
		})

		Convey("when request does not fullfil the validation", func() {
			Convey("it should not create a subscription between two friends", func() {
				s := getANewSubscription()
				s.Requestor = ""
				defer deleteSubs(s)
				reqData, _ := json.Marshal(s)
				req, _ := http.NewRequest("POST", "/subscription/create", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusUnprocessableEntity)
			})
		})
	})

}

func TestSubscriptionBlock(t *testing.T) {
	g := gin.Default()
	LoadRoutes(g)
	Convey("Subscription BLOCK", t, func() {
		Convey("when request fullfils the validation", func() {
			Convey("it should block a subscription", func() {
				s := getANewSubscription()
				defer deleteSubs(s)
				addSubs(s)
				reqData, _ := json.Marshal(s)
				req, _ := http.NewRequest("DELETE", "/subscription/block", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				resBytes, _ := ioutil.ReadAll(res.Body)
				So(string(resBytes), ShouldEqual, `{"success":true}`)
				So(res.Code, ShouldEqual, 200)
				dS, _ := models.GetSubscription(s.Requestor, s.Target)
				So(dS.Target, ShouldEqual, s.Target)
				So(dS.Requestor, ShouldEqual, s.Requestor)
				So(dS.Active, ShouldBeFalse)
			})
		})

		Convey("when request does not fullfil the validation", func() {
			Convey("it should not create a subscription between two friends", func() {
				s := getANewSubscription()
				s.Requestor = ""
				defer deleteSubs(s)
				reqData, _ := json.Marshal(s)
				req, _ := http.NewRequest("DELETE", "/subscription/block", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusUnprocessableEntity)
			})
		})
	})

}

func getANewSubscription() Subscribe {
	s := Subscribe{}
	s.Requestor = uuid.New()
	s.Target = uuid.New()

	return s
}

func addSubs(s Subscribe) {
	dS := models.Subscription{}
	dS.Requestor = s.Requestor
	dS.Target = s.Target

	dS.Upsert()
}

func deleteSubs(s Subscribe) {

	dS := models.Subscription{}
	dS.Requestor = s.Requestor
	dS.Target = s.Target

	dS.Delete()
}
