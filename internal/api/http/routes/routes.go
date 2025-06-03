package routes

import (
	"github.com/g-stro/tech-task/internal/api/http/handler"
	"net/http"
)

type Handlers struct {
	investment *handler.InvestmentHandler
	reporting  *handler.ReportingHandler
}

func NewHandlers(
	investmentHandler *handler.InvestmentHandler,
	reportingHandler *handler.ReportingHandler,
) *Handlers {
	return &Handlers{
		investment: investmentHandler,
		reporting:  reportingHandler,
	}
}

func RegisterRoutes(mux *http.ServeMux, handlers *Handlers) {
	mux.HandleFunc("/investments/", handlers.investment.HandleInvestmentRequest)
	mux.HandleFunc("/accounts/{id}/investments/", handlers.reporting.HandleReportingRequest)
}
