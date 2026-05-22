package dto

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required,notblank"`
	Price float64 `json:"price" validate:"gt=0"`
	Stock int     `json:"stock" validate:"gte=0"`
}

type UpdateProductRequest struct {
	Name  *string  `json:"name" validate:"omitempty,notblank"`
	Price *float64 `json:"price" validate:"omitempty,gt=0"`
	Stock *int     `json:"stock" validate:"omitempty gte=0"`
}
