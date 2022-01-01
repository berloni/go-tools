package tools

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

// ValidateEmail checks if a given email is valid or not
func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 && len(email) > 254 {
		return errors.New("invalid email")
	}
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email")
	}

	// check MX
	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return errors.New("invalid email")
	}

	return nil
}
