package token

import (
	"context"
	"time"
)

type Tokener[T any] interface {
	GenerateToken(ctx context.Context, payload T, expiryTime *time.Duration) (string, error)
	Validate(ctx context.Context, token string) (T, error)
}
