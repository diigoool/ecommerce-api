package models

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	TotalPrice float64

	Items []OrderItem `gorm:"foreignKey:OrderID"`
}
