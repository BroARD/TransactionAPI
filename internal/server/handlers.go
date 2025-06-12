package server

import (
	transHttp "TransactionAPI/internal/transactions/delivery/http"
	transRepo "TransactionAPI/internal/transactions/repository"
	transUseCase "TransactionAPI/internal/transactions/usecase"

	walletRepo "TransactionAPI/internal/wallets/repository"
	walletUseCase "TransactionAPI/internal/wallets/usecase"
	walletHttp "TransactionAPI/internal/wallets/delivery/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error{

	if s.db == nil {
		s.logger.Info("database connection is not initialized")
	}
	//TODO: Init repositories
	transactionsRepository := transRepo.NewTransRepository(s.db)
	walletsRepository := walletRepo.NewWalletRepository(s.db)
	walletRedisRepository := walletRepo.NewWalletRedisRepository(s.redisClient)
	//TODO: Init UseCases
	transactionUseCase := transUseCase.NewTransUseCase(transactionsRepository, s.logger, walletsRepository)
	walletUseCase := walletUseCase.NewWalletUseCase(walletsRepository, s.logger, walletRedisRepository)
	//TODO: Init handlers
	transactionHandlers := transHttp.NewTransactionHandlers(transactionUseCase, s.logger)
	walletHandlers := walletHttp.NewWalletHandlers(walletUseCase, s.logger)

	v1 := e.Group("/api")
	transGroup := v1.Group("")
	walletGroup := v1.Group("/wallet")

	transHttp.MapTransationsRoutes(transGroup, transactionHandlers)
	walletHttp.MapWalletRoutes(walletGroup, walletHandlers)

	routes := e.Routes()

	for _, route := range routes {
		s.logger.Infof("Method: %s, Path: %s\n", route.Method, route.Path)
	}
	return nil
}