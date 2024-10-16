package migrations

import "gorm.io/gorm"

type Order struct {
	*gorm.Model
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}
