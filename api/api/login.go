package api

import (
	"encoding/json"
	"net/http"

	"alles/boxes/api/apierr"
	"alles/boxes/google"
)

// POST /login
func (h handlers) login(w http.ResponseWriter, r *http.Request) {
	// parse body
	var body struct {
		Code  string
		State string
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		apierr.Respond(w, apierr.InvalidBody)
		return
	}

	// get google profile from code
	profile, err := google.GetProfile(body.Code)
	if err != nil || !profile.EmailVerified {
		apierr.Respond(w, apierr.BadLogin)
		return
	}

	// create or get account with google id
	account, err := h.db.AccountCreate(r.Context(), profile.Name, profile.Email, profile.Id, body.State)
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// create session
	session, err := h.db.SessionCreate(r.Context(), account.Id, getAddress(r), r.Header.Get("user-agent"))
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// response
	respond(w, struct {
		Token string `json:"token"`
	}{
		Token: session.Token,
	})
}
