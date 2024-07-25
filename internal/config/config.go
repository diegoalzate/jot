package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port    int
	Discord struct {
		Oauth OauthConfig
		Bot   struct {
			Token string
		}
	}
	Github  OauthConfig
	Session struct {
		Secret string
	}
}

type OauthConfig struct {
	Key    string
	Secret string
}

func New() Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	githubClientKey := os.Getenv("GITHUB_KEY")
	githubClientSecret := os.Getenv("GITHUB_SECRET")
	discordClientKey := os.Getenv("DISCORD_KEY")
	discordClientSecret := os.Getenv("DISCORD_SECRET")
	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	cookieSecret := os.Getenv("COOKIE_SECRET")

	return Config{
		Port: port,
		Discord: struct {
			Oauth OauthConfig
			Bot   struct {
				Token string
			}
		}{
			Oauth: OauthConfig{
				Key:    discordClientKey,
				Secret: discordClientSecret,
			},
			Bot: struct{ Token string }{
				Token: discordBotToken,
			},
		},
		Github: OauthConfig{
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
