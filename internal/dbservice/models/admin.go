package models

type Admin struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"unique"`	
	State string 
}