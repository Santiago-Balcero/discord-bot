package handlers

import (
	"log"

	"discord-spotify-bot/config"

	"github.com/bwmarrin/discordgo"
)

func HandleInteraction(
	discord *discordgo.Session,
	interaction *discordgo.InteractionCreate,
) {
	client := config.Spotify
	data := interaction.ApplicationCommandData()

	log.Println(
		"Interaction received - user:",
		data.Resolved.Users,
		"- command:",
		data.Name,
		"- value:",
		data.Options[0].StringValue(),
	)

	switch data.Name {
	case "artist":
		GetArtist(
			discord,
			interaction,
			data,
			client,
		)
	case "podcast":
		GetPodcast(
			discord,
			interaction,
			data,
			client,
		)
	}
}
