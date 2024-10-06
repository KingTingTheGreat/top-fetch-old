package handlers

import (
	"net/http"

	"github.com/kingtingthegreat/top-fetch/tmplts"
)

func DocumentationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tmplts.LayoutComponent(tmplts.Documentation(), "Documentation").Render(r.Context(), w)
}
