package activity

import (
	"bytes"
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/url"
	"time"
)

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

	nf := accounts.NewFollows(actor, name)
	err := nf.New()
	if err != nil {
		return nil
	}

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
	url := fmt.Sprintf("https://%s/inbox", h.Hostname())
	method := "POST"

	payload := bytes.NewBuffer(a.Data)
	client := &http.Client{}
	fmt.Println(payload)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	date := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Add("Host", h.Hostname())
	req.Header.Add("Date", date)
	req.Header.Set("User-Agent", fmt.Sprintf("hvxahv/%s; %s", viper.GetString("version"), a.Local))
	req.Header.Set("Content-Type", "application/activity+json")

	block := security.PriKEY{
		Type: security.RSA,
		Key:  a.Key,
	}

	ns := security.NewSign(a.KeyID, block, req, a.Data)
	ns.SignRequest()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()
	switch res.StatusCode {
	case 200:
	case 201:
	case 202:
	default:
		fmt.Errorf("http post status: %d", res.StatusCode)
	}
	log.Printf("successful post: %s %d", url, res.StatusCode)
}
