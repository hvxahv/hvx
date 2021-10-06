package activity

import (
	"bytes"
	"fmt"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func NewDelete(articleID uint) *activitypub.Delete {
	actor := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), "hvturingga")
	id := fmt.Sprintf("%s/article/%s", actor, strconv.Itoa(int(articleID)))
	to := []string{"https://mas.to/users/hvturingga/"}

	return &activitypub.Delete{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      fmt.Sprintf("%s#delete", id),
		Type:    "Delete",
		Actor:   actor,
		To:      to,
		Object: struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			AtomUri string `json:"atomUri"`
		}{
			Id:      id,
			Type:    "Tombstone",
			AtomUri: id,
		},
	}
}

func NewArticle(articleID uint, content string) *activitypub.Create {
	actor := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), "hvturingga")
	id := fmt.Sprintf("%s/article/%s", actor, strconv.Itoa(int(articleID)))
	to := []string{"https://mas.to/users/hvturingga/"}

	return &activitypub.Create{
		Context:   "https://www.w3.org/ns/activitystreams",
		Id:        fmt.Sprintf("%s/activity", id),
		Type:      "Create",
		Actor:     actor,
		Published: time.Time{},
		To:        to,
		Cc:        nil,
		Object: struct {
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
			Attachment       []interface{} `json:"attachment"`
			Tag              []interface{} `json:"tag"`
			Replies          struct {
				Id    string `json:"id"`
				Type  string `json:"type"`
				First struct {
					Type   string        `json:"type"`
					Next   string        `json:"next"`
					PartOf string        `json:"partOf"`
					Items  []interface{} `json:"items"`
				} `json:"first"`
			} `json:"replies"`
		}{
			Id:               id,
			Type:             "Note",
			Summary:          nil,
			InReplyTo:        nil,
			Published:        time.Time{},
			Url:              id,
			AttributedTo:     "",
			To:               to,
			Cc:               nil,
			Sensitive:        false,
			AtomUri:          "",
			InReplyToAtomUri: nil,
			Conversation:     "",
			Content:          content,
			Attachment:       nil,
			Tag:              nil,
			Replies: struct {
				Id    string `json:"id"`
				Type  string `json:"type"`
				First struct {
					Type   string        `json:"type"`
					Next   string        `json:"next"`
					PartOf string        `json:"partOf"`
					Items  []interface{} `json:"items"`
				} `json:"first"`
			}{},
		},
	}
}

func NewFollowRequest(actor, object string) *activitypub.Follow {
	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id  = fmt.Sprintf("https://%s/%s", viper.GetString("localhost"), uuid.New().String())
		a   = fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), actor)
	)

	return &activitypub.Follow{
		Context: ctx,
		Id:      id,
		Type:    "Follow",
		Actor:   a,
		Object:  object,
	}
}

// NewFollowAccept
// name: LOCAL ACTOR NAME,
// actor: REMOTE ACTOR LINK,
// oid: CONTEXT ID,
// object: LOCAL ACTOR LINK.
func NewFollowAccept(name, actor, oid string) *activitypub.Accept {
	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id  = fmt.Sprintf("https://%s/u/%s#accepts/follows/%s", viper.GetString("localhost"), name, uuid.New().String())
		a   = fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)
	)

	//nf := accounts.NewFollower(name, acct.ID)
	//if err := nf.New(); err != nil {
	//	log.Println(err)
	//}

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
			Id:     oid,
			Type:   "Follow",
			Actor:  actor,
			Object: a,
		},
	}
}

func (a *ActivityRequest) Send() {
	h, err := url.Parse(a.TargetURL)
	if err != nil {
		log.Fatal(err)
	}

	uri := fmt.Sprintf("https://%s/users/hvturingga/inbox", h.Hostname())
	method := "POST"

	payload := bytes.NewBuffer(a.Data)
	client := &http.Client{}

	fmt.Println(payload)

	req, err := http.NewRequest(method, uri, payload)
	if err != nil {
		fmt.Println(err)
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Add("Host", h.Hostname())
	req.Header.Add("Date", date)
	req.Header.Set("User-Agent", fmt.Sprintf("hvxahv/%s; %s", viper.GetString("version"), a.Local))
	req.Header.Set("Content-Type", "application/activity+json")

	block := activitypub.PriKEY{
		Type: activitypub.RSA,
		Key:  a.Key,
	}

	ns := activitypub.NewSign(a.KeyID, block, req, a.Data)
	ns.SignRequest()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if err := res.Body.Close(); err != nil {
		log.Println(err)
	}
	switch res.StatusCode {
	case 200:
	case 201:
	case 202:
	default:
		_ = fmt.Errorf("http post status: %d", res.StatusCode)
	}
	log.Printf("successful post: %s %d", uri, res.StatusCode)
}
