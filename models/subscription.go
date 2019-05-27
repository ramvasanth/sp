package models

import "time"

//Subscription - model
type Subscription struct {
	Requestor   string `gorm:"primary_key"`
	Target      string `gorm:"primary_key"`
	Active      bool
	CreatedTime int64
	UpdatedTime int64
}

func (Subscription) table() string {

	return "subscriptions"
}

//Upsert - upsert a subscription
func (s *Subscription) Upsert() error {
	count := 0
	db.Model(s).Where("requestor = ? and target = ?", s.Requestor, s.Target).Count(&count)
	if count == 0 {
		s.CreatedTime = time.Now().Unix()
		return db.Model(s).Create(s).Error
	}

	s.UpdatedTime = time.Now().Unix()

	return s.Update()
}

//Update - update subscription
func (s *Subscription) Update() error {

	return db.Save(s).Error
}

//Delete - delete a subscription
func (s *Subscription) Delete() error {

	return db.Delete(s).Error
}

//GetSubscription - get a subscription
func GetSubscription(request, target string) (Subscription, error) {
	s := Subscription{}
	err := db.Model(s).Where("requestor = ? and target = ?", request, target).Find(&s).Error

	return s, err
}
