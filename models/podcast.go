package models

type Podcast struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Publisher     string    `json:"publisher"`
	Url           string    `json:"url"`
	Image         string    `json:"image"`
	Episodes      []Episode `json:"episodes"`
	EpisodesCount int       `json:"episodesCount"`
	DurationMs    int       `json:"durationMs"`
}

type Episode struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DurationMs  int    `json:"durationMs"`
	ReleaseDate string `json:"releaseDate"`
}
