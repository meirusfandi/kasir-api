package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")
	products, err := h.service.GetAllProducts()
	if err == sql.ErrNoRows {
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "success",
			"code":    "200",
			"message": "No products found",
			"data":    []models.Product{},
		})
		return
	}

	if err != nil {
		log.Println("error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("products", products)

	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"code":    "200",
		"message": "Get all products",
		"data":    products,
	})
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	result, err := h.service.CreateProduct(product)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"code":    "201",
		"message": "Product created successfully",
		"data":    result,
	})
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
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	log.Println("id", id)

	w.Header().Set("Content-Type", "application/json")
	product, err := h.service.GetProductByID(id)
	log.Println("product", product)

	if err != nil {
		if err.Error() == "Product not found" {
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "success",
				"code":    "200",
				"message": fmt.Sprintf("Data with id %d not found", id),
				"data":    nil,
			})
			return
		}
		log.Println("error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("product", product)

	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"code":    "200",
		"message": "Get product by ID",
		"data":    product,
	})
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	product.ID = id
	_, err = h.service.UpdateProduct(product)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"code":    "200",
		"message": "Product updated successfully",
		"data":    product,
	})
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"code":    "200",
		"message": "Product deleted successfully",
	})
}
