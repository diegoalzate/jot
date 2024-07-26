package server

import (
	"net/http"

	"github.com/diegoalzate/jot/cmd/ui"
	"github.com/diegoalzate/jot/cmd/web/auth"
	"github.com/diegoalzate/jot/cmd/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() (http.Handler, error) {
	auth.SetupAuthProviders(&s.config)

	r := chi.NewRouter()

	r.Use(sessionManager.LoadAndSave)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// file server
	fileServer := http.FileServer(http.FS(ui.Files))
	r.Handle("/assets/*", fileServer)

	middleware := auth.New(s.db, s.session)
	handlers, err := handlers.New(s.db, s.config, s.session)

	if err != nil {
		return r, err
	}

	// views
	r.Get("/", middleware.WithUser(handlers.ViewHome))

	// api
	r.Get("/health", handlers.HealthCheck)
	r.Get("/auth/{provider}/callback", handlers.OauthCallback)
	r.Get("/auth/{provider}/logout", handlers.Logout)
	r.Get("/auth/{provider}", handlers.Login)

	return r, nil
}
