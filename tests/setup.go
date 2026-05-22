package tests

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"log"

	"github.com/joho/godotenv"
)

func SetupTestDB() {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)
}
