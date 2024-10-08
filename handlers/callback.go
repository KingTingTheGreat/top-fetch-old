package handlers

import (
	"net/http"
	"os"

	"github.com/kingtingthegreat/top-fetch/db"
	"github.com/kingtingthegreat/top-fetch/spotify"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no authorization code"))
		return
	}

	clientId := os.Getenv("SPOTIFY_CLIEN_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectUri := os.Getenv("SPOTIFY_REDIRECT_URI")

	accessToken, refreshToken, err := spotify.ExchangeCode(clientId, clientSecret, redirectUri, code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong. please try again."))
		return
	}

	spotifyId, newAccessToken, err := spotify.GetUserProfile(clientId, clientSecret, accessToken, refreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong. please try again."))
		return
	}
	if newAccessToken != "" {
		accessToken = newAccessToken
	}

	user := db.DBUser{
		SpotifyId:    spotifyId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	id, err := db.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong. please try again."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}
