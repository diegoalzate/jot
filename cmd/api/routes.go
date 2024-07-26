package api

import (
	"net/http"

	"github.com/diegoalzate/jot/cmd/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) RegisterRoutes() (http.Handler, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handlers := handlers.New(s.db, s.config)

	// api
	r.Post("/api/servers", handlers.CreateServer)

	return r, nil
}
