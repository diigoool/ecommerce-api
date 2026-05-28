package dto

type AddToCartRequest struct {
	ProductID uint `json:"product_id" validate:"required" example:"1"`
	Qty       int  `json:"qty" validate:"gt=0" example:"10"`
}

type RemoveCartItemRequest struct {
	ProductID uint `json:"product_id" validate:"required" example:"2"`
}
