package services

import (
	"context"
	"fmt"

	"discord-spotify-bot/models"

	"github.com/zmb3/spotify/v2"
)

func AnalyseTrack(client *spotify.Client, track *models.Track) error {
	ctx := context.Background()
	features, err := client.GetAudioFeatures(ctx, spotify.ID(track.Id))
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
