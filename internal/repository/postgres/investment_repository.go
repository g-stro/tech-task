package postgres

import (
	"database/sql"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
	"log"
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
		log.Printf("failed to begin transaction: %v", err)
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
        VALUES ($1, $2, $3, $4, $5)
        RETURNING created_at, updated_at`

	err = tx.QueryRow(stmt,
		investment.ID,
		investment.AccountID,
		investment.FundID,
		investment.Amount,
		investment.Status,
	).Scan(&investment.CreatedAt, &investment.UpdatedAt)
	if err != nil {
		log.Printf("failed to insert investment: %v", err)
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return nil, err
	}

	return investment, nil
}

func (r InvestmentRepository) GetByAccountID(id uuid.UUID) ([]*model.Investment, error) {
	query := `
        SELECT id, account_id, fund_id, amount, status, created_at, updated_at
        FROM investments
        WHERE account_id = $1
    `

	rows, err := r.db.conn.DB.Query(query, id)
	if err != nil {
		log.Printf("error executing query: %e", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Printf("error closing rows: %e", err)
		}
	}(rows)

	var investments []*model.Investment
	for rows.Next() {
		var investment model.Investment
		err = rows.Scan(
			&investment.ID,
			&investment.AccountID,
			&investment.FundID,
			&investment.Amount,
			&investment.Status,
			&investment.CreatedAt,
			&investment.UpdatedAt)
		if err != nil {
			log.Printf("error scanning row: %e", err)
			return nil, err
		}
		investments = append(investments, &investment)
	}

	if len(investments) == 0 {
		return nil, nil
	}

	return investments, nil
}
