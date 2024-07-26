package api

import (
	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
)

type AuthContext struct {
	db     database.Service
	config config.Config
}

func WithServer() {}
