package usecase

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/transactions"
	"TransactionAPI/pkg/logging"
	"context"
)

type transUC struct {
	transRepo transactions.Repository
	logger    logging.Logger
}

func NewTransUseCase(transRepo transactions.Repository, logger logging.Logger) transactions.UseCase {
	return &transUC{transRepo: transRepo, logger: logger}
}

func (u *transUC) Create(ctx context.Context, trans *models.Transaction) (*models.Transaction, error) {
	return u.transRepo.Create(ctx, trans)
}

func (u *transUC) GetTransactionsByCount(ctx context.Context, trans_count int) ([]models.Transaction, error) {
	return u.transRepo.GetTransactionsByCount(ctx, trans_count)
}
