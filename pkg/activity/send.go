package activity

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"hvxahv/api/cli/account"
	"hvxahv/pkg/models"
	"log"
	"net/http"
	"net/url"
	"time"
)

// TODO ........
// SendActivity 发送活动
func SendActivity(data *models.SendActivity) {
	payload := bytes.NewBuffer(data.Data)
	cli := &http.Client {
	}
	fmt.Println(payload)
	req, err := http.NewRequest(data.Method, data.UserAddress, payload)

	if err != nil {
		fmt.Println(err)
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	// 解析 Url 获取 hostname
	h, err := url.Parse(data.EndActor)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Host", h.Hostname())
	req.Header.Add("Date", date)

	r, err := account.GetActorClient(data.Name)
	if err != nil {
		log.Println(err)
	}

	block := PrivateKey{
		Key: []byte(r.PrivateKey),
	}
	SignRequest(data.EndActor, block, req, data.Data)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := cli.Do(req)
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
	log.Printf("successful post: %s %d", data.EndInbox, res.StatusCode)
}

