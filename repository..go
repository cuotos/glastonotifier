package main

import (
	"time"

	"github.com/go-redis/redis"
)

type Repository interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
}

type repository struct {
	Client redis.Cmdable
}

func NewRedisRepository(Client redis.Cmdable) Repository {
	return &repository{Client}
}

func (r *repository) Set(key string, value interface{}, exp time.Duration) error {
	return r.Client.Set(key, value, exp).Err()
}

func (r *repository) Get(key string) (string, error) {
	get := r.Client.Get(key)
	return get.Result()
}
