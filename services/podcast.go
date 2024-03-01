package services

import (
	"fmt"

	"github.com/Santiago-Balcero/discord-bot/utils"
	"github.com/zmb3/spotify"
)

func AnalysePodcast(client *spotify.Client, podcastId string) {
	podcast, err := client.GetShow(podcastId)
	utils.CheckError(err)

	fmt.Println("Podcast name:", podcast.Name)
	fmt.Println("Description:", podcast.Description)
	fmt.Println("Publisher:", podcast.Publisher)
	fmt.Println("Episodes endpoint:", podcast.Episodes.Endpoint)

	episodes, err := client.GetShowEpisodes(podcastId)
	utils.CheckError(err)

	for _, ep := range episodes.Episodes {
		fmt.Println("\tEpisode name:", ep.Name)
	}
}
