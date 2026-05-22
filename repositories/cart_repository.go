package repositories

import (
	"ecommerce-api/models"
	"errors"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetOrCreateCart(userID uint) (models.Cart, error)
	AddToCart(cartID, productID uint, qty int) error
	GetCart(userID uint) (models.Cart, error)
	RemoveItem(cartID, productID uint) error
	GetCartItem(cartID, productID uint) (models.CartItem, error)
	GetProductByID(productID uint) (models.Product, error)
}

type PostgresCartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &PostgresCartRepository{DB: db}
}

func (r *PostgresCartRepository) GetOrCreateCart(userID uint) (models.Cart, error) {
	var cart models.Cart

	err := r.DB.Where("user_id = ?", userID).First(&cart).Error

	if err == nil {
		return cart, nil
	}

	cart = models.Cart{UserID: userID}
	err = r.DB.Create(&cart).Error

	return cart, err

}

func (r *PostgresCartRepository) AddToCart(cartID, productID uint, qty int) error {
	var item models.CartItem

	err := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error

	if err == nil {
		item.Quantity += qty
		return r.DB.Save(&item).Error
	}

	item = models.CartItem{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  qty,
	}

	return r.DB.Create(&item).Error

}

func (r *PostgresCartRepository) GetCart(userID uint) (models.Cart, error) {
	var cart models.Cart

	err := r.DB.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error

	return cart, err
}

func (r *PostgresCartRepository) RemoveItem(cartID, productID uint) error {

	result := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&models.CartItem{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("item not found in cart")
	}

	return nil

}

func (r *PostgresCartRepository) GetProductByID(productID uint) (models.Product, error) {

	var product models.Product

	err := r.DB.First(&product, productID).Error

	return product, err
}

func (r *PostgresCartRepository) GetCartItem(cartID, productID uint) (models.CartItem, error) {

	var item models.CartItem

	err := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error

	return item, err
}
