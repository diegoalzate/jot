package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

type ApiServer struct {
	db     database.Service
	config config.Config
}

func NewApiServer() *http.Server {
	NewServer := &ApiServer{
		config: config.New(),
		db:     database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.config.Port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
