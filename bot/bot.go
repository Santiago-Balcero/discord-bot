package bot

import (
	"log"
	"os"
	"os/signal"
	"strings"

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
	discord.Open()
	defer discord.Close()

	log.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
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
	log.Println("Message: ", msg)

	// prevent bot responding to its own messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.ToLower(message.Content) == "new":
		discord.ChannelMessageSend(message.ChannelID, "New game is about to start! Enter name of opponent.")
	case strings.ToLower(message.Content) == "end":
		discord.ChannelMessageSend(message.ChannelID, "Game has ended!")
	}
}
