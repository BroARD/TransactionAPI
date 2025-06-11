package usecase

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/wallets"
	"TransactionAPI/pkg/logging"
	"context"
)

type walletUseCase struct {
	repo wallets.Repository
	logger logging.Logger
}

func NewWalletUseCase(repo wallets.Repository, logger logging.Logger) wallets.UseCase {
	return &walletUseCase{repo: repo, logger: logger}
}

// Create implements wallets.UseCase.
func (u *walletUseCase) Create(ctx context.Context, wallet models.Wallet) (models.Wallet, error) {
	return u.repo.Create(ctx, wallet)
}

// GetWalletByID implements wallets.UseCase.
func (u *walletUseCase) GetWalletByID(ctx context.Context, wallet_id string) (models.Wallet, error) {
	return u.repo.GetWalletByID(ctx, wallet_id)
}

