package errors

import "net/http"

func NewUnauthorizedError(code, message, detail string) HandledError {
	return NewHandledError(http.StatusUnauthorized, code, message, detail)
}
func NewValidationError(code, message, detail string) HandledError {
	return NewHandledError(http.StatusBadRequest, code, message, detail)
}
func NewUnprocessableEntity(code, message, detail string) HandledError {
	return NewHandledError(http.StatusUnprocessableEntity, code, message, detail)
}
func NewDefaultValidationError(message, detail string) HandledError {
	return NewHandledError(http.StatusBadRequest, "VLD0001", message, detail)
}

var ErrInvalidSession = NewUnauthorizedError("COMM0001", "Unauthorized", "invalid session")
var ErrEmptyPayload = NewValidationError("COMM0002", "Please fill data", "request body is empty")
var ErrUnprocessablePayload = NewUnprocessableEntity("COMM0003", "Data submitted is invalid", "payload is unprocessable")
