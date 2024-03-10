package services

import (
	"fmt"

	"discord-spotify-bot/models"

	"github.com/zmb3/spotify"
)

func AnalyseTrack(client *spotify.Client, track *models.Track) error {
	features, err := client.GetAudioFeatures(spotify.ID(track.Id))
	if err != nil {
		return fmt.Errorf("error in GetAudioFeatures: %v", err)
	}

	track.Danceability = features[0].Danceability
	track.Energy = features[0].Energy
	track.Acousticness = features[0].Acousticness
	track.Loudness = features[0].Loudness
	track.Liveness = features[0].Liveness
	track.Instrumentalness = features[0].Instrumentalness
	track.Liveness = features[0].Tempo

	return nil
}
