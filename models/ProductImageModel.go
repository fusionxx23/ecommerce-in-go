package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ID        int64  `gorm:"primaryKey"`
	Url       string `gorm:"not null"`
	ProductID int64  `gorm:"not null"`
}
