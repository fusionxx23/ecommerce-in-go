package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ID        int64
	CartID    int64
	Cart      Cart
	ProductID int64 `gorm:"not null"`
	Product   Product
	Quantity  int `gorm:"not null"`
}
