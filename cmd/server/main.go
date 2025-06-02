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
	// ProcessInvestment DB connection
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to establish database connection %v", err)
	}
	defer db.Close()

	// ProcessInvestment repositories
	investmentRepo := postgres.NewInvestmentRepository(db)
	accountRepo := postgres.NewAccountRepository(db)
	fundRepo := postgres.NewFundRepository(db)

	// ProcessInvestment services
	investmentService := service.NewInvestmentService(investmentRepo, accountRepo, fundRepo)

	// ProcessInvestment handlers
	investmentHandler := handler.NewInvestmentHandler(investmentService)

	// ProcessInvestment multiplexer
	mux := http.NewServeMux()

	// Register routes
	handlers := routes.NewHandlers(investmentHandler)
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
