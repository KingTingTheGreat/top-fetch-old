package router

import (
	"net/http"
	"os"

	"github.com/kingtingthegreat/top-fetch/handlers"
	"github.com/kingtingthegreat/top-fetch/spotify"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	// router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", handlers.HomePageHandler)
	router.HandleFunc("/documentation", handlers.DocumentationHandler)

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.HandleFunc("GET /sign-in", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, spotify.AuthUrl(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_REDIRECT_URI")), http.StatusSeeOther)
	})
	router.HandleFunc("GET /callback", handlers.CallbackHandler)

	router.HandleFunc("GET /track", handlers.TrackHandler)

	router.HandleFunc("/404", handlers.NotFoundHandler)

	return router
}
