package repository

import (
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
)

type AccountRepository interface {
	GetByID(id uuid.UUID) (*model.Account, error)
}
