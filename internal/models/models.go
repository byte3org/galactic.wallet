package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type PaymentModel struct {
	BaseModel
	Amount    int       `json:"amount"`
	IsSuccess bool      `json:"is_success"`
	WalletID  uuid.UUID `json:wallet`
	Wallet    WalletModel
}

type WalletModel struct {
	BaseModel
	User    uuid.UUID `json:"id" gorm:"uuid"`
	Balance int       `json:"balance"`
}
