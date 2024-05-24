package main

import (
	"discord-spotify-bot/bot"
	"discord-spotify-bot/config"
	"discord-spotify-bot/server"
)

func main() {
	config.LoadConfig()
	config.ConnectToDB()
	server.InitServer()
	config.CreateSpotifyClient()
	bot.Run()
}
