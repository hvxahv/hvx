/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package activitypub

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"log"
)

//type WebFingerLinks struct {
//	Rel      string `json:"rel"`
//	Type     string `json:"type,omitempty"`
//	Href     string `json:"href,omitempty"`
//	Template string `json:"template,omitempty"`
//}

type WebFinger struct {
	Subject string   `json:"subject"`
	Aliases []string `json:"aliases"`
	Links   []struct {
		Rel      string `json:"rel"`
		Type     string `json:"type,omitempty"`
		Href     string `json:"href,omitempty"`
		Template string `json:"template,omitempty"`
	} `json:"links"`
}

// NewWebFingerUrl Returns a standard remote webFinger query url.
// for example: https://mas.to/.well-known/webfinger?resource=hvturingga
func NewWebFingerUrl(host, resource string) string {
	return fmt.Sprintf("https://%s/.well-known/webfinger?resource=%s", host, resource)
}

// GetRemoteWebFinger Obtain the WebFinger of the remote
// instance through the email address format.
func GetRemoteWebFinger(resource string) *WebFinger {
	w, err := resty.New().R().Get(NewWebFingerUrl(GetHost(resource), resource))
	if err != nil {
		log.Fatal(err)
	}
	var wf WebFinger
	if err := json.Unmarshal(w.Body(), &wf); err != nil {
		return nil
	}
	return &wf
}

// NewWebFinger Receive the username and combine the object
// into the standard json data returned by webFinger. In order
// to return to the query that the instance owns this actor.
func NewWebFinger(name string, isChan bool) *WebFinger {
	address := viper.GetString("localhost")
	sub := fmt.Sprintf("acct:%s@%s", name, address)

	var href string
	if isChan {
		href = fmt.Sprintf("https://%s/c/%s", address, name)
	} else {
		href = fmt.Sprintf("https://%s/u/%s", address, name)
	}
	return &WebFinger{
		Subject: sub,
		Aliases: nil,
		Links: []struct {
			Rel      string `json:"rel"`
			Type     string `json:"type,omitempty"`
			Href     string `json:"href,omitempty"`
			Template string `json:"template,omitempty"`
		}{
			{
				Rel:      "self",
				Type:     "application/activity+json",
				Href:     href,
				Template: "",
			},
		},
	}
}
