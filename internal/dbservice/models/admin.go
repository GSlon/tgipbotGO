package models

type Admin struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int `gorm:"unique" json:userid`	
	State string `json:"state"` 
}
