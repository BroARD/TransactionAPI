package transactions

import (
	"TransactionAPI/internal/models"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, trans *models.Transaction) (*models.Transaction, error)
	GetTransactionsByCount(ctx context.Context, trans_count int) ([]models.Transaction, error) 
}