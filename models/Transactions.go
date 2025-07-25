package models

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model

	Id          int       `json:"ID"`
	AccountID   int       `json:"account_id" gorm:"index" validate:"required"`
	Amount      float64   `json:"amount" validate:"required,gt=0"`
	Currency    string    `json:"currency" validate:"required,oneof=COP USD EUR"`
	Type        string    `json:"type" validate:"required,oneof=deposit withdrawal"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	//Account     Accounts  `gorm:"foreignKey:AccountID"`
}
