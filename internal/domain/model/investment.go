package model

import (
	"github.com/google/uuid"
	"time"
)

type Investment struct {
	ID        uuid.UUID `json:"id" db:"id"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`
	FundID    uuid.UUID `json:"fund_id" db:"fund_id"`
	Amount    float64   `json:"amount" db:"amount"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
