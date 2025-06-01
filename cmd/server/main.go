package main

import (
	"github.com/g-stro/tech-task/internal/api/http/handler"
	"github.com/g-stro/tech-task/internal/api/http/routes"
	"github.com/g-stro/tech-task/internal/repository/postgres"
	"github.com/g-stro/tech-task/internal/service"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create DB connection
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to establish database connection %v", err)
	}
	defer db.Close()

	// Create repositories
	customerRepo := postgres.NewCustomerRepository(db)

	// Create services
	customerService := service.NewCustomerService(customerRepo)

	// Create handlers
	customerHandler := handler.NewCustomerHandler(customerService)

	// Create multiplexer
	mux := http.NewServeMux()

	// Register routes
	handlers := routes.NewHandlers(customerHandler)
	routes.RegisterRoutes(mux, handlers)

	server := &http.Server{
		Addr:    ":" + os.Getenv("SERVICE_PORT"),
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
