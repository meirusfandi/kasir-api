package handler

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetCategories(w, r)
	case "POST":
		h.CreateCategory(w, r)
	default:
		helpers.SendResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Get all categories", categories)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result, err := h.categoryService.CreateCategory(category)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusCreated, "Category created successfully", result)
}

func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetCategory(w, r)
	case "PUT":
		h.UpdateCategory(w, r)
	case "DELETE":
		h.DeleteCategory(w, r)
	default:
		helpers.SendResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}
	category, err := h.categoryService.GetCategoryByID(id)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Get category by ID", category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	// get data dari request
	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	category.ID = id
	result, err := h.categoryService.UpdateCategory(category)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Update category by ID", result)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	err = h.categoryService.DeleteCategory(id)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Category deleted successfully", nil)
}
