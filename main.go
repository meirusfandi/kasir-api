package main

import (
	"fmt"
	"io"
	"kasir-api/database"
	"kasir-api/handler"
	"kasir-api/models"
	"kasir-api/repository"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	// Setup viper
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := models.Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		PORT:       viper.GetString("PORT"),
	}

	// Setup database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := database.InitDB(connStr)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

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

	// Setup repository, service, and handler
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// handle on products
	http.Handle("/api/v1/products", logMiddleware(productHandler.HandleProducts))
	http.Handle("/api/v1/products/", logMiddleware(productHandler.HandleProductByID))

	// handle on categories
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	http.Handle("/api/v1/categories", logMiddleware(categoryHandler.HandleCategories))
	http.Handle("/api/v1/categories/", logMiddleware(categoryHandler.HandleCategoryByID))

	// Transaction
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	http.Handle("/api/v1/checkout", logMiddleware(transactionHandler.HandleCheckout)) // POST

	// Report
	reportRepo := repository.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handler.NewReportHandler(reportService)

	http.Handle("/api/report/hari-ini", logMiddleware(reportHandler.GetDailyReport))
	http.Handle("/api/report", logMiddleware(reportHandler.GetReport))

	fmt.Println("Starting server on :" + config.PORT)
	if err := http.ListenAndServe(":"+config.PORT, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
