package request

import (
	"encoding/json"
	"herman-technical-julo/internal/errors"
	"io"
	"net/http"
)

func DecodeBody(r *http.Request, payload any) error {
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		if err == io.EOF {
			return errors.ErrEmptyPayload
		}

		return err
	}

	return nil
}
