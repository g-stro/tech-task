package service

import (
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/g-stro/tech-task/internal/domain/repository"
)

type CustomerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetCustomer(id string) (*model.Customer, error) {
	customer, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
