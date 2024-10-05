package handlers

import (
	"net/http"

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

	accessToken, refreshToken, err := spotify.ExchangeCode(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong. please try again."))
		return
	}

	spotifyId, newAccessToken, err := spotify.GetUserProfile(accessToken, refreshToken)
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
