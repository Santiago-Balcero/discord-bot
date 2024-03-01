package main

import (
	"log"

	bot "github.com/Santiago-Balcero/discord-bot/bot"
	"github.com/Santiago-Balcero/discord-bot/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.LoadConfig()

	bot.Run()
}
