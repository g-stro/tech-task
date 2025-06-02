package handler

import (
	"github.com/g-stro/tech-task/internal/service"

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

func (h *InvestmentHandler) createInvestment(w http.ResponseWriter, r *http.Request) {
	// TODO
}
