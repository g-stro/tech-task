package handler

import (
	"encoding/json"
	"github.com/g-stro/tech-task/internal/api/http/contract"
	"github.com/g-stro/tech-task/internal/service"
	"net/http"
	"strings"
)

type CustomerHandler struct {
	svc *service.CustomerService
}

func NewCustomerHandler(svc *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{svc: svc}
}

func (h *CustomerHandler) HandleCustomerRequest(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	if id == "" {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getCustomer(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request, id string) {
	customer, err := h.svc.GetCustomer(id)
	if err != nil {
		http.Error(w, "failed to get customer", http.StatusInternalServerError)
		return
	}

	if customer == nil {
		http.Error(w, "customer not found", http.StatusNotFound)
		return
	}

	resp := &contract.CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}

	// Respond with customer data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
