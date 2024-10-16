package model

type Order struct {
	ID     uint   `json:"id"`
	Item   string `json:"item"`
	Status string `json:"Status"`
}
