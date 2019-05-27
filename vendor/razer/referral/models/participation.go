package models

import "time"

//Participation model
type Participation struct {
	ID            string `gorm:"primary_key;index" json:"id"`
	CampaignID    string `json:"campaign_id" sql:"size:255" gorm:"index"`
	PromoCode     string `json:"promo_code" sql:"size:50"`
	Referee       string `json:"referee"`
	Referrer      string `json:"referrer"`
	CreatedAt     int64  `json:"created_at"`
	WebhookStatus bool
	WebhookInfo   string
}

//Create - save a new promo
func (p *Participation) Create() error {

	p.CreatedAt = time.Now().Unix()
	return db.Table(p.table()).Create(p).Error
}

//IsParticipated -
func (p *Participation) IsParticipated() (bool, error) {
	count := 0
	err := db.Table(p.table()).Where("campaign_id = ? and referee = ? and referrer = ?", p.CampaignID, p.Referee, p.Referrer).Count(&count).Error
	if count == 0 {
		return false, err
	}

	return true, err
}

//Update -
func (p *Participation) Update() error {

	return db.Table(p.table()).Save(p).Error
}

//GetParticipations - get a Participants list
func GetParticipations(campaignID string, webhookFailure bool) ([]*Participation, error) {
	participations := []*Participation{}
	if webhookFailure {
		err := db.Table(Participation{}.table()).Where("campaign_id = ? and webhook_status = ?", campaignID, true).Find(&participations).Error
		return participations, err
	}

	err := db.Table(Participation{}.table()).Where("campaign_id = ?", campaignID).Find(&participations).Error
	return participations, err
}

//Delete - delete a  campaign participation
func (p *Participation) Delete() error {
	if p.ID == "" {

		return nil
	}

	return db.Delete(p).Error
}

func (Participation) table() string {

	return "participations"
}
