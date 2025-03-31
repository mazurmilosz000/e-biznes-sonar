package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
