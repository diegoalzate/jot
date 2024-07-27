package server

import (
	"log"
	"net/http"
	"strings"
)

func (s *Server) protectedRoute(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check auth headers for discord token
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			log.Printf("authorization header is not defined")
			http.Error(w, "authorization header is not defined", http.StatusBadRequest)
			return
		}

		bearerToken := strings.Fields(authHeader)[1]

		if bearerToken == "" {
			log.Printf("bearer token is not defined")
			http.Error(w, "bearer token is not defined", http.StatusBadRequest)
			return
		}

		if bearerToken != s.config.Discord.Bot.Token {
			log.Printf("bearer token is not authorized")
			http.Error(w, "bearer token is not authorized", http.StatusUnauthorized)
			return
		}

		fn(w, r)
		return
	}
}
