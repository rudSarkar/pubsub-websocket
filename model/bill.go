package model

type Bill struct {
	CustomerID uint    `json:"customer_id"`
	Amount     float64 `json:"amount"`
}
