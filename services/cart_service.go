package services

import (
	"ecommerce-api/dto"
	"ecommerce-api/mapper"
	"ecommerce-api/repositories"
	"ecommerce-api/utils"
)

type CartService struct {
	Repo repositories.CartRepository
}

func NewCartService(repo repositories.CartRepository) *CartService {
	return &CartService{Repo: repo}
}

func (s *CartService) AddToCart(userID uint, req dto.AddToCartRequest) error {

	if req.Qty <= 0 {
		return utils.NewBadRequestError("invalid quantity")
	}

	cart, err := s.Repo.GetOrCreateCart(userID)

	if err != nil {
		return err
	}

	product, err := s.Repo.GetProductByID(req.ProductID)

	if err != nil {
		return utils.NewNotFoundError("product not found")
	}

	existingItem, err := s.Repo.GetCartItem(cart.ID, req.ProductID)

	totalQty := req.Qty

	if err == nil {
		totalQty += existingItem.Quantity
	}

	if totalQty > product.Stock {
		return utils.NewBadRequestError("stock not enough")
	}

	return s.Repo.AddToCart(cart.ID, req.ProductID, req.Qty)

}

func (s *CartService) GetCart(userID uint) (dto.CartResponse, error) {
	cart, err := s.Repo.GetCart(userID)
	if err != nil {
		return dto.CartResponse{}, err
	}

	response := mapper.ToCartResponse(cart)

	return response, nil
}

func (s *CartService) RemoveItem(userID uint, req dto.RemoveCartItemRequest) error {

	cart, err := s.Repo.GetOrCreateCart(userID)

	if err != nil {
		return err
	}

	return s.Repo.RemoveItem(cart.ID, req.ProductID)
}
