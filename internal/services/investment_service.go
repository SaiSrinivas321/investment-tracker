package services

import (
	"database/sql"
	"fmt"
	"investment-tracker/internal/models"
	"strings"
)

type InvestmentService struct {
	db *sql.DB
}

func NewInvestmentService(db *sql.DB) *InvestmentService {
	return &InvestmentService{db: db}
}

func (s *InvestmentService) AddInvestment(inv models.Investment) (models.Investment, error) {
	query := `INSERT INTO investments (asset_type, asset_name, quantity, invested_amount, account_name)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	err := s.db.QueryRow(query, inv.AssetType, inv.AssetName, inv.Quantity, inv.InvestedAmount, inv.AccountName).
		Scan(&inv.ID, &inv.CreatedAt)

	if err != nil {
		return inv, err
	}

	return inv, nil
}

func (s *InvestmentService) ListInvestments() ([]models.Investment, error) {
	rows, err := s.db.Query(`SELECT id, asset_type, asset_name, quantity, invested_amount, account_name, created_at FROM investments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	investments := []models.Investment{}
	for rows.Next() {
		var inv models.Investment
		err := rows.Scan(&inv.ID, &inv.AssetType, &inv.AssetName, &inv.Quantity, &inv.InvestedAmount, &inv.AccountName, &inv.CreatedAt)
		if err != nil {
			return nil, err
		}
		investments = append(investments, inv)
	}

	return investments, nil
}

// AggregateInvestments aggregates investment data based on filters and group by fields
func (s *InvestmentService) AggregateInvestments(filters map[string]interface{}, groupByFields []string) (map[string]interface{}, error) {
	// Construct the base query
	query := `SELECT asset_type, account_name, SUM(invested_amount) AS total_investment 
              FROM investments`

	// Prepare conditions for the WHERE clause
	var conditions []string
	var args []interface{}
	placeholderIndex := 1 // PostgreSQL parameterized query placeholder index

	// Add filter conditions to the WHERE clause
	for key, value := range filters {
		conditions = append(conditions, fmt.Sprintf("%s = $%d", key, placeholderIndex))
		args = append(args, value)
		placeholderIndex++
	}

	// Append WHERE clause if any conditions are present
	if len(conditions) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditions, " AND "))
	}

	// Ensure asset_type is always included in the GROUP BY clause
	groupByFields = appendIfMissing(groupByFields, "asset_type")
	groupByFields = appendIfMissing(groupByFields, "account_name") // Ensure account_name is also in GROUP BY if selected

	// Add GROUP BY clause
	query += fmt.Sprintf(" GROUP BY %s", strings.Join(groupByFields, ", "))

	// Execute the query with the provided filters and arguments
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Collect results from the query
	var aggregates []models.AggregateInvestment
	var cumulativeInvestment float64

	for rows.Next() {
		var agg models.AggregateInvestment
		if err := rows.Scan(&agg.AssetType, &agg.AccountName, &agg.TotalInvestment); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		aggregates = append(aggregates, agg)
		cumulativeInvestment += agg.TotalInvestment
	}

	// Construct the response with the cumulative investment and the individual aggregates
	response := map[string]interface{}{
		"cuminvestment": cumulativeInvestment,
		"investments":   aggregates,
	}

	// Return the aggregated results along with the cumulative investment
	return response, nil
}

// appendIfMissing ensures "asset_type" and "account_name" are always included in the GROUP BY clause
func appendIfMissing(slice []string, value string) []string {
	for _, v := range slice {
		if v == value {
			return slice
		}
	}
	return append(slice, value)
}
