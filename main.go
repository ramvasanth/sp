package main

import (
	"fmt"
	"log"

	"github.com/ramvasanth/sp/config"
	"github.com/ramvasanth/sp/controllers"
	"github.com/ramvasanth/sp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	loadDeps()
	log.Println("dependecies are loaded")
	g := gin.New()
	g.Use(gin.Recovery())
	controllers.LoadRoutes(g)

	listenAddress := fmt.Sprintf("0.0.0.0:%d", 3000)
	log.Println("service is running at ", listenAddress)
	g.Run(listenAddress)
}

func loadDeps() {
	cfg, err := config.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config.Initialize(cfg)

	db, err := models.New()
	if err != nil {
		log.Fatal(err)
	}
	models.Set(db)
}
