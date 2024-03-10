package config

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var Spotify spotify.Client

func CreateSpotifyClient() {
	authConfig := &clientcredentials.Config{
		ClientID:     SpotifyId,
		ClientSecret: SpotifySecret,
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatal("Error creating spotify client: ", err)
	}
	Spotify = spotify.Authenticator{}.NewClient(accessToken)
}
