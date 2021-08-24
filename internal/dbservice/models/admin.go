package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	ID uint `gorm:"primaryKey"`
	UserName string `gorm:"unique"`
	PasswordHash string `gorm:"not null"`
}