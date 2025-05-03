package main

import (
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
