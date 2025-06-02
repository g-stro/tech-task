package repository

import (
	"github.com/g-stro/tech-task/internal/domain/model"
)

type InvestmentRepository interface {
	Create(*model.Investment) (*model.Investment, error)
}
