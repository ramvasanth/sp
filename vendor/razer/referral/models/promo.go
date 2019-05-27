package models

import "time"

//Promo model
type Promo struct {
	Code      string `json:"code" sql:"size:50" gorm:"primary_key;index"`
	UUID      string `json:"uuid" form:"uuid" binding:"required" sql:"size:255" gorm:"index"`
	Length    int    `json:"length" form:"length" binding:"required" sql:"-"`
	CreatedAt int64  `json:"created_at"`
}

//Create - save a new promo
func (p *Promo) Create() error {

	p.CreatedAt = time.Now().Unix()
	return db.Table(p.table()).Create(p).Error
}

//Get - get a promo
func (p *Promo) Get(code, uuid string) error {

	return db.Table(p.table()).Where("code = ? and uuid = ?", code, uuid).First(p).Error
}

//Load - get a promo
func (p *Promo) Load() error {

	return db.Table(p.table()).Where("code = ? ", p.Code).First(p).Error
}

//Delete - delete a  promo code
func (p *Promo) Delete() error {
	if p.Code == "" {

		return nil
	}

	return db.Delete(p).Error
}

func (Promo) table() string {

	return "promos"
}
