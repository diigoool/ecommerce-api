package services_test

import (
	"ecommerce-api/dto"
	"ecommerce-api/mocks"
	"ecommerce-api/services"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToCart_InvalidQty(t *testing.T) {

	mockRepo := new(mocks.CartRepositoryMock)

	service := services.NewCartService(mockRepo)

	req := dto.AddToCartRequest{
		ProductID: 1,
		Qty:       -1,
	}

	err := service.AddToCart(1, req)

	assert.Error(t, err)
	assert.Equal(t, "invalid quantity", err.Error())

}
