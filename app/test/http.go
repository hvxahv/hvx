package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	httpsig "hvxahv/pkg/activitypub"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Req(da string, pk []byte) {

	url := "https://mstdn.social/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())
	obj := map[string]string {
		"id": "https://4e54ea0be52f.ngrok.io/"+ idr,
		"type": "Note",
		"published": time.Now().UTC().Format(http.TimeFormat),
		"attributedTo": "https://4e54ea0be52f.ngrok.io/actor",
		"inReplyTo": "https://mstdn.social/@hvturingga/105515197741965407",
		"content": fmt.Sprintf("<p>Hello %s world</p>", idr),
		"to": "https://www.w3.org/ns/activitystreams#Public",
	}

	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",

		"id": "https://4e54ea0be52f.ngrok.io/create-" + idr,
		"type": "Create",
		"actor": "https://4e54ea0be52f.ngrok.io/actor",
		"object": obj,
	}
	byterData, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}


	payload := bytes.NewBuffer(byterData)
	client := &http.Client {
	}
	fmt.Println(payload)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}


	date := time.Now().UTC().Format(http.TimeFormat)

	req.Header.Add("Host", "mstdn.social")
	req.Header.Add("Date", date)

	//req.Header.Add("Signature", header)
	block := httpsig.PrivateKey{
		Key: pk,
	}
	httpsig.SignRequest("https://4e54ea0be52f.ngrok.io/actor", block, req, byterData)
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
		fmt.Errorf("https post status: %d", res.StatusCode)
	}
	log.Printf("successful post: %s %d", url, res.StatusCode)
	log.Println("请求出现错误",req)
}


// 发送私信
func Req2(da string, pk []byte) {

	url := "https://mstdn.social/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())
	//obj := map[string]string {
	//	"id": "https://activitypub.disism.com/"+ idr,
	//	"type": "Note",
	//	"published": time.Now().UTC().Format(https.TimeFormat),
	//	"attributedTo": "https://activitypub.disism.com/actor",
	//	"inReplyTo": "https://mastodon.social/@hvturingga/104812740119120055",
	//	//"content": da,
	//	"to": "https://www.w3.org/ns/activitystreams#Like",
	//}

	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",

		"id": "https://services.disism.com/create-" + idr,
		"type": "Like",
		"actor": "https://services.disism.com/actor",
		"object": "https://mastodon.social/@hvturingga/104812740119120055",
	}
	byterData, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}


	payload := bytes.NewBuffer(byterData)
	client := &http.Client {
	}
	fmt.Println(payload)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}


	date := time.Now().UTC().Format(http.TimeFormat)

	req.Header.Add("Host", "mastodon.social")
	req.Header.Add("Date", date)

	//req.Header.Add("Signature", header)
	block := httpsig.PrivateKey{
		Key: pk,
	}
	httpsig.SignRequest("https://services.disism.com/actor", block, req, byterData)
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
		fmt.Errorf("https post status: %d", res.StatusCode)
	}
	log.Printf("successful post: %s %d", url, res.StatusCode)
}
