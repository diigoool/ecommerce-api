package mocks

import (
	"ecommerce-api/models"

	"github.com/stretchr/testify/mock"
)

type CartRepositoryMock struct {
	mock.Mock
}

func (m *CartRepositoryMock) GetOrCreateCart(userID uint) (models.Cart, error) {

	args := m.Called(userID)

	return args.Get(0).(models.Cart), args.Error(1)

}

func (m *CartRepositoryMock) GetProductByID(productID uint) (models.Product, error) {

	args := m.Called(productID)

	return args.Get(0).(models.Product), args.Error(1)
}

func (m *CartRepositoryMock) GetCartItem(cartID, productID uint) (models.CartItem, error) {

	args := m.Called(cartID, productID)

	return args.Get(0).(models.CartItem), args.Error(1)
}

func (m *CartRepositoryMock) AddToCart(cartID, productID uint, qty int) error {

	args := m.Called(cartID, productID, qty)

	return args.Error(0)

}

func (m *CartRepositoryMock) GetCart(userID uint) (models.Cart, error) {
	args := m.Called(userID)

	return args.Get(0).(models.Cart), args.Error(1)
}

func (m *CartRepositoryMock) RemoveItem(cartID, productID uint) error {
	args := m.Called(cartID, productID)

	return args.Error(0)

}
