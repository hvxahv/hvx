package activity

import (
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/internal/articles"
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
	fmt.Println(string(body))

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
		// 请求关注。
		//{
		//	"@context":"https://www.w3.org/ns/activitystreams",
		//	"id":"https://mas.to/ad263d19-ed74-46d1-8827-ffa6ea2cb893",
		//	"type":"Follow",
		//	"actor":"https://mas.to/users/hvturingga",
		//	"object":"https://ec45-2408-832f-20b4-3480-c58f-a0f9-96e6-54eb.ngrok.io/u/hvturingga"
		//}
		f := activitypub.Follow{}
		err2 := json.Unmarshal(body, &f)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println("请求关注")
		la := accounts.NewActorUrl(f.Actor)
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
		// 撤回了请求。
		//	{
		//		"@context":"https://www.w3.org/ns/activitystreams",
		//		"id":"https://mas.to/users/hvturingga#follows/120713/undo",
		//		"type":"Undo",
		//		"actor":"https://mas.to/users/hvturingga",
		//		"object":{
		//			"id":"https://mas.to/ad263d19-ed74-46d1-8827-ffa6ea2cb893",
		//			"type":"Follow","actor":"https://mas.to/users/hvturingga",
		//			"object":"https://ec45-2408-832f-20b4-3480-c58f-a0f9-96e6-54eb.ngrok.io/u/hvturingga"
		//		}
		//	}
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
		fmt.Printf("拒绝了你的请求")
		fmt.Println("接收了你的请求:", i.Object)
		//nm := NewInbox(i.Actor, i.Type, i.Id, name)

	case "Accept":
		// 同意请求。
		//	{
		//		"@context":"https://www.w3.org/ns/activitystreams",
		//		"id":"https://mas.to/users/hvturingga#accepts/follows/120712",
		//		"type":"Accept",
		//		"actor":"https://mas.to/users/hvturingga",
		//		"object":{
		//			"id":"https://ec45-2408-832f-20b4-3480-c58f-a0f9-96e6-54eb.ngrok.io/3c338626-4588-440d-a3de-13b5e5918bd6",
		//			"type":"Follow",
		//			"actor":"https://ec45-2408-832f-20b4-3480-c58f-a0f9-96e6-54eb.ngrok.io/u/hvturingga",
		//			"object":"https://mas.to/users/hvturingga"
		//		}
		//	}
		fmt.Println("接受了你的请求:", i.Object)
		a := activitypub.Accept{}
		if err := json.Unmarshal(body, &a); err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))

		la := accounts.NewActorUrl(a.Object.Actor)
		LID, err3 := la.FindActorByUrl()
		if err3 != nil {
			return
		}

		// Following...
		nf := accounts.NewFollows(LID.ID, actor.ID)
		if err := nf.New(); err != nil {
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

			to := map[string]interface{} {
				"to": c.Object.To,
			}

			n := articles.Articles{
				ActivityID: c.Id,
				AuthorID:   LID.ID,
				URL:        c.Object.Url,
				Article:    c.Object.Content,
				//Attachment: &articles.Attachment{
				//	Attachment: c.Object.Attachment,
				//},
				TO:         to,
				//CC:         &articles.CC{CC: c.Object.Cc},
				Statuses:   true,
				NSFW:       false,
				Visibility: false,
			}
			if err := n.New(); err != nil {
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

// 推送给用户的通知

// 提及
// {"@context":["https://www.w3.org/ns/activitystreams",{"ostatus":"http://ostatus.org#","atomUri":"ostatus:atomUri","inReplyToAtomUri":"ostatus:inReplyToAtomUri","conversation":"ostatus:conversation","sensitive":"as:sensitive","toot":"http://joinmastodon.org/ns#","votersCount":"toot:votersCount"}],"id":"https://mas.to/users/hvturingga/statuses/107122649066661445/activity","type":"Create","actor":"https://mas.to/users/hvturingga","published":"2021-10-18T12:58:25Z","to":["https://mas.to/users/hvturingga/followers"],"cc":["https://fb09-175-162-3-122.ngrok.io/u/hvturingga"],"object":{"id":"https://mas.to/users/hvturingga/statuses/107122649066661445","type":"Note","summary":null,"inReplyTo":null,"published":"2021-10-18T12:58:25Z","url":"https://mas.to/@hvturingga/107122649066661445","attributedTo":"https://mas.to/users/hvturingga","to":["https://mas.to/users/hvturingga/followers"],"cc":["https://fb09-175-162-3-122.ngrok.io/u/hvturingga"],"sensitive":false,"atomUri":"https://mas.to/users/hvturingga/statuses/107122649066661445","inReplyToAtomUri":null,"conversation":"tag:mas.to,2021-10-18:objectId=53402191:objectType=Conversation","content":"<p><span class=\"h-card\"><a href=\"https://fb09-175-162-3-122.ngrok.io/u/hvturingga\" class=\"u-url mention\">@<span>hvturingga@fb09-175-162-3-122.ngrok.io</span></a></span> 悲しみの忘れ方。</p>","contentMap":{"ja":"<p><span class=\"h-card\"><a href=\"https://fb09-175-162-3-122.ngrok.io/u/hvturingga\" class=\"u-url mention\">@<span>hvturingga@fb09-175-162-3-122.ngrok.io</span></a></span> 悲しみの忘れ方。</p>"},"attachment":[],"tag":[{"type":"Mention","href":"https://fb09-175-162-3-122.ngrok.io/u/hvturingga","name":"@hvturingga@fb09-175-162-3-122.ngrok.io"}],"replies":{"id":"https://mas.to/users/hvturingga/statuses/107122649066661445/replies","type":"Collection","first":{"type":"CollectionPage","next":"https://mas.to/users/hvturingga/statuses/107122649066661445/replies?only_other_accounts=true&page=true","partOf":"https://mas.to/users/hvturingga/statuses/107122649066661445/replies","items":[]}}}}

// 评论
// {"@context":["https://www.w3.org/ns/activitystreams",{"ostatus":"http://ostatus.org#","atomUri":"ostatus:atomUri","inReplyToAtomUri":"ostatus:inReplyToAtomUri","conversation":"ostatus:conversation","sensitive":"as:sensitive","toot":"http://joinmastodon.org/ns#","votersCount":"toot:votersCount"}],"id":"https://mas.to/users/hvturingga/statuses/107122822497703019/activity","type":"Create","actor":"https://mas.to/users/hvturingga","published":"2021-10-18T13:42:32Z","to":["https://mas.to/users/hvturingga/followers"],"cc":["https://fb09-175-162-3-122.ngrok.io/u/hvturingga"],"object":{"id":"https://mas.to/users/hvturingga/statuses/107122822497703019","type":"Note","summary":null,"inReplyTo":"https://fb09-175-162-3-122.ngrok.io/u/hvturingga/article/698898124205195265","published":"2021-10-18T13:42:32Z","url":"https://mas.to/@hvturingga/107122822497703019","attributedTo":"https://mas.to/users/hvturingga","to":["https://mas.to/users/hvturingga/followers"],"cc":["https://fb09-175-162-3-122.ngrok.io/u/hvturingga"],"sensitive":false,"atomUri":"https://mas.to/users/hvturingga/statuses/107122822497703019","inReplyToAtomUri":"https://fb09-175-162-3-122.ngrok.io/u/hvturingga/article/698898124205195265","conversation":"tag:mas.to,2021-10-04:objectId=52441016:objectType=Conversation","content":"<p><span class=\"h-card\"><a href=\"https://fb09-175-162-3-122.ngrok.io/u/hvturingga\" class=\"u-url mention\">@<span>hvturingga@fb09-175-162-3-122.ngrok.io</span></a></span> HAHA</p>","contentMap":{"zhCn":"<p><span class=\"h-card\"><a href=\"https://fb09-175-162-3-122.ngrok.io/u/hvturingga\" class=\"u-url mention\">@<span>hvturingga@fb09-175-162-3-122.ngrok.io</span></a></span> HAHA</p>"},"attachment":[],"tag":[{"type":"Mention","href":"https://fb09-175-162-3-122.ngrok.io/u/hvturingga","name":"@hvturingga@fb09-175-162-3-122.ngrok.io"}],"replies":{"id":"https://mas.to/users/hvturingga/statuses/107122822497703019/replies","type":"Collection","first":{"type":"CollectionPage","next":"https://mas.to/users/hvturingga/statuses/107122822497703019/replies?only_other_accounts=true&page=true","partOf":"https://mas.to/users/hvturingga/statuses/107122822497703019/replies","items":[]}}}}

