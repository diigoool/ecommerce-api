package mapper

import (
	"ecommerce-api/dto"
	"ecommerce-api/models"
)

func ToProductResponse(product models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}

func ToProductResponses(products []models.Product) []dto.ProductResponse {
	var responses []dto.ProductResponse

	for _, product := range products {

		responses = append(
			responses,
			ToProductResponse(product),
		)
	}

	return responses
}
