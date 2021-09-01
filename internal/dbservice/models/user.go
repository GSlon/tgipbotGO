package models

import (
	"gorm.io/gorm"
) 

type User struct {
	ID int `gorm:"primaryKey" json:"id"`
	UserID int `gorm:"unique" json:"userid"`
	ChatID int64 `gorm:"unique" json:"chatid"`	
	State string `json:"state"` 	// для ответов на запросы бота
}

type UserLog struct {
	gorm.Model
	UserID int `json:"userid"`
	User User `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
	IP string `gorm:"not null" json:"ip"`
	Info string `json:"info"`
}

