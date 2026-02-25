package kv

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type UpstashStoreClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewUpstashStoreClient(cfg *Config) (*UpstashStoreClient, error) {
	opt, err := redis.ParseURL(cfg.StoreURL)
	if err != nil {
		return nil, err
	}

	return &UpstashStoreClient{
		client: redis.NewClient(opt),
		ctx:    context.Background(),
	}, nil
}

func (u *UpstashStoreClient) Set(key string, value string, ttlSeconds int) error {
	return u.client.Set(u.ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func (u *UpstashStoreClient) Get(key string) (string, error) {
	return u.client.Get(u.ctx, key).Result()
}

func (u *UpstashStoreClient) Delete(key string) error {
	return u.client.Del(u.ctx, key).Err()
}

func (u *UpstashStoreClient) Exists(key string) (int64, error) {
	return u.client.Exists(u.ctx, key).Result()
}

func (u *UpstashStoreClient) Ping() error {
	return u.client.Ping(u.ctx).Err()
}
