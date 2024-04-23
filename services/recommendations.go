package services

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
)

func GetRecommendations(client *spotify.Client) {
	ctx := context.Background()

	seed := spotify.Seeds{
		Artists: []spotify.ID{
			spotify.ID(""),
			spotify.ID(""),
		},
		Tracks: []spotify.ID{
			spotify.ID(""), // Nao force
			spotify.ID(""), // Panthera
		},
		Genres: []string{"Funk"},
	}

	trackAttributes := spotify.NewTrackAttributes()
	trackAttributes.MinTempo(120)
	trackAttributes.MaxTempo(150)
	trackAttributes.MinDanceability(0.7)
	trackAttributes.MinEnergy(0.6)

	recommendations, _ := client.GetRecommendations(
		ctx,
		seed,
		trackAttributes,
	)

	for _, track := range recommendations.Tracks {
		fmt.Printf(
			"Recommendation: %s - %v\n",
			track.Name,
			track.ExternalURLs["spotify"],
		)
	}
}
