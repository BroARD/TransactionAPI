package repository

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/transactions"
	"context"

	"gorm.io/gorm"
)

type transRepo struct {
	db *gorm.DB
}

func NewTransRepository(db *gorm.DB) transactions.Repository {
	return &transRepo{db: db}
}

func (r *transRepo) Create(ctx context.Context, trans *models.Transaction) (*models.Transaction, error) {
	err := r.db.WithContext(ctx).Create(trans).Error
	return trans, err
}

func (r *transRepo) GetTransactionsByCount(ctx context.Context, trans_count int) ([]models.Transaction, error) {
	var transList []models.Transaction
	err := r.db.WithContext(ctx).Order("created_at DESC").Limit(trans_count).Find(&transList).Error
	return transList, err
}
