package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"alles/boxes/api"
	"alles/boxes/env"
	"alles/boxes/google"
	"alles/boxes/receiver"
	"alles/boxes/store"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// connect to database
	conn, err := pgxpool.New(context.Background(), env.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}
	db := store.Store{Conn: conn}

	// router
	r := chi.NewRouter()
	r.Mount("/api", api.NewRouter(db))
	r.Post("/receive", receiver.Handler(db))

	// redirect to google auth
	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		refCookie, err := r.Cookie("ref")
		ref := ""
		if err == nil {
			ref = refCookie.Value
		}

		http.Redirect(w, r, google.GenerateUrl(ref), http.StatusTemporaryRedirect)
	})

	// start http server
	fmt.Println("starting http server on :3000")
	http.ListenAndServe(":3000", r)
}
