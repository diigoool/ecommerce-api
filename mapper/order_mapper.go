package mapper

import (
	"ecommerce-api/dto"
	"ecommerce-api/models"
)

func ToOrderResponse(order models.Order) dto.OrderResponse {
	var items []dto.OrderItemResponse

	for _, item := range order.Items {
		items = append(items, dto.OrderItemResponse{
			ProductID:   item.ProductID,
			ProductName: item.Product.Name,
			Quantity:    item.Quantity,
			Price:       item.Price,
		})
	}

	return dto.OrderResponse{
		ID:         order.ID,
		TotalPrice: order.TotalPrice,
		Items:      items,
	}

}

func ToOrderResponses(
	orders []models.Order,
) []dto.OrderResponse {

	var responses []dto.OrderResponse

	for _, order := range orders {

		responses = append(
			responses,
			ToOrderResponse(order),
		)
	}

	return responses
}
