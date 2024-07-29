package auth

import (
	"fmt"

	"github.com/diegoalzate/jot/internal/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
)

func SetupAuthProviders(config *config.Config) {
	goth.UseProviders(
		discord.New(
			config.Discord.Oauth.Key,
			config.Discord.Oauth.Secret,
			fmt.Sprintf("http://localhost:%v/auth/discord/callback", config.Ports.Web),
			discord.ScopeIdentify, discord.ScopeEmail, discord.ScopeGuilds,
		),
		github.New(
			config.Github.Key,
			config.Github.Secret,
			fmt.Sprintf("http://localhost:%v/auth/github/callback", config.Ports.Web),
		),
	)

	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(config.Session.Secret))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = isProd

	gothic.Store = store
}
