package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var DatabaseUrl = os.Getenv("DATABASE_URL")

var GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
var GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
var GoogleRedirect = os.Getenv("GOOGLE_REDIRECT")

var ReceiveSecret = os.Getenv("RECEIVE_SECRET")

var EmailDomain = os.Getenv("EMAIL_DOMAIN")
