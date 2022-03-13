package main

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestSand(t *testing.T) {
	b := Band{Name: "The Band", Stage: "A Stage", State: StronglyRumoured}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	r := NewRedisRepository(client)

	err := r.Set("band1", b, time.Duration(time.Minute))
	assert.NoError(t, err)
}
