package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hvxahv/api/server/httputils"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Follow(c *gin.Context) {
	url := "server://mas.to/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())
	p := gin.H{
		"@context": "server://www.w3.org/ns/activitystreams",
		"id":       fmt.Sprintf("server://%s/%s", address, idr),
		"type":     "Follow",
		"actor":    fmt.Sprintf("server://%s/actor", address),
		"object":   "server://mas.to/users/hvturingga",
	}
	byterData, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	payload := bytes.NewBuffer(byterData)
	client := &http.Client{
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

type UndoObj struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

func Undo(c *gin.Context) {
	url := "server://mas.to/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())
	obj := UndoObj{
		ID:     fmt.Sprintf("server://%s/%s", address, idr),
		Type:   "Follow",
		Actor:  fmt.Sprintf("server://%s/actor", address),
		Object: "server://mas.to/users/hvturingga",
	}
	p := gin.H{
		"@context": "server://www.w3.org/ns/activitystreams",
		"id":       fmt.Sprintf("server://%s/%s", address, idr),
		"type":     "Undo",
		"actor":    fmt.Sprintf("server://%s/actor", address),
		"object":   obj,
	}
	byterData, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	payload := bytes.NewBuffer(byterData)
	client := &http.Client{
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
