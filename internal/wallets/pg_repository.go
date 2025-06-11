package wallets

import (
	"TransactionAPI/internal/models"
	"context"
)

type Repository interface {
	Create(ctx context.Context, wallet models.Wallet) (models.Wallet, error)
	GetWalletByID(ctx context.Context, wallet_id string) (models.Wallet, error)
	UpdateAmount(ctx context.Context, wallet models.Wallet, new_amount float64) error
}