package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
	"io/ioutil"
)

func InboxHandler(c *gin.Context) {
	name := c.Param("actor")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	fmt.Println(name, string(body))
	//
	//d, err := json.Marshal(push.NewData(
	//	"Notify",
	//	fmt.Sprintf("YOU HAVE A NEW FOLLOWER."),
	//	"http://49.233.26.52:9000/avatar/7b77d2b8-eccf-4886-ac0f-1806ab747261-%E5%B0%8F%E6%9E%97%E7%94%B1%E4%BE%9D.jpg",
	//	"INBOX"),
	//)
	//
	//if err := notify.NewPush(727491255195172865, 727491502310457345, d).Push(); err != nil {
	//	log.Println(err)
	//	return
	//}

	//activity.Types(name, body)
	//activity.NewActivities(name, body).Handler()
}

func OutboxHandler(c *gin.Context) {
	name := middleware.GetUsername(c)
	t := c.PostForm("type")
	o := c.PostForm("object")

	fmt.Println(name, t, o)
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}
