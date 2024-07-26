package auth

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/diegoalzate/jot/internal/database"
	"github.com/diegoalzate/jot/internal/query"
)

type AuthHandler func(http.ResponseWriter, *http.Request, query.User)

type Middleware struct {
	db      database.Service
	session *scs.SessionManager
}

func NewMiddleware(db database.Service, session *scs.SessionManager) Middleware {
	return Middleware{
		db:      db,
		session: session,
	}
}
