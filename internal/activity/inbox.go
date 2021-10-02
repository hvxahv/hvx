package activity

import (
	"encoding/json"
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
)

// TODO - Redesign the INBOX data structure.






type InboxData struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

type Inboxes struct {
	gorm.Model

	Actor     string `gorm:"type:text;actor"`
	EventType string `gorm:"type:text;event_type"`
	EventID   string `gorm:"type:text;event_id"`
	AccountID uint   `gorm:"primaryKey;type:bigint;account_id"`
}

func (i *Inboxes) FindInboxByAccountID() (*[]Inboxes, error) {
	db := cockroach.GetDB()

	var inboxes []Inboxes
	if err := db.Debug().Table("inboxes").Where("account_id", i.AccountID).Find(&inboxes).Error; err != nil {
		return nil, errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return &inboxes, nil
}

func (i *Inboxes) New() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("inboxes").Create(&i).Error; err != nil {
		return errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return nil
}

type Inbox interface {
	New() error

	FindInboxByAccountID() (*[]Inboxes, error)
}

func NewInbox(actor, types, eventID, username string) (*Inboxes, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_INBOX_DATABASE")
	}

	cli, conn, err := client.Accounts()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	account, err := cli.FindAccountsByUsername(context.Background(), &pb.AccountUsername{Username: username})
	if err != nil {
		return nil, err
	}

	return &Inboxes{Actor: actor, EventType: types, EventID: eventID, AccountID: uint(account.Id)}, nil
}

func NewInboxAccountID(id uint) *Inboxes {
	return &Inboxes{AccountID: id}
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

		inbox, err := NewInbox(f.Actor, f.Type, f.Id, name)
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
