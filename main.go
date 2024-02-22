package main

import (
	"log"
	"os"

	bot "github.com/Santiago-Balcero/discord-bot/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
	token := os.Getenv("BOT_TOKEN")
	bot.BotToken = token
	bot.Run()
}
