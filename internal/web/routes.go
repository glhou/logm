package web

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", dashboardHandler)
	mux.HandleFunc("/partials/stats", statsHandler)
}
