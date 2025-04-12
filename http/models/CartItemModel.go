package models

import (
	"github.com/fusionxx23/ecommerce-go/http/database"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ID               int64
	CartID           int64
	Cart             Cart
	ProductVariantID int64 `gorm:"not null"`
	ProductVariant   ProductVariant
	Quantity         int `gorm:"not null"`
}

func InsertCartItem(db *gorm.DB, cartItem *CartItem) error {
	if err := db.Create(cartItem).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCartItemQuantity(newQuantity int, cartId int64, productId int64) error {
	if err := database.DB.Model(&CartItem{}).Where("cart_id = ? AND product_id = ?", cartId, productId).Update("quantity", newQuantity).Error; err != nil {
		return err
	}
	return nil
}
