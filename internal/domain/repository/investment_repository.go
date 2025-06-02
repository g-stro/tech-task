package repository

import (
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
)

type InvestmentRepository interface {
	Create(*model.Investment) (*model.Investment, error)
	GetByAccountID(id uuid.UUID) ([]*model.Investment, error)
}
