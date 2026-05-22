package handlers

import (
	"ecommerce-api/dto"
	"ecommerce-api/mapper"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	Service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	pageStr := query.Get("page")
	limitStr := query.Get("limit")
	search := query.Get("q")

	page := 1
	limit := 10

	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			utils.RespondError(w, utils.NewBadRequestError("invalid page"))
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			utils.RespondError(w, utils.NewBadRequestError("invalid limit"))
			return
		}

	}

	if limit > 100 {
		limit = 100
	}

	products, err := h.Service.GetAllProduct(page, limit, search)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	responses := mapper.ToProductResponses(products)

	utils.JSONWithMeta(w, http.StatusOK, responses, page, limit)

}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(idStr)

	if err != nil || idInt <= 0 {
		utils.RespondError(w, utils.NewBadRequestError("invalid format id"))
		return
	}

	product, err := h.Service.GetProductById(uint(idInt))

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	response := mapper.ToProductResponse(product)

	utils.JSON(w, http.StatusOK, response)

}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest

	if r.Header.Get("Content-Type") != "application/json" {
		utils.Error(w, http.StatusUnsupportedMediaType, "INVALID_CONTENT_TYPE", "Content-Type must be application/json")
		return
	}

	if err := utils.DecodeJSON(w, r, &req); err != nil {
		utils.RespondError(w, utils.NewBadRequestError("invalid body request"))
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		errors := utils.FormatValidationError(err)
		utils.RespondError(w, utils.NewValidationError(errors))
		return
	}

	result, err := h.Service.CreateProduct(req)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	response := mapper.ToProductResponse(result)

	utils.JSON(w, http.StatusCreated, response)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil || idInt <= 0 {
		utils.RespondError(w, utils.NewBadRequestError("invalid id"))
		return
	}

	if err := h.Service.DeleteProduct(uint(idInt)); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.JSON(w, http.StatusOK, "Product deleted")
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(idStr)

	if err != nil || idInt <= 0 {
		utils.RespondError(w, utils.NewBadRequestError("invalid id"))
		return
	}

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		utils.RespondError(w, utils.NewBadRequestError("content-type must be application/json"))
		return
	}

	var req dto.UpdateProductRequest

	if err := utils.DecodeJSON(w, r, &req); err != nil {
		utils.RespondError(w, utils.NewBadRequestError(err.Error()))
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		errors := utils.FormatValidationError(err)
		utils.RespondError(w, utils.NewValidationError(errors))
		return
	}

	result, err := h.Service.UpdateProduct(uint(idInt), req)

	if err != nil {
		utils.RespondError(w, err)
		return
	}

	response := mapper.ToProductResponse(result)

	utils.JSON(w, http.StatusOK, response)

}
