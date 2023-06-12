package errors

import "net/http"

func NewUnauthorizedError(code, message, detail string) HandledError {
	return NewHandledError(http.StatusUnauthorized, code, message, detail)
}

var ErrInvalidSession = NewUnauthorizedError("COMM0001", "Unauthorized", "invalid session")
