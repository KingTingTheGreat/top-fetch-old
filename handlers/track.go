package handlers

import (
	"net/http"

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

	user, err := db.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id. user not found."))
		return
	}

	track, newAccessToken, err := spotify.GetUserTopTrack(user.AccessToken, user.RefreshToken)
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
