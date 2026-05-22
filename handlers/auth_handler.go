package handlers

import (
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "INVALID_BODY", "invalid request")
		return
	}

	result, err := h.Service.Register(req.Username, req.Password, req.Email, req.Role)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	utils.JSON(w, http.StatusCreated, result)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "INVALID_BODY", "invalid request")
		return
	}

	token, err := h.Service.Login(req.Username, req.Password)

	if err != nil {
		utils.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
