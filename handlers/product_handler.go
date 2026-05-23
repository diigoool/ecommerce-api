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

// GetProducts godoc
// @Summary Get all products
// @Description get all products
// @Tags products
// @Security BearerAuth
// @Produce json
// @Param q query string false "name search by name"
// @Param limit query int false "limit" default(10)
// @Param page query int false "page" default(1)
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/product [get]
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

// GetProductByID godoc
// @Summary Get Product By Id
// @Description get product by id
// @Tags products
// @Security BearerAuth
// @Produce json
// @Param id path uint true "id"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/product/{id} [get]
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

// CreateProduct godoc
// @Summary Create product
// @Description create new product
// @Tags products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Create Product Request"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/product/create [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
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

// DeleteProduct godoc
// @Summary Delete product
// @Description delete product
// @Tags products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path uint true "id"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/product/delete/{id} [delete]
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

// UpdateProduct godoc
// @Summary Update product
// @Description update product
// @Tags products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.UpdateProductRequest true "Update Product Request"
// @Param id path uint true "id"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /api/product/update/{id} [patch]
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
