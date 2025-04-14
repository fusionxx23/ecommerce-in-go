package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey"`
	Url         string `gorm:"not null"`
	ProductID   int64  `gorm:"not null"`
	Orientation string `gorm:"not null"`      // "landscape" or "portrait"
	Optimized   bool   `gorm:"default:false"` // true if the image has been optimized
}

func UpdateProductImage(db *gorm.DB, id string, orientation string) error {
	err := db.Model(&ProductImage{}).Where("id = ?", id).Updates(map[string]any{
		"optimized":   true,
		"orientation": orientation,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
