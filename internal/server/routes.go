package server

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/diegoalzate/jot/cmd/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// file server
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	// views
	r.Get("/web", templ.Handler(web.HomePage()).ServeHTTP)

	// api
	r.Get("/health", s.healthHandler)

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
