package handlers

import (
	"encoding/json"
	"fmt"
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

// AggregateInvestmentsHandler handles requests for aggregated investment data
func AggregateInvestmentsHandler(service *services.InvestmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse filters and groupBy parameters from the URL query string
		filters := parseFilters(r)
		groupByFields := parseGroupByFields(r)

		// Call service function to get aggregated investment data
		aggregates, err := service.AggregateInvestments(filters, groupByFields)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to fetch aggregates: %v", err), http.StatusInternalServerError)
			return
		}

		// Set response header and encode response as JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(aggregates); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		}
	}
}

// parseFilters parses the query parameters into a map, excluding group_by.
func parseFilters(r *http.Request) map[string]interface{} {
	filters := make(map[string]interface{})
	for key, values := range r.URL.Query() {
		// Skip group_by parameter, it's handled separately
		if key != "group_by" && len(values) > 0 {
			filters[key] = values[0] // Use first value for simplicity
		}
	}
	return filters
}

// parseGroupByFields parses the group_by parameters, defaulting to "asset_type".
func parseGroupByFields(r *http.Request) []string {
	groupByFields := r.URL.Query()["group_by"]
	if len(groupByFields) == 0 {
		groupByFields = []string{"asset_type"} // Default group by asset_type if none provided
	}
	return groupByFields
}
