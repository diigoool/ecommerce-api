package handlers

import (
	"ecommerce-api/mapper"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"net/http"
)

type OrderHandler struct {
	Service *services.OrderService
}

func NewOrderHandler(s *services.OrderService) *OrderHandler {
	return &OrderHandler{Service: s}
}

// Checkout godoc
// @Summary Checkout
// @Description Checkout
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Accept json
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/checkout [post]
func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {

	userID, ok := utils.GetUserID(r.Context())

	if !ok {
		utils.RespondError(w, utils.NewUnauthorizedError("user not found"))
		return
	}

	order, err := h.Service.Checkout(userID)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	response := mapper.ToOrderResponse(order)

	utils.JSON(w, http.StatusCreated, response)

}

// GetOrders godoc
// @Summary Get orders
// @Description get orders
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/orders [get]
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetUserID(r.Context())

	if !ok {
		utils.RespondError(w, utils.NewUnauthorizedError("user not found"))
		return
	}

	orders, err := h.Service.GetOrders(userID)

	if err != nil {
		utils.RespondError(w, err)
		return
	}
	response := mapper.ToOrderResponses(orders)

	utils.JSON(w, 200, response)

}
