package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID        int64
	UserID    int64 `gorm:"default:null"`
	User      User
	CartItems []CartItem
}
