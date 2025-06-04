package postgres

import (
	"database/sql"
	"errors"
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

	investmentStmt := `
        INSERT INTO investments (id, account_id, amount, status)
        VALUES ($1, $2, $3, $4)
        RETURNING created_at, updated_at`

	err = tx.QueryRow(investmentStmt,
		investment.ID,
		investment.AccountID,
		investment.Amount,
		investment.Status,
	).Scan(&investment.CreatedAt, &investment.UpdatedAt)
	if err != nil {
		log.Printf("failed to insert investment: %v", err)
		return nil, err
	}

	investmentFundsStmt := `
		INSERT INTO investment_funds (id, investment_id, fund_id, amount)
		VALUES ($1, $2, $3, $4)`

	for _, fund := range investment.Funds {
		investmentFundID := uuid.New()
		res, err := tx.Exec(investmentFundsStmt,
			investmentFundID,
			investment.ID,
			fund.ID,
			fund.Amount)
		if err != nil {
			log.Printf("failed to insert investment funds: %v", err)
			return nil, err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Printf("failed to get rows affected: %v", err)
			return nil, err
		}
		if rowsAffected != 1 {
			log.Printf("expected 1 row affected, got %v", rowsAffected)
			return nil, errors.New("failed to insert investment funds")
		}
	}

	if err = tx.Commit(); err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return nil, err
	}

	return investment, nil
}

// GetByAccountID retrieves all investments for a given account
func (r InvestmentRepository) GetByAccountID(id uuid.UUID) ([]*model.Investment, error) {
	investmentQuery := `
        SELECT id, account_id, amount, status, created_at, updated_at
        FROM investments
        WHERE account_id = $1
    `

	rows, err := r.db.conn.DB.Query(investmentQuery, id)
	if err != nil {
		log.Printf("error executing query: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}(rows)

	var investments []*model.Investment
	for rows.Next() {
		var investment model.Investment
		err = rows.Scan(
			&investment.ID,
			&investment.AccountID,
			&investment.Amount,
			&investment.Status,
			&investment.CreatedAt,
			&investment.UpdatedAt)
		if err != nil {
			log.Printf("error scanning row: %v", err)
			return nil, err
		}

		funds, err := r.GetInvestmentFundsByID(investment.ID)
		if err != nil {
			return nil, err
		}
		investment.Funds = append(investment.Funds, funds...)
		investments = append(investments, &investment)
	}

	if len(investments) == 0 {
		return []*model.Investment{}, nil
	}

	return investments, nil
}

// GetInvestmentFundsByID retrieves all funds for a given investment
func (r InvestmentRepository) GetInvestmentFundsByID(id uuid.UUID) ([]*model.Fund, error) {
	query := `
        SELECT f.id, f.name, f.category, f.currency, if.amount, f.risk_return
        FROM funds f
        LEFT JOIN investment_funds if ON f.id = if.fund_id
        WHERE if.investment_id = $1
	`

	rows, err := r.db.conn.DB.Query(query, id)
	if err != nil {
		log.Printf("error executing query: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}(rows)

	var funds []*model.Fund
	for rows.Next() {
		var fund model.Fund
		err = rows.Scan(
			&fund.ID,
			&fund.Name,
			&fund.Category,
			&fund.Currency,
			&fund.Amount,
			&fund.RiskReturn)
		if err != nil {
			log.Printf("error scanning row: %v", err)
			return nil, err
		}
		funds = append(funds, &fund)
	}

	if len(funds) == 0 {
		return []*model.Fund{}, nil
	}

	if err != nil {
		log.Printf("error executing query: %v", err)
		return nil, err
	}

	return funds, nil
}
