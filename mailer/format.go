package mailer

import (
	"fmt"
	"github.com/hvxahv/hvx/errors"
	"strings"
)

type Format struct {
	Username string
	Domain   string
}

func ParseEmailAddress(email string) (*Format, error) {
	at := strings.LastIndex(email, "@")
	if at >= 0 {
		username, domain := email[:at], email[at+1:]

		return &Format{
			Username: username,
			Domain:   domain,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Error: %s is an invalid email address", email))
}
