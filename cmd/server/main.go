package main

import (
	"log"
	"net/http"

	"investment-tracker/internal/db"
	"investment-tracker/internal/handlers"
	"investment-tracker/internal/services"
)

func main() {
	// Initialize DB
	err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	// Start investment service
	investmentService := services.NewInvestmentService(db.GetDB())

	// Handlers
	http.HandleFunc("/investments", handlers.InvestmentsHandler(investmentService))
	http.HandleFunc("/investments/aggregate", handlers.AggregateInvestmentsHandler(investmentService))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
