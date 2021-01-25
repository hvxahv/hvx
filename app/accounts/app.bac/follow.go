package app_bac

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

func Follow(c *gin.Context) {
	url := "https://mas.to/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())
	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"id": fmt.Sprintf("https://%s/%s", address, idr),
		"type": "Follow",
		"actor": fmt.Sprintf("https://%s/actor", address),
		"object": "https://mas.to/users/hvturingga",
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


type UndoObj struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

func Undo(c *gin.Context) {
	url := "https://mas.to/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())
	obj := UndoObj{
		ID: fmt.Sprintf("https://%s/%s", address, idr),
		Type:  "Follow",
		Actor: fmt.Sprintf("https://%s/actor", address),
		Object: "https://mas.to/users/hvturingga",
	}
	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"id": fmt.Sprintf("https://%s/%s", address, idr),
		"type": "Undo",
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

	req.Header.Add("Host", "mas.to")
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
