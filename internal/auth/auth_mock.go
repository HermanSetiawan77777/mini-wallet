package auth

import "context"

type AuthServiceMock struct {
	AuthenticateFunc func(ctx context.Context, customerXid string) (token string, err error)
	ValidateFunc     func(ctx context.Context, token string) (session *Session, isValid bool, err error)
}

func (m *AuthServiceMock) Authenticate(ctx context.Context, customerXid string) (token string, err error) {
	if m.AuthenticateFunc != nil {
		return m.AuthenticateFunc(ctx, customerXid)
	}

	return "token", nil
}

func (m *AuthServiceMock) Validate(ctx context.Context, token string) (session *Session, isValid bool, err error) {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(ctx, token)
	}

	return &Session{}, true, nil
}
