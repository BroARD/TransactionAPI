package wallets

import (
	"TransactionAPI/internal/models"
	"context"
)

type RedisRepository interface {
	GetWalletByIDCtx(ctx context.Context, key string) (*models.Wallet, error)
	SetWalletCtx(ctx context.Context, key string, seconds int, wallet *models.Wallet) error
	DeleteWalletCtx(ctx context.Context, key string) error
}
