package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(h.db.Health())
	_, _ = w.Write(jsonResp)
}
