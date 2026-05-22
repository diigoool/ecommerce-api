package repositories

import (
	"ecommerce-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type PostgresUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User

	err := r.DB.Where("username = ?", username).First(&user).Error

	return user, err

}

func (r *PostgresUserRepository) CreateUser(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}
