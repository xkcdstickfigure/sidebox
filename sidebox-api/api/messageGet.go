package api

import (
	"net/http"
	"time"

	"alles/boxes/api/apierr"

	"github.com/go-chi/chi/v5"
)

// GET /message/{id}
func (h handlers) messageGet(w http.ResponseWriter, r *http.Request) {
	// session
	session, err := h.auth(r)
	if err != nil {
		apierr.Respond(w, apierr.BadAuthorization)
		return
	}

	// get message
	id := chi.URLParam(r, "id")
	message, err := h.db.MessageGet(r.Context(), id)
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// get inbox and check owner
	inbox, err := h.db.InboxGet(r.Context(), message.InboxId)
	if err != nil || inbox.AccountId != session.AccountId {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// response
	respond(w, struct {
		Id          string    `json:"id"`
		FromName    string    `json:"fromName"`
		FromAddress string    `json:"fromAddress"`
		Subject     string    `json:"subject"`
		Body        string    `json:"body"`
		Html        bool      `json:"html"`
		Date        time.Time `json:"date"`
	}{
		Id:          message.Id,
		FromName:    message.FromName,
		FromAddress: message.FromAddress,
		Subject:     message.Subject,
		Body:        message.Body,
		Html:        message.Html,
		Date:        message.Date,
	})
}
