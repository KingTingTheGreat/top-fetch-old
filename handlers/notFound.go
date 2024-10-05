package handlers

import (
	"net/http"
	"github.com/kingtingthegreat/top-fetch/tmplts"

	"github.com/a-h/templ"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templ.Handler(tmplts.LayoutString("Page Not Found", "404")).ServeHTTP(w, r)
}
