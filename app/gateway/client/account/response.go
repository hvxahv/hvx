package account

import (
	"github.com/gin-gonic/gin"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/pkg/bot"
)

// AccountsHandlerResponse 处理 Handler 返回的状态
func AccountsHandlerResponse(r *pb.NewAccountReply, c *gin.Context) {
	switch {
	case r.Reply == 202:
		c.JSON(202, gin.H{
			"state": "202",
			"message": "用户已存在",
		})
	case r.Reply == 200:
		c.JSON(200, gin.H{
			"state": "200",
			"message": "注册成功",
		})
		go bot.NewAccountNotice("新增加了一个用户")
	case r.Reply == 500:
		c.JSON(500, gin.H{
			"state": "500",
			"message": "注册失败",
		})
	default:

	}
}