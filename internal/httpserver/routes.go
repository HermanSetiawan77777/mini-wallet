package httpserver

import (
	"herman-technical-julo/internal/app"
	"net/http"

	"github.com/gorilla/mux"
)

func buildRoutes(appContainer *app.Application) http.Handler {
	root := mux.NewRouter()

	return root
}
