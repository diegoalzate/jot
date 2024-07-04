package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/diegoalzate/jot/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	db     database.Service
	config Config
}

func NewServer() *http.Server {
	NewServer := &Server{
		config: newConfig(),
		db:     database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.config.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
