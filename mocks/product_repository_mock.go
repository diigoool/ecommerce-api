package mocks

import (
	"ecommerce-api/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) GetAllProduct(page, limit int, search string) ([]models.Product, error) {
	args := m.Called(page, limit, search)

	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) FindById(id uint) (models.Product, error) {

	args := m.Called(id)

	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) CreateProduct(product models.Product) (models.Product, error) {

	args := m.Called(product)

	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) DeleteProduct(id uint) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *ProductRepositoryMock) UpdateProduct(id uint, product models.Product) (models.Product, error) {

	args := m.Called(id, product)

	return args.Get(0).(models.Product), args.Error(1)
}
