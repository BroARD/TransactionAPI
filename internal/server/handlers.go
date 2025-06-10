package server

import (
	transHttp "TransactionAPI/internal/transactions/delivery/http"
	"TransactionAPI/internal/transactions/repository"
	"TransactionAPI/internal/transactions/usecase"
	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error{

	if s.db == nil {
		s.logger.Info("database connection is not initialized")
	}
	//TODO: Init repositories
	transactionsRepository := repository.NewTransRepository(s.db)
	//TODO: Init useCases
	transactionUseCase := usecase.NewTransUseCase(transactionsRepository, s.logger)
	//TODO: Init handlers
	transactionHandlers := transHttp.NewTransactionHandlers(transactionUseCase, s.logger)
	v1 := e.Group("/api")
	transGroup := v1.Group("/")

	transHttp.MapTransationsRoutes(transGroup, transactionHandlers)
	return nil
}