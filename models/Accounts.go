package models

import (
	"time"

	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model

	Id           int            `json:"ID"`
	Name         string         `json:"name"`
	Create_at    time.Time      `json:"created_at"`
	Transactions []Transactions `json:"transactions" gorm:"foreignKey:AccountID"`
}
