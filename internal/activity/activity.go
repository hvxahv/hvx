package activity

import (
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/url"
)

// All Activity Types inherit the properties of the base Activity type.
// Some specific Activity Types are subtypes or specializations of more generalized Activity Types
// (for instance, the Invite Activity Type is a more specific form of the Offer Activity Type).
// The Activity Types include:
// https://www.w3.org/TR/activitystreams-vocabulary/#activity-types

type Activity struct {
	Context string      `json:"@context"`
	Id      string      `json:"id"`
	Type    string      `json:"type"`
	Actor   string      `json:"actor"`
	Object  interface{} `json:"object"`
}

// Types handler for inbox activity.
func Types(name string, body []byte) {
	i := Activity{}

	if err := json.Unmarshal(body, &i); err != nil {
		fmt.Println(err)
	}

	u := accounts.NewActorUrl(i.Actor)
	actor, err := u.FindActorByUrl()
	if err != nil {
		remote, err2 := activitypub.FetchRemoteActor(i.Actor)
		if err2 != nil {
			return
		}
		actor = remote
	}

	switch i.Type {
	case "Follow":
		f := activitypub.Follow{}
		err2 := json.Unmarshal(body, &f)
		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Printf("%s 给 %s 发送了请求; 请求关注", f.Actor, f.Object)
		la := accounts.NewActorUrl(f.Object)
		LID, err3 := la.FindActorByUrl()
		if err3 != nil {
			return 
		}
		inbox, err := NewInbox(actor.ID, f.Type, f.Id, LID.ID)

		if err != nil {
			log.Println(err)
		}
		if err := inbox.New(); err != nil {
			log.Println(err)
		}

	case "Undo":
		fmt.Printf("取消了请求")
		fmt.Println("得到的接口数据:", i.Object)
		//nm := NewInbox(i.Actor, i.Type, i.Id, name)

	case "Reject":
		fmt.Printf("拒绝了你的请求")
		fmt.Println("接收了你的请求:", i.Object)
		//nm := NewInbox(i.Actor, i.Type, i.Id, name)

	case "Accept":
		fmt.Println("接受了你的请求:", i.Object)
		a := activitypub.Accept{}
		err2 := json.Unmarshal(body, &a)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(a)

	case "Create":
		fmt.Println("创建了一条消息")
		c := activitypub.Create{}
		if err := json.Unmarshal(body, &c); err != nil {
			log.Println(err)
		}
		fmt.Println("CONTEXT: ", c.Context)
		fmt.Println("ACTOR: ", c.Actor)
		fmt.Println("TYPE: ", c.Type)
		fmt.Println("ID: ", c.Id)
		fmt.Println("PUBLISHED: ", c.Published)
		fmt.Println("CC: ", c.Cc)
		fmt.Println("TO: ", c.To)

		fmt.Println("OBJECT: ", c.Object)

		fmt.Println("Id:", c.Object.Id)
		fmt.Println("Type:", c.Object.Type)
		fmt.Println("Summary:", c.Object.Summary)
		fmt.Println("InReplyTo:", c.Object.InReplyTo)
		fmt.Println("Url:", c.Object.Url)
		fmt.Println("AttributedTo:", c.Object.AttributedTo)
		fmt.Println("To:", c.Object.To)
		fmt.Println("Cc:", c.Object.Cc)
		fmt.Println("Sensitive:", c.Object.Sensitive)
		fmt.Println("AtomUri:", c.Object.AtomUri)
		fmt.Println("InReplyToAtomUri:", c.Object.InReplyToAtomUri)
		fmt.Println("Conversation:", c.Object.Conversation)
		fmt.Println("Content:", c.Object.Content)
		fmt.Println("InReplyToAtomUri:", c.Object.InReplyToAtomUri)

		switch c.Object.Type {
		case "Note":
			fmt.Println("得到了一条 Note")
		}

	}
}

//CONTEXT:  [https://www.w3.org/ns/activitystreams map[atomUri:ostatus:atomUri conversation:ostatus:conversation inReplyToAtomUri:ostatus:inReplyToAtomUri ostatus:http://ostatus.org# sensitive:as:sensitive toot:http://joinmastodon.org/ns# votersCount:toot:votersCount]]
//ACTOR:  https://mas.to/users/hvturingga
//TYPE:  Create
//ID:  https://mas.to/users/hvturingga/statuses/106947625962126298/activity
//PUBLISHED:  2021-09-17 15:07:45 +0000 UTC
//CC:  []
//TO:  [https://mas.to/users/hvturingga/followers]

/*
Id: https://mas.to/users/hvturingga/statuses/106947728146032819
Type: Note
Summary: <nil>
InReplyTo: <nil>
Url: https://mas.to/@hvturingga/106947728146032819
AttributedTo: https://mas.to/users/hvturingga
To: [https://mas.to/users/hvturingga/followers]
Cc: []
Sensitive: false
AtomUri: https://mas.to/users/hvturingga/statuses/106947728146032819
InReplyToAtomUri: <nil>
Conversation: tag:mas.to,2021-09-17:objectId=51323475:objectType=Conversation
Content: <p>你的名字叫做希望</p>
InReplyToAtomUri: <nil>
InReplyToAtomUri: <nil>
*/

type ActivityRequest struct {
	KeyID     string
	TargetURL string
	Local     string
	Data      []byte
	Key       []byte
}

type Request interface {
	// Send request to remote server.
	Send()

	// Follow ActivityPub follow method.
	Follow()

	// Accept ... TODO - Implement the method...
	Accept()
}

type InboxWithCtx struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}
type Receive interface {
	Inbox(name string)
}

func getPrivk() string {
	acct := &accounts.Accounts{}
	if err2 := cockroach.GetDB().
		Debug().
		Table("accounts").
		Where("username = ?", "hvturingga").
		First(acct).Error; err2 != nil {
		log.Println(gorm.ErrMissingWhereClause)
	}
	return acct.PrivateKey
}

// NewActivityRequest Receive the current actor name,
// the other party's URL,
// the requested data and the current user's private key.
func NewActivityRequest(actor string, object string, data []byte, key []byte) *ActivityRequest {
	h, err := url.Parse(object)
	if err != nil {
		log.Fatal(err)
	}

	targetURL := fmt.Sprintf("https://%s/inbox", h.Hostname())
	keyID := fmt.Sprintf("%s#main-key", actor)

	return &ActivityRequest{
		KeyID:     keyID,
		TargetURL: targetURL,
		Local:     fmt.Sprintf(viper.GetString("localhost")),
		Data:      data,
		Key:       key,
	}
}

func (a *ActivityRequest) Follow() {
	a.Send()
}

func (a *ActivityRequest) Accept() {
	a.Send()
}

func (a *ActivityRequest) Create() {
	a.Send()
}
