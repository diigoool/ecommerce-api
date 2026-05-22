package repositories

import (
	"ecommerce-api/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	Checkout(userID uint) (models.Order, error)
	GetOrders(userID uint) ([]models.Order, error)
}

type PostgresOrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &PostgresOrderRepository{DB: db}
}

func (r *PostgresOrderRepository) Checkout(userID uint) (models.Order, error) {
	var order models.Order

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		// 1.ambil cart + items
		var cart models.Cart

		if err := tx.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
			return err
		}

		if len(cart.Items) == 0 {
			return gorm.ErrInvalidData
		}

		// 2. hitung total
		total := 0.0
		for _, item := range cart.Items {
			var product models.Product

			// Lock Product Row
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, item.ProductID).Error; err != nil {
				return err
			}
			// Check Stock
			if product.Stock < item.Quantity {
				return errors.New("insufficient stock")
			}

			// Pengurangan Stock
			product.Stock -= item.Quantity

			if err := tx.Save(&product).Error; err != nil {
				return err
			}

			total += float64(item.Quantity) * item.Product.Price
		}

		// 3. create order
		order = models.Order{
			UserID:     userID,
			TotalPrice: total,
		}

		// 4. insert order item
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for _, item := range cart.Items {
			orderItem := models.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Product.Price,
			}

			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}

		}

		// 5. clear cart

		if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return order, err
	}

	err = r.DB.Preload("Items.Product").First(&order, order.ID).Error

	return order, err
}

func (r *PostgresOrderRepository) GetOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order

	err := r.DB.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error

	return orders, err

}
