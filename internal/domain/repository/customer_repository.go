package repository

import (
	"github.com/g-stro/tech-task/internal/domain/model"
)

type CustomerRepository interface {
	GetByID(id string) (*model.Customer, error)
}
