package services

import (
	"ecommerce-api/dto"
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"ecommerce-api/utils"
	"strings"
)

type ProductService struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetAllProduct(page, limit int, search string) ([]models.Product, error) {

	if page < 1 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	return s.Repo.GetAllProduct(page, limit, search)
}

func (s *ProductService) GetProductById(id uint) (models.Product, error) {
	product, err := s.Repo.FindById(id)
	if err != nil {
		return product, utils.NewNotFoundError("product not found")
	}
	return product, nil
}

func (s *ProductService) CreateProduct(req dto.CreateProductRequest) (models.Product, error) {

	if strings.TrimSpace(req.Name) == "" {
		return models.Product{}, utils.NewBadRequestError("product name is required")
	}

	if req.Price <= 0 {
		return models.Product{}, utils.NewBadRequestError("price must be greater than 0")
	}

	if req.Stock < 0 {
		return models.Product{},
			utils.NewBadRequestError("stock must be greater than or equal 0")
	}

	product := models.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	return s.Repo.CreateProduct(product)

}

func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.Repo.FindById(id)
	if err != nil {
		return utils.NewNotFoundError("product not found")
	}

	return s.Repo.DeleteProduct(id)
}

func (s *ProductService) UpdateProduct(id uint, req dto.UpdateProductRequest) (models.Product, error) {

	// cek apakah product ada
	product, err := s.Repo.FindById(id)

	if err != nil {
		return product, utils.NewNotFoundError("product not found")
	}

	if req.Name != nil {
		// validasi name hanya jika dikirim
		if strings.TrimSpace(*req.Name) == "" {
			return product, utils.NewBadRequestError("product name is required")
		}
		product.Name = *req.Name
	}

	if req.Price != nil {
		// validasi price hanya jika diisi
		if *req.Price <= 0 {
			return product, utils.NewBadRequestError("price must be greater than 0")
		}
		product.Price = *req.Price
	}

	if req.Stock != nil {

		// validasi stock
		if *req.Stock < 0 {
			return product, utils.NewBadRequestError("stock cannot be negative")
		}
		product.Stock = *req.Stock
	}

	return s.Repo.UpdateProduct(id, product)
}
