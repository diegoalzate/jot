package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/diegoalzate/jot/internal/config"
)

func main() {
	config := config.New()

	discord, err := discordgo.New("Bot " + config.Discord.Bot.Token)

	if err != nil {
		fmt.Println("failed to start discord bot: ", err)
		return
	}

	// add for discord
	discord.AddHandler(guildCreate)
	discord.AddHandler(guildDelete)
	discord.AddHandler(messageCreate)

	// add permissions we require
	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func guildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	fmt.Println("bot was installed on a server")
}

func guildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	fmt.Println("bot was removed from a server")
}
