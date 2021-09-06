package activitypub

import (
	"time"
)

// Create Activity.
// https://www.w3.org/TR/activitypub/#create-activity-outbox
// {
//  "@context": "https://www.w3.org/ns/activitystreams",
//  "type": "Create",
//  "id": "https://example.net/~mallory/87374",
//  "actor": "https://example.net/~mallory",
//  "object": {
//    "id": "https://example.com/~mallory/note/72",
//    "type": "Note",
//    "attributedTo": "https://example.net/~mallory",
//    "content": "This is a note",
//    "published": "2015-02-10T15:04:55Z",
//    "to": ["https://example.org/~john/"],
//    "cc": ["https://example.com/~erik/followers",
//           "https://www.w3.org/ns/activitystreams#Public"]
//  },
//  "published": "2015-02-10T15:04:55Z",
//  "to": ["https://example.org/~john/"],
//  "cc": ["https://example.com/~erik/followers",
//         "https://www.w3.org/ns/activitystreams#Public"]
// }

// Activity ...
type Activity struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	Id      string `json:"id"`
	Actor   string `json:"actor"`
	Object  struct {
		Id           string    `json:"id"`
		Type         string    `json:"type"`
		AttributedTo string    `json:"attributedTo"`
		Content      string    `json:"content"`
		Published    time.Time `json:"published"`
		To           []string  `json:"to"`
		Cc           []string  `json:"cc"`
	} `json:"object"`
	Published time.Time `json:"published"`
	To        []string  `json:"to"`
	Cc        []string  `json:"cc"`
}

// Delete Activity.
// {
//  "@context": "https://www.w3.org/ns/activitystreams",
//  "id": "https://example.com/~alice/note/72",
//  "type": "Tombstone",
//  "published": "2015-02-10T15:04:55Z",
//  "updated": "2015-02-10T15:04:55Z",
//  "deleted": "2015-02-10T15:04:55Z"
// }
//

// ActivityDel ...
type ActivityDel struct {
	Context   string    `json:"@context"`
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
	Deleted   time.Time `json:"deleted"`
}

// Follow
//{
//	"@context":"https://www.w3.org/ns/activitystreams",
//	"id":"https://mas.to/e27a4e0e-a0a0-400e-a395-6b0e60f08291",
//	"type":"Follow",
//	"actor":"https://mas.to/users/hvturingga",
//	"object":"https://07ee-2408-832f-20b2-be60-7c3c-bb0d-7b8b-bb20.ngrok.io/u/hvturingga"
//}

type Follow struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

// Reply
//{
//	"@context":"https://www.w3.org/ns/activitystreams",
//	"id":"https://mas.to/users/hvturingga#follows/113972/undo",
//	"type":"Undo",
//	"actor":"https://mas.to/users/hvturingga",
//	"object":{
//		"id":"https://mas.to/30ff54b1-c2dd-482c-ad70-43a775476584",
//		"type":"Follow","actor":"https://mas.to/users/hvturingga",
//		"object":"https://07ee-2408-832f-20b2-be60-7c3c-bb0d-7b8b-bb20.ngrok.io/u/hvturingga"
//	}
//}

type Reply struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Actor  string `json:"actor"`
		Object string `json:"object"`
	} `json:"object"`
}

type Object struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}
