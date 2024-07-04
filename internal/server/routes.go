package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/diegoalzate/jot/cmd/web"
	"github.com/diegoalzate/jot/internal/query"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
)

type authHandler func(http.ResponseWriter, *http.Request, query.User)

var sessionManager *scs.SessionManager

func (s *Server) RegisterRoutes() http.Handler {
	// sets up goth internal store
	s.config.setupCookieAuth()
	// setup cookie auth
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	r := chi.NewRouter()
	r.Use(sessionManager.LoadAndSave)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// file server
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	// views
	r.Get("/", s.withUser(func(w http.ResponseWriter, r *http.Request, u query.User) {
		isLoggedIn := false

		log.Print(u)
		if u.ID != (uuid.UUID{}) {
			isLoggedIn = true
		}

		web.HomePage(isLoggedIn, u).Render(r.Context(), w)
	}))

	// api
	r.Get("/api/health", s.healthHandler)
	r.Get("/api/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		provider := r.PathValue("provider")
		r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

		// logs out of gothic store
		gothUser, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			log.Printf("[ERR]: %#v", err)
			http.Error(w, "failed to complete authentication callback", http.StatusInternalServerError)
			return
		}
		// app session management
		err = sessionManager.RenewToken(r.Context())
		if err != nil {
			http.Error(w, "failed to renew token", http.StatusInternalServerError)
			return
		}

		// check if user exists
		q := query.New(s.db.Conn)

		dbIdentity, err := q.GetIdentity(context.Background(), query.GetIdentityParams{
			ProviderID: gothUser.UserID,
			Provider:   gothUser.Provider,
		})

		if err != nil && err != sql.ErrNoRows {
			log.Printf("[ERR]: %#v", err)
			http.Error(w, "failed to get identity", http.StatusInternalServerError)
			return
		}

		if err == sql.ErrNoRows {
			// create user
			id, err := uuid.NewV7()

			if err != nil {
				log.Printf("[ERR]: %#v", err)
				http.Error(w, "failed to create uuid", http.StatusInternalServerError)
				return
			}

			newUser, err := q.CreateUser(
				context.Background(),
				query.CreateUserParams{
					ID:        id,
					Username:  gothUser.NickName,
					Email:     gothUser.Email,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			)

			if err != nil {
				log.Printf("[ERR]: %#v", err)
				http.Error(w, "failed to create user", http.StatusInternalServerError)
				return
			}

			id, err = uuid.NewV7()

			if err != nil {
				log.Printf("[ERR]: %#v", err)
				http.Error(w, "failed to create uuid", http.StatusInternalServerError)
				return
			}

			rawData, err := json.Marshal(gothUser.RawData)

			if err != nil {
				log.Printf("[ERR]: %#v", err)
				http.Error(w, "failed to convert raw data to bytes", http.StatusInternalServerError)
				return
			}

			_, err = q.CreateIdentity(context.Background(), query.CreateIdentityParams{
				ID:           id,
				UserID:       newUser.ID,
				Provider:     gothUser.Provider,
				ProviderID:   gothUser.UserID,
				IdentityData: rawData,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			})

			log.Printf("[SUCCESS]:user: %#v", newUser)
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}

		dbUser, err := q.GetUserById(context.Background(), dbIdentity.UserID)

		if err != nil {
			log.Printf("[ERR]: %#v", err)
			http.Error(w, "failed to get user", http.StatusInternalServerError)
			return
		}

		// redirect home
		sessionManager.Put(r.Context(), "user_id", dbUser.ID.String())
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	})

	r.Get("/api/auth/{provider}/logout", func(w http.ResponseWriter, r *http.Request) {
		provider := r.PathValue("provider")
		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		gothic.Logout(w, r)
		sessionManager.Destroy(r.Context())
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	r.Get("/api/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		provider := r.PathValue("provider")
		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			log.Println("[ALREADY LOGGED IN]", gothUser)
		} else {
			url, err := gothic.GetAuthURL(w, r)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Add("HX-Redirect", url)
		}
	})

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) withUser(fn authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookieUserId := sessionManager.GetString(r.Context(), "user_id")
		log.Print("cookie id", cookieUserId)
		if cookieUserId == "" {
			fn(w, r, query.User{})
			return
		}

		q := query.New(s.db.Conn)

		userID, err := uuid.Parse(cookieUserId)
		if err != nil {
			http.Error(w, "failed to parse user ID", http.StatusInternalServerError)
			return
		}

		dbUser, err := q.GetUserById(context.Background(), userID)

		if err != nil {
			fn(w, r, query.User{})
			return
		}

		fn(w, r, dbUser)
		return
	}
}
