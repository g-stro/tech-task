package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
	"log"
)

type AccountRepository struct {
	db *DB
}

func NewAccountRepository(db *DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetByID(id uuid.UUID) (*model.Account, error) {
	query := `
	   SELECT id, customer_id, account_type, account_number, status, created_at, updated_at
	   FROM accounts
	   WHERE id = $1
	`

	var account model.Account
	err := r.db.conn.DB.QueryRow(query, id).Scan(
		&account.ID,
		&account.CustomerID,
		&account.AccountType,
		&account.AccountNumber,
		&account.Status,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("account %v not found", account.ID)
		return nil, nil
	}
	if err != nil {
		log.Printf("error retrieving account %v: %e", account.ID, err)
		return nil, fmt.Errorf("failed to fetch account with ID %v: %w", id, err)
	}

	return &account, nil
}
