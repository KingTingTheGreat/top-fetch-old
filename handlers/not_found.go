package handlers

import (
	"github.com/kingtingthegreat/top-fetch/tmplts"
	"net/http"

	"github.com/a-h/templ"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templ.Handler(tmplts.LayoutString("Page Not Found", "404")).ServeHTTP(w, r)
}
