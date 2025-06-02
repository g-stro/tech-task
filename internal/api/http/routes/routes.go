package routes

import (
	"github.com/g-stro/tech-task/internal/api/http/handler"
	"net/http"
)

type Handlers struct {
	investment *handler.InvestmentHandler
	// ...
}

func NewHandlers(
	customerHandler *handler.InvestmentHandler,
	// ...
) *Handlers {
	return &Handlers{
		investment: customerHandler,
		// ...
	}
}

func RegisterRoutes(mux *http.ServeMux, handlers *Handlers) {
	// Customer routes
	mux.HandleFunc("/investments/", handlers.investment.HandleInvestmentRequest)
	// ...
}
