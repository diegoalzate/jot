package server

import (
	"net/http"

	"github.com/diegoalzate/jot/cmd/ui"
	"github.com/diegoalzate/jot/cmd/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() (http.Handler, error) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// file server
	fileServer := http.FileServer(http.FS(ui.Files))
	r.Handle("/assets/*", fileServer)

	handlers, err := handlers.New(s.db, s.config, s.session)

	if err != nil {
		return r, err
	}

	// views
	r.Get("/", handlers.ViewHome)

	// api
	r.Get("/health", handlers.HealthCheck)
	return r, nil
}
