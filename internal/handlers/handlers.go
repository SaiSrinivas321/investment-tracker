package handlers

import (
	"encoding/json"
	"investment-tracker/internal/models"
	"investment-tracker/internal/services"
	"net/http"
)

func InvestmentsHandler(service *services.InvestmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			addInvestment(w, r, service)
		case http.MethodGet:
			listInvestments(w, service)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func addInvestment(w http.ResponseWriter, r *http.Request, service *services.InvestmentService) {
	var inv models.Investment
	err := json.NewDecoder(r.Body).Decode(&inv)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdInvestment, err := service.AddInvestment(inv)
	if err != nil {
		http.Error(w, "Failed to insert investment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdInvestment)
}

func listInvestments(w http.ResponseWriter, service *services.InvestmentService) {
	investments, err := service.ListInvestments()
	if err != nil {
		http.Error(w, "Failed to fetch investments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(investments)
}

func AggregateInvestmentsHandler(service *services.InvestmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aggregates, err := service.AggregateInvestments()
		if err != nil {
			http.Error(w, "Failed to fetch aggregates", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(aggregates)
	}
}
