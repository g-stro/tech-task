package routes

import (
	"github.com/g-stro/tech-task/internal/api/http/handler"
	"net/http"
)

type Handlers struct {
	Customer *handler.CustomerHandler
	// ...
}

func NewHandlers(
	customerHandler *handler.CustomerHandler,
	// ...
) *Handlers {
	return &Handlers{
		Customer: customerHandler,
		// ...
	}
}

func RegisterRoutes(mux *http.ServeMux, handlers *Handlers) {
	// Customer routes
	mux.HandleFunc("/customers/", handlers.Customer.HandleCustomerRequest)
	// ...
}
