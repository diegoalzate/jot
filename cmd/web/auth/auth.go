package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/diegoalzate/jot/cmd/ui"
	"github.com/diegoalzate/jot/internal/database"
	"github.com/diegoalzate/jot/internal/query"
	"github.com/google/uuid"
)

type AuthHandler func(http.ResponseWriter, *http.Request, query.User)

type AuthContext struct {
	db      database.Service
	session *scs.SessionManager
}

func New(db database.Service, session *scs.SessionManager) AuthContext {
	return AuthContext{
		db:      db,
		session: session,
	}
}

func (a *AuthContext) WithUser(fn AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookieUserId := a.session.GetString(r.Context(), "user_id")
		if cookieUserId == "" {
			log.Print("user cookie is undefined")
			ui.LoginPage().Render(r.Context(), w)
			return
		}

		q := query.New(a.db.Conn)

		userID, err := uuid.Parse(cookieUserId)

		if err != nil {
			log.Printf("failed to parse userid: %v", err)
			ui.LoginPage().Render(r.Context(), w)
			return
		}

		dbUser, err := q.GetUserById(context.Background(), userID)

		if err != nil {
			log.Printf("failed to get user: %v", err)
			ui.LoginPage().Render(r.Context(), w)
			return
		}

		fn(w, r, dbUser)
		return
	}
}
