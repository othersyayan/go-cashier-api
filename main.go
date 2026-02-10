package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	netHttp "net/http"
	"os"
	"strings"

	"go-cashier-api/database"
	"go-cashier-api/handlers"
	"go-cashier-api/repositories"
	"go-cashier-api/services"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

func corsMiddleware(next netHttp.Handler) netHttp.Handler {
	return netHttp.HandlerFunc(func(w netHttp.ResponseWriter, r *netHttp.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(netHttp.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	mux := http.NewServeMux()

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:        viper.GetString("PORT"),
		DatabaseUrl: viper.GetString("DATABASE_URL"),
	}

	db, errDB := database.Connect(config.DatabaseUrl)
	if errDB != nil {
		log.Fatal("Failed to initialize database:", errDB)
	}

	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	mux.HandleFunc("/health", func(w netHttp.ResponseWriter, r *netHttp.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	mux.HandleFunc("/api/products", productHandler.HandleProducts)
	mux.HandleFunc("/api/products/", productHandler.HandleProductByID)

	mux.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	mux.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	mux.HandleFunc("/api/transactions/checkout", transactionHandler.HandleTransactions)

	mux.HandleFunc("/api/report/today", reportHandler.HandleTodayReport)
	mux.HandleFunc("/api/report", reportHandler.HandleReport)

	handler := corsMiddleware(mux)

	fmt.Println("Server running di localhost:" + config.Port)

	err := netHttp.ListenAndServe(":"+config.Port, handler)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
