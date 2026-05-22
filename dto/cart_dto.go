package dto

type AddToCartRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Qty       int  `json:"qty" validate:"gt=0"`
}

type RemoveCartItemRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
}
