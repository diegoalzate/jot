package auth

import (
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
			config.Discord.Key,
			config.Discord.Secret,
			"http://localhost:8080/api/auth/discord/callback",
			discord.ScopeIdentify, discord.ScopeEmail, discord.ScopeGuilds,
		),
		github.New(
			config.Github.Key,
			config.Github.Secret,
			"http://localhost:8080/api/auth/github/callback",
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
