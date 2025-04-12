package models

import (
	"github.com/fusionxx23/ecommerce-go/http/database"
	"gorm.io/gorm"
)

type ProductVariant struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	ProductID int64 `gorm:"not null" json:"product_id"`
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

func SelectProductVariant(id int64) (ProductVariant, error) {
	productVariant := ProductVariant{}
	if err := database.DB.Model(&ProductVariant{}).Where("id = ?", id).First(&productVariant).Error; err != nil {
		return ProductVariant{}, err
	}

	return productVariant, nil
}

func SelectProductVariants(productId int64) ([]ProductVariant, error) {
	productVariants := []ProductVariant{}

	if err := database.DB.Model(&ProductVariant{}).Where("product_id = ?", productId).Select("id", "product_id", "name", "quantity").Find(&productVariants).Error; err != nil {
		return []ProductVariant{}, err
	}

	return productVariants, nil
}
