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

func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {

	userID, ok := utils.GetUserID(r.Context())

	if !ok {
		utils.RespondError(w, utils.NewUnauthorizedError("user not found"))
		return
	}

	order, err := h.Service.Checkout(userID)

	if err != nil {
		utils.RespondError(w, utils.NewBadRequestError("cart is empty"))
		return
	}

	response := mapper.ToOrderResponse(order)

	utils.JSON(w, 200, response)

}

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
