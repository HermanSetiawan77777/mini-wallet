package authenticator

import (
	"context"
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/errors"
	"net/http"
	"strings"
)

type RequestAuthenticator interface {
	Validate(w http.ResponseWriter, r *http.Request) (context.Context, error)
}

type JULOWebAuthenticator struct {
	authService auth.AuthIService
}

func NewJULOWebAuthenticator(authService auth.AuthIService) *JULOWebAuthenticator {
	return &JULOWebAuthenticator{authService}
}

func (a *JULOWebAuthenticator) Validate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	authorization := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authorization) < 2 {
		return nil, errors.ErrInvalidSession
	}

	bearerToken := authorization[1]

	currentSession, isValid, err := a.authService.Validate(r.Context(), bearerToken)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.ErrInvalidSession
	}

	ctx := auth.InjectSessionToContext(r.Context(), currentSession)

	return ctx, nil
}
