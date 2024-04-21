package handlers

import (
	"discord-spotify-bot/services"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
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

	podcastStr, err := services.GetPodcast(&client, podcastName)
	if err != nil {
		log.Println("[Podcast handler] Error in GetPodcast:", err)
		_, _ = discord.FollowupMessageCreate(
			interaction.Interaction,
			true,
			&discordgo.WebhookParams{
				Content: podcastStr,
			},
		)
		return
	}

	podcastLog := strings.ReplaceAll(podcastStr, "\n", " ")
	podcastLog = strings.ReplaceAll(podcastLog, "\t", " ")
	log.Println("[Podcast handler] /podcast response:", podcastLog)

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
