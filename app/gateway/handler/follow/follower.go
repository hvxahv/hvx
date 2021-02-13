package follow

import (
	"github.com/gin-gonic/gin"
	"hvxahv/pkg/activitypub"
	"hvxahv/pkg/utils"
)

// GetFollowerHandler 获取关注者
func GetFollowerHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	r := activitypub.GetFollow(name, "follower")
	c.JSON(200, gin.H{
		"res": r,
		"count": len(r),
	})
}
