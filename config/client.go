package config

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func GetClient() (spotify.Client, error) {
	authConfig := &clientcredentials.Config{
		ClientID:     SpotifyId,
		ClientSecret: SpotifySecret,
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		return spotify.Authenticator{}.NewClient(nil),
			fmt.Errorf("error creating spotify client")
	}

	return spotify.Authenticator{}.NewClient(accessToken), nil
}
