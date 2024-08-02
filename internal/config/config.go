package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port int
}

func New() Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return Config{
		Port: port,
	}

}
