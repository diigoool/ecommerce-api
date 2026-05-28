package dto

type RegisterResponse struct {
	ID       string `json:"id" example:"1"`
	Username string `json:"username" example:"Johndoe"`
	Email    string `json:"email" example:"johndoe@example.com"`
	Role     string `json:"role" example:"user"`
}
