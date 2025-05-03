package models

import "time"

type Investment struct {
	ID             int       `json:"id"`
	AssetType      string    `json:"asset_type"`
	AssetName      string    `json:"asset_name"`
	Quantity       float64   `json:"quantity"`
	InvestedAmount float64   `json:"invested_amount"`
	AccountName    string    `json:"account_name"`
	CreatedAt      time.Time `json:"created_at"`
}

type AggregateInvestment struct {
	AssetType       string  `json:"asset_type"`
	AccountName     string  `json:"account_name"`
	TotalQuantity   float64 `json:"total_quantity"`
	TotalInvestment float64 `json:"total_investment"`
}
