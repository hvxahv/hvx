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
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/mailer"
	"github.com/spf13/viper"
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

// NewWebFingerAddress
// FOR EXAMPLE: https://mas.to/.well-known/webfinger?resource=acct:hvturingga@halfmemories.com
func NewWebFingerAddress(host, resource string) string {
	return fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s", host, resource)
}

// GetWebFingerHandler Obtain the WebFinger of the remote
// instance through the email address format.
func GetWebFingerHandler(address string) (*WebFinger, error) {
	format, err := mailer.ParseEmailAddress(address)
	if err != nil {
		return nil, err
	}
	g, err := resty.New().R().Get(NewWebFingerAddress(format.Domain, address))
	if err != nil {
		return nil, err
	}
	if g.StatusCode() != 200 {
		return nil, errors.New(errors.ErrWebfinger)
	}
	var w WebFinger
	if err := json.Unmarshal(g.Body(), &w); err != nil {
		return nil, err
	}
	return &w, nil
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
