package models

import (
	"github.com/fusionxx23/ecommerce-go/http/database"
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey"`
	ImageId     int64  `gorm:"not null;autoIncrement"`
	ProductID   int64  `gorm:"not null"`
	Orientation string // "landscape" or "portrait"
	Optimized   bool   `gorm:"default:false"` // true if the image has been optimized
}

func InsertProductImage(productImage *ProductImage) error {
	tx := database.DB.Create(productImage)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductImage(productImageID int64) error {
	if err := database.DB.Delete(&ProductImage{}, productImageID).Error; err != nil {
		return err
	}
	return nil
}
