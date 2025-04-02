package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ID        uint
	CartID    uint
	Cart      Cart
	ProductID uint `gorm:"not null"`
	Product   Product
	Quantity  int `gorm:"not null"`
}
