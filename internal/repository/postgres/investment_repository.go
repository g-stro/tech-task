package postgres

import (
	"fmt"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
)

type InvestmentRepository struct {
	db *DB
}

func NewInvestmentRepository(db *DB) *InvestmentRepository {
	return &InvestmentRepository{db: db}
}

func (r InvestmentRepository) Create(investment *model.Investment) (*model.Investment, error) {
	tx, err := r.db.conn.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if investment.ID == uuid.Nil {
		investment.ID = uuid.New()
	}

	stmt := `
        INSERT INTO investments (id, account_id, fund_id, amount, status)
        VALUES ($1, $2, $3, $4, &5)
        RETURNING id`

	err = tx.QueryRow(stmt,
		investment.ID,
		investment.FundID,
		investment.AccountID,
		investment.Amount,
		investment.Status,
	).Scan(&investment.CreatedAt, &investment.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert investment: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return investment, nil
}
