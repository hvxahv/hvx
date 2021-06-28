package activitypub

import (
	"fmt"
	"github.com/spf13/viper"
)

// WebFinger JSON example
// {
//    "subject": "acct:hvturingga@2cf915078b27.ngrok.io",
//    "links": [
//        {
//            "rel": "self",
//            "type": "application/activity+json",
//            "href": "https://2cf915078b27.ngrok.io/actor"
//        }
//    ]
// }

type webFinger struct {
	Subject string `json:"subject"`
	Links   *[]links `json:"links"`
}

type links struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

// NewWebFinger Receive the username and combine the object into the standard json data returned by webfinger.
// In order to return to the queryer that the instance owns this actor.
func NewWebFinger(name string) *webFinger {
	address := viper.GetString("localhost")
	sub := fmt.Sprintf("acct:%s@%s", name, address)

	l := &[]links{
		{
			Rel:  "self",
			Type: "application/activity+json",
			Href: fmt.Sprintf("https://%s/u/%s", address, name),
		},
	}
	return &webFinger{Subject: sub, Links: l}
}
