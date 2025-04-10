package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey"`
	Url         string `gorm:"not null"`
	ProductID   int64  `gorm:"not null"`
	Orientation string `gorm:"not null"` // "landscape" or "portrait"
}

func InsertProductImage(db *gorm.DB, productImage *ProductImage) error {
	if err := db.Create(productImage).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProductImage(db *gorm.DB, productImageID int64) error {
	if err := db.Delete(&ProductImage{}, productImageID).Error; err != nil {
		return err
	}
	return nil
}
