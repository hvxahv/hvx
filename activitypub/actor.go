/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package activitypub

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hvxahv/hvx/errors"

	"github.com/spf13/viper"
)

// EXAMPLE 9
//{
//  "@context": ["https://www.w3.org/ns/activitystreams",
//               {"@language": "ja"}],
//  "type": "Person",
//  "id": "https://kenzoishii.example.com/",
//  "following": "https://kenzoishii.example.com/following.json",
//  "followers": "https://kenzoishii.example.com/followers.json",
//  "liked": "https://kenzoishii.example.com/liked.json",
//  "inbox": "https://kenzoishii.example.com/inbox.json",
//  "outbox": "https://kenzoishii.example.com/feed.json",
//  "preferredUsername": "kenzoishii",
//  "name": "石井健蔵",
//  "summary": "この方はただの例です",
//  "icon": [
//    "https://kenzoishii.example.com/image/165987aklre4"
//  ]
//}

//{
//    "@context": [
//        "https://www.w3.org/ns/activitystreams",
//        "https://w3id.org/security/v1",
//        {
//            "manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
//            "toot": "http://joinmastodon.org/ns#",
//            "featured": {
//                "@id": "toot:featured",
//                "@type": "@id"
//            },
//            "featuredTags": {
//                "@id": "toot:featuredTags",
//                "@type": "@id"
//            },
//            "alsoKnownAs": {
//                "@id": "as:alsoKnownAs",
//                "@type": "@id"
//            },
//            "movedTo": {
//                "@id": "as:movedTo",
//                "@type": "@id"
//            },
//            "schema": "http://schema.org#",
//            "PropertyValue": "schema:PropertyValue",
//            "value": "schema:value",
//            "IdentityProof": "toot:IdentityProof",
//            "discoverable": "toot:discoverable",
//            "Device": "toot:Device",
//            "Ed25519Signature": "toot:Ed25519Signature",
//            "Ed25519Key": "toot:Ed25519Key",
//            "Curve25519Key": "toot:Curve25519Key",
//            "EncryptedMessage": "toot:EncryptedMessage",
//            "publicKeyBase64": "toot:publicKeyBase64",
//            "deviceId": "toot:deviceId",
//            "claim": {
//                "@type": "@id",
//                "@id": "toot:claim"
//            },
//            "fingerprintKey": {
//                "@type": "@id",
//                "@id": "toot:fingerprintKey"
//            },
//            "identityKey": {
//                "@type": "@id",
//                "@id": "toot:identityKey"
//            },
//            "devices": {
//                "@type": "@id",
//                "@id": "toot:devices"
//            },
//            "messageFranking": "toot:messageFranking",
//            "messageType": "toot:messageType",
//            "cipherText": "toot:cipherText",
//            "suspended": "toot:suspended",
//            "focalPoint": {
//                "@container": "@list",
//                "@id": "toot:focalPoint"
//            }
//        }
//    ],
//    "id": "https://mas.to/users/hvturingga",
//    "type": "Person",
//    "following": "https://mas.to/users/hvturingga/following",
//    "followers": "https://mas.to/users/hvturingga/followers",
//    "inbox": "https://mas.to/users/hvturingga/inbox",
//    "outbox": "https://mas.to/users/hvturingga/outbox",
//    "featured": "https://mas.to/users/hvturingga/collections/featured",
//    "featuredTags": "https://mas.to/users/hvturingga/collections/tags",
//    "preferredUsername": "hvturingga",
//    "name": "",
//    "summary": "<p></p>",
//    "url": "https://mas.to/@hvturingga",
//    "manuallyApprovesFollowers": true,
//    "discoverable": false,
//    "published": "2021-01-07T00:00:00Z",
//    "devices": "https://mas.to/users/hvturingga/collections/devices",
//    "publicKey": {
//        "id": "https://mas.to/users/hvturingga#main-key",
//        "owner": "https://mas.to/users/hvturingga",
//        "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAug/IRF3dYbJJod02ucJJ\n/HMQQQbpIY41HCpuy1AotRBCOTbsfaQ0/4KeSeRedLImV5KtfHwv5mqB/jSLWs6c\ndI7c7wOppQr5FKeYqhLcdHVDLNO56py6GehFLtBKmQv/aLjnxAMeXcWSJOker9DZ\nBr0PxUScIFMS3CF4pjklVX5yIaorP1QY+PEw7mHSCar1erR5As5BiuMKiWuVQepO\nHho5H/ulw9w0CXMmcRcb9lbMOHf8Agx7r3p8fyoktOqcDgaTbg1sHf6pjrgLkmNK\n3pKN2aouWBXoU4MmuqcImu7c+x+CX78x4d+2jrE48pQVmqo+caKN4SLqi0O2yMDO\nMQIDAQAB\n-----END PUBLIC KEY-----\n"
//    },
//    "tag": [],
//    "attachment": [],
//    "endpoints": {
//        "sharedInbox": "https://mas.to/inbox"
//    },
//    "icon": {
//        "type": "Image",
//        "mediaType": "image/jpeg",
//        "url": "https://media.mas.to/masto-public/accounts/avatars/000/233/556/original/b5cb0332006740ef.jpg"
//    }
//}

