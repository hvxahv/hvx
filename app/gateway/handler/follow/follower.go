package follow

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/server/middleware"
	"hvxahv/pkg/activitypub"
)

// GetFollowerHandler 获取关注者
func GetFollowerHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	r := activitypub.GetFollow(name, "follower")
	c.JSON(200, gin.H{
		"res": r,
		"count": len(r),
	})
}
