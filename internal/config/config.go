package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Ports struct {
		Api int
		Web int
	}

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
	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))
	webPort, _ := strconv.Atoi(os.Getenv("WEB_PORT"))
	githubClientKey := os.Getenv("GITHUB_KEY")
	githubClientSecret := os.Getenv("GITHUB_SECRET")
	discordClientKey := os.Getenv("DISCORD_KEY")
	discordClientSecret := os.Getenv("DISCORD_SECRET")
	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	cookieSecret := os.Getenv("COOKIE_SECRET")

	return Config{
		Ports: struct {
			Api int
			Web int
		}{
			Api: apiPort,
			Web: webPort,
		},
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
