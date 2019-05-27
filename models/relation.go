package models

import (
	"time"
)

//Relation model
type Relation struct {
	Person      string `gorm:"primary_key"`
	Friend      string `gorm:"primary_key"`
	CreatedTime int64
	UpdatedTime int64
}

//Upsert - add relation
func (r *Relation) Upsert() error {
	count := 0
	db.Model(r).Where("person = ? and friend = ?", r.Person, r.Friend).Count(&count)
	if count == 0 {
		r.CreatedTime = time.Now().Unix()
		return db.Model(r).Create(r).Error
	}

	r.UpdatedTime = time.Now().Unix()

	return r.Update()
}

//Update - update relation
func (r *Relation) Update() error {

	return db.Save(r).Error
}

func (Relation) table() string {

	return "relations"
}

//GetRelation - get a relation
func GetRelation(person, friend string) (Relation, error) {
	r := Relation{}
	err := db.Model(r).Where("person = ? and friend = ?", person, friend).Find(&r).Error

	return r, err
}

//GetFriends - get all the friends
func GetFriends(person string) ([]string, error) {
	relations := []Relation{}
	err := db.Model(Relation{}).Where("person = ?", person).Find(&relations).Error
	friends := []string{}
	for _, relation := range relations {

		friends = append(friends, relation.Friend)
	}

	return friends, err
}

//CommonFriends - list common friend //TO-DO - use sql query instead of this
func CommonFriends(friend1, friend2 string) ([]string, error) {
	f1Relations := []Relation{}
	f2Relations := []Relation{}
	common := []string{}
	err := db.Table(Relation{}.table()).Where("person = ?", friend1).Find(&f1Relations).Error
	if err != nil {
		return nil, err
	}
	err = db.Table(Relation{}.table()).Where("person = ?", friend2).Find(&f2Relations).Error
	if err != nil {
		return nil, err
	}
	for _, relation1 := range f1Relations {
		for _, relation2 := range f2Relations {
			if relation1.Friend == relation2.Friend {
				common = append(common, relation1.Friend)
				break
			}
		}
	}

	return common, nil
}

//Delete - delete a relation
func (r *Relation) Delete() error {

	return db.Delete(r).Error
}
