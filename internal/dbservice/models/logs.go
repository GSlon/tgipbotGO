package models

import (
	"gorm.io/gorm"
) 

// таблица для логгирования ошибок
type ErrorLog struct {
	gorm.Model
	Info string `json:"info"`
}
