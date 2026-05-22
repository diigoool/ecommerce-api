package dto

type CartProductResponse struct {
	ID    uint    `json:"product_id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CartItemResponse struct {
	ID        uint                `json:"cart_items_id"`
	CartID    uint                `json:"cart_id"`
	ProductID uint                `json:"product_id"`
	Quantity  int                 `json:"quantity"`
	Product   CartProductResponse `json:"product"`
}

type CartResponse struct {
	ID     uint               `json:"cart_id"`
	UserID uint               `json:"user_id"`
	Items  []CartItemResponse `json:"items"`
}
