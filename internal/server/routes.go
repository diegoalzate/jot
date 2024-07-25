package server

import (
	"net/http"

	"github.com/diegoalzate/jot/cmd/web"
	"github.com/diegoalzate/jot/internal/auth"
	"github.com/diegoalzate/jot/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	auth.SetupAuthProviders(&s.config)

	r := chi.NewRouter()

	r.Use(sessionManager.LoadAndSave)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// file server
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	middleware := auth.NewMiddleware(s.db, s.session)
	handlers := handlers.New(s.db, s.session, s.config)

	// views
	r.Get("/", middleware.WithUser(handlers.ViewHome))

	// api
	r.Get("/api/health", handlers.HealthCheck)
	r.Get("/api/auth/{provider}/callback", handlers.OauthCallback)
	r.Get("/api/auth/{provider}/logout", handlers.Logout)
	r.Get("/api/auth/{provider}", handlers.Login)

	return r
}
