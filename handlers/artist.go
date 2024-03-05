package handlers

import (
	"log"
	"strings"

	"github.com/Santiago-Balcero/discord-bot/config"
	"github.com/Santiago-Balcero/discord-bot/models"
	services "github.com/Santiago-Balcero/discord-bot/services"
	"github.com/bwmarrin/discordgo"
)

func GetArtist(discord *discordgo.Session, message *discordgo.MessageCreate) {
	client, err := config.GetClient()
	if err != nil {
		log.Println("Error:", err)
		discord.ChannelMessageSend(
			message.ChannelID,
			"Service unavailable. Try again.",
		)
	}
	msg := models.Message{
		Author: models.Author{
			ID:       message.Author.ID,
			Email:    message.Author.Email,
			Locale:   message.Author.Locale,
			Username: message.Author.Username,
			Verified: message.Author.Verified,
		},
		Content: message.Content,
	}

	// prevent bot responding to its own messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	log.Println("Message received:", msg)

	messageContent := strings.ToLower(message.Content)

	switch {
	case messageContent[:11] == "!artist:":
		discord.ChannelMessageSend(
			message.ChannelID,
			"Searching artist data...",
		)
		artistName := strings.Split(messageContent, ":")[1]
		artistName = strings.TrimSpace(artistName)
		log.Println("Request for !artist:", artistName)
		artistStr, err := services.GetArtist(&client, artistName)
		if err != nil {
			log.Println("[ARTIST HANDLER] Error in GetArtist:", err)
			discord.ChannelMessageSend(
				message.ChannelID,
				"Can't find artist data. Try again.",
			)
			return
		}
		log.Println("[ARTIST HANDLER] !artist response:", artistStr)
		discord.ChannelMessageSend(message.ChannelID, "Artist data:\n")
		discord.ChannelMessageSend(message.ChannelID, artistStr)
	default:
		log.Println("Invalid command")
		discord.ChannelMessageSend(message.ChannelID, "Invalid command.")
	}
}
