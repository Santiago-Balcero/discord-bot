package handlers

import (
	"discord-spotify-bot/services"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify"
)

func GetPodcast(
	discord *discordgo.Session,
	interaction *discordgo.InteractionCreate,
	data discordgo.ApplicationCommandInteractionData,
	client spotify.Client,
) {
	podcastName := strings.TrimSpace(data.Options[0].StringValue())
	log.Println("Podcast name:", podcastName)

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

	podcastStr, err := services.AnalysePodcast(&client, podcastName)
	if err != nil {
		log.Println("[Podcast handler] Error in GetPodcast:", err)
		_, _ = discord.FollowupMessageCreate(
			interaction.Interaction,
			true,
			&discordgo.WebhookParams{
				Content: fmt.Sprintf(
					"Podcast not found: %s",
					podcastName,
				),
			},
		)
		return
	}

	log.Println("[Podcast handler] /podcast response:", podcastStr)

	_, err = discord.FollowupMessageCreate(
		interaction.Interaction,
		true,
		&discordgo.WebhookParams{
			Content: podcastStr,
		},
	)
	if err != nil {
		log.Println("Error sending response message:", err)
	}
}
