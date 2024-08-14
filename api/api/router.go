package api

import (
	"alles/boxes/store"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db store.Store) chi.Router {
	r := chi.NewRouter()
	h := handlers{db}

	r.Post("/login", h.login)
	r.Get("/account", h.account)
	r.Post("/inbox", h.inboxCreate)
	r.Get("/inbox/{id}", h.inboxGet)
	r.Post("/inbox/{id}/name", h.inboxSetName)
	r.Post("/inbox/{id}/muted", h.inboxSetMuted)
	r.Delete("/inbox/{id}", h.inboxDelete)
	r.Get("/message/{id}", h.messageGet)

	return r
}

type handlers struct {
	db store.Store
}
