package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRespository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		redis: client,
	}
}

func (r *RedisRepository) Store(ctx context.Context, key string, b Band) error {
	bandBytes, err := json.Marshal(b)

	r.redis.Set(ctx, key, bandBytes, 0)

	return err
}

func (r *RedisRepository) Get(ctx context.Context, key string) (Band, error) {

	var b Band

	val, err := r.redis.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return b, fmt.Errorf("redis key not found %s", key)
	case err != nil:
		return b, fmt.Errorf("redis get failed: %w", err)
	case val == "":
		return b, fmt.Errorf("redis key found but no value %s", key)
	}

	if err := json.Unmarshal([]byte(val), &b); err != nil {
		return b, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return b, nil
}
