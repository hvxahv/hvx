package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hvxahv/app/activitypub/app"
)

func main() {
	r := gin.Default()
	r.GET("/actor", func(c *gin.Context) {
		pem := `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzcU2iJkBv6GpgU02t9/6
w1ZrKVhlw/iBLN5RmsC7K1jVLEPGCc0XWGg26ab8ZomKQfq2DpjS25L1aBjfxkWW
nJ13PnLbCTkoMiVGinEFpwuDcONN93YwZuPTqzrwZCbPc5L2owlI30AVXfl3zacm
SgwXHISyJ3Lri5RUbUv1EMMYcRxwPxPiD1AFZFTfCoB32Tv8vr4zqUJNeh+rnU8B
VZfgwLHWwjWazVxTKdirQz8HoAraARs4pxpI0Y74FSnMANfnOuwevstV+14hFtOb
WuCJurVuZkBga9pm2gsFCAdyifaNoXZMojsKSh0+d6ENuWx0247PxbQjpSNZ+atv
AwIDAQAB
-----END PUBLIC KEY-----`
		con := []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"}
		publicKey := map[string]string{
			"id": "https://a58e36568ae3.ngrok.io/actor#main-key",
			"owner": fmt.Sprintf("https://%s/actor", address),
			"publicKeyPem": pem,
		}

		c.JSON(200, gin.H{
			"@context": con,
			"id": fmt.Sprintf("https://%s/actor", address),
			"type": "Person",
			"preferredUsername": "hvturingga",
			"inbox": fmt.Sprintf("https://%s/actor/inbox", address),
			"publicKey": publicKey,
		})
	})

	r.GET("/.well-known/webfinger", func(c *gin.Context) {

		links := map[string]string{
			"rel": "self",
			"type": "application/activity+json",
			"href": fmt.Sprintf("https://%s/actor", address),
		}
		lls := make([]map[string]string,1)
		lls[0] = links


		c.JSON(200, gin.H{
			"subject": fmt.Sprintf( "acct:hvturingga@%s", address),
			"links": lls,
		})
	})
	r.GET("/@hvturingga/inspect", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		fmt.Println("接收到了消息cHECK")
	})
	r.POST("/users/hvturingga/inbox", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		fmt.Println("接收到了消息")
	})
	r.POST("/", func(c *gin.Context) {

		fmt.Println(c.Request.Header)
		fmt.Println("接收到了消息")
	})

	// 发送一条 HELLO 消息
	r.POST("/hello", app.HandleHello2)
	r.POST("/like", app.LikeHandler)
	r.POST("/start", app.NewOutbox)
	r.POST("/follow", app.Follow)
	r.POST("/undo", app.Undo)
	r.POST("/create", app.CreateHandler)
	// 接受同意消息
	r.POST("accept", app.AcceptHandler)
	// 获取收到的信息

	r.POST("/actor/inbox", app.Inbox)
	r.POST("/status", app.Status)
	r.Run(":8088")
}
