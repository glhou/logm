package api

import (
	"net/http"

	"github.com/glhou/logm/internal/core"
)

func RegisterRoutes(mux *http.ServeMux, app *core.App) {
	mux.HandleFunc("/api/log", newLogHandler(app))
	mux.HandleFunc("/api/logs", listLogHandler(app))
}
