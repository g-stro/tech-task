package service

import (
	"errors"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/g-stro/tech-task/internal/domain/repository"
	"github.com/google/uuid"
	"log"
)

type ReportingService struct {
	investmentRepo repository.InvestmentRepository
	accountRepo    repository.AccountRepository
	fundRepo       repository.FundRepository
}

func NewReportingService(investRepo repository.InvestmentRepository, accRepo repository.AccountRepository, fundRepo repository.FundRepository) *ReportingService {
	return &ReportingService{investmentRepo: investRepo, accountRepo: accRepo, fundRepo: fundRepo}
}

// GetInvestmentsByAccountID retrieves all investments with funds for a given account
func (s *ReportingService) GetInvestmentsByAccountID(accountID uuid.UUID) ([]*model.Investment, error) {
	// Check account exists
	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		log.Printf("error retrieving account %v for investment details: %v", accountID, err)
		return nil, err
	}
	if account == nil {
		return nil, errors.New("account not found")
	}
	investments, err := s.investmentRepo.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}
	log.Printf("funds: %v", investments[0].Funds)

	if len(investments) == 0 {
		return []*model.Investment{}, nil
	}

	return investments, nil
}
