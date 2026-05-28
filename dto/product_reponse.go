package dto

type ProductResponse struct {
	ID    uint    `json:"id" example:"1"`
	Name  string  `json:"product_name" example:"Laptop"`
	Price float64 `json:"price" example:"1000000"`
	Stock int     `json:"stock" example:"5"`
}
