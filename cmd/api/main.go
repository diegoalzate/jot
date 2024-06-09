package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	DISCORD_BOT_TOKEN := os.Getenv("DISCORD_BOT_TOKEN")

	if DISCORD_BOT_TOKEN == "" {
		fmt.Println("DISCORD_BOT_TOKEN is required to run this app")
		return
	}

	discord, err := discordgo.New("Bot " + DISCORD_BOT_TOKEN)

	if err != nil {
		fmt.Println("failed to start discord bot: ", err)
		return
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}

		if m.Content == "pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}
	})

	// add permissions we require
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// open ws connection to discord
	err = discord.Open()
	defer discord.Close()

	if err != nil {
		fmt.Println("failed to start ws connection: ", err)
		return
	}

	// wait here until CTRL + C or other term signal is received
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
