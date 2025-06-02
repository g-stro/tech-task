package service

import (
	"fmt"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/google/uuid"
	"reflect"
	"strings"
	"testing"
)

func TestReportingService_GetInvestmentsByAccountID(t *testing.T) {
	tests := []struct {
		name               string
		id                 uuid.UUID
		mockInvestmentRepo *MockInvestmentRepository
		mockAccountRepo    *MockAccountRepository
		mockFundRepo       *MockFundRepository
		want               []*InvestmentDetails
		wantErr            bool
	}{
		{
			name:               "Success - Investments retrieved for account",
			id:                 fixedAccountID,
			mockInvestmentRepo: NewMockInvestmentRepository(),
			mockAccountRepo:    NewMockAccountRepository(),
			mockFundRepo:       NewMockFundRepository(),
			want: []*InvestmentDetails{
				{
					ID:        fixedInvestmentID,
					Amount:    20000.00,
					Status:    statusPending,
					CreatedAt: fixedTime,
					UpdatedAt: fixedTime,
					Fund: &model.Fund{
						ID:         fixedFundID,
						Name:       fundNameCushonEquity,
						Category:   fundTypeEQUITY,
						Currency:   fundCurrencyGBP,
						RiskReturn: fundRiskLOW,
						CreatedAt:  fixedTime,
						UpdatedAt:  fixedTime,
					},
				},
			},
			wantErr: false,
		},
		{
			name:               "Fail - Account not found",
			id:                 fixedAccountID,
			mockInvestmentRepo: NewMockInvestmentRepository(),
			mockAccountRepo:    &MockAccountRepository{},
			mockFundRepo:       NewMockFundRepository(),
			want:               nil,
			wantErr:            true,
		},
		{
			name:               "Investments not found for account",
			id:                 fixedAccountID,
			mockInvestmentRepo: &MockInvestmentRepository{},
			mockAccountRepo:    NewMockAccountRepository(),
			mockFundRepo:       &MockFundRepository{},
			want:               []*InvestmentDetails{},
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ReportingService{
				investmentRepo: tt.mockInvestmentRepo,
				accountRepo:    tt.mockAccountRepo,
				fundRepo:       tt.mockFundRepo,
			}
			got, err := s.GetInvestmentsByAccountID(fixedAccountID)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetInvestmentsByAccountID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInvestmentsByAccountID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func printSlice[T any](items []*T) string {
	var builder strings.Builder
	for _, item := range items {
		builder.WriteString(fmt.Sprintf("%+v\n", *item))
	}
	return builder.String()
}
