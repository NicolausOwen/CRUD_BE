package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	BuyerID   uint   `json:"buyer_id"`
	Status    string `json:"status"`
}
