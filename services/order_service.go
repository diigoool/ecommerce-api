package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
)

type OrderService struct {
	Repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) *OrderService {
	return &OrderService{Repo: repo}
}

func (s *OrderService) Checkout(userID uint) (models.Order, error) {
	return s.Repo.Checkout(userID)
}

func (s *OrderService) GetOrders(userID uint) ([]models.Order, error) {
	return s.Repo.GetOrders(userID)
}
