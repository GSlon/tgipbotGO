package models

import (
	"gorm.io/gorm"
) 

type User struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"unique"`
	ChatID uint `gorm:"unique"`	
	State string `gorm: "default:default_state"`// для ответов на запросы бота
}

type UserLog struct {
	gorm.Model
	UserID uint
	User User `gorm:"constraint:OnDelete:CASCADE;"`
	IP string `gorm:"not null"`
	Info string 
}

