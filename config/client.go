package config

import (
	"context"
	"log"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
)

var Spotify spotify.Client

func CreateSpotifyClient() {
	ctx := context.Background()
	authConfig := &clientcredentials.Config{
		ClientID:     SpotifyId,
		ClientSecret: SpotifySecret,
		TokenURL:     spotifyauth.TokenURL,
	}

	accessToken, err := authConfig.Token(ctx)
	if err != nil {
		log.Fatal("Error creating spotify client: ", err)
	}
	httpClient := spotifyauth.New().Client(ctx, accessToken)
	Spotify = *spotify.New(httpClient)
}
