package repository

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/transactions"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type transRedisRepo struct {
	redisClient *redis.Client
}

func NewTransRedisRepo(redisClient *redis.Client) transactions.RedisRepository {
	return &transRedisRepo{redisClient: redisClient}
}

// DeletTransCtx implements transactions.RedisRepository.
func (t *transRedisRepo) DeletTransCtx(ctx context.Context, key string) error {
	if err := t.redisClient.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

// GetTransacionByCountCtx implements transactions.RedisRepository.
func (t *transRedisRepo) GetTransacionByCountCtx(ctx context.Context, key string) (*models.Transaction, error) {
	transByte, err := t.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	transBase := &models.Transaction{}
	if err := json.Unmarshal(transByte, transBase); err != nil {
		return nil, err
	}

	return transBase, nil
}

// SetTransactionCtx implements transactions.RedisRepository.
func (t *transRedisRepo) SetTransactionCtx(ctx context.Context, key string, seconds int, trans models.Transaction) error {
	transBytes, err := json.Marshal(trans)
	if err != nil {
		return err
	}
	if err = t.redisClient.Set(ctx, key, transBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}

	return nil
}

