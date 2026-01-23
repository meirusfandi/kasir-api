package routes

import (
	"encoding/json"
	"kasir-api/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func TestApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Hello World",
		"code":    "200",
	})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		log.Println("GET request")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "success",
			"code":    "200",
			"message": "Get all products",
			"data":    models.Products,
		})
	case "POST":
		log.Println("POST request")
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Println("Error decoding request: ", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		product.ID = len(models.Products) + 1
		models.Products = append(models.Products, product)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "success",
			"code":    "201",
			"message": "Product created successfully",
			"data":    product,
		})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	for _, product := range models.Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "success",
				"code":    "200",
				"message": "Get product by ID",
				"data":    product,
			})
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, product := range models.Products {
		if product.ID == id {
			log.Println("Product found")
			updateProduct.ID = id
			models.Products[i] = updateProduct
			log.Println("Product updated successfully")
			log.Println(updateProduct)
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "success",
				"code":    "200",
				"message": "Product updated successfully",
				"data":    updateProduct,
			})
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	for i, product := range models.Products {
		if product.ID == id {
			models.Products = append(models.Products[:i], models.Products[i+1:]...)
			json.NewEncoder(w).Encode(map[string]any{
				"status":  "success",
				"code":    "200",
				"message": "Product deleted successfully",
			})
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
