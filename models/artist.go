package models

import "time"

type Artist struct {
	ArtistId             int       `json:"artistId"`
	SpotifyId            string    `json:"spotifyId"`
	Name                 string    `json:"name"`
	Popularity           int       `json:"popularity"`
	Albums               []Album   `json:"albums"`
	Singles              []Album   `json:"singles"`
	Compilations         []Album   `json:"compilations"`
	Genres               []string  `json:"genres"`
	Url                  string    `json:"url"`
	Followers            int       `json:"followers"`
	Image                string    `json:"image"`
	MaxDanceability      float32   `json:"maxDanceability"`
	MaxDanceabilityTrack string    `json:"maxDanceabilityTrack"`
	MaxEnergy            float32   `json:"maxEnergy"`
	MaxEnergyTrack       string    `json:"maxEnergyTrack"`
	AlbumsCount          int       `json:"albumsCount"`
	SinglesCount         int       `json:"singlesCount"`
	CompilationsCount    int       `json:"compilationsCount"`
	TracksCount          int       `json:"tracksCount"`
	DurationMs           int       `json:"durationMs"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

type Album struct {
	AlbumId              int       `json:"albumId"`
	SpotifyId            string    `json:"spotifyId"`
	Name                 string    `json:"name"`
	Type                 string    `json:"type"`
	ReleaseDate          string    `json:"releaseDate"`
	Url                  string    `json:"url"`
	Image                string    `json:"image"`
	Tracks               []Track   `json:"tracks"`
	MaxDanceability      float32   `json:"maxDanceability"`
	MaxDanceabilityTrack string    `json:"maxDanceabilityTrack"`
	MaxEnergy            float32   `json:"maxEnergy"`
	MaxEnergyTrack       string    `json:"maxEnergyTrack"`
	TracksCount          int       `json:"tracksCount"`
	DurationMs           int       `json:"durationMs"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

type Track struct {
	TrackId          int       `json:"trackId"`
	SpotifyId        string    `json:"spotifyId"`
	Name             string    `json:"name"`
	Url              string    `json:"url"`
	Danceability     float32   `json:"danceability"`
	Energy           float32   `json:"energy"`
	Acousticness     float32   `json:"acousticness"`
	Loudness         float32   `json:"loudness"`
	Liveness         float32   `json:"liveness"`
	Instrumentalness float32   `json:"instrumentalness"`
	Tempo            float64   `json:"tempo"`
	DurationMs       int       `json:"durationMs"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
