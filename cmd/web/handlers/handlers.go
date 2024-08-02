package handlers

import (
	"errors"

	"github.com/alexedwards/scs/v2"
	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
)

type Handlers struct {
	db      database.Service
	config  config.Config
	session *scs.SessionManager
}

func New(db database.Service, config config.Config, session *scs.SessionManager) (Handlers, error) {
	if session == nil {
		return Handlers{}, errors.New("session can not be nil")
	}

	return Handlers{
		db:      db,
		config:  config,
		session: session,
	}, nil
}
