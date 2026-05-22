package seeders

import (
	"ecommerce-api/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {

	var user models.User

	err := db.
		Where("email = ?", "admin@example.com").
		First(&user).Error

	// admin sudah ada
	if err == nil {
		log.Println("admin already exists")
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("admin123"),
		bcrypt.DefaultCost,
	)

	admin := models.User{
		Username: "admin",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Println("failed seed admin:", err)
		return
	}

	log.Println("admin seeded")
}
