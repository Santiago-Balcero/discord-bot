package models

import (
	"fmt"
	"strings"
)

type Artist struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Popularity           int      `json:"popularity"`
	Albums               []Album  `json:"albums"`
	Genres               []string `json:"genres"`
	Url                  string   `json:"url"`
	Followers            int      `json:"followers"`
	Image                string   `json:"image"`
	MaxDanceability      float32  `json:"maxDanceability"`
	MaxDanceabilityTrack string   `json:"maxDanceabilityTrack"`
	MaxEnergy            float32  `json:"maxEnergy"`
	MaxEnergyTrack       string   `json:"maxEnergyTrack"`
	AlbumsCount          int      `json:"albumsCount"`
	TracksCount          int      `json:"tracksCount"`
}

type Album struct {
	Name                 string  `json:"name"`
	Type                 string  `json:"type"`
	ReleaseDate          string  `json:"releaseDate"`
	Tracks               []Track `json:"tracks"`
	MaxDanceability      float32 `json:"maxDanceability"`
	MaxDanceabilityTrack string  `json:"maxDanceabilityTrack"`
	MaxEnergy            float32 `json:"maxEnergy"`
	MaxEnergyTrack       string  `json:"maxEnergyTrack"`
	TracksCount          int     `json:"tracksCount"`
}

type Track struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	Danceability     float32 `json:"danceability"`
	Energy           float32 `json:"energy"`
	Acousticness     float32 `json:"acousticness"`
	Loudness         float32 `json:"loudness"`
	Liveness         float32 `json:"liveness"`
	Instrumentalness float32 `json:"instrumentalness"`
	Tempo            float64 `json:"tempo"`
}

func (a *Artist) ToString() string {
	genres := "Genres: not found"
	if len(a.Genres) > 0 {
		genres = fmt.Sprintf("Genres: %s", strings.Join(a.Genres, ", "))
	}
	popularity := fmt.Sprintf("Popularity: %v", a.Popularity)
	followers := fmt.Sprintf("Followers: %v", a.Followers)
	danceability := fmt.Sprintf("To dance: %s", a.MaxDanceabilityTrack)
	energy := fmt.Sprintf("To jump: %s", a.MaxEnergyTrack)
	albums := fmt.Sprintf("Total albums: %v", a.AlbumsCount)
	tracks := fmt.Sprintf("Total tracks: %v", a.TracksCount)
	albumsInfo := "Albums:\n"
	for i := range a.Albums {
		tracksText := "tracks"
		if a.Albums[i].TracksCount == 1 {
			tracksText = "track"
		}
		albumsInfo += fmt.Sprintf(
			"\t• %s (%s), %s, %v %s.\n",
			a.Albums[i].Name,
			strings.Split(a.Albums[i].ReleaseDate, "-")[0],
			a.Albums[i].Type,
			a.Albums[i].TracksCount,
			tracksText,
		)
	}

	artist := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		strings.ToUpper(a.Name),
		a.Url,
		genres,
		popularity,
		followers,
		danceability,
		energy,
		albums,
		tracks,
		albumsInfo,
	)
	return artist
}
