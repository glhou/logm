package middleware

import (
	"net/http"

	"github.com/glhou/logm/internal/core"
	"go.uber.org/zap"
)

func LoggingMiddleware(app *core.App, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("Request", zap.String("method", r.Method), zap.String("url", r.URL.Path))
		next.ServeHTTP(w, r)
		app.Logger.Debug("Handled", zap.String("method", r.Method), zap.String("url", r.URL.Path))
	})
}
