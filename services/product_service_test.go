package services_test

import (
	"ecommerce-api/dto"
	"ecommerce-api/mocks"
	"ecommerce-api/models"
	"ecommerce-api/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Testing Create Product
func TestCreateProduct_InvalidName(t *testing.T) {

	mockRepo := new(mocks.ProductRepositoryMock)

	service := services.NewProductService(mockRepo)

	req := dto.CreateProductRequest{
		Name:  "",
		Price: 1000,
		Stock: 10,
	}

	_, err := service.CreateProduct(req)

	assert.Error(t, err)
	assert.Equal(t, "product name is required", err.Error())
}

func TestCreateProduct_InvalidPrice(t *testing.T) {

	mockRepo := new(mocks.ProductRepositoryMock)

	service := services.NewProductService(mockRepo)

	req := dto.CreateProductRequest{
		Name:  "testing",
		Price: -1000,
		Stock: 1,
	}

	_, err := service.CreateProduct(req)

	assert.Error(t, err)
	assert.Equal(t, "price must be greater than 0", err.Error())
}

func TestCreateProduct_InvalidStock(t *testing.T) {

	mockRepo := new(mocks.ProductRepositoryMock)

	service := services.NewProductService(mockRepo)

	req := dto.CreateProductRequest{
		Name:  "testing",
		Price: 10000,
		Stock: -1,
	}

	_, err := service.CreateProduct(req)

	assert.Error(t, err)
	assert.Equal(t, "stock must be greater than or equal 0", err.Error())
}

func TestCreateProduct_RepositoryError(t *testing.T) {

	mockRepo := new(mocks.ProductRepositoryMock)

	service := services.NewProductService(
		mockRepo,
	)

	req := dto.CreateProductRequest{
		Name:  "Laptop",
		Price: 10000,
		Stock: 5,
	}

	mockRepo.
		On(
			"CreateProduct",
			mock.Anything,
		).
		Return(
			models.Product{},
			errors.New("database error"),
		)

	_, err := service.CreateProduct(req)

	assert.Error(t, err)

	assert.Equal(
		t,
		"database error",
		err.Error(),
	)
}

func TestCreateProduct_Success(t *testing.T) {

	mockRepo := new(mocks.ProductRepositoryMock)

	service := services.NewProductService(
		mockRepo,
	)

	req := dto.CreateProductRequest{
		Name:  "Laptop",
		Price: 10000,
		Stock: 5,
	}

	mockRepo.
		On(
			"CreateProduct",
			mock.Anything,
		).
		Return(
			models.Product{
				ID:    1,
				Name:  req.Name,
				Price: req.Price,
				Stock: req.Stock,
			},
			nil,
		)

	result, err := service.CreateProduct(req)

	assert.NoError(t, err)

	assert.Equal(
		t,
		"Laptop",
		result.Name,
	)

	mockRepo.AssertExpectations(t)
}
