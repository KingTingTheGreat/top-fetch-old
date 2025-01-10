package handlers

import (
	"net/http"
	"os"

	"github.com/kingtingthegreat/top-fetch-old/db"
	"github.com/kingtingthegreat/top-fetch-old/spotify"
	"github.com/kingtingthegreat/top-fetch-old/tmplts"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}

	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectUri := os.Getenv("SPOTIFY_REDIRECT_URI")

	accessToken, refreshToken, err := spotify.ExchangeCode(clientId, clientSecret, redirectUri, code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}

	spotifyId, newAccessToken, err := spotify.GetUserProfile(clientId, clientSecret, accessToken, refreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}
	if newAccessToken != "" {
		accessToken = newAccessToken
	}

	// recover existing id
	user, err := db.GetUserBySpotifyId(spotifyId)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		tmplts.LayoutComponent(tmplts.Callback(user.Id), "Top Fetch").Render(r.Context(), w)
		return
	}

	// new user, create a new id for them
	user = db.DBUser{
		SpotifyId:    spotifyId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	id, err := db.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmplts.LayoutString("something went wrong. please try again.", "Internal Server Error").Render(r.Context(), w)
		return
	}

	tmplts.LayoutComponent(tmplts.Callback(id), "Top Fetch").Render(r.Context(), w)
	return
}
