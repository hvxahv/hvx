package v1alpha1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"hvxahv/internal/accounts"
	"log"
	"strings"
)

// getUser Get the username in the request url such,
// as "/.well-known/webfinger?resource=acct:hvturingga@0efb43b41a8a.ngrok.io"
// Will get hvturingga,
// If the match fails, it will return a custom username not found error.
func getUser(resource string) (string, error) {
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

func WebFingerHandler(c *gin.Context) {
	res := c.Query("resource")
	// Perform some filtering operations from the request to obtain the user name,
	// and then search for the user name to find whether the user exists in the database.
	// Currently only tested mastodon has not supported other ActivityPub implementations.
	name, err := getUser(res)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(name)
	a := accounts.NewAccountByName(name)
	query, err := a.Query()
	if err != nil {
		return 
	}
	fmt.Println(query)
	//c.JSON(200, gin.H{
	//	"response": query,
	//})
}
