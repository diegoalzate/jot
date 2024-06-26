package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/diegoalzate/jot/cmd/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	NewAuth()

	// file server
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	// views
	r.Get("/", templ.Handler(web.HomePage()).ServeHTTP)

	// api
	r.Get("/api/health", s.healthHandler)
	r.Get("/api/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		provider := r.PathValue("provider")
		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, "failed to complete authentication callback", http.StatusInternalServerError)
			return
		}
		log.Println(user)
		// redirect
	})

	r.Get("/api/auth/{provider}/logout", func(w http.ResponseWriter, r *http.Request) {
		provider := r.PathValue("provider")
		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		gothic.Logout(w, r)
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
			gothic.BeginAuthHandler(w, r)
		}
	})

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
