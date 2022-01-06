package activity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/channel"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/spf13/viper"
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
	var objectID uint
	var actorID uint

	// Get ActorID by NAME.
	account, err := account.NewAccountsUsername(name).GetAccountByUsername()
	if err != nil {
		return
	}
	objectID = account.ActorID

	a := Activity{}
	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Printf("UNMARSHAL ACTICITY TYPE ERROR:%v", err)
	}

	//actor, err := account.NewActorUri(a.Actor).GetByActorUri()
	//if err != nil {
	//	resp, err2 := activitypub.GetRemoteActor(a.Actor)
	//	if err2 != nil {
	//		return
	//	}
	//	actorID = resp.ID
	//}
	//actorID = actor.ID

	switch a.Type {
	case "Follow":
		f := activitypub.Follow{}
		if err := json.Unmarshal(body, &f); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("请求关注")
		if err := NewFollowRequests(a.Id, actorID, objectID).Create(); err != nil {
			fmt.Println(err)
			return
		}

	case "Undo":
		undo := activitypub.Undo{}
		if err := json.Unmarshal(body, &undo); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("撤回了消息")
		if err := NewFollowRequestsActivityID(undo.Object.Id).Delete(); err != nil {
			fmt.Println(err)
			return
		}

	case "Accept":
		accept := activitypub.Accept{}
		if err := json.Unmarshal(body, &accept); err != nil {
			fmt.Println(err)
		}

		fmt.Println("接受了你的请求:")
		if err := NewFollowAccepts(accept.Id, actorID, objectID, accept.Object.Id).Create(); err != nil {
			fmt.Println(err)
			return
		}

	case "Reject":
		reject := activitypub.Reject{}
		if err := json.Unmarshal(body, &reject); err != nil {
			fmt.Println(err)
		}

		if reject.Object.Type == "Follow" {
			fmt.Println("移除了你的关注")
			if err := NewFollows(actorID, objectID).Remove(); err != nil {
				fmt.Println(err)
				return
			}
		}

		//case "Create":
		//
		//	fmt.Println("创建了一条消息")
		//	c := activitypub.Create{}
		//	if err := json.Unmarshal(body, &c); err != nil {
		//		log.Println(err)
		//	}
		//
		//	fmt.Println(string(body))
		//	fmt.Println("CONTEXT: ", c.Context)
		//	fmt.Println("ACTOR: ", c.Actor)
		//	fmt.Println("TYPE: ", c.Type)
		//	fmt.Println("ID: ", c.Id)
		//	fmt.Println("PUBLISHED: ", c.Published)
		//	fmt.Println("CC: ", c.Cc)
		//	fmt.Println("TO: ", c.To)
		//
		//	fmt.Println("OBJECT: ", c.Object)
		//
		//	fmt.Println("Id:", c.Object.Id)
		//	fmt.Println("Type:", c.Object.Type)
		//	fmt.Println("Summary:", c.Object.Summary)
		//	fmt.Println("InReplyTo:", c.Object.InReplyTo)
		//	fmt.Println("Url:", c.Object.Url)
		//	fmt.Println("AttributedTo:", c.Object.AttributedTo)
		//	fmt.Println("To:", c.Object.To)
		//	fmt.Println("Cc:", c.Object.Cc)
		//	fmt.Println("Sensitive:", c.Object.Sensitive)
		//	fmt.Println("AtomUri:", c.Object.AtomUri)
		//	fmt.Println("InReplyToAtomUri:", c.Object.InReplyToAtomUri)
		//	fmt.Println("Conversation:", c.Object.Conversation)
		//	fmt.Println("Content:", c.Object.Content)
		//	fmt.Println("InReplyToAtomUri:", c.Object.InReplyToAtomUri)
		//
		//	switch c.Object.Type {
		//	case "Note":
		//		fmt.Println("得到了一条 Note")
		//
		//		la := account.NewActorUrl(c.Actor)
		//		LID, err3 := la.FindActorByUrl()
		//		if err3 != nil {
		//			return
		//		}
		//
		//		to := map[string]interface{}{
		//			"to": c.Object.To,
		//		}
		//
		//		if len(c.Object.Cc) != 0 {
		//			fmt.Println("这是条消息提及：", c.Object.Cc)
		//
		//			if c.Object.InReplyTo != "" {
		//				fmt.Println("这是一条评论消息，对于：", c.Object.InReplyTo)
		//				nc := article.NewConversations(c.Id, LID.ID, c.Object.InReplyTo, c.Object.Content, remoteActor.ID)
		//				if err := nc.New(); err != nil {
		//					log.Println(err)
		//					return
		//				}
		//				return
		//			}
		//		}
		//
		//		n := article.Articles{
		//			AuthorID:   LID.ID,
		//			URL:        c.Object.Url,
		//			Article:    c.Object.Content,
		//			//Attachment: &article.Attachment{
		//			//	Attachment: c.Object.Attachment,
		//			//},
		//			TO: to,
		//			//CC:         &article.CC{CC: c.Object.Cc},
		//			Statuses:   true,
		//			NSFW:       false,
		//			Visibility: false,
		//		}
		//		if err := n.Create(); err != nil {
		//			log.Println(err)
		//			return
		//		}
		//
		//	}
		//case "Delete":
		//	fmt.Println("一个删除事件")
		//	fmt.Println(string(body))
		//	d := activitypub.Delete{}
		//	if err := json.Unmarshal(body, &d); err != nil {
		//		log.Println(err)
		//	}
		//	da := article.NewArticleURL(d.Object.Id)
		//	if err := da.DeleteByURL(); err != nil {
		//		log.Println(err)
		//		return
		//	}
		//}
	}
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

