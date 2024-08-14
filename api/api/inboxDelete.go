package api

import (
	"net/http"

	"alles/boxes/api/apierr"

	"github.com/go-chi/chi/v5"
)

// DELETE /inbox/{id}
func (h handlers) inboxDelete(w http.ResponseWriter, r *http.Request) {
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

	// delete inbox
	err = h.db.InboxDelete(r.Context(), inbox.Id)
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// response
	respond(w, struct {}{})
}
