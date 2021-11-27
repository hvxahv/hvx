package activity

import (
	"encoding/json"
	"fmt"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
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
	account, err := accounts.NewAccountsUsername(name).GetAccountByUsername()
	if err != nil {
		return
	}
	objectID = account.ActorID

	a := Activity{}
	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Printf("UNMARSHAL ACTICITY TYPE ERROR:%v", err)
	}

	actor, err := accounts.NewActorUri(a.Actor).GetActorByUri()
	if err != nil {
		resp, err2 := activitypub.GetRemoteActor(a.Actor)
		if err2 != nil {
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
		//		la := accounts.NewActorUrl(c.Actor)
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
		//				nc := articles.NewConversations(c.Id, LID.ID, c.Object.InReplyTo, c.Object.Content, remoteActor.ID)
		//				if err := nc.New(); err != nil {
		//					log.Println(err)
		//					return
		//				}
		//				return
		//			}
		//		}
		//
		//		n := articles.Articles{
		//			AuthorID:   LID.ID,
		//			URL:        c.Object.Url,
		//			Article:    c.Object.Content,
		//			//Attachment: &articles.Attachment{
		//			//	Attachment: c.Object.Attachment,
		//			//},
		//			TO: to,
		//			//CC:         &articles.CC{CC: c.Object.Cc},
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
		//	da := articles.NewArticleURL(d.Object.Id)
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

