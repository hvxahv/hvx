package activitypub

import "time"

// Follow Indicates that the actor is "following" the object.
// Following is defined in the sense typically used within
// Social systems in which the actor is interested in any
// activity performed by or on the object. The target and origin
// typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-follow
type Follow struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

// Undo Indicates that the actor is undoing the object.
// In most cases, the object will be an Activity describing
// some previously performed action (for instance, a person
// may have previously "liked" an article but, for whatever
// reason, might choose to undo that like at some later point in time).
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-undo
type Undo struct {
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

// Accept Indicates that the actor accepts the object. The target
// property can be used in certain circumstances to indicate the
// context into which the object has been accepted.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-accept
type Accept struct {
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

// Reject Indicates that the actor is rejecting the object.
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-reject
type Reject struct {
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

func NewContext() []interface{} {
	return []interface{}{
		"https://www.w3.org/ns/activitystreams",
	}
}

// Create Indicates that the actor has created the object.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-create
type Create struct {
	Context   []interface{} `json:"@context"`
	Id        string        `json:"id"`
	Type      string        `json:"type"`
	Actor     string        `json:"actor"`
	Published time.Time     `json:"published"`
	To        []string      `json:"to"`
	Cc        []string      `json:"cc"`
	Object    struct {
		Id               string      `json:"id"`
		Type             string      `json:"type"`
		Title            string      `json:"title"`
		Summary          interface{} `json:"summary"`
		InReplyTo        interface{} `json:"inReplyTo"`
		Published        time.Time   `json:"published"`
		Url              string      `json:"url"`
		AttributedTo     string      `json:"attributedTo"`
		To               []string    `json:"to"`
		Cc               []string    `json:"cc"`
		Sensitive        bool        `json:"sensitive"`
		AtomUri          string      `json:"atomUri"`
		InReplyToAtomUri interface{} `json:"inReplyToAtomUri"`
		Conversation     string      `json:"conversation"`
		Content          string      `json:"content"`
		ContentMap       struct {
			En string `json:"en"`
		} `json:"contentMap"`
		Attachment []interface{} `json:"attachment"`
		Tag        []interface{} `json:"tag"`
		Replies    struct {
			Id    string `json:"id"`
			Type  string `json:"type"`
			First struct {
				Type   string        `json:"type"`
				Next   string        `json:"next"`
				PartOf string        `json:"partOf"`
				Items  []interface{} `json:"items"`
			} `json:"first"`
		} `json:"replies"`
	} `json:"object"`
	Signature struct {
		Type           string    `json:"type"`
		Creator        string    `json:"creator"`
		Created        time.Time `json:"created"`
		SignatureValue string    `json:"signatureValue"`
	} `json:"signature"`
}

type Delete struct {
	Context []interface{} `json:"@context"`
	Id      string        `json:"id"`
	Type    string        `json:"type"`
	Actor   string        `json:"actor"`
	To      []string      `json:"to"`
	Object  struct {
		Id      string `json:"id"`
		Type    string `json:"type"`
		AtomUri string `json:"atomUri"`
	} `json:"object"`
	Signature struct {
		Type           string    `json:"type"`
		Creator        string    `json:"creator"`
		Created        time.Time `json:"created"`
		SignatureValue string    `json:"signatureValue"`
	} `json:"signature"`
}
