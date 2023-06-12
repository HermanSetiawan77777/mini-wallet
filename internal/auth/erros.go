package auth

import "herman-technical-julo/internal/errors"

var ErrCredentialsInvalid = errors.NewValidationError("ErrCredentialsInvalid", "Customer xid is invalid", "invalid credentials")
