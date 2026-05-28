package handlers

import (
	"ecommerce-api/dto"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	Service *services.CartService
}

func NewCartHandler(s *services.CartService) *CartHandler {
	return &CartHandler{Service: s}
}

// AddToCart godoc
// @Summary Add to cart
// @Description add to cart
// @Tags carts
// @Security BearerAuth
// @Produce json
// @Accept json
// @Param request body dto.AddToCartRequest true "Add To Cart"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/carts [post]
func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetUserID(r.Context())

	if !ok {
		utils.RespondError(w, utils.NewUnauthorizedError("user not found"))
		return
	}

	var req dto.AddToCartRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, utils.NewBadRequestError("invalid body request"))
		return
	}

	err := h.Service.AddToCart(userID, req)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.JSON(w, 200, "added to cart")

}

// GetCart godoc
// @Summary Get cart
// @Description get cart
// @Tags carts
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/carts [get]
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetUserID(r.Context())

	if !ok {
		utils.RespondError(w, utils.NewUnauthorizedError("user not found"))
		return
	}

	cart, err := h.Service.GetCart(userID)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.JSON(w, 200, cart)

}

// RemoveItem godoc
// @Summary remove item
// @Description remove item
// @Tags carts
// @Security BearerAuth
// @Produce json
// @Accept json
// @Param request body dto.RemoveCartItemRequest true "Remove Item"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/carts [delete]
func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetUserID(r.Context())

	if !ok {
		utils.RespondError(w, utils.NewUnauthorizedError("user not found"))
		return
	}

	var req dto.RemoveCartItemRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, utils.NewBadRequestError("invalid body request"))
		return
	}

	err := h.Service.RemoveItem(userID, req)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.JSON(w, 200, "item removed")

}