// ChannelTypes Get the request to subscribe to the channel and reply immediately.
func ChannelTypes(name string, body []byte) {
	var objectID uint
	var actorID uint

	// Get ActorID by NAME.
	ca, err2 := channel.NewChannelsByLink(name).GetActorDataByLink()
	if err2 != nil {
		return
	}

	objectID = ca.ID

	a := Activity{}
	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Printf("UNMARSHAL ACTICITY TYPE ERROR:%v", err)
	}

	actor, err := account.NewActorUri(a.Actor).GetByActorUri()
	fmt.Println("没找到错误", err)
	if err != nil {
		resp, err := activitypub.GetRemoteActor(a.Actor)
		if err != nil {
			return
		}
		actorID = resp.ID
	}
	actorID = actor.ID

	switch a.Type {
	case "Follow":
		f := activitypub.Follow{}
		if err := json.Unmarshal(body, &f); err != nil {
			fmt.Println(err)
			return
		}

		SubAccept(name, a.Id, actorID, objectID)
		// Todo - + channel subscriber.

	case "Reject":
		reject := activitypub.Reject{}
		if err := json.Unmarshal(body, &reject); err != nil {
			fmt.Println(err)
		}

		if reject.Object.Type == "Follow" {
			fmt.Println("取消订阅")
			// Todo - + Remove channel subscriber.
		}

	}
}

func SubAccept(name, activityId string, actorID, objectID uint) {

	fa, addr := NewSubAccept(name, activityId, actorID)
	data, err := json.Marshal(fa)
	if err != nil {
		return
	}

	fmt.Println(addr)

	if err := NewChannelAPData(name, addr, data).Send(); err != nil {
		return
	}
}

func NewSubAccept(actor, activityID string, objectID uint) (*activitypub.Accept, string) {
	o, err := account.NewActorID(objectID).GetByActorID()
	if err != nil {
		return nil, ""
	}

	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id  = fmt.Sprintf("https://%s/c/%s#accepts/follows/%s", viper.GetString("localhost"), actor, uuid.New().String())
		a   = fmt.Sprintf("https://%s/c/%s", viper.GetString("localhost"), actor)
	)

	return &activitypub.Accept{
		Context: ctx,
		Id:      id,
		Type:    "Accept",
		Actor:   a,
		Object: struct {
			Id     string `json:"id"`
			Type   string `json:"type"`
			Actor  string `json:"actor"`
			Object string `json:"object"`
		}{
			Id:     activityID,
			Type:   "Follow",
			Actor:  o.Url,
			Object: a,
		},
	}, o.Inbox
}
