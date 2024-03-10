package services

import (
	"fmt"

	"discord-spotify-bot/utils"
	"github.com/zmb3/spotify"
)

func GetRecommendations(client *spotify.Client) {
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

	country := "BR"
	options := spotify.Options{
		Country: &country,
	}

	recommendations, err := client.GetRecommendations(
		seed,
		trackAttributes,
		&options,
	)
	utils.CheckError(err)

	for _, track := range recommendations.Tracks {
		fmt.Printf(
			"Recommendation: %s - %v\n",
			track.Name,
			track.ExternalURLs["spotify"],
		)
	}
}
