package wallets

import (
	"TransactionAPI/internal/models"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, wallet models.Wallet) (models.Wallet, error)
	GetWalletByID(ctx context.Context, wallet_id string) (models.Wallet, error)
}