package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID                    int64  `gorm:"primaryKey"`
	Email                 string `gorm:"unique;not null"`
	DefaultDeliveryInfoId int64
	RefreshToken          string // for JWT refresh token
}
