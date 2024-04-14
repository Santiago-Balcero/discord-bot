package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"discord-spotify-bot/models"
	"discord-spotify-bot/utils"

	"github.com/zmb3/spotify/v2"
)

func GetPodcast(client *spotify.Client, podcastName string) (string, error) {
	ctx := context.Background()
	podcastData, err := searchPodcast(ctx, client, podcastName)
	if err != nil {
		return "", err
	}
	err = analysePodcast(ctx, client, &podcastData)
	if err != nil {
		return "", err
	}

	return "OK", nil
}

func searchPodcast(
	ctx context.Context,
	client *spotify.Client,
	podcastName string,
) (models.Podcast, error) {
	podcast := models.Podcast{}

	options := []spotify.RequestOption{
		spotify.Market("CO"),
	}

	result, err := client.Search(ctx, podcastName, spotify.SearchTypeShow, options...)
	if err != nil {
		return podcast, fmt.Errorf("error searching podcast: %v", err)
	}

	podcastFound := false

	for _, s := range result.Shows.Shows {
		name := utils.ClearString(s.Name)
		if name == strings.ReplaceAll(podcastName, " ", "") {
			podcast.Id = s.ID.String()
			podcast.Name = s.Name
			podcast.Description = s.Description
			podcast.Publisher = s.Publisher
			podcast.Url = s.ExternalURLs["spotify"]
			podcast.Image = string(s.Images[2].URL)
			log.Println("Podcast found:", podcast.Name, podcast.Url)
			podcastFound = true
			break
		}
	}

	if !podcastFound {
		return podcast, fmt.Errorf("podcast not found")
	}

	return podcast, nil
}

func analysePodcast(
	ctx context.Context,
	client *spotify.Client,
	podcastData *models.Podcast,
) error {
	loc, _ := time.LoadLocation("America/Bogota")
	startTime := time.Now().In(loc)

	episodes, err := client.GetShowEpisodes(ctx, podcastData.Id)
	if err != nil {
		return fmt.Errorf("error in GetShowEpisodes: %v", err)
	}

	for _, ep := range episodes.Episodes {
		podcastData.EpisodesCount++
		episode := models.Episode{
			Id:          ep.ID.String(),
			Name:        ep.Name,
			Description: ep.Description,
			DurationMs:  ep.Duration_ms,
			ReleaseDate: ep.ReleaseDate,
		}
		podcastData.DurationMs += ep.Duration_ms
		podcastData.Episodes = append(podcastData.Episodes, episode)
	}

	fetchTime := time.Since(startTime)
	log.Println("Podcast data was fetched in:", fetchTime)
	return nil
}
