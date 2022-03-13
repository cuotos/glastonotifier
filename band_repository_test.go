package main

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	key = "key"
	val = "value"
)
var (
	client *redis.Client
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	code := m.Run()
	os.Exit(code)
}

func TestSet(t *testing.T) {
	key = "key"
	b := Band{
		Name: "The Band",
	}

	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, b, exp).Return(redis.NewStatusResult("", nil))

	r := NewRedisRepository(mock)
	err := r.Set(key, b, exp)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(val, nil))

	r := NewRedisRepository(mock)
	res, err := r.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}