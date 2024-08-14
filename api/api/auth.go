package api

import (
	"errors"
	"net/http"

	"alles/boxes/store"
)

func (h handlers) auth(r *http.Request) (store.Session, error) {
	token := r.Header.Get("authorization")
	if token == "" {
		return store.Session{}, errors.New("no session token")
	}

	return h.db.SessionGetByToken(r.Context(), token)
}
