package models

import "gorm.io/gorm"

type DeliveryInfo struct {
	gorm.Model
	ID          int64
	UserID      int64
	User        User
	City        string
	Apartment   string
	Address     string
	Email       string
	PhoneNumber string
	Name        string
	State       string
	PostalCode  string
}
