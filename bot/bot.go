package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/Santiago-Balcero/discord-bot/config"
	"github.com/Santiago-Balcero/discord-bot/handlers"
	"github.com/bwmarrin/discordgo"
)

func Run() {
	token := config.BotToken
	// create session
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	// add an event handler
	discord.AddHandler(handlers.GetArtist)

	// open session
	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening Discord session:", err)
	}
	defer discord.Close()

	log.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Bot shut down")
}
