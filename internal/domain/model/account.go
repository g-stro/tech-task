package model

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CustomerID    uuid.UUID `json:"customer_id" db:"customer_id"`
	AccountType   string    `json:"account_type" db:"account_type"`
	AccountNumber string    `json:"account_number" db:"account_number"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
