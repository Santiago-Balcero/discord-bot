package config

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var SpotifyId string
var SpotifySecret string
var BotToken string
var DiscordAppId string
var DbUser string
var DbPassword string
var DbHost string
var DbName string
var DbDriver string
var DbPort string

func LoadConfig() {
	if dockerEnv := os.Getenv("DOCKER_ENV"); dockerEnv == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env: ", err)
		}
	}

	port := os.Getenv("PORT")
	log.Println("Assigned port:", port)
	go http.ListenAndServe(":"+port, nil)
	SpotifyId = os.Getenv("SPOTIFY_ID")
	SpotifySecret = os.Getenv("SPOTIFY_KEY")
	BotToken = os.Getenv("BOT_TOKEN")
	DiscordAppId = os.Getenv("DISCORD_APP_ID")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbName = os.Getenv("DB_NAME")
	DbDriver = os.Getenv("DB_DRIVER")
}
