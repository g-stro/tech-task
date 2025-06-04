package handler

import (
	"encoding/json"
	"github.com/g-stro/tech-task/internal/api/http/contract"
	"github.com/g-stro/tech-task/internal/domain/model"
	"github.com/g-stro/tech-task/internal/service"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type InvestmentHandler struct {
	svc *service.InvestmentService
}

func NewInvestmentHandler(svc *service.InvestmentService) *InvestmentHandler {
	return &InvestmentHandler{svc: svc}
}

func (h *InvestmentHandler) HandleInvestmentRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createInvestment(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// CreateInvestment creates a new investment with a single fund.
// Note: The fund amount will match the investment amount. This would need to be changed once multiple funds are supported.
func (h *InvestmentHandler) createInvestment(w http.ResponseWriter, r *http.Request) {
	var req contract.InvestmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error deconding JSON: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	investment := &model.Investment{
		ID:        uuid.New(),
		AccountID: req.AccountID,
		Funds:     []*model.Fund{{ID: req.FundID, Amount: req.Amount}}, // Only 1 fund for now
		Amount:    req.Amount,
		Status:    "PENDING", // Default
	}

	result, err := h.svc.ProcessInvestment(investment)
	if err != nil {
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	resp := contract.InvestmentResponse{
		ID: result.ID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
