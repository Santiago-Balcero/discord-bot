package services

import (
	"context"
	"fmt"
	"log"

	"discord-spotify-bot/utils"

	"github.com/zmb3/spotify/v2"
)

func GetPodcast(client *spotify.Client, podcastName string) (string, error) {
	ctx := context.Background()
	podcast, err := searchPodcast(ctx, client, podcastName)
	if err != nil {
		return "", err
	}
	data, err := analysePodcast(ctx, client, podcast)
	if err != nil {
		return "", err
	}
	return data, nil
}

func searchPodcast(
	ctx context.Context,
	client *spotify.Client,
	podcastName string,
) (string, error) {
	result, err := client.Search(ctx, podcastName, spotify.SearchTypeShow)
	if err != nil {
		return "", fmt.Errorf("error searching podcast: %v", err)
	}

	podcastFound := false
	log.Println(result.Shows)

	for _, a := range result.Shows.Shows {
		// TODO as in artist service
		log.Println(a)
	}

	if !podcastFound {
		return "", fmt.Errorf("podcast not found")
	}

	return "", nil
}

func analysePodcast(
	ctx context.Context,
	client *spotify.Client,
	podcastId string,
) (string, error) {
	podcast, err := client.GetShow(ctx, spotify.ID(podcastId))
	utils.CheckError(err)

	fmt.Println("Podcast name:", podcast.Name)
	fmt.Println("Description:", podcast.Description)
	fmt.Println("Publisher:", podcast.Publisher)
	fmt.Println("Episodes endpoint:", podcast.Episodes.Endpoint)

	episodes, err := client.GetShowEpisodes(ctx, podcastId)
	utils.CheckError(err)

	for _, ep := range episodes.Episodes {
		fmt.Println("\tEpisode name:", ep.Name)
	}
	return "Analyse Podcast EOF", nil
}
