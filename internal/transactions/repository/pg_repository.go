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
	if err != nil {
		return &models.Transaction{}, err
	}
	return trans, err
}

func (r *transRepo) GetTransactionsByCount(ctx context.Context, trans_count int) ([]models.Transaction, error) {
	var transList []models.Transaction
	err := r.db.Order("created_at DESC").Limit(trans_count).Find(&transList).Error
	if err != nil {
		return []models.Transaction{}, err
	}
	return transList, nil
}
