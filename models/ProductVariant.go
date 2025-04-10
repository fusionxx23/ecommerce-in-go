package models

import "gorm.io/gorm"

type ProductVariant struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	ProductID int64 `gorm:"not null"`
	Name      string
}
