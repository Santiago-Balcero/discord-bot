package handlers

import (
	"fmt"
	"log"
	"strings"

	"discord-spotify-bot/config"
	"discord-spotify-bot/services"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify"
)

func HandleInteraction(
	discord *discordgo.Session,
	interaction *discordgo.InteractionCreate,
) {
	client := config.Spotify
	data := interaction.ApplicationCommandData()

	log.Println(
		"Interaction received - command:",
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
	}
}

func GetArtist(
	discord *discordgo.Session,
	interaction *discordgo.InteractionCreate,
	data discordgo.ApplicationCommandInteractionData,
	client spotify.Client,
) {
	artistName := strings.TrimSpace(data.Options[0].StringValue())
	log.Println("Artist name:", artistName)

	// To keep active interaction
	err := discord.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		},
	)
	if err != nil {
		log.Println("Error in InteractionRespond:", err)
	}

	artistStr, err := services.GetArtist(&client, artistName)
	if err != nil {
		log.Println("[Artist handler] Error in GetArtist:", err)
		_, _ = discord.FollowupMessageCreate(
			interaction.Interaction,
			true,
			&discordgo.WebhookParams{
				Content: fmt.Sprintf(
					"Artist not found: %s",
					artistName,
				),
			},
		)
		return
	}

	artistLog := strings.ReplaceAll(artistStr, "\n", " ")
	artistLog = strings.ReplaceAll(artistLog, "\t", " ")
	log.Println("[Artist handler] /artist response:", artistLog)

	_, err = discord.FollowupMessageCreate(
		interaction.Interaction,
		true,
		&discordgo.WebhookParams{
			Content: artistStr,
		},
	)
	if err != nil {
		log.Println("Error sending response message:", err)
	}
}
