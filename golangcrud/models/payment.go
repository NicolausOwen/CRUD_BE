package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID uint    `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}
