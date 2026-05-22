package dto

type OrderItemResponse struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type OrderResponse struct {
	ID         uint                `json:"id"`
	TotalPrice float64             `json:"total_price"`
	Items      []OrderItemResponse `json:"items"`
}
