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
	"strconv"
)
// AcceptHandler ...
func AcceptHandler(in *models.Accept) int {
	domain := viper.GetString("activitypub")

	idr := strconv.Itoa(rand.Int())
	uad := fmt.Sprintf("https://%s/u/%s", domain, in.Name)
	randId := fmt.Sprintf("https://%s/%s", domain, idr)

	obj := map[string]string {
		"id": in.RequestId,
		"type": "Follow",
		"actor": in.Actor,
		"object": uad,
	}
	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"id": randId,
		"type": "Accept",
		"actor": uad,
		"object": obj,
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	eib := fmt.Sprintf("%s/inbox", in.Actor)
	method := "POST"

	sa := *models.NewSendActivity(data, eib, method, in.Name, uad, in.Actor)
	r := activity.SendActivity(&sa)
	return r

}
