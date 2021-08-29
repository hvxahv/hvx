package messages

import (
	"bytes"
	"fmt"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/httpsig"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/url"
	"time"
)

func NewFollow(actor, object string) *activitypub.Follow {
	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id = fmt.Sprintf("https://%s/%s", viper.GetString("localhost"), uuid.New().String())
		a = fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), actor)
	)
	
	return &activitypub.Follow{
		Context: ctx,
		Id:      id,
		Type:    "Follow",
		Actor:   a,
		Object:  object,
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

	block := httpsig.PriKEY{
		Type: httpsig.RSA,
		Key: a.Key,
	}

	ns := httpsig.NewSign(a.KeyID, block, req, a.Data)
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
