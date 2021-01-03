package account

import (
	"github.com/gin-gonic/gin"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/app/gateway/pkg/tools"
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
			"message": "REGISTRATION SUCCESS",
		})
		go tools.SendTGNotice("Notice...")
	case r.Reply == 500:
		c.JSON(500, gin.H{
			"state": "500",
			"message": "REGISTRATION FAILED",
		})
	default:

	}
}