package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ID        uint
	OrderID   uint `gorm:"not null"`
	Order     Order
	ProductID uint `gorm:"not null"`
	Product   Product
	Quantity  int `gorm:"not null"`
}
