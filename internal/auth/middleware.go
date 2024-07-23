package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/diegoalzate/jot/cmd/web"
	"github.com/diegoalzate/jot/internal/query"
	"github.com/google/uuid"
)

func (m *Middleware) WithUser(fn AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookieUserId := m.session.GetString(r.Context(), "user_id")
		if cookieUserId == "" {
			log.Print("user cookie is undefined")
			web.LoginPage().Render(r.Context(), w)
			return
		}

		q := query.New(m.db.Conn)

		userID, err := uuid.Parse(cookieUserId)

		if err != nil {
			log.Printf("failed to parse userid: %v", err)
			web.LoginPage().Render(r.Context(), w)
			return
		}

		dbUser, err := q.GetUserById(context.Background(), userID)

		if err != nil {
			log.Printf("failed to get user: %v", err)
			web.LoginPage().Render(r.Context(), w)
			return
		}

		fn(w, r, dbUser)
		return
	}
}
