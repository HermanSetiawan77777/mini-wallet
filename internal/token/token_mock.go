package token

import (
	"context"
	"time"
)

type TokenerMock[T any] struct {
	GenerateTokenFunc func(ctx context.Context, payload T, expiryTime *time.Duration) (string, error)
	ValidateFunc      func(ctx context.Context, token string) (T, error)
}

func (m *TokenerMock[T]) GenerateToken(ctx context.Context, payload T, expiryTime *time.Duration) (string, error) {
	if m.GenerateTokenFunc != nil {
		return m.GenerateTokenFunc(ctx, payload, expiryTime)
	}

	return "token", nil
}

func (m *TokenerMock[T]) Validate(ctx context.Context, token string) (T, error) {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(ctx, token)
	}

	var t T
	return t, nil
}
