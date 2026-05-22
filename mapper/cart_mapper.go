package mapper

import (
	"ecommerce-api/dto"
	"ecommerce-api/models"
)

func ToCartResponse(cart models.Cart) dto.CartResponse {

	var items []dto.CartItemResponse

	for _, item := range cart.Items {

		items = append(items, dto.CartItemResponse{
			ID:        item.ID,
			CartID:    item.CartID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,

			Product: dto.CartProductResponse{
				ID:    item.Product.ID,
				Name:  item.Product.Name,
				Price: item.Product.Price,
			},
		})
	}

	return dto.CartResponse{
		ID:     cart.ID,
		UserID: cart.UserID,
		Items:  items,
	}
}
