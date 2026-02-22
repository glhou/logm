package web

import (
	"net/http"

	"github.com/glhou/logm/internal/core"
	"github.com/glhou/logm/internal/web/api"
)

func RegisterRoutes(mux *http.ServeMux, app *core.App) {
	mux.HandleFunc("/", dashboardHandler(app))
	mux.HandleFunc("/partials/stats", statsHandler(app))
	api.RegisterRoutes(mux, app)
}
