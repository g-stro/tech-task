package service

import (
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
	"time"
)

var fixedTime = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
var fixedCustomerID = uuid.MustParse("6aa6cb6c-6054-4943-a0a7-f279cf6ceabd")
var fixedInvalidCustomerID = uuid.MustParse("12345678-6054-4943-a0a7-f279cf6ceabd")
var fixedAccountID = uuid.MustParse("d7ee4877-7645-461a-b2cc-2f2c8f6a7284")
var fixedFundID = uuid.MustParse("cb91e975-d8bc-423b-bc99-fa6f396c2eaf")
var fixedInvestmentID = uuid.MustParse("f9d7338c-62e9-4f60-8d3a-9082f1616a23")

const (
	accountNumber         = "1234567890"
	accountStatusACTIVE   = "ACTIVE"
	accountStatusINACTIVE = "INACTIVE"
	accountTypeCashISA    = "CASH_ISA"

	fundNameCushonEquity = "Cushon Equities Fund"
	fundTypeEQUITY       = "EQUITY"
	fundCurrencyGBP      = "GBP"
	fundRiskLOW          = "LOW"
)

// Mocks and helpers
type MockAccountRepository struct {
	accounts map[uuid.UUID]*model.Account
	err      error
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		accounts: map[uuid.UUID]*model.Account{
			fixedAccountID: StandardAccount(),
		},
	}
}

func StandardAccount() *model.Account {
	return &model.Account{
		ID:            fixedAccountID,
		CustomerID:    fixedCustomerID,
		AccountType:   accountTypeISA,
		AccountNumber: accountNumber,
		Status:        accountStatusACTIVE,
		CreatedAt:     fixedTime,
		UpdatedAt:     fixedTime,
	}
}

func (m *MockAccountRepository) GetByID(id uuid.UUID) (*model.Account, error) {
	if m.err != nil {
		return nil, m.err
	}
	account, exists := m.accounts[id]
	if !exists {
		return nil, nil
	}
	return account, nil
}

type MockFundRepository struct {
	funds map[uuid.UUID]*model.Fund
	err   error
}

func NewMockFundRepository() *MockFundRepository {
	return &MockFundRepository{
		funds: map[uuid.UUID]*model.Fund{
			fixedFundID: StandardFund(),
		},
	}
}

func StandardFund() *model.Fund {
	return &model.Fund{
		ID:         fixedFundID,
		Name:       fundNameCushonEquity,
		Category:   fundTypeEQUITY,
		Currency:   fundCurrencyGBP,
		RiskReturn: fundRiskLOW,
		CreatedAt:  fixedTime,
		UpdatedAt:  fixedTime,
	}
}

func (m *MockFundRepository) GetByID(id uuid.UUID) (*model.Fund, error) {
	if m.err != nil {
		return nil, m.err
	}
	fund, exists := m.funds[id]
	if !exists {
		return nil, nil
	}
	return fund, nil
}
