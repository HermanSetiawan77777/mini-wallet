package http

import (
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/errors"
	"herman-technical-julo/internal/httpserver/request"
	"herman-technical-julo/internal/httpserver/response"
	"net/http"
)

type GetTokenRequest struct {
	CustomerXid string `json:"customer_xid"`
}

type GetTokenResponse struct {
	Token string `json:"token"`
}

func HandleGetToken(service auth.AuthIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody := &GetTokenRequest{}

		err := request.DecodeBody(r, &requestBody)
		if err != nil {
			if err == errors.ErrEmptyPayload {
				response.WithError(w, err)
				return
			}
			response.WithError(w, errors.ErrUnprocessablePayload)
			return
		}

		token, err := service.Authenticate(r.Context(), requestBody.CustomerXid)
		if err != nil {
			response.WithError(w, err)
			return
		}
		response.WithData(w, http.StatusOK, &GetTokenResponse{
			Token: token,
		}, "success")
	}
}
