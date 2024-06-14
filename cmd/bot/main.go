package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v62/github"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	DISCORD_BOT_TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	GITHUB_PERSONAL_ACCESS_TOKEN := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")

	if DISCORD_BOT_TOKEN == "" {
		fmt.Println("DISCORD_BOT_TOKEN is required to run this app")
		return
	}

	if GITHUB_PERSONAL_ACCESS_TOKEN == "" {
		fmt.Println("GITHUB_PERSONAL_ACCESS_TOKEN is required to run this app")
		return
	}

	discord, err := discordgo.New("Bot " + DISCORD_BOT_TOKEN)
	client := github.NewClient(nil).WithAuthToken(GITHUB_PERSONAL_ACCESS_TOKEN)

	if err != nil {
		fmt.Println("failed to start discord bot: ", err)
		return
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			// lets see a repo's issues
			repos, _, err := client.Repositories.ListByAuthenticatedUser(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			}

			for _, repo := range repos {
				s.ChannelMessageSend(m.ChannelID, *repo.Name)
			}

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
