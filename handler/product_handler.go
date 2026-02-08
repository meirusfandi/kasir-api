package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/services"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAllProducts(w, r)
	case http.MethodPost:
		h.CreateProduct(w, r)
	}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := h.service.GetAllProducts(name)
	if err == sql.ErrNoRows {
		helpers.SendResponse(w, http.StatusOK, "No products found", []models.Product{})
		return
	}

	if err != nil {
		log.Println("error", err)
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Get all products", products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result, err := h.service.CreateProduct(product)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusCreated, "Product created successfully", result)
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetProduct(w, r)
	case http.MethodPut:
		h.UpdateProduct(w, r)
	case http.MethodDelete:
		h.DeleteProduct(w, r)
	}
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	log.Println("idStr", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("error", err)
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}
	log.Println("id", id)

	w.Header().Set("Content-Type", "application/json")
	product, err := h.service.GetProductByID(id)
	log.Println("product", product)

	if err != nil {
		if err.Error() == "Product not found" {
			helpers.SendResponse(w, http.StatusOK, fmt.Sprintf("Data with id %d not found", id), nil)
			return
		}
		log.Println("error", err)
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}
	log.Println("product", product)

	helpers.SendResponse(w, http.StatusOK, "Get product by ID", product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	// get data dari request
	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	product.ID = id
	_, err = h.service.UpdateProduct(product)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Product updated successfully", product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Product deleted successfully", nil)
}
