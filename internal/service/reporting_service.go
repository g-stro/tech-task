package service

import (
	"errors"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/g-stro/tech-task/internal/domain/repository"
	"github.com/google/uuid"
	"log"
	"time"
)

type ReportingService struct {
	investmentRepo repository.InvestmentRepository
	accountRepo    repository.AccountRepository
	fundRepo       repository.FundRepository
}

func NewReportingService(investRepo repository.InvestmentRepository, accRepo repository.AccountRepository, fundRepo repository.FundRepository) *ReportingService {
	return &ReportingService{investmentRepo: investRepo, accountRepo: accRepo, fundRepo: fundRepo}
}

type InvestmentDetails struct {
	ID        uuid.UUID   `json:"id"`
	Amount    float64     `json:"amount"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Fund      *model.Fund `json:"fund"`
}

func (s *ReportingService) GetInvestmentsByAccountID(accountID uuid.UUID) ([]*InvestmentDetails, error) {
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

	if len(investments) == 0 {
		return []*InvestmentDetails{}, nil
	}

	// Create a presentable response for the user
	result := make([]*InvestmentDetails, 0, len(investments))
	for _, investment := range investments {
		// Get associated fund
		fund, err := s.fundRepo.GetByID(investment.FundID)
		if err != nil {
			log.Printf("error retrieving fund for investment %v: %e", investment.ID, err)
			return nil, err
		}

		result = append(result, &InvestmentDetails{
			ID:        investment.ID,
			Amount:    investment.Amount,
			Status:    investment.Status,
			CreatedAt: investment.CreatedAt,
			UpdatedAt: investment.UpdatedAt,
			Fund:      fund,
		})
	}

	return result, nil
}
