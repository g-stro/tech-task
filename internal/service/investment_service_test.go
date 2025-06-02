package service

import (
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

type MockInvestmentRepository struct {
	investments       map[string][]*model.Investment
	createdInvestment *model.Investment
	err               error
}

func (m *MockInvestmentRepository) GetByAccountID(id string) ([]*model.Investment, error) {
	if m.err != nil {
		return nil, m.err
	}
	investments, exists := m.investments[id]
	if !exists {
		return nil, nil
	}
	return investments, nil
}

func (m *MockInvestmentRepository) Create(investment *model.Investment) (*model.Investment, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.createdInvestment = investment
	m.createdInvestment.ID = fixedInvestmentID
	return m.createdInvestment, nil
}

func TestInvestmentService_Create(t *testing.T) {
	tests := []struct {
		name               string
		investment         *model.Investment // change to DTO?
		mockInvestmentRepo *MockInvestmentRepository
		mockAccountRepo    *MockAccountRepository
		mockFundRepo       *MockFundRepository
		want               *model.Investment
		wantErr            bool
	}{
		{
			name: "Success - Investment created",
			investment: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    20000.00,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo:    NewMockAccountRepository(),
			mockFundRepo:       NewMockFundRepository(),
			want: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    20000.00,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			wantErr: false,
		},
		{
			name: "Fail - Amound exceeds allowance",
			investment: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    25000.00,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo:    NewMockAccountRepository(),
			mockFundRepo:       NewMockFundRepository(),
			want:               nil,
			wantErr:            true,
		},
		{
			name: "Fail - Amount less than minimum investment amount",
			investment: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    0.99,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo:    NewMockAccountRepository(),
			mockFundRepo:       NewMockFundRepository(),
			want:               nil,
			wantErr:            true,
		},
		{
			name: "Fail - Account not found",
			investment: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    20000.00,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo:    &MockAccountRepository{},
			mockFundRepo:       NewMockFundRepository(),
			want:               nil,
			wantErr:            true,
		},
		{
			name: "Fail - Account not ACTIVE",
			investment: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    20000.00,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo: &MockAccountRepository{
				accounts: map[uuid.UUID]*model.Account{
					fixedAccountID: {
						ID:            fixedAccountID,
						CustomerID:    fixedCustomerID,
						AccountType:   accountTypeISA,
						AccountNumber: accountNumber,
						Status:        accountStatusINACTIVE,
						CreatedAt:     fixedTime,
						UpdatedAt:     fixedTime,
					},
				},
			},
			mockFundRepo: NewMockFundRepository(),
			want:         nil,
			wantErr:      true,
		},
		{
			name: "Fail - Account is not an ISA",
			investment: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    25000.00,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo: &MockAccountRepository{
				accounts: map[uuid.UUID]*model.Account{
					fixedAccountID: {
						ID:            fixedAccountID,
						CustomerID:    fixedCustomerID,
						AccountType:   accountTypeCashISA,
						AccountNumber: accountNumber,
						Status:        accountStatusINACTIVE,
						CreatedAt:     fixedTime,
						UpdatedAt:     fixedTime,
					},
				},
			},
			mockFundRepo: NewMockFundRepository(),
			want:         nil,
			wantErr:      true,
		},
		{
			name: "Fail - Fund not found",
			investment: &model.Investment{
				ID:        fixedInvestmentID,
				AccountID: fixedAccountID,
				FundID:    fixedFundID,
				Amount:    25000.00,
				Status:    statusPending,
				CreatedAt: fixedTime,
				UpdatedAt: fixedTime,
			},
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo:    NewMockAccountRepository(),
			mockFundRepo:       &MockFundRepository{},
			want:               nil,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InvestmentService{
				investmentRepo: tt.mockInvestmentRepo,
				accountRepo:    tt.mockAccountRepo,
				fundRepo:       tt.mockFundRepo,
			}
			got, err := s.ProcessInvestment(tt.investment)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInvestment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessInvestment() got = %v, want %v", got, tt.want)
			}
		})
	}
}
