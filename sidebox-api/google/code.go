package google

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"alles/boxes/env"

	"google.golang.org/api/idtoken"
)

type Token struct {
	Scope   string `json:"scope"`
	IdToken string `json:"id_token"`
}

type Profile struct {
	Id            string
	Name          string
	Email         string
	EmailVerified bool
}

func GetProfile(code string) (Profile, error) {
	// make token request
	values := url.Values{}
	values.Set("client_id", env.GoogleClientId)
	values.Set("client_secret", env.GoogleClientSecret)
	values.Set("redirect_uri", env.GoogleRedirect)
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)

	resp, err := http.PostForm("https://oauth2.googleapis.com/token", values)
	if err != nil {
		return Profile{}, err
	} else if resp.StatusCode != 200 {
		return Profile{}, errors.New("token request failed")
	}

	// parse token response
	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return Profile{}, err
	}

	// compare scope
	tokenScope := strings.Split(token.Scope, " ")
	sort.Strings(tokenScope)
	if strings.Join(tokenScope, " ") != scope {
		return Profile{}, errors.New("invalid scope")
	}

	// parse id token
	payload, err := idtoken.Validate(context.Background(), token.IdToken, env.GoogleClientId)
	if err != nil {
		return Profile{}, err
	}

	// return
	return Profile{
		Id:            payload.Claims["sub"].(string),
		Name:          payload.Claims["name"].(string),
		Email:         payload.Claims["email"].(string),
		EmailVerified: payload.Claims["email_verified"].(bool),
	}, nil
}
