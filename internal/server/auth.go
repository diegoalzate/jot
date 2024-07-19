package server

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
)

func (config *Config) setupCookieAuth() {
	goth.UseProviders(
		discord.New(
			config.discord.clientKey,
			config.discord.clientSecret,
			"http://localhost:8080/api/auth/discord/callback",
			discord.ScopeIdentify, discord.ScopeEmail, discord.ScopeGuilds,
		),
		github.New(
			config.github.clientKey,
			config.github.clientSecret,
			"http://localhost:8080/api/auth/github/callback",
		),
	)

	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(config.session.cookieSecret))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = isProd

	gothic.Store = store
}
