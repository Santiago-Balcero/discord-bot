package handlers

import (
	"discord-spotify-bot/services"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
)

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
				Content: artistStr,
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
