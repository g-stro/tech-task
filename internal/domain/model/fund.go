package model

import (
	"github.com/google/uuid"
	"time"
)

type Fund struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Category   string    `json:"category" db:"category"`
	Currency   string    `json:"currency" db:"currency"`
	RiskReturn string    `json:"risk_return" db:"risk_return"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
