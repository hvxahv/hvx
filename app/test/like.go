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

func LikeReq(da string, pk []byte) {

	url := "https://mas.to/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())

	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",

		"id": "https://4e54ea0be52f.ngrok.io/create-" + idr,
		"type": "Like",
		"actor": "https://4e54ea0be52f.ngrok.io/actor",
		"object": "https://mas.to/@hvturingga/105515218721573733",
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
		fmt.Errorf("http post status: %d", res.StatusCode)
	}
	log.Printf("successful post: %s %d", url, res.StatusCode)
	log.Print(res)
}