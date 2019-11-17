package models

type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"productName"`
	Seller      string `json:"seller"`
	Price       int    `json:"price"`
}
