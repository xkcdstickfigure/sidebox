package email

import (
	"regexp"
)

func ValidateAddress(str string) bool {
	return regexp.MustCompile(`^[a-z0-9._=+\-]+@[a-z0-9.\-]+\.[a-z]{2,12}$`).MatchString(str)
}
