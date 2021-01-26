package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/response"
	"hvxahv/pkg/structs"
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
		"id": fmt.Sprintf("https://%s/actor#main-key", address),
		"owner": fmt.Sprintf("https://%s/actor/%s", address, name),
		"publicKeyPem": r.PublicKey,
	}

	c.JSON(200, gin.H{
		"@context": con,
		"id": fmt.Sprintf("https://%s/u/%s", address, name),
		"type": "Person",
		"preferredUsername": r.Username,
		"inbox": fmt.Sprintf("https://%s/u/%s/inbox", address, name),
		"outbox": fmt.Sprintf("https://%s/u/%s/outbox", address, name),
		"publicKey": publicKey,
	})
}

// WebFingerResponse 它是 Activitypub 协议的 webfinger 的 JSON-LD 标准数据返回
func WebFingerResponse(c *gin.Context, r *pb.AccountData) {
	address := viper.GetString("activitypub")
	name := r.Username

	links := []structs.WebFingerLinks{
		{
			Rel: "self",
			Type: "application/activity+json",
			Href: fmt.Sprintf("https://%s/u/%s", address, name),
		},
	}
	finger := &structs.WebFinger{
		Subject: fmt.Sprintf("acct:%s@%s", name, address),
		Links: links,
	}
	c.JSON(200, finger)
}