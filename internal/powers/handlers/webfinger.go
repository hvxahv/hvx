package handlers

import (
	"github.com/pkg/errors"
	"strings"
)

// getUser Get the username in the request url such,
// as "/.well-known/webfinger?resource=acct:hvturingga@0efb43b41a8a.ngrok.io" Will get hvturingga,
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

// WebFinger 和 WebFingerLinks 组成标准 Activitypub 的 JSON-LD 协议
type WebFinger struct {
	Subject string           `json:"subject"`
	Links   []WebFingerLinks `json:"links"`
}

// WebFingerLinks 供 WebFinger 使用
type WebFingerLinks struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

//
//func WebFingerHandler(c *gin.Context) {
//	res := c.Query("resource")
//	fmt.Println(res)
//	// Perform some filtering operations from the request to obtain the user name,
//	// and then search for the user name to find whether the user exists in the database.
//	// Currently only tested mastodon has not supported other ActivityPub implementations.
//	name, err := getUser(res)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	fmt.Println(name)
//	a := accounts.NewAccountByName(name)
//	query, err := a.Query()
//	if err != nil {
//		return
//	}
//	fmt.Println(query)
//
//	address := "ad907fce3836.ngrok.io"
//
//	links := []WebFingerLinks{
//		{
//			Rel:  "self",
//			Type: "application/activitypub+json",
//			Href: fmt.Sprintf("https://%s/u/%s", address, name),
//		},
//	}
//	finger := &WebFinger{
//		Subject: fmt.Sprintf("acct:%s@%s", name, address),
//		Links:   links,
//	}
//	log.Println(finger)
//	c.JSON(200, finger)
//}
