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
