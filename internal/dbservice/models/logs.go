package models

import (
	"gorm.io/gorm"
) 

type ErrorLog struct {
	gorm.Model
	Info string
}