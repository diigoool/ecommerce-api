package dto

type ProductResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
