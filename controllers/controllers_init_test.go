package controllers

import (
	"log"

	"github.com/ramvasanth/sp/config"
	"github.com/ramvasanth/sp/models"
)

func init() {
	cfg, err := config.Load("../.envtest")
	if err != nil {
		log.Fatal(err)
	}
	config.Initialize(cfg)
	db, err := models.New()
	if err != nil {
		log.Fatal(err)
	}
	models.Set(db)
	models.Migrate()
}
