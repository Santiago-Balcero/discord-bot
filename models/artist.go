package models

type Artist struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Popularity           int      `json:"popularity"`
	Albums               []Album  `json:"albums"`
	Singles              []Album  `json:"singles"`
	Compilations         []Album  `josn:"compilations"`
	Genres               []string `json:"genres"`
	Url                  string   `json:"url"`
	Followers            int      `json:"followers"`
	Image                string   `json:"image"`
	MaxDanceability      float32  `json:"maxDanceability"`
	MaxDanceabilityTrack string   `json:"maxDanceabilityTrack"`
	MaxEnergy            float32  `json:"maxEnergy"`
	MaxEnergyTrack       string   `json:"maxEnergyTrack"`
	AlbumsCount          int      `json:"albumsCount"`
	SinglesCount         int      `json:"singlesCount"`
	CompilationsCount    int      `json:"compilationsCount"`
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
