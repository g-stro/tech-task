package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
	"log"
)

type FundRepository struct {
	db *DB
}

func NewFundRepository(db *DB) *FundRepository {
	return &FundRepository{db: db}
}

func (r FundRepository) GetByID(id uuid.UUID) (*model.Fund, error) {
	query := `
        SELECT id, name, category, currency, risk_return
        FROM funds
        WHERE id = $1
    `

	var fund model.Fund
	err := r.db.conn.DB.QueryRow(query, id).Scan(
		&fund.ID,
		&fund.Name,
		&fund.Category,
		&fund.Currency,
		&fund.RiskReturn,
	)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("fund %v not found", id)
		return nil, nil
	}
	if err != nil {
		log.Printf("error retrieving fund %v: %v", id, err)
		return nil, fmt.Errorf("failed to fetch fund with ID %v: %w", id, err)
	}

	return &fund, nil
}
