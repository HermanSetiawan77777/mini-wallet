package mux

import (
	"herman-technical-julo/internal/httpserver/middleware/authenticator"
	"herman-technical-julo/internal/httpserver/response"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthenticateRequest(authenticator authenticator.RequestAuthenticator) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, err := authenticator.Validate(w, r)
			if err != nil {
				response.WithError(w, err)
				return
			}

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
