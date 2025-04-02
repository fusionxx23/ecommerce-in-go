package models

import "gorm.io/gorm"

type CheckoutSession struct {
	gorm.Model
	ID             int64
	UserID         int64
	User           User
	CartID         int64 `gorm:"not null"`
	Cart           Cart
	DeliveryInfoID int64
	DeliveryInfo   DeliveryInfo
}
