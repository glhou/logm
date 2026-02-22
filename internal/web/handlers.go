package web

import (
	"net/http"

	"github.com/glhou/logm/views"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	views.Dashboard().Render(r.Context(), w)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<div class="card">
	<p> Total logs: 0 </p>
	<p> Errors: 0 </p>
</div>
		`))
}
