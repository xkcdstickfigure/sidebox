package api

import (
	"net/http"
	"time"

	"alles/boxes/api/apierr"
	"alles/boxes/env"

	"github.com/go-chi/chi/v5"
)

// GET /inbox/{id}
func (h handlers) inboxGet(w http.ResponseWriter, r *http.Request) {
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

	// mark inbox as read
	h.db.InboxSetUnread(r.Context(), inbox.Id, false)

	// list messages
	messages, err := h.db.MessageList(r.Context(), inbox.Id)
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// response
	type resMessage struct {
		Id          string    `json:"id"`
		FromName    string    `json:"fromName"`
		FromAddress string    `json:"fromAddress"`
		Subject     string    `json:"subject"`
		Date        time.Time `json:"date"`
	}

	resMessages := []resMessage{}
	for _, m := range messages {
		resMessages = append(resMessages, resMessage{
			Id:          m.Id,
			FromName:    m.FromName,
			FromAddress: m.FromAddress,
			Subject:     m.Subject,
			Date:        m.Date,
		})
	}

	respond(w, struct {
		Id       string       `json:"id"`
		Name     string       `json:"name"`
		Address  string       `json:"address"`
		Muted    bool         `json:"muted"`
		Messages []resMessage `json:"messages"`
	}{
		Id:       inbox.Id,
		Name:     inbox.Name,
		Address:  inbox.Code + "@" + env.EmailDomain,
		Muted:    inbox.Muted,
		Messages: resMessages,
	})
}
