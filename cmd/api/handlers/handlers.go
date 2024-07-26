package handlers

import (
	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
)

type HandlerContext struct {
	db     database.Service
	config config.Config
}

func New(db database.Service, config config.Config) HandlerContext {
	return HandlerContext{
		db:     db,
		config: config,
	}
}
