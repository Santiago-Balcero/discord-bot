package config

import (
	"log"
	"os"

	"discord-spotify-bot/utils"

	"github.com/joho/godotenv"
)

var SpotifyId string
var SpotifySecret string
var BotToken string
var DiscordAppId string

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env: ", err)
	}
	utils.CheckError(err)
	SpotifyId = os.Getenv("SPOTIFY_ID")
	SpotifySecret = os.Getenv("SPOTIFY_KEY")
	BotToken = os.Getenv("BOT_TOKEN")
	DiscordAppId = os.Getenv("DISCORD_APP_ID")
}
