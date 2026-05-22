package tests

import (
	"bytes"
	"ecommerce-api/config"
	"ecommerce-api/handlers"
	"ecommerce-api/repositories"
	"ecommerce-api/services"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	repo := repositories.NewUserRepository(config.DB)

	service := services.NewAuthService(repo)

	handler := handlers.NewAuthHandler(service)

	r := chi.NewRouter()

	r.Post("/api/auth/login", handler.Login)

	payload := map[string]string{
		"username": "user",
		"password": "user123",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestRegister_Success(t *testing.T) {

	config.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")

	err := godotenv.Load("../.env")

	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	repo := repositories.NewUserRepository(config.DB)

	service := services.NewAuthService(repo)

	handler := handlers.NewAuthHandler(service)

	r := chi.NewRouter()

	r.Post("/api/auth/register", handler.Register)

	payload := map[string]string{
		"email":    "user@example.com",
		"username": "user",
		"password": "user123",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/register",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestLogin_Failed(t *testing.T) {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	repo := repositories.NewUserRepository(config.DB)

	service := services.NewAuthService(repo)

	handler := handlers.NewAuthHandler(service)

	r := chi.NewRouter()

	r.Post("/api/auth/login", handler.Login)

	payload := map[string]string{
		"username": "user",
		"password": "wrongpassword",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

}

func TestRegister_InvalidRegister(t *testing.T) {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	repo := repositories.NewUserRepository(config.DB)

	service := services.NewAuthService(repo)

	handler := handlers.NewAuthHandler(service)

	r := chi.NewRouter()

	r.Post("/api/auth/register", handler.Register)

	payload := map[string]string{
		"email":    "",
		"username": "",
		"password": "",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/register",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

}
