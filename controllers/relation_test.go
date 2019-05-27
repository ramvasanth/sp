package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/ramvasanth/sp/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRelationCreate(t *testing.T) {
	g := gin.Default()
	LoadRoutes(g)
	Convey("Relation CREATE", t, func() {
		Convey("when request fullfils the validation", func() {
			Convey("it should create a relation between two friends", func() {
				f := getANewRelation()
				defer deleteRelation(f)
				reqData, _ := json.Marshal(f)
				req, _ := http.NewRequest("POST", "/relation/create", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, 200)
				r, _ := models.GetRelation(f.Friends[0], f.Friends[1])
				So(r.Person, ShouldEqual, f.Friends[0])
				So(r.Friend, ShouldEqual, f.Friends[1])
			})
		})
		Convey("when request does not fullfil the validation", func() {
			Convey("it should not create a relation between two friends", func() {
				f := getANewRelation()
				f.Friends[0] = ""
				reqData, _ := json.Marshal(f)
				req, _ := http.NewRequest("POST", "/relation/create", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, http.StatusUnprocessableEntity)
			})
		})

		Convey("when duplicate relation is made", func() {
			Convey("it should update the existing relation", func() {
				f := getANewRelation()
				defer deleteRelation(f)
				reqData, _ := json.Marshal(f)
				req, _ := http.NewRequest("POST", "/relation/create", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				So(res.Code, ShouldEqual, 200)
				r, _ := models.GetRelation(f.Friends[0], f.Friends[1])
				So(r.Person, ShouldEqual, f.Friends[0])
				So(r.Friend, ShouldEqual, f.Friends[1])
			})
		})
	})

}
func TestRelationGetFriends(t *testing.T) {
	g := gin.Default()
	LoadRoutes(g)
	Convey("Relation GetFriends", t, func() {
		Convey("when request fullfils the validation", func() {
			Convey("it should return the friends list", func() {
				r1 := getANewRelation()
				addRelation(r1)
				defer deleteRelation(r1)

				r2 := getANewRelation()
				r2.Friends[0] = r1.Friends[0]
				addRelation(r2)
				defer deleteRelation(r2)
				f := Friend{}
				f.Email = r1.Friends[0]
				reqData, _ := json.Marshal(f)
				req, _ := http.NewRequest("POST", "/relation/get", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				resBytes, _ := ioutil.ReadAll(res.Body)
				fList := FriendsList{}
				json.Unmarshal(resBytes, &fList)
				So(res.Code, ShouldEqual, 200)
				So(fList.Success, ShouldBeTrue)
				So(fList.Count, ShouldEqual, 2)
				cFriends := []string{r1.Friends[1], r2.Friends[1]}
				So(fList.Friends[0], ShouldBeIn, cFriends)
				So(fList.Friends[1], ShouldBeIn, cFriends)
			})
		})

		Convey("when request does not fullfil the validation", func() {
			Convey("it should return validation error", func() {
				f := Friend{}
				f.Email = ""
				reqData, _ := json.Marshal(f)
				req, _ := http.NewRequest("POST", "/relation/get", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)

				So(res.Code, ShouldEqual, http.StatusUnprocessableEntity)
			})
		})

		Convey("when friends list is not found", func() {
			Convey("it should return not found", func() {
				f := Friend{}
				f.Email = uuid.New()
				reqData, _ := json.Marshal(f)
				req, _ := http.NewRequest("POST", "/relation/get", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)

				So(res.Code, ShouldEqual, http.StatusNotFound)
			})
		})
	})
}

func TestRelationGetCommonFriends(t *testing.T) {
	g := gin.Default()
	LoadRoutes(g)
	Convey("Relation GetCommonFriends", t, func() {
		Convey("when request fullfils the validation", func() {
			Convey("it should return the common friends list", func() {
				r1 := getANewRelation()
				addRelation(r1)
				defer deleteRelation(r1)

				r2 := getANewRelation()
				addRelation(r2)
				defer deleteRelation(r2)

				// add common relation
				c1 := getANewRelation()
				c1.Friends[0] = r1.Friends[0]
				addRelation(c1)
				defer deleteRelation(c1)

				c2 := getANewRelation()
				c2.Friends[0] = r2.Friends[0]
				c2.Friends[1] = c1.Friends[1]
				addRelation(c2)
				defer deleteRelation(c2)

				f := Friends{}
				f.Friends = append(f.Friends, r1.Friends[0])
				f.Friends = append(f.Friends, r2.Friends[0])

				reqData, _ := json.Marshal(f)
				req, _ := http.NewRequest("POST", "/relation/common", bytes.NewBuffer(reqData))
				res := httptest.NewRecorder()
				g.ServeHTTP(res, req)
				resBytes, _ := ioutil.ReadAll(res.Body)
				fList := FriendsList{}
				json.Unmarshal(resBytes, &fList)
				fmt.Println(string(resBytes))
				So(res.Code, ShouldEqual, 200)
				So(fList.Success, ShouldBeTrue)
				So(fList.Count, ShouldEqual, 1)
				cFriends := []string{c1.Friends[1], c2.Friends[1]}
				So(fList.Friends[0], ShouldBeIn, cFriends)
			})

			Convey("when common friends list is not found", func() {
				Convey("it should return not found", func() {
					f := Friends{}
					f.Friends = append(f.Friends, uuid.New())
					f.Friends = append(f.Friends, uuid.New())

					reqData, _ := json.Marshal(f)
					req, _ := http.NewRequest("POST", "/relation/common", bytes.NewBuffer(reqData))
					res := httptest.NewRecorder()
					g.ServeHTTP(res, req)

					So(res.Code, ShouldEqual, http.StatusNotFound)
				})
			})
		})
	})
}

func getANewRelation() Friends {
	f := Friends{}
	f.Friends = append(f.Friends, uuid.New())
	f.Friends = append(f.Friends, uuid.New())

	return f
}

func deleteRelation(f Friends) {
	r := models.Relation{}
	r.Person = f.Friends[0]
	r.Friend = f.Friends[1]

	r.Delete()
}

func addRelation(f Friends) {
	r := models.Relation{}
	r.Person = f.Friends[0]
	r.Friend = f.Friends[1]

	r.Upsert()
}
