package follow

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/server/middleware"
	"hvxahv/pkg/activitypub"
)

// GetFollowingHandler ... 获取正在关注的用户
func GetFollowingHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	r := activitypub.GetFollow(name, "following")
	c.JSON(200, gin.H{
		"res":   r,
		"count": len(r),
	})
}
