package handlers

import (
	"errors"

	"github.com/alexedwards/scs/v2"
	"github.com/diegoalzate/jot/internal/config"
	"github.com/diegoalzate/jot/internal/database"
)

type Web struct {
	db      database.Service
	config  config.Config
	session *scs.SessionManager
}

func New(db database.Service, config config.Config, session *scs.SessionManager) (Web, error) {
	if session == nil {
		return Web{}, errors.New("session can not be nil")
	}

	return Web{
		db:      db,
		config:  config,
		session: session,
	}, nil
}
