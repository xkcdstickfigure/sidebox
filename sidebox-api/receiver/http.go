package receiver

import (
	"net/http"

	"alles/boxes/env"
	"alles/boxes/store"
)

func Handler(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auth
		authHeader := r.Header.Get("authorization")
		if authHeader != env.ReceiveSecret {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// process
		err := process(r.Context(), db, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// response
		w.WriteHeader(204)
	}
}
