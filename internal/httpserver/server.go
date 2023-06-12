package httpserver

import (
	"fmt"
	"herman-technical-julo/internal/app"
	"net/http"
)

func NewServer(port string, appContainer *app.Application) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: buildRoutes(appContainer),
	}
}
