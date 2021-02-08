package activity

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"hvxahv/api/client/account"
	"hvxahv/pkg/models"
	"log"
	"net/http"
	"net/url"
	"time"
)

// SendActivity 发送活动
func SendActivity(data *models.SendActivity) int {
	payload := bytes.NewBuffer(data.Data)
	cli := &http.Client {}

	req, err := http.NewRequest(data.Method, data.EndInbox, payload)
	if err != nil {
		fmt.Println(err)
	}

	// 解析 host
	h, err := url.Parse(data.EndActor)
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	// h.Hostname 类似 disism.com
	req.Header.Add("Host", h.Hostname())
	req.Header.Add("Date", date)

	r, err := account.GetActorClient(data.Name)
	if err != nil {
		log.Println(err)
	}

	block := PrivateKey{
		Key: []byte(r.PrivateKey),
	}
	SignRequest(data.UserAddress, block, req, data.Data)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()

	return res.StatusCode
}

