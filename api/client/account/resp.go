package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/db"
	"hvxahv/pkg/models"
	"hvxahv/pkg/response"
	"log"
)

// NewAccountsResponse 创建用户的返回值处理, 它接收 Accounts 的服务返回的状态码
// 将返回的状态码进行处理并将相应返回
func NewAccountsResponse(c *gin.Context, r *pb.NewAccountReply) {
	switch {
	case r.Reply == 200:
		response.SimpleResponse(c, "200", "注册成功")
		// 注册成功后，将注册信息发送给 BOT
		go bot.NewAccountNotice("新增加了一个用户")
	case r.Reply == 202:
		response.SimpleResponse(c, "202", "用户已存在")
	case r.Reply == 500:
		response.SimpleResponse(c, "500", "注册失败")
	default:

	}
}

// AccountsResponse 返回用户获取到的它的账户的信息
func AccountsResponse(c *gin.Context, r *pb.AccountData) {
	c.JSON(200, gin.H{
		"name": r.Name,
	})
}

// ActorResponse 它是 Activitypub 协议的 Actor 的 JSON-LD 标准数据返回
func ActorResponse(c *gin.Context, r *pb.AccountData) {
	name := r.Username
	address := viper.GetString("activitypub")

	con := []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"}
	publicKey := map[string]string{
		"id": r.Id,
		"owner": fmt.Sprintf("https://%s/actor/%s", address, name),
		"publicKeyPem": r.PublicKey,
	}

	c.JSON(200, gin.H{
		"@context":          con,
		"type":              "Person",
		"id":                fmt.Sprintf("https://%s/u/%s", address, name),
		"following":         formatLink("following", name),
		"followers":         formatLink("followers", name),
		"preferredUsername": r.Username,
		"name":              r.Name,
		"inbox":             formatLink("inbox", name),
		"outbox":            formatLink("outbox", name),
		"publicKey":         publicKey,
		"icon":              "https://i.mydramalist.com/EpDnpc.jpg",
	})
}
func formatLink(route, name string) string {
	address := viper.GetString("activitypub")
	return fmt.Sprintf("https://%s/u/%s/%s", address, name, route)
}
// WebFingerResponse 它是 Activitypub 协议的 webfinger 的 JSON-LD 标准数据返回
func WebFingerResponse(c *gin.Context, r *pb.AccountData) {
	address := viper.GetString("activitypub")
	name := r.Username

	links := []models.WebFingerLinks{
		{
			Rel: "self",
			Type: "application/activity+json",
			Href: fmt.Sprintf("https://%s/u/%s", address, name),
		},
	}
	finger := &models.WebFinger{
		Subject: fmt.Sprintf("acct:%s@%s", name, address),
		Links: links,
	}
	log.Println(finger)
	c.JSON(200, finger)
}


type Create struct {
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object []*Object
}
type Object struct {
	Type         string `json:"type"`
	AttributedTo string `json:"attributedTo,omitempty"`
	InReplyTo    string `json:"inReplyTo"`
	Content      string `json:"content,omitempty"`
	To           string `json:"to,omitempty"`
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
}
func OutboxResponse(c *gin.Context) {

	obj := &Object{
		Type: "Note",
		AttributedTo: "https://dc3a16810ea3.ngrok.io/u/hvturingga",
		InReplyTo: "https://dc3a16810ea3.ngrok.io/u/hvturingga/outbox/7ca154ff",
		Content: "<p>Hello world</p>",
		To:  "https://www.w3.org/ns/activitystreams#Public",
	}
	obj2 := &Object{
		Type: "Note",
		AttributedTo: "https://littr.git/api/accounts/anonymous",
		InReplyTo: "https://dc3a16810ea3.ngrok.io/u/hvturingga/outbox/7ca154ff",
		Content: "<p>Hello wdsadsadasdavdsvdwsdvcdorld</p>",
		To:  "https://www.w3.org/ns/activitystreams#Public",
	}
	xxx := []*Object{obj, obj2}

	c.JSON(200, gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "Sally's notes",
		"type": "OrderedCollection",
		"totalItems": 4,
		"orderedItems": xxx,
	})
}


func FollowersResponse(c *gin.Context) {
	name := c.Param("user")
	rdb := db.GetRDB()
	res, err := redis.Int(rdb.Do("GET", fmt.Sprintf("%s-following", name)))
	if err != nil {
		log.Println("Redis 获取 Actor 数据失败:", err)
	}

	c.JSON(200, gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "Sally followed John",
		"type": "OrderedCollection",
		"totalItems": res,
		"orderedItems": "",
	})

}

func FollowingResponse(c *gin.Context) {

}