package repositories

import (
	"ecommerce-api/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct(page, limit int, search string) ([]models.Product, error)
	FindById(id uint) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id uint) error
	UpdateProduct(id uint, product models.Product) (models.Product, error)
}

type PostgresProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &PostgresProductRepository{DB: db}
}

func (r *PostgresProductRepository) GetAllProduct(page, limit int, search string) ([]models.Product, error) {
	var products []models.Product

	offset := (page - 1) * limit

	query := r.DB.Model(&models.Product{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	err := query.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *PostgresProductRepository) FindById(id uint) (models.Product, error) {
	var product models.Product

	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *PostgresProductRepository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.DB.Create(&product).Error
	return product, err
}

func (r *PostgresProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}

func (r *PostgresProductRepository) UpdateProduct(id uint, product models.Product) (models.Product, error) {
	err := r.DB.Save(&product).Error

	return product, err
}
