package service

import (
	"errors"
	"fmt"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/g-stro/tech-task/internal/domain/repository"
	"github.com/google/uuid"
	"log"
)

const (
	maxAllowance   = 20000.00
	minInvestment  = 1.00
	accountTypeISA = "ISA"
	statusPending  = "PENDING"
)

type InvestmentService struct {
	investmentRepo repository.InvestmentRepository
	accountRepo    repository.AccountRepository
	fundRepo       repository.FundRepository
}

func NewInvestmentService(investRepo repository.InvestmentRepository, accRepo repository.AccountRepository, fundRepo repository.FundRepository) *InvestmentService {
	return &InvestmentService{investmentRepo: investRepo, accountRepo: accRepo, fundRepo: fundRepo}
}

// ProcessInvestment validates and persists the investment before generating an event for further processing by other services (not implemented)
func (s *InvestmentService) ProcessInvestment(investment *model.Investment) (*model.Investment, error) {
	// Validate the investments
	if err := s.validateInvestment(investment.AccountID, investment.FundID, investment.Amount); err != nil {
		log.Printf("error validating investment: %v", err)
		return nil, err
	}

	investment.Status = statusPending // 'PENDING' until processed
	// Save the investment
	savedInvestment, err := s.investmentRepo.Create(investment)
	if err != nil {
		return nil, err
	}

	// TODO - Send the investment for processing with other services (out of scope)

	return savedInvestment, nil
}

// validateInvestment performs validation for a new investment
func (s *InvestmentService) validateInvestment(accountID, fundID uuid.UUID, amount float64) error {
	// Check account exists and is active
	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("account not found")
	}

	if account.Status != "ACTIVE" {
		return errors.New("account is not active")
	}

	// Check account is an ISA
	if account.AccountType != accountTypeISA {
		return errors.New("account is not an ISA")
	}

	// Check if fund exists
	// Single fund for now, a future change would involve fetching and validating multiple funds
	fund, err := s.fundRepo.GetByID(fundID)
	if err != nil || fund == nil {
		return errors.New("fund not found")
	}

	// Validate amount
	if amount < minInvestment {
		return fmt.Errorf("amount must be greater than %v", minInvestment)
	}

	// Very naive; assumes the customer has not made other deposits into any existing ISAs
	if amount > maxAllowance {
		return fmt.Errorf("amount must be less than %v", maxAllowance)
	}

	return nil
}
