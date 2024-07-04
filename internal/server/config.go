package server

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	port    int
	discord struct {
		clientKey    string
		clientSecret string
	}
	github struct {
		clientKey    string
		clientSecret string
	}
	session struct {
		cookieSecret string
	}
}

func newConfig() Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	githubClientKey := os.Getenv("GITHUB_KEY")
	githubClientSecret := os.Getenv("GITHUB_SECRET")
	discordClientKey := os.Getenv("DISCORD_KEY")
	discordClientSecret := os.Getenv("DISCORD_SECRET")
	cookieSecret := os.Getenv("COOKIE_SECRET")

	return Config{
		port: port,
		discord: struct {
			clientKey    string
			clientSecret string
		}{
			clientKey:    discordClientKey,
			clientSecret: discordClientSecret,
		},
		github: struct {
			clientKey    string
			clientSecret string
		}{
			clientKey:    githubClientKey,
			clientSecret: githubClientSecret,
		},
		session: struct {
			cookieSecret string
		}{
			cookieSecret: cookieSecret,
		},
	}

}
