package handlers

import (
	"net/http"

	"github.com/diegoalzate/jot/cmd/ui"
)

func (h *Handlers) ViewHome(w http.ResponseWriter, r *http.Request) {
	ui.HomePage(struct{}{}).Render(r.Context(), w)
	return
}