type Actor struct {
	Context                   []interface{} `json:"@context"`
	Id                        string        `json:"id"`
	Type                      string        `json:"type"`
	Following                 string        `json:"following"`
	Followers                 string        `json:"followers"`
	Inbox                     string        `json:"inbox"`
	Outbox                    string        `json:"outbox"`
	Featured                  string        `json:"featured"`
	FeaturedTags              string        `json:"featuredTags"`
	PreferredUsername         string        `json:"preferredUsername"`
	Name                      string        `json:"name"`
	Summary                   string        `json:"summary"`
	Url                       string        `json:"url"`
	ManuallyApprovesFollowers bool          `json:"manuallyApprovesFollowers"`
	Discoverable              bool          `json:"discoverable"`
	Published                 time.Time     `json:"published"`
	Devices                   string        `json:"devices"`
	PublicKey                 struct {
		Id           string `json:"id"`
		Owner        string `json:"owner"`
		PublicKeyPem string `json:"publicKeyPem"`
	} `json:"publicKey"`
	Tag        []interface{} `json:"tag"`
	Attachment []interface{} `json:"attachment"`
	Endpoints  struct {
		SharedInbox string `json:"sharedInbox"`
	} `json:"endpoints"`
	Icon struct {
		Type      string `json:"type"`
		MediaType string `json:"mediaType"`
		Url       string `json:"url"`
	} `json:"icon"`
}

// JSON LD
// https://www.w3.org/TR/2014/REC-json-ld-20140116/

func NewActorContext() []interface{} {
	arr := make([]interface{}, 0)
	ctx := []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"}
	for _, i := range ctx {
		arr = append(arr, i)
	}
	return arr
}

// NewActor Return standard ActivityPub protocol user data.
func NewActor(preferredUsername, name, summary, publicKey, avatar string) *Actor {
	var (
		addr = viper.GetString("domain")

		id  = fmt.Sprintf("https://%s/u/%s", addr, preferredUsername)
		kid = fmt.Sprintf("%s#main-key", id)
		box = fmt.Sprintf("https://%s/u/%s/", addr, preferredUsername)
	)

	actor := &Actor{
		Context:                   NewActorContext(),
		Id:                        id,
		Type:                      "Person",
		Following:                 "",
		Followers:                 "",
		Inbox:                     box + "inbox",
		Outbox:                    box + "outbox",
		Featured:                  "",
		FeaturedTags:              "",
		PreferredUsername:         preferredUsername,
		Name:                      name,
		Summary:                   summary,
		Url:                       "",
		ManuallyApprovesFollowers: false,
		Discoverable:              false,
		Published:                 time.Time{},
		Devices:                   "",
		PublicKey: struct {
			Id           string `json:"id"`
			Owner        string `json:"owner"`
			PublicKeyPem string `json:"publicKeyPem"`
		}{
			Id:           kid,
			Owner:        id,
			PublicKeyPem: publicKey,
		},
		Tag:        nil,
		Attachment: nil,
		Endpoints: struct {
			SharedInbox string `json:"sharedInbox"`
		}{},
		Icon: struct {
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
			Url       string `json:"url"`
		}{
			Type:      "Image",
			MediaType: "image/jpg",
			Url:       avatar,
		},
	}
	return actor
}

// GetActorName Get the username in the request url such,
// as "/.well-known/webFinger?resource=acct:hvturingga@halfmemories.com" Will get hvturingga,
// If the match fails, it will return a custom username not found error.
func GetActorName(resource string) string {
	if strings.HasPrefix(resource, "acct:") {
		resource = resource[5:]
		if ali := strings.IndexByte(resource, '@'); ali != -1 {
			resource = resource[:ali]
		}
	}

	return resource
}

// GetActorByWebfinger Get remote Actor via WebFinger.
func GetActorByWebfinger(wf *WebFinger) (*resty.Response, error) {
	g, err := resty.New().
		R().
		SetHeader("Content-Type", "application/activitypub+json; charset=utf-8").
		SetHeader("Accept", "application/ld+json").
		Get(wf.Links[0].Href)
	if err != nil {
		return nil, err
	}
	if g.StatusCode() != 200 {
		return nil, errors.New(errors.ErrWebfinger)
	}
	return g, nil
}

// IsRemote Get host to determine whether it is a remote
// instance user (not a user of this instance).
//func IsRemote(resource string) bool {
//	host := GetActorHost(resource)
//	if !strings.Contains(resource, "@") {
//		return false
//	}
//
//	if host != viper.GetString("localhost") {
//		return true
//	}
//	return false
//
//}

// GetActorHost Get the host of the received resource.
// hvturingga@halfmemories.com Will get halfmemories.com,
//func GetActorHost(resource string) string {
//	if strings.HasPrefix(resource, "acct:") {
//		ali := strings.IndexByte(resource, '@')
//		if ali != -1 {
//			return resource[ali+1:]
//		}
//	}
//
//	return resource[5:]
//}
