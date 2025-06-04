package model

import (
	"github.com/google/uuid"
)

type Fund struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Category   string    `json:"category" db:"category"`
	Currency   string    `json:"currency" db:"currency"`
	Amount     float64   `json:"amount" db:"amount"`
	RiskReturn string    `json:"risk_return" db:"risk_return"`
}
