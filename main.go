package main

import (
	"encoding/json"
	"fmt"
	_ "go-cashier-api/docs"
	"go-cashier-api/internal/delivery/http"
	"go-cashier-api/internal/repository"
	"go-cashier-api/internal/usecase"
	netHttp "net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

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
	repo := repository.NewInMemoryProductRepository()
	productUsecase := usecase.NewProductUsecase(repo)
	productHandler := http.NewProductHandler(productUsecase)

	categoryRepo := repository.NewInMemoryCategoryRepository()
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := http.NewCategoryHandler(categoryUsecase)

	productHandler.RegisterRoutes()
	categoryHandler.RegisterRoutes()

	netHttp.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	netHttp.HandleFunc("/health", func(w netHttp.ResponseWriter, r *netHttp.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)

	handler := corsMiddleware(netHttp.DefaultServeMux)
	err := netHttp.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
