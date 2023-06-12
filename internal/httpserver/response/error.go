package response

import (
	"herman-technical-julo/internal/errors"
	"net/http"
)

func WithError(w http.ResponseWriter, err error) {
	handledError := defineError(err)
	withJSON(w, handledError.HttpStatus, map[string]any{
		"error": handledError,
	})
}

func defineError(err error) errors.HandledError {
	var definedError errors.HandledError
	switch err := err.(type) {
	case errors.HandledError:
		definedError = err
	default:
		definedError = errors.NewDefaultHandledError(http.StatusInternalServerError, err.Error())
	}

	return definedError
}
