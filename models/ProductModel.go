package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID            int64 `gorm:"primaryKey"`
	Name          string
	Price         string
	Description   string
	ThumbnailID   int64
	Thumbnail     ProductImage
	ProductImages []ProductImage
}
