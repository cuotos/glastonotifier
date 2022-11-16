package main

import (
	"context"
)

type Repository interface {
	Store(context.Context, string, Band) error
	Get(context.Context, string) (Band, error)
}
