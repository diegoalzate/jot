package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port    int
	Discord struct {
		Key    string
		Secret string
	}
	Github struct {
		Key    string
		Secret string
	}
	Session struct {
		Secret string
	}
}

func New() Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	githubClientKey := os.Getenv("GITHUB_KEY")
	githubClientSecret := os.Getenv("GITHUB_SECRET")
	discordClientKey := os.Getenv("DISCORD_KEY")
	discordClientSecret := os.Getenv("DISCORD_SECRET")
	cookieSecret := os.Getenv("COOKIE_SECRET")

	return Config{
		Port: port,
		Discord: struct {
			Key    string
			Secret string
		}{
			Key:    discordClientKey,
			Secret: discordClientSecret,
		},
		Github: struct {
			Key    string
			Secret string
		}{
			Key:    githubClientKey,
			Secret: githubClientSecret,
		},
		Session: struct {
			Secret string
		}{
			Secret: cookieSecret,
		},
	}

}
