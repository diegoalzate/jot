package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/diegoalzate/jot/cmd/web"
	"github.com/diegoalzate/jot/internal/query"
)

func (h *Handlers) ViewHome(w http.ResponseWriter, r *http.Request, u query.User) {
	// get discord servers user has or the bot is already installed on
	accessToken := h.session.GetString(r.Context(), "discord_token")
	discordClient, err := discordgo.New("Bearer " + accessToken)
	if err != nil {
		log.Fatal("could not create discord client")
		return
	}

	guilds, err := discordClient.UserGuilds(20, "", "", false)

	if err != nil {
		log.Fatal("could not get guilds")
		return
	}

	adminGuilds := []*discordgo.UserGuild{}

	for _, guild := range guilds {
		if guild.Permissions&discordgo.PermissionManageServer != 0 {
			adminGuilds = append(adminGuilds, guild)
		}
	}

	installLink := fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=%v", h.config.Discord.Oauth.Key)

	web.HomePage(u, adminGuilds, installLink).Render(r.Context(), w)
	return
}
