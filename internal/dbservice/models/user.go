package models

import (
	"gorm.io/gorm"
) 

type User struct {
	ID int `gorm:"primaryKey"`
	UserID int `gorm:"unique"`
	ChatID int64 `gorm:"unique"`	
	State string 	// для ответов на запросы бота
}

type UserLog struct {
	gorm.Model
	UserID int
	User User `gorm:"constraint:OnDelete:CASCADE;"`
	IP string `gorm:"not null"`
	Info string 
}

