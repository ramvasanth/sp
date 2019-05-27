package models

import (
	"errors"
	"time"
)

//Campaign model
type Campaign struct {
	ID          string `json:"id" sql:"size:255" gorm:"primary_key;index"`
	Name        string `json:"name" sql:"size:255" binding:"required"`
	ServiceName string `json:"service_name" sql:"size:255" binding:"required"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	StartDate   int64  `json:"start_date" binding:"required"`
	EndDate     int64  `json:"end_date" binding:"required"`
}

var ErrDuplicateCampaign = errors.New("duplicate campaign")
var ErrCampaignSearch = errors.New("indicate atleast one search criteria")

//Create - save a new campaign
func (c *Campaign) Create() error {
	c.CreatedAt = time.Now().Unix()

	if c.campaignExist() {
		return ErrDuplicateCampaign
	}

	return db.Table(c.table()).Create(c).Error
}

//Delete - delete a  campaign
func (c *Campaign) Delete() error {
	if c.ID == "" {
		return nil
	}
	return db.Table(c.table()).Delete(c).Error
}

//Get - get a campaign
func (c *Campaign) Get(id string) error {

	return db.Table(c.table()).Where("id = ?", id).First(c).Error
}

//Update - update a campaign
func (c *Campaign) Update() error {
	count := 0
	db.Table(c.table()).Where("name = ? and id <> ?", c.Name, c.ID).Count(&count)
	if count > 0 {
		return ErrDuplicateCampaign
	}

	c.UpdatedAt = time.Now().Unix()

	return db.Table(c.table()).Save(c).Error
}

func (c *Campaign) campaignExist() bool {
	count := 0
	db.Table(c.table()).Where("name = ?", c.Name).Count(&count)
	if count > 0 {
		return true
	}

	return false
}

func (Campaign) table() string {

	return "campaigns"
}

//GetCampaigns - get all campaigns based on search conditions
func GetCampaigns(startDate, endDate int64, name string) ([]*Campaign, error) {
	campaigns := []*Campaign{}
	if startDate > 0 && endDate > 0 {
		if name != "" {
			err := db.Table(Campaign{}.table()).Where("start_date >= ? and  end_date <= ? and name LIKE ?", startDate, endDate, "%"+name+"%").Find(&campaigns).Error
			return campaigns, err
		}

		err := db.Table(Campaign{}.table()).Where("start_date >= ? and  end_date <= ?", startDate, endDate).Find(&campaigns).Error
		return campaigns, err
	}

	if startDate > 0 {
		if name != "" {
			err := db.Table(Campaign{}.table()).Where("start_date >= ? and name LIKE ?", startDate, "%"+name+"%").Find(&campaigns).Error
			return campaigns, err
		}
		err := db.Table(Campaign{}.table()).Where("start_date >= ? ", startDate).Find(&campaigns).Error
		return campaigns, err
	}

	if endDate > 0 {
		if name != "" {
			err := db.Table(Campaign{}.table()).Where("end_date <= ? and name LIKE ?", endDate, "%"+name+"%").Find(&campaigns).Error
			return campaigns, err
		}

		err := db.Table(Campaign{}.table()).Where("end_date <= ? ", endDate).Find(&campaigns).Error
		return campaigns, err
	}

	if name != "" {
		err := db.Table(Campaign{}.table()).Where("name LIKE ?", "%"+name+"%").Find(&campaigns).Error
		return campaigns, err
	}

	return campaigns, ErrCampaignSearch
}
