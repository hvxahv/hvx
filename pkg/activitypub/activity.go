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

// https://www.w3.org/TR/activitypub/#create-activity-outbox

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

/**


 */

// Accept Indicates that the actor accepts the object.
// The target property can be used in certain circumstances to indicate the context into which the object
// has been accepted.
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

// Add Indicates that the actor has added the object to the target.
// If the target property is not explicitly specified, the target would need to be determined implicitly by context.
// The origin can be used to identify the context from which the object originated.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-add
type Add struct {
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Actor   struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"actor"`
	Object string `json:"object"`
}

// Announce Indicates that the actor is calling the target's attention the object.
// The origin typically has no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-announce
type Announce struct {
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Actor   struct {
		Type string `json:"type"`
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"actor"`
	Object struct {
		Type     string `json:"type"`
		Actor    string `json:"actor"`
		Location struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"location"`
	} `json:"object"`
}

// Block Indicates that the actor is blocking the object.
// Blocking is a stronger form of Ignore.
// The typical use is to support social systems that allow one user to block activities or content of other users.
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-block
type Block struct {
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

// Create Indicates that the actor has created the object.
// Object creation without a Create Activity
// https://www.w3.org/TR/activitypub/#create-activity-outbox
//type Create struct {
//	Context string `json:"@context"`
//	Type    string `json:"type"`
//	Id      string `json:"id"`
//	Actor   string `json:"actor"`
//	Object  struct {
//		Id           string    `json:"id"`
//		Type         string    `json:"type"`
//		AttributedTo string    `json:"attributedTo"`
//		Content      string    `json:"content"`
//		Published    time.Time `json:"published"`
//		To           []string  `json:"to"`
//		Cc           []string  `json:"cc"`
//	} `json:"object"`
//	Published time.Time `json:"published"`
//	To        []string  `json:"to"`
//	Cc        []string  `json:"cc"`
//}

// Delete Indicates that the actor has deleted the object.
// If specified, the origin indicates the context from which the object was deleted.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-delete
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
}

// Follow Indicates that the actor is "following" the object.
// Following is defined in the sense typically used within Social systems in which the actor is interested in any activity performed by or on the object.
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-follow
//type Follow struct {
//	Context string `json:"@context"`
//	Summary string `json:"summary"`
//	Type    string `json:"type"`
//	Actor   struct {
//		Type string `json:"type"`
//		Name string `json:"name"`
//	} `json:"actor"`
//	Object struct {
//		Type string `json:"type"`
//		Name string `json:"name"`
//	} `json:"object"`
//}

// Flag Indicates that the actor is "flagging" the object.
// Flagging is defined in the sense common to many social platforms as reporting content as being inappropriate for any number of reasons.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-flag
type Flag struct {
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"object"`
}

// Ignore Indicates that the actor is ignoring the object.
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-ignore
type Ignore struct {
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Actor   struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"actor"`
	Object string `json:"object"`
}

// Reject Indicates that the actor is rejecting the object.
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-reject
type Reject struct {
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Actor   struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"actor"`
	Object struct {
		Type   string `json:"type"`
		Actor  string `json:"actor"`
		Object struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"object"`
	} `json:"object"`
}

// Undo  Indicates that the actor is undoing the object.
// In most cases, the object will be an Activity describing some previously performed action (for instance, a person may have previously "liked" an article but,
// for whatever reason, might choose to undo that like at some later point in time).
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-undo
type Undo struct {
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  struct {
		Type   string `json:"type"`
		Actor  string `json:"actor"`
		Object string `json:"object"`
		Target string `json:"target"`
	} `json:"object"`
}

// Update Indicates that the actor has updated the object.
// Note, however, that this vocabulary does not define a mechanism for describing the actual set of modifications made to object.
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-update
//type Update struct {
//	Context string `json:"@context"`
//	Summary string `json:"summary"`
//	Type    string `json:"type"`
//	Actor   struct {
//		Type string `json:"type"`
//		Name string `json:"name"`
//	} `json:"actor"`
//	Object string `json:"object"`
//}

type Create struct {
	Context   []interface{} `json:"@context"`
	Id        string        `json:"id"`
	Type      string        `json:"type"`
	Actor     string        `json:"actor"`
	Published time.Time     `json:"published"`
	To        []string      `json:"to"`
	Cc        []interface{} `json:"cc"`
	Object    struct {
		Id               string        `json:"id"`
		Type             string        `json:"type"`
		Summary          interface{}   `json:"summary"`
		InReplyTo        interface{}   `json:"inReplyTo"`
		Published        time.Time     `json:"published"`
		Url              string        `json:"url"`
		AttributedTo     string        `json:"attributedTo"`
		To               []string      `json:"to"`
		Cc               []interface{} `json:"cc"`
		Sensitive        bool          `json:"sensitive"`
		AtomUri          string        `json:"atomUri"`
		InReplyToAtomUri interface{}   `json:"inReplyToAtomUri"`
		Conversation     string        `json:"conversation"`
		Content          string        `json:"content"`
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
}

// Update Indicates that the actor has updated the object.
// Note, however, that this vocabulary does not define a mechanism for describing the actual set of modifications made to object.
// The target and origin typically have no defined meaning.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-update
type Update struct {
	Context []interface{} `json:"@context"`
	Id      string        `json:"id"`
	Type    string        `json:"type"`
	Actor   string        `json:"actor"`
	To      []string      `json:"to"`
	Object  struct {
		Id                        string    `json:"id"`
		Type                      string    `json:"type"`
		Following                 string    `json:"following"`
		Followers                 string    `json:"followers"`
		Inbox                     string    `json:"inbox"`
		Outbox                    string    `json:"outbox"`
		Featured                  string    `json:"featured"`
		FeaturedTags              string    `json:"featuredTags"`
		PreferredUsername         string    `json:"preferredUsername"`
		Name                      string    `json:"name"`
		Summary                   string    `json:"summary"`
		Url                       string    `json:"url"`
		ManuallyApprovesFollowers bool      `json:"manuallyApprovesFollowers"`
		Discoverable              bool      `json:"discoverable"`
		Published                 time.Time `json:"published"`
		Devices                   string    `json:"devices"`
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
	} `json:"object"`
	Signature struct {
		Type           string    `json:"type"`
		Creator        string    `json:"creator"`
		Created        time.Time `json:"created"`
		SignatureValue string    `json:"signatureValue"`
	} `json:"signature"`
}
