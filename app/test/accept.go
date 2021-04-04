package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	httpsig "hvxahv/pkg/activitypub"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func AcceptHandler(c *gin.Context) {
	idid := c.PostForm("id")
	url := "http://mas.to/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())


	obj := map[string]string {
		"id": fmt.Sprintf("http://mas.to/%s", idid),
		"type": "Follow",
		"actor": "http://mas.to/users/hvturingga",
		"object": "http://ba7b8a81471d.ngrok.io/u/hvturingga",
	}
	p := gin.H{
		"@context": "http://www.w3.org/ns/activitystreams",
		"id": fmt.Sprintf("http://%s/%s", address, idr),
		"type": "Accept",
		"actor": "http://ba7b8a81471d.ngrok.io/u/hvturingga",
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

	req.Header.Add("Host", "mas.to")
	req.Header.Add("Date", date)

	//req.Header.Add("Signature", header)

	block := httpsig.PrivateKey{
		Key: GetKey(),
	}
	httpsig.SignRequest("http://ba7b8a81471d.ngrok.io/u/hvturingga", block, req, byterData)
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
