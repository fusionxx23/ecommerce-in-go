package models

import (
	"github.com/fusionxx23/ecommerce-go/http/database"
	"gorm.io/gorm"
)

type ProductCategoryRef struct {
	gorm.Model
	ID         int64 `gorm:"primaryKey"`
	ProductId  int64
	CategoryId int64
	Product    Product  `gorm:"foreignKey:ProductId"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

func InsertProductCategoryRef(productCategoryRef *ProductCategoryRef) error {
	if err := database.DB.Create(productCategoryRef).Error; err != nil {
		return err
	}
	return nil
}
