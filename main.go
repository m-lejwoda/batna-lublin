package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/m-lejwoda/batna-lublin/internal/db"
	handler "github.com/m-lejwoda/batna-lublin/internal/handlers"
	scrappers "github.com/m-lejwoda/batna-lublin/internal/scrapers"
	service "github.com/m-lejwoda/batna-lublin/internal/services"
	"github.com/m-lejwoda/batna-lublin/internal/storage"
)

func main() {
	fmt.Println("Start application")
	DatabaseURL := os.Getenv("DATABASE_URL")
	fmt.Println("sdasdsa", DatabaseURL)
	database, err := db.Connect(DatabaseURL)

	if err != nil {
		log.Fatalf("Error connection")
	}
	defer database.Close()
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "lalak":
			fmt.Println("Lalak")
			scrappers.GetLalakFlats()
			return

		case "multimel":
			fmt.Println("Multimel")
			return

		default:
			fmt.Println("Not working")
			return
		}
	}
	newInvestmentStorage := storage.NewInvestmentStorage(database)
	newInvestmentService := service.NewInvestmentService(newInvestmentStorage)
	newInvestmentHandler := handler.NewInvestmentHandler(newInvestmentService)
	mux := http.NewServeMux()
	handler.RegisterInvestmentRoutes(mux, &newInvestmentHandler)
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Error")
	}
}
