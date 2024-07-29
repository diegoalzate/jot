package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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

type CreateDiscordServerRequest struct {
	DiscordID string `json:"discord_id"`
	Name      string `json:"name"`
}

func guildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	fmt.Println("bot was installed on a server")
	reqUrl := "http://localhost:8080/api/servers"
	requestBody := CreateDiscordServerRequest{
		DiscordID: e.Guild.ID,
		Name:      e.Guild.Name,
	}

	rawBody, err := json.Marshal(requestBody)

	if err != nil {
		log.Printf("failed to marshall request body: %v", err.Error())
		s.ChannelMessageSend(e.SystemChannelID, "there was an error setting up the bot, please start again")
	}

	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(rawBody))

	if err != nil {
		log.Printf("failed to create request: %v", err.Error())
		s.ChannelMessageSend(e.SystemChannelID, "there was an error setting up the bot, please start again")
	}

	authToken := strings.Fields(s.Identify.Token)[1]

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authToken))

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	_, err = client.Do(req)

	if err != nil {
		log.Printf("failed to send request: %v", err.Error())
		s.ChannelMessageSend(e.SystemChannelID, "there was an error setting up the bot, please start again")
	}

	// TODO: handle response if server was already installed
	s.ChannelMessageSend(e.SystemChannelID, "congrats, continue setup on page")
}

func guildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	fmt.Println("bot was removed from a server")
}
