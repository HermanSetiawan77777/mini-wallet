package token

import "herman-technical-julo/internal/errors"

var ErrTokenInvalid = errors.NewUnauthorizedError("ErrTokenInvalid", "Token is invalid", "invalid token")
