package models

import (
	"os"
	"sync"

	"github.com/ramvasanth/sp/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//db is a mysql database service and is used by all the models
var db *gorm.DB

var mysqlOnce sync.Once

var ErrRecordNotFound = gorm.ErrRecordNotFound

//New - initiate mysql database service
func New() (*gorm.DB, error) {
	dbs, err := gorm.Open("mysql", config.Get().MysqlURL)
	if err != nil {

		return nil, err
	}
	dbs.DB()
	err = dbs.DB().Ping()

	if err != nil {
		return nil, err
	}

	dbs.DB().SetMaxIdleConns(10)
	dbs.DB().SetMaxOpenConns(100)
	if config.Get().WorkerRunMode == "dev" ||
		config.Get().WorkerRunMode == "testing" {
		dbs.LogMode(false)
	}

	return dbs, nil
}

//Set the global database  service
func Set(dbs *gorm.DB) {
	if dbs != nil {
		setDB := func() {
			db = dbs
			if os.Getenv("MIGRATE_DB") == "yes" {
				Migrate()
			}
		}

		mysqlOnce.Do(setDB)
	}
}

//Migrate - migrates the models
func Migrate() {

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Relation{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Subscription{})
}
