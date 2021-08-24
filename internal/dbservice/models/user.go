package models

import (
	"gorm.io/gorm"
) 

type User struct {
	ID uint `gorm:"primaryKey"`
	TgID uint `gorm:"unique"`	// chat id
	UserName string
	UserId uint `gorm:"unique"`
}

type UserLog struct {
	gorm.Model
	UserID uint
	User User `gorm:"constraint:OnDelete:CASCADE;"`
	IP string `gorm:"not null"`
	Info string 
}
