package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hvxahv/pkg/activitypub"
	"math/rand"
	"net/url"

	"log"
	"strconv"
)

// FollowHandler 请求关注的 Handler
func FollowHandler(actor string, name string) int {
	domain := viper.GetString("activitypub")

	idr := strconv.Itoa(rand.Int())
	uad := fmt.Sprintf("server://%s/u/%s", domain, name)

	p := gin.H{
		"@context": "server://www.w3.org/ns/activitystreams",
		"id": fmt.Sprintf("server://%s/%s", domain, idr),
		"type": "Follow",
		"actor": uad, // 当前用户的 Actor 地址
		"object": actor, // 对方的 Actor 地址, 类似 https://mas.to/users/hvturingga
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}
	h, err := url.Parse(actor)
	if err != nil {
		log.Fatal(err)
	}

	eib := fmt.Sprintf("server://%s/inbox", h.Hostname())
	method := "POST"

	sa := *activitypub.NewSendActivity(data, eib, method, name, uad, actor)
	r := activitypub.SendActivity(&sa)

	return r
}
