package models

type Admin struct {
	ID uint `gorm:"primaryKey"`
	UserId uint `gorm:"unique"`	
}