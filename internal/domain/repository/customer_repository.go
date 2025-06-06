package repository

import (
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	GetByID(id uuid.UUID) (*model.Customer, error)
}
