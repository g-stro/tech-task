package repository

import (
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
)

type FundRepository interface {
	GetByID(id uuid.UUID) (*model.Fund, error)
}
