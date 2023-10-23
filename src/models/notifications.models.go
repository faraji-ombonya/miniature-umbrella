package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Sender    string         `json:"sender"`
	Receiver  string         `json:"receiver"`
	Body      string         `json:"body"`
	Channel   string         `json:"channel"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdateAt  time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
