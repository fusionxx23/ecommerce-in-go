package models

import (
	"github.com/fusionxx23/ecommerce-go/database"
	"gorm.io/gorm"
)

type ProductVariant struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	ProductID int64 `gorm:"not null"`
	Name      string
	Quantity  int16 `gorm:"default:0"`
}

func InsertProductVariant(productVariant *ProductVariant) error {
	if err := database.DB.Create(productVariant).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProductVariant(productVariantID int64) error {
	if err := database.DB.Delete(&ProductVariant{}, productVariantID).Error; err != nil {
		return err
	}
	return nil
}

func SelectProductVariant(productId int64) (ProductVariant, error) {
	productVariant := ProductVariant{}
	if err := database.DB.Model(&ProductVariant{}).Where("id = ?", "").First(&productVariant).Error; err != nil {
		return ProductVariant{}, err
	}

	return productVariant, nil
}
