package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	db     database.Service
	config config.Config
}

func New() (*http.Server, error) {
	NewServer := &Server{
		config: config.New(),
		db:     database.New(),
	}

	handler, err := NewServer.RegisterRoutes()

	if err != nil {
		return &http.Server{}, err
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.config.Ports.Api),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
