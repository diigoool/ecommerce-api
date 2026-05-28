package dto

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoiSm9obmRvZSIsImV4cCI6MTc4MDAzMzIyMX0.Da9lgarVgBNxa-ET7obQTVvlCPLReyHlFQV5XBF5Fmk"`
}

type RegisterResponse struct {
	ID       string `json:"id" example:"1"`
	Username string `json:"username" example:"Johndoe"`
	Email    string `json:"email" example:"johndoe@example.com"`
	Role     string `json:"role" example:"user"`
}
