package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
)

type AccountRepository struct {
	db *DB
}

func NewAccountRepository(db *DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetByID(id uuid.UUID) (*model.Account, error) {
	query := `
	   SELECT id, customer_id, account_type, account_number, status, created_at, updated_at)
	   FROM accounts
	   WHERE account_id = $1
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
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account with ID %v: %w", id, err)
	}

	return nil, nil
}
