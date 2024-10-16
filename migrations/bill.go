package migrations

import "gorm.io/gorm"

type Bill struct {
	*gorm.Model
	CustomerID uint    `json:"customer_id"`
	Amount     float64 `json:"amount"`
}
