package auth

import (
	"context"
	"herman-technical-julo/internal/errors"
)

type key string

var sessionKey key = "currentSession"

func InjectSessionToContext(ctx context.Context, session *Session) context.Context {
	ctx = context.WithValue(ctx, sessionKey, session)
	return ctx
}

// Extract the session from the given context.
//
// The returned session will always be present and valid if error is not nil
func GetSessionFromContext(ctx context.Context) (*Session, error) {
	session, ok := ctx.Value(sessionKey).(*Session)
	if !ok {
		return nil, errors.ErrInvalidSession
	}
	if session == nil {
		return nil, errors.ErrInvalidSession
	}
	if session.Validate() != nil {
		return nil, errors.ErrInvalidSession
	}

	return session, nil
}
