package database

import (
	"log"

	"github.com/byte3/galactic.payment/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Initialize(config *config.Config) {
	var err error
	dsn := "host=localhost user=rxored dbname=bookclub port=9920 sslmode=disable TimeZone=Asia/Colombo"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.Environment == "production" {
		Db.Logger.LogMode(logger.Error)
	}

	if err = Db.AutoMigrate(
	// migrating models

	); err != nil {
		log.Fatal(err.Error())
	}
}
