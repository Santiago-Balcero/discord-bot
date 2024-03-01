package config

import (
	"os"

	"github.com/Santiago-Balcero/discord-bot/utils"
	"github.com/joho/godotenv"
)

var SpotifyId string
var SpotifySecret string
var BotToken string

func LoadConfig() {
	err := godotenv.Load(".env")
	utils.CheckError(err)
	SpotifyId = os.Getenv("SPOTIFY_ID")
	SpotifySecret = os.Getenv("SPOTIFY_KEY")
	BotToken = os.Getenv("BOT_TOKEN")
}
