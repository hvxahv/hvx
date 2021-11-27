package activity

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/url"
	"time"
)


type ActivityRequest struct {
	KeyID     string
	inboxAddr string
	Local     string
	Data      []byte
	Key       []byte
}

// NewActivityRequest Receive the current actor name,
// the other party's URL,
// the requested data and the current user's private key.
func NewActivityRequest(actor, inbox string, data []byte) *ActivityRequest {
	acct, err := accounts.NewAccountsUsername(actor).GetAccountByUsername()
	if err != nil {
		return nil
	}

	keyID := fmt.Sprintf("https://%s/u/%s#main-key", viper.GetString("localhost"), actor)

	return &ActivityRequest{
		KeyID:     keyID,
		inboxAddr: inbox,
		Local:     fmt.Sprintf(viper.GetString("localhost")),
		Data:      data,
		Key:       []byte(acct.PrivateKey),
	}
}

func (a *ActivityRequest) Send() error {
	h, err := url.Parse(a.inboxAddr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	method := "POST"

	payload := bytes.NewBuffer(a.Data)
	client := &http.Client{}

	fmt.Println(payload)

	req, err := http.NewRequest(method, a.inboxAddr, payload)
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
	req.Header.Set("Signature", ns.SignRequest())
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
	log.Printf("successful post: %s %d", a.inboxAddr, res.StatusCode)
	return nil
}

func (a *ActivityRequest) Resty() error {
	h, err := url.Parse(a.inboxAddr)
	if err != nil {
		return err
	}

	resty := resty.New()

	R := resty.R().
		SetHeader("Host", h.Hostname()).
		SetHeader("Date", time.Now().UTC().Format(http.TimeFormat)).
		SetHeader("Content-Type", "application/activity+json").
		SetHeader("User-Agent", fmt.Sprintf("hvxahv/%s; %s", viper.GetString("version"), a.Local))

	block := activitypub.PriKEY{
		Type: activitypub.RSA,
		Key:  a.Key,
	}

	ns := activitypub.NewSign(a.KeyID, block, R.RawRequest, a.Data)
	sign := ns.SignRequest()

	fmt.Println("加密完成")
	resp, err := R.
		SetHeader("Signature", sign).
		Post(a.inboxAddr)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}