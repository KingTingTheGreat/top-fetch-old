package spotify

import (
	"encoding/json"
	"fmt"
	"log"
)

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}

type Artist struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Uri  string `json:"uri"`
}

type Album struct {
	Name        string   `json:"name"`
	Href        string   `json:"href"`
	Uri         string   `json:"uri"`
	Artists     []Artist `json:"artists"`
	ReleaseDate string   `json:"release_date"`
	TotalTracks int      `json:"total_tracks"`
	Images      []Image  `json:"images"`
}

type Track struct {
	Name        string   `json:"name"`
	Album       Album    `json:"album"`
	Artists     []Artist `json:"artists"`
	Uri         string   `json:"uri"`
	Href        string   `json:"href"`
	TrackNumber int      `json:"track_number"`
	Popularity  int      `json:"popularity"`
}

type ProfileResponse struct {
	DisplayName string `json:"display_name"`
	Href        string `json:"href"`
	SpotifyId   string `json:"id"`
}

// Takes in AccessToken and RefreshToken.
// Returns user's SpotifyId.
// If the provided AccessToken has expired, return the new AccessToken.
func GetUserProfile(clientId, clientSecret, accessToken, refreshToken string) (string, string, error) {
	body, newAccessToken, err := spotifyRequest(clientId, clientSecret, accessToken, refreshToken, "https://api.spotify.com/v1/me")
	if err != nil {
		return "", "", err
	}

	var userProfileRes ProfileResponse
	if err := json.Unmarshal(body, &userProfileRes); err != nil {
		return "", "", err
	}

	return userProfileRes.SpotifyId, newAccessToken, nil
}

type TopTracksResponse struct {
	Items []Track `json:"items"`
	Total int     `json:"total"`
	Limit int     `json:"limit"`
}

// Takes in AccessToken and RefreshToken.
// Returns user's top song of type Track.
// If the provided AccessToken has expired, return the new AccessToken.
func GetUserTopTrack(clientId, clientSecret, accessToken, refreshToken string) (Track, string, error) {
	log.Println("GetUserTopTrack()")
	body, newAccessToken, err := spotifyRequest(clientId, clientSecret, accessToken, refreshToken, "https://api.spotify.com/v1/me/top/tracks?time_range=short_term&limit=1")
	if err != nil {
		return Track{}, "", err
	}

	log.Println("unmarshalling json")
	var topTracksResponse TopTracksResponse
	if err := json.Unmarshal(body, &topTracksResponse); err != nil {
		return Track{}, "", err
	}

	if len(topTracksResponse.Items) == 0 {
		return Track{}, "", fmt.Errorf("no top track found")
	}

	return topTracksResponse.Items[0], newAccessToken, nil

}
