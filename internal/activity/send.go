package activity

import (
	"bytes"
	"fmt"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/channels"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/url"
	"time"
)

type APData struct {
	ID         string
	Addr       string
	Data       []byte
	PrivateKey []byte
}
func NewChannelAPData(link, inbox string, data []byte) *APData {
	ca, err := channels.NewChannelsByLink(link).GetByLink()
	if err != nil {
		fmt.Println(err)
	}

	id := fmt.Sprintf("https://%s/c/%s#main-key", viper.GetString("localhost"), link)
	return &APData{
		ID:         id,
		Addr:       inbox,
		Data:       data,
		PrivateKey: []byte(ca.PrivateKey),
	}
}

// NewAPData Instantiate ActivityPub data and return formatted data for sending request.
func NewAPData(actor, inbox string, data []byte) *APData {
	acct, err := accounts.NewAccountsUsername(actor).GetAccountByUsername()
	if err != nil {
		return nil
	}

	id := fmt.Sprintf("https://%s/u/%s#main-key", viper.GetString("localhost"), actor)
	return &APData{
		ID:         id,
		Addr:       inbox,
		Data:       data,
		PrivateKey: []byte(acct.PrivateKey),
	}
}

func (a *APData) Send() error {
	h, err := url.Parse(a.Addr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	client := &http.Client{}

	req, err := http.NewRequest("POST", a.Addr, bytes.NewBuffer(a.Data))
	if err != nil {
		fmt.Println(err)
	}

	ua := fmt.Sprintf("hvxahv/%s; %s", viper.GetString("version"), viper.GetString("localhost"))

	date := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Add("Host", h.Hostname())
	req.Header.Add("Date", date)
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/activity+json")

	block := activitypub.PriKEY{
		Type: activitypub.RSA,
		Key:  a.PrivateKey,
	}

	ns := activitypub.NewSign(a.ID, block, req, a.Data)

	req.Header.Set("Signature", ns.SignRequest())

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := res.Body.Close(); err != nil {
		return err
	}
	switch res.StatusCode {
	case 200:
	case 201:
	case 202:
	default:
		_ = fmt.Errorf("http post status: %d", res.StatusCode)
	}

	log.Printf("successful post: %s %d", a.Addr, res.StatusCode)
	return nil
}

type Send interface {
	Send() error
}