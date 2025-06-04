package contract

import (
	"github.com/google/uuid"
)

type InvestmentRequest struct {
	AccountID uuid.UUID `json:"account_id"`
	FundID    uuid.UUID `json:"fund_id"`
	Amount    float64   `json:"amount"`
}

type InvestmentResponse struct {
	ID string `json:"id"`
}

type InvestmentDetailsResponse struct {
	ID        string                    `json:"id"`
	Amount    float64                   `json:"amount"`
	Status    string                    `json:"status"`
	CreatedAt string                    `json:"created_at"`
	UpdatedAt string                    `json:"updated_at"`
	Funds     []InvestmentFundsResponse `json:"funds"`
}

type InvestmentFundsResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	Currency   string  `json:"currency"`
	Amount     float64 `json:"amount"`
	RiskReturn string  `json:"risk_return"`
}
