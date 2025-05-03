package services

import (
	"database/sql"
	"investment-tracker/internal/models"
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

func (s *InvestmentService) AggregateInvestments() ([]models.AggregateInvestment, error) {
	rows, err := s.db.Query(`SELECT asset_type, account_name, SUM(quantity) as total_quantity, SUM(invested_amount) as total_investment
	                       FROM investments
	                       GROUP BY asset_type, account_name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	aggregates := []models.AggregateInvestment{}
	for rows.Next() {
		var agg models.AggregateInvestment
		err := rows.Scan(&agg.AssetType, &agg.AccountName, &agg.TotalQuantity, &agg.TotalInvestment)
		if err != nil {
			return nil, err
		}
		aggregates = append(aggregates, agg)
	}

	return aggregates, nil
}
