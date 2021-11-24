package activitypub

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"log"
)

// https://webfinger.net.
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

type WebFingerLinks struct {
	Rel      string `json:"rel"`
	Type     string `json:"type,omitempty"`
	Href     string `json:"href,omitempty"`
	Template string `json:"template,omitempty"`
}

type WebFingerData struct {
	Subject string           `json:"subject"`
	Aliases []string         `json:"aliases"`
	Links   []WebFingerLinks `json:"links"`
}

// NewWebFingerUrl Returns a standard remote webFinger query url.
// for example: https://mas.to/.well-known/webfinger?resource=hvturingga
func NewWebFingerUrl(host, resource string) string {
	return fmt.Sprintf("https://%s/.well-known/webfinger?resource=%s", host, resource)
}

// NewWebFingerData  WebFinger data and links form the JSON-LD protocol of the standard ActivityPub.
func NewWebFingerData(subject, address, name string) *WebFingerData {
	return &WebFingerData{
		Subject: subject,
		Aliases: nil,
		Links:   NewWebFingerLinks(address, name),
	}
}

func NewWebFingerLinks(address, name string) []WebFingerLinks {
	return []WebFingerLinks{{
		Rel:      "self",
		Type:     "application/activity+json",
		Href:     fmt.Sprintf("https://%s/u/%s", address, name),
		Template: "",
	}}
}

// NewWebFinger Receive the username and combine the object into the standard json data returned by webFinger.
// In order to return to the query that the instance owns this actor.
func NewWebFinger(name string) *WebFingerData {
	address := viper.GetString("localhost")
	sub := fmt.Sprintf("acct:%s@%s", name, address)

	wf := NewWebFingerData(sub, address, name)

	return wf
}

// GetWebFinger Obtain the WebFinger of the remote instance through the email address format.
func GetWebFinger(resource string) *WebFingerData {
	w, err := resty.New().R().Get(NewWebFingerUrl(GetHost(resource), resource))
	if err != nil {
		log.Fatal(err)
	}
	var wf WebFingerData
	if err := json.Unmarshal(w.Body(), &wf); err != nil {
		return nil
	}
	return &wf
}