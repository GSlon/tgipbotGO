package models

import (
	"gorm.io/gorm"
)

func setup(db *gorm.DB) {
	db.AutoMigrate(&User{}, &UserLog{}, &Admin{}, &ErrorLog{})

}