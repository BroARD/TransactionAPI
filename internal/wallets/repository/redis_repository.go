package repository

import (
	"TransactionAPI/internal/models"
	"TransactionAPI/internal/wallets"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type walletRedisRepository struct {
	redisClient *redis.Client
}

func NewWalletRedisRepository(redisClient *redis.Client) wallets.RedisRepository {
	return &walletRedisRepository{redisClient: redisClient}
}

func (w *walletRedisRepository) DeleteWalletCtx(ctx context.Context, key string) error {
	if err := w.redisClient.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

func (w *walletRedisRepository) GetWalletByIDCtx(ctx context.Context, key string) (*models.Wallet, error) {
	walletByte, err := w.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	walletBase := &models.Wallet{}
	if err := json.Unmarshal(walletByte, walletBase); err != nil {
		return nil, err
	}

	return walletBase, nil
}


func (w *walletRedisRepository) SetWalletCtx(ctx context.Context, key string, seconds int, wallet *models.Wallet) error {
	transBytes, err := json.Marshal(wallet)
	if err != nil {
		return err
	}
	if err = w.redisClient.Set(ctx, key, transBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}

	return nil
}

