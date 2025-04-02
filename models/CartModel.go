package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID     int64
	UserID uint
	User   User
}
