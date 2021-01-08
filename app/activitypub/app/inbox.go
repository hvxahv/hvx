package app

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Followers struct {
	Context string `json:"@context"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Actor  string `json:"actor"`
		Object string `json:"object"`
	} `json:"object"`
}

func Inbox(c *gin.Context) {
	//f := make(map[string]interface{})
	//if err := c.BindJSON(&f); err != nil {
	//	log.Println(err)
	//}
	//log.Println(f["type"])
	//switch f["type"] {
	//case "Follow":
	//	log.Printf("%s关注了你", f["actor"])
	//case "Undo":
	//	log.Printf("%s取消关注你了，废物", f["actor"])
	//}
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))

}
