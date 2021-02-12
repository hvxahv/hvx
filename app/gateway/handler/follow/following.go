package follow

import (
	"github.com/gin-gonic/gin"
	"hvxahv/pkg/activity"
	"hvxahv/pkg/utils"
)

// GetFollowingHandler ... 获取正在关注的用户
func GetFollowingHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	r := activity.GetFollow(name, "following")
	c.JSON(200, gin.H{
		"res": r,
		"count": len(r),
	})
}
