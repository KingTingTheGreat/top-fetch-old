package router

import (
	"net/http"
	"os"

	"github.com/kingtingthegreat/top-fetch/handlers"
	"github.com/kingtingthegreat/top-fetch/spotify"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	if os.Getenv("ENVIRONMENT") == "dev" {
		fileServer := http.FileServer(http.Dir("./public"))
		router.Handle("/public/", http.StripPrefix("/public/", fileServer))
	}

	router.HandleFunc("/", handlers.HomePageHandler)

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.HandleFunc("GET /sign-in", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, spotify.AuthUrl(), http.StatusSeeOther)
	})
	router.HandleFunc("GET /callback", handlers.CallbackHandler)

	router.HandleFunc("GET /track", handlers.TrackHandler)

	router.HandleFunc("/404", handlers.NotFoundHandler)

	return router
}
