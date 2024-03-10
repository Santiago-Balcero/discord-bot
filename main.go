package main

import (
	bot "github.com/Santiago-Balcero/discord-bot/bot"
	"github.com/Santiago-Balcero/discord-bot/config"
)

func main() {
	config.LoadConfig()
	config.CreateSpotifyClient()
	bot.Run()
}
