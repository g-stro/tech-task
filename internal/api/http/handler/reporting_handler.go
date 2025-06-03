package handler

import (
	"encoding/json"
	"github.com/g-stro/tech-task/internal/api/http/contract"
	"github.com/g-stro/tech-task/internal/service"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type ReportingHandler struct {
	svc *service.ReportingService
}

func NewReportingHandler(svc *service.ReportingService) *ReportingHandler {
	return &ReportingHandler{svc: svc}
}

func (h *ReportingHandler) HandleReportingRequest(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		log.Printf("")
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetInvestments(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetInvestments returns investment details for a specific account
func (h *ReportingHandler) GetInvestments(w http.ResponseWriter, r *http.Request, accID string) {
	id, err := uuid.Parse(accID)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	investmentDetails, err := h.svc.GetInvestmentsByAccountID(id)
	if err != nil {
		http.Error(w, "error retrieving investment details", http.StatusInternalServerError)
		return
	}

	if len(investmentDetails) == 0 {
		log.Printf("no investment details found for account %v", id)
		http.Error(w, "no investment details found", http.StatusNotFound)
		return
	}

	// Create a presentable response for the user
	resp := make([]*contract.InvestmentDetailsResponse, 0, len(investmentDetails))
	for _, details := range investmentDetails {
		resp = append(resp, &contract.InvestmentDetailsResponse{
			ID:        details.ID.String(),
			Amount:    details.Amount,
			Status:    details.Status,
			CreatedAt: details.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: details.UpdatedAt.Format("2006-01-02 15:04:05"),
			Fund:      *details.Fund,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
