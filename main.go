package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/m-lejwoda/batna-lublin/internal/db"
	handler "github.com/m-lejwoda/batna-lublin/internal/handlers"
	service "github.com/m-lejwoda/batna-lublin/internal/services"
	"github.com/m-lejwoda/batna-lublin/internal/storage"
)

func main() {
	fmt.Println("Start application")
	DatabaseURL := os.Getenv("DATABASE_URL")
	database, err := db.Connect(DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}
	defer database.Close()
	newInvestmentStorage := storage.NewInvestmentStorage(database)
	newInvestmentService := service.NewInvestmentService(newInvestmentStorage)
	newInvestmentHandler := handler.NewInvestmentHandler(newInvestmentService)
	mux := http.NewServeMux()
	handler.RegisterInvestmentRoutes(mux, &newInvestmentHandler)
}
