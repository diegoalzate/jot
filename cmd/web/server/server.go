package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

var sessionManager *scs.SessionManager

type Server struct {
	db      database.Service
	session *scs.SessionManager
	config  config.Config
}

func New() (*http.Server, error) {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	NewServer := &Server{
		config:  config.New(),
		db:      database.New(),
		session: sessionManager,
	}
	handler, err := NewServer.RegisterRoutes()

	if err != nil {
		return &http.Server{}, err
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.config.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
