package handlers

import (
	"net/http"
	"os"

	"github.com/kingtingthegreat/top-fetch/db"
	"github.com/kingtingthegreat/top-fetch/spotify"
)

func TrackHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no id provided"))
		return
	}

	user, err := db.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id. user not found."))
		return
	}

	track, newAccessToken, err := spotify.GetUserTopTrack(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"), user.AccessToken, user.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong. please try again."))
		return
	}

	if newAccessToken != "" {
		user.AccessToken = newAccessToken
		_, err := db.InsertUser(user)
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something went wrong. please try again."))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(track.Name + " - " + track.Artists[0].Name + "\x1d" + track.Album.Images[0].Url))
}
