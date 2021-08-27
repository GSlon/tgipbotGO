package models

type Admin struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	UserID int `gorm:"unique"`	
	State string 
}