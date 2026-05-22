package models

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
