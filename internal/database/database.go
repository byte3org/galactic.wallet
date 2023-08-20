package database

import (
	"log"

	"github.com/byte3/galactic.wallet/config"
	"github.com/byte3/galactic.wallet/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Initialize(config *config.Config) {
	var err error
	dsn := "host=localhost user=galactic dbname=GalacticWallet port=9920 sslmode=disable TimeZone=Asia/Colombo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.Environment == "production" {
		db.Logger.LogMode(logger.Error)
	}

	if err = db.AutoMigrate(
		// migrating models
		&models.PaymentModel{},
		&models.WalletModel{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

func GetDatabase() *gorm.DB {
	return db
}
