package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ID        int64
	OrderID   int64 `gorm:"not null"`
	Order     Order
	ProductID int64 `gorm:"not null"`
	Product   Product
	Quantity  int `gorm:"not null"`
}
