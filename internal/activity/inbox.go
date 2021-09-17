package activity

import (
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/activitypub"
	"gorm.io/gorm"
	"log"
)

type InboxData struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

type Inbox struct {
	gorm.Model
	Actor     string `gorm:"type:varchar(999);actor"`
	EventType string `gorm:"type:varchar(999);event_type"`
	EventID   string `gorm:"type:varchar(999);event_id"`
	Username  string `gorm:"primaryKey;type:varchar(999);username"`
}

func (i *Inbox) FetchInbox() {
	inbox, err := FetchInboxCollectionByName(i.Username)
	if err != nil {
		return 
	}
	fmt.Println(inbox)

}

func (i *Inbox) NewInbox() {
	if err := NewInboxToDB(i); err != nil {
		return 
	}
}

type INBOX interface {
	// NewInbox inbox data.
	NewInbox()

	// FetchInbox The inbox is discovered through the property of an actor's profile.
	// The MUST be an OrderedCollection.
	FetchInbox()
}

func NewInbox(actor string, types string, eventID string, username string) *Inbox {
	return &Inbox{Actor: actor, EventType: types, EventID: eventID, Username: username}
}

func NewInboxByName(username string) *Inbox {
	return &Inbox{Username: username}
}

func InboxEventHandler(name string, body []byte) {
	i := InboxData{}

	err := json.Unmarshal(body, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	fmt.Println("TYPE: ", i.Type)
	fmt.Printf("%s 给 %s 发送了请求; 请求 ID : %s", i.Actor, name, i.Id)

	switch i.Type {
	case "Follow":
		f := activitypub.Follow{}
		err2 := json.Unmarshal(body, &f)
		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println(f)
		fmt.Printf("%s 请求关注你", f.Actor)

		nm := NewInbox(f.Actor, f.Type, f.Id, name)
		nm.NewInbox()

	case "Undo":
		fmt.Printf("取消了请求")
		fmt.Println("得到的接口数据:", i.Object)
		nm := NewInbox(i.Actor, i.Type, i.Id, name)
		nm.NewInbox()

	case "Reject":
		fmt.Printf("拒绝了你的请求")
		fmt.Println("接收了你的请求:", i.Object)
		nm := NewInbox(i.Actor, i.Type, i.Id, name)
		nm.NewInbox()

	case "Accept":
		fmt.Println("接受了你的请求:", i.Object)
		a := activitypub.Accept{}
		err2 := json.Unmarshal(body, &a)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(a)

		nm := NewInbox(a.Actor, a.Type, a.Id, name)
		nm.NewInbox()
		
		nf := accounts.NewFollows(name, i.Actor)
		err3 := nf.New()
		if err3 != nil {
			log.Println(err3)
		}

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
