package bot

import (
	"log"
	"os"
	"os/signal"

	"discord-spotify-bot/config"
	"discord-spotify-bot/handlers"

	"github.com/bwmarrin/discordgo"
)

func Run() {
	// create session
	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "artist",
			Description: "Get information about an artist",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "The name of the artist",
					Required:    true,
				},
			},
		},
		{
			Name:        "podcast",
			Description: "Get information about a podcast",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "The name of the podcast",
					Required:    true,
				},
			},
		},
	}
	// set commands
	_, err = discord.ApplicationCommandBulkOverwrite(
		config.DiscordAppId,
		"",
		commands,
	)
	if err != nil {
		log.Fatal("Error creating bot commands: ", err)
	}

	// add an event handler
	discord.AddHandler(handlers.GetArtist)

	// open websocket session
	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}
	defer discord.Close()

	log.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Bot shut down")
}
