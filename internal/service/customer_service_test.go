package service

import (
	"errors"
	"github.com/g-stro/tech-task/internal/domain/model"
	"reflect"
	"testing"
	"time"
)

var fixedTime = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

type MockCustomerRepository struct {
	customers map[string]*model.Customer
	err       error
}

func (m *MockCustomerRepository) GetByID(id string) (*model.Customer, error) {
	if m.err != nil {
		return nil, m.err
	}
	customer, exists := m.customers[id]
	if !exists {
		return nil, nil
	}
	return customer, nil
}

func TestCustomerService_GetCustomer(t *testing.T) {
	tests := []struct {
		name       string
		customerID string
		mockRepo   *MockCustomerRepository
		want       *model.Customer
		wantErr    bool
	}{
		{
			name:       "Success - Customer found",
			customerID: "6aa6cb6c-6054-4943-a0a7-f279cf6ceabd",
			mockRepo: &MockCustomerRepository{
				customers: map[string]*model.Customer{
					"6aa6cb6c-6054-4943-a0a7-f279cf6ceabd": {
						ID:        "6aa6cb6c-6054-4943-a0a7-f279cf6ceabd",
						FirstName: "John",
						LastName:  "Doe",
						Email:     "john.doe@example.com",
						CreatedAt: fixedTime,
					},
				},
			},
			want: &model.Customer{
				ID:        "6aa6cb6c-6054-4943-a0a7-f279cf6ceabd",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				CreatedAt: fixedTime,
			},
			wantErr: false,
		},
		{
			name:       "Fail - Customer not found",
			customerID: "xxxxxxxx-6054-4943-a0a7-f279cf6ceabd",
			mockRepo:   &MockCustomerRepository{},
			want:       nil,
			wantErr:    false,
		},
		{
			name:       "Fail - Repository error",
			customerID: "xxxxxxxx-6054-4943-a0a7-f279cf6ceabd",
			mockRepo: &MockCustomerRepository{
				err: errors.New("repository error"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CustomerService{
				repo: tt.mockRepo,
			}
			got, err := s.GetCustomer(tt.customerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
