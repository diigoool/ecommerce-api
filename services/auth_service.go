package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"ecommerce-api/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(username, password, email, role string) (models.User, error) {
	if username == "" || password == "" {
		return models.User{}, errors.New("Username and Password required")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: username,
		Password: string(hashed),
		Email:    email,
		Role:     "user",
	}

	return s.Repo.CreateUser(user)

}

func (s *AuthService) Login(username, password string) (string, error) {

	user, err := s.Repo.FindByUsername(username)

	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("Invalid Credential")
	}

	// create token

	token, err := utils.GenerateToken(user.ID, user.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}
