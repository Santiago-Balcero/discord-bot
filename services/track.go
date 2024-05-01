package services

import (
	"context"
	"fmt"
	"log"

	"discord-spotify-bot/interfaces"
	"discord-spotify-bot/models"

	"github.com/zmb3/spotify/v2"
)

func AnalyseTrack(client interfaces.SpotifyClient, track *models.Track) error {
	ctx := context.Background()
	features, err := client.GetAudioFeatures(ctx, spotify.ID(track.SpotifyId))
	if err != nil {
		return fmt.Errorf("error in GetAudioFeatures: %v", err)
	}
	if features[0] == nil {
		log.Println("No features found for track:", track.Name)
		return nil
	}

	track.Danceability = features[0].Danceability
	track.Energy = features[0].Energy
	track.Acousticness = features[0].Acousticness
	track.Loudness = features[0].Loudness
	track.Liveness = features[0].Liveness
	track.Liveness = features[0].Tempo
	track.Instrumentalness = features[0].Instrumentalness
	return nil
}
