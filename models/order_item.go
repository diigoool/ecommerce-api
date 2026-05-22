package models

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64

	Product Product `gorm:"foreignKey:ProductID"`
}
