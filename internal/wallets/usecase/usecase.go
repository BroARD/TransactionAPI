package usecase

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/wallets"
	"TransactionAPI/pkg/logging"
	"context"
	"fmt"
)

const (
	basePrefix = "api-wallet:"
	cacheDuration = 3600
)

type walletUseCase struct {
	repo wallets.Repository
	redisRepo wallets.RedisRepository
	logger logging.Logger
}

func NewWalletUseCase(repo wallets.Repository, logger logging.Logger, redisRepo wallets.RedisRepository) wallets.UseCase {
	return &walletUseCase{repo: repo, logger: logger, redisRepo: redisRepo}
}

// Create implements wallets.UseCase.
func (u *walletUseCase) Create(ctx context.Context, wallet models.Wallet) (models.Wallet, error) {
	return u.repo.Create(ctx, wallet)
}

// GetWalletByID implements wallets.UseCase.
func (u *walletUseCase) GetWalletByID(ctx context.Context, wallet_id string) (models.Wallet, error) {
	walletBase, err := u.redisRepo.GetWalletByIDCtx(ctx, u.getKeyWithPrefix(wallet_id))
	if err != nil {
		u.logger.Errorf("newsUC.GetNewsByID.GetNewsByIDCtx: %v", err)
	}
	if walletBase != nil {
		return *walletBase, nil
	}

	w, err := u.repo.GetWalletByID(ctx, wallet_id)
	if err != nil {
		return models.Wallet{}, err
	}

	if err = u.redisRepo.SetWalletCtx(ctx, u.getKeyWithPrefix(wallet_id), cacheDuration, &w); err != nil {
		u.logger.Errorf("newsUC.GetNewsByID.SetNewsCtx: %s", err)
	}

	return w, nil
}

func (u *walletUseCase) getKeyWithPrefix(key string) string{
	return fmt.Sprintf("%s:%s", basePrefix, key)
}

