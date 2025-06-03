package postgres

import (
	"database/sql"
	"errors"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
	"log"
)

type CustomerRepository struct {
	db *DB
}

func NewCustomerRepository(db *DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetByID(id uuid.UUID) (*model.Customer, error) {
	query := `
        SELECT id, first_name, last_name, email, created_at
        FROM customers
        WHERE id = $1
    `

	var customer model.Customer
	err := r.db.conn.DB.QueryRow(query, id).Scan(
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("customer %s not found", id)
		return nil, nil
	}
	if err != nil {
		log.Printf("error fetching customer %s: %v", id, err)
		return nil, err
	}

	return &customer, nil
}
