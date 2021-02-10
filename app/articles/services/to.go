package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hvxahv/pkg/activity"
	"hvxahv/pkg/models"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// SendActivity 将活动组合成 json 发送给用户的关注者
func SendActivity(id uint, author string, content string) {
	addr := viper.GetString("activitypub")
	idr := strconv.Itoa(rand.Int())
	date := time.Now().UTC().Format(http.TimeFormat)


	articleId := fmt.Sprintf("https://%s/u/%s/%s", addr, author, string(id))
	authorUrl := fmt.Sprintf("https://%s/u/%s", addr, author)
	activityId := fmt.Sprintf("https://%s/u/%s/%s", addr, author, idr)

	to := []string{"https://mas.to/users/hvturingga", "https://mstdn.social/users/hvturingga"}
	cc := []string{"https://www.w3.org/ns/activitystreams#Public", "https://mstdn.social/users/hvturingga"}

	obj := gin.H{
		"id": articleId,
		"type": "Note",
		"attributedTo": authorUrl,
		"content": content,
		"published": date,
		"to": to,
		"cc": cc,
	}
	hd := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Create",
		"id": activityId,
		"actor": authorUrl,
		"object": obj,
		"published": date,
		"to": to,
		"cc": cc,

	}

	data, err := json.Marshal(hd)
	if err != nil {
		log.Println(err)
	}

	for _, i := range to {
		eib := fmt.Sprintf("%s/inbox", i)
		method := "POST"

		sa := *models.NewSendActivity(data, eib, method, author, authorUrl, i)
		r := activity.SendActivity(&sa)
		log.Println("-------------发送创建事件到远程服务器--------------->", r)
	}

}


