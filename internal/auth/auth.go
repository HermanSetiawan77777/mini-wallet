package auth

import (
	"context"
	"herman-technical-julo/internal/errors"
)

type Session struct {
	WalletId string
}

func (s *Session) Validate() error {
	if s.WalletId == "" {
		return errors.ErrInvalidSession
	}

	return nil
}

type AuthIService interface {
	Validate(ctx context.Context, token string) (session *Session, isValid bool, err error)
}

type AuthService struct {
}
