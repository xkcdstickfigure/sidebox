package api

import (
	"encoding/json"
	"net/http"

	"alles/boxes/api/apierr"

	"github.com/go-chi/chi/v5"
)

// POST /inbox/{id}/name
func (h handlers) inboxSetName(w http.ResponseWriter, r *http.Request) {
	// parse body
	var body struct {
		Name string
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.Name == "" {
		apierr.Respond(w, apierr.InvalidBody)
		return
	}

	// session
	session, err := h.auth(r)
	if err != nil {
		apierr.Respond(w, apierr.BadAuthorization)
		return
	}

	// get inbox and check owner
	id := chi.URLParam(r, "id")
	inbox, err := h.db.InboxGet(r.Context(), id)
	if err != nil || inbox.AccountId != session.AccountId {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// update inbox
	h.db.InboxSetName(r.Context(), inbox.Id, body.Name)

	// response
	respond(w, struct{}{})
}
