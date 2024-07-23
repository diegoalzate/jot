package handlers

import (
	"github.com/alexedwards/scs/v2"
	"github.com/diegoalzate/jot/internal/database"
)

type Handlers struct {
	db      database.Service
	session *scs.SessionManager
}

func New(db database.Service, session *scs.SessionManager) Handlers {
	return Handlers{
		db:      db,
		session: session,
	}
}
