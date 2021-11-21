package activity

import (
	"encoding/json"
	"fmt"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/articles"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
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
	// Get local actor ID
	acctName := accounts.NewAccountsName(name)
	localActor, err3 := acctName.FindAccountByUsername()
	if err3 != nil {
		return
	}

	fmt.Println(string(body))

	i := Activity{}

	if err := json.Unmarshal(body, &i); err != nil {
		fmt.Printf("UNMARSHAL ACTICITY TYPE ERROR:%v", err)
	}

	u := accounts.NewActorUrl(i.Actor)
	remoteActor, err := u.FindActorByUrl()
	if err != nil {
		remote, err2 := activitypub.FetchRemoteActor(i.Actor)
		if err2 != nil {
			return
		}
		remoteActor = remote
	}

	switch i.Type {
	case "Follow":
		f := activitypub.Follow{}
		err2 := json.Unmarshal(body, &f)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println("请求关注")

		if err := NewFollow(i.Id, "Follow", remoteActor.ID, localActor.ActorID).New(); err != nil {
			log.Println(err)
			return
		}

	case "Undo":
		fmt.Printf("撤回了消息")
		fmt.Println("得到的接口数据:", i.Object)
		fmt.Println(string(body))
		undo := activitypub.Undo{}
		err2 := json.Unmarshal(body, &undo)
		if err2 != nil {
			fmt.Println(err2)
		}

		d := NewInboxesActivityID(undo.Object.Id)
		if err := d.Delete(); err != nil {
			log.Println(err)
			return
		}
		fmt.Println("删除消息成功")

	case "Reject":
		reject := activitypub.Reject{}
		err2 := json.Unmarshal(body, &reject)
		if err2 != nil {
			fmt.Println(err2)
		}

		if reject.Object.Type == "Follow" {
			fmt.Println("移除了你的关注")
			if err := accounts.NewFollows(localActor.ActorID, remoteActor.ID).Remove(); err != nil {
				log.Println(err)
				return
			}
		}

	case "Accept":
		fmt.Println("接受了你的请求:", i.Object)
		a := activitypub.Accept{}
		if err := json.Unmarshal(body, &a); err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))

		// Following...
		nf := accounts.NewFollows(localActor.ActorID, remoteActor.ID)
		if err := nf.New(); err != nil {
			return
		}
		if err := NewAccept(i.Id, "Accept", remoteActor.ID, localActor.ActorID, a.Object.Id).New(); err != nil {
			log.Println(err)
			return
		}

	case "Create":

		fmt.Println("创建了一条消息")
		c := activitypub.Create{}
		if err := json.Unmarshal(body, &c); err != nil {
			log.Println(err)
		}

		fmt.Println(string(body))
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

			la := accounts.NewActorUrl(c.Actor)
			LID, err3 := la.FindActorByUrl()
			if err3 != nil {
				return
			}

			to := map[string]interface{}{
				"to": c.Object.To,
			}

			if len(c.Object.Cc) != 0 {
				fmt.Println("这是条消息提及：", c.Object.Cc)

				if c.Object.InReplyTo != "" {
					fmt.Println("这是一条评论消息，对于：", c.Object.InReplyTo)
					nc := articles.NewConversations(c.Id, LID.ID, c.Object.InReplyTo, c.Object.Content, remoteActor.ID)
					if err := nc.New(); err != nil {
						log.Println(err)
						return 
					}
					return
				}
			}

			n := articles.Articles{
				AuthorID:   LID.ID,
				URL:        c.Object.Url,
				Article:    c.Object.Content,
				//Attachment: &articles.Attachment{
				//	Attachment: c.Object.Attachment,
				//},
				TO: to,
				//CC:         &articles.CC{CC: c.Object.Cc},
				Statuses:   true,
				NSFW:       false,
				Visibility: false,
			}
			if err := n.Create(); err != nil {
				log.Println(err)
				return
			}

		}
	case "Delete":
		fmt.Println("一个删除事件")
		fmt.Println(string(body))
		d := activitypub.Delete{}
		if err := json.Unmarshal(body, &d); err != nil {
			log.Println(err)
		}
		da := articles.NewArticleURL(d.Object.Id)
		if err := da.DeleteByURL(); err != nil {
			log.Println(err)
			return
		}
	}
}

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

	Create()

	Article()
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

func (a *ActivityRequest) Article() {
	a.Send()
}
