package web

import (
	"net/http"

	"github.com/glhou/logm/internal/core"
	"github.com/glhou/logm/views"
)

func dashboardHandler(a *core.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		views.Dashboard().Render(r.Context(), w)
	}
}

func statsHandler(a *core.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
<div class="card">
	<p> Total logs: 0 </p>
	<p> Errors: 0 </p>
</div>
		`))
	}
}
