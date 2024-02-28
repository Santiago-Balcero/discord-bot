package bot

import (
	"log"
	"os"
	"os/signal"
	"strings"

	services "github.com/Santiago-Balcero/discord-bot/services/spotify"
	"github.com/bwmarrin/discordgo"
)

type Message struct {
	Author  Author
	Content string
}

type Author struct {
	ID       string
	Email    string
	Locale   string
	Username string
	Verified bool
}

var BotToken string

func checkErr(e error) {
	if e != nil {
		log.Fatal("Error: ", e)
	}
}

func Run() {
	// create session
	discord, err := discordgo.New("Bot " + BotToken)
	checkErr(err)

	// add an event handler
	discord.AddHandler(newMessage)

	// open session
	err = discord.Open()
	checkErr(err)
	defer discord.Close()

	log.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Bot shut down")
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	msg := Message{
		Author: Author{
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

	log.Println("Message: ", msg)

	messageContent := strings.ToLower(message.Content)

	switch {
	case strings.Contains(messageContent, "get-artist"):
		discord.ChannelMessageSend(message.ChannelID, "Searching artist data...")
		artistName := strings.Split(messageContent, ":")[1]
		artistName = strings.TrimSpace(artistName)
		artistName = strings.ReplaceAll(artistName, " ", "-")
		log.Println("Request for get-artist:", artistName)
		artistData, err := services.GetArtist(artistName)
		if err != nil {
			log.Println("Error:", err)
			discord.ChannelMessageSend(message.ChannelID, "Can't give you artist data. Try again.")
			return
		}
		log.Println("get-artist response:", artistData)
		discord.ChannelMessageSend(message.ChannelID, artistData.ToString())
	}
}
