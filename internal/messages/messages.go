package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/url"
	"time"
)

//{
//	"@context":"https://www.w3.org/ns/activitystreams",
//	"id":"https://mas.to/users/hvturingga#follows/113972/undo",
//	"type":"Undo",
//	"actor":"https://mas.to/users/hvturingga",
//	"object":{
//		"id":"https://mas.to/30ff54b1-c2dd-482c-ad70-43a775476584",
//		"type":"Follow","actor":"https://mas.to/users/hvturingga",
//		"object":"https://07ee-2408-832f-20b2-be60-7c3c-bb0d-7b8b-bb20.ngrok.io/u/hvturingga"
//	}
//}

type Reply struct {
	Context string `json:"@context"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Actor  string `json:"actor"`
		Object string `json:"object"`
	} `json:"object"`
}


type ActivityRequest struct {
	KeyName   string
	TargetURL string
	Local     string
	Data      []byte
	Key       []byte
}

var targetHost = "https://mas.to"

func (a *ActivityRequest) Follow () {
	nf := activitypub.NewFollow("Follow", "hvturingga")
	m := activitypub.Follow{
		Context: nf.Context,
		Id:      nf.Id,
		Type:    nf.Type,
		Actor:   nf.Actor,
		Object:  "https://mas.to/users/hvturingga",
	}

	data, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}

	acct := &accounts.Accounts{}
	if err2 := cockroach.GetDB().
		Debug().
		Table("accounts").
		Where("username = ?", "hvturingga").
		First(acct).Error; err2 != nil {
		log.Println(gorm.ErrMissingWhereClause)
	}

	host := fmt.Sprintf(viper.GetString("localhost"))

	req := ActivityRequest{
		KeyName:   fmt.Sprintf("%s#main-key", nf.Actor),
		TargetURL: fmt.Sprintf("%s/inbox", targetHost),
		Local:     host,
		Data:      data,
		Key:       []byte(acct.PrivateKey),
	}

	req.SendMessage()
}

//func (a *Message) Inbox(name string) {

	//fmt.Printf("%s 给 %s 发送了请求", i.Actor, name)
	//switch i.Type {
	//case "Follow":
	//	fmt.Printf("请求关注")
	//case "Undo":
	//	fmt.Printf("取消了请求")
	//case "Reject":
	//	fmt.Printf("拒绝了你的请求")
	//}
//}

type Message interface {
	Follow()
}


func (a *ActivityRequest) SendMessage() {
	//date := time.Now().UTC().Format(http.TimeFormat)
	//req.Header.Add("Host", "mstdn.social")
	//req.Header.Add("Date", date)
	//req.Header.Set("User-Agent", "hvxahv/0.0.1; "+ar.Local)
	//req.Header.Set("Content-Type", ContentType)

	h, err := url.Parse(a.TargetURL)
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("https://%s/inbox", h.Hostname())
	method := "POST"



	payload := bytes.NewBuffer(a.Data)
	client := &http.Client {
	}
	fmt.Println(payload)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}


	date := time.Now().UTC().Format(http.TimeFormat)

	req.Header.Add("Host", "mas.to")
	req.Header.Add("Date", date)

	//req.Header.Add("Signature", header)
	block := security.PrivateKey{
		Key: a.Key,
	}

	security.SignRequest(a.KeyName, block, req, a.Data)
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