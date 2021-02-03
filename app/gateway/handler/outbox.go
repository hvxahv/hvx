package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"hvxahv/api/client/account"
	httpsig "hvxahv/pkg/activitypub"
	"hvxahv/pkg/utils"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func FollowersAcceptHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	actor := c.PostForm("actor")
	id := c.PostForm("id")


	// 将 url 传过来的数据进行过滤，得到真正的用户名
	//if strings.HasPrefix(actor, "acct:") {
	//	name = name[5:]
	//}
	//ali := strings.IndexByte(name, '@')
	//if ali != -1 {
	//	name = name[:ali]
	//}
	//log.Println(name,actor, id, ali)


	SendFollower(c, name, id, actor)

}

func SendFollower(c *gin.Context, name, id, actor string) {

	idr := strconv.Itoa(rand.Int())
	url := fmt.Sprintf("%s/inbox", actor)
	method := "POST"

	// u == 当前用户
	u := fmt.Sprintf("https://%s/u/%s", viper.GetString("activitypub"),name)
	randId := fmt.Sprintf("https://%s/%s", viper.GetString("activitypub"), idr)
	obj := map[string]string {
		"id": id,
		"type": "Follow",
		"actor": actor,
		"object": u,
	}
	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"id": randId,
		"type": "Accept",
		"actor": u,
		"object": obj,
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}


	payload := bytes.NewBuffer(data)
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

	r, err := account.GetActorClient(name)
	if err != nil {
		log.Println(err)
	}

	block := httpsig.PrivateKey{
		Key: []byte(r.PrivateKey),
	}
	httpsig.SignRequest(u, block, req, data)
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
	c.JSON(200, gin.H{
		"status": "200",
		"message": "ok",
	})
}

