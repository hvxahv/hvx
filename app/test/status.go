package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	httpsig "hvxahv/pkg/activity"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ContentMap struct {
	Zh string `json:"zh"`
}
type First struct {
	Type   string        `json:"type"`
	Next   string        `json:"next"`
	PartOf string        `json:"partOf"`
	Items  []interface{} `json:"items"`
}
type Replies struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	First `json:"first"`
}

type StatusObj struct {
		ID               string      `json:"id"`
		Type             string      `json:"type"`
		Summary          interface{} `json:"summary"`
		InReplyTo        interface{} `json:"inReplyTo"`
		Published        time.Time     `json:"published"`
		URL              string        `json:"url"`
		AttributedTo     string        `json:"attributedTo"`
		To               []string      `json:"to"`
		Cc               []string      `json:"cc"`
		Sensitive        bool          `json:"sensitive"`
		AtomURI          string        `json:"atomUri"`
		InReplyToAtomURI interface{}   `json:"inReplyToAtomUri"`
		Conversation     string        `json:"conversation"`
		Content          string        `json:"content"`
		ContentMap       ContentMap    `json:"contentMap"`
		Attachment       []interface{} `json:"attachment"`
		Tag              []interface{} `json:"tag"`
		Replies          Replies       `json:"replies"`
}

func Status(c *gin.Context) {
	url := fmt.Sprintf("https://%s/inbox", address)
	method := "POST"

	idr := strconv.Itoa(rand.Int())


	obj := StatusObj{
		ID:         fmt.Sprintf("https://%s/users/hvturingga/statuses/111/activity", address),
		Type:       "Create",
		Summary:    fmt.Sprintf("https://%s/users/hvturingga", address),
		Published:  time.Now(),
		To:         []string{"https://www.w3.org/ns/activitystreams#Public"},
		Cc:         []string{fmt.Sprintf("https://%s/users/hvturingga/followers", address)},
		Sensitive:  false,
		Content:    "<p>这是一条测试消息</p>",
		ContentMap: ContentMap{Zh: "<p>我发送了一条测试消息</p>"},
		Replies: Replies{
			ID: fmt.Sprintf("https://%s/users/hvturingga/statuses/111/replies", address),
			Type: "Collection",
			First: First{
				Type: "CollectionPage",
			},
		},

	}
	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"id": fmt.Sprintf("https://%s/%s", address, idr),
		"type": "Follow",
		"actor": fmt.Sprintf("https://%s/actor", address),
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

	req.Header.Add("Host", address)
	req.Header.Add("Date", date)

	//req.Header.Add("Signature", header)
	block := httpsig.PrivateKey{
		Key: GetKey(),
	}
	httpsig.SignRequest(fmt.Sprintf("https://%s/actor", address), block, req, byterData)
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
	log.Println(req)
}
