package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID              int64 `gorm:"primaryKey"`
	Name            string
	Price           string
	Description     string
	ThumbnailID     int64
	Thumbnail       ProductImage
	ProductImages   []ProductImage
	ProductVariants []ProductVariant
}

func InsertProduct(db *gorm.DB, product *Product) error {
	if err := db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(db *gorm.DB, productID int64) error {
	if err := db.Delete(&Product{}, productID).Error; err != nil {
		return err
	}
	// find all ProductImages and delete them
	if err := db.Where("product_id = ?", productID).Delete(&ProductImage{}).Error; err != nil {
		return err
	}
	// find all ProductVariants with ProductID relationship and delete them
	if err := db.Where("product_id = ?", productID).Delete(&ProductVariant{}).Error; err != nil {
		return err
	}
	return nil
}
