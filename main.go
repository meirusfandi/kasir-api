package main

import (
	"fmt"
	"io"
	"kasir-api/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Logging setup
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	multiWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multiWriter)

	// Middleware for logging
	logMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
			next(w, r)
		}
	}
	// handle on products
	http.HandleFunc("/api/v1/products", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			routes.GetProducts(w, r)
		case "POST":
			routes.GetProducts(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	http.HandleFunc("/api/v1/products/", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
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
	}))

	// handle on categories
	http.HandleFunc("/api/v1/categories", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			routes.GetCategories(w, r)
		case "POST":
			routes.GetCategories(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/api/v1/categories/", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			routes.GetCategory(w, r)
		case "PUT":
			routes.UpdateCategory(w, r)
		case "DELETE":
			routes.DeleteCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
