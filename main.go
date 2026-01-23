package main

import (
	"fmt"
	"kasir-api/routes"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			routes.GetProducts(w, r)
		case "POST":
			routes.GetProducts(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			routes.GetProduct(w, r)
		case "PUT":
			routes.UpdateProduct(w, r)
		case "DELETE":
			routes.DeleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
