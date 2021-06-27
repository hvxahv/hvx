package activitypub

import (
	"github.com/pkg/errors"
	"strings"
)

// GetActor Get the username in the request url such,
// as "/.well-known/webfinger?resource=acct:hvturingga@0efb43b41a8a.ngrok.io" Will get hvturingga,
// If the match fails, it will return a custom username not found error.
func GetActor(resource string) (string, error) {
	if strings.HasPrefix(resource, "acct:") {
		resource = resource[5:]
		if ali := strings.IndexByte(resource, '@'); ali != -1 {
			resource = resource[:ali]
		}
	} else {
		return "", errors.New("Failed to get username.")
	}

	return resource, nil
}
