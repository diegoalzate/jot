package server

import (
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
)

func setupAuthProvider() {
	goth.UseProviders(
		discord.New(os.Getenv("DISCORD_CLIENT_ID"), os.Getenv("DISCORD_CLIENT_SECRET"), "http://localhost:8080/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:8080/auth/github/callback"),
	)
}

func newAuthStore() sessions.Store {
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = isProd
	return store
}

func NewAuth() {
	setupAuthProvider()
	store := newAuthStore()
	gothic.Store = store
}
