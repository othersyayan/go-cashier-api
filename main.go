package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	netHttp.HandleFunc("/health", func(w netHttp.ResponseWriter, r *netHttp.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	netHttp.HandleFunc("/api/products", productHandler.HandleProducts)

	fmt.Println("Server running di localhost:" + config.Port)

	err := netHttp.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
