package database

import (
	"log"

	"github.com/byte3/galactic.wallet/config"
	"github.com/byte3/galactic.wallet/internal/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Initialize(config *config.Config) {
	var err error
	dsn := "host=localhost user=galactic dbname=GalacticWallet port=9920 sslmode=disable TimeZone=Asia/Colombo"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.Environment == "production" {
		Db.Logger.LogMode(logger.Error)
	}

	if err = Db.AutoMigrate(
		// migrating models
		&models.PaymentModel{},
		&models.WalletModel{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

func SelectWalletByUserId(id uuid.UUID) (models.WalletModel, error) {
	wallet := models.WalletModel{}
	result := Db.Where("User = ?", id).First(&wallet)
	log.Print(wallet)
	return wallet, result.Error
}

func CreateWallet(wallet *models.WalletModel) (int, error) {
	result := Db.Create(wallet)
	return int(result.RowsAffected), result.Error
}

func UpdateWalletBalance(wallet *models.WalletModel) (int, error) {
	result := Db.Save(wallet)
	return int(result.RowsAffected), result.Error
}

func CreatePayment(payment *models.PaymentModel) (int, error) {
	result := Db.Create(payment)
	return int(result.RowsAffected), result.Error
}
