package models

import (
	"fmt"

	"github.com/fusionxx23/ecommerce-go/http/database"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

func InsertCategory(category *Category) error {
	var existingCategory Category
	if err := database.DB.Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
		return fmt.Errorf("category with name '%s' already exists", category.Name)
	}
	if err := database.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategory(categoryID int64) error {
	if err := database.DB.Delete(&Category{}, categoryID).Error; err != nil {
		return err
	}
	return nil
}

func SelectCategories() ([]Category, error) {
	categories := []Category{}
	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
