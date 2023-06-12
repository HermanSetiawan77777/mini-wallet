package httpserver

import (
	"herman-technical-julo/internal/app"
	authhttp "herman-technical-julo/internal/auth/http"
	"herman-technical-julo/internal/httpserver/response"
	"net/http"

	"github.com/gorilla/mux"
)

func buildRoutes(appContainer *app.Application) http.Handler {
	root := mux.NewRouter()

	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.WithData(w, http.StatusOK, "Server is running")
	}).Methods(http.MethodGet)

	root.HandleFunc("/api/v1/tokens", authhttp.HandleGetToken(appContainer.Services.AuthService)).Methods(http.MethodPost)

	return root
}
