package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SpotifyId string
var SpotifySecret string
var BotToken string
var DiscordAppId string

func LoadConfig() {
	if dockerEnv := os.Getenv("DOCKER_ENV"); dockerEnv == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env: ", err)
		}
	}

	port := os.Getenv("PORT")
	SpotifyId = os.Getenv("SPOTIFY_ID")
	SpotifySecret = os.Getenv("SPOTIFY_KEY")
	BotToken = os.Getenv("BOT_TOKEN")
	DiscordAppId = os.Getenv("DISCORD_APP_ID")
}
