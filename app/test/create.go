package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"hvxahv/api/server/httputils"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func CreateHandler(c *gin.Context) {
	url := "server://mas.to/inbox"
	method := "POST"


	idr := strconv.Itoa(rand.Int())

	date := time.Now().UTC().Format(http.TimeFormat)
	obj := gin.H {
		"id": fmt.Sprintf("server://%s/%s", address, idr),
		"type": "Note",
		"published": date,
		"attributedTo": fmt.Sprintf("server://%s/actor", address),
		"content": "这是一条测试数据",
		"to": []string{"server://www.w3.org/ns/activitystreams#Public"},
	}

	p := gin.H{
		"@context": "server://www.w3.org/ns/activitystreams",
		"id": fmt.Sprintf("server://%s/create-%s", address, idr),
		"type": "Create",
		"actor": fmt.Sprintf("server://%s/actor", address),
		"to": []string{"server://www.w3.org/ns/activitystreams#Public"},
		"cc": []string{"server://mas.to/users/hvturingga"},
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


	req.Header.Add("Host", "mas.to")
	req.Header.Add("Date", date)

	//req.Header.Add("Signature", header)

	block := httputils.PrivateKey{
		Key: GetKey(),
	}
	httputils.SignRequest(fmt.Sprintf("server://%s/actor", address), block, req, byterData)
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
		fmt.Errorf("server post status: %d", res.StatusCode)
	}
	log.Printf("successful post: %s %d", url, res.StatusCode)
	log.Println(req)
}
