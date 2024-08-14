package google

import (
	"net/url"

	"alles/boxes/env"
)

const scope = "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile openid"

func GenerateUrl(state string) string {
	values := url.Values{}
	values.Set("client_id", env.GoogleClientId)
	values.Set("redirect_uri", env.GoogleRedirect)
	values.Set("response_type", "code")
	values.Set("scope", scope)
	values.Set("state", state)

	return "https://accounts.google.com/o/oauth2/v2/auth?" + values.Encode()
}
