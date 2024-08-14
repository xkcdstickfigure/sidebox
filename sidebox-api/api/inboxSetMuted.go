package api

import (
	"encoding/json"
	"net/http"

	"alles/boxes/api/apierr"

	"github.com/go-chi/chi/v5"
)

// POST /inbox/{id}/muted
func (h handlers) inboxSetMuted(w http.ResponseWriter, r *http.Request) {
	// parse body
	var body struct {
		Muted bool
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
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
	h.db.InboxSetMuted(r.Context(), inbox.Id, body.Muted)

	// response
	respond(w, struct{}{})
}
