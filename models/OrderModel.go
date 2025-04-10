package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID     int64
	UserID int64 `gorm:"not null"`
	User   User
}
