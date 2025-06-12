package transactions

import (
	"TransactionAPI/internal/models"
	"context"
)

type RedisRepository interface {
	GetTransacionByCountCtx(ctx context.Context, key string) (*models.Transaction, error)
	SetTransactionCtx(ctx context.Context, key string, seconds int, trans models.Transaction) error
	DeletTransCtx(ctx context.Context, key string) error
}
