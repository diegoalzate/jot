package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/diegoalzate/jot/internal/query"
	"github.com/google/uuid"
)

type CreateDiscordServerRequest struct {
	DiscordID string `json:"discord_id"`
	Name      string `json:"name"`
}

func (h *HandlerContext) CreateDiscordServer(w http.ResponseWriter, r *http.Request) {
	var inputServer CreateDiscordServerRequest

	q := query.New(h.db.Conn)

	err := json.NewDecoder(r.Body).Decode(&inputServer)

	if err != nil {
		log.Print("failed to decode discord server")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.NewV7()

	if err != nil {
		log.Printf("[ERR]: %#v", err.Error())
		http.Error(w, "failed to create uuid", http.StatusInternalServerError)
		return
	}

	newServer, err := q.CreateDiscordServer(r.Context(), query.CreateDiscordServerParams{
		ID:        id,
		DiscordID: inputServer.DiscordID,
		Name:      inputServer.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	log.Printf("created server: %v", newServer)
	return
}
