package main

import (
	"log"
	"net/http"

	"github.com/glhou/logm/internal/core"
	"github.com/glhou/logm/internal/web"
	"github.com/glhou/logm/internal/web/middleware"
)

func main() {
	mux := http.NewServeMux()

	app := core.GetApp()

	web.RegisterRoutes(mux, app)

	wrappedMux := middleware.LoggingMiddleware(app, mux)

	app.Logger.Info("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
