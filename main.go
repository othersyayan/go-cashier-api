package main

import (
	"encoding/json"
	"fmt"
	"go-cashier-api/internal/delivery/http"
	"go-cashier-api/internal/repository"
	"go-cashier-api/internal/usecase"
	netHttp "net/http"
)

func main() {
	repo := repository.NewInMemoryProductRepository()
	productUsecase := usecase.NewProductUsecase(repo)
	productHandler := http.NewProductHandler(productUsecase)

	categoryRepo := repository.NewInMemoryCategoryRepository()
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := http.NewCategoryHandler(categoryUsecase)

	productHandler.RegisterRoutes()
	categoryHandler.RegisterRoutes()

	netHttp.HandleFunc("/health", func(w netHttp.ResponseWriter, r *netHttp.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:8080")

	err := netHttp.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
