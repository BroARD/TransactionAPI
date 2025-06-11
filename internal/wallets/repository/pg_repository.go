package repository

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/wallets"
	"context"

	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) wallets.Repository {
	return &walletRepository{db: db}
}

func (r *walletRepository) Create(ctx context.Context, wallet models.Wallet) (models.Wallet, error) {
	err := r.db.WithContext(ctx).Create(wallet).Error
	return wallet, err
}

func (r *walletRepository) GetWalletByID(ctx context.Context, wallet_id string) (models.Wallet, error) {
	var wallet models.Wallet
	err := r.db.WithContext(ctx).First(&wallet, "id = ?", wallet_id).Error
	return wallet, err
}

func (r *walletRepository) UpdateAmount(ctx context.Context, wallet models.Wallet, new_amount float64) error {
	err := r.db.WithContext(ctx).Model(wallet).Update("amount", new_amount).Error
	return err
}
