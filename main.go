package main

import (
	"discord-spotify-bot/bot"
	"discord-spotify-bot/config"
)

func main() {
	config.LoadConfig()
	config.CreateSpotifyClient()
	bot.Run()
}
