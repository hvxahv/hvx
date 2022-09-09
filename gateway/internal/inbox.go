package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func InboxHandler(c *gin.Context) {
	fmt.Println(c.Request.Header)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(502, gin.H{
			"status": err.Error(),
		})
	}
	fmt.Println(string(body))

	//name := c.Param("actor")
	//var body []byte
	//body, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	c.JSON(502, gin.H{
	//		"status": err.Error(),
	//	})
	//}
	//client, err := clientv1.New(c, microsvc.NewGRPCAddress("activity").Get())
	//if err != nil {
	//	c.JSON(502, gin.H{
	//		"status": err.Error(),
	//	})
	//	return
	//}
	//defer client.Close()
	//
	//inbox, err := activity.NewActivityClient(client.Conn).Inbox(c, &activity.InboxRequest{
	//	Name: name,
	//	Data: body,
	//})
	//if err != nil {
	//	c.JSON(502, gin.H{
	//		"status": err.Error(),
	//	})
	//	return
	//}
	//c.JSON(200, inbox)

	//if err := proxy.NewProxy(c, "/u/"+c.Param("actor")+"/inbox", address.Public).Proxy(); err != nil {
	//	c.JSON(502, gin.H{
	//		"error": "502_BAD_GATEWAY",
	//	})
	//	return
	//}
}
