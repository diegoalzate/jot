package server

import (
	"net/http"

	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
	"github.com/diegoalzate/jot/internal/query"
)

type AuthHandler func(http.ResponseWriter, *http.Request, query.DiscordServer)

type AuthContext struct {
	db     database.Service
	config config.Config
}

func WithDiscordServer(fn AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check auth headers for discord token
	}
}
