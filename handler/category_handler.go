package handler

import (
	"encoding/json"
	"kasir-api/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "success",
			"message": "Hello World",
			"code":    "200",
			"data":    models.Categories,
		})
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		var category models.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "error",
				"message": "Error decoding request body",
				"code":    "400",
			})
			return
		}
		category.ID = len(models.Categories) + 1
		models.Categories = append(models.Categories, category)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "success",
			"message": "Category created successfully",
			"code":    "201",
			"data":    category,
		})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	log.Println("Get category by ID", idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	log.Println("Get category by ID", id)
	for _, category := range models.Categories {
		if category.ID == id {
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "success",
				"message": "Get category by ID",
				"code":    "200",
				"data":    category,
			})
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	log.Println("Update category by ID", idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	log.Println("Update category by ID", id)
	for i, category := range models.Categories {
		if category.ID == id {
			log.Println("Category found")
			updateCategory.ID = id
			models.Categories[i] = updateCategory
			log.Println("Category updated successfully")
			log.Println(updateCategory)
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "success",
				"message": "Category updated successfully",
				"code":    "200",
				"data":    updateCategory,
			})
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	log.Println("Delete category by ID", idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	log.Println("Delete category by ID", id)
	for i, category := range models.Categories {
		if category.ID == id {
			models.Categories = append(models.Categories[:i], models.Categories[i+1:]...)
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "success",
				"message": "Category deleted successfully",
				"code":    "200",
			})
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}
