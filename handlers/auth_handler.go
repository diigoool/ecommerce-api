package handlers

import (
	"ecommerce-api/dto"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"net/http"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// Register godoc
// @Summary Register user
// @Description register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register Request"
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	var req dto.RegisterRequest

	if err := utils.DecodeJSON(w, r, &req); err != nil {
		utils.RespondError(
			w,
			utils.NewBadRequestError("invalid request body"),
		)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {

		errors := utils.FormatValidationError(err)

		utils.RespondError(
			w,
			utils.NewValidationError(errors),
		)
		return
	}

	result, err := h.Service.Register(
		req.Username,
		req.Password,
		req.Email,
		"user",
	)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.JSON(w, http.StatusCreated, result)
}

// Login godoc
// @Summary Login user
// @Description login user and return jwt token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req dto.LoginRequest

	if err := utils.DecodeJSON(w, r, &req); err != nil {
		utils.RespondError(
			w,
			utils.NewBadRequestError("invalid request body"),
		)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {

		errors := utils.FormatValidationError(err)

		utils.RespondError(
			w,
			utils.NewValidationError(errors),
		)
		return
	}

	token, err := h.Service.Login(
		req.Username,
		req.Password,
	)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
