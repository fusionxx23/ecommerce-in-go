package models

import (
	"github.com/fusionxx23/ecommerce-go/http/database"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                   int64  `gorm:"primaryKey"`
	Name                 string `gorm:"not null"`
	Price                string `gorm:"not null"`
	Slug                 string `gorm:"not null"`
	Description          string `gorm:"not null"`
	SecondaryThumbnailId int64
	ThumbnailID          int64
	Thumbnail            ProductImage
	SecondaryThumbnail   ProductImage
	ProductImages        []ProductImage
	ProductVariants      []ProductVariant
}

func InsertProduct(product *Product) error {
	if err := database.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(productID int64) error {
	if err := database.DB.Delete(&Product{}, productID).Error; err != nil {
		return err
	}
	// find all ProductImages and delete them
	if err := database.DB.Where("product_id = ?", productID).Delete(&ProductImage{}).Error; err != nil {
		return err
	}
	// find all ProductVariants with ProductID relationship and delete them
	if err := database.DB.Where("product_id = ?", productID).Delete(&ProductVariant{}).Error; err != nil {
		return err
	}
	return nil
}

func SelectProductFromSlug(slug string) (Product, error) {
	product := Product{}
	if err := database.DB.Model(&Product{}).Preload("ProductImages").Preload("ProductVariants").Where("id = ?", slug).First(&product).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}
func SelectAllProducts(limit int, offset int) ([]Product, error) {
	products := []Product{}
	if err := database.DB.Model(&Product{}).
		Preload("ProductImages").
		Preload("ProductVariants").
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
