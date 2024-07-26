package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/diegoalzate/jot/internal/query"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
)

func (h *Web) OauthCallback(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	// logs out of gothic store
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("[ERR]: %#v", err)
		http.Error(w, "failed to complete authentication callback", http.StatusInternalServerError)
		return
	}
	// app session management
	err = h.session.RenewToken(r.Context())
	if err != nil {
		http.Error(w, "failed to renew token", http.StatusInternalServerError)
		return
	}

	h.session.Put(r.Context(), "discord_token", gothUser.AccessToken)

	// check if user exists
	q := query.New(h.db.Conn)

	dbIdentity, err := q.GetIdentity(context.Background(), query.GetIdentityParams{
		ProviderID: gothUser.UserID,
		Provider:   gothUser.Provider,
	})

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[ERR]: %#v", err)
		http.Error(w, "failed to get identity", http.StatusInternalServerError)
		return
	}

	if err == sql.ErrNoRows {
		// create user
		id, err := uuid.NewV7()

		if err != nil {
			log.Printf("[ERR]: %#v", err)
			http.Error(w, "failed to create uuid", http.StatusInternalServerError)
			return
		}

		newUser, err := q.CreateUser(
			context.Background(),
			query.CreateUserParams{
				ID:        id,
				Username:  gothUser.NickName,
				Email:     gothUser.Email,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		)

		if err != nil {
			log.Printf("[ERR]: %#v", err)
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		id, err = uuid.NewV7()

		if err != nil {
			log.Printf("[ERR]: %#v", err)
			http.Error(w, "failed to create uuid", http.StatusInternalServerError)
			return
		}

		rawData, err := json.Marshal(gothUser.RawData)

		if err != nil {
			log.Printf("[ERR]: %#v", err)
			http.Error(w, "failed to convert raw data to bytes", http.StatusInternalServerError)
			return
		}

		_, err = q.CreateIdentity(context.Background(), query.CreateIdentityParams{
			ID:           id,
			UserID:       newUser.ID,
			Provider:     gothUser.Provider,
			ProviderID:   gothUser.UserID,
			IdentityData: rawData,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		log.Printf("[SUCCESS]:user: %#v", newUser)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	dbUser, err := q.GetUserById(context.Background(), dbIdentity.UserID)

	if err != nil {
		log.Printf("[ERR]: %#v", err)
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	// redirect home
	h.session.Put(r.Context(), "user_id", dbUser.ID.String())
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	return
}

func (h *Web) Logout(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	gothic.Logout(w, r)
	h.session.Destroy(r.Context())
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Web) Login(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.Println("[ALREADY LOGGED IN]", gothUser)
	} else {
		url, err := gothic.GetAuthURL(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add("HX-Redirect", url)
	}
}
